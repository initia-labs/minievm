package keeper_test

import (
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/require"

	"cosmossdk.io/math"
	transfertypes "github.com/cosmos/ibc-go/v10/modules/apps/transfer/types"
	opchildtypes "github.com/initia-labs/OPinit/x/opchild/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/holiman/uint256"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/initia-labs/minievm/x/evm/contracts/erc20"
	"github.com/initia-labs/minievm/x/evm/contracts/erc20_factory"
	"github.com/initia-labs/minievm/x/evm/contracts/erc20_wrapper"
	"github.com/initia-labs/minievm/x/evm/contracts/infinite_loop_erc20"
	"github.com/initia-labs/minievm/x/evm/keeper"
	"github.com/initia-labs/minievm/x/evm/types"
)

func deployERC20(t *testing.T, ctx sdk.Context, input TestKeepers, caller common.Address, symbol string) common.Address {
	abi, err := erc20_factory.Erc20FactoryMetaData.GetAbi()
	require.NoError(t, err)

	inputBz, err := abi.Pack("createERC20", symbol, symbol, uint8(6))
	require.NoError(t, err)

	factoryAddr, err := input.EVMKeeper.GetERC20FactoryAddr(ctx)
	require.NoError(t, err)

	ret, _, err := input.EVMKeeper.EVMCall(ctx, caller, factoryAddr, inputBz, nil, nil, nil)
	require.NoError(t, err)

	return common.BytesToAddress(ret[12:])
}

func deployERC20WithSalt(t *testing.T, ctx sdk.Context, input TestKeepers, caller common.Address, symbol string) common.Address {
	salt := func() [32]byte {
		var salt [32]byte
		rand.Read(salt[:])
		return salt
	}()
	factoryAddr, err := input.EVMKeeper.GetERC20FactoryAddr(ctx)
	require.NoError(t, err)

	abi, err := erc20_factory.Erc20FactoryMetaData.GetAbi()
	require.NoError(t, err)

	// compute the address of the contract
	inputBz, err := abi.Pack("computeERC20Address", caller, symbol, symbol, uint8(6), salt)
	require.NoError(t, err)
	expected, err := input.EVMKeeper.EVMStaticCall(ctx, caller, factoryAddr, inputBz, nil)
	require.NoError(t, err)

	inputBz, err = abi.Pack("createERC20WithSalt", symbol, symbol, uint8(6), salt)
	require.NoError(t, err)

	ret, _, err := input.EVMKeeper.EVMCall(ctx, caller, factoryAddr, inputBz, nil, nil, nil)
	require.NoError(t, err)

	require.Equal(t, expected[12:], ret[12:])

	return common.BytesToAddress(ret[12:])
}

func burnERC20(t *testing.T, ctx sdk.Context, input TestKeepers, caller, from common.Address, amount sdk.Coin, expectErr bool) {
	erc20ContractAddr, err := types.DenomToContractAddr(ctx, &input.EVMKeeper, amount.Denom)
	require.NoError(t, err)

	abi, err := erc20.Erc20MetaData.GetAbi()
	require.NoError(t, err)

	inputBz, err := abi.Pack("approve", caller, amount.Amount.BigInt())
	require.NoError(t, err)

	_, _, err = input.EVMKeeper.EVMCall(ctx, from, erc20ContractAddr, inputBz, nil, nil, nil)
	require.NoError(t, err)

	inputBz, err = abi.Pack("burnFrom", from, amount.Amount.BigInt())
	require.NoError(t, err)

	_, _, err = input.EVMKeeper.EVMCall(ctx, caller, erc20ContractAddr, inputBz, nil, nil, nil)
	if expectErr {
		require.Error(t, err)
	} else {
		require.NoError(t, err)
	}
}

func mintERC20(t *testing.T, ctx sdk.Context, input TestKeepers, caller, recipient common.Address, amount sdk.Coin, expectErr bool) {
	abi, err := erc20.Erc20MetaData.GetAbi()
	require.NoError(t, err)

	inputBz, err := abi.Pack("mint", recipient, amount.Amount.BigInt())
	require.NoError(t, err)

	erc20ContractAddr, err := types.DenomToContractAddr(ctx, &input.EVMKeeper, amount.Denom)
	require.NoError(t, err)

	_, _, err = input.EVMKeeper.EVMCall(ctx, caller, erc20ContractAddr, inputBz, nil, nil, nil)
	if expectErr {
		require.Error(t, err)
	} else {
		require.NoError(t, err)
	}
}

