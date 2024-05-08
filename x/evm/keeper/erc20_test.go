package keeper_test

import (
	"testing"

	"cosmossdk.io/math"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/initia-labs/minievm/x/evm/contracts/erc20"
	"github.com/initia-labs/minievm/x/evm/contracts/erc20_factory"
	"github.com/initia-labs/minievm/x/evm/keeper"
	"github.com/initia-labs/minievm/x/evm/types"
)

func deployERC20(t *testing.T, ctx sdk.Context, input TestKeepers, caller common.Address, denom string) common.Address {
	abi, err := erc20_factory.Erc20FactoryMetaData.GetAbi()
	require.NoError(t, err)

	inputBz, err := abi.Pack("createERC20", denom, denom, uint8(6))
	require.NoError(t, err)

	ret, _, err := input.EVMKeeper.EVMCall(ctx, caller, types.ERC20FactoryAddress(), inputBz)
	require.NoError(t, err)

	return common.BytesToAddress(ret[12:])
}

func mintERC20(t *testing.T, ctx sdk.Context, input TestKeepers, caller, recipient common.Address, amount sdk.Coin) {
	abi, err := erc20.Erc20MetaData.GetAbi()
	require.NoError(t, err)

	inputBz, err := abi.Pack("mint", recipient, amount.Amount.BigInt())
	require.NoError(t, err)

	erc20ContractAddr, err := types.DenomToContractAddr(ctx, &input.EVMKeeper, amount.Denom)
	require.NoError(t, err)

	_, _, err = input.EVMKeeper.EVMCall(ctx, caller, erc20ContractAddr, inputBz)
	require.NoError(t, err)
}

func Test_MintBurn(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()
	evmAddr := common.BytesToAddress(addr.Bytes())

	erc20Keeper, err := keeper.NewERC20Keeper(&input.EVMKeeper)
	require.NoError(t, err)

	// deploy erc20 contract
	fooContractAddr := deployERC20(t, ctx, input, evmAddr, "foo")
	fooDenom, err := types.ContractAddrToDenom(ctx, &input.EVMKeeper, fooContractAddr)
	require.NoError(t, err)
	require.Equal(t, "evm/"+fooContractAddr.Hex()[2:], fooDenom)

	// cannot mint erc20 from cosmos side
	cacheCtx, _ := ctx.CacheContext()
	err = erc20Keeper.MintCoins(cacheCtx, addr, sdk.NewCoins(
		sdk.NewCoin("bar", math.NewInt(200)),
		sdk.NewCoin(fooDenom, math.NewInt(100)),
	))
	require.Error(t, err)

	// mint success
	err = erc20Keeper.MintCoins(ctx, addr, sdk.NewCoins(
		sdk.NewCoin("bar", math.NewInt(200)),
	))
	require.NoError(t, err)

	// mint erc20
	mintERC20(t, ctx, input, evmAddr, evmAddr, sdk.NewCoin(fooDenom, math.NewInt(100)))

	amount, err := erc20Keeper.GetBalance(ctx, addr, "bar")
	require.NoError(t, err)
	require.Equal(t, math.NewInt(200), amount)

	amount, err = erc20Keeper.GetBalance(ctx, addr, fooDenom)
	require.NoError(t, err)
	require.Equal(t, math.NewInt(100), amount)

	// erc20(foo) coins will be sent to community pool
	err = erc20Keeper.BurnCoins(ctx, addr, sdk.NewCoins(
		sdk.NewCoin("bar", math.NewInt(50)),
		sdk.NewCoin(fooDenom, math.NewInt(50)),
	))
	require.NoError(t, err)

	res, _, err := erc20Keeper.GetPaginatedBalances(ctx, nil, addr)
	require.NoError(t, err)
	require.Equal(t, sdk.NewCoins(
		sdk.NewCoin("bar", math.NewInt(150)),
		sdk.NewCoin(fooDenom, math.NewInt(100)),
	), res)

	// check community pool
	require.Equal(t, math.NewInt(50), input.CommunityPoolKeeper.CommunityPool.AmountOf(fooDenom))
}

func Test_SendCoins(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()
	_, _, addr2 := keyPubAddr()

	erc20Keeper, err := keeper.NewERC20Keeper(&input.EVMKeeper)
	require.NoError(t, err)

	err = erc20Keeper.MintCoins(ctx, addr, sdk.NewCoins(
		sdk.NewCoin("bar", math.NewInt(200)),
		sdk.NewCoin("foo", math.NewInt(100)),
	))
	require.NoError(t, err)

	err = erc20Keeper.SendCoins(ctx, addr, addr2, sdk.NewCoins(
		sdk.NewCoin("bar", math.NewInt(100)),
		sdk.NewCoin("foo", math.NewInt(50)),
	))
	require.NoError(t, err)

	res, _, err := erc20Keeper.GetPaginatedBalances(ctx, nil, addr)
	require.NoError(t, err)
	require.Equal(t, sdk.NewCoins(
		sdk.NewCoin("bar", math.NewInt(100)),
		sdk.NewCoin("foo", math.NewInt(50)),
	), res)

	res2, _, err := erc20Keeper.GetPaginatedBalances(ctx, nil, addr2)
	require.NoError(t, err)
	require.Equal(t, sdk.NewCoins(
		sdk.NewCoin("bar", math.NewInt(100)),
		sdk.NewCoin("foo", math.NewInt(50)),
	), res2)
}

