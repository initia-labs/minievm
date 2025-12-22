package backend

import (
	"context"
	"errors"
	"io"
	"time"

	"cosmossdk.io/collections"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"

	rpctypes "github.com/initia-labs/minievm/jsonrpc/types"
	"github.com/initia-labs/minievm/x/evm/keeper"
	"github.com/initia-labs/minievm/x/evm/types"

	cmtrpcclient "github.com/cometbft/cometbft/rpc/client"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (b *JSONRPCBackend) SendRawTransaction(input hexutil.Bytes) (common.Hash, error) {
	tx := new(coretypes.Transaction)
	if err := tx.UnmarshalBinary(input); err != nil {
		return common.Hash{}, err
	}

	if !tx.Protected() {
		// Ensure only eip155 signed transactions are submitted if EIP155Required is set.
		return common.Hash{}, errors.New("only replay-protected (EIP-155) transactions allowed over RPC")
	}

	if err := b.SendTx(tx); err != nil {
		return common.Hash{}, err
	}

	return tx.Hash(), nil
}

// SendRawTransactionSync send a raw Ethereum transaction and wait for the result synchronously.
func (b *JSONRPCBackend) SendRawTransactionSync(input hexutil.Bytes, timeoutInMS int64) (map[string]any, error) {
	timeoutInMS = min(timeoutInMS, b.cfg.HTTPTimeout.Milliseconds())

	txHash, err := b.SendRawTransaction(input)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutInMS)*time.Millisecond)
	defer cancel()
	timer := time.NewTicker(100 * time.Millisecond)
	defer timer.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil, NewTimeoutError("transaction not mined within the specified timeout")
		case <-timer.C:
			receipt, err := b.GetTransactionReceipt(txHash)
			if err != nil {
				return nil, err
			}
			if receipt != nil {
				return receipt, nil
			}
		}
	}
}

func (b *JSONRPCBackend) SendTx(tx *coretypes.Transaction) error {
	queryCtx, closer, err := b.getQueryCtx()
	if closer != nil {
		defer closer.Close()
	}
	if err != nil {
		return NewReadinessError(err.Error())
	}

	cosmosTx, err := keeper.NewTxUtils(b.app.EVMKeeper).ConvertEthereumTxToCosmosTx(queryCtx, tx)
	if err != nil {
		return err
	}

	txBytes, err := b.app.TxEncode(cosmosTx)
	if err != nil {
		return err
	}

	res, err := b.clientCtx.BroadcastTxSync(txBytes)
	if err != nil {
		return err
	} else if res.Code != 0 {
		return errors.New(res.RawLog)
	}

	return nil
}

// getQueryCtx returns a query context for the current block height.
// This function should only be used when interacting with keepers, as it creates a context specifically configured for keeper queries.
func (b *JSONRPCBackend) getQueryCtx() (context.Context, io.Closer, error) {
	return b.app.CreateQueryContext(0, false)
}

// getQueryCtxWithHeight returns a query context for the given block height.
// This function should only be used when interacting with keepers, as it creates a context specifically configured for keeper queries.
func (b *JSONRPCBackend) getQueryCtxWithHeight(height uint64) (context.Context, io.Closer, error) {
	// check whether the given height is bigger than the latest block height
	num, err := b.BlockNumber()
	if err != nil {
		return nil, nil, err
	}
	if height > uint64(num) {
		return nil, nil, errors.New("requested height is greater than the latest block height")
	}
	if height == uint64(num) {
		height = 0
	}

	return b.app.CreateQueryContext(int64(height), false)
}

// GetTransactionByHash returns the transaction with the given hash.
func (b *JSONRPCBackend) GetTransactionByHash(hash common.Hash) (*rpctypes.RPCTransaction, error) {
	return b.getTransaction(hash)
}