func transferERC20(t *testing.T, ctx sdk.Context, input TestKeepers, caller, recipient common.Address, amount sdk.Coin, expectErr bool) {
	abi, err := erc20.Erc20MetaData.GetAbi()
	require.NoError(t, err)

	inputBz, err := abi.Pack("transfer", recipient, amount.Amount.BigInt())
	require.NoError(t, err)

	erc20ContractAddr, err := types.DenomToContractAddr(ctx, &input.EVMKeeper, amount.Denom)
	require.NoError(t, err)

	_, _, err = input.EVMKeeper.EVMCall(ctx, caller, erc20ContractAddr, inputBz, nil, nil, nil)
	if expectErr {
		require.Error(t, err)
	} else {
		require.NoError(t, err)
	}

}

func approveERC20(t *testing.T, ctx sdk.Context, input TestKeepers, caller, spender common.Address, amount sdk.Coin, expectErr bool) {
	abi, err := erc20.Erc20MetaData.GetAbi()
	require.NoError(t, err)

	inputBz, err := abi.Pack("approve", spender, amount.Amount.BigInt())
	require.NoError(t, err)

	erc20ContractAddr, err := types.DenomToContractAddr(ctx, &input.EVMKeeper, amount.Denom)
	require.NoError(t, err)

	_, _, err = input.EVMKeeper.EVMCall(ctx, caller, erc20ContractAddr, inputBz, nil, nil, nil)
	if expectErr {
		require.Error(t, err)
	} else {
		require.NoError(t, err)
	}
}

func transferFromERC20(t *testing.T, ctx sdk.Context, input TestKeepers, caller, from, to common.Address, amount sdk.Coin, expectErr bool) {
	abi, err := erc20.Erc20MetaData.GetAbi()
	require.NoError(t, err)

	inputBz, err := abi.Pack("transferFrom", from, to, amount.Amount.BigInt())
	require.NoError(t, err)

	erc20ContractAddr, err := types.DenomToContractAddr(ctx, &input.EVMKeeper, amount.Denom)
	require.NoError(t, err)

	_, _, err = input.EVMKeeper.EVMCall(ctx, caller, erc20ContractAddr, inputBz, nil, nil, nil)
	if expectErr {
		require.Error(t, err)
	} else {
		require.NoError(t, err)
	}
}

func callERC20WrapperWithoutDispatch(
	t *testing.T,
	ctx sdk.Context,
	input TestKeepers,
	caller common.Address,
	inputBz []byte,
) []types.ExecuteRequest {
	t.Helper()

	wrapperAddr, err := input.EVMKeeper.GetERC20WrapperAddr(ctx)
	require.NoError(t, err)

	evmCtx, evm, cleanup, err := input.EVMKeeper.CreateEVM(ctx, caller)
	require.NoError(t, err)
	defer cleanup()

	_, _, err = evm.Call(caller, wrapperAddr, inputBz, 50_000_000, uint256.NewInt(0))
	require.NoError(t, err)

	stateDB := evm.StateDB.(types.StateDB)
	require.NoError(t, stateDB.Commit())

	requests := sdk.UnwrapSDKContext(evmCtx).Value(types.CONTEXT_KEY_EXECUTE_REQUESTS).(*[]types.ExecuteRequest)
	return *requests
}

func updateMetadataERC20(t *testing.T, ctx sdk.Context, input TestKeepers, caller common.Address, denom, name, symbol string, decimals uint8) error {
	abi, err := erc20.Erc20MetaData.GetAbi()
	require.NoError(t, err)

	inputBz, err := abi.Pack("updateMetadata", name, symbol, decimals)
	require.NoError(t, err)

	erc20ContractAddr, err := types.DenomToContractAddr(ctx, &input.EVMKeeper, denom)
	require.NoError(t, err)

	_, _, err = input.EVMKeeper.EVMCall(ctx, caller, erc20ContractAddr, inputBz, nil, nil, nil)
	return err
}

