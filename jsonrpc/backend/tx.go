package backend

import (
	"context"
	"errors"
	"fmt"

	"cosmossdk.io/collections"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"

	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
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

	if err := b.SendTx(tx); err != nil {
		return common.Hash{}, err
	}

	return tx.Hash(), nil
}

func (b *JSONRPCBackend) SendTx(tx *coretypes.Transaction) error {
	queryCtx, err := b.getQueryCtx()
	if err != nil {
		return err
	}

	cosmosTx, err := keeper.NewTxUtils(b.app.EVMKeeper).ConvertEthereumTxToCosmosTx(queryCtx, tx)
	if err != nil {
		return err
	}

	txBytes, err := b.app.TxEncode(cosmosTx)
	if err != nil {
		return err
	}

	authTx, ok := cosmosTx.(authsigning.Tx)
	if !ok {
		return NewInternalError("failed to convert cosmosTx to authsigning.Tx")
	}

	sigs, err := authTx.GetSignaturesV2()
	if err != nil || len(sigs) != 1 {
		b.logger.Error("failed to get signatures from authsigning.Tx", "err", err)
		return NewInternalError("failed to get signatures from authsigning.Tx")
	}

	sig := sigs[0]
	txSeq := sig.Sequence
	accSeq := uint64(0)
	sender := sdk.AccAddress(sig.PubKey.Address().Bytes())

	senderHex := hexutil.Encode(sender.Bytes())

	// hold mutex for each sender
	b.acquireAccMut(senderHex)
	defer b.releaseAccMut(senderHex)

	checkCtx := b.app.GetContextForCheckTx(nil)
	if acc := b.app.AccountKeeper.GetAccount(checkCtx, sender); acc != nil {
		accSeq = acc.GetSequence()
	}

	if accSeq > txSeq {
		return fmt.Errorf("%w: next nonce %v, tx nonce %v", core.ErrNonceTooLow, accSeq, txSeq)
	}

	b.logger.Debug("enqueue tx", "sender", senderHex, "txSeq", txSeq, "accSeq", accSeq)
	cacheKey := fmt.Sprintf("%s-%d", senderHex, txSeq)
	_ = b.queuedTxs.Add(cacheKey, txBytes)

	// check if there are queued txs which can be sent
	for {
		cacheKey := fmt.Sprintf("%s-%d", senderHex, accSeq)
		if txBytes, ok := b.queuedTxs.Get(cacheKey); ok {
			_ = b.queuedTxs.Remove(cacheKey)

			b.logger.Debug("broadcast queued tx", "sender", senderHex, "txSeq", accSeq)
			res, err := b.clientCtx.BroadcastTxSync(txBytes)
			if err != nil {
				return err
			}
			if res.Code != 0 {
				return sdkerrors.ErrInvalidRequest.Wrapf("tx failed with code: %d: raw_log: %s", res.Code, res.RawLog)
			}
		} else {
			break
		}

		accSeq++
	}

	return nil
}

func (b *JSONRPCBackend) getQueryCtx() (context.Context, error) {
	return b.app.CreateQueryContext(0, false)
}

func (b *JSONRPCBackend) getQueryCtxWithHeight(height uint64) (context.Context, error) {
	// check whether the given height is bigger than the latest block height
	num, err := b.BlockNumber()
	if err != nil {
		return nil, err
	}
	if height >= uint64(num) {
		height = 0
	}

	return b.app.CreateQueryContext(int64(height), false)
}

// GetTransactionByHash returns the transaction with the given hash.
func (b *JSONRPCBackend) GetTransactionByHash(hash common.Hash) (*rpctypes.RPCTransaction, error) {
	queryCtx, err := b.getQueryCtx()
	if err != nil {
		return nil, err
	}

	tx, err := b.app.EVMIndexer().TxByHash(queryCtx, hash)
	if err != nil && errors.Is(err, collections.ErrNotFound) {
		return nil, nil
	} else if err != nil {
		b.logger.Error("failed to get transaction by hash", "err", err)
		return nil, NewTxIndexingError()
	}

	return tx, nil
}