// GetTransactionCount returns the number of transactions at the given block number.
func (b *JSONRPCBackend) GetTransactionCount(address common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*hexutil.Uint64, error) {
	sdkAddr := sdk.AccAddress(address[:])

	var blockNumber rpc.BlockNumber
	if blockHash, ok := blockNrOrHash.Hash(); ok {
		blockNumberU64, err := b.blockNumberByHash(blockHash)
		if err != nil && errors.Is(err, collections.ErrNotFound) {
			return nil, nil
		} else if err != nil {
			b.logger.Error("failed to get block number by hash", "err", err)
			return nil, err
		}

		blockNumber = rpc.BlockNumber(blockNumberU64)
	} else {
		blockNumber, _ = blockNrOrHash.Number()
	}

	seq := uint64(0)
	var queryCtx context.Context
	if blockNumber == rpc.PendingBlockNumber {
		queryCtx = b.app.GetContextForCheckTx(nil)
	} else {
		if blockNumber < 0 {
			blockNumber = 0
		}

		var err error
		var closer io.Closer
		queryCtx, closer, err = b.getQueryCtxWithHeight(uint64(blockNumber.Int64()))
		if closer != nil {
			defer closer.Close()
		}
		if err != nil {
			return nil, err
		}
	}

	acc := b.app.AccountKeeper.GetAccount(queryCtx, sdkAddr)
	if acc != nil {
		seq = acc.GetSequence()
	}

	return (*hexutil.Uint64)(&seq), nil
}

// GetTransactionReceipt returns the transaction receipt for the given transaction hash.
func (b *JSONRPCBackend) GetTransactionReceipt(hash common.Hash) (map[string]any, error) {
	rpcTx, err := b.getTransaction(hash)
	if err != nil {
		return nil, err
	} else if rpcTx == nil || rpcTx.BlockNumber == nil || rpcTx.BlockNumber.ToInt() == nil {
		return nil, nil // tx is not found or in pending/queued state
	}

	if indexed, err := b.isBlockIndexed(rpcTx.BlockNumber.ToInt().Uint64()); err != nil || !indexed {
		return nil, err
	}

	receipt, err := b.getReceipt(hash)
	if err != nil {
		return nil, err
	} else if receipt == nil {
		return nil, nil
	}

	return marshalReceipt(receipt, rpcTx), nil
}

// GetTransactionByBlockHashAndIndex returns the transaction at the given block hash and index.
func (b *JSONRPCBackend) GetTransactionByBlockHashAndIndex(hash common.Hash, idx hexutil.Uint) (*rpctypes.RPCTransaction, error) {
	blockNumber, err := b.resolveBlockNrOrHash(rpc.BlockNumberOrHash{BlockHash: &hash})
	if err != nil && errors.Is(err, collections.ErrNotFound) {
		return nil, nil
	} else if err != nil {
		b.logger.Error("failed to get block number by hash", "err", err)
		return nil, err
	}

	return b.GetTransactionByBlockNumberAndIndex(rpc.BlockNumber(blockNumber), idx)
}

// GetTransactionByBlockNumberAndIndex returns the transaction at the given block number and index.
func (b *JSONRPCBackend) GetTransactionByBlockNumberAndIndex(blockNum rpc.BlockNumber, idx hexutil.Uint) (*rpctypes.RPCTransaction, error) {
	blockNumber, err := b.resolveBlockNr(blockNum)
	if err != nil {
		return nil, err
	}
	if txs, ok := b.blockTxsCache.Get(blockNumber); ok {
		if int(idx) >= len(txs) {
			return nil, nil
		}

		return txs[idx], nil
	}

	if indexed, err := b.isBlockIndexed(blockNumber); err != nil || !indexed {
		return nil, err
	}

	txhash, err := b.app.EVMIndexer().TxHashByBlockAndIndex(b.ctx, blockNumber, uint64(idx))
	if err != nil && errors.Is(err, collections.ErrNotFound) {
		return nil, nil
	} else if err != nil {
		b.logger.Error("failed to get transaction by block and index", "err", err)
		return nil, NewInternalError("failed to get transaction by block and index")
	}

	return b.getTransaction(txhash)
}

// GetBlockTransactionCountByHash returns the number of transactions in a block from a block matching the given block hash.
func (b *JSONRPCBackend) GetBlockTransactionCountByHash(hash common.Hash) (*hexutil.Uint, error) {
	block, err := b.GetBlockByHash(hash, true)
	if err != nil {
		return nil, err
	} else if block == nil {
		return nil, nil
	}

	numTxs := hexutil.Uint(len(block["transactions"].([]*rpctypes.RPCTransaction)))
	return &numTxs, nil
}

// GetBlockTransactionCountByNumber returns the number of transactions in a block from a block matching the given block number.
func (b *JSONRPCBackend) GetBlockTransactionCountByNumber(blockNum rpc.BlockNumber) (*hexutil.Uint, error) {
	block, err := b.GetBlockByNumber(blockNum, true)
	if err != nil {
		return nil, err
	} else if block == nil {
		return nil, nil
	}

	numTxs := hexutil.Uint(len(block["transactions"].([]*rpctypes.RPCTransaction)))
	return &numTxs, nil
}

