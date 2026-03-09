package indexer

import (
	"fmt"
	"sync"
	"sync/atomic"

	cmtmempool "github.com/cometbft/cometbft/mempool"
	"github.com/ethereum/go-ethereum/common"

	rpctypes "github.com/initia-labs/minievm/jsonrpc/types"
)

// MempoolTxCache is an in-memory cache that tracks pending and queued
// transactions by address+nonce and by hash, enabling fast lookups for txpool queries.
type MempoolTxCache struct {
	mu sync.RWMutex

	// address -> nonce -> tx, separated by pool
	pending map[common.Address]map[uint64]*rpctypes.RPCTransaction
	queued  map[common.Address]map[uint64]*rpctypes.RPCTransaction

	// hash -> tx for fast tx lookup by hash
	byHash map[common.Hash]*rpctypes.RPCTransaction

	// pool tracks which pool a tx hash belongs to: true=pending, false=queued
	pool map[common.Hash]bool

	// txKeyToHash tracks cosmosTxKey -> eth hash for fast removal by cosmos tx key
	txKeyToHash map[string]common.Hash

	pendingCount atomic.Int64
	queuedCount  atomic.Int64
}

// NewMempoolTxCache creates a new MempoolTxCache.
func NewMempoolTxCache() *MempoolTxCache {
	return &MempoolTxCache{
		pending:     make(map[common.Address]map[uint64]*rpctypes.RPCTransaction),
		queued:      make(map[common.Address]map[uint64]*rpctypes.RPCTransaction),
		byHash:      make(map[common.Hash]*rpctypes.RPCTransaction),
		pool:        make(map[common.Hash]bool),
		txKeyToHash: make(map[string]common.Hash),
	}
}