func Test_BalanceOf(t *testing.T) {
	ctx, input := createDefaultTestInput(t)

	_, _, addr := keyPubAddr()
	_, _, addr2 := keyPubAddr()

	input.Faucet.Fund(ctx, addr, sdk.NewCoin("foo", math.NewInt(100)))

	amount, err := input.EVMKeeper.ERC20Keeper().GetBalance(ctx, addr, "foo")
	require.NoError(t, err)
	require.Equal(t, math.NewInt(100), amount)

	amount, err = input.EVMKeeper.ERC20Keeper().GetBalance(ctx, addr2, "foo")
	require.NoError(t, err)
	require.Equal(t, math.NewInt(0), amount)
}

func Test_TransferToModuleAccount(t *testing.T) {
	ctx, input := createDefaultTestInput(t)

	_, _, addr := keyPubAddr()
	evmAddr := common.BytesToAddress(addr.Bytes())

	input.Faucet.Fund(ctx, addr, sdk.NewCoin("foo", math.NewInt(100)))

	feeCollectorAddr := authtypes.NewModuleAddress(authtypes.FeeCollectorName)
	transferERC20(t, ctx, input, evmAddr, common.BytesToAddress(feeCollectorAddr.Bytes()), sdk.NewCoin("foo", math.NewInt(50)), true)

	_, _, addr2 := keyPubAddr()
	evmAddr2 := common.BytesToAddress(addr2.Bytes())
	transferERC20(t, ctx, input, evmAddr, evmAddr2, sdk.NewCoin("foo", math.NewInt(50)), false)
}

func Test_MintToModuleAccount(t *testing.T) {
	ctx, input := createDefaultTestInput(t)

	_, _, addr := keyPubAddr()
	evmAddr := common.BytesToAddress(addr.Bytes())

	// deploy erc20 contract
	fooContractAddr := deployERC20(t, ctx, input, evmAddr, "foo")
	fooDenom, err := types.ContractAddrToDenom(ctx, &input.EVMKeeper, fooContractAddr)
	require.NoError(t, err)
	require.Equal(t, "evm/"+fooContractAddr.Hex()[2:], fooDenom)

	feeCollectorAddr := authtypes.NewModuleAddress(authtypes.FeeCollectorName)
	mintERC20(t, ctx, input, evmAddr, common.BytesToAddress(feeCollectorAddr.Bytes()), sdk.NewCoin(fooDenom, math.NewInt(50)), true)

	_, _, addr2 := keyPubAddr()
	evmAddr2 := common.BytesToAddress(addr2.Bytes())
	mintERC20(t, ctx, input, evmAddr, evmAddr2, sdk.NewCoin(fooDenom, math.NewInt(50)), false)

	// deploy erc20 contract with salt
	fooContractAddr = deployERC20WithSalt(t, ctx, input, evmAddr, "foo")
	fooDenom2, err := types.ContractAddrToDenom(ctx, &input.EVMKeeper, fooContractAddr)
	require.NoError(t, err)
	require.Equal(t, "evm/"+fooContractAddr.Hex()[2:], fooDenom2)

	feeCollectorAddr = authtypes.NewModuleAddress(authtypes.FeeCollectorName)
	mintERC20(t, ctx, input, evmAddr, common.BytesToAddress(feeCollectorAddr.Bytes()), sdk.NewCoin(fooDenom2, math.NewInt(50)), true)

	_, _, addr3 := keyPubAddr()
	evmAddr2 = common.BytesToAddress(addr3.Bytes())
	mintERC20(t, ctx, input, evmAddr, evmAddr2, sdk.NewCoin(fooDenom2, math.NewInt(50)), false)
}

