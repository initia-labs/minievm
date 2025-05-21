package backend

import (
	"errors"

	"cosmossdk.io/collections"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"

	rpctypes "github.com/initia-labs/minievm/jsonrpc/types"
)

func (b *JSONRPCBackend) BlockNumber() (hexutil.Uint64, error) {
	lh, err := b.app.EVMIndexer().GetLastIndexedHeight(b.ctx)
	if err != nil {
		return 0, err
	}
	return hexutil.Uint64(lh), nil
}

func (b *JSONRPCBackend) resolveBlockNrOrHash(blockNrOrHash rpc.BlockNumberOrHash) (uint64, error) {
	if blockHash, ok := blockNrOrHash.Hash(); ok {
		return b.blockNumberByHash(blockHash)
	} else if blockNumber, ok := blockNrOrHash.Number(); !ok || blockNumber < 0 {
		num, err := b.BlockNumber()
		if err != nil {
			return 0, err
		}

		return uint64(num), nil
	} else if blockNumber == 0 {
		return uint64(1), nil
	} else {
		return uint64(blockNumber), nil
	}
}

func (b *JSONRPCBackend) resolveBlockNr(blockNr rpc.BlockNumber) (uint64, error) {
	if blockNr < rpc.BlockNumber(0) {
		num, err := b.BlockNumber()
		if err != nil {
			return 0, err
		}

		return uint64(num), nil
	} else if blockNr == rpc.BlockNumber(0) {
		return uint64(1), nil
	} else {
		return uint64(blockNr), nil
	}
}

func (b *JSONRPCBackend) GetHeaderByNumber(ethBlockNum rpc.BlockNumber) (*coretypes.Header, error) {
	blockNumber, err := b.resolveBlockNr(ethBlockNum)
	if err != nil {
		return nil, err
	}

	if indexed, err := b.isBlockIndexed(blockNumber); err != nil || !indexed {
		return nil, err
	}

	if header, ok := b.headerCache.Get(blockNumber); ok {
		return header, nil
	}

	queryCtx, err := b.getQueryCtx()
	if err != nil {
		return nil, err
	}

	header, err := b.app.EVMIndexer().BlockHeaderByNumber(queryCtx, blockNumber)
	if err != nil && errors.Is(err, collections.ErrNotFound) {
		return nil, nil
	} else if err != nil {
		b.logger.Error("failed to get block header by number", "err", err)
		return nil, NewInternalError("failed to get block header by number")
	}

	// cache the header
	_ = b.headerCache.Add(blockNumber, header)
	return header, nil
}

func (b *JSONRPCBackend) GetHeaderByHash(hash common.Hash) (*coretypes.Header, error) {
	blockNumber, err := b.resolveBlockNrOrHash(rpc.BlockNumberOrHash{BlockHash: &hash})
	if err != nil && errors.Is(err, collections.ErrNotFound) {
		return nil, nil
	} else if err != nil {
		b.logger.Error("failed to get block number by hash", "err", err)
		return nil, err
	}

	return b.GetHeaderByNumber(rpc.BlockNumber(blockNumber))
}

func (b *JSONRPCBackend) GetBlockByNumber(ethBlockNum rpc.BlockNumber, fullTx bool) (map[string]interface{}, error) {
	blockNumber, err := b.resolveBlockNr(ethBlockNum)
	if err != nil {
		return nil, err
	}

	header, err := b.GetHeaderByNumber(ethBlockNum)
	if err != nil {
		b.logger.Error("failed to get block header by number", "err", err)
		return nil, err
	}
	if header == nil {
		return nil, nil
	}

	txs, err := b.getBlockTransactions(blockNumber)
	if err != nil {
		return nil, err
	}

	return formatBlock(header, txs, fullTx), nil
}

func (b *JSONRPCBackend) GetBlockByHash(hash common.Hash, fullTx bool) (map[string]interface{}, error) {
	blockNumber, err := b.resolveBlockNrOrHash(rpc.BlockNumberOrHash{BlockHash: &hash})
	if err != nil && errors.Is(err, collections.ErrNotFound) {
		return nil, nil
	} else if err != nil {
		b.logger.Error("failed to get block number by hash", "err", err)
		return nil, err
	}
	return b.GetBlockByNumber(rpc.BlockNumber(blockNumber), fullTx)
}

func (b *JSONRPCBackend) blockNumberByHash(hash common.Hash) (uint64, error) {
	if number, ok := b.blockHashCache.Get(hash); ok {
		return number, nil
	}

	queryCtx, err := b.getQueryCtx()
	if err != nil {
		return 0, err
	}

	number, err := b.app.EVMIndexer().BlockHashToNumber(queryCtx, hash)
	if err != nil {
		b.logger.Error("failed to get block number by hash", "err", err)
		return 0, NewInternalError("failed to get block number by hash")
	}

	_ = b.blockHashCache.Add(hash, number)
	return number, nil
}

func formatBlock(header *coretypes.Header, txs []*rpctypes.RPCTransaction, fullTx bool) map[string]interface{} {
	fields := formatHeader(header)

	if fullTx {
		fields["transactions"] = txs
	} else {
		txsHashes := []common.Hash{}
		for _, tx := range txs {
			txsHashes = append(txsHashes, tx.Hash)
		}
		fields["transactions"] = txsHashes
	}

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