// HandleInserted handles an EventTxInserted. If the tx is already in the queued pool (promotion), then moves it.
// Otherwise, adds to pending.
func (c *MempoolTxCache) HandleInserted(rpcTx *rpctypes.RPCTransaction, cosmosTxKey string) {
	if rpcTx == nil {
		return
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	hash := rpcTx.Hash
	from := rpcTx.From
	nonce := uint64(rpcTx.Nonce)

	// if already in queued, promote to pending
	if isPending, exists := c.pool[hash]; exists && !isPending {
		c.removeFromQueuedLocked(from, nonce, hash)
		c.addToPendingLocked(from, nonce, hash, rpcTx, cosmosTxKey)
		return
	}

	// fresh insert to pending
	c.addToPendingLocked(from, nonce, hash, rpcTx, cosmosTxKey)
}

// AddQueued adds a transaction to the queued pool.
func (c *MempoolTxCache) AddQueued(rpcTx *rpctypes.RPCTransaction, cosmosTxKey string) {
	if rpcTx == nil {
		return
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	hash := rpcTx.Hash
	from := rpcTx.From
	nonce := uint64(rpcTx.Nonce)

	// if already tracked, remove the old entry first
	if _, exists := c.pool[hash]; exists {
		return
	}

	c.addToQueuedLocked(from, nonce, hash, rpcTx, cosmosTxKey)
}

// Remove removes a transaction from the pool it's in.
func (c *MempoolTxCache) Remove(cosmosTxKey string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	hash, ok := c.txKeyToHash[cosmosTxKey]
	if !ok {
		return
	}

	rpcTx, ok := c.byHash[hash]
	if !ok {
		delete(c.txKeyToHash, cosmosTxKey)
		return
	}

	from := rpcTx.From
	nonce := uint64(rpcTx.Nonce)

	if isPending, exists := c.pool[hash]; exists {
		if isPending {
			c.removeFromPendingLocked(from, nonce, hash)
		} else {
			c.removeFromQueuedLocked(from, nonce, hash)
		}
	}

	delete(c.txKeyToHash, cosmosTxKey)
}

// FindByHash looks up a transaction by its eth hash across both pools. O(1).
func (c *MempoolTxCache) FindByHash(hash common.Hash) *rpctypes.RPCTransaction {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.byHash[hash]
}

// PendingContent returns the pending pool.
func (c *MempoolTxCache) PendingContent() map[string]map[string]*rpctypes.RPCTransaction {
	c.mu.RLock()
	defer c.mu.RUnlock()

	result := make(map[string]map[string]*rpctypes.RPCTransaction, len(c.pending))
	for addr, nonces := range c.pending {
		addrHex := addr.Hex()
		dump := make(map[string]*rpctypes.RPCTransaction, len(nonces))
		for nonce, tx := range nonces {
			dump[fmt.Sprintf("%d", nonce)] = tx
		}
		result[addrHex] = dump
	}

	return result
}

// QueuedContent returns the queued pool.
func (c *MempoolTxCache) QueuedContent() map[string]map[string]*rpctypes.RPCTransaction {
	c.mu.RLock()
	defer c.mu.RUnlock()

	result := make(map[string]map[string]*rpctypes.RPCTransaction, len(c.queued))
	for addr, nonces := range c.queued {
		addrHex := addr.Hex()
		dump := make(map[string]*rpctypes.RPCTransaction, len(nonces))
		for nonce, tx := range nonces {
			dump[fmt.Sprintf("%d", nonce)] = tx
		}
		result[addrHex] = dump
	}

	return result
}

// AllPending returns a flat list of all pending transactions.
func (c *MempoolTxCache) AllPending() []*rpctypes.RPCTransaction {
	c.mu.RLock()
	defer c.mu.RUnlock()

	result := make([]*rpctypes.RPCTransaction, 0, c.pendingCount.Load())
	for _, nonces := range c.pending {
		for _, tx := range nonces {
			result = append(result, tx)
		}
	}

	return result
}

// PendingCount returns the number of pending transactions.
func (c *MempoolTxCache) PendingCount() int {
	return int(c.pendingCount.Load())
}

// QueuedCount returns the number of queued transactions.
func (c *MempoolTxCache) QueuedCount() int {
	return int(c.queuedCount.Load())
}

func (c *MempoolTxCache) addToPendingLocked(from common.Address, nonce uint64, hash common.Hash, rpcTx *rpctypes.RPCTransaction, cosmosTxKey string) {
	if _, ok := c.pending[from]; !ok {
		c.pending[from] = make(map[uint64]*rpctypes.RPCTransaction)
	}

	c.pending[from][nonce] = rpcTx
	c.byHash[hash] = rpcTx
	c.pool[hash] = true
	c.txKeyToHash[cosmosTxKey] = hash
	c.pendingCount.Add(1)
}

func (c *MempoolTxCache) addToQueuedLocked(from common.Address, nonce uint64, hash common.Hash, rpcTx *rpctypes.RPCTransaction, cosmosTxKey string) {
	if _, ok := c.queued[from]; !ok {
		c.queued[from] = make(map[uint64]*rpctypes.RPCTransaction)
	}

	c.queued[from][nonce] = rpcTx
	c.byHash[hash] = rpcTx
	c.pool[hash] = false
	c.txKeyToHash[cosmosTxKey] = hash
	c.queuedCount.Add(1)
}

func (c *MempoolTxCache) removeFromPendingLocked(from common.Address, nonce uint64, hash common.Hash) {
	if nonces, ok := c.pending[from]; ok {
		delete(nonces, nonce)
		if len(nonces) == 0 {
			delete(c.pending, from)
		}
	}

	delete(c.byHash, hash)
	delete(c.pool, hash)

	c.pendingCount.Add(-1)
}

func (c *MempoolTxCache) removeFromQueuedLocked(from common.Address, nonce uint64, hash common.Hash) {
	if nonces, ok := c.queued[from]; ok {
		delete(nonces, nonce)
		if len(nonces) == 0 {
			delete(c.queued, from)
		}
	}

	delete(c.byHash, hash)
	delete(c.pool, hash)

	c.queuedCount.Add(-1)
}

// TxConverter converts raw cosmos tx bytes into an RPCTransaction.
// Returns nil for non evm tx or on error.
type TxConverter func(txBytes []byte) *rpctypes.RPCTransaction

// StartEventConsumer processes mempool events and updates the cache.
func (c *MempoolTxCache) StartEventConsumer(events <-chan cmtmempool.AppMempoolEvent, convert TxConverter, pendingTxChan chan<- *rpctypes.RPCTransaction) {
	go func() {
		for event := range events {
			cosmosTxKey := fmt.Sprintf("%x", event.TxKey)

			switch event.Type {
			case cmtmempool.EventTxInserted:
				rpcTx := convert(event.Tx)
				if rpcTx != nil {
					c.HandleInserted(rpcTx, cosmosTxKey)
					if pendingTxChan != nil {
						select {
						case pendingTxChan <- rpcTx:
						default:
						}
					}
				}

			case cmtmempool.EventTxQueued:
				rpcTx := convert(event.Tx)
				if rpcTx != nil {
					c.AddQueued(rpcTx, cosmosTxKey)
				}

			case cmtmempool.EventTxRemoved:
				c.Remove(cosmosTxKey)
			}
		}
	}()
}
