package filters

import (
	"context"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	ethfilters "github.com/ethereum/go-ethereum/eth/filters"
	"github.com/ethereum/go-ethereum/rpc"

	rpctypes "github.com/initia-labs/minievm/jsonrpc/types"
)

type subscription struct {
	id     rpc.ID
	ty     ethfilters.Type
	crit   ethfilters.FilterCriteria
	fullTx bool

	// for listening
	headerChan chan *coretypes.Header
	logsChan   chan []*coretypes.Log
	txChan     chan *rpctypes.RPCTransaction
	hashChan   chan common.Hash

	// for lifecycle
	installed chan struct{} // closed when the subscription is installed
	err       chan error    // closed when the subscription is uninstalled
	unsubOnce sync.Once
}

// NewHeads send a notification each time a new (header) block is appended to the chain.
func (api *FilterAPI) NewHeads(ctx context.Context) (*rpc.Subscription, error) {
	notifier, supported := rpc.NotifierFromContext(ctx)
	if !supported {
		return &rpc.Subscription{}, rpc.ErrNotificationsUnsupported
	}

	var (
		rpcSub     = notifier.CreateSubscription()
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

	go func() {
		defer api.uninstallSubscription(s)

		for {
			select {
			case h := <-headerChan:
				_ = notifier.Notify(rpcSub.ID, h)
			case <-rpcSub.Err():
				return
			}
		}
	}()

	return rpcSub, nil
}

func (api *FilterAPI) Logs(ctx context.Context, crit ethfilters.FilterCriteria) (*rpc.Subscription, error) {
	notifier, supported := rpc.NotifierFromContext(ctx)
	if !supported {
		return &rpc.Subscription{}, rpc.ErrNotificationsUnsupported
	}

	if len(crit.Topics) > maxTopics {
		return &rpc.Subscription{}, errExceedMaxTopics
	} else if len(crit.Addresses) > api.backend.FilterMaxAddresses() {
		return &rpc.Subscription{}, errExceedMaxAddrs
	}

	var (
		rpcSub   = notifier.CreateSubscription()
		logsChan = make(chan []*coretypes.Log)
	)

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
		return &rpc.Subscription{}, errInvalidBlockRange
	}

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

	go func() {
		defer api.uninstallSubscription(s)
		for {
			select {
			case logs := <-logsChan:
				logs = filterLogs(logs, s.crit.FromBlock, s.crit.ToBlock, s.crit.Addresses, s.crit.Topics)
				for _, log := range logs {
					log := log
					_ = notifier.Notify(rpcSub.ID, &log)
				}
			case <-rpcSub.Err(): // client send an unsubscribe request
				return
			}
		}
	}()

	return rpcSub, nil
}

// NewPendingTransactions creates a subscription that is triggered each time a
// transaction enters the transaction pool. If fullTx is true the full tx is
// sent to the client, otherwise the hash is sent.
func (api *FilterAPI) NewPendingTransactions(ctx context.Context, fullTx *bool) (*rpc.Subscription, error) {
	notifier, supported := rpc.NotifierFromContext(ctx)
	if !supported {
		return &rpc.Subscription{}, rpc.ErrNotificationsUnsupported
	}

	var (
		rpcSub   = notifier.CreateSubscription()
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

	go func() {
		defer api.uninstallSubscription(s)

		for {
			select {
			case rpcTx := <-txChan:
				_ = notifier.Notify(rpcSub.ID, rpcTx)
			case hash := <-hashChan:
				_ = notifier.Notify(rpcSub.ID, hash)
			case <-rpcSub.Err():
				return
			}
		}
	}()

	return rpcSub, nil
}

func (api *FilterAPI) installSubscription(s *subscription) {
	api.install <- s
	<-s.installed
}

func (api *FilterAPI) uninstallSubscription(s *subscription) {
	s.unsubOnce.Do(func() {
	uninstallLoop:
		for {
			// write uninstall request and consume logs/hashes. This prevents
			// the eventLoop broadcast method to deadlock when writing to the
			// filter event channel while the subscription loop is waiting for
			// this method to return (and thus not reading these events).
			select {
			case api.uninstall <- s:
				break uninstallLoop
			case <-s.logsChan:
			case <-s.txChan:
			case <-s.hashChan:
			case <-s.headerChan:
			}
		}

		// wait for filter to be uninstalled in work loop before returning
		// this ensures that the manager won't use the event channel which
		// will probably be closed by the client asap after this method returns.
		<-s.err
	})
}
