package backend_test

import (
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"

	rpctypes "github.com/initia-labs/minievm/jsonrpc/types"
	"github.com/initia-labs/minievm/tests"
	"github.com/initia-labs/minievm/x/evm/contracts/erc20"
	evmtypes "github.com/initia-labs/minievm/x/evm/types"
)

func Test_GasPrice(t *testing.T) {
	input := setupBackend(t)
	app, _, backend, _, _ := input.app, input.addrs, input.backend, input.addrs, input.privKeys

	ctx, err := app.CreateQueryContext(0, false)
	require.NoError(t, err)
	params, err := app.OPChildKeeper.Params.Get(ctx)
	require.NoError(t, err)

	minGasPrice := params.MinGasPrices[0].Amount.TruncateInt().BigInt()

	time.Sleep(3*time.Second + 500*time.Millisecond)

	gasPrice, err := backend.GasPrice()
	require.NoError(t, err)
	require.Equal(t, minGasPrice, gasPrice.ToInt())
}

func Test_MaxPriorityFeePerGas(t *testing.T) {
	input := setupBackend(t)
	_, _, backend, _, _ := input.app, input.addrs, input.backend, input.addrs, input.privKeys

	maxPriorityFeePerGas, err := backend.MaxPriorityFeePerGas()
	require.NoError(t, err)
	require.Equal(t, big.NewInt(0), maxPriorityFeePerGas.ToInt())
}

func Test_EstimateGas(t *testing.T) {
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
	contractEVMAddr := common.BytesToAddress(contractAddr)
	tx, _ = tests.GenerateMintERC20Tx(t, app, privKeys[0], contractEVMAddr, addrs[0], new(big.Int).SetUint64(1_000_000_000_000))
	tx2, _ := tests.GenerateTransferERC20Tx(t, app, privKeys[0], contractEVMAddr, addrs[1], new(big.Int).SetUint64(1_000_000), tests.SetNonce(2))
	_, finalizeRes = tests.ExecuteTxs(t, app, tx, tx2)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)
	tests.CheckTxResult(t, finalizeRes.TxResults[1], true)

	// call transfer function
	abi, err := erc20.Erc20MetaData.GetAbi()
	require.NoError(t, err)

	inputBz, err := abi.Pack("transfer", addrs[1], big.NewInt(1_000_000))
	require.NoError(t, err)

	time.Sleep(3 * time.Second)

	// query to latest block
	gasEstimated, err := backend.EstimateGas(rpctypes.TransactionArgs{
		From:  &addrs[0],
		To:    &contractEVMAddr,
		Input: (*hexutil.Bytes)(&inputBz),
		Value: nil,
		Nonce: nil,
	}, nil, nil)
	require.NoError(t, err)
	require.Greater(t, uint64(gasEstimated), uint64(finalizeRes.TxResults[1].GasUsed))

	gasEstimated, err = backend.EstimateGas(rpctypes.TransactionArgs{
		From:  &addrs[0],
		To:    &contractEVMAddr,
		Input: (*hexutil.Bytes)(&inputBz),
		Value: nil,
		Nonce: nil,
		AccessList: &types.AccessList{
			types.AccessTuple{
				Address: contractEVMAddr,
				StorageKeys: []common.Hash{
					common.HexToHash("0x0"),
					common.HexToHash("0x1"),
					common.HexToHash("0x2"),
					common.HexToHash("0x3"),
				},
			},
		},
	}, nil, nil)
	require.NoError(t, err)
	require.Greater(t, uint64(gasEstimated), uint64(finalizeRes.TxResults[1].GasUsed))

}
