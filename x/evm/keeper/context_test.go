package keeper_test

import (
	"strings"
	"testing"

	"cosmossdk.io/math"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/eth/tracers/logger"
	"github.com/holiman/uint256"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/initia-labs/minievm/x/evm/contracts/counter"
	"github.com/initia-labs/minievm/x/evm/contracts/erc20"
	"github.com/initia-labs/minievm/x/evm/types"

	"github.com/stretchr/testify/require"
)

func Test_Create(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()

	counterBz, err := hexutil.Decode(counter.CounterBin)
	require.NoError(t, err)

	caller := common.BytesToAddress(addr.Bytes())
	retBz, contractAddr, _, err := input.EVMKeeper.EVMCreate(ctx, caller, counterBz, nil)
	require.NoError(t, err)
	require.NotEmpty(t, retBz)
	require.Len(t, contractAddr, 20)
}

func Test_CreateWithValue(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()

	// fund addr
	input.Faucet.Fund(ctx, addr, sdk.NewInt64Coin(sdk.DefaultBondDenom, 1000000000))

	counterBz, err := hexutil.Decode(counter.CounterBin)
	require.NoError(t, err)

	caller := common.BytesToAddress(addr.Bytes())
	tracerOutput := new(strings.Builder)
	tracer := logger.NewJSONLogger(&logger.Config{
		EnableMemory:     false,
		DisableStack:     false,
		DisableStorage:   false,
		EnableReturnData: true,
	}, tracerOutput)

	retBz, contractAddr, _, err := input.EVMKeeper.EVMCreateWithTracer(ctx, caller, counterBz, uint256.NewInt(100), nil, tracer)
	require.NoError(t, err)
	require.NotEmpty(t, retBz)
	require.Len(t, contractAddr, 20)

	// check balance
	balance, err := input.EVMKeeper.ERC20Keeper().GetBalance(ctx, contractAddr.Bytes(), sdk.DefaultBondDenom)
	require.NoError(t, err)
	require.Equal(t, balance, math.NewInt(100))
}

func Test_Call(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()

	// fund addr
	input.Faucet.Fund(ctx, addr, sdk.NewInt64Coin(sdk.DefaultBondDenom, 1000000000))

	counterBz, err := hexutil.Decode(counter.CounterBin)
	require.NoError(t, err)

	caller := common.BytesToAddress(addr.Bytes())
	retBz, contractAddr, _, err := input.EVMKeeper.EVMCreate(ctx, caller, counterBz, nil)
	require.NoError(t, err)
	require.NotEmpty(t, retBz)
	require.Len(t, contractAddr, 20)

	parsed, err := counter.CounterMetaData.GetAbi()
	require.NoError(t, err)

	queryInputBz, err := parsed.Pack("count")
	require.NoError(t, err)

	queryRes, logs, err := input.EVMKeeper.EVMCall(ctx, caller, contractAddr, queryInputBz, nil)
	require.NoError(t, err)
	require.Equal(t, uint256.NewInt(0).Bytes32(), [32]byte(queryRes))
	require.Empty(t, logs)

	inputBz, err := parsed.Pack("increase")
	require.NoError(t, err)

	// call with value
	res, logs, err := input.EVMKeeper.EVMCall(ctx, caller, contractAddr, inputBz, uint256.NewInt(100))
	require.NoError(t, err)
	require.Empty(t, res)
	require.NotEmpty(t, logs)

	// check balance
	balance, err := input.EVMKeeper.ERC20Keeper().GetBalance(ctx, contractAddr.Bytes(), sdk.DefaultBondDenom)
	require.NoError(t, err)
	require.Equal(t, balance, math.NewInt(100))

	queryRes, logs, err = input.EVMKeeper.EVMCall(ctx, caller, contractAddr, queryInputBz, nil)
	require.NoError(t, err)
	require.Equal(t, uint256.NewInt(1).Bytes32(), [32]byte(queryRes))
	require.Empty(t, logs)

	// calling not existing function
	erc20ABI, err := erc20.Erc20MetaData.GetAbi()
	require.NoError(t, err)

	queryInputBz, err = erc20ABI.Pack("balanceOf", caller)
	require.NoError(t, err)

	_, _, err = input.EVMKeeper.EVMCall(ctx, caller, contractAddr, queryInputBz, nil)
	require.ErrorContains(t, err, types.ErrReverted.Error())
}
