package filters

import (
	"context"
	"errors"
	"sync"
	"time"

	"cosmossdk.io/log"

	"github.com/initia-labs/minievm/app"
	"github.com/initia-labs/minievm/jsonrpc/backend"
	rpctypes "github.com/initia-labs/minievm/jsonrpc/types"

	"github.com/ethereum/go-ethereum/common"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	ethfilters "github.com/ethereum/go-ethereum/eth/filters"
	"github.com/ethereum/go-ethereum/rpc"
)

var (
	errFilterNotFound    = errors.New("filter not found")
	errInvalidBlockRange = errors.New("invalid block range params")
	errExceedMaxTopics   = errors.New("exceed max topics")
)

// The maximum number of topic criteria allowed, vm.LOG4 - vm.LOG0
const maxTopics = 4

type filter struct {
	hashes []common.Hash
	txs    []*rpctypes.RPCTransaction
	crit   ethfilters.FilterCriteria
	logs   []*coretypes.Log

	// lastUsed is the time the filter was last used
	lastUsed time.Time

	// subscription lifecycle
	s *subscription
}

// FilterAPI is the eth_ filter namespace API
type FilterAPI struct {
	app     *app.MinitiaApp
	backend *backend.JSONRPCBackend

	logger log.Logger

	filtersMut sync.Mutex

	filters       map[rpc.ID]*filter
	subscriptions map[rpc.ID]*subscription

	// Channels for subscription managements
	install   chan *subscription // install filter for event notification
	uninstall chan *subscription // remove filter for event notification

	// channels for block and log events
	blockChan   chan *coretypes.Header
	logsChan    chan []*coretypes.Log
	pendingChan chan *rpctypes.RPCTransaction
}

// NewFiltersAPI returns a new instance
func NewFilterAPI(app *app.MinitiaApp, backend *backend.JSONRPCBackend, logger log.Logger) *FilterAPI {
	logger = logger.With("api", "filter")
	api := &FilterAPI{
		app:     app,
		backend: backend,

		logger: logger,

		install:   make(chan *subscription),
		uninstall: make(chan *subscription),

		filters:       make(map[rpc.ID]*filter),
		subscriptions: make(map[rpc.ID]*subscription),
	}

	go api.clearUnusedFilters()

	api.blockChan, api.logsChan, api.pendingChan = app.EVMIndexer().Subscribe()
	go api.eventLoop()

	return api
}

// clearUnusedFilters removes filters that have not been used for 5 minutes
func (api *FilterAPI) clearUnusedFilters() {
	timeout := api.backend.FilterTimeout()
	if timeout == 0 {
		timeout = 5 * time.Minute
	}

	ticker := time.NewTicker(timeout)
	defer ticker.Stop()

	var toUninstall []*subscription
	for {
		<-ticker.C
		api.filtersMut.Lock()
		for id, f := range api.filters {
			if time.Since(f.lastUsed) > timeout {
				toUninstall = append(toUninstall, f.s)
				delete(api.filters, id)
			}
		}
		api.filtersMut.Unlock()

		// Unsubscribes are processed outside the lock to avoid the following scenario:
		// event loop attempts broadcasting events to still active filters while
		// Unsubscribe is waiting for it to process the uninstall request.
		for _, s := range toUninstall {
			api.uninstallSubscription(s)
		}
		toUninstall = nil
	}
}

func (api *FilterAPI) eventLoop() {
	for {
		select {
		case block := <-api.blockChan:
			for _, s := range api.subscriptions {
				if s.ty == ethfilters.BlocksSubscription {
					s.headerChan <- block
				}
			}
		case logs := <-api.logsChan:
			if len(logs) == 0 {
				continue
			}

			for _, s := range api.subscriptions {
				if s.ty == ethfilters.LogsSubscription {
					// logs will be filtered in the subscription in the goroutine
					s.logsChan <- logs
				}
			}
		case tx := <-api.pendingChan:
			for _, s := range api.subscriptions {
				if s.ty == ethfilters.PendingTransactionsSubscription {
					if s.fullTx {
						s.txChan <- tx
					} else {
						s.hashChan <- tx.Hash
					}
				}
			}
		// subscription managements
		case s := <-api.install:
			api.subscriptions[s.id] = s
			close(s.installed)
		case s := <-api.uninstall:
			delete(api.subscriptions, s.id)
			close(s.err)
		}
	}
}

