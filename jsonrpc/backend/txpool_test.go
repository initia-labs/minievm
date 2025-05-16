package backend_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/initia-labs/minievm/tests"
	evmkeeper "github.com/initia-labs/minievm/x/evm/keeper"
	evmtypes "github.com/initia-labs/minievm/x/evm/types"
	"github.com/stretchr/testify/require"
)

func Test_TxPoolContextFrom(t *testing.T) {
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

	// Acc: 1, Nonce: 0
	tx10, txHash10 := tests.GenerateTransferERC20Tx(t, app, privKeys[1], common.BytesToAddress(contractAddr), addrs[1], new(big.Int).SetUint64(1_000_000), tests.SetNonce(0))
	evmTx10, _, err := evmkeeper.NewTxUtils(app.EVMKeeper).ConvertCosmosTxToEthereumTx(ctx, tx10)
	require.NoError(t, err)
	require.NotNil(t, evmTx10)

	txBz, err = evmTx10.MarshalBinary()
	require.NoError(t, err)
	_, err = backend.SendRawTransaction(txBz)
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

	// 2 in queued and 0 in pending
	txPool, err := backend.TxPoolContentFrom(addrs[0])
	require.NoError(t, err)
	require.Empty(t, txPool["pending"])
	require.Len(t, txPool["queued"], 2)
	require.Equal(t, txPool["queued"]["4"].Hash, txHash04)
	require.Equal(t, txPool["queued"]["6"].Hash, txHash06)

	// 1 in pending and 0 in queued
	txPool, err = backend.TxPoolContentFrom(addrs[1])
	require.NoError(t, err)
	require.Len(t, txPool["pending"], 1)
	require.Equal(t, txPool["pending"]["0"].Hash, txHash10)
	require.Empty(t, txPool["queued"])

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

	// execute recheck
	err = input.cometRPC.RecheckTx()
	require.NoError(t, err)

	// 2 in pending and 1 in queued
	txPool, err = backend.TxPoolContentFrom(addrs[0])
	require.NoError(t, err)
	require.Len(t, txPool["pending"], 2)
	require.Equal(t, txPool["pending"]["3"].Hash, txHash03)
	require.Equal(t, txPool["pending"]["4"].Hash, txHash04)
	require.Len(t, txPool["queued"], 1)
	require.Equal(t, txPool["queued"]["6"].Hash, txHash06)

	// 1 in pending and 0 in queued
	txPool, err = backend.TxPoolContentFrom(addrs[1])
	require.NoError(t, err)
	require.Len(t, txPool["pending"], 1)
	require.Equal(t, txPool["pending"]["0"].Hash, txHash10)
	require.Empty(t, txPool["queued"])

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

	// execute recheck
	err = input.cometRPC.RecheckTx()
	require.NoError(t, err)

	// 4 in pending and 0 in queued
	txPool, err = backend.TxPoolContentFrom(addrs[0])
	require.NoError(t, err)
	require.Len(t, txPool["pending"], 4)
	require.Equal(t, txPool["pending"]["3"].Hash, txHash03)
	require.Equal(t, txPool["pending"]["4"].Hash, txHash04)
	require.Equal(t, txPool["pending"]["5"].Hash, txHash05)
	require.Equal(t, txPool["pending"]["6"].Hash, txHash06)
	require.Empty(t, txPool["queued"])

	// 1 in pending
	txPool, err = backend.TxPoolContentFrom(addrs[1])
	require.NoError(t, err)
	require.Len(t, txPool["pending"], 1)
	require.Equal(t, txPool["pending"]["0"].Hash, txHash10)
	require.Empty(t, txPool["queued"])
}