// GetRawTransactionByHash returns the bytes of the transaction for the given hash.
func (b *JSONRPCBackend) GetRawTransactionByHash(hash common.Hash) (hexutil.Bytes, error) {
	rpcTx, err := b.getTransaction(hash)
	if rpcTx == nil {
		return nil, err
	}

	return rpcTx.ToTransaction().MarshalBinary()
}

// GetRawTransactionByBlockHashAndIndex returns the bytes of the transaction for the given block hash and index.
func (b *JSONRPCBackend) GetRawTransactionByBlockHashAndIndex(blockHash common.Hash, index hexutil.Uint) (hexutil.Bytes, error) {
	rpcTx, err := b.GetTransactionByBlockHashAndIndex(blockHash, index)
	if err != nil {
		return nil, err
	} else if rpcTx == nil {
		return nil, nil
	}

	return rpcTx.ToTransaction().MarshalBinary()
}

func (b *JSONRPCBackend) PendingTransactions() ([]*rpctypes.RPCTransaction, error) {
	queryCtx, closer, err := b.getQueryCtx()
	if closer != nil {
		defer closer.Close()
	}
	if err != nil {
		return nil, err
	}

	mc, ok := b.clientCtx.Client.(cmtrpcclient.MempoolClient)
	if !ok {
		return nil, errors.New("mempool client not available")
	}

	res, err := mc.UnconfirmedTxs(b.ctx, nil)
	if err != nil {
		return nil, err
	}

	result := make([]*rpctypes.RPCTransaction, 0, len(res.Txs))
	for _, txBz := range res.Txs {
		tx, err := b.clientCtx.TxConfig.TxDecoder()(txBz)
		if err != nil {
			return nil, err
		}

		sdkCtx := sdk.UnwrapSDKContext(queryCtx)
		ethTx, _, err := keeper.NewTxUtils(b.app.EVMKeeper).ConvertCosmosTxToEthereumTx(sdkCtx, tx)
		if err != nil {
			return nil, err
		}
		if ethTx != nil {
			result = append(
				result,
				rpctypes.NewRPCTransaction(ethTx, common.Hash{}, 0, 0, ethTx.ChainId()),
			)
		}
	}

	return result, nil
}

func (b *JSONRPCBackend) GetBlockReceipts(blockNrOrHash rpc.BlockNumberOrHash) ([]map[string]any, error) {
	blockNumber, err := b.resolveBlockNrOrHash(blockNrOrHash)
	if err != nil && errors.Is(err, collections.ErrNotFound) {
		return nil, nil
	} else if err != nil {
		b.logger.Error("failed to get block number by hash", "err", err)
		return nil, err
	}

	if indexed, err := b.isBlockIndexed(blockNumber); err != nil || !indexed {
		return nil, err
	}

	txs, err := b.getBlockTransactions(blockNumber)
	if err != nil {
		return nil, err
	}
	receipts, err := b.getBlockReceipts(blockNumber)
	if err != nil {
		return nil, err
	}
	if len(txs) != len(receipts) {
		// something is wrong, clear the cache
		b.blockTxsCache.Remove(blockNumber)
		b.blockReceiptsCache.Remove(blockNumber)
		b.logger.Error("mismatched number of transactions and receipts", "height", blockNumber)

		return nil, NewInternalError("mismatched number of transactions and receipts")
	}

	result := make([]map[string]any, len(receipts))
	for i, receipt := range receipts {
		result[i] = marshalReceipt(receipt, txs[i])
	}

	return result, nil
}

// getTransaction retrieves the lookup along with the transaction itself associate
// with the given transaction hash.
func (b *JSONRPCBackend) getTransaction(hash common.Hash) (*rpctypes.RPCTransaction, error) {
	if tx, ok := b.txLookupCache.Get(hash); ok {
		return tx, nil
	}

	// check if the transaction is in the queued txs
	if tx := b.app.EVMIndexer().TxInQueued(hash); tx != nil {
		return tx, nil
	}

	// check if the transaction is in the pending txs
	if tx := b.app.EVMIndexer().TxInPending(hash); tx != nil {
		return tx, nil
	}

	tx, err := b.app.EVMIndexer().TxByHash(b.ctx, hash)
	if err != nil && errors.Is(err, collections.ErrNotFound) {
		return nil, nil
	} else if err != nil {
		b.logger.Error("failed to get transaction by hash", "err", err)
		return nil, NewInternalError("failed to get transaction by hash")
	}

	_ = b.txLookupCache.Add(hash, tx)
	return tx, nil
}

