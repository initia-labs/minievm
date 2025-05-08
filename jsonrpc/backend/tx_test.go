package backend_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/initia-labs/minievm/tests"
	evmkeeper "github.com/initia-labs/minievm/x/evm/keeper"
	evmtypes "github.com/initia-labs/minievm/x/evm/types"
	"github.com/stretchr/testify/require"
)

func Test_SendRawTransaction_EIP155(t *testing.T) {
	input := setupBackend(t)

	txBz, err := hexutil.Decode("0xf8a58085174876e800830186a08080b853604580600e600039806000f350fe7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe03601600081602082378035828234f58015156039578182fd5b8082525050506014600cf31ba02222222222222222222222222222222222222222222222222222222222222222a02222222222222222222222222222222222222222222222222222222222222222")
	require.NoError(t, err)

	_, err = input.backend.SendRawTransaction(txBz)
	require.ErrorContains(t, err, "EIP-155")
}

func Test_SendRawTransaction(t *testing.T) {
	input := setupBackend(t)
	app, _, backend, addrs, privKeys := input.app, input.addrs, input.backend, input.addrs, input.privKeys

	tx, _ := tests.GenerateCreateERC20Tx(t, app, privKeys[0])
	_, finalizeRes := tests.ExecuteTxs(t, app, tx)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	events := finalizeRes.TxResults[0].Events
	createEvent := events[len(events)-3]
	require.Equal(t, evmtypes.EventTypeContractCreated, createEvent.GetType())

	contractAddr, err := hexutil.Decode(createEvent.Attributes[0].Value)
	require.NoError(t, err)

	// mint 1_000_000 tokens to the first address
	tx, _ = tests.GenerateMintERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[0], new(big.Int).SetUint64(1_000_000_000_000))
	tx2, _ := tests.GenerateMintERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[0], new(big.Int).SetUint64(1_000_000_000_000), tests.SetNonce(2))
	_, finalizeRes = tests.ExecuteTxs(t, app, tx, tx2)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)
	tests.CheckTxResult(t, finalizeRes.TxResults[1], true)

	ctx, err := app.CreateQueryContext(0, false)
	require.NoError(t, err)

	// Acc: 0, Nonce: 4
	tx04, txHash04 := tests.GenerateTransferERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[1], new(big.Int).SetUint64(1_000_000), tests.SetNonce(4))
	evmTx04, _, err := evmkeeper.NewTxUtils(app.EVMKeeper).ConvertCosmosTxToEthereumTx(ctx, tx04)
	require.NoError(t, err)
	require.NotNil(t, evmTx04)

	txBz, err := evmTx04.MarshalBinary()
	require.NoError(t, err)
	_, err = backend.SendRawTransaction(txBz)
	require.NoError(t, err)

	// execute recheck txs
	err = input.cometRPC.RecheckTx()
	require.NoError(t, err)

	// Acc: 1, Nonce: 0
	tx10, txHash10 := tests.GenerateTransferERC20Tx(t, app, privKeys[1], common.BytesToAddress(contractAddr), addrs[1], new(big.Int).SetUint64(1_000_000), tests.SetNonce(0))
	evmTx10, _, err := evmkeeper.NewTxUtils(app.EVMKeeper).ConvertCosmosTxToEthereumTx(ctx, tx10)
	require.NoError(t, err)
	require.NotNil(t, evmTx10)

	txBz, err = evmTx10.MarshalBinary()
	require.NoError(t, err)
	_, err = backend.SendRawTransaction(txBz)
	require.NoError(t, err)

	// execute recheck txs
	err = input.cometRPC.RecheckTx()
	require.NoError(t, err)

	// Acc: 0, Nonce: 6
	tx06, txHash06 := tests.GenerateTransferERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[1], new(big.Int).SetUint64(1_000_000), tests.SetNonce(6))
	evmTx06, _, err := evmkeeper.NewTxUtils(app.EVMKeeper).ConvertCosmosTxToEthereumTx(ctx, tx06)
	require.NoError(t, err)
	require.NotNil(t, evmTx06)

	txBz, err = evmTx06.MarshalBinary()
	require.NoError(t, err)
	_, err = backend.SendRawTransaction(txBz)
	require.NoError(t, err)

	// execute recheck txs
	err = input.cometRPC.RecheckTx()
	require.NoError(t, err)

	// 1 in pending, and 2 in queued
	txPool, err := backend.TxPoolContent()
	require.NoError(t, err)
	require.Len(t, txPool["pending"][addrs[1].Hex()], 1)
	require.Equal(t, txPool["pending"][addrs[1].Hex()]["0"].Hash, txHash10)
	require.Len(t, txPool["queued"][addrs[0].Hex()], 2)
	require.Equal(t, txPool["queued"][addrs[0].Hex()]["4"].Hash, txHash04)
	require.Equal(t, txPool["queued"][addrs[0].Hex()]["6"].Hash, txHash06)

	// sending Nonce 3 should make Nonce 4 to be pending

	// Acc: 0, Nonce: 3
	tx03, txHash03 := tests.GenerateTransferERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[1], new(big.Int).SetUint64(1_000_000), tests.SetNonce(3))
	evmTx03, _, err := evmkeeper.NewTxUtils(app.EVMKeeper).ConvertCosmosTxToEthereumTx(ctx, tx03)
	require.NoError(t, err)
	require.NotNil(t, evmTx03)

	txBz, err = evmTx03.MarshalBinary()
	require.NoError(t, err)
	_, err = backend.SendRawTransaction(txBz)
	require.NoError(t, err)

	// execute recheck txs
	err = input.cometRPC.RecheckTx()
	require.NoError(t, err)

	// 3 in pending and 1 in queued
	txPool, err = backend.TxPoolContent()
	require.NoError(t, err)
	require.Len(t, txPool["pending"][addrs[0].Hex()], 2)
	require.Equal(t, txPool["pending"][addrs[0].Hex()]["3"].Hash, txHash03)
	require.Equal(t, txPool["pending"][addrs[0].Hex()]["4"].Hash, txHash04)
	require.Len(t, txPool["pending"][addrs[1].Hex()], 1)
	require.Equal(t, txPool["pending"][addrs[1].Hex()]["0"].Hash, txHash10)
	require.Len(t, txPool["queued"][addrs[0].Hex()], 1)
	require.Equal(t, txPool["queued"][addrs[0].Hex()]["6"].Hash, txHash06)

	// sending Nonce 5 should make Nonce 6 to be pending

	// Acc: 0, Nonce: 5
	tx05, txHash05 := tests.GenerateTransferERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[1], new(big.Int).SetUint64(1_000_000), tests.SetNonce(5))
	evmTx05, _, err := evmkeeper.NewTxUtils(app.EVMKeeper).ConvertCosmosTxToEthereumTx(ctx, tx05)
	require.NoError(t, err)
	require.NotNil(t, evmTx05)

	txBz, err = evmTx05.MarshalBinary()
	require.NoError(t, err)
	_, err = backend.SendRawTransaction(txBz)
	require.NoError(t, err)

	// execute recheck txs
	err = input.cometRPC.RecheckTx()
	require.NoError(t, err)

	// 5 in pending and 0 in queued
	txPool, err = backend.TxPoolContent()
	require.NoError(t, err)
	require.Len(t, txPool["pending"][addrs[0].Hex()], 4)
	require.Equal(t, txPool["pending"][addrs[0].Hex()]["3"].Hash, txHash03)
	require.Equal(t, txPool["pending"][addrs[0].Hex()]["4"].Hash, txHash04)
	require.Equal(t, txPool["pending"][addrs[0].Hex()]["5"].Hash, txHash05)
	require.Equal(t, txPool["pending"][addrs[0].Hex()]["6"].Hash, txHash06)
	require.Len(t, txPool["pending"][addrs[1].Hex()], 1)
	require.Equal(t, txPool["pending"][addrs[1].Hex()]["0"].Hash, txHash10)
	require.Empty(t, txPool["queued"])
}

