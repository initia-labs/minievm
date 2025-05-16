package indexer

import (
	"context"
	"sync"

	corestoretypes "cosmossdk.io/core/store"
	"cosmossdk.io/store/cachekv"
	"cosmossdk.io/store/dbadapter"
	storetypes "cosmossdk.io/store/types"
	"github.com/allegro/bigcache/v3"
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
	cache *bigcache.BigCache      // Fast in-memory cache for frequently accessed data
	batch dbm.Batch               // Batch accumulator for database operations
	db    dbm.DB                  // Underlying persistent database
	mtx   sync.RWMutex
}

// NewCacheStoreWithBatch creates a new cache store with batched write capabilities.
func NewCacheStoreWithBatch(db dbm.DB, capacity int) *CacheStoreWithBatch {
	// default with no eviction and custom hard max cache capacity
	cacheCfg := bigcache.DefaultConfig(0)
	cacheCfg.Verbose = false
	cacheCfg.HardMaxCacheSize = capacity

	cache, err := bigcache.New(context.Background(), cacheCfg)
	if err != nil {
		panic(err)
	}

	return &CacheStoreWithBatch{
		store: cachekv.NewStore(dbadapter.Store{DB: db}),
		cache: cache,
		db:    db,
	}
}

// Get retrieves a value for the given key, prioritizing cache lookups for performance.
// Returns a nil iff key doesn't exist, errors on a nil key.
func (c *CacheStoreWithBatch) Get(key []byte) ([]byte, error) {
	storetypes.AssertValidKey(key)

	if value, err := c.cache.Get(string(key)); err == nil {
		return value, nil
	}

	c.mtx.RLock()
	defer c.mtx.RUnlock()

	// if not in cache, get from kvstore
	value := c.store.Get(key)
	if value == nil {
		return nil, nil
	}

	// cache for future reads, ignore error
	_ = c.cache.Set(string(key), value)

	return value, nil
}

// Has checks if a key exists. Errors on a nil key.
func (c *CacheStoreWithBatch) Has(key []byte) (bool, error) {
	_, err := c.cache.Get(string(key))
	if err == nil {
		return true, nil
	}

	c.mtx.RLock()
	defer c.mtx.RUnlock()

	value := c.store.Get(key)
	if value == nil {
		return false, nil
	}

	// ignore cache error
	_ = c.cache.Set(string(key), value)

	return true, nil
}

// Set stores a key-value pair and adds it to the writing batch.
// Automatically writes the batch to the underlying database if the batch threshold is reached.
func (c *CacheStoreWithBatch) Set(key, value []byte) error {
	storetypes.AssertValidKey(key)
	storetypes.AssertValidValue(value)

	c.mtx.Lock()
	defer c.mtx.Unlock()

	// update both caches
	_ = c.cache.Set(string(key), value)
	c.store.Set(key, value)

	// ensure batch is available
	if c.batch == nil {
		c.batch = c.db.NewBatch()
	}

	batchSizeAfter, err := c.estimateSizeAfterSetting(key, value)
	if err != nil {
		return err
	}
	if batchSizeAfter > DefaultBatchFlushThreshold {
		c.writeBatchUnlocked()
		c.batch = c.db.NewBatch()
	}

	// add to batch for persistence
	return c.batch.Set(key, value)
}

// Delete removes a key-value pair from all store layers and schedules the deletion in the writing batch.
// Like Set, it uses threshold-based batch commits for optimal performance.
func (c *CacheStoreWithBatch) Delete(key []byte) error {
	storetypes.AssertValidKey(key)

	c.mtx.Lock()
	defer c.mtx.Unlock()

	// update both caches
	_ = c.cache.Delete(string(key))
	c.store.Delete(key)

	// ensure batch is available
	if c.batch == nil {
		c.batch = c.db.NewBatch()
	}

	batchSizeAfter, err := c.estimateSizeAfterSetting(key, []byte{})
	if err != nil {
		return err
	}
	if batchSizeAfter > DefaultBatchFlushThreshold {
		c.writeBatchUnlocked()
		c.batch = c.db.NewBatch()
	}

	// add to batch for persistence
	return c.batch.Delete(key)
}

// cacheIterator extends the standard iterator with cache awareness
type cacheIterator struct {
	storetypes.Iterator
	cache *bigcache.BigCache
}

// Value returns the value at the current position
func (i *cacheIterator) Value() []byte {
	key := i.Key()

	// try cache first
	if value, err := i.cache.Get(string(key)); err == nil {
		return value
	}

	return i.Iterator.Value()
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

	return &cacheIterator{
		Iterator: c.store.Iterator(start, end),
		cache:    c.cache,
	}, nil
}

// ReverseIterator iterates over a domain of keys in descending order. End is exclusive.
// Start must be less than end, or the Iterator is invalid.
// Iterator must be closed by caller.
// CONTRACT: No writes may happen within a domain while an iterator exists over it.
// Exceptionally allowed for cachekv.Store, safe to write in the modules.
func (c *CacheStoreWithBatch) ReverseIterator(start, end []byte) (storetypes.Iterator, error) {
	c.mtx.RLock()
	defer c.mtx.RUnlock()

	return &cacheIterator{
		Iterator: c.store.ReverseIterator(start, end),
		cache:    c.cache,
	}, nil
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
	if c.batch == nil {
		return
	}

	_ = c.batch.Write()
	_ = c.batch.Close()

	// clear the batch after writing
	c.batch = nil

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