func Test_TxPoolStatus(t *testing.T) {
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
	tx04, _ := tests.GenerateTransferERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[1], new(big.Int).SetUint64(1_000_000), tests.SetNonce(4))
	evmTx04, _, err := evmkeeper.NewTxUtils(app.EVMKeeper).ConvertCosmosTxToEthereumTx(ctx, tx04)
	require.NoError(t, err)
	require.NotNil(t, evmTx04)

	txBz, err := evmTx04.MarshalBinary()
	require.NoError(t, err)
	_, err = backend.SendRawTransaction(txBz)
	require.NoError(t, err)

	// execute recheck
	err = input.cometRPC.RecheckTx()
	require.NoError(t, err)

	// Acc: 1, Nonce: 0
	tx10, _ := tests.GenerateTransferERC20Tx(t, app, privKeys[1], common.BytesToAddress(contractAddr), addrs[1], new(big.Int).SetUint64(1_000_000), tests.SetNonce(0))
	evmTx10, _, err := evmkeeper.NewTxUtils(app.EVMKeeper).ConvertCosmosTxToEthereumTx(ctx, tx10)
	require.NoError(t, err)
	require.NotNil(t, evmTx10)

	txBz, err = evmTx10.MarshalBinary()
	require.NoError(t, err)
	_, err = backend.SendRawTransaction(txBz)
	require.NoError(t, err)

	// execute recheck
	err = input.cometRPC.RecheckTx()
	require.NoError(t, err)

	// Acc: 0, Nonce: 6
	tx06, _ := tests.GenerateTransferERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[1], new(big.Int).SetUint64(1_000_000), tests.SetNonce(6))
	evmTx06, _, err := evmkeeper.NewTxUtils(app.EVMKeeper).ConvertCosmosTxToEthereumTx(ctx, tx06)
	require.NoError(t, err)
	require.NotNil(t, evmTx06)

	txBz, err = evmTx06.MarshalBinary()
	require.NoError(t, err)
	_, err = backend.SendRawTransaction(txBz)
	require.NoError(t, err)

	// execute recheck
	err = input.cometRPC.RecheckTx()
	require.NoError(t, err)

	// 2 in queued and 0 in pending
	status, err := backend.TxPoolStatus()
	require.NoError(t, err)
	require.Equal(t, 1, int(status["pending"]))
	require.Equal(t, 2, int(status["queued"]))

	// sending Nonce 3 should make Nonce 4 to be pending

	// Acc: 0, Nonce: 3
	tx03, _ := tests.GenerateTransferERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[1], new(big.Int).SetUint64(1_000_000), tests.SetNonce(3))
	evmTx03, _, err := evmkeeper.NewTxUtils(app.EVMKeeper).ConvertCosmosTxToEthereumTx(ctx, tx03)
	require.NoError(t, err)
	require.NotNil(t, evmTx03)

	txBz, err = evmTx03.MarshalBinary()
	require.NoError(t, err)
	_, err = backend.SendRawTransaction(txBz)
	require.NoError(t, err)

	// execute recheck
	err = input.cometRPC.RecheckTx()
	require.NoError(t, err)

	// 3 in pending and 1 in queued
	status, err = backend.TxPoolStatus()
	require.NoError(t, err)
	require.Equal(t, 3, int(status["pending"]))
	require.Equal(t, 1, int(status["queued"]))

	// sending Nonce 5 should make Nonce 6 to be pending

	// Acc: 0, Nonce: 5
	tx05, _ := tests.GenerateTransferERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[1], new(big.Int).SetUint64(1_000_000), tests.SetNonce(5))
	evmTx05, _, err := evmkeeper.NewTxUtils(app.EVMKeeper).ConvertCosmosTxToEthereumTx(ctx, tx05)
	require.NoError(t, err)
	require.NotNil(t, evmTx05)

	txBz, err = evmTx05.MarshalBinary()
	require.NoError(t, err)
	_, err = backend.SendRawTransaction(txBz)
	require.NoError(t, err)

	// execute recheck
	err = input.cometRPC.RecheckTx()
	require.NoError(t, err)

	// 5 in pending and 0 in queued
	status, err = backend.TxPoolStatus()
	require.NoError(t, err)
	require.Equal(t, 5, int(status["pending"]))
	require.Equal(t, 0, int(status["queued"]))
}

