package indexer

import (
	"context"

	corestoretypes "cosmossdk.io/core/store"
	"cosmossdk.io/errors"
	cachekv "cosmossdk.io/store/cachekv"
	storetypes "cosmossdk.io/store/types"

	bigcache "github.com/allegro/bigcache/v3"
)

var _ corestoretypes.KVStore = (*CacheStore)(nil)

type CacheStore struct {
	store storetypes.CacheKVStore
	cache *bigcache.BigCache
}

func NewCacheStore(store storetypes.KVStore, capacity int) *CacheStore {
	// default with no eviction and custom hard max cache capacity
	cacheCfg := bigcache.DefaultConfig(0)
	cacheCfg.Verbose = false
	cacheCfg.HardMaxCacheSize = capacity

	cache, err := bigcache.New(context.Background(), cacheCfg)
	if err != nil {
		panic(err)
	}

	return &CacheStore{
		store: cachekv.NewStore(store),
		cache: cache,
	}
}

// Get returns nil iff key doesn't exist. Errors on nil key.
func (c CacheStore) Get(key []byte) ([]byte, error) {
	storetypes.AssertValidKey(key)

	if value, err := c.cache.Get(string(key)); err == nil {
		return value, nil
	}

	// get from store and write to cache
	value := c.store.Get(key)
	if value == nil {
		return nil, nil
	}

	// ignore cache error
	_ = c.cache.Set(string(key), value)

	return value, nil
}

// Has checks if a key exists. Errors on nil key.
func (c CacheStore) Has(key []byte) (bool, error) {
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
func (c CacheStore) Set(key, value []byte) error {
	storetypes.AssertValidKey(key)
	storetypes.AssertValidValue(value)

	err := c.cache.Set(string(key), value)
	if err != nil {
		return errors.Wrap(err, "failed to set cache")
	}
	c.store.Set(key, value)

	return nil
}

// Delete deletes the key. Errors on nil key.
func (c CacheStore) Delete(key []byte) error {
	storetypes.AssertValidKey(key)

	err := c.cache.Delete(string(key))
	if err != nil && errors.IsOf(err, bigcache.ErrEntryNotFound) {
		return errors.Wrap(err, "failed to delete cache")
	}
	c.store.Delete(key)

	return nil
}

// Iterator iterates over a domain of keys in ascending order. End is exclusive.
// Start must be less than end, or the Iterator is invalid.
// Iterator must be closed by caller.
// To iterate over entire domain, use store.Iterator(nil, nil)
// CONTRACT: No writes may happen within a domain while an iterator exists over it.
// Exceptionally allowed for cachekv.Store, safe to write in the modules.
func (c CacheStore) Iterator(start, end []byte) (storetypes.Iterator, error) {
	return c.store.Iterator(start, end), nil
}

// ReverseIterator iterates over a domain of keys in descending order. End is exclusive.
// Start must be less than end, or the Iterator is invalid.
// Iterator must be closed by caller.
// CONTRACT: No writes may happen within a domain while an iterator exists over it.
// Exceptionally allowed for cachekv.Store, safe to write in the modules.
func (c CacheStore) ReverseIterator(start, end []byte) (storetypes.Iterator, error) {
	return c.store.ReverseIterator(start, end), nil
}

func (c CacheStore) Write() {
	c.store.Write()
}