func Test_BurnFromModuleAccount(t *testing.T) {
	ctx, input := createDefaultTestInput(t)

	// register fee collector module account
	input.AccountKeeper.GetModuleAccount(ctx, authtypes.FeeCollectorName)

	_, _, addr := keyPubAddr()
	evmAddr := common.BytesToAddress(addr.Bytes())
	_, _, addr2 := keyPubAddr()
	evmAddr2 := common.BytesToAddress(addr2.Bytes())

	erc20Keeper, err := keeper.NewERC20Keeper(&input.EVMKeeper)
	require.NoError(t, err)

	// deploy erc20 contract
	fooContractAddr := deployERC20(t, ctx, input, evmAddr, "foo")
	fooDenom, err := types.ContractAddrToDenom(ctx, &input.EVMKeeper, fooContractAddr)
	require.NoError(t, err)
	require.Equal(t, "evm/"+fooContractAddr.Hex()[2:], fooDenom)

	// mint coins
	feeCollectorAddr := authtypes.NewModuleAddress(authtypes.FeeCollectorName)
	mintERC20(t, ctx, input, evmAddr, evmAddr, sdk.NewCoin(fooDenom, math.NewInt(100)), false)
	erc20Keeper.SendCoins(ctx, addr, feeCollectorAddr, sdk.NewCoins(sdk.NewCoin(fooDenom, math.NewInt(50))))
	erc20Keeper.SendCoins(ctx, addr, addr2, sdk.NewCoins(sdk.NewCoin(fooDenom, math.NewInt(50))))

	// should not be able to burn from module account
	burnERC20(t, ctx, input, evmAddr, common.BytesToAddress(feeCollectorAddr.Bytes()), sdk.NewCoin(fooDenom, math.NewInt(50)), true)

	// should be able to burn from other account
	burnERC20(t, ctx, input, evmAddr, evmAddr2, sdk.NewCoin(fooDenom, math.NewInt(50)), false)

	fooContractAddr2 := deployERC20WithSalt(t, ctx, input, evmAddr, "foo")
	fooDenom2, err := types.ContractAddrToDenom(ctx, &input.EVMKeeper, fooContractAddr2)
	require.NoError(t, err)
	require.Equal(t, "evm/"+fooContractAddr2.Hex()[2:], fooDenom2)

	// mint coins
	mintERC20(t, ctx, input, evmAddr, evmAddr, sdk.NewCoin(fooDenom2, math.NewInt(100)), false)
	erc20Keeper.SendCoins(ctx, addr, feeCollectorAddr, sdk.NewCoins(sdk.NewCoin(fooDenom2, math.NewInt(50))))
	erc20Keeper.SendCoins(ctx, addr, addr2, sdk.NewCoins(sdk.NewCoin(fooDenom2, math.NewInt(50))))

	// should not be able to burn from module account
	burnERC20(t, ctx, input, evmAddr, common.BytesToAddress(feeCollectorAddr.Bytes()), sdk.NewCoin(fooDenom2, math.NewInt(50)), true)

	// should be able to burn from other account
	burnERC20(t, ctx, input, evmAddr, evmAddr2, sdk.NewCoin(fooDenom2, math.NewInt(50)), false)
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
	mintERC20(t, ctx, input, evmAddr, evmAddr, sdk.NewCoin(fooDenom, math.NewInt(100)), false)

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
		sdk.NewCoin(fooDenom, math.NewInt(50)),
	), res)

	// check community pool
	require.Equal(t, math.NewInt(50), input.CommunityPoolKeeper.CommunityPool.AmountOf(fooDenom))

}

