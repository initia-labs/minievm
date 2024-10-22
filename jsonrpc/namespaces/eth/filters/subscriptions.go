package filters

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	ethfilters "github.com/ethereum/go-ethereum/eth/filters"
	"github.com/ethereum/go-ethereum/rpc"

	rpctypes "github.com/initia-labs/minievm/jsonrpc/types"
)

type subscription struct {
	ty     ethfilters.Type
	crit   ethfilters.FilterCriteria
	fullTx bool

	headerChan chan *coretypes.Header
	logsChan   chan []*coretypes.Log
	txChan     chan *rpctypes.RPCTransaction
	hashChan   chan common.Hash
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
	api.filtersMut.Lock()
	api.subscriptions[id] = &subscription{
		ty:         ethfilters.BlocksSubscription,
		headerChan: headerChan,
	}
	api.filtersMut.Unlock()

	go func() {
		defer api.clearSubscription(id)

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
	api.filtersMut.Lock()
	api.subscriptions[id] = &subscription{
		ty:   ethfilters.LogsSubscription,
		crit: crit,

		logsChan: logsChan,
	}
	api.filtersMut.Unlock()

	go func() {
		defer api.clearSubscription(id)
		for {
			select {
			case logs := <-logsChan:
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
	api.filtersMut.Lock()
	api.subscriptions[id] = &subscription{
		ty:       ethfilters.PendingTransactionsSubscription,
		fullTx:   fullTx != nil && *fullTx,
		txChan:   txChan,
		hashChan: hashChan,
	}
	api.filtersMut.Unlock()

	go func() {
		defer api.clearSubscription(id)

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

func (api *FilterAPI) clearSubscription(id rpc.ID) {
	api.filtersMut.Lock()
	delete(api.subscriptions, id)
	api.filtersMut.Unlock()
}
