package keeper_test

import (
	"fmt"
	"testing"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/initia-labs/minievm/x/evm/contracts/counter"
	"github.com/initia-labs/minievm/x/evm/contracts/i_cosmos"
	"github.com/initia-labs/minievm/x/evm/keeper"
	"github.com/initia-labs/minievm/x/evm/types"
	"github.com/stretchr/testify/require"
)

func Test_ExecuteCosmosMessage(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()
	_, _, addr2 := keyPubAddr()
	evmAddr := common.BytesToAddress(addr.Bytes())

	erc20Keeper, err := keeper.NewERC20Keeper(&input.EVMKeeper)
	require.NoError(t, err)

	// mint native coin
	err = erc20Keeper.MintCoins(ctx, addr, sdk.NewCoins(
		sdk.NewCoin("bar", math.NewInt(200)),
	))
	require.NoError(t, err)

	abi, err := i_cosmos.ICosmosMetaData.GetAbi()
	require.NoError(t, err)

	inputBz, err := abi.Pack("execute_cosmos", fmt.Sprintf(`
		{
			"@type": "/cosmos.bank.v1beta1.MsgSend",
			"from_address": "%s",
			"to_address": "%s",
			"amount": [
				{
					"denom": "bar",
					"amount": "100"
				}
			]
		}
	`, addr, addr2))
	require.NoError(t, err)

	_, _, err = input.EVMKeeper.EVMCall(ctx, evmAddr, types.CosmosPrecompileAddress, inputBz, nil)
	require.NoError(t, err)

	balance := input.BankKeeper.GetBalance(ctx, addr2, "bar")
	require.Equal(t, math.NewInt(100), balance.Amount)
}

func Test_QueryCosmosMessage(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()
	evmAddr := common.BytesToAddress(addr.Bytes())

	erc20Keeper, err := keeper.NewERC20Keeper(&input.EVMKeeper)
	require.NoError(t, err)

	// mint native coin
	err = erc20Keeper.MintCoins(ctx, addr, sdk.NewCoins(
		sdk.NewCoin("bar", math.NewInt(200)),
	))
	require.NoError(t, err)

	abi, err := i_cosmos.ICosmosMetaData.GetAbi()
	require.NoError(t, err)

	inputBz, err := abi.Pack("query_cosmos", "/cosmos.bank.v1beta1.Query/Balance", fmt.Sprintf(`{
		"address": "%s",
		"denom": "bar"
	}`, addr))
	require.NoError(t, err)

	retBz, _, err := input.EVMKeeper.EVMCall(ctx, evmAddr, types.CosmosPrecompileAddress, inputBz, nil)
	require.NoError(t, err)

	unpackedRet, err := abi.Methods["query_cosmos"].Outputs.Unpack(retBz)
	require.NoError(t, err)

	var ret banktypes.QueryBalanceResponse
	err = input.EncodingConfig.Codec.UnmarshalJSON([]byte(unpackedRet[0].(string)), &ret)
	require.NoError(t, err)

	require.Equal(t, math.NewInt(200), ret.Balance.Amount)
}

func Test_QueryCosmosFromContract(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()

	erc20Keeper, err := keeper.NewERC20Keeper(&input.EVMKeeper)
	require.NoError(t, err)

	// mint native coin
	err = erc20Keeper.MintCoins(ctx, addr, sdk.NewCoins(
		sdk.NewCoin("bar", math.NewInt(200)),
	))
	require.NoError(t, err)

	counterBz, err := hexutil.Decode(counter.CounterBin)
	require.NoError(t, err)

	caller := common.BytesToAddress(addr.Bytes())
	retBz, contractAddr, _, err := input.EVMKeeper.EVMCreate(ctx, caller, counterBz, nil)
	require.NoError(t, err)
	require.NotEmpty(t, retBz)
	require.Len(t, contractAddr, 20)

	abi, err := counter.CounterMetaData.GetAbi()
	require.NoError(t, err)

	inputBz, err := abi.Pack("query_cosmos", "/cosmos.bank.v1beta1.Query/Balance", fmt.Sprintf(`{
		"address": "%s",
		"denom": "bar"
	}`, addr))
	require.NoError(t, err)

	retBz, _, err = input.EVMKeeper.EVMCall(ctx, contractAddr, types.CosmosPrecompileAddress, inputBz, nil)
	require.NoError(t, err)

	unpackedRet, err := abi.Methods["query_cosmos"].Outputs.Unpack(retBz)
	require.NoError(t, err)

	var ret banktypes.QueryBalanceResponse
	err = input.EncodingConfig.Codec.UnmarshalJSON([]byte(unpackedRet[0].(string)), &ret)
	require.NoError(t, err)

	require.Equal(t, math.NewInt(200), ret.Balance.Amount)
}

func Test_ToDenom(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()
	evmAddr := common.BytesToAddress(addr.Bytes())

	erc20Keeper, err := keeper.NewERC20Keeper(&input.EVMKeeper)
	require.NoError(t, err)

	// mint native coin
	err = erc20Keeper.MintCoins(ctx, addr, sdk.NewCoins(
		sdk.NewCoin("bar", math.NewInt(200)),
	))
	require.NoError(t, err)

	contractAddr, err := input.EVMKeeper.GetContractAddrByDenom(ctx, "bar")
	require.NoError(t, err)

	abi, err := i_cosmos.ICosmosMetaData.GetAbi()
	require.NoError(t, err)

	inputBz, err := abi.Pack("to_denom", contractAddr)
	require.NoError(t, err)

	retBz, _, err := input.EVMKeeper.EVMCall(ctx, evmAddr, types.CosmosPrecompileAddress, inputBz, nil)
	require.NoError(t, err)

	unpackedRet, err := abi.Methods["to_denom"].Outputs.Unpack(retBz)
	require.NoError(t, err)

	require.Equal(t, "bar", unpackedRet[0].(string))
}

func Test_ToERC20(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()
	evmAddr := common.BytesToAddress(addr.Bytes())

	erc20Keeper, err := keeper.NewERC20Keeper(&input.EVMKeeper)
	require.NoError(t, err)

	// mint native coin
	err = erc20Keeper.MintCoins(ctx, addr, sdk.NewCoins(
		sdk.NewCoin("bar", math.NewInt(200)),
	))
	require.NoError(t, err)

	contractAddr, err := input.EVMKeeper.GetContractAddrByDenom(ctx, "bar")
	require.NoError(t, err)

	abi, err := i_cosmos.ICosmosMetaData.GetAbi()
	require.NoError(t, err)

	inputBz, err := abi.Pack("to_erc20", "bar")
	require.NoError(t, err)

	retBz, _, err := input.EVMKeeper.EVMCall(ctx, evmAddr, types.CosmosPrecompileAddress, inputBz, nil)
	require.NoError(t, err)

	unpackedRet, err := abi.Methods["to_erc20"].Outputs.Unpack(retBz)
	require.NoError(t, err)

	require.Equal(t, contractAddr, unpackedRet[0].(common.Address))
}