func Test_GetTransactionCount(t *testing.T) {
	input := setupBackend(t)
	app, _, backend, addrs, privKeys := input.app, input.addrs, input.backend, input.addrs, input.privKeys

	tx, _ := tests.GenerateCreateERC20Tx(t, app, privKeys[0])
	_, finalizeRes := tests.ExecuteTxs(t, app, tx)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	events := finalizeRes.TxResults[0].Events
	createEvent := events[len(events)-3]
	require.Equal(t, evmtypes.EventTypeContractCreated, createEvent.GetType())

	contractAddr, err := hexutil.Decode(createEvent.Attributes[0].Value)
	require.NoError(t, err)

	// mint 1_000_000 tokens to the first address
	tx, _ = tests.GenerateMintERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[0], new(big.Int).SetUint64(1_000_000_000_000))
	tx2, _ := tests.GenerateMintERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[0], new(big.Int).SetUint64(1_000_000_000_000), tests.SetNonce(2))
	_, finalizeRes = tests.ExecuteTxs(t, app, tx, tx2)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)
	tests.CheckTxResult(t, finalizeRes.TxResults[1], true)

	count, err := backend.GetTransactionCount(common.BytesToAddress(addrs[0].Bytes()), rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(app.LastBlockHeight()-1)))
	require.NoError(t, err)
	require.Equal(t, "0x1", count.String())

	count, err = backend.GetTransactionCount(common.BytesToAddress(addrs[0].Bytes()), rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(app.LastBlockHeight())))
	require.NoError(t, err)
	require.Equal(t, "0x3", count.String())

	// try wrong block hash
	_, err = backend.GetTransactionCount(common.BytesToAddress(addrs[0].Bytes()), rpc.BlockNumberOrHashWithHash(common.Hash{}, false))
	require.Error(t, err)

	// try with valid block hash
	header, err := backend.GetHeaderByNumber(rpc.BlockNumber(app.LastBlockHeight()))
	require.NoError(t, err)

	count, err = backend.GetTransactionCount(common.BytesToAddress(addrs[0].Bytes()), rpc.BlockNumberOrHashWithHash(header.Hash(), false))
	require.NoError(t, err)
	require.Equal(t, "0x3", count.String())
}