func Test_BurnMultipleCoins(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()
	evmAddr := common.BytesToAddress(addr.Bytes())

	erc20Keeper, err := keeper.NewERC20Keeper(&input.EVMKeeper)
	require.NoError(t, err)

	contractAddr0 := deployERC20(t, ctx, input, evmAddr, "foo")
	denom0, _ := types.ContractAddrToDenom(ctx, &input.EVMKeeper, contractAddr0)

	contractAddr1 := deployERC20(t, ctx, input, evmAddr, "bar")
	denom1, _ := types.ContractAddrToDenom(ctx, &input.EVMKeeper, contractAddr1)
	// cannot mint erc20 from cosmos side
	cacheCtx, _ := ctx.CacheContext()
	err = erc20Keeper.MintCoins(cacheCtx, addr, sdk.NewCoins(
		sdk.NewCoin(denom0, math.NewInt(100)),
		sdk.NewCoin(denom1, math.NewInt(100)),
	))
	require.Error(t, err)

	mintERC20(t, ctx, input, evmAddr, evmAddr, sdk.NewCoin(denom0, math.NewInt(100)), false)
	mintERC20(t, ctx, input, evmAddr, evmAddr, sdk.NewCoin(denom1, math.NewInt(100)), false)

	res, _, err := erc20Keeper.GetPaginatedBalances(ctx, nil, addr)
	require.NoError(t, err)
	require.Equal(t, sdk.NewCoins(
		sdk.NewCoin(denom0, math.NewInt(100)),
		sdk.NewCoin(denom1, math.NewInt(100)),
	), res)

	require.True(t, input.CommunityPoolKeeper.CommunityPool.IsZero())
	err = erc20Keeper.BurnCoins(ctx, addr, sdk.NewCoins(
		sdk.NewCoin(denom0, math.NewInt(50)),
		sdk.NewCoin(denom1, math.NewInt(50)),
	))
	require.NoError(t, err)

	require.Equal(t, math.NewInt(50), input.CommunityPoolKeeper.CommunityPool.AmountOf(denom0))
	require.Equal(t, math.NewInt(50), input.CommunityPoolKeeper.CommunityPool.AmountOf(denom1))

	res, _, err = erc20Keeper.GetPaginatedBalances(ctx, nil, addr)
	require.NoError(t, err)
	require.Equal(t, sdk.NewCoins(
		sdk.NewCoin(denom0, math.NewInt(50)),
		sdk.NewCoin(denom1, math.NewInt(50)),
	), res)
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
	mintERC20(t, ctx, input, evmAddr, evmAddr, sdk.NewCoin(fooDenom, math.NewInt(100)), false)

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
		require.True(t, supply.Denom == "bar" || supply.Denom == fooDenom || supply.Denom == sdk.DefaultBondDenom)
		switch supply.Denom {
		case "bar":
			require.Equal(t, math.NewInt(200), supply.Amount)
		case fooDenom:
			require.Equal(t, math.NewInt(100), supply.Amount)
		case sdk.DefaultBondDenom:
			require.Equal(t, math.NewInt(1_000_000), supply.Amount)
		}
		return false, nil
	})

	supply, _, err := erc20Keeper.GetPaginatedSupply(ctx, nil)
	require.NoError(t, err)
	require.Equal(t, sdk.NewCoins(
		sdk.NewCoin("bar", math.NewInt(200)),
		sdk.NewCoin(fooDenom, math.NewInt(100)),
		sdk.NewCoin(sdk.DefaultBondDenom, math.NewInt(1_000_000)),
	), supply)
}

