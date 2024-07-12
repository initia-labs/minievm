package backend

import (
	"context"
	"errors"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

	return tx.Hash(), b.SendTx(tx)
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

	res, err := b.clientCtx.BroadcastTxSync(txBytes)
	if err != nil {
		return err
	}
	if res.Code != 0 {
		return sdkerrors.ErrInvalidRequest.Wrapf("tx failed with code: %d: raw_log: %s", res.Code, res.RawLog)
	}

	return nil
}

func (b *JSONRPCBackend) getQueryCtx() (context.Context, error) {
	return b.app.CreateQueryContext(0, false)
}

func (b *JSONRPCBackend) getQueryCtxWithHeight(height uint64) (context.Context, error) {
	return b.app.CreateQueryContext(int64(height), false)
}

// GetTransactionByHash returns the transaction with the given hash.
func (b *JSONRPCBackend) GetTransactionByHash(hash common.Hash) (*rpctypes.RPCTransaction, error) {
	queryCtx, err := b.getQueryCtx()
	if err != nil {
		return nil, err
	}

	return b.app.EVMIndexer().TxByHash(queryCtx, hash)
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
		if err != nil {
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
		var err error
		queryCtx, err = b.app.CreateQueryContext(0, false)
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
	if err != nil {
		return nil, err
	}

	receipt, err := b.app.EVMIndexer().TxReceiptByHash(queryCtx, hash)
	if err != nil {
		return nil, err
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
	if err != nil {
		return nil, err
	}

	return b.app.EVMIndexer().TxByBlockAndIndex(queryCtx, number, uint64(idx))
}

// GetTransactionByBlockNumberAndIndex returns the transaction at the given block number and index.
func (b *JSONRPCBackend) GetTransactionByBlockNumberAndIndex(blockNum rpc.BlockNumber, idx hexutil.Uint) (*rpctypes.RPCTransaction, error) {
	queryCtx, err := b.getQueryCtx()
	if err != nil {
		return nil, err
	}

	number := uint64(blockNum.Int64())
	return b.app.EVMIndexer().TxByBlockAndIndex(queryCtx, number, uint64(idx))
}

// GetBlockTransactionCountByHash returns the number of transactions in a block from a block matching the given block hash.
func (b *JSONRPCBackend) GetBlockTransactionCountByHash(hash common.Hash) (*hexutil.Uint, error) {
	block, err := b.GetBlockByHash(hash, true)
	if err != nil {
		return nil, err
	}

	numTxs := hexutil.Uint(len(block["transactions"].([]*rpctypes.RPCTransaction)))
	return &numTxs, nil
}

// GetBlockTransactionCountByNumber returns the number of transactions in a block from a block matching the given block number.
func (b *JSONRPCBackend) GetBlockTransactionCountByNumber(blockNum rpc.BlockNumber) (*hexutil.Uint, error) {
	block, err := b.GetBlockByNumber(blockNum, true)
	if err != nil {
		return nil, err
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
	if err != nil {
		return nil, err
	}

	return rpcTx.ToTransaction().MarshalBinary()
}

// GetRawTransactionByBlockHashAndIndex returns the bytes of the transaction for the given block hash and index.
func (b *JSONRPCBackend) GetRawTransactionByBlockHashAndIndex(blockHash common.Hash, index hexutil.Uint) (hexutil.Bytes, error) {
	rpcTx, err := b.GetTransactionByBlockHashAndIndex(blockHash, index)
	if err != nil {
		return nil, err
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

// marshalReceipt marshals a transaction receipt into a JSON object.
func marshalReceipt(receipt *coretypes.Receipt, tx *rpctypes.RPCTransaction) map[string]interface{} {
	fields := map[string]interface{}{
		"blockHash":         tx.BlockHash,
		"blockNumber":       hexutil.Big(*tx.BlockNumber),
		"transactionHash":   tx.Hash,
		"transactionIndex":  hexutil.Uint64(*tx.TransactionIndex),
		"from":              tx.From,
		"to":                tx.To,
		"gasUsed":           hexutil.Uint64(receipt.GasUsed),
		"cumulativeGasUsed": hexutil.Uint64(receipt.CumulativeGasUsed),
		"contractAddress":   nil,
		"logs":              receipt.Logs,
		"logsBloom":         receipt.Bloom,
		"type":              hexutil.Uint(coretypes.LegacyTxType),
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
