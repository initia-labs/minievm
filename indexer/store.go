package indexer

import (
	"sync"

	corestoretypes "cosmossdk.io/core/store"
	"cosmossdk.io/store/cachekv"
	"cosmossdk.io/store/dbadapter"
	storetypes "cosmossdk.io/store/types"
	dbm "github.com/cosmos/cosmos-db"
)

var _ corestoretypes.KVStore = (*CacheStoreWithBatch)(nil)

const (
	// DefaultBatchFlushThreshold defines the maximum batch size (in bytes) before automatically writing to disk.
	DefaultBatchFlushThreshold = 4 * 1024 * 1024 // 4MB
)

// CacheStoreWithBatch is a CacheKVStore implementation that combines multi-level caching with batched
// database operations.
type CacheStoreWithBatch struct {
	store storetypes.CacheKVStore // Core CacheKVStore interface for standard operations
	batch dbm.Batch               // Batch accumulator for database operations
	db    dbm.DB                  // Underlying persistent database
	mtx   sync.RWMutex
}

// NewCacheStoreWithBatch creates a new cache store with batched write capabilities.
func NewCacheStoreWithBatch(db dbm.DB) *CacheStoreWithBatch {
	return &CacheStoreWithBatch{
		store: cachekv.NewStore(dbadapter.Store{DB: db}),
		db:    db,
		batch: db.NewBatch(),
	}
}

// Get retrieves a value for the given key, prioritizing cache lookups for performance.
// Returns a nil iff key doesn't exist, errors on a nil key.
func (c *CacheStoreWithBatch) Get(key []byte) ([]byte, error) {
	storetypes.AssertValidKey(key)

	c.mtx.RLock()
	defer c.mtx.RUnlock()

	return c.store.Get(key), nil
}

// Has checks if a key exists. Errors on a nil key.
func (c *CacheStoreWithBatch) Has(key []byte) (bool, error) {
	storetypes.AssertValidKey(key)

	c.mtx.RLock()
	defer c.mtx.RUnlock()

	return c.store.Get(key) != nil, nil
}

// Set stores a key-value pair and adds it to the writing batch.
// Automatically writes the batch to the underlying database if the batch threshold is reached.
func (c *CacheStoreWithBatch) Set(key, value []byte) error {
	storetypes.AssertValidKey(key)
	storetypes.AssertValidValue(value)

	c.mtx.Lock()
	defer c.mtx.Unlock()

	batchSizeAfter, err := c.estimateSizeAfterSetting(key, value)
	if err != nil {
		return err
	}
	if batchSizeAfter > DefaultBatchFlushThreshold {
		c.writeBatchUnlocked()
	}

	// update cache store
	c.store.Set(key, value)

	// add to batch for persistence
	return c.batch.Set(key, value)
}

// Delete removes a key-value pair from all store layers and schedules the deletion in the writing batch.
// Like Set, it uses threshold-based batch commits for optimal performance.
func (c *CacheStoreWithBatch) Delete(key []byte) error {
	storetypes.AssertValidKey(key)

	c.mtx.Lock()
	defer c.mtx.Unlock()

	batchSizeAfter, err := c.estimateSizeAfterSetting(key, []byte{})
	if err != nil {
		return err
	}
	if batchSizeAfter > DefaultBatchFlushThreshold {
		c.writeBatchUnlocked()
	}

	// update cache store
	c.store.Delete(key)

	// add to batch for persistence
	return c.batch.Delete(key)
}

// Iterator iterates over a domain of keys in ascending order. End is exclusive.
// Start must be less than end, or the Iterator is invalid.
// Iterator must be closed by caller.
// To iterate over entire domain, use store.Iterator(nil, nil)
// CONTRACT: No writes may happen within a domain while an iterator exists over it.
// Exceptionally allowed for cachekv.Store, safe to write in the modules.
func (c *CacheStoreWithBatch) Iterator(start, end []byte) (storetypes.Iterator, error) {
	c.mtx.RLock()
	defer c.mtx.RUnlock()

	return c.store.Iterator(start, end), nil
}

// ReverseIterator iterates over a domain of keys in descending order. End is exclusive.
// Start must be less than end, or the Iterator is invalid.
// Iterator must be closed by caller.
// CONTRACT: No writes may happen within a domain while an iterator exists over it.
// Exceptionally allowed for cachekv.Store, safe to write in the modules.
func (c *CacheStoreWithBatch) ReverseIterator(start, end []byte) (storetypes.Iterator, error) {
	c.mtx.RLock()
	defer c.mtx.RUnlock()

	return c.store.ReverseIterator(start, end), nil
}

// Write explicitly flushes the current batch to persistent storage.
func (c *CacheStoreWithBatch) Write() {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	c.writeBatchUnlocked()
}

// writeBatchUnlocked writes the batch to persistent storage and resets both the batch and cache,
// assuming that the caller has already acquired the appropriate locks.
func (c *CacheStoreWithBatch) writeBatchUnlocked() {
	_ = c.batch.Write()
	_ = c.batch.Close()

	// clear the batch after writing
	c.batch = c.db.NewBatch()

	// also clear the cache after writing
	c.store = cachekv.NewStore(dbadapter.Store{DB: c.db})
}

// estimateSizeAfterSetting estimates the batch's size after setting a key / value
func (c *CacheStoreWithBatch) estimateSizeAfterSetting(key []byte, value []byte) (int, error) {
	currentSize, err := c.batch.GetByteSize()
	if err != nil {
		return 0, err
	}

	// add 100 here just to overcompensate for overhead
	return currentSize + len(key) + len(value) + 100, nil
}
