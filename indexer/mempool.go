package indexer

import (
	"github.com/ethereum/go-ethereum/common"
	coretypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/jellydator/ttlcache/v3"

	rpctypes "github.com/initia-labs/minievm/jsonrpc/types"
)

// TxInPending returns true if the transaction with the given hash is in the mempool.
func (indexer *EVMIndexerImpl) TxInPending(hash common.Hash) *rpctypes.RPCTransaction {
	item := indexer.pendingTxs.Get(hash)
	if item == nil {
		return nil
	}

	return item.Value()
}

func (indexer *EVMIndexerImpl) TxInQueued(hash common.Hash) *rpctypes.RPCTransaction {
	item := indexer.queuedTxs.Get(hash)
	if item == nil {
		return nil
	}

	return item.Value()
}

func (e *EVMIndexerImpl) PushPendingTx(tx *coretypes.Transaction) {
	rpcTx := rpctypes.NewRPCTransaction(tx, common.Hash{}, 0, 0, tx.ChainId())

	// push to pending txs
	e.pendingTxs.Set(tx.Hash(), rpcTx, ttlcache.DefaultTTL)

	// emit the transaction to all pending channels
	go func() {
		for _, pendingChan := range e.pendingChans {
			pendingChan <- rpcTx
		}
	}()
}

func (e *EVMIndexerImpl) PushQueuedTx(tx *coretypes.Transaction) {
	rpcTx := rpctypes.NewRPCTransaction(tx, common.Hash{}, 0, 0, tx.ChainId())

	// push to queued txs
	e.queuedTxs.Set(tx.Hash(), rpcTx, ttlcache.DefaultTTL)
}

func (e *EVMIndexerImpl) PendingTxs() []*rpctypes.RPCTransaction {
	items := e.pendingTxs.Items()
	result := make([]*rpctypes.RPCTransaction, 0, len(items))
	for _, item := range items {
		result = append(result, item.Value())
	}
	return result
}

func (e *EVMIndexerImpl) QueuedTxs() []*rpctypes.RPCTransaction {
	items := e.queuedTxs.Items()
	result := make([]*rpctypes.RPCTransaction, 0, len(items))
	for _, item := range items {
		result = append(result, item.Value())
	}
	return result
}

func (e *EVMIndexerImpl) NumPendingTxs() int {
	return e.pendingTxs.Len()
}

func (e *EVMIndexerImpl) NumQueuedTxs() int {
	return e.queuedTxs.Len()
}

func (e *EVMIndexerImpl) RemovePendingTx(hash common.Hash) {
	e.pendingTxs.Delete(hash)
}

func (e *EVMIndexerImpl) RemoveQueuedTx(hash common.Hash) {
	e.queuedTxs.Delete(hash)
}