func TestERC20Keeper_GetMetadata(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()
	evmAddr := common.BytesToAddress(addr.Bytes())

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

	factoryAbi, err := erc20_factory.Erc20FactoryMetaData.GetAbi()
	require.NoError(t, err)

	callBz, err := factoryAbi.Pack("createERC20", "hey", "hey", uint8(18))
	require.NoError(t, err)

	erc20WrapperAddr, err := input.EVMKeeper.ERC20FactoryAddr.Get(ctx)
	require.NoError(t, err)
	retBz, _, err := input.EVMKeeper.EVMCall(ctx, evmAddr, common.BytesToAddress(erc20WrapperAddr), callBz, nil, nil, nil)
	require.NoError(t, err)
	require.NotEmpty(t, retBz)

	ret, err := factoryAbi.Unpack("createERC20", retBz)
	require.NoError(t, err)

	contractAddr := ret[0].(common.Address)
	denom, err := types.ContractAddrToDenom(ctx, &input.EVMKeeper, contractAddr)
	require.NoError(t, err)

	metadata, err = erc20Keeper.GetMetadata(ctx, denom)
	require.NoError(t, err)
	require.Equal(t, "hey", metadata.Name)
	require.Equal(t, "hey", metadata.Symbol)
	require.Equal(t, uint32(18), metadata.DenomUnits[1].Exponent)
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
	mintERC20(t, ctx, input, evmAddr, evmAddr, sdk.NewCoin(fooDenom, math.NewInt(100)), false)
	mintERC20(t, ctx, input, evmAddr, evmAddr2, sdk.NewCoin(fooDenom, math.NewInt(200)), false)

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

func Test_Approve(t *testing.T) {
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
	mintERC20(t, ctx, input, evmAddr, evmAddr, sdk.NewCoin(fooDenom, math.NewInt(100)), false)

	// approve erc20
	approveERC20(t, ctx, input, evmAddr, evmAddr2, sdk.NewCoin(fooDenom, math.NewInt(50)), false)

	// transferFrom erc20
	transferFromERC20(t, ctx, input, evmAddr2, evmAddr, evmAddr2, sdk.NewCoin(fooDenom, math.NewInt(50)), false)

	amount, err := erc20Keeper.GetBalance(ctx, addr, fooDenom)
	require.NoError(t, err)
	require.Equal(t, math.NewInt(50), amount)
	amount, err = erc20Keeper.GetBalance(ctx, addr2, fooDenom)
	require.NoError(t, err)
	require.Equal(t, math.NewInt(50), amount)

	// should fail to transferFrom more than approved
	transferFromERC20(t, ctx, input, evmAddr2, evmAddr, evmAddr2, sdk.NewCoin(fooDenom, math.NewInt(50)), true)
}

func Test_ERC20Wrapper_OPWithdrawEscapesReceiverJSONInjection(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	opchildtypes.RegisterInterfaces(input.EncodingConfig.InterfaceRegistry)

	_, _, addr := keyPubAddr()
	evmAddr := common.BytesToAddress(addr.Bytes())

	wrapperAddr, err := input.EVMKeeper.GetERC20WrapperAddr(ctx)
	require.NoError(t, err)

	localToken := deployERC20(t, ctx, input, evmAddr, "local")
	localDenom, err := types.ContractAddrToDenom(ctx, &input.EVMKeeper, localToken)
	require.NoError(t, err)

	mintERC20(t, ctx, input, evmAddr, evmAddr, sdk.NewCoin(localDenom, math.NewInt(1)), false)
	approveERC20(t, ctx, input, evmAddr, wrapperAddr, sdk.NewCoin(localDenom, math.NewInt(1)), false)

	abi, err := erc20_wrapper.Erc20WrapperMetaData.GetAbi()
	require.NoError(t, err)

	receiver := `init1atk","amount":{"denom":"uusdc","amount":"999999999999"},"to":"init1atk`
	inputBz, err := abi.Pack("toRemoteAndOPWithdraw", receiver, localDenom, math.NewInt(1).BigInt(), uint64(250_000))
	require.NoError(t, err)

	requests := callERC20WrapperWithoutDispatch(t, ctx, input, evmAddr, inputBz)
	require.Len(t, requests, 1)

	msg, ok := requests[0].Msg.(*opchildtypes.MsgInitiateTokenWithdrawal)
	require.True(t, ok)
	require.Equal(t, receiver, msg.To)
	require.Equal(t, math.NewInt(1), msg.Amount.Amount)
	require.NotEqual(t, "uusdc", msg.Amount.Denom)
}

func Test_ERC20Wrapper_IBCTransferEscapesJSONInjection(t *testing.T) {
	testCases := []struct {
		name     string
		channel  string
		receiver string
	}{
		{
			name:     "channel",
			channel:  `channel-0","token":{"denom":"ibc/USDC","amount":"1000000"},"source_channel":"channel-0`,
			receiver: "init1receiver",
		},
		{
			name:     "receiver",
			channel:  "channel-0",
			receiver: `attacker","token":{"denom":"ibc/USDC","amount":"1000000"},"receiver":"attacker`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx, input := createDefaultTestInput(t)
			transfertypes.RegisterInterfaces(input.EncodingConfig.InterfaceRegistry)

			_, _, addr := keyPubAddr()
			evmAddr := common.BytesToAddress(addr.Bytes())

			wrapperAddr, err := input.EVMKeeper.GetERC20WrapperAddr(ctx)
			require.NoError(t, err)

			localToken := deployERC20(t, ctx, input, evmAddr, "local")
			localDenom, err := types.ContractAddrToDenom(ctx, &input.EVMKeeper, localToken)
			require.NoError(t, err)

			mintERC20(t, ctx, input, evmAddr, evmAddr, sdk.NewCoin(localDenom, math.NewInt(1)), false)
			approveERC20(t, ctx, input, evmAddr, wrapperAddr, sdk.NewCoin(localDenom, math.NewInt(1)), false)

			abi, err := erc20_wrapper.Erc20WrapperMetaData.GetAbi()
			require.NoError(t, err)

			inputBz, err := abi.Pack(
				"toRemoteAndIBCTransfer",
				localDenom,
				math.NewInt(1).BigInt(),
				tc.channel,
				tc.receiver,
				math.NewInt(1000).BigInt(),
				"{}",
				uint64(250_000),
			)
			require.NoError(t, err)

			requests := callERC20WrapperWithoutDispatch(t, ctx, input, evmAddr, inputBz)
			require.Len(t, requests, 1)

			msg, ok := requests[0].Msg.(*transfertypes.MsgTransfer)
			require.True(t, ok)
			require.Equal(t, tc.channel, msg.SourceChannel)
			require.Equal(t, tc.receiver, msg.Receiver)
			require.Equal(t, math.NewInt(1), msg.Token.Amount)
			require.NotEqual(t, "ibc/USDC", msg.Token.Denom)
		})
	}
}