func Test_GetSupply(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()
	_, _, addr2 := keyPubAddr()
	evmAddr := common.BytesToAddress(addr.Bytes())

	erc20Keeper, err := keeper.NewERC20Keeper(&input.EVMKeeper)
	require.NoError(t, err)

	// deploy erc20 contract
	fooContractAddr := deployERC20(t, ctx, input, evmAddr, "foo")
	fooDenom, err := types.ContractAddrToDenom(ctx, &input.EVMKeeper, fooContractAddr)
	require.NoError(t, err)
	require.Equal(t, "evm/"+fooContractAddr.Hex()[2:], fooDenom)

	// mint erc20
	mintERC20(t, ctx, input, evmAddr, evmAddr, sdk.NewCoin(fooDenom, math.NewInt(100)))

	// mint native coin
	err = erc20Keeper.MintCoins(ctx, addr, sdk.NewCoins(
		sdk.NewCoin("bar", math.NewInt(200)),
	))
	require.NoError(t, err)

	err = erc20Keeper.SendCoins(ctx, addr, addr2, sdk.NewCoins(
		sdk.NewCoin("bar", math.NewInt(100)),
		sdk.NewCoin(fooDenom, math.NewInt(50)),
	))
	require.NoError(t, err)

	amount, err := erc20Keeper.GetSupply(ctx, fooDenom)
	require.NoError(t, err)
	require.Equal(t, math.NewInt(100), amount)

	has, err := erc20Keeper.HasSupply(ctx, fooDenom)
	require.NoError(t, err)
	require.True(t, has)

	amount, err = erc20Keeper.GetSupply(ctx, "bar")
	require.NoError(t, err)
	require.Equal(t, math.NewInt(200), amount)

	has, err = erc20Keeper.HasSupply(ctx, "bar")
	require.NoError(t, err)
	require.True(t, has)

	erc20Keeper.IterateSupply(ctx, func(supply sdk.Coin) (bool, error) {
		require.True(t, supply.Denom == "bar" || supply.Denom == fooDenom)
		switch supply.Denom {
		case "bar":
			require.Equal(t, math.NewInt(200), supply.Amount)
		case fooDenom:
			require.Equal(t, math.NewInt(100), supply.Amount)
		}
		return false, nil
	})

	supply, _, err := erc20Keeper.GetPaginatedSupply(ctx, nil)
	require.NoError(t, err)
	require.Equal(t, sdk.NewCoins(
		sdk.NewCoin("bar", math.NewInt(200)),
		sdk.NewCoin(fooDenom, math.NewInt(100)),
	), supply)
}

func TestERC20Keeper_GetMetadata(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()

	erc20Keeper, err := keeper.NewERC20Keeper(&input.EVMKeeper)
	require.NoError(t, err)

	err = erc20Keeper.MintCoins(ctx, addr, sdk.NewCoins(
		sdk.NewCoin("bar", math.NewInt(200)),
		sdk.NewCoin("foo", math.NewInt(100)),
	))
	require.NoError(t, err)

	supply, err := erc20Keeper.GetSupply(ctx, "foo")
	require.NoError(t, err)
	require.Equal(t, math.NewInt(100), supply)

	metadata, err := erc20Keeper.GetMetadata(ctx, "foo")
	require.NoError(t, err)

	require.Equal(t, banktypes.Metadata{
		Description: "",
		Base:        "foo",
		Display:     "foo",
		Name:        "foo",
		Symbol:      "foo",
		URI:         "",
		URIHash:     "",
		DenomUnits: []*banktypes.DenomUnit{
			{
				Denom:    "foo",
				Exponent: 0,
			},
		},
	}, metadata)
}

func Test_IterateAccountBalances(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()
	_, _, addr2 := keyPubAddr()

	evmAddr := common.BytesToAddress(addr.Bytes())
	evmAddr2 := common.BytesToAddress(addr2.Bytes())

	erc20Keeper, err := keeper.NewERC20Keeper(&input.EVMKeeper)
	require.NoError(t, err)

	// deploy erc20 contract
	fooContractAddr := deployERC20(t, ctx, input, evmAddr, "foo")
	fooDenom, err := types.ContractAddrToDenom(ctx, &input.EVMKeeper, fooContractAddr)
	require.NoError(t, err)
	require.Equal(t, "evm/"+fooContractAddr.Hex()[2:], fooDenom)

	// mint erc20
	mintERC20(t, ctx, input, evmAddr, evmAddr, sdk.NewCoin(fooDenom, math.NewInt(100)))
	mintERC20(t, ctx, input, evmAddr, evmAddr2, sdk.NewCoin(fooDenom, math.NewInt(200)))

	// mint native coin
	err = erc20Keeper.MintCoins(ctx, addr, sdk.NewCoins(
		sdk.NewCoin("bar", math.NewInt(200)),
	))
	require.NoError(t, err)
	err = erc20Keeper.MintCoins(ctx, addr2, sdk.NewCoins(
		sdk.NewCoin("bar", math.NewInt(400)),
	))
	require.NoError(t, err)

	erc20Keeper.IterateAccountBalances(ctx, addr, func(balance sdk.Coin) (bool, error) {
		require.True(t, balance.Denom == "bar" || balance.Denom == fooDenom)
		switch balance.Denom {
		case "bar":
			require.Equal(t, math.NewInt(200), balance.Amount)
		case fooDenom:
			require.Equal(t, math.NewInt(100), balance.Amount)
		}
		return false, nil
	})

	erc20Keeper.IterateAccountBalances(ctx, addr2, func(balance sdk.Coin) (bool, error) {
		require.True(t, balance.Denom == "bar" || balance.Denom == fooDenom)
		switch balance.Denom {
		case "bar":
			require.Equal(t, math.NewInt(400), balance.Amount)
		case fooDenom:
			require.Equal(t, math.NewInt(200), balance.Amount)
		}
		return false, nil
	})
}