// NewPendingTransactionFilter creates a filter that fetches pending transactions
// as transactions enter the pending state.
//
// It is part of the filter package because this filter can be used through the
// `eth_getFilterChanges` polling method that is also used for log filters.
func (api *FilterAPI) NewPendingTransactionFilter(fullTx *bool) (rpc.ID, error) {
	var (
		txChan   = make(chan *rpctypes.RPCTransaction)
		hashChan = make(chan common.Hash)
	)

	id := rpc.NewID()
	s := &subscription{
		id:     id,
		ty:     ethfilters.PendingTransactionsSubscription,
		fullTx: fullTx != nil && *fullTx,

		// for listening
		txChan:   txChan,
		hashChan: hashChan,

		// for lifecycle
		installed: make(chan struct{}),
		err:       make(chan error),
	}
	api.installSubscription(s)

	api.filtersMut.Lock()
	api.filters[id] = &filter{
		txs:    make([]*rpctypes.RPCTransaction, 0),
		hashes: make([]common.Hash, 0),
		s:      s,
	}
	api.filtersMut.Unlock()

	go func() {
		for {
			select {
			case rpcTx := <-txChan:
				api.filtersMut.Lock()
				if f, found := api.filters[id]; found {
					f.txs = append(f.txs, rpcTx)
				}
				api.filtersMut.Unlock()
			case hash := <-hashChan:
				api.filtersMut.Lock()
				if f, found := api.filters[id]; found {
					f.hashes = append(f.hashes, hash)
				}
				api.filtersMut.Unlock()
			case <-s.err: // subsciprtion is uninstalled
				return
			}
		}
	}()

	return id, nil
}

// NewBlockFilter creates a filter that fetches blocks that are imported into the chain.
// It is part of the filter package since polling goes with eth_getFilterChanges.
func (api *FilterAPI) NewBlockFilter() (rpc.ID, error) {
	var (
		headerChan = make(chan *coretypes.Header)
	)

	id := rpc.NewID()
	s := &subscription{
		id: id,
		ty: ethfilters.BlocksSubscription,

		// for listening
		headerChan: headerChan,

		// for lifecycle
		installed: make(chan struct{}),
		err:       make(chan error),
	}
	api.installSubscription(s)

	api.filtersMut.Lock()
	api.filters[id] = &filter{
		hashes: make([]common.Hash, 0),
		s:      s,
	}
	api.filtersMut.Unlock()

	go func() {
		for {
			select {
			case header := <-headerChan:
				api.filtersMut.Lock()
				if f, found := api.filters[id]; found {
					f.hashes = append(f.hashes, header.Hash())
				}
				api.filtersMut.Unlock()
			case <-s.err: // subsciprtion is uninstalled
				return
			}
		}
	}()

	return id, nil
}

// NewFilter creates a new filter and returns the filter id. It can be
// used to retrieve logs when the state changes. This method cannot be
// used to fetch logs that are already stored in the state.
//
// Default criteria for the from and to block are "latest".
// Using "latest" as block number will return logs for mined blocks.
//
// In case "fromBlock" > "toBlock" an error is returned.
func (api *FilterAPI) NewFilter(crit ethfilters.FilterCriteria) (rpc.ID, error) {
	if len(crit.Topics) > maxTopics {
		return "", errExceedMaxTopics
	}

	var from, to rpc.BlockNumber
	if crit.FromBlock == nil {
		from = rpc.LatestBlockNumber
	} else {
		from = rpc.BlockNumber(crit.FromBlock.Int64())
	}
	if crit.ToBlock == nil {
		to = rpc.LatestBlockNumber
	} else {
		to = rpc.BlockNumber(crit.ToBlock.Int64())
	}

	// we don't support pending logs
	if !(from == rpc.LatestBlockNumber && to == rpc.LatestBlockNumber) &&
		!(from >= 0 && to >= 0 && to >= from) &&
		!(from >= 0 && to == rpc.LatestBlockNumber) {
		return "", errInvalidBlockRange
	}

	var (
		logsChan = make(chan []*coretypes.Log)
	)

	id := rpc.NewID()
	s := &subscription{
		id:   id,
		ty:   ethfilters.LogsSubscription,
		crit: crit,

		// for listening
		logsChan: logsChan,

		// for lifecycle
		installed: make(chan struct{}),
		err:       make(chan error),
	}
	api.installSubscription(s)

	api.filtersMut.Lock()
	api.filters[id] = &filter{
		crit: crit, lastUsed: time.Now(),
		logs: make([]*coretypes.Log, 0),
		s:    s,
	}
	api.filtersMut.Unlock()

	go func() {
		for {
			select {
			case logs := <-logsChan:
				logs = filterLogs(logs, s.crit.FromBlock, s.crit.ToBlock, s.crit.Addresses, s.crit.Topics)
				api.filtersMut.Lock()
				if f, found := api.filters[id]; found {
					f.logs = append(f.logs, logs...)
				}
				api.filtersMut.Unlock()
			case <-s.err: // subsciprtion is uninstalled
				return
			}
		}
	}()

	return id, nil
}

