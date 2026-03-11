package indexer

import (
	"context"
	"errors"
	"time"

	"cosmossdk.io/collections"
	evmconfig "github.com/initia-labs/minievm/x/evm/config"
	"golang.org/x/sync/errgroup"

	"github.com/ethereum/go-ethereum/common"
	coretypes "github.com/ethereum/go-ethereum/core/types"
)

// doPrune records a prune target and notifies the prune worker.
func (e *EVMIndexerImpl) doPrune(ctx context.Context, height uint64) {
	_ = ctx

	for {
		prev := e.pruneRequestedHeight.Load()
		if height <= prev {
			break
		}
		if e.pruneRequestedHeight.CompareAndSwap(prev, height) {
			break
		}
	}

	// Coalesce wakeups; worker always loads the latest requested height.
	select {
	case e.pruneNotifyCh <- struct{}{}:
	default:
	}
}

func (e *EVMIndexerImpl) pruneLoop() {
	defer close(e.pruneDoneCh)

	ctx := context.Background()
	for {
		select {
		case <-e.pruneStopCh:
			return
		case <-e.pruneNotifyCh:
		}

		for {
			targetHeight := e.pruneRequestedHeight.Load()
			if targetHeight <= e.lastPruneTriggerHeight.Load() {
				break
			}

			e.pruningRunning.Store(true)
			err := e.prune(ctx, targetHeight)
			e.pruningRunning.Store(false)
			if err != nil {
				e.logger.Error("failed to prune", "height", targetHeight, "err", err)
				// Back off on repeated failures, but remain responsive to shutdown.
				select {
				case <-time.After(100 * time.Millisecond):
				case <-e.pruneStopCh:
					return
				}
				continue
			}

			e.logger.Debug("prune finished", "height", targetHeight)

			// If a newer target arrived while pruning, keep draining requests.
			if e.pruneRequestedHeight.Load() <= targetHeight {
				break
			}
		}
	}
}

// prune removes old blocks and transactions from the indexer.
func (e *EVMIndexerImpl) prune(ctx context.Context, curHeight uint64) error {
	minHeight := curHeight - e.retainHeight
	if minHeight == 0 || minHeight >= curHeight {
		e.lastPruneTriggerHeight.Store(curHeight)
		return nil
	}

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return e.pruneBlocks(ctx, minHeight)
	})
	g.Go(func() error {
		return e.pruneTxs(ctx, minHeight)
	})
	g.Go(func() error {
		return e.pruneBloomBits(ctx, minHeight)
	})

	if err := g.Wait(); err != nil {
		return err
	}

	// write the changes to the store
	e.store.Write()

	// update the last prune trigger height
	e.lastPruneTriggerHeight.Store(curHeight)

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
		if err := e.TxStartLogIndexMap.Remove(ctx, txHash.Bytes()); err != nil && !errors.Is(err, collections.ErrNotFound) {
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

// pruneBloomBits removes old bloom bits from the indexer.
func (e *EVMIndexerImpl) pruneBloomBits(ctx context.Context, minHeight uint64) error {
	prunedSections := (minHeight + 1) / evmconfig.SectionSize
	if prunedSections == 0 {
		return nil
	}
	section := prunedSections - 1
	return e.BloomBits.Clear(ctx, collections.NewPrefixUntilPairRange[uint64, uint32](section))
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