func Test_GetTransactionReceipt(t *testing.T) {
	input := setupBackend(t)
	app, _, backend, addrs, privKeys := input.app, input.addrs, input.backend, input.addrs, input.privKeys

	tx, _ := tests.GenerateCreateERC20Tx(t, app, privKeys[0])
	_, finalizeRes := tests.ExecuteTxs(t, app, tx)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	events := finalizeRes.TxResults[0].Events
	createEvent := events[len(events)-3]
	require.Equal(t, evmtypes.EventTypeContractCreated, createEvent.GetType())

	contractAddr, err := hexutil.Decode(createEvent.Attributes[0].Value)
	require.NoError(t, err)

	// mint 1_000_000 tokens to the first address
	tx, txHash := tests.GenerateMintERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[0], new(big.Int).SetUint64(1_000_000_000_000))
	tx2, txHash2 := tests.GenerateMintERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[0], new(big.Int).SetUint64(1_000_000_000_000), tests.SetNonce(2))
	_, finalizeRes = tests.ExecuteTxs(t, app, tx, tx2)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)
	tests.CheckTxResult(t, finalizeRes.TxResults[1], true)

	header, err := backend.GetHeaderByNumber(rpc.BlockNumber(app.LastBlockHeight()))
	require.NoError(t, err)

	receipt, err := backend.GetTransactionReceipt(txHash)
	require.NoError(t, err)
	require.NotNil(t, receipt)
	require.Equal(t, header.Hash(), *receipt["blockHash"].(*common.Hash))
	require.Equal(t, header.Number.Uint64(), uint64(receipt["blockNumber"].(hexutil.Uint64)))
	require.Equal(t, txHash, receipt["transactionHash"])
	require.Equal(t, uint64(1), uint64(receipt["transactionIndex"].(hexutil.Uint64)))

	receipt2, err := backend.GetTransactionReceipt(txHash2)
	require.NoError(t, err)
	require.NotNil(t, receipt2)
	require.Equal(t, header.Hash(), *receipt2["blockHash"].(*common.Hash))
	require.Equal(t, header.Number.Uint64(), uint64(receipt2["blockNumber"].(hexutil.Uint64)))
	require.Equal(t, txHash2, receipt2["transactionHash"])
	require.Equal(t, uint64(2), uint64(receipt2["transactionIndex"].(hexutil.Uint64)))
}