// GetLogs returns logs matching the given argument that are stored within the state.
func (api *FilterAPI) GetLogs(ctx context.Context, crit ethfilters.FilterCriteria) ([]*coretypes.Log, error) {
	if len(crit.Topics) > maxTopics {
		return nil, errExceedMaxTopics
	}
	var filter *Filter
	if crit.BlockHash != nil {
		// Block filter requested, construct a single-shot filter
		filter = newBlockFilter(api.logger, api.backend, *crit.BlockHash, crit.Addresses, crit.Topics)
	} else {
		// Convert the RPC block numbers into internal representations
		begin := rpc.LatestBlockNumber.Int64()
		if crit.FromBlock != nil {
			begin = crit.FromBlock.Int64()
		}
		end := rpc.LatestBlockNumber.Int64()
		if crit.ToBlock != nil {
			end = crit.ToBlock.Int64()
		}
		if begin > 0 && end > 0 && begin > end {
			return nil, errInvalidBlockRange
		}
		// Construct the range filter
		filter = newRangeFilter(api.logger, api.backend, begin, end, crit.Addresses, crit.Topics)
	}

	// Run the filter and return all the logs
	logs, err := filter.Logs(ctx)
	if err != nil {
		return nil, err
	}
	return returnLogs(logs), err
}

// UninstallFilter removes the filter with the given filter id.
func (api *FilterAPI) UninstallFilter(id rpc.ID) bool {
	api.filtersMut.Lock()
	f, found := api.filters[id]
	if found {
		delete(api.filters, id)
	}
	api.filtersMut.Unlock()
	if found {
		api.uninstallSubscription(f.s)
	}

	return found
}

// GetFilterLogs returns the logs for the filter with the given id.
// If the filter could not be found an empty array of logs is returned.
func (api *FilterAPI) GetFilterLogs(ctx context.Context, id rpc.ID) ([]*coretypes.Log, error) {
	api.filtersMut.Lock()
	f, found := api.filters[id]
	api.filtersMut.Unlock()

	if !found || f.s.ty != ethfilters.LogsSubscription {
		return nil, errFilterNotFound
	}

	var bloomFilter *Filter
	if f.crit.BlockHash != nil {
		// Block filter requested, construct a single-shot filter
		bloomFilter = newBlockFilter(api.logger, api.backend, *f.crit.BlockHash, f.crit.Addresses, f.crit.Topics)
	} else {
		// Convert the RPC block numbers into internal representations
		begin := rpc.LatestBlockNumber.Int64()
		if f.crit.FromBlock != nil {
			begin = f.crit.FromBlock.Int64()
		}
		end := rpc.LatestBlockNumber.Int64()
		if f.crit.ToBlock != nil {
			end = f.crit.ToBlock.Int64()
		}
		// Construct the range filter
		bloomFilter = newRangeFilter(api.logger, api.backend, begin, end, f.crit.Addresses, f.crit.Topics)
	}

	// Run the filter and return all the logs
	logs, err := bloomFilter.Logs(ctx)
	if err != nil {
		return nil, err
	}

	return returnLogs(logs), nil
}

// GetFilterChanges returns the logs for the filter with the given id since
// last time it was called. This can be used for polling.
//
// For pending transaction and block filters the result is []common.Hash.
// (pending)Log filters return []Log.
func (api *FilterAPI) GetFilterChanges(id rpc.ID) (interface{}, error) {
	api.filtersMut.Lock()
	defer api.filtersMut.Unlock()

	f, ok := api.filters[id]
	if !ok {
		return []interface{}{}, errFilterNotFound
	}

	f.lastUsed = time.Now()

	switch f.s.ty {
	case ethfilters.BlocksSubscription:
		hashes := f.hashes
		f.hashes = nil

		return returnHashes(hashes), nil
	case ethfilters.LogsSubscription:
		logs := f.logs
		f.logs = nil

		return returnLogs(logs), nil
	case ethfilters.PendingTransactionsSubscription:
		if f.s.fullTx {
			txs := f.txs
			f.txs = nil

			return txs, nil
		}

		hashes := f.hashes
		f.hashes = nil

		return returnHashes(hashes), nil
	}

	return []interface{}{}, errFilterNotFound
}

// returnLogs is a helper that will return an empty log array in case the given logs array is nil,
// otherwise the given logs array is returned.
func returnLogs(logs []*coretypes.Log) []*coretypes.Log {
	if logs == nil {
		return []*coretypes.Log{}
	}
	return logs
}

// returnHashes is a helper that will return an empty hash array case the given hash array is nil,
// otherwise the given hashes array is returned.
func returnHashes(hashes []common.Hash) []common.Hash {
	if hashes == nil {
		return []common.Hash{}
	}
	return hashes
}
