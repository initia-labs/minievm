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
	DefaultBatchFlushThreshold = 4 * 1024 * 1024 // 4MB
)

type CacheStoreWithBatch struct {
	store storetypes.CacheKVStore
	cache *bigcache.BigCache
	batch dbm.Batch
	db    dbm.DB
	mtx   sync.Mutex
}

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

// Get returns nil iff key doesn't exist. Errors on nil key.
func (c *CacheStoreWithBatch) Get(key []byte) ([]byte, error) {
	storetypes.AssertValidKey(key)

	if value, err := c.cache.Get(string(key)); err == nil {
		return value, nil
	}

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

	value := c.store.Get(key)
	if value == nil {
		return false, nil
	}

	// ignore cache error
	_ = c.cache.Set(string(key), value)

	return true, nil
}

// Set sets the key. Errors on nil key or value.
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
		c.Write()
		c.batch = c.db.NewBatch()
	}

	// add to batch for persistence
	return c.batch.Set(key, value)
}

// Delete deletes the key. Errors on a nil key.
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
		c.Write()
		c.batch = c.db.NewBatch()
	}

	// add to batch for persistence
	return c.batch.Delete(key)
}

// cacheIterator wraps a storetypes.Iterator to add caching
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

	// get from the underlying iterator
	value := i.Iterator.Value()
	if value != nil {
		_ = i.cache.Set(string(key), value)
	}

	return value
}

// Iterator iterates over a domain of keys in ascending order. End is exclusive.
// Start must be less than end, or the Iterator is invalid.
// Iterator must be closed by caller.
// To iterate over entire domain, use store.Iterator(nil, nil)
// CONTRACT: No writes may happen within a domain while an iterator exists over it.
// Exceptionally allowed for cachekv.Store, safe to write in the modules.
func (c *CacheStoreWithBatch) Iterator(start, end []byte) (storetypes.Iterator, error) {
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
	return &cacheIterator{
		Iterator: c.store.ReverseIterator(start, end),
		cache:    c.cache,
	}, nil
}

// Write writes the batch to the store.
func (c *CacheStoreWithBatch) Write() {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	if c.batch == nil {
		return
	}

	_ = c.batch.Write()
	_ = c.batch.Close()

	// clear the batch after writing
	c.batch = nil
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
