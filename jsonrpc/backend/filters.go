package backend

import (
	"github.com/ethereum/go-ethereum/common"
	coretypes "github.com/ethereum/go-ethereum/core/types"

	rpctypes "github.com/initia-labs/minievm/jsonrpc/types"
)

// GetLogs returns all the logs from all the ethereum transactions in a block.
func (b *JSONRPCBackend) GetLogs(hash common.Hash) ([][]*coretypes.Log, error) {

	queryCtx, err := b.getQueryCtx()
	if err != nil {
		return nil, err
	}

	height, err := b.app.EVMIndexer().BlockHashToNumber(queryCtx, hash)
	if err != nil {
		return nil, err
	}

	h := int64(height)
	return b.GetLogsByHeight(&h)
}

// GetLogsByHeight returns all the logs from all the ethereum transactions in a block.
func (b *JSONRPCBackend) GetLogsByHeight(height *int64) ([][]*coretypes.Log, error) {

	blockLogs := [][]*coretypes.Log{}
	blockNumber := uint64(*height)

	queryCtx, err := b.getQueryCtx()
	if err != nil {
		return nil, err
	}

	blockHeader, err := b.app.EVMIndexer().BlockHeaderByNumber(queryCtx, blockNumber)
	if err != nil {
		return nil, err
	}

	txs := []*rpctypes.RPCTransaction{}
	b.app.EVMIndexer().IterateBlockTxs(queryCtx, blockNumber, func(tx *rpctypes.RPCTransaction) (bool, error) {
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
			log.BlockNumber = blockNumber
			log.TxHash = tx.Hash
			log.Index = uint(idx)
			log.TxIndex = receipt.TransactionIndex
		}
		blockLogs = append(blockLogs, logs)
	}

	return blockLogs, nil
}

/*
func (b *JSONRPCBackend) GetFilterLogs(ctx context.Context, id rpc.ID) ([]*coretypes.Log, error) {
	blockLogs := []*coretypes.Log{}
	return blockLogs, nil

}
*/
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
