package keeper_test

import (
	"strings"
	"testing"

	"cosmossdk.io/math"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/eth/tracers/logger"
	"github.com/holiman/uint256"
	"golang.org/x/crypto/sha3"

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
	retBz, contractAddr, _, err := input.EVMKeeper.EVMCreate(ctx, caller, counterBz, nil, nil)
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

	retBz, contractAddr, _, err := input.EVMKeeper.EVMCreateWithTracer(ctx, caller, counterBz, uint256.NewInt(100), nil, nil, tracer)
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
	retBz, contractAddr, _, err := input.EVMKeeper.EVMCreate(ctx, caller, counterBz, nil, nil)
	require.NoError(t, err)
	require.NotEmpty(t, retBz)
	require.Len(t, contractAddr, 20)

	parsed, err := counter.CounterMetaData.GetAbi()
	require.NoError(t, err)

	queryInputBz, err := parsed.Pack("count")
	require.NoError(t, err)

	queryRes, logs, err := input.EVMKeeper.EVMCall(ctx, caller, contractAddr, queryInputBz, nil, nil)
	require.NoError(t, err)
	require.Equal(t, uint256.NewInt(0).Bytes32(), [32]byte(queryRes))
	require.Empty(t, logs)

	inputBz, err := parsed.Pack("increase")
	require.NoError(t, err)

	// call with value
	res, logs, err := input.EVMKeeper.EVMCall(ctx, caller, contractAddr, inputBz, uint256.NewInt(100), nil)
	require.NoError(t, err)
	require.Empty(t, res)
	require.NotEmpty(t, logs)

	// check balance
	balance, err := input.EVMKeeper.ERC20Keeper().GetBalance(ctx, contractAddr.Bytes(), sdk.DefaultBondDenom)
	require.NoError(t, err)
	require.Equal(t, balance, math.NewInt(100))

	queryRes, logs, err = input.EVMKeeper.EVMCall(ctx, caller, contractAddr, queryInputBz, nil, nil)
	require.NoError(t, err)
	require.Equal(t, uint256.NewInt(1).Bytes32(), [32]byte(queryRes))
	require.Empty(t, logs)

	// calling not existing function
	erc20ABI, err := erc20.Erc20MetaData.GetAbi()
	require.NoError(t, err)

	queryInputBz, err = erc20ABI.Pack("balanceOf", caller)
	require.NoError(t, err)

	_, _, err = input.EVMKeeper.EVMCall(ctx, caller, contractAddr, queryInputBz, nil, nil)
	require.ErrorContains(t, err, types.ErrReverted.Error())
}

func Test_GetHash(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()

	// change number of retain block hashes to 257
	params, err := input.EVMKeeper.Params.Get(ctx)
	require.NoError(t, err)
	params.NumRetainBlockHashes = 257
	err = input.EVMKeeper.Params.Set(ctx, params)
	require.NoError(t, err)

	// fund addr
	input.Faucet.Fund(ctx, addr, sdk.NewInt64Coin(sdk.DefaultBondDenom, 1000000000))

	counterBz, err := hexutil.Decode(counter.CounterBin)
	require.NoError(t, err)

	caller := common.BytesToAddress(addr.Bytes())
	retBz, contractAddr, _, err := input.EVMKeeper.EVMCreate(ctx, caller, counterBz, nil, nil)
	require.NoError(t, err)
	require.NotEmpty(t, retBz)
	require.Len(t, contractAddr, 20)

	parsed, err := counter.CounterMetaData.GetAbi()
	require.NoError(t, err)

	// set block hash
	hash99 := sha3.Sum256([]byte("block99"))
	hash100 := sha3.Sum256([]byte("block100"))
	hash101 := sha3.Sum256([]byte("block101"))
	hash356 := sha3.Sum256([]byte("block356"))
	require.NoError(t, input.EVMKeeper.TrackBlockHash(ctx, 356, common.BytesToHash(hash99[:])))
	require.NoError(t, input.EVMKeeper.TrackBlockHash(ctx, 100, common.BytesToHash(hash100[:])))
	require.NoError(t, input.EVMKeeper.TrackBlockHash(ctx, 101, common.BytesToHash(hash101[:])))
	require.NoError(t, input.EVMKeeper.TrackBlockHash(ctx, 356, common.BytesToHash(hash356[:])))

	// set current block height
	ctx = ctx.WithBlockHeight(357)

	// query 99 should return empty hash because only resent 257 block hashes are stored
	queryInputBz, err := parsed.Pack("get_blockhash", uint64(99))
	require.NoError(t, err)

	queryRes, logs, err := input.EVMKeeper.EVMCall(ctx, caller, contractAddr, queryInputBz, nil, nil)
	require.NoError(t, err)
	require.Equal(t, [32]byte{}, [32]byte(queryRes))
	require.Empty(t, logs)

	// valid query
	queryInputBz, err = parsed.Pack("get_blockhash", uint64(100))
	require.NoError(t, err)

	queryRes, logs, err = input.EVMKeeper.EVMCall(ctx, caller, contractAddr, queryInputBz, nil, nil)
	require.NoError(t, err)
	require.Equal(t, hash100, [32]byte(queryRes))
	require.Empty(t, logs)

	// vaild query
	queryInputBz, err = parsed.Pack("get_blockhash", uint64(101))
	require.NoError(t, err)

	queryRes, logs, err = input.EVMKeeper.EVMCall(ctx, caller, contractAddr, queryInputBz, nil, nil)
	require.NoError(t, err)
	require.Equal(t, hash101, [32]byte(queryRes))
	require.Empty(t, logs)

	// vaild query
	queryInputBz, err = parsed.Pack("get_blockhash", uint64(356))
	require.NoError(t, err)

	queryRes, logs, err = input.EVMKeeper.EVMCall(ctx, caller, contractAddr, queryInputBz, nil, nil)
	require.NoError(t, err)
	require.Equal(t, hash356, [32]byte(queryRes))
	require.Empty(t, logs)

	// return empty bytes if block height is greater than current block height
	queryInputBz, err = parsed.Pack("get_blockhash", uint64(357))
	require.NoError(t, err)

	queryRes, logs, err = input.EVMKeeper.EVMCall(ctx, caller, contractAddr, queryInputBz, nil, nil)
	require.NoError(t, err)
	require.Equal(t, [32]byte{}, [32]byte(queryRes))
	require.Empty(t, logs)
}
