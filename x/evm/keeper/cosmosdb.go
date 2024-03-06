package keeper

import (
	"context"
	"errors"

	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/ethdb"
)

var _ ethdb.KeyValueStore = &CosmosDB{}
var _ ethdb.Snapshot = &CosmosDB{}

func (k Keeper) newStateDB(ctx context.Context) (vm.StateDB, error) {
	vmRoot, err := k.VMRoot.Get(ctx)
	if err != nil && errors.Is(err, collections.ErrNotFound) {
		vmRoot = coretypes.EmptyRootHash[:]
	} else if err != nil {
		return nil, err
	}

	return state.New(
		common.Hash(vmRoot),
		state.NewDatabase(
			rawdb.NewDatabase(
				&CosmosDB{ctx: ctx, vmStore: k.VMStore},
			),
		),
		nil,
	)
}

// CosmosDB is a wrapper of ethdb.KeyValueStore
type CosmosDB struct {
	ctx     context.Context
	vmStore collections.Map[[]byte, []byte]
}

// Release implements ethdb.Snapshot.
func (*CosmosDB) Release() {}

// Close is not supported on a cosmos database.
func (db *CosmosDB) Close() error {
	return nil
}

// Compact is not supported on a cosmos database.
func (*CosmosDB) Compact(start []byte, limit []byte) error {
	return nil
}

// Delete inserts the a key removal into the batch for later committing.
func (db *CosmosDB) Delete(key []byte) error {
	return db.vmStore.Remove(db.ctx, key)
}

// Get implements ethdb.KeyValueStore.
func (db *CosmosDB) Get(key []byte) ([]byte, error) {
	return db.vmStore.Get(db.ctx, key)
}

// Has implements ethdb.KeyValueStore.
func (db *CosmosDB) Has(key []byte) (bool, error) {
	return db.vmStore.Has(db.ctx, key)
}

// Put inserts the given value into the batch for later committing.
func (db *CosmosDB) Put(key []byte, value []byte) error {
	return db.vmStore.Set(db.ctx, key, value)
}

// NewBatch implements ethdb.KeyValueStore.
func (db *CosmosDB) NewBatch() ethdb.Batch {
	return &batch{db: db, writes: []keyvalue{}}
}

// NewBatchWithSize implements ethdb.KeyValueStore.
func (db *CosmosDB) NewBatchWithSize(size int) ethdb.Batch {
	return &batch{db: db, writes: make([]keyvalue, 0, size)}
}

// NewIterator implements ethdb.KeyValueStore.
func (db *CosmosDB) NewIterator(prefix []byte, start []byte) ethdb.Iterator {
	iter, err := db.vmStore.Iterate(db.ctx, new(collections.Range[[]byte]).Prefix(prefix).StartInclusive(start))
	if err != nil {
		return &iterator{
			inner: collections.Iterator[[]byte, []byte]{},
			err:   err,
		}
	}

	return &iterator{
		inner: iter,
	}
}

// NewSnapshot implements ethdb.KeyValueStore.
func (db *CosmosDB) NewSnapshot() (ethdb.Snapshot, error) {
	sdkCtx := sdk.UnwrapSDKContext(db.ctx)
	cacheCtx, _ := sdkCtx.CacheContext()

	return &CosmosDB{ctx: cacheCtx, vmStore: db.vmStore}, nil
}

// Stat returns a particular internal stat of the database.
func (*CosmosDB) Stat(property string) (string, error) {
	return "", errors.New("unknown property")
}

////////////////////////////////////////////
// iterator implementation

var _ ethdb.Iterator = &iterator{}

type iterator struct {
	inner collections.Iterator[[]byte, []byte]
	err   error
}

// Error implements ethdb.Iterator.
func (iter *iterator) Error() error {
	return iter.err
}

// Key implements ethdb.Iterator.
func (iter *iterator) Key() []byte {
	if iter.err != nil {
		return nil
	}

	key, err := iter.inner.Key()
	if err != nil {
		// store error for Error()
		iter.err = err
		return nil
	}

	return key
}

// Value implements ethdb.Iterator.
func (iter *iterator) Value() []byte {
	if iter.err != nil {
		return nil
	}

	val, err := iter.inner.Value()
	if err != nil {
		// store error for Error()
		iter.err = err
		return nil
	}

	return val
}

// Next implements ethdb.Iterator.
func (iter *iterator) Next() bool {
	if iter.err != nil {
		return false
	}

	iter.inner.Next()
	return iter.inner.Valid()
}

// Release implements ethdb.Iterator.
func (iter *iterator) Release() {
	if iter.err != nil {
		return
	}

	// error maybe ignored
	iter.err = iter.inner.Close()
}

////////////////////////////////////////////
// batch implementation

// keyvalue is a key-value tuple tagged with a deletion field to allow creating
// memory-database write batches.
type keyvalue struct {
	key    []byte
	value  []byte
	delete bool
}

// batch is a write-only memory batch that commits changes to its host
// database when Write is called. A batch cannot be used concurrently.
type batch struct {
	db     *CosmosDB
	writes []keyvalue
	size   int
}

// Put inserts the given value into the batch for later committing.
func (b *batch) Put(key, value []byte) error {
	b.writes = append(b.writes, keyvalue{key, common.CopyBytes(value), false})
	b.size += len(key) + len(value)
	return nil
}

// Delete inserts the a key removal into the batch for later committing.
func (b *batch) Delete(key []byte) error {
	b.writes = append(b.writes, keyvalue{key, nil, true})
	b.size += len(key)
	return nil
}

// ValueSize retrieves the amount of data queued up for writing.
func (b *batch) ValueSize() int {
	return b.size
}

// Write flushes any accumulated data to the memory database.
func (b *batch) Write() error {
	for _, keyvalue := range b.writes {
		switch keyvalue.delete {
		case true:
			if err := b.db.Delete(keyvalue.key); err != nil {
				return err
			}
		case false:
			if err := b.db.Put(keyvalue.key, keyvalue.value); err != nil {
				return err
			}
		}
	}

	return nil
}

// Reset resets the batch for reuse.
func (b *batch) Reset() {
	b.writes = b.writes[:0]
	b.size = 0
}

// Replay replays the batch contents.
func (b *batch) Replay(w ethdb.KeyValueWriter) error {
	for _, keyvalue := range b.writes {
		switch keyvalue.delete {
		case true:
			if err := w.Delete([]byte(keyvalue.key)); err != nil {
				return err
			}
		case false:
			if err := w.Put([]byte(keyvalue.key), keyvalue.value); err != nil {
				return err
			}
		}
	}

	return nil
}
