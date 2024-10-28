package backend

import (
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
)

// GetLogsByHeight returns all the logs from all the ethereum transactions in a block.
func (b *JSONRPCBackend) GetLogsByHeight(height uint64) ([]*coretypes.Log, error) {
	if blockLogs, ok := b.logsCache.Get(height); ok {
		return blockLogs, nil
	}

	blockHeader, err := b.GetHeaderByNumber(rpc.BlockNumber(height))
	if err != nil {
		return nil, err
	} else if blockHeader == nil {
		return nil, nil
	}

	txs, err := b.getBlockTransactions(height)
	if err != nil {
		return nil, err
	}
	receipts, err := b.getBlockReceipts(height)
	if err != nil {
		return nil, err
	}
	if len(txs) != len(receipts) {
		return nil, NewInternalError("mismatched number of transactions and receipts")
	}

	blockLogs := []*coretypes.Log{}
	for i, tx := range txs {
		receipt := receipts[i]
		logs := receipt.Logs
		for idx, log := range logs {
			log.BlockHash = blockHeader.Hash()
			log.BlockNumber = height
			log.TxHash = tx.Hash
			log.Index = uint(idx)
			log.TxIndex = receipt.TransactionIndex
		}
		blockLogs = append(blockLogs, logs...)
	}

	// cache the logs
	_ = b.logsCache.Add(height, blockLogs)
	return blockLogs, nil
}
