package backend

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"

	rpctypes "github.com/initia-labs/minievm/jsonrpc/types"
)

func (b *JSONRPCBackend) BlockNumber() (hexutil.Uint64, error) {
	res, err := b.clientCtx.Client.Status(b.ctx)
	if err != nil {
		return 0, err
	}

	return hexutil.Uint64(res.SyncInfo.LatestBlockHeight), nil
}

func (b *JSONRPCBackend) resolveBlockNrOrHash(blockNrOrHash rpc.BlockNumberOrHash) (uint64, error) {
	if blockHash, ok := blockNrOrHash.Hash(); ok {
		queryCtx, err := b.getQueryCtx()
		if err != nil {
			return 0, err
		}

		return b.app.EVMIndexer().BlockHashToNumber(queryCtx, blockHash)
	} else if blockNumber, ok := blockNrOrHash.Number(); !ok || blockNumber < 0 {
		num, err := b.BlockNumber()
		if err != nil {
			return 0, err
		}

		return uint64(num), nil
	} else {
		return uint64(blockNumber), nil
	}
}

func (b *JSONRPCBackend) resolveBlockNr(blockNr rpc.BlockNumber) (uint64, error) {
	if blockNr < 0 {
		num, err := b.BlockNumber()
		if err != nil {
			return 0, err
		}

		return uint64(num), nil
	} else {
		return uint64(blockNr), nil
	}
}

func (b *JSONRPCBackend) GetHeaderByNumber(ethBlockNum rpc.BlockNumber) (*coretypes.Header, error) {
	blockNumber, err := b.resolveBlockNr(ethBlockNum)
	if err != nil {
		return nil, err
	}

	queryCtx, err := b.getQueryCtx()
	if err != nil {
		return nil, err
	}

	header, err := b.app.EVMIndexer().BlockHeaderByNumber(queryCtx, blockNumber)
	if err != nil {
		return nil, err
	}

	return header, nil

}

func (b *JSONRPCBackend) GetBlockByNumber(ethBlockNum rpc.BlockNumber, fullTx bool) (map[string]interface{}, error) {
	blockNumber, err := b.resolveBlockNr(ethBlockNum)
	if err != nil {
		return nil, err
	}

	queryCtx, err := b.getQueryCtx()
	if err != nil {
		return nil, err
	}

	header, err := b.app.EVMIndexer().BlockHeaderByNumber(queryCtx, blockNumber)
	if err != nil {
		return nil, err
	}

	txs := []*rpctypes.RPCTransaction{}
	if fullTx {
		b.app.EVMIndexer().IterateBlockTxs(queryCtx, blockNumber, func(tx *rpctypes.RPCTransaction) (bool, error) {
			txs = append(txs, tx)
			return false, nil
		})
	}

	return formatBlock(header, txs), nil
}

func (b *JSONRPCBackend) GetBlockByHash(hash common.Hash, fullTx bool) (map[string]interface{}, error) {
	queryCtx, err := b.getQueryCtx()
	if err != nil {
		return nil, err
	}

	header, err := b.app.EVMIndexer().BlockHeaderByHash(queryCtx, hash)
	if err != nil {
		return nil, err
	}

	txs := []*rpctypes.RPCTransaction{}
	if fullTx {
		blockNumber := header.Number.Uint64()
		b.app.EVMIndexer().IterateBlockTxs(queryCtx, blockNumber, func(tx *rpctypes.RPCTransaction) (bool, error) {
			txs = append(txs, tx)
			return false, nil
		})
	}

	return formatBlock(header, txs), nil
}

func formatBlock(header *coretypes.Header, txs []*rpctypes.RPCTransaction) map[string]interface{} {
	fields := formatHeader(header)
	fields["transactions"] = txs

	// empty values
	fields["size"] = hexutil.Uint64(0)
	fields["uncles"] = []common.Hash{}
	fields["withdrawals"] = coretypes.Withdrawals{}

	return fields
}

// formatHeader converts the given header to the RPC output .
func formatHeader(head *coretypes.Header) map[string]interface{} {
	result := map[string]interface{}{
		"number":           (*hexutil.Big)(head.Number),
		"hash":             head.Hash(),
		"parentHash":       head.ParentHash,
		"nonce":            head.Nonce,
		"mixHash":          head.MixDigest,
		"sha3Uncles":       head.UncleHash,
		"logsBloom":        head.Bloom,
		"stateRoot":        head.Root,
		"miner":            head.Coinbase,
		"difficulty":       (*hexutil.Big)(head.Difficulty),
		"extraData":        hexutil.Bytes(head.Extra),
		"gasLimit":         hexutil.Uint64(head.GasLimit),
		"gasUsed":          hexutil.Uint64(head.GasUsed),
		"timestamp":        hexutil.Uint64(head.Time),
		"transactionsRoot": head.TxHash,
		"receiptsRoot":     head.ReceiptHash,
	}
	if head.BaseFee != nil {
		result["baseFeePerGas"] = (*hexutil.Big)(head.BaseFee)
	}
	if head.WithdrawalsHash != nil {
		result["withdrawalsRoot"] = head.WithdrawalsHash
	}
	if head.BlobGasUsed != nil {
		result["blobGasUsed"] = hexutil.Uint64(*head.BlobGasUsed)
	}
	if head.ExcessBlobGas != nil {
		result["excessBlobGas"] = hexutil.Uint64(*head.ExcessBlobGas)
	}
	if head.ParentBeaconRoot != nil {
		result["parentBeaconBlockRoot"] = head.ParentBeaconRoot
	}
	return result
}