func Test_GetTransaction(t *testing.T) {
	input := setupBackend(t)
	app, _, backend, addrs, privKeys := input.app, input.addrs, input.backend, input.addrs, input.privKeys

	tx, _ := tests.GenerateCreateERC20Tx(t, app, privKeys[0])
	_, finalizeRes := tests.ExecuteTxs(t, app, tx)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	events := finalizeRes.TxResults[0].Events
	createEvent := events[len(events)-3]
	require.Equal(t, evmtypes.EventTypeContractCreated, createEvent.GetType())

	contractAddr, err := hexutil.Decode(createEvent.Attributes[0].Value)
	require.NoError(t, err)

	// mint 1_000_000 tokens to the first address
	tx, txHash := tests.GenerateMintERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[0], new(big.Int).SetUint64(1_000_000_000_000))
	tx2, txHash2 := tests.GenerateMintERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[0], new(big.Int).SetUint64(1_000_000_000_000), tests.SetNonce(2))
	_, finalizeRes = tests.ExecuteTxs(t, app, tx, tx2)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)
	tests.CheckTxResult(t, finalizeRes.TxResults[1], true)

	txByHash, err := backend.GetTransactionByHash(txHash)
	require.NoError(t, err)

	header, err := backend.GetHeaderByNumber(rpc.BlockNumber(app.LastBlockHeight()))
	require.NoError(t, err)

	txByHashAndIndex, err := backend.GetTransactionByBlockHashAndIndex(header.Hash(), 1)
	require.NoError(t, err)
	require.Equal(t, txByHash, txByHashAndIndex)
	txByNumberAndIndex, err := backend.GetTransactionByBlockNumberAndIndex(rpc.BlockNumber(app.LastBlockHeight()), 1)
	require.NoError(t, err)
	require.Equal(t, txByHash, txByNumberAndIndex)

	txByHash2, err := backend.GetTransactionByHash(txHash2)
	require.NoError(t, err)

	txByHashAndIndex2, err := backend.GetTransactionByBlockHashAndIndex(header.Hash(), 2)
	require.NoError(t, err)
	require.Equal(t, txByHash2, txByHashAndIndex2)
	txByNumberAndIndex2, err := backend.GetTransactionByBlockNumberAndIndex(rpc.BlockNumber(app.LastBlockHeight()), 2)
	require.NoError(t, err)
	require.Equal(t, txByHash2, txByNumberAndIndex2)

	// try with wrong block hash
	_, err = backend.GetTransactionByBlockHashAndIndex(common.Hash{}, 1)
	require.Error(t, err)
}

func Test_BlockTransactionCount(t *testing.T) {
	input := setupBackend(t)
	app, _, backend, addrs, privKeys := input.app, input.addrs, input.backend, input.addrs, input.privKeys

	tx, _ := tests.GenerateCreateERC20Tx(t, app, privKeys[0])
	_, finalizeRes := tests.ExecuteTxs(t, app, tx)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	events := finalizeRes.TxResults[0].Events
	createEvent := events[len(events)-3]
	require.Equal(t, evmtypes.EventTypeContractCreated, createEvent.GetType())

	contractAddr, err := hexutil.Decode(createEvent.Attributes[0].Value)
	require.NoError(t, err)

	// mint 1_000_000 tokens to the first address
	tx, _ = tests.GenerateMintERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[0], new(big.Int).SetUint64(1_000_000_000_000))
	tx2, _ := tests.GenerateMintERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[0], new(big.Int).SetUint64(1_000_000_000_000), tests.SetNonce(2))
	_, finalizeRes = tests.ExecuteTxs(t, app, tx, tx2)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)
	tests.CheckTxResult(t, finalizeRes.TxResults[1], true)

	header, err := backend.GetHeaderByNumber(rpc.BlockNumber(app.LastBlockHeight()))
	require.NoError(t, err)

	count, err := backend.GetBlockTransactionCountByHash(header.Hash())
	require.NoError(t, err)
	require.Equal(t, "0x2", count.String())

	count, err = backend.GetBlockTransactionCountByNumber(rpc.BlockNumber(app.LastBlockHeight()))
	require.NoError(t, err)
	require.Equal(t, "0x2", count.String())
}

