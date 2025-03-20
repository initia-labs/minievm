package keeper_test

import (
	"fmt"
	"strings"
	"testing"

	"cosmossdk.io/math"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/vm"
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

func Test_Create2(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()

	counterBz, err := hexutil.Decode(counter.CounterBin)
	require.NoError(t, err)

	caller := common.BytesToAddress(addr.Bytes())

	retBz, contractAddr, _, err := input.EVMKeeper.EVMCreate2(ctx, caller, counterBz, nil, uint256.NewInt(1), nil)
	require.NoError(t, err)
	require.NotEmpty(t, retBz)
	require.Len(t, contractAddr, 20)

	_, _, _, err = input.EVMKeeper.EVMCreate2(ctx, caller, counterBz, nil, uint256.NewInt(1), nil)
	require.ErrorContains(t, err, vm.ErrContractAddressCollision.Error())
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
	require.Len(t, logs, int(2))

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
	require.NoError(t, input.EVMKeeper.TrackBlockHash(ctx, 99, common.BytesToHash(hash99[:])))
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

	// valid query
	queryInputBz, err = parsed.Pack("get_blockhash", uint64(101))
	require.NoError(t, err)

	queryRes, logs, err = input.EVMKeeper.EVMCall(ctx, caller, contractAddr, queryInputBz, nil, nil)
	require.NoError(t, err)
	require.Equal(t, hash101, [32]byte(queryRes))
	require.Empty(t, logs)

	// valid query
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

func Test_RecursiveDepth(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()

	counterBz, err := hexutil.Decode(counter.CounterBin)
	require.NoError(t, err)

	// deploy counter contract
	caller := common.BytesToAddress(addr.Bytes())
	retBz, contractAddr, _, err := input.EVMKeeper.EVMCreate(ctx, caller, counterBz, nil, nil)
	require.NoError(t, err)
	require.NotEmpty(t, retBz)
	require.Len(t, contractAddr, 20)

	// call recursive function
	parsed, err := counter.CounterMetaData.GetAbi()
	require.NoError(t, err)

	inputBz, err := parsed.Pack("recursive", uint64(types.MAX_RECURSIVE_DEPTH-1))
	require.NoError(t, err)

	_, logs, err := input.EVMKeeper.EVMCall(ctx, caller, contractAddr, inputBz, nil, nil)
	require.NoError(t, err)
	require.Equal(t, 1<<types.MAX_RECURSIVE_DEPTH-1, len(logs))

	// exceed max recursive depth
	inputBz, err = parsed.Pack("recursive", uint64(types.MAX_RECURSIVE_DEPTH))
	require.NoError(t, err)

	_, _, err = input.EVMKeeper.EVMCall(ctx, caller, contractAddr, inputBz, nil, nil)
	require.ErrorContains(t, err, types.ErrExceedMaxRecursiveDepth.Error())
}

func Test_RevertAfterExecuteCosmos(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()

	counterBz, err := hexutil.Decode(counter.CounterBin)
	require.NoError(t, err)

	// deploy counter contract
	caller := common.BytesToAddress(addr.Bytes())
	retBz, contractAddr, _, err := input.EVMKeeper.EVMCreate(ctx, caller, counterBz, nil, nil)
	require.NoError(t, err)
	require.NotEmpty(t, retBz)
	require.Len(t, contractAddr, 20)

	// call execute cosmos function
	parsed, err := counter.CounterMetaData.GetAbi()
	require.NoError(t, err)

	denom := sdk.DefaultBondDenom
	amount := math.NewInt(1000000000)
	input.Faucet.Mint(ctx, contractAddr.Bytes(), sdk.NewCoin(denom, amount))

	// call execute_cosmos with revert
	inputBz, err := parsed.Pack("execute_cosmos",
		fmt.Sprintf(`{"@type":"/cosmos.bank.v1beta1.MsgSend","from_address":"%s","to_address":"%s","amount":[{"denom":"%s","amount":"%s"}]}`,
			sdk.AccAddress(contractAddr.Bytes()).String(),
			addr.String(), // caller
			denom,
			amount,
		),
		uint64(120_000),
		true,
	)
	require.NoError(t, err)

	_, _, err = input.EVMKeeper.EVMCall(ctx, caller, contractAddr, inputBz, nil, nil)
	require.ErrorContains(t, err, vm.ErrExecutionReverted.Error())
	require.ErrorContains(t, err, "revert reason dummy value for test")

	// check balance
	require.Equal(t, amount, input.BankKeeper.GetBalance(ctx, sdk.AccAddress(contractAddr.Bytes()), denom).Amount)
	require.Equal(t, math.ZeroInt(), input.BankKeeper.GetBalance(ctx, addr, denom).Amount)

	// call execute_cosmos without revert
	inputBz, err = parsed.Pack("execute_cosmos",
		fmt.Sprintf(`{"@type":"/cosmos.bank.v1beta1.MsgSend","from_address":"%s","to_address":"%s","amount":[{"denom":"%s","amount":"%s"}]}`,
			sdk.AccAddress(contractAddr.Bytes()).String(),
			addr.String(), // caller
			denom,
			amount,
		),
		uint64(120_000),
		false,
	)
	require.NoError(t, err)

	_, _, err = input.EVMKeeper.EVMCall(ctx, caller, contractAddr, inputBz, nil, nil)
	require.NoError(t, err)

	require.Equal(t, math.ZeroInt(), input.BankKeeper.GetBalance(ctx, sdk.AccAddress(contractAddr.Bytes()), denom).Amount)
	require.Equal(t, amount, input.BankKeeper.GetBalance(ctx, addr, denom).Amount)
}

func Test_ExecuteCosmosWithOptions(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()

	counterBz, err := hexutil.Decode(counter.CounterBin)
	require.NoError(t, err)

	// deploy counter contract
	caller := common.BytesToAddress(addr.Bytes())
	retBz, contractAddr, _, err := input.EVMKeeper.EVMCreate(ctx, caller, counterBz, nil, nil)
	require.NoError(t, err)
	require.NotEmpty(t, retBz)
	require.Len(t, contractAddr, 20)

	// call execute cosmos function
	parsed, err := counter.CounterMetaData.GetAbi()
	require.NoError(t, err)

	denom := sdk.DefaultBondDenom
	amount := math.NewInt(1000000000)
	input.Faucet.Mint(ctx, contractAddr.Bytes(), sdk.NewCoin(denom, amount))

	// case 1. call execute_cosmos with options by sending more than balance to revert the cosmos execute
	inputBz, err := parsed.Pack("execute_cosmos_with_options",
		fmt.Sprintf(`{"@type":"/cosmos.bank.v1beta1.MsgSend","from_address":"%s","to_address":"%s","amount":[{"denom":"%s","amount":"%s"}]}`,
			sdk.AccAddress(contractAddr.Bytes()).String(),
			addr.String(), // caller
			denom,
			amount.AddRaw(1),
		),
		uint64(120_000),
		true,
		uint64(100),
	)
	require.NoError(t, err)

	_, logs, err := input.EVMKeeper.EVMCall(ctx, caller, contractAddr, inputBz, nil, nil)
	require.NoError(t, err)
	require.Len(t, logs, 1)

	// check balance
	require.Equal(t, amount, input.BankKeeper.GetBalance(ctx, sdk.AccAddress(contractAddr.Bytes()), denom).Amount)
	require.Equal(t, math.ZeroInt(), input.BankKeeper.GetBalance(ctx, addr, denom).Amount)

	// check events
	events := ctx.EventManager().Events()
	var shouldEnter bool
	for _, event := range events {
		if event.Type == types.EventTypeSubmsg {
			shouldEnter = true
			require.Equal(t, "false", event.Attributes[0].Value)
			break
		}
	}
	require.True(t, shouldEnter)

	// check callback logs
	log := logs[0]
	require.Equal(t, contractAddr.Hex(), log.Address)
	require.Equal(t, parsed.Events["callback_received"].ID.Hex(), log.Topics[0])
	require.Equal(t, "0x00000000000000000000000000000000000000000000000000000000000000640000000000000000000000000000000000000000000000000000000000000000", log.Data)

	// case 2. call execute_cosmos with options by sending less than balance to revert the cosmos execute but not allow revert
	ctx = ctx.WithEventManager(sdk.NewEventManager())
	inputBz, err = parsed.Pack("execute_cosmos_with_options",
		fmt.Sprintf(`{"@type":"/cosmos.bank.v1beta1.MsgSend","from_address":"%s","to_address":"%s","amount":[{"denom":"%s","amount":"%s"}]}`,
			sdk.AccAddress(contractAddr.Bytes()).String(),
			addr.String(), // caller
			denom,
			amount.AddRaw(1),
		),
		uint64(120_000),
		false,
		uint64(101),
	)
	require.NoError(t, err)

	_, _, err = input.EVMKeeper.EVMCall(ctx, caller, contractAddr, inputBz, nil, nil)
	require.Error(t, err)

	// case 3. call execute_cosmos with options by sending valid amount
	ctx = ctx.WithEventManager(sdk.NewEventManager())
	inputBz, err = parsed.Pack("execute_cosmos_with_options",
		fmt.Sprintf(`{"@type":"/cosmos.bank.v1beta1.MsgSend","from_address":"%s","to_address":"%s","amount":[{"denom":"%s","amount":"%s"}]}`,
			sdk.AccAddress(contractAddr.Bytes()).String(),
			addr.String(), // caller
			denom,
			amount,
		),
		uint64(120_000),
		true,
		uint64(102),
	)
	require.NoError(t, err)

	_, logs, err = input.EVMKeeper.EVMCall(ctx, caller, contractAddr, inputBz, nil, nil)
	require.NoError(t, err)
	require.Equal(t, len(logs), 1)

	// check events
	events = ctx.EventManager().Events()
	shouldEnter = false
	for _, event := range events {
		if event.Type == types.EventTypeSubmsg {
			shouldEnter = true
			require.Equal(t, "true", event.Attributes[0].Value)
			break
		}
	}
	require.True(t, shouldEnter)

	// check callback logs
	log = logs[len(logs)-1]
	require.Equal(t, contractAddr.Hex(), log.Address)
	require.Equal(t, parsed.Events["callback_received"].ID.Hex(), log.Topics[0])
	require.Equal(t, "0x00000000000000000000000000000000000000000000000000000000000000660000000000000000000000000000000000000000000000000000000000000001", log.Data)
}

func Test_Recursive_Audit_ExecuteRequestsNotCleanedOnRevert(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()

	counterBz, err := hexutil.Decode(counter.CounterBin)
	require.NoError(t, err)

	// deploy counter contract
	caller := common.BytesToAddress(addr.Bytes())
	retBz, contractAddr, _, err := input.EVMKeeper.EVMCreate(ctx, caller, counterBz, nil, nil)
	require.NoError(t, err)
	require.NotEmpty(t, retBz)
	require.Len(t, contractAddr, 20)

	// call recursive_revert function
	parsed, err := counter.CounterMetaData.GetAbi()
	require.NoError(t, err)

	inputBz, err := parsed.Pack("recursive_revert", uint64(1))
	require.NoError(t, err)

	_, logs, err := input.EVMKeeper.EVMCall(ctx, caller, contractAddr, inputBz, nil, nil)
	require.NoError(t, err)
	require.Len(t, logs, 1)

	// check logs
	log := logs[0]
	require.Equal(t, contractAddr.Hex(), log.Address)
	require.Equal(t, parsed.Events["recursive_called"].ID.Hex(), log.Topics[0])
	require.Equal(t, "0x0000000000000000000000000000000000000000000000000000000000000001", log.Data)
}
