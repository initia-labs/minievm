package indexer

import (
	"encoding/binary"
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	rpctypes "github.com/initia-labs/minievm/jsonrpc/types"
)

// BlockHeaderByHash implements EVMIndexer.
func (e *EVMIndexerImpl) BlockHeaderByHash(hash common.Hash) (*coretypes.Header, error) {
	bz, err := e.db.Get(keyBlockHashToNumber(hash.Bytes()))
	if err != nil {
		return nil, err
	}

	blockNumber := binary.BigEndian.Uint64(bz)
	return e.BlockHeaderByNumber(blockNumber)
}

// BlockHeaderByNumber implements EVMIndexer.
func (e *EVMIndexerImpl) BlockHeaderByNumber(blockNumber uint64) (*coretypes.Header, error) {
	bz, err := e.db.Get(keyBlock(blockNumber))
	if err != nil {
		return nil, err
	}

	var header coretypes.Header
	if err := json.Unmarshal(bz, &header); err != nil {
		return nil, err
	}

	return &header, nil
}

// TxByBlockAndIndex implements EVMIndexer.
func (e *EVMIndexerImpl) TxByBlockAndIndex(blockHeight uint64, index uint64) (*rpctypes.RPCTransaction, error) {
	bz, err := e.db.Get(keyBlockAndIndexToTxHash(blockHeight, index))
	if err != nil {
		return nil, err
	}

	txHash := common.BytesToHash(bz)
	return e.TxByHash(txHash)
}

// TxByHash implements EVMIndexer.
func (e *EVMIndexerImpl) TxByHash(hash common.Hash) (*rpctypes.RPCTransaction, error) {
	bz, err := e.db.Get(keyTx(hash.Bytes()))
	if err != nil {
		return nil, err
	}

	var tx rpctypes.RPCTransaction
	if err := json.Unmarshal(bz, &tx); err != nil {
		return nil, err
	}

	return &tx, nil
}
