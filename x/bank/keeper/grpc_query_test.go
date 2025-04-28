package keeper_test

import (
	"testing"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/ethereum/go-ethereum/common"

	"github.com/initia-labs/minievm/x/evm/contracts/erc20_factory"
	evmtypes "github.com/initia-labs/minievm/x/evm/types"

	"github.com/stretchr/testify/require"
)

func TestQueryBalance(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := testdata.KeyTestPubAddr()

	_, err := input.BankKeeper.Balance(ctx, &types.QueryBalanceRequest{})
	require.Error(t, err)

	_, err = input.BankKeeper.Balance(ctx, &types.QueryBalanceRequest{Address: addr.String()})
	require.Error(t, err)

	testDenom := testDenoms[0]
	req := types.NewQueryBalanceRequest(addr, testDenom)
	res, err := input.BankKeeper.Balance(ctx, req)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.True(t, res.Balance.IsZero())

	origCoins := sdk.NewCoins(sdk.NewCoin(testDenom, math.NewInt(50)), sdk.NewCoin(bondDenom, math.NewInt(30)))
	acc := input.AccountKeeper.NewAccountWithAddress(ctx, addr)

	input.AccountKeeper.SetAccount(ctx, acc)
	input.Faucet.Fund(ctx, acc.GetAddress(), origCoins...)

	res, err = input.BankKeeper.Balance(ctx, req)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.True(t, res.Balance.IsEqual(sdk.NewCoin(testDenom, math.NewInt(50))))
}

func TestQueryAllBalances(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := testdata.KeyTestPubAddr()
	_, err := input.BankKeeper.AllBalances(ctx, &types.QueryAllBalancesRequest{})
	require.Error(t, err)

	pageReq := &query.PageRequest{
		Key:        nil,
		Limit:      1,
		CountTotal: false,
	}
	req := types.NewQueryAllBalancesRequest(addr, pageReq, false)
	res, err := input.BankKeeper.AllBalances(ctx, req)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.True(t, res.Balances.IsZero())

	testDenom := testDenoms[0]
	rewardCoin := sdk.NewCoin(testDenom, math.NewInt(50))
	bondCoin := sdk.NewCoin(bondDenom, math.NewInt(30))

	origCoins := sdk.NewCoins(rewardCoin, bondCoin)
	acc := input.AccountKeeper.NewAccountWithAddress(ctx, addr)

	input.AccountKeeper.SetAccount(ctx, acc)
	input.Faucet.Fund(ctx, acc.GetAddress(), origCoins...)

	res, err = input.BankKeeper.AllBalances(ctx, req)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, res.Balances.Len(), 1)
	require.NotNil(t, res.Pagination.NextKey)

	t.Log("query second page with nextkey")
	pageReq = &query.PageRequest{
		Key:        res.Pagination.NextKey,
		Limit:      1,
		CountTotal: true,
	}
	req = types.NewQueryAllBalancesRequest(addr, pageReq, false)
	res, err = input.BankKeeper.AllBalances(ctx, req)
	require.NoError(t, err)
	require.Equal(t, res.Balances.Len(), 1)
	require.Nil(t, res.Pagination.NextKey)
}

func TestQueryTotalSupply(t *testing.T) {
	ctx, input := createDefaultTestInput(t)

	mintCoins := sdk.NewCoins(sdk.NewInt64Coin("test", 400000000))
	expectedTotalSupply := initialTotalSupply().Add(mintCoins...)
	require.NoError(t, input.BankKeeper.MintCoins(ctx, authtypes.Minter, mintCoins))

	res, err := input.BankKeeper.TotalSupply(ctx, &types.QueryTotalSupplyRequest{})
	require.NoError(t, err)
	require.NotNil(t, res)

	require.Equal(t, expectedTotalSupply, res.Supply)
}

func TestQueryTotalSupplyOf(t *testing.T) {
	ctx, input := createDefaultTestInput(t)

	test1Supply := sdk.NewInt64Coin("foo", 4000000)
	test2Supply := sdk.NewInt64Coin("bar", 700000000)
	mintCoins := sdk.NewCoins(test1Supply, test2Supply)
	require.
		NoError(t, input.BankKeeper.MintCoins(ctx, authtypes.Minter, mintCoins))

	_, err := input.BankKeeper.SupplyOf(ctx, &types.QuerySupplyOfRequest{})
	require.Error(t, err)

	res, err := input.BankKeeper.SupplyOf(ctx, &types.QuerySupplyOfRequest{Denom: test1Supply.Denom})
	require.NoError(t, err)
	require.NotNil(t, res)

	require.Equal(t, test1Supply, res.Amount)
}

func TestQueryParams(t *testing.T) {
	ctx, input := createDefaultTestInput(t)

	res, err := input.BankKeeper.Params(ctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, input.BankKeeper.GetParams(ctx), res.GetParams())
}

func TestQueryDenomMetadata(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()
	evmAddr := common.BytesToAddress(addr.Bytes())

	res, err := input.BankKeeper.DenomMetadata(ctx, &types.QueryDenomMetadataRequest{
		Denom: bondDenom,
	})
	require.NoError(t, err)
	metadata := res.Metadata
	require.Equal(t, sdk.DefaultBondDenom, metadata.Base)
	require.Equal(t, sdk.DefaultBondDenom, metadata.Display)
	require.Equal(t, sdk.DefaultBondDenom, metadata.Name)
	require.Equal(t, []*types.DenomUnit{
		{
			Denom:    bondDenom,
			Exponent: 18,
		},
	}, metadata.DenomUnits)

	// test createERC20
	factoryAbi, err := erc20_factory.Erc20FactoryMetaData.GetAbi()
	require.NoError(t, err)

	callBz, err := factoryAbi.Pack("createERC20", "hey", "hey", uint8(18))
	require.NoError(t, err)

	erc20WrapperAddr, err := input.EVMKeeper.ERC20FactoryAddr.Get(ctx)
	require.NoError(t, err)
	retBz, _, err := input.EVMKeeper.EVMCall(ctx, evmAddr, common.BytesToAddress(erc20WrapperAddr), callBz, nil, nil)
	require.NoError(t, err)
	require.NotEmpty(t, retBz)

	ret, err := factoryAbi.Unpack("createERC20", retBz)
	require.NoError(t, err)

	contractAddr := ret[0].(common.Address)
	denom, err := evmtypes.ContractAddrToDenom(ctx, &input.EVMKeeper, contractAddr)
	require.NoError(t, err)

	res, err = input.BankKeeper.DenomMetadata(ctx, &types.QueryDenomMetadataRequest{
		Denom: denom,
	})
	require.NoError(t, err)
	require.Equal(t, "hey", res.Metadata.Name)
	require.Equal(t, "hey", res.Metadata.Symbol)
	require.Equal(t, uint32(18), res.Metadata.DenomUnits[1].Exponent)

	res2, err := input.BankKeeper.DenomMetadataByQueryString(ctx, &types.QueryDenomMetadataByQueryStringRequest{
		Denom: denom,
	})
	require.NoError(t, err)
	require.Equal(t, "hey", res2.Metadata.Name)
	require.Equal(t, "hey", res2.Metadata.Symbol)
	require.Equal(t, uint32(18), res2.Metadata.DenomUnits[1].Exponent)
}
