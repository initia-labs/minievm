package backend

import (
	coretypes "github.com/ethereum/go-ethereum/core/types"

	rpctypes "github.com/initia-labs/minievm/jsonrpc/types"
)

// GetLogsByHeight returns all the logs from all the ethereum transactions in a block.
func (b *JSONRPCBackend) GetLogsByHeight(height uint64) ([][]*coretypes.Log, error) {

	blockLogs := [][]*coretypes.Log{}

	queryCtx, err := b.getQueryCtx()
	if err != nil {
		return nil, err
	}

	blockHeader, err := b.app.EVMIndexer().BlockHeaderByNumber(queryCtx, height)
	if err != nil {
		return nil, err
	}

	txs := []*rpctypes.RPCTransaction{}
	b.app.EVMIndexer().IterateBlockTxs(queryCtx, height, func(tx *rpctypes.RPCTransaction) (bool, error) {
		txs = append(txs, tx)
		return false, nil
	})

	for _, tx := range txs {
		receipt, err := b.app.EVMIndexer().TxReceiptByHash(queryCtx, tx.Hash)
		if err != nil {
			return nil, err
		}
		logs := receipt.Logs
		for idx, log := range logs {
			log.BlockHash = blockHeader.Hash()
			log.BlockNumber = height
			log.TxHash = tx.Hash
			log.Index = uint(idx)
			log.TxIndex = receipt.TransactionIndex
		}
		blockLogs = append(blockLogs, logs)
	}

	return blockLogs, nil
}

// RPCFilterCap is the limit for total number of filters that can be created
func (b *JSONRPCBackend) RPCFilterCap() int32 {
	return b.cfg.FilterCap
}

// RPCFilterCap is the limit for total number of filters that can be created
func (b *JSONRPCBackend) RPCLogsCap() int32 {
	return b.cfg.FilterCap
}

// RPCFilterCap is the limit for total number of filters that can be created
func (b *JSONRPCBackend) RPCBlockRangeCap() int32 {
	return b.cfg.FilterCap
}