func Test_GetRawTransaction(t *testing.T) {
	input := setupBackend(t)
	app, _, backend, addrs, privKeys := input.app, input.addrs, input.backend, input.addrs, input.privKeys

	tx, _ := tests.GenerateCreateERC20Tx(t, app, privKeys[0])
	_, finalizeRes := tests.ExecuteTxs(t, app, tx)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	events := finalizeRes.TxResults[0].Events
	createEvent := events[len(events)-3]
	require.Equal(t, evmtypes.EventTypeContractCreated, createEvent.GetType())

	contractAddr, err := hexutil.Decode(createEvent.Attributes[0].Value)
	require.NoError(t, err)

	// mint 1_000_000 tokens to the first address
	tx, txHash := tests.GenerateMintERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[0], new(big.Int).SetUint64(1_000_000_000_000))
	tx2, _ := tests.GenerateMintERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[0], new(big.Int).SetUint64(1_000_000_000_000), tests.SetNonce(2))
	_, finalizeRes = tests.ExecuteTxs(t, app, tx, tx2)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)
	tests.CheckTxResult(t, finalizeRes.TxResults[1], true)

	txByHash, err := backend.GetRawTransactionByHash(txHash)
	require.NoError(t, err)

	header, err := backend.GetHeaderByNumber(rpc.BlockNumber(app.LastBlockHeight()))
	require.NoError(t, err)

	txByHashAndIndex, err := backend.GetRawTransactionByBlockHashAndIndex(header.Hash(), 1)
	require.NoError(t, err)
	require.Equal(t, txByHash, txByHashAndIndex)

	ctx, err := app.CreateQueryContext(0, false)
	require.NoError(t, err)
	evmTx, _, err := evmkeeper.NewTxUtils(app.EVMKeeper).ConvertCosmosTxToEthereumTx(ctx, tx)
	require.NoError(t, err)
	require.NotNil(t, evmTx)

	txBz, err := evmTx.MarshalBinary()
	require.NoError(t, err)
	require.Equal(t, txBz, []byte(txByHash))
}

