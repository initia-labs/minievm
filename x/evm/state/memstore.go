package state

import (
	"fmt"

	corestoretypes "cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"cosmossdk.io/store"
	storemetrics "cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
	dbm "github.com/cosmos/cosmos-db"
)

const (
	storeKey = "stateDB"
)

func newMemStore() (storetypes.MultiStore, storetypes.StoreKey) {
	memStoreKey := storetypes.NewMemoryStoreKey(storeKey)
	memStore := store.NewCommitMultiStore(dbm.NewMemDB(), log.NewNopLogger(), storemetrics.NewNoOpMetrics())
	memStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	if err := memStore.LoadLatestVersion(); err != nil {
		panic(fmt.Sprintf("failed to initialize memory store: %v", err))
	}

	return memStore, memStoreKey
}

// CoreKVStore is a wrapper of Core/Store kvstore interface
// Remove after https://github.com/cosmos/cosmos-sdk/issues/14714 is closed
type coreKVStore struct {
	kvStore storetypes.KVStore
}

// NewKVStore returns a wrapper of Core/Store kvstore interface
// Remove once store migrates to core/store kvstore interface
func newKVStore(store storetypes.KVStore) corestoretypes.KVStore {
	return coreKVStore{kvStore: store}
}

// Get returns nil iff key doesn't exist. Errors on nil key.
func (store coreKVStore) Get(key []byte) ([]byte, error) {
	return store.kvStore.Get(key), nil
}

// Has checks if a key exists. Errors on nil key.
func (store coreKVStore) Has(key []byte) (bool, error) {
	return store.kvStore.Has(key), nil
}

// Set sets the key. Errors on nil key or value.
func (store coreKVStore) Set(key, value []byte) error {
	store.kvStore.Set(key, value)
	return nil
}

// Delete deletes the key. Errors on nil key.
func (store coreKVStore) Delete(key []byte) error {
	store.kvStore.Delete(key)
	return nil
}

// Iterator iterates over a domain of keys in ascending order. End is exclusive.
// Start must be less than end, or the Iterator is invalid.
// Iterator must be closed by caller.
// To iterate over entire domain, use store.Iterator(nil, nil)
// CONTRACT: No writes may happen within a domain while an iterator exists over it.
// Exceptionally allowed for cachekv.Store, safe to write in the modules.
func (store coreKVStore) Iterator(start, end []byte) (store.Iterator, error) {
	return store.kvStore.Iterator(start, end), nil
}

// ReverseIterator iterates over a domain of keys in descending order. End is exclusive.
// Start must be less than end, or the Iterator is invalid.
// Iterator must be closed by caller.
// CONTRACT: No writes may happen within a domain while an iterator exists over it.
// Exceptionally allowed for cachekv.Store, safe to write in the modules.
func (store coreKVStore) ReverseIterator(start, end []byte) (store.Iterator, error) {
	return store.kvStore.ReverseIterator(start, end), nil
}
