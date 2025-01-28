package indexer

import (
	"context"
	"math/big"

	"cosmossdk.io/collections"
	"github.com/ethereum/go-ethereum/common"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	rpctypes "github.com/initia-labs/minievm/jsonrpc/types"
)

// BlockHeaderByNumber implements EVMIndexer.
func (e *EVMIndexerImpl) BlockHeaderByNumber(ctx context.Context, blockNumber uint64) (*coretypes.Header, error) {
	if blockNumber == 0 {
		// this is a special case for genesis block
		return &coretypes.Header{
			Number: big.NewInt(0),
			Bloom:  coretypes.Bloom{},
		}, nil
	}

	blockHeader, err := e.BlockHeaderMap.Get(ctx, blockNumber)
	if err != nil {
		return nil, err
	}

	return &blockHeader, nil
}

// TxHashByBlockAndIndex implements EVMIndexer.
func (e *EVMIndexerImpl) TxHashByBlockAndIndex(ctx context.Context, blockHeight uint64, index uint64) (common.Hash, error) {
	txHashBz, err := e.BlockAndIndexToTxHashMap.Get(ctx, collections.Join(blockHeight, index))
	if err != nil {
		return common.Hash{}, err
	}

	return common.BytesToHash(txHashBz), nil
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

// IterateBlockTxs implements EVMIndexer.
func (e *EVMIndexerImpl) IterateBlockTxReceipts(ctx context.Context, blockHeight uint64, cb func(tx *coretypes.Receipt) (bool, error)) error {
	return e.BlockAndIndexToTxHashMap.Walk(ctx, collections.NewPrefixedPairRange[uint64, uint64](blockHeight), func(key collections.Pair[uint64, uint64], txHashBz []byte) (bool, error) {
		txHash := common.BytesToHash(txHashBz)
		txRecept, err := e.TxReceiptByHash(ctx, txHash)
		if err != nil {
			return true, err
		}

		return cb(txRecept)
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

// CosmosTxHashByTxHash implements EVMIndexer.
func (e *EVMIndexerImpl) CosmosTxHashByTxHash(ctx context.Context, hash common.Hash) ([]byte, error) {
	return e.TxHashToCosmosTxHash.Get(ctx, hash.Bytes())
}

// TxHashByCosmosTxHash implements EVMIndexer.
func (e *EVMIndexerImpl) TxHashByCosmosTxHash(ctx context.Context, hash []byte) (common.Hash, error) {
	bz, err := e.CosmosTxHashToTxHash.Get(ctx, hash)
	if err != nil {
		return common.Hash{}, err
	}

	return common.BytesToHash(bz), nil
}
