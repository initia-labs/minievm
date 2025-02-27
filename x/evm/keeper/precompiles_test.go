package keeper_test

import (
	"encoding/json"
	"fmt"
	"math/big"
	"testing"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	abiapi "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/vm"

	"github.com/initia-labs/minievm/x/evm/contracts/connect_oracle"
	"github.com/initia-labs/minievm/x/evm/contracts/counter"
	"github.com/initia-labs/minievm/x/evm/contracts/i_cosmos"
	"github.com/initia-labs/minievm/x/evm/contracts/i_jsonutils"
	"github.com/initia-labs/minievm/x/evm/keeper"
	"github.com/initia-labs/minievm/x/evm/types"

	"github.com/stretchr/testify/require"

	connecttypes "github.com/skip-mev/connect/v2/pkg/types"
	marketmaptypes "github.com/skip-mev/connect/v2/x/marketmap/types"
	oracletypes "github.com/skip-mev/connect/v2/x/oracle/types"
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
	`, addr, addr2), uint64(150_000))
	require.NoError(t, err)

	_, _, err = input.EVMKeeper.EVMCall(ctx, evmAddr, types.CosmosPrecompileAddress, inputBz, nil, nil)
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

	retBz, _, err := input.EVMKeeper.EVMCall(ctx, evmAddr, types.CosmosPrecompileAddress, inputBz, nil, nil)
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
	retBz, contractAddr, _, err := input.EVMKeeper.EVMCreate(ctx, caller, counterBz, nil, nil)
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

	retBz, _, err = input.EVMKeeper.EVMCall(ctx, contractAddr, types.CosmosPrecompileAddress, inputBz, nil, nil)
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

	retBz, _, err := input.EVMKeeper.EVMCall(ctx, evmAddr, types.CosmosPrecompileAddress, inputBz, nil, nil)
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

	contractAddr, err := types.DenomToContractAddr(ctx, &input.EVMKeeper, "bar")
	require.NoError(t, err)

	abi, err := i_cosmos.ICosmosMetaData.GetAbi()
	require.NoError(t, err)

	inputBz, err := abi.Pack("to_erc20", "bar")
	require.NoError(t, err)

	retBz, _, err := input.EVMKeeper.EVMCall(ctx, evmAddr, types.CosmosPrecompileAddress, inputBz, nil, nil)
	require.NoError(t, err)

	unpackedRet, err := abi.Methods["to_erc20"].Outputs.Unpack(retBz)
	require.NoError(t, err)

	require.Equal(t, contractAddr, unpackedRet[0].(common.Address))
}

func Test_JSONMerge(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()
	evmAddr := common.BytesToAddress(addr.Bytes())

	abi, err := i_jsonutils.IJsonutilsMetaData.GetAbi()
	require.NoError(t, err)

	inputBz, err := abi.Pack("merge_json", `{"a": 1, "b": 2}`, `{"b": 3, "c": 4}`)
	require.NoError(t, err)

	retBz, _, err := input.EVMKeeper.EVMCall(ctx, evmAddr, types.JSONUtilsPrecompileAddress, inputBz, nil, nil)
	require.NoError(t, err)

	unpackedRet, err := abi.Methods["merge_json"].Outputs.Unpack(retBz)
	require.NoError(t, err)

	require.Equal(t, `{"a":1,"b":3,"c":4}`, unpackedRet[0].(string))
}

func Test_PrecompileRevertError(t *testing.T) {
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
			addr.String(), // try to call with wrong signer
			addr.String(), // caller
			denom,
			amount,
		),
		uint64(150_000),
		false,
	)
	require.NoError(t, err)

	_, _, err = input.EVMKeeper.EVMCall(ctx, caller, contractAddr, inputBz, nil, nil)
	require.ErrorContains(t, err, vm.ErrExecutionReverted.Error())
	require.ErrorContains(t, err, sdkerrors.ErrUnauthorized.Error())

	// check balance
	require.Equal(t, amount, input.BankKeeper.GetBalance(ctx, sdk.AccAddress(contractAddr.Bytes()), denom).Amount)
	require.Equal(t, math.ZeroInt(), input.BankKeeper.GetBalance(ctx, addr, denom).Amount)
}

func mustMarshalJSON(t *testing.T, obj interface{}) []byte {
	bz, err := json.Marshal(obj)
	require.NoError(t, err)
	return bz
}

func Test_JSONUnmarshalObject(t *testing.T) {
	testCases := []struct {
		name        string
		input       []byte
		expected    i_jsonutils.IJSONUtilsJSONObject
		expectedErr bool
	}{
		{
			name:     "empty map",
			input:    []byte(`{}`),
			expected: i_jsonutils.IJSONUtilsJSONObject{Elements: []i_jsonutils.IJSONUtilsJSONElement{}},
		},
		{
			name:     "simple map",
			input:    []byte(`{"a": 1, "b": 2}`),
			expected: i_jsonutils.IJSONUtilsJSONObject{Elements: []i_jsonutils.IJSONUtilsJSONElement{{Key: "a", Value: mustMarshalJSON(t, 1)}, {Key: "b", Value: mustMarshalJSON(t, 2)}}},
		},
		{
			name:     "simple map sorted",
			input:    []byte(`{"b": 1, "a": 2}`),
			expected: i_jsonutils.IJSONUtilsJSONObject{Elements: []i_jsonutils.IJSONUtilsJSONElement{{Key: "a", Value: mustMarshalJSON(t, 2)}, {Key: "b", Value: mustMarshalJSON(t, 1)}}},
		},
		{
			name:     "nested map",
			input:    []byte(`{"a": 1, "b": {"c": 2}}`),
			expected: i_jsonutils.IJSONUtilsJSONObject{Elements: []i_jsonutils.IJSONUtilsJSONElement{{Key: "a", Value: mustMarshalJSON(t, 1)}, {Key: "b", Value: mustMarshalJSON(t, map[string]int{"c": 2})}}},
		},
		{
			name:        "invalid json",
			input:       []byte(`{"a": 1, "b": 2`),
			expectedErr: true,
		},
		{
			name:        "slice",
			input:       []byte(`[1,2,3]`),
			expectedErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx, input := createDefaultTestInput(t)
			_, _, addr := keyPubAddr()
			evmAddr := common.BytesToAddress(addr.Bytes())

			abi, err := i_jsonutils.IJsonutilsMetaData.GetAbi()
			require.NoError(t, err)

			inputBz, err := abi.Pack("unmarshal_to_object", tc.input)
			require.NoError(t, err)

			retBz, _, err := input.EVMKeeper.EVMCall(ctx, evmAddr, types.JSONUtilsPrecompileAddress, inputBz, nil, nil)
			if tc.expectedErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			unpackedRet, err := abi.Methods["unmarshal_to_object"].Outputs.Unpack(retBz)
			require.NoError(t, err)

			res := *abiapi.ConvertType(unpackedRet[0], new(i_jsonutils.IJSONUtilsJSONObject)).(*i_jsonutils.IJSONUtilsJSONObject)
			require.Equal(t, tc.expected, res)
		})
	}
}

func Test_ConnectOracle_GetPrice(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()
	evmAddr := common.BytesToAddress(addr.Bytes())

	input.MarketMapKeeper.MarketMap["BTC/USD"] = marketmaptypes.Market{
		Ticker:          marketmaptypes.NewTicker("BTC", "USD", 6, 0, true),
		ProviderConfigs: []marketmaptypes.ProviderConfig{},
	}
	input.MarketMapKeeper.MarketMap["ETH/USD"] = marketmaptypes.Market{
		Ticker:          marketmaptypes.NewTicker("ETH", "USD", 6, 0, true),
		ProviderConfigs: []marketmaptypes.ProviderConfig{},
	}

	abi, err := connect_oracle.ConnectOracleMetaData.GetAbi()
	require.NoError(t, err)

	oracleAddr, err := input.EVMKeeper.GetConnectOracleAddr(ctx)
	require.NoError(t, err)

	// 1. get price in error case
	inputBz, err := abi.Pack("get_price", `BTC/USD`)
	require.NoError(t, err)

	ret, _, err := input.EVMKeeper.EVMCall(ctx, evmAddr, oracleAddr, inputBz, nil, nil)
	require.ErrorContains(t, err, types.ErrPrecompileFailed.Error())
	require.ErrorContains(t, err, vm.ErrExecutionReverted.Error())
	require.ErrorContains(t, err, "no price / nonce reported for CurrencyPair")

	inputBz, err = abi.Pack("get_price", `Error`)
	require.NoError(t, err)

	ret, _, err = input.EVMKeeper.EVMCall(ctx, evmAddr, oracleAddr, inputBz, nil, nil)
	require.ErrorContains(t, err, types.ErrPrecompileFailed.Error())
	require.ErrorContains(t, err, vm.ErrExecutionReverted.Error())
	require.ErrorContains(t, err, "incorrectly formatted CurrencyPair")

	// 2. get price in correct case
	// set BTC/USD price to 1000
	btcCp, err := connecttypes.CurrencyPairFromString("BTC/USD")
	require.NoError(t, err)
	err = input.OracleKeeper.SetPriceForCurrencyPair(ctx, btcCp, oracletypes.QuotePrice{
		Price:          math.NewInt(1000_000_000),
		BlockTimestamp: ctx.BlockTime(),
		BlockHeight:    uint64(ctx.BlockHeight()),
	})
	require.NoError(t, err)
	// set ETH/USD price to 100
	ethCp, err := connecttypes.CurrencyPairFromString("ETH/USD")
	require.NoError(t, err)
	err = input.OracleKeeper.SetPriceForCurrencyPair(ctx, ethCp, oracletypes.QuotePrice{
		Price:          math.NewInt(100_000_000),
		BlockTimestamp: ctx.BlockTime(),
		BlockHeight:    uint64(ctx.BlockHeight()),
	})
	require.NoError(t, err)

	// compute id and nonce
	btcId, ok := input.OracleKeeper.GetIDForCurrencyPair(ctx, btcCp)
	require.True(t, ok)
	btcNonce, err := input.OracleKeeper.GetNonceForCurrencyPair(ctx, btcCp)
	require.NoError(t, err)
	ethId, ok := input.OracleKeeper.GetIDForCurrencyPair(ctx, ethCp)
	require.True(t, ok)
	ethNonce, err := input.OracleKeeper.GetNonceForCurrencyPair(ctx, ethCp)
	require.NoError(t, err)

	// try get price
	inputBz, err = abi.Pack("get_price", `BTC/USD`)
	require.NoError(t, err)

	ret, _, err = input.EVMKeeper.EVMCall(ctx, evmAddr, oracleAddr, inputBz, nil, nil)
	require.NoError(t, err)

	unpackedRet, err := abi.Methods["get_price"].Outputs.Unpack(ret)
	require.NoError(t, err)

	res := *abiapi.ConvertType(unpackedRet[0], new(connect_oracle.ConnectOraclePrice)).(*connect_oracle.ConnectOraclePrice)
	require.Equal(t, connect_oracle.ConnectOraclePrice{
		Price:     big.NewInt(1000_000_000),
		Timestamp: big.NewInt(ctx.BlockTime().UnixNano()),
		Height:    uint64(ctx.BlockHeight()),
		Nonce:     btcNonce,
		Decimal:   6,
		Id:        btcId,
	}, res)

	// try get prices
	inputBz, err = abi.Pack("get_prices", []string{`BTC/USD`, `ETH/USD`})
	require.NoError(t, err)

	ret, _, err = input.EVMKeeper.EVMCall(ctx, evmAddr, oracleAddr, inputBz, nil, nil)
	require.NoError(t, err)

	unpackedRet, err = abi.Methods["get_prices"].Outputs.Unpack(ret)
	require.NoError(t, err)

	resArr := *abiapi.ConvertType(unpackedRet[0], new([]connect_oracle.ConnectOraclePrice)).(*[]connect_oracle.ConnectOraclePrice)
	require.Equal(t, []connect_oracle.ConnectOraclePrice{
		{
			Price:     big.NewInt(1000_000_000),
			Timestamp: big.NewInt(ctx.BlockTime().UnixNano()),
			Height:    uint64(ctx.BlockHeight()),
			Nonce:     btcNonce,
			Decimal:   6,
			Id:        btcId,
		},
		{
			Price:     big.NewInt(100_000_000),
			Timestamp: big.NewInt(ctx.BlockTime().UnixNano()),
			Height:    uint64(ctx.BlockHeight()),
			Nonce:     ethNonce,
			Decimal:   6,
			Id:        ethId,
		},
	}, resArr)
}
