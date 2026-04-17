package indexer

import (
	"context"
	"errors"
	"math/big"

	"cosmossdk.io/collections"
	"github.com/ethereum/go-ethereum/core/bloombits"
	coretypes "github.com/ethereum/go-ethereum/core/types"

	evmconfig "github.com/initia-labs/minievm/x/evm/config"
)

// doBloomIndexing records a bloom indexing target and notifies the bloom worker.
func (e *EVMIndexerImpl) doBloomIndexing(ctx context.Context, height uint64) {
	_ = ctx

	for {
		prev := e.bloomRequestedHeight.Load()
		if height <= prev {
			break
		}
		if e.bloomRequestedHeight.CompareAndSwap(prev, height) {
			break
		}
	}

	// Coalesce wakeups; worker always loads the latest requested height.
	select {
	case e.bloomNotifyCh <- struct{}{}:
	default:
	}
}

func (e *EVMIndexerImpl) bloomLoop() {
	defer close(e.bloomDoneCh)

	ctx, cancel := context.WithCancel(context.Background())
	for {
		select {
		case <-e.bloomStopCh:
			cancel()
			return
		case <-e.bloomNotifyCh:
		}

		e.bloomIndexingRunning.Store(true)
		for {
			targetHeight := e.bloomRequestedHeight.Load()
			lastIndexedHeight := e.lastBloomIndexedHeight.Load()
			if targetHeight <= lastIndexedHeight {
				e.logger.Debug("bloom indexing finished", "height", targetHeight)
				break
			}

			prevIndexedHeight := lastIndexedHeight
			if err := e.bloomIndexing(ctx, targetHeight); err != nil {
				e.logger.Error("failed to do bloom indexing", "height", targetHeight, "err", err)
				break
			}

			currIndexedHeight := e.lastBloomIndexedHeight.Load()
			// If a newer target arrived while indexing, continue with the latest target.
			if e.bloomRequestedHeight.Load() > targetHeight {
				continue
			}
			// No section was indexed; wait for a new block-triggered notification.
			if currIndexedHeight <= prevIndexedHeight {
				break
			}
			// Keep draining until we index up to the target height.
			if currIndexedHeight < targetHeight {
				continue
			}

			e.logger.Debug("bloom indexing finished", "height", targetHeight)
			break
		}
		e.bloomIndexingRunning.Store(false)
	}
}

// bloomIndexing generates the bloom index if the current section is complete.
func (e *EVMIndexerImpl) bloomIndexing(ctx context.Context, height uint64) error {
	section, err := e.PeekBloomBitsNextSection()
	if err != nil {
		return err
	}
	if (height / evmconfig.SectionSize) <= section {
		return nil
	}

	e.logger.Info("Processing new bloom indexing section", "section", section)

	gen, err := bloombits.NewGenerator(uint(evmconfig.SectionSize))
	if err != nil {
		return err
	}

	for i := range evmconfig.SectionSize {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		height := section*evmconfig.SectionSize + i
		header, err := e.BlockHeaderByNumber(height)
		if err != nil && errors.Is(err, collections.ErrNotFound) {
			// pruned block, create a dummy header
			header = &coretypes.Header{
				Number: new(big.Int).SetUint64(height),
				Bloom:  coretypes.Bloom{},
			}
		} else if err != nil {
			return err
		}

		if err := gen.AddBloom(uint(header.Number.Uint64()-section*evmconfig.SectionSize), header.Bloom); err != nil {
			return err
		}
	}

	// write the bloom bits to the store
	for i := range coretypes.BloomBitLength {
		bits, err := gen.Bitset(uint(i))
		if err != nil {
			return err
		}

		if err := e.RecordBloomBits(section, uint32(i), bits); err != nil {
			return err
		}
	}

	// increment the section number; if this fails, the section will be reprocessed
	if err := e.NextBloomBitsSection(); err != nil {
		return err
	}

	// update the last bloom indexed height to the end of the indexed section
	e.lastBloomIndexedHeight.Store((section + 1) * evmconfig.SectionSize)

	return nil
}

// ReadBloomBits reads the bloom bits for the given index, section and hash.
func (e *EVMIndexerImpl) ReadBloomBits(section uint64, index uint32) ([]byte, error) {
	bloomBits, err := e.BloomBits.Get(storageCtx, collections.Join(section, index))
	if err != nil {
		return nil, err
	}

	return bloomBits, nil
}

// RecordBloomBits records the bloom bits for the given index, section and hash.
func (e *EVMIndexerImpl) RecordBloomBits(section uint64, index uint32, bloomBits []byte) error {
	return e.BloomBits.Set(storageCtx, collections.Join(section, index), bloomBits)
}

// NextBloomBitsSection increments the section number.
func (e *EVMIndexerImpl) NextBloomBitsSection() error {
	_, err := e.BloomBitsNextSection.Next(storageCtx)
	return err
}

// PeekBloomBitsNextSection returns the next section number to be processed.
func (e *EVMIndexerImpl) PeekBloomBitsNextSection() (uint64, error) {
	return e.BloomBitsNextSection.Peek(storageCtx)
}

// Check if bloom indexing is running
func (e *EVMIndexerImpl) IsBloomIndexingRunning() bool {
	return e.bloomIndexingRunning.Load()
}
