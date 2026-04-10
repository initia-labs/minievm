package indexer

import (
	"cosmossdk.io/collections"
	"github.com/ethereum/go-ethereum/common"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	rpctypes "github.com/initia-labs/minievm/jsonrpc/types"
)

// BlockHeaderByNumber implements EVMIndexer.
func (e *EVMIndexerImpl) BlockHeaderByNumber(blockNumber uint64) (*coretypes.Header, error) {
	blockHeader, err := e.BlockHeaderMap.Get(storageCtx, blockNumber)
	if err != nil {
		return nil, err
	}

	return &blockHeader, nil
}

// TxHashByBlockAndIndex implements EVMIndexer.
func (e *EVMIndexerImpl) TxHashByBlockAndIndex(blockHeight uint64, index uint64) (common.Hash, error) {
	txHashBz, err := e.BlockAndIndexToTxHashMap.Get(storageCtx, collections.Join(blockHeight, index))
	if err != nil {
		return common.Hash{}, err
	}

	return common.BytesToHash(txHashBz), nil
}

// TxByHash implements EVMIndexer.
func (e *EVMIndexerImpl) TxByHash(hash common.Hash) (*rpctypes.RPCTransaction, error) {
	tx, err := e.TxMap.Get(storageCtx, hash.Bytes())
	if err != nil {
		return nil, err
	}

	return &tx, nil
}

// IterateBlockTxs implements EVMIndexer.
func (e *EVMIndexerImpl) IterateBlockTxs(blockHeight uint64, cb func(tx *rpctypes.RPCTransaction) (bool, error)) error {
	return e.BlockAndIndexToTxHashMap.Walk(storageCtx, collections.NewPrefixedPairRange[uint64, uint64](blockHeight), func(key collections.Pair[uint64, uint64], txHashBz []byte) (bool, error) {
		txHash := common.BytesToHash(txHashBz)
		tx, err := e.TxByHash(txHash)
		if err != nil {
			return true, err
		}

		return cb(tx)
	})
}

// IterateBlockTxReceipts implements EVMIndexer.
func (e *EVMIndexerImpl) IterateBlockTxReceipts(blockHeight uint64, cb func(tx *coretypes.Receipt) (bool, error)) error {
	return e.BlockAndIndexToTxHashMap.Walk(storageCtx, collections.NewPrefixedPairRange[uint64, uint64](blockHeight), func(key collections.Pair[uint64, uint64], txHashBz []byte) (bool, error) {
		txHash := common.BytesToHash(txHashBz)
		txReceipt, err := e.TxReceiptByHash(txHash)
		if err != nil {
			return true, err
		}

		return cb(txReceipt)
	})
}

// TxReceiptByHash implements EVMIndexer.
func (e *EVMIndexerImpl) TxReceiptByHash(hash common.Hash) (*coretypes.Receipt, error) {
	receipt, err := e.TxReceiptMap.Get(storageCtx, hash.Bytes())
	return &receipt, err
}

// TxStartLogIndexByHash implements EVMIndexer.
// Returns the block-scoped index of the first log for the given tx.
// Returns collections.ErrNotFound if not stored (e.g. indexed before this field was introduced).
func (e *EVMIndexerImpl) TxStartLogIndexByHash(hash common.Hash) (uint64, error) {
	return e.TxStartLogIndexMap.Get(storageCtx, hash.Bytes())
}

// StoreTxStartLogIndex implements EVMIndexer.
func (e *EVMIndexerImpl) StoreTxStartLogIndex(hash common.Hash, index uint64) error {
	return e.TxStartLogIndexMap.Set(storageCtx, hash.Bytes(), index)
}

// BlockHashToNumber implements EVMIndexer.
func (e *EVMIndexerImpl) BlockHashToNumber(hash common.Hash) (uint64, error) {
	return e.BlockHashToNumberMap.Get(storageCtx, hash.Bytes())
}

// CosmosTxHashByTxHash implements EVMIndexer.
func (e *EVMIndexerImpl) CosmosTxHashByTxHash(hash common.Hash) ([]byte, error) {
	return e.TxHashToCosmosTxHash.Get(storageCtx, hash.Bytes())
}

// TxHashByCosmosTxHash implements EVMIndexer.
func (e *EVMIndexerImpl) TxHashByCosmosTxHash(hash []byte) (common.Hash, error) {
	bz, err := e.CosmosTxHashToTxHash.Get(storageCtx, hash)
	if err != nil {
		return common.Hash{}, err
	}

	return common.BytesToHash(bz), nil
}
