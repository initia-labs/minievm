package indexer

import (
	"context"

	"cosmossdk.io/collections"
	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"golang.org/x/sync/errgroup"

	"github.com/ethereum/go-ethereum/common"
	coretypes "github.com/ethereum/go-ethereum/core/types"
)

const (
	pruneStoreKey = iota
)

// doPrune triggers pruning in a goroutine. If pruning is already running,
// it does nothing.
func (e *EVMIndexerImpl) doPrune(ctx context.Context, height uint64) {
	if running := e.pruningRunning.Swap(true); running {
		return
	}

	go func(ctx context.Context, height uint64) {
		defer e.pruningRunning.Store(false)
		if err := e.prune(ctx, height); err != nil {
			e.logger.Error("failed to prune", "err", err)
		}

		e.logger.Debug("prune finished", "height", height)
	}(ctx, height)
}

// prune removes old blocks and transactions from the indexer.
func (e *EVMIndexerImpl) prune(ctx context.Context, curHeight uint64) error {
	// use branch context to perform batch operations
	batchStore := e.store.store.CacheWrap()
	ctx = sdk.UnwrapSDKContext(ctx).WithValue(pruneStoreKey, newCoreKVStore(interface{}(batchStore).(storetypes.KVStore)))

	minHeight := curHeight - e.retainHeight
	if minHeight <= 0 || minHeight >= curHeight {
		return nil
	}

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return e.pruneBlocks(ctx, minHeight)
	})
	g.Go(func() error {
		return e.pruneTxs(ctx, minHeight)
	})

	if err := g.Wait(); err != nil {
		return err
	}

	// write batch
	batchStore.Write()

	return nil
}

// pruneBlocks removes old block headers from the indexer.
func (e *EVMIndexerImpl) pruneBlocks(ctx context.Context, minHeight uint64) error {
	// record block hashes
	var blockHashes []common.Hash
	rn := new(collections.Range[uint64]).EndInclusive(minHeight)
	err := e.BlockHeaderMap.Walk(ctx, rn, func(key uint64, value coretypes.Header) (stop bool, err error) {
		blockHashes = append(blockHashes, value.Hash())
		return false, nil
	})
	if err != nil {
		return err
	}

	// clear block headers within range
	err = e.BlockHeaderMap.Clear(ctx, rn)
	if err != nil {
		return err
	}

	// clear block hash to number map
	for _, blockHash := range blockHashes {
		if err := e.BlockHashToNumberMap.Remove(ctx, blockHash.Bytes()); err != nil {
			return err
		}
	}

	return nil
}

// pruneTxs removes old transactions from the indexer.
func (e *EVMIndexerImpl) pruneTxs(ctx context.Context, minHeight uint64) error {
	// record tx hashes
	var txHashes []common.Hash
	rnPair := collections.NewPrefixUntilPairRange[uint64, uint64](minHeight)
	err := e.BlockAndIndexToTxHashMap.Walk(ctx, rnPair, func(key collections.Pair[uint64, uint64], txHashBz []byte) (bool, error) {
		txHash := common.BytesToHash(txHashBz)
		txHashes = append(txHashes, txHash)
		return false, nil
	})
	if err != nil {
		return err
	}

	// clear block txs within range
	err = e.BlockAndIndexToTxHashMap.Clear(ctx, rnPair)
	if err != nil {
		return err
	}

	// clear txs and receipts and cosmos tx hash mappings
	for _, txHash := range txHashes {
		if err := e.TxMap.Remove(ctx, txHash.Bytes()); err != nil {
			return err
		}
		if err := e.TxReceiptMap.Remove(ctx, txHash.Bytes()); err != nil {
			return err
		}
		if cosmosTxHash, err := e.TxHashToCosmosTxHash.Get(ctx, txHash.Bytes()); err == nil {
			if err := e.TxHashToCosmosTxHash.Remove(ctx, txHash.Bytes()); err != nil {
				return err
			}
			if err := e.CosmosTxHashToTxHash.Remove(ctx, cosmosTxHash); err != nil {
				return err
			}
		} else if err != collections.ErrNotFound {
			return err
		}
	}

	return nil
}

//////////////////////// TESTING INTERFACE ////////////////////////

// Set custom retain height
func (e *EVMIndexerImpl) SetRetainHeight(height uint64) {
	e.retainHeight = height
}

// Check if pruning is running
func (e *EVMIndexerImpl) IsPruningRunning() bool {
	return e.pruningRunning.Load()
}

// Clear cache for testing
func (e *EVMIndexerImpl) ClearCache() error {
	return e.store.cache.Reset()
}