func (b *JSONRPCBackend) getReceipt(hash common.Hash) (*coretypes.Receipt, error) {
	if receipt, ok := b.receiptCache.Get(hash); ok {
		return receipt, nil
	}

	receipt, err := b.app.EVMIndexer().TxReceiptByHash(b.ctx, hash)
	if err != nil && errors.Is(err, collections.ErrNotFound) {
		return nil, nil
	} else if err != nil {
		b.logger.Error("failed to get transaction receipt by hash", "err", err)
		return nil, NewInternalError("failed to get transaction receipt by hash")
	}

	_ = b.receiptCache.Add(hash, receipt)
	return receipt, nil
}

func (b *JSONRPCBackend) getBlockTransactions(blockNumber uint64) ([]*rpctypes.RPCTransaction, error) {
	if txs, ok := b.blockTxsCache.Get(blockNumber); ok {
		return txs, nil
	}

	txs := []*rpctypes.RPCTransaction{}
	err := b.app.EVMIndexer().IterateBlockTxs(b.ctx, blockNumber, func(tx *rpctypes.RPCTransaction) (bool, error) {
		txs = append(txs, tx)
		return false, nil
	})
	if err != nil {
		b.logger.Error("failed to get block transactions", "err", err)
		return nil, NewInternalError("failed to get block transactions")
	}

	// cache the transactions
	_ = b.blockTxsCache.Add(blockNumber, txs)
	return txs, nil
}

func (b *JSONRPCBackend) getBlockReceipts(blockNumber uint64) ([]*coretypes.Receipt, error) {
	if receipt, ok := b.blockReceiptsCache.Get(blockNumber); ok {
		return receipt, nil
	}

	receipt := []*coretypes.Receipt{}
	err := b.app.EVMIndexer().IterateBlockTxReceipts(b.ctx, blockNumber, func(recept *coretypes.Receipt) (bool, error) {
		receipt = append(receipt, recept)
		return false, nil
	})
	if err != nil {
		b.logger.Error("failed to get block receipts", "err", err)
		return nil, NewInternalError("failed to get block receipts")
	}

	// cache the receipts
	_ = b.blockReceiptsCache.Add(blockNumber, receipt)
	return receipt, nil
}

// marshalReceipt marshals a transaction receipt into a JSON object.
func marshalReceipt(receipt *coretypes.Receipt, tx *rpctypes.RPCTransaction) map[string]any {
	for idx, log := range receipt.Logs {
		log.Index = uint(idx)
		if tx.BlockHash != nil {
			log.BlockHash = *tx.BlockHash
		}
		if tx.BlockNumber != nil {
			log.BlockNumber = tx.BlockNumber.ToInt().Uint64()
		}
		log.TxHash = tx.Hash
		if tx.TransactionIndex != nil {
			log.TxIndex = uint(*tx.TransactionIndex)
		}
	}

	fields := map[string]any{
		"blockHash":         tx.BlockHash,
		"blockNumber":       hexutil.Uint64(tx.BlockNumber.ToInt().Uint64()),
		"transactionHash":   tx.Hash,
		"transactionIndex":  *tx.TransactionIndex,
		"from":              tx.From,
		"to":                tx.To,
		"gasUsed":           hexutil.Uint64(receipt.GasUsed),
		"cumulativeGasUsed": hexutil.Uint64(receipt.CumulativeGasUsed),
		"contractAddress":   nil,
		"logs":              receipt.Logs,
		"logsBloom":         receipt.Bloom,
		"type":              hexutil.Uint(tx.Type),
		"effectiveGasPrice": (*hexutil.Big)(receipt.EffectiveGasPrice),
	}

	// Assign receipt status or post state.
	if len(receipt.PostState) > 0 {
		fields["root"] = hexutil.Bytes(receipt.PostState)
	} else {
		fields["status"] = hexutil.Uint(receipt.Status)
	}
	if receipt.Logs == nil {
		fields["logs"] = []*types.Log{}
	}

	// If the ContractAddress is 20 0x0 bytes, assume it is not a contract creation
	if receipt.ContractAddress != (common.Address{}) {
		fields["contractAddress"] = receipt.ContractAddress
	}

	return fields
}
