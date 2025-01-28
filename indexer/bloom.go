package indexer

import (
	"context"

	"cosmossdk.io/collections"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/bloombits"
	coretypes "github.com/ethereum/go-ethereum/core/types"

	evmconfig "github.com/initia-labs/minievm/x/evm/config"
)

// doBloomIndexing triggers bloom indexing in a goroutine. If bloom indexing is already running,
// it does nothing.
func (e *EVMIndexerImpl) doBloomIndexing(ctx context.Context, height uint64) {
	if running := e.bloomIndexingRunning.Swap(true); running {
		return
	}

	go func(ctx context.Context, height uint64) {
		defer e.bloomIndexingRunning.Store(false)
		if err := e.bloomIndexing(ctx, height); err != nil {
			e.logger.Error("failed to do bloom indexing", "err", err)
		}

		e.logger.Debug("bloom indexing finished", "height", height)
	}(ctx, height)
}

// bloomIndexing generates the bloom index if the current section is complete.
func (e *EVMIndexerImpl) bloomIndexing(ctx context.Context, height uint64) error {
	section, err := e.PeekBloomBitsNextSection(ctx)
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

	lastHead := common.Hash{}
	for i := uint64(0); i < evmconfig.SectionSize; i++ {
		header, err := e.BlockHeaderByNumber(ctx, section*evmconfig.SectionSize+i)
		if err != nil {
			return err
		}

		if err := gen.AddBloom(uint(header.Number.Uint64()-section*evmconfig.SectionSize), header.Bloom); err != nil {
			return err
		}

		lastHead = header.Hash()
	}

	// write the bloom bits to the store
	for i := 0; i < coretypes.BloomBitLength; i++ {
		bits, err := gen.Bitset(uint(i))
		if err != nil {
			return err
		}

		e.RecordBloomBits(ctx, section, uint32(i), lastHead, bits)
	}

	// increment the section number; if this fails, the section will be reprocessed
	if err := e.NextBloomBitsSection(ctx); err != nil {
		return err
	}

	return nil
}

// ReadBloomBits reads the bloom bits for the given index, section and hash.
func (e *EVMIndexerImpl) ReadBloomBits(ctx context.Context, section uint64, index uint32, hash common.Hash) ([]byte, error) {
	bloomBits, err := e.BloomBits.Get(ctx, collections.Join3(section, index, hash.Bytes()))
	if err != nil {
		return nil, err
	}

	return bloomBits, nil
}

// RecordBloomBits records the bloom bits for the given index, section and hash.
func (e *EVMIndexerImpl) RecordBloomBits(ctx context.Context, section uint64, index uint32, hash common.Hash, bloomBits []byte) error {
	return e.BloomBits.Set(ctx, collections.Join3(section, index, hash.Bytes()), bloomBits)
}

// NextBloomBitsSection increments the section number.
func (e *EVMIndexerImpl) NextBloomBitsSection(ctx context.Context) error {
	_, err := e.BloomBitsNextSection.Next(ctx)
	return err
}

// PeekBloomBitsNextSection returns the next section number to be processed.
func (e *EVMIndexerImpl) PeekBloomBitsNextSection(ctx context.Context) (uint64, error) {
	return e.BloomBitsNextSection.Peek(ctx)
}

// Check if bloom indexing is running
func (e *EVMIndexerImpl) IsBloomIndexingRunning() bool {
	return e.bloomIndexingRunning.Load()
}
