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
	errExceedFilterCap   = errors.New("exceed filter cap")
)

// The maximum number of topic criteria allowed, vm.LOG4 - vm.LOG0
const maxTopics = 4

type filter struct {
	ty     ethfilters.Type
	hashes []common.Hash
	fullTx bool
	txs    []*rpctypes.RPCTransaction
	crit   ethfilters.FilterCriteria
	logs   []*coretypes.Log

	// lastUsed is the time the filter was last used
	lastUsed time.Time
}

// FilterAPI is the eth_ filter namespace API
type FilterAPI struct {
	app     *app.MinitiaApp
	backend *backend.JSONRPCBackend

	logger log.Logger

	filtersMut sync.Mutex

	filters       map[rpc.ID]*filter
	subscriptions map[rpc.ID]*subscription

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

		filters:       make(map[rpc.ID]*filter),
		subscriptions: make(map[rpc.ID]*subscription),
	}

	go api.clearUnusedFilters()

	api.blockChan, api.logsChan, api.pendingChan = app.EVMIndexer().Subscribe()
	go api.subscribeEvents()

	return api
}

// clearUnusedFilters removes filters that have not been used for 5 minutes
func (api *FilterAPI) clearUnusedFilters() {
	const timeout = 5 * time.Minute

	for {
		time.Sleep(timeout)
		api.filtersMut.Lock()
		for id, f := range api.filters {
			if time.Since(f.lastUsed) > 5*time.Minute {
				delete(api.filters, id)
			}
		}
		api.filtersMut.Unlock()
	}
}

func (api *FilterAPI) subscribeEvents() {
	for {
		select {
		case block := <-api.blockChan:
			api.filtersMut.Lock()
			for _, f := range api.filters {
				if f.ty == ethfilters.BlocksSubscription {
					f.hashes = append(f.hashes, block.Hash())
				}
			}
			for _, s := range api.subscriptions {
				if s.ty == ethfilters.BlocksSubscription {
					s.headerChan <- block
				}
			}
			api.filtersMut.Unlock()
		case logs := <-api.logsChan:
			if len(logs) == 0 {
				continue
			}

			api.filtersMut.Lock()
			for _, f := range api.filters {
				if f.ty == ethfilters.LogsSubscription {
					logs := filterLogs(logs, f.crit.FromBlock, f.crit.ToBlock, f.crit.Addresses, f.crit.Topics)
					if len(logs) > 0 {
						f.logs = append(f.logs, logs...)
					}
				}
			}
			for _, s := range api.subscriptions {
				if s.ty == ethfilters.LogsSubscription {
					logs := filterLogs(logs, s.crit.FromBlock, s.crit.ToBlock, s.crit.Addresses, s.crit.Topics)
					if len(logs) > 0 {
						s.logsChan <- logs
					}
				}
			}
			api.filtersMut.Unlock()
		case tx := <-api.pendingChan:
			api.filtersMut.Lock()
			for _, f := range api.filters {
				if f.ty == ethfilters.PendingTransactionsSubscription {
					if f.fullTx {
						f.txs = append(f.txs, tx)
					} else {
						f.hashes = append(f.hashes, tx.Hash)
					}
				}
			}
			for _, s := range api.subscriptions {
				if s.ty == ethfilters.PendingTransactionsSubscription {
					if s.fullTx {
						s.txChan <- tx
					} else {
						s.hashChan <- tx.Hash
					}
				}
			}
			api.filtersMut.Unlock()
		}
	}
}

// NewPendingTransactionFilter creates a filter that fetches pending transactions
// as transactions enter the pending state.
//
// It is part of the filter package because this filter can be used through the
// `eth_getFilterChanges` polling method that is also used for log filters.
func (api *FilterAPI) NewPendingTransactionFilter(fullTx *bool) (rpc.ID, error) {
	if len(api.filters) >= int(api.backend.RPCFilterCap()) {
		return "", errExceedFilterCap
	}

	id := rpc.NewID()
	api.filtersMut.Lock()
	api.filters[id] = &filter{
		ty:     ethfilters.PendingTransactionsSubscription,
		fullTx: fullTx != nil && *fullTx,
		txs:    make([]*rpctypes.RPCTransaction, 0),
		hashes: make([]common.Hash, 0),
	}
	api.filtersMut.Unlock()

	return id, nil
}

// NewBlockFilter creates a filter that fetches blocks that are imported into the chain.
// It is part of the filter package since polling goes with eth_getFilterChanges.
func (api *FilterAPI) NewBlockFilter() (rpc.ID, error) {
	if len(api.filters) >= int(api.backend.RPCFilterCap()) {
		return "", errExceedFilterCap
	}

	id := rpc.NewID()
	api.filtersMut.Lock()
	api.filters[id] = &filter{
		ty:     ethfilters.BlocksSubscription,
		hashes: make([]common.Hash, 0),
	}
	api.filtersMut.Unlock()

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
	if len(api.filters) >= int(api.backend.RPCFilterCap()) {
		return "", errExceedFilterCap
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

	id := rpc.NewID()
	api.filtersMut.Lock()
	api.filters[id] = &filter{
		ty:   ethfilters.LogsSubscription,
		crit: crit, lastUsed: time.Now(),
		logs: make([]*coretypes.Log, 0),
	}
	api.filtersMut.Unlock()

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
	logs, err := filter.Logs(ctx, int64(api.backend.RPCBlockRangeCap()))
	if err != nil {
		return nil, err
	}
	return returnLogs(logs), err
}

// UninstallFilter removes the filter with the given filter id.
func (api *FilterAPI) UninstallFilter(id rpc.ID) bool {
	api.filtersMut.Lock()
	_, found := api.filters[id]
	delete(api.filters, id)
	api.filtersMut.Unlock()
	return found
}

// GetFilterLogs returns the logs for the filter with the given id.
// If the filter could not be found an empty array of logs is returned.
func (api *FilterAPI) GetFilterLogs(ctx context.Context, id rpc.ID) ([]*coretypes.Log, error) {
	api.filtersMut.Lock()
	f, found := api.filters[id]
	api.filtersMut.Lock()

	if !found || f.ty != ethfilters.LogsSubscription {
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
	logs, err := bloomFilter.Logs(ctx, int64(api.backend.RPCBlockRangeCap()))
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

	switch f.ty {
	case ethfilters.BlocksSubscription:
		hashes := f.hashes
		f.hashes = nil

		return returnHashes(hashes), nil
	case ethfilters.LogsSubscription:
		logs := f.logs
		f.logs = nil

		return returnLogs(logs), nil
	case ethfilters.PendingTransactionsSubscription:
		if f.fullTx {
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
