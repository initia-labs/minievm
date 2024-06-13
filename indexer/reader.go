package indexer

import (
	"context"

	"cosmossdk.io/collections"
	"github.com/ethereum/go-ethereum/common"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	rpctypes "github.com/initia-labs/minievm/jsonrpc/types"
)

// BlockHeaderByHash implements EVMIndexer.
func (e *EVMIndexerImpl) BlockHeaderByHash(ctx context.Context, hash common.Hash) (*coretypes.Header, error) {
	blockNumber, err := e.BlockHashToNumberMap.Get(ctx, hash.Bytes())
	if err != nil {
		return nil, err
	}

	return e.BlockHeaderByNumber(ctx, blockNumber)
}

// BlockHeaderByNumber implements EVMIndexer.
func (e *EVMIndexerImpl) BlockHeaderByNumber(ctx context.Context, blockNumber uint64) (*coretypes.Header, error) {
	blockHeader, err := e.BlockHeaderMap.Get(ctx, blockNumber)
	if err != nil {
		return nil, err
	}

	return &blockHeader, nil
}

// TxByBlockAndIndex implements EVMIndexer.
func (e *EVMIndexerImpl) TxByBlockAndIndex(ctx context.Context, blockHeight uint64, index uint64) (*rpctypes.RPCTransaction, error) {
	txHashBz, err := e.BlockAndIndexToTxHashMap.Get(ctx, collections.Join(blockHeight, index))
	if err != nil {
		return nil, err
	}

	txHash := common.BytesToHash(txHashBz)
	return e.TxByHash(ctx, txHash)
}

// TxByHash implements EVMIndexer.
func (e *EVMIndexerImpl) TxByHash(ctx context.Context, hash common.Hash) (*rpctypes.RPCTransaction, error) {
	tx, err := e.TxMap.Get(ctx, hash.Bytes())
	if err != nil {
		return nil, err
	}

	return &tx, nil
}

// IterateBlockTxs implements EVMIndexer.
func (e *EVMIndexerImpl) IterateBlockTxs(ctx context.Context, blockHeight uint64, cb func(tx *rpctypes.RPCTransaction) (bool, error)) error {
	return e.BlockAndIndexToTxHashMap.Walk(ctx, collections.NewPrefixedPairRange[uint64, uint64](blockHeight), func(key collections.Pair[uint64, uint64], txHashBz []byte) (bool, error) {
		txHash := common.BytesToHash(txHashBz)
		tx, err := e.TxByHash(ctx, txHash)
		if err != nil {
			return true, err
		}

		return cb(tx)
	})
}

// TxReceiptByHash implements EVMIndexer.
func (e *EVMIndexerImpl) TxReceiptByHash(ctx context.Context, hash common.Hash) (*coretypes.Receipt, error) {
	receipt, err := e.TxReceiptMap.Get(ctx, hash.Bytes())
	return &receipt, err
}

// BlockHashToNumber implements EVMIndexer.
func (e *EVMIndexerImpl) BlockHashToNumber(ctx context.Context, hash common.Hash) (uint64, error) {
	return e.BlockHashToNumberMap.Get(ctx, hash.Bytes())
}