// GetTransactionCount returns the number of transactions at the given block number.
func (b *JSONRPCBackend) GetTransactionCount(address common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*hexutil.Uint64, error) {
	sdkAddr := sdk.AccAddress(address[:])

	var blockNumber rpc.BlockNumber
	if blockHash, ok := blockNrOrHash.Hash(); ok {
		queryCtx, err := b.getQueryCtx()
		if err != nil {
			return nil, err
		}

		blockNumberU64, err := b.app.EVMIndexer().BlockHashToNumber(queryCtx, blockHash)
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
		queryCtx, err = b.getQueryCtxWithHeight(uint64(blockNumber.Int64()))
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
func (b *JSONRPCBackend) GetTransactionReceipt(hash common.Hash) (map[string]interface{}, error) {
	queryCtx, err := b.getQueryCtx()
	if err != nil {
		return nil, err
	}

	tx, err := b.app.EVMIndexer().TxByHash(queryCtx, hash)
	if err != nil && errors.Is(err, collections.ErrNotFound) {
		return nil, nil
	} else if err != nil {
		b.logger.Error("failed to get transaction by hash", "err", err)
		return nil, NewTxIndexingError()
	}

	receipt, err := b.app.EVMIndexer().TxReceiptByHash(queryCtx, hash)
	if err != nil && errors.Is(err, collections.ErrNotFound) {
		return nil, nil
	} else if err != nil {
		b.logger.Error("failed to get transaction receipt by hash", "err", err)
		return nil, NewTxIndexingError()
	}

	return marshalReceipt(receipt, tx), nil
}

// GetTransactionByBlockHashAndIndex returns the transaction at the given block hash and index.
func (b *JSONRPCBackend) GetTransactionByBlockHashAndIndex(hash common.Hash, idx hexutil.Uint) (*rpctypes.RPCTransaction, error) {
	queryCtx, err := b.getQueryCtx()
	if err != nil {
		return nil, err
	}

	number, err := b.app.EVMIndexer().BlockHashToNumber(queryCtx, hash)
	if err != nil && errors.Is(err, collections.ErrNotFound) {
		return nil, nil
	} else if err != nil {
		b.logger.Error("failed to get block number by hash", "err", err)
		return nil, err
	}

	rpcTx, err := b.app.EVMIndexer().TxByBlockAndIndex(queryCtx, number, uint64(idx))
	if err != nil && errors.Is(err, collections.ErrNotFound) {
		return nil, nil
	} else if err != nil {
		b.logger.Error("failed to get transaction by block and index", "err", err)
		return nil, err
	}

	return rpcTx, nil
}

// GetTransactionByBlockNumberAndIndex returns the transaction at the given block number and index.
func (b *JSONRPCBackend) GetTransactionByBlockNumberAndIndex(blockNum rpc.BlockNumber, idx hexutil.Uint) (*rpctypes.RPCTransaction, error) {
	queryCtx, err := b.getQueryCtx()
	if err != nil {
		return nil, err
	}

	number := uint64(blockNum.Int64())
	rpcTx, err := b.app.EVMIndexer().TxByBlockAndIndex(queryCtx, number, uint64(idx))
	if err != nil && errors.Is(err, collections.ErrNotFound) {
		return nil, nil
	} else if err != nil {
		b.logger.Error("failed to get transaction by block and index", "err", err)
		return nil, err
	}

	return rpcTx, nil
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
	queryCtx, err := b.getQueryCtx()
	if err != nil {
		return nil, err
	}

	rpcTx, err := b.app.EVMIndexer().TxByHash(queryCtx, hash)
	if err != nil && errors.Is(err, collections.ErrNotFound) {
		return nil, nil
	} else if err != nil {
		b.logger.Error("failed to get raw transaction by hash", "err", err)
		return nil, NewTxIndexingError()
	} else if rpcTx == nil {
		return nil, nil
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
	chainID, err := b.ChainID()
	if err != nil {
		return nil, err
	}

	queryCtx, err := b.getQueryCtx()
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
				rpctypes.NewRPCTransaction(ethTx, common.Hash{}, 0, 0, chainID),
			)
		}
	}

	return result, nil
}

func (b *JSONRPCBackend) getBlockTransactions(blockNumber uint64) ([]*rpctypes.RPCTransaction, error) {
	queryCtx, err := b.getQueryCtx()
	if err != nil {
		return nil, err
	}

	txs := []*rpctypes.RPCTransaction{}
	err = b.app.EVMIndexer().IterateBlockTxs(queryCtx, blockNumber, func(tx *rpctypes.RPCTransaction) (bool, error) {
		txs = append(txs, tx)
		return false, nil
	})
	if err != nil {
		return nil, err
	}

	return txs, nil
}

func (b *JSONRPCBackend) getBLockReceipts(blockNumber uint64) ([]*coretypes.Receipt, error) {
	queryCtx, err := b.getQueryCtx()
	if err != nil {
		return nil, err
	}

	recepts := []*coretypes.Receipt{}
	err = b.app.EVMIndexer().IterateBlockTxRecepts(queryCtx, blockNumber, func(recept *coretypes.Receipt) (bool, error) {
		recepts = append(recepts, recept)
		return false, nil
	})
	if err != nil {
		return nil, err
	}

	return recepts, nil
}

func (b *JSONRPCBackend) GetBlockReceipts(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) ([]map[string]interface{}, error) {
	blockNumber, err := b.resolveBlockNrOrHash(blockNrOrHash)
	if err != nil {
		return nil, err
	}

	txs, err := b.getBlockTransactions(blockNumber)
	if err != nil {
		return nil, err
	}

	receipts, err := b.getBLockReceipts(blockNumber)
	if err != nil {
		return nil, err
	}

	if len(txs) != len(receipts) {
		return nil, fmt.Errorf("receipts length mismatch: %d vs %d", len(txs), len(receipts))
	}

	result := make([]map[string]interface{}, len(receipts))
	for i, receipt := range receipts {
		result[i] = marshalReceipt(receipt, txs[i])
	}

	return result, nil
}

// marshalReceipt marshals a transaction receipt into a JSON object.
func marshalReceipt(receipt *coretypes.Receipt, tx *rpctypes.RPCTransaction) map[string]interface{} {
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

	fields := map[string]interface{}{
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