func Test_TxPoolInspect(t *testing.T) {
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
	tx04, _ := tests.GenerateTransferERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[1], new(big.Int).SetUint64(1_000_000), tests.SetNonce(4))
	evmTx04, _, err := evmkeeper.NewTxUtils(app.EVMKeeper).ConvertCosmosTxToEthereumTx(ctx, tx04)
	require.NoError(t, err)
	require.NotNil(t, evmTx04)

	txBz, err := evmTx04.MarshalBinary()
	require.NoError(t, err)
	_, err = backend.SendRawTransaction(txBz)
	require.NoError(t, err)

	// execute recheck
	err = input.cometRPC.RecheckTx()
	require.NoError(t, err)

	// Acc: 1, Nonce: 0
	tx10, _ := tests.GenerateTransferERC20Tx(t, app, privKeys[1], common.BytesToAddress(contractAddr), addrs[1], new(big.Int).SetUint64(1_000_000), tests.SetNonce(0))
	evmTx10, _, err := evmkeeper.NewTxUtils(app.EVMKeeper).ConvertCosmosTxToEthereumTx(ctx, tx10)
	require.NoError(t, err)
	require.NotNil(t, evmTx10)

	txBz, err = evmTx10.MarshalBinary()
	require.NoError(t, err)
	_, err = backend.SendRawTransaction(txBz)
	require.NoError(t, err)

	// execute recheck
	err = input.cometRPC.RecheckTx()
	require.NoError(t, err)

	// Acc: 0, Nonce: 6
	tx06, _ := tests.GenerateTransferERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[1], new(big.Int).SetUint64(1_000_000), tests.SetNonce(6))
	evmTx06, _, err := evmkeeper.NewTxUtils(app.EVMKeeper).ConvertCosmosTxToEthereumTx(ctx, tx06)
	require.NoError(t, err)
	require.NotNil(t, evmTx06)

	txBz, err = evmTx06.MarshalBinary()
	require.NoError(t, err)
	_, err = backend.SendRawTransaction(txBz)
	require.NoError(t, err)

	// execute recheck
	err = input.cometRPC.RecheckTx()
	require.NoError(t, err)

	// 1 in pending, and 2 in queued
	txPool, err := backend.TxPoolInspect()
	require.NoError(t, err)
	require.Len(t, txPool["pending"][addrs[1].Hex()], 1)
	require.Len(t, txPool["queued"][addrs[0].Hex()], 2)

	// sending Nonce 3 should make Nonce 4 to be pending

	// Acc: 0, Nonce: 3
	tx03, _ := tests.GenerateTransferERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[1], new(big.Int).SetUint64(1_000_000), tests.SetNonce(3))
	evmTx03, _, err := evmkeeper.NewTxUtils(app.EVMKeeper).ConvertCosmosTxToEthereumTx(ctx, tx03)
	require.NoError(t, err)
	require.NotNil(t, evmTx03)

	txBz, err = evmTx03.MarshalBinary()
	require.NoError(t, err)
	_, err = backend.SendRawTransaction(txBz)
	require.NoError(t, err)

	// execute recheck
	err = input.cometRPC.RecheckTx()
	require.NoError(t, err)

	// 3 in pending and 1 in queued
	txPool, err = backend.TxPoolInspect()
	require.NoError(t, err)
	require.Len(t, txPool["pending"][addrs[0].Hex()], 2)
	require.Len(t, txPool["pending"][addrs[1].Hex()], 1)
	require.Len(t, txPool["queued"][addrs[0].Hex()], 1)

	// sending Nonce 5 should make Nonce 6 to be pending

	// Acc: 0, Nonce: 5
	tx05, _ := tests.GenerateTransferERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[1], new(big.Int).SetUint64(1_000_000), tests.SetNonce(5))
	evmTx05, _, err := evmkeeper.NewTxUtils(app.EVMKeeper).ConvertCosmosTxToEthereumTx(ctx, tx05)
	require.NoError(t, err)
	require.NotNil(t, evmTx05)

	txBz, err = evmTx05.MarshalBinary()
	require.NoError(t, err)
	_, err = backend.SendRawTransaction(txBz)
	require.NoError(t, err)

	// execute recheck
	err = input.cometRPC.RecheckTx()
	require.NoError(t, err)

	// 5 in pending and 0 in queued
	txPool, err = backend.TxPoolInspect()
	require.NoError(t, err)
	require.Len(t, txPool["pending"][addrs[0].Hex()], 4)
	require.Len(t, txPool["pending"][addrs[1].Hex()], 1)
	require.Empty(t, txPool["queued"])
}
