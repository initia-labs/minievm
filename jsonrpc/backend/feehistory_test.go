package backend_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/initia-labs/minievm/tests"
	evmtypes "github.com/initia-labs/minievm/x/evm/types"

	"github.com/stretchr/testify/require"
)

func Test_FeeHistory(t *testing.T) {
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
	_, finalizeRes = tests.ExecuteTxs(t, app, tx)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	// multiple transfers
	ctx, closer, err := app.CreateQueryContext(0, false)
	if closer != nil {
		defer closer.Close()
	}
	require.NoError(t, err)

	params, err := app.OPChildKeeper.Params.Get(ctx)
	require.NoError(t, err)

	minGasPrice := params.MinGasPrices[0]

	gasLimit := uint64(1_000_000)
	baseFee_ := minGasPrice.Amount.TruncateInt().BigInt()
	baseFeeCap := new(big.Int).Mul(baseFee_, big.NewInt(int64(gasLimit)))
	gasTipCap := new(big.Int).Mul(minGasPrice.Amount.TruncateInt().BigInt(), big.NewInt(1_000))
	gasFeeCap := new(big.Int).Add(baseFeeCap, gasTipCap) // add extra tip

	nonce := 2
	txs := make([]sdk.Tx, 5)
	for i := 0; i < 5; i++ {
		txs[i], _ = tests.GenerateTransferERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[1], new(big.Int).SetUint64(1_000_000), tests.SetNonce(uint64(nonce+i)), tests.SetGasLimit(gasLimit), tests.SetGasFeeCap(gasFeeCap), tests.SetGasTipCap(gasTipCap))
	}
	_, finalizeRes = tests.ExecuteTxs(t, app, txs...)

	for i := 0; i < 5; i++ {
		tests.CheckTxResult(t, finalizeRes.TxResults[i], true)
	}

	header, err := backend.GetHeaderByNumber(rpc.BlockNumber(app.LastBlockHeight()))
	require.NoError(t, err)

	gasUsedRatioExp := float64(header.GasUsed) / float64(header.GasLimit)
	oldestBlock, reward, baseFee, gasUsedRatio, blobBaseFee, blobGasUsedRatio, err := backend.FeeHistory(1, rpc.BlockNumber(app.LastBlockHeight()), []float64{40, 60})
	require.NoError(t, err)
	require.Equal(t, uint64(app.LastBlockHeight()), oldestBlock.Uint64())
	require.Equal(t, [][]*big.Int{{gasTipCap, gasTipCap}}, reward)
	require.Equal(t, []*big.Int{baseFee_, baseFee_}, baseFee)
	require.Equal(t, []float64{gasUsedRatioExp}, gasUsedRatio)
	require.Equal(t, []*big.Int{big.NewInt(0), big.NewInt(0)}, blobBaseFee)
	require.Equal(t, []float64{0}, blobGasUsedRatio)
}
