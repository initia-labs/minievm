package backend_test

import (
	"math/big"
	"testing"

	cmttypes "github.com/cometbft/cometbft/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/initia-labs/minievm/tests"
	evmtypes "github.com/initia-labs/minievm/x/evm/types"

	"github.com/stretchr/testify/require"
)

func Test_CosmosTxHashByTxHash(t *testing.T) {
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
	tx, evmTxHash := tests.GenerateMintERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[0], new(big.Int).SetUint64(1_000_000_000_000))
	tx2, evmTxHash2 := tests.GenerateTransferERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[1], new(big.Int).SetUint64(1_000_000), tests.SetNonce(2))
	_, finalizeRes = tests.ExecuteTxs(t, app, tx, tx2)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)
	tests.CheckTxResult(t, finalizeRes.TxResults[1], true)

	txBytes, err := app.TxEncode(tx)
	require.NoError(t, err)
	txBytes2, err := app.TxEncode(tx2)
	require.NoError(t, err)

	cmtTx := cmttypes.Tx(txBytes)
	cosmosTxHash := cmtTx.Hash()
	cmtTx2 := cmttypes.Tx(txBytes2)
	cosmosTxHash2 := cmtTx2.Hash()

	txHash, err := backend.CosmosTxHashByTxHash(evmTxHash)
	require.NoError(t, err)
	require.Equal(t, cosmosTxHash, txHash)

	txHash2, err := backend.CosmosTxHashByTxHash(evmTxHash2)
	require.NoError(t, err)
	require.Equal(t, cosmosTxHash2, txHash2)
}

func Test_TxHashByCosmosTxHash(t *testing.T) {
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
	tx, evmTxHash := tests.GenerateMintERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[0], new(big.Int).SetUint64(1_000_000_000_000))
	tx2, evmTxHash2 := tests.GenerateTransferERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[1], new(big.Int).SetUint64(1_000_000), tests.SetNonce(2))
	_, finalizeRes = tests.ExecuteTxs(t, app, tx, tx2)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)
	tests.CheckTxResult(t, finalizeRes.TxResults[1], true)

	txBytes, err := app.TxEncode(tx)
	require.NoError(t, err)
	txBytes2, err := app.TxEncode(tx2)
	require.NoError(t, err)

	cmtTx := cmttypes.Tx(txBytes)
	cosmosTxHash := cmtTx.Hash()
	cmtTx2 := cmttypes.Tx(txBytes2)
	cosmosTxHash2 := cmtTx2.Hash()

	txHash, err := backend.TxHashByCosmosTxHash(cosmosTxHash)
	require.NoError(t, err)
	require.Equal(t, evmTxHash, txHash)

	txHash2, err := backend.TxHashByCosmosTxHash(cosmosTxHash2)
	require.NoError(t, err)
	require.Equal(t, evmTxHash2, txHash2)
}