func Test_ERC20MetadataUpdate(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()

	evmAddr := common.BytesToAddress(addr.Bytes())
	authorityAddr, err := input.AccountKeeper.AddressCodec().StringToBytes(input.EVMKeeper.GetAuthority())
	require.NoError(t, err)
	authorityEVMAddr := common.BytesToAddress(authorityAddr)

	erc20Keeper, err := keeper.NewERC20Keeper(&input.EVMKeeper)
	require.NoError(t, err)

	// deploy erc20 contract
	fooContractAddr := deployERC20(t, ctx, input, evmAddr, "foo")
	fooDenom, err := types.ContractAddrToDenom(ctx, &input.EVMKeeper, fooContractAddr)
	require.NoError(t, err)
	require.Equal(t, "evm/"+fooContractAddr.Hex()[2:], fooDenom)

	// update metadata should fail because deployer is not 0x1
	err = updateMetadataERC20(t, ctx, input, authorityEVMAddr, fooDenom, "new name", "new symbol", 18)
	require.Error(t, err)

	// create erc20 contract with deployer 0x1
	fooDenom = "foo"
	err = input.EVMKeeper.ERC20Keeper().CreateERC20(ctx, fooDenom, 6)
	require.NoError(t, err)

	// update metadata
	err = updateMetadataERC20(t, ctx, input, authorityEVMAddr, fooDenom, "new name", "new symbol", 18)
	require.NoError(t, err)
	metadata, err := erc20Keeper.GetMetadata(ctx, fooDenom)
	require.NoError(t, err)
	require.Equal(t, "new name", metadata.Name)
	require.Equal(t, "new symbol", metadata.Symbol)
	require.Equal(t, uint32(18), metadata.DenomUnits[1].Exponent)

	// update metadata again should fail
	err = updateMetadataERC20(t, ctx, input, authorityEVMAddr, fooDenom, "new name", "new symbol", 18)
	require.Error(t, err)
}

func Test_ERC20TransferGas(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()
	params, err := input.EVMKeeper.Params.Get(ctx)
	require.NoError(t, err)
	params.AllowCustomERC20 = true
	require.NoError(t, input.EVMKeeper.Params.Set(ctx, params))

	evmAddr := common.BytesToAddress(addr.Bytes())

	bz, err := hexutil.Decode(infinite_loop_erc20.InfiniteLoopErc20MetaData.Bin)
	require.NoError(t, err)

	abi, err := infinite_loop_erc20.InfiniteLoopErc20MetaData.GetAbi()
	require.NoError(t, err)

	inputBz, err := abi.Constructor.Inputs.Pack("test", "test", uint8(18))
	require.NoError(t, err)

	bz = append(bz, inputBz...)
	_, contractAddr, _, err := input.EVMKeeper.EVMCreate(ctx, evmAddr, bz, nil, nil)
	require.NoError(t, err)

	testDenom, err := types.ContractAddrToDenom(ctx, &input.EVMKeeper, contractAddr)
	require.NoError(t, err)
	require.Equal(t, "evm/"+contractAddr.Hex()[2:], testDenom)

	// mint token to address
	inputBz, err = abi.Pack("mint", evmAddr, math.NewInt(100).BigInt())
	require.NoError(t, err)
	_, _, err = input.EVMKeeper.EVMCall(ctx, evmAddr, contractAddr, inputBz, nil, nil, nil)
	require.NoError(t, err)

	// should fail to send coins due to out of gas
	err = input.EVMKeeper.ERC20Keeper().SendCoins(ctx, addr, addr, sdk.NewCoins(sdk.NewCoin(testDenom, math.NewInt(100))))
	require.ErrorContains(t, err, "out of gas")
}