func Test_PendingTransactions(t *testing.T) {
	input := setupBackend(t)
	app, _, backend, addrs, privKeys := input.app, input.addrs, input.backend, input.addrs, input.privKeys

	tx, _ := tests.GenerateCreateERC20Tx(t, app, privKeys[0])
	_, finalizeRes := tests.ExecuteTxs(t, app, tx)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	events := finalizeRes.TxResults[0].Events
	createEvent := events[len(events)-3]
	require.Equal(t, evmtypes.EventTypeContractCreated, createEvent.GetType())

	contractAddr, err := hexutil.Decode(createEvent.Attributes[0].Value)
	require.NoError(t, err)

	// mint 1_000_000 tokens to the first address
	tx, _ = tests.GenerateMintERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[0], new(big.Int).SetUint64(1_000_000_000_000))
	tx2, _ := tests.GenerateMintERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[0], new(big.Int).SetUint64(1_000_000_000_000), tests.SetNonce(2))
	_, finalizeRes = tests.ExecuteTxs(t, app, tx, tx2)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)
	tests.CheckTxResult(t, finalizeRes.TxResults[1], true)

	ctx, err := app.CreateQueryContext(0, false)
	require.NoError(t, err)

	// Acc: 0, Nonce: 3
	tx03, txHash03 := tests.GenerateTransferERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[1], new(big.Int).SetUint64(1_000_000), tests.SetNonce(3))
	evmTx03, _, err := evmkeeper.NewTxUtils(app.EVMKeeper).ConvertCosmosTxToEthereumTx(ctx, tx03)
	require.NoError(t, err)
	require.NotNil(t, evmTx03)

	txBz, err := evmTx03.MarshalBinary()
	require.NoError(t, err)
	_, err = backend.SendRawTransaction(txBz)
	require.NoError(t, err)

	// Acc: 0, Nonce: 4
	tx04, txHash04 := tests.GenerateTransferERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[1], new(big.Int).SetUint64(1_000_000), tests.SetNonce(4))
	evmTx04, _, err := evmkeeper.NewTxUtils(app.EVMKeeper).ConvertCosmosTxToEthereumTx(ctx, tx04)
	require.NoError(t, err)
	require.NotNil(t, evmTx04)

	txBz, err = evmTx04.MarshalBinary()
	require.NoError(t, err)
	_, err = backend.SendRawTransaction(txBz)
	require.NoError(t, err)

	// Acc: 1, Nonce: 0
	tx10, txHash10 := tests.GenerateTransferERC20Tx(t, app, privKeys[1], common.BytesToAddress(contractAddr), addrs[1], new(big.Int).SetUint64(1_000_000), tests.SetNonce(0))
	evmTx10, _, err := evmkeeper.NewTxUtils(app.EVMKeeper).ConvertCosmosTxToEthereumTx(ctx, tx10)
	require.NoError(t, err)
	require.NotNil(t, evmTx10)

	txBz, err = evmTx10.MarshalBinary()
	require.NoError(t, err)
	_, err = backend.SendRawTransaction(txBz)
	require.NoError(t, err)

	pendingTxs, err := backend.PendingTransactions()
	require.NoError(t, err)
	require.Len(t, pendingTxs, 3)

	txHashes := []common.Hash{txHash03, txHash04, txHash10}
	for _, tx := range pendingTxs {
		require.Contains(t, txHashes, tx.Hash)
	}
}

func Test_BlockReceipts(t *testing.T) {
	input := setupBackend(t)
	app, _, backend, addrs, privKeys := input.app, input.addrs, input.backend, input.addrs, input.privKeys

	tx, _ := tests.GenerateCreateERC20Tx(t, app, privKeys[0])
	_, finalizeRes := tests.ExecuteTxs(t, app, tx)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	events := finalizeRes.TxResults[0].Events
	createEvent := events[len(events)-3]
	require.Equal(t, evmtypes.EventTypeContractCreated, createEvent.GetType())

	contractAddr, err := hexutil.Decode(createEvent.Attributes[0].Value)
	require.NoError(t, err)

	// mint 1_000_000 tokens to the first address
	tx, txHash := tests.GenerateMintERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[0], new(big.Int).SetUint64(1_000_000_000_000))
	tx2, txHash2 := tests.GenerateMintERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[0], new(big.Int).SetUint64(1_000_000_000_000), tests.SetNonce(2))
	_, finalizeRes = tests.ExecuteTxs(t, app, tx, tx2)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)
	tests.CheckTxResult(t, finalizeRes.TxResults[1], true)

	header, err := backend.GetHeaderByNumber(rpc.BlockNumber(app.LastBlockHeight()))
	require.NoError(t, err)

	receipts, err := backend.GetBlockReceipts(rpc.BlockNumberOrHashWithHash(header.Hash(), false))
	require.NoError(t, err)
	require.Len(t, receipts, 2)

	receipt := receipts[0]
	require.NotNil(t, receipt)
	require.Equal(t, header.Hash(), *receipt["blockHash"].(*common.Hash))
	require.Equal(t, header.Number.Uint64(), uint64(receipt["blockNumber"].(hexutil.Uint64)))
	require.Equal(t, txHash, receipt["transactionHash"])
	require.Equal(t, uint64(1), uint64(receipt["transactionIndex"].(hexutil.Uint64)))

	receipt2 := receipts[1]
	require.NotNil(t, receipt2)
	require.Equal(t, header.Hash(), *receipt2["blockHash"].(*common.Hash))
	require.Equal(t, header.Number.Uint64(), uint64(receipt2["blockNumber"].(hexutil.Uint64)))
	require.Equal(t, txHash2, receipt2["transactionHash"])
	require.Equal(t, uint64(2), uint64(receipt2["transactionIndex"].(hexutil.Uint64)))
}