func Test_ERC20StaticCallGas(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()

	params, err := input.EVMKeeper.Params.Get(ctx)
	require.NoError(t, err)
	params.AllowCustomERC20 = true
	require.NoError(t, input.EVMKeeper.Params.Set(ctx, params))

	evmAddr := common.BytesToAddress(addr.Bytes())
	bz, err := hexutil.Decode(infinite_loop_erc20.InfiniteLoopErc20MetaData.Bin)
	require.NoError(t, err)

	abi, err := infinite_loop_erc20.InfiniteLoopErc20MetaData.GetAbi()
	require.NoError(t, err)

	inputBz, err := abi.Constructor.Inputs.Pack("test", "test", uint8(18))
	require.NoError(t, err)

	bz = append(bz, inputBz...)
	_, contractAddr, _, err := input.EVMKeeper.EVMCreate(ctx, evmAddr, bz, nil, nil)
	require.NoError(t, err)

	testDenom, err := types.ContractAddrToDenom(ctx, &input.EVMKeeper, contractAddr)
	require.NoError(t, err)
	require.Equal(t, "evm/"+contractAddr.Hex()[2:], testDenom)

	// mint token to address
	inputBz, err = abi.Pack("mint", evmAddr, math.NewInt(100).BigInt())
	require.NoError(t, err)
	_, _, err = input.EVMKeeper.EVMCall(ctx, evmAddr, contractAddr, inputBz, nil, nil, nil)
	require.NoError(t, err)

	// return 0 balance due to out of gas
	balance, err := input.EVMKeeper.ERC20Keeper().GetBalance(ctx, addr, testDenom)
	require.NoError(t, err)
	require.Equal(t, math.NewInt(0), balance)

	// return 0 supply due to out of gas
	supply, err := input.EVMKeeper.ERC20Keeper().GetSupply(ctx, testDenom)
	require.NoError(t, err)
	require.Equal(t, math.NewInt(0), supply)

	// return 0 decimals due to out of gas
	decimals, err := input.EVMKeeper.ERC20Keeper().GetDecimals(ctx, testDenom)
	require.NoError(t, err)
	require.Equal(t, uint8(0), decimals)

	// return empty strings for metadata
	metadata, err := input.EVMKeeper.ERC20Keeper().GetMetadata(ctx, testDenom)
	require.NoError(t, err)
	require.Equal(t, banktypes.Metadata{
		Name:    "",
		Symbol:  "",
		Base:    testDenom,
		Display: testDenom,
		DenomUnits: []*banktypes.DenomUnit{
			{
				Denom:    testDenom,
				Exponent: 0,
			},
			{
				Denom:    "",
				Exponent: 0,
			},
		},
	}, metadata)

	// supply should not contain the token
	err = input.EVMKeeper.ERC20Keeper().IterateSupply(ctx, func(supply sdk.Coin) (bool, error) {
		require.True(t, supply.Denom != testDenom || supply.Amount.IsZero())
		return false, nil
	})
	require.NoError(t, err)

	// balance should not contain the token
	err = input.EVMKeeper.ERC20Keeper().IterateAccountBalances(ctx, addr, func(balance sdk.Coin) (bool, error) {
		require.True(t, balance.Denom != testDenom)
		return false, nil
	})
	require.NoError(t, err)

	// paginated supply should not contain the token
	totalSupply, _, err := input.EVMKeeper.ERC20Keeper().GetPaginatedSupply(ctx, nil)
	require.NoError(t, err)
	require.True(t, totalSupply.AmountOf(testDenom).IsZero())

	// paginated balances should not contain the token
	balances, _, err := input.EVMKeeper.ERC20Keeper().GetPaginatedBalances(ctx, nil, addr)
	require.NoError(t, err)
	require.True(t, balances.AmountOf(testDenom).IsZero())
}
