package keeper

import (
	"context"
	"errors"
	"math/big"
	"strings"

	"cosmossdk.io/collections"
	moderrors "cosmossdk.io/errors"
	"cosmossdk.io/math"
	storetypes "cosmossdk.io/store/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/initia-labs/minievm/x/evm/contracts/erc20"
	"github.com/initia-labs/minievm/x/evm/contracts/erc20_factory"
	"github.com/initia-labs/minievm/x/evm/contracts/erc20_wrapper"
	"github.com/initia-labs/minievm/x/evm/types"
)

var _ types.IERC20Keeper = &ERC20Keeper{}

type ERC20Keeper struct {
	*Keeper
	ERC20Bin        []byte
	ERC20ABI        *abi.ABI
	ERC20FactoryABI *abi.ABI
	ERC20WrapperABI *abi.ABI
}

func NewERC20Keeper(k *Keeper) (types.IERC20Keeper, error) {
	erc20ABI, err := erc20.Erc20MetaData.GetAbi()
	if err != nil {
		return ERC20Keeper{}, err
	}

	factoryABI, err := erc20_factory.Erc20FactoryMetaData.GetAbi()
	if err != nil {
		return ERC20Keeper{}, err
	}

	wrapperABI, err := erc20_wrapper.Erc20WrapperMetaData.GetAbi()
	if err != nil {
		return ERC20Keeper{}, err
	}

	erc20Bin, err := hexutil.Decode(erc20.Erc20Bin)
	if err != nil {
		return ERC20Keeper{}, err
	}

	return &ERC20Keeper{k, erc20Bin, erc20ABI, factoryABI, wrapperABI}, nil
}

// GetERC20ABI implements IERC20Keeper.
func (k ERC20Keeper) GetERC20ABI() *abi.ABI {
	return k.ERC20ABI
}

// GetERC20FactoryABI implements IERC20Keeper.
func (K ERC20Keeper) GetERC20FactoryABI() *abi.ABI {
	return K.ERC20FactoryABI
}

// GetERC20WrapperABI implements IERC20Keeper.
func (K ERC20Keeper) GetERC20WrapperABI() *abi.ABI {
	return K.ERC20WrapperABI
}

// BurnCoins implements IERC20Keeper.
func (k ERC20Keeper) BurnCoins(ctx context.Context, addr sdk.AccAddress, amount sdk.Coins) error {
	evmAddr, err := k.convertToEVMAddress(ctx, addr, false)
	if err != nil {
		return err
	}

	communityPoolFunds := sdk.NewCoins()
	for _, coin := range amount {
		// if a coin is not generated from 0x1, then send the coin to community pool
		// because we don't have burn capability.
		if types.IsERC20Denom(coin.Denom) {
			communityPoolFunds = communityPoolFunds.Add(coin)
			continue
		}

		// burn coins
		contractAddr, err := types.DenomToContractAddr(ctx, k, coin.Denom)
		if err != nil {
			return err
		}

		inputBz, err := k.ERC20ABI.Pack("sudoBurn", evmAddr, coin.Amount.BigInt())
		if err != nil {
			return types.ErrFailedToPackABI.Wrap(err.Error())
		}

		// ignore the return values
		_, _, err = k.EVMCall(ctx, types.StdAddress, contractAddr, inputBz, nil, nil)
		if err != nil {
			return err
		}
	}

	if !communityPoolFunds.IsZero() {
		if err := k.communityPoolKeeper.FundCommunityPool(ctx, communityPoolFunds, evmAddr.Bytes()); err != nil {
			return err
		}
	}

	return nil
}

// MintCoins implements IERC20Keeper.
func (k ERC20Keeper) MintCoins(ctx context.Context, addr sdk.AccAddress, amount sdk.Coins) error {
	evmAddr, err := k.convertToEVMAddress(ctx, addr, false)
	if err != nil {
		return err
	}

	for _, coin := range amount {
		denom := coin.Denom
		if types.IsERC20Denom(denom) {
			return moderrors.Wrapf(types.ErrInvalidRequest, "cannot mint erc20 coin: %s", coin.Denom)
		}

		// check whether the erc20 contract exists or not and create it if not
		if found, err := k.ERC20ContractAddrsByDenom.Has(ctx, denom); err != nil {
			return err
		} else if !found {
			err := k.CreateERC20(ctx, denom, 0)
			if err != nil {
				return err
			}
		}

		// mint coin
		contractAddr, err := types.DenomToContractAddr(ctx, k, denom)
		if err != nil {
			return err
		}
		inputBz, err := k.ERC20ABI.Pack("sudoMint", evmAddr, coin.Amount.BigInt())
		if err != nil {
			return types.ErrFailedToPackABI.Wrap(err.Error())
		}

		// ignore the return values
		_, _, err = k.EVMCall(ctx, types.StdAddress, contractAddr, inputBz, nil, nil)
		if err != nil {
			return err
		}
	}

	return nil
}

// TokenCreationFn is a helper function to create a new ERC20 token if it doesn't exist.
func (k ERC20Keeper) TokenCreationFn(ctx context.Context, denom string, decimals uint8) error {
	found, err := k.ERC20ContractAddrsByDenom.Has(ctx, denom)
	if err != nil {
		return err
	} else if found {
		return nil
	}

	return k.CreateERC20(ctx, denom, decimals)
}

func (k ERC20Keeper) CreateERC20(ctx context.Context, denom string, decimals uint8) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	factoryAddr, err := k.GetERC20FactoryAddr(ctx)
	if err != nil {
		return err
	}

	inputBz, err := k.ERC20FactoryABI.Pack("createERC20", denom, denom, uint8(decimals))
	if err != nil {
		return types.ErrFailedToPackABI.Wrap(err.Error())
	}

	ret, _, err := k.EVMCall(ctx, types.StdAddress, factoryAddr, inputBz, nil, nil)
	if err != nil {
		return err
	}

	res, err := k.ERC20FactoryABI.Unpack("createERC20", ret)
	if err != nil {
		return types.ErrFailedToUnpackABI.Wrap(err.Error())
	}

	// store created erc20 contract address <> denom mapping
	contractAddr := res[0].(common.Address)
	if err := k.ERC20DenomsByContractAddr.Set(ctx, contractAddr.Bytes(), denom); err != nil {
		return err
	}
	if err := k.ERC20ContractAddrsByDenom.Set(ctx, denom, contractAddr.Bytes()); err != nil {
		return err
	}

	// emit erc20 created event
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeERC20Created,
			sdk.NewAttribute(types.AttributeKeyDenom, denom),
			sdk.NewAttribute(types.AttributeKeyContract, hexutil.Encode(ret[12:])),
		),
	)

	return nil
}

// SendCoins implements IERC20Keeper.
func (k ERC20Keeper) SendCoins(ctx context.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error {
	evmFromAddr, err := k.convertToEVMAddress(ctx, fromAddr, true)
	if err != nil {
		return err
	}
	evmToAddr, err := k.convertToEVMAddress(ctx, toAddr, false)
	if err != nil {
		return err
	}

	for _, coin := range amt {
		contractAddr, err := types.DenomToContractAddr(ctx, k, coin.Denom)
		if err != nil {
			return err
		}

		inputBz, err := k.ERC20ABI.Pack("sudoTransfer", evmFromAddr, evmToAddr, coin.Amount.BigInt())
		if err != nil {
			return types.ErrFailedToPackABI.Wrap(err.Error())
		}

		// ignore the return values
		_, _, err = k.EVMCall(ctx, types.StdAddress, contractAddr, inputBz, nil, nil)
		if err != nil {
			return err
		}
	}

	return nil
}

// GetBalance implements IERC20Keeper.
func (k ERC20Keeper) GetBalance(ctx context.Context, addr sdk.AccAddress, denom string) (math.Int, error) {
	evmAddr, err := k.convertToEVMAddress(ctx, addr, false)
	if err != nil {
		return math.ZeroInt(), err
	}

	contractAddr, err := types.DenomToContractAddr(ctx, k, denom)
	if err != nil && errors.Is(err, collections.ErrNotFound) {
		return math.ZeroInt(), nil
	} else if err != nil {
		return math.ZeroInt(), err
	}

	return k.balanceOf(ctx, evmAddr, contractAddr), nil
}

// GetMetadata implements IERC20Keeper.
func (k ERC20Keeper) GetMetadata(ctx context.Context, denom string) (banktypes.Metadata, error) {
	contractAddr, err := types.DenomToContractAddr(ctx, k, denom)
	if err != nil {
		return banktypes.Metadata{}, err
	}

	name := k.name(ctx, contractAddr)
	symbol := k.symbol(ctx, contractAddr)
	decimals := k.decimals(ctx, contractAddr)
	denomUnits := []*banktypes.DenomUnit{
		{
			Denom:    denom,
			Exponent: 0,
		},
		{
			Denom:    symbol,
			Exponent: uint32(decimals),
		},
	}
	if denom == symbol {
		denomUnits = denomUnits[1:]
	}

	base := denom
	display := denom
	if decimals == 0 {
		if !strings.Contains(denom, "/") && denom[0] == 'u' {
			display = denom[1:]
			denomUnits = append(denomUnits, &banktypes.DenomUnit{
				Denom:    display,
				Exponent: 6,
			})
		} else if !strings.Contains(denom, "/") && denom[0] == 'm' {
			display = denom[1:]
			denomUnits = append(denomUnits, &banktypes.DenomUnit{
				Denom:    display,
				Exponent: 3,
			})
		}
	}

	return banktypes.Metadata{
		Name:       name,
		Symbol:     symbol,
		Base:       base,
		Display:    display,
		DenomUnits: denomUnits,
	}, nil
}

// HasMetadata implements IERC20Keeper.
func (k ERC20Keeper) HasMetadata(ctx context.Context, denom string) (bool, error) {
	contractAddr, err := types.DenomToContractAddr(ctx, k, denom)
	if err != nil && errors.Is(err, collections.ErrNotFound) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return k.ERC20s.Has(ctx, contractAddr.Bytes())
}

// GetPaginatedBalances implements IERC20Keeper.
func (k ERC20Keeper) GetPaginatedBalances(ctx context.Context, pageReq *query.PageRequest, addr sdk.AccAddress) (sdk.Coins, *query.PageResponse, error) {
	evmAddr, err := k.convertToEVMAddress(ctx, addr, false)
	if err != nil {
		return nil, nil, err
	}

	coins, res, err := query.CollectionPaginate(ctx, k.ERC20Stores, pageReq, func(key collections.Pair[[]byte, []byte], _ collections.NoValue) (sdk.Coin, error) {
		denom, err := types.ContractAddrToDenom(ctx, k, common.BytesToAddress(key.K2()))
		if err != nil {
			return sdk.Coin{}, err
		}

		balance := k.balanceOf(ctx, common.BytesToAddress(key.K1()), common.BytesToAddress(key.K2()))
		return sdk.NewCoin(denom, balance), nil
	}, func(opt *query.CollectionsPaginateOptions[collections.Pair[[]byte, []byte]]) {
		prefix := collections.PairPrefix[[]byte, []byte](evmAddr.Bytes())
		opt.Prefix = &prefix
	})

	return sdk.Coins(coins).Sort(), res, err
}

// GetPaginatedSupply implements IERC20Keeper.
func (k ERC20Keeper) GetPaginatedSupply(ctx context.Context, pageReq *query.PageRequest) (sdk.Coins, *query.PageResponse, error) {
	coins, res, err := query.CollectionPaginate(ctx, k.ERC20s, pageReq, func(contractAddr []byte, _ collections.NoValue) (sdk.Coin, error) {
		denom, err := types.ContractAddrToDenom(ctx, k, common.BytesToAddress(contractAddr))
		if err != nil {
			return sdk.Coin{}, err
		}

		supply := k.totalSupply(ctx, common.BytesToAddress(contractAddr))
		return sdk.NewCoin(denom, supply), nil
	})

	return sdk.Coins(coins).Sort(), res, err
}

// GetSupply implements IERC20Keeper.
func (k ERC20Keeper) GetSupply(ctx context.Context, denom string) (math.Int, error) {
	contractAddr, err := types.DenomToContractAddr(ctx, k, denom)
	if err != nil {
		return math.ZeroInt(), err
	}

	return k.totalSupply(ctx, contractAddr), nil
}

// HasSupply implements IERC20Keeper.
func (k ERC20Keeper) HasSupply(ctx context.Context, denom string) (bool, error) {
	contractAddr, err := types.DenomToContractAddr(ctx, k, denom)
	if err != nil && errors.Is(err, collections.ErrNotFound) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return k.ERC20s.Has(ctx, contractAddr.Bytes())
}

// IterateAccountBalances implements IERC20Keeper.
func (k ERC20Keeper) IterateAccountBalances(ctx context.Context, addr sdk.AccAddress, cb func(sdk.Coin) (bool, error)) error {
	evmAddr, err := k.convertToEVMAddress(ctx, addr, false)
	if err != nil {
		return err
	}

	return k.ERC20Stores.Walk(ctx, collections.NewPrefixedPairRange[[]byte, []byte](evmAddr.Bytes()), func(key collections.Pair[[]byte, []byte]) (stop bool, err error) {
		denom, err := types.ContractAddrToDenom(ctx, k, common.BytesToAddress(key.K2()))
		if err != nil {
			return false, nil
		}

		// not return zero balance
		balance := k.balanceOf(ctx, common.BytesToAddress(key.K1()), common.BytesToAddress(key.K2()))
		if balance.IsZero() {
			return false, nil
		}

		return cb(sdk.NewCoin(denom, balance))
	})
}

// IterateSupply implements IERC20Keeper.
func (k ERC20Keeper) IterateSupply(ctx context.Context, cb func(supply sdk.Coin) (bool, error)) error {
	return k.ERC20s.Walk(ctx, nil, func(contractAddr []byte) (stop bool, err error) {
		denom, err := types.ContractAddrToDenom(ctx, k, common.BytesToAddress(contractAddr))
		if err != nil {
			return false, nil
		}

		supply := k.totalSupply(ctx, common.BytesToAddress(contractAddr))
		return cb(sdk.NewCoin(denom, supply))
	})
}

// GetDecimals returns the number of decimals for an ERC20 token with the given denom.
func (k ERC20Keeper) GetDecimals(ctx context.Context, denom string) (uint8, error) {
	contractAddr, err := types.DenomToContractAddr(ctx, k, denom)
	if err != nil {
		return 0, err
	}

	return k.decimals(ctx, contractAddr), nil
}

// Decimals returns the number of decimals for an ERC20 token with the given contract address.
func (k ERC20Keeper) Decimals(ctx context.Context, contractAddr common.Address) uint8 {
	return k.decimals(ctx, contractAddr)
}

// erc20StaticCallGas is the gas limit for static EVM calls.
const erc20StaticCallGas = 100000

// prepareStaticCall creates a context with limited gas for static EVM calls.
// It returns the context and a cleanup function that must be called after the call.
// The cleanup function ensures accurate gas accounting by consuming any used gas.
func prepareStaticCall(ctx context.Context, description string) (context.Context, func()) {
	gasMeter := sdk.UnwrapSDKContext(ctx).GasMeter()
	gasLimit := min(gasMeter.GasRemaining(), erc20StaticCallGas)
	ctx = sdk.UnwrapSDKContext(ctx).WithGasMeter(storetypes.NewGasMeter(gasLimit))

	return ctx, func() {
		gasConsumed := sdk.UnwrapSDKContext(ctx).GasMeter().GasConsumedToLimit()
		gasMeter.ConsumeGas(gasConsumed, description)
	}
}

// balanceOf is a helper function that returns the balance of an ERC20 token for a given address.
// It performs a static call to the token contract's balanceOf() function.
// If any error occurs during the static call (e.g. out of gas, contract reverts),
// or if the return value cannot be unpacked, it returns 0 as a safe default.
func (k ERC20Keeper) balanceOf(ctx context.Context, addr, contractAddr common.Address) (b math.Int) {
	ctx, cleanup := prepareStaticCall(ctx, "erc20 balanceOf")
	defer func() {
		// ignore the panic
		if r := recover(); r != nil {
			b = math.ZeroInt()
		}

		cleanup()
	}()

	inputBz, err := k.ERC20ABI.Pack("balanceOf", addr)
	if err != nil {
		return math.ZeroInt()
	}

	retBz, err := k.EVMStaticCall(ctx, types.NullAddress, contractAddr, inputBz, nil)
	if err != nil {
		return math.ZeroInt()
	}

	res, err := k.ERC20ABI.Unpack("balanceOf", retBz)
	if err != nil {
		return math.ZeroInt()
	}

	balance, ok := res[0].(*big.Int)
	if !ok {
		return math.ZeroInt()
	}

	return math.NewIntFromBigInt(balance)
}

// totalSupply is a helper function that returns the total supply of an ERC20 token.
// It performs a static call to the token contract's totalSupply() function.
// If any error occurs during the static call (e.g. out of gas, contract reverts),
// or if the return value cannot be unpacked, it returns 0 as a safe default.
func (k ERC20Keeper) totalSupply(ctx context.Context, contractAddr common.Address) (s math.Int) {
	ctx, cleanup := prepareStaticCall(ctx, "erc20 totalSupply")
	defer func() {
		// ignore the panic
		if r := recover(); r != nil {
			s = math.ZeroInt()
		}

		cleanup()
	}()

	inputBz, err := k.ERC20ABI.Pack("totalSupply")
	if err != nil {
		return math.ZeroInt()
	}

	retBz, err := k.EVMStaticCall(ctx, types.NullAddress, contractAddr, inputBz, nil)
	if err != nil {
		return math.ZeroInt()
	}

	res, err := k.ERC20ABI.Unpack("totalSupply", retBz)
	if err != nil {
		return math.ZeroInt()
	}

	totalSupply, ok := res[0].(*big.Int)
	if !ok {
		return math.ZeroInt()
	}

	return math.NewIntFromBigInt(totalSupply)
}

// name is a helper function that returns the name of an ERC20 token.
// It performs a static call to the token contract's name() function.
// If any error occurs during the static call (e.g. out of gas, contract reverts),
// or if the return value cannot be unpacked, it returns an empty string as a safe default.
func (k ERC20Keeper) name(ctx context.Context, contractAddr common.Address) (n string) {
	ctx, cleanup := prepareStaticCall(ctx, "erc20 name")
	defer func() {
		// ignore the panic
		if r := recover(); r != nil {
			n = ""
		}

		cleanup()
	}()

	inputBz, err := k.ERC20ABI.Pack("name")
	if err != nil {
		return ""
	}

	retBz, err := k.EVMStaticCall(ctx, types.NullAddress, contractAddr, inputBz, nil)
	if err != nil {
		return ""
	}

	res, err := k.ERC20ABI.Unpack("name", retBz)
	if err != nil {
		return ""
	}

	name, ok := res[0].(string)
	if !ok {
		return ""
	}

	return name
}

// symbol is a helper function that returns the symbol of an ERC20 token.
// It performs a static call to the token contract's symbol() function.
// If any error occurs during the static call (e.g. out of gas, contract reverts),
// or if the return value cannot be unpacked, it returns an empty string as a safe default.
func (k ERC20Keeper) symbol(ctx context.Context, contractAddr common.Address) (s string) {
	ctx, cleanup := prepareStaticCall(ctx, "erc20 symbol")
	defer func() {
		// ignore the panic
		if r := recover(); r != nil {
			s = ""
		}

		cleanup()
	}()

	inputBz, err := k.ERC20ABI.Pack("symbol")
	if err != nil {
		return ""
	}

	retBz, err := k.EVMStaticCall(ctx, types.NullAddress, contractAddr, inputBz, nil)
	if err != nil {
		return ""
	}

	res, err := k.ERC20ABI.Unpack("symbol", retBz)
	if err != nil {
		return ""
	}

	symbol, ok := res[0].(string)
	if !ok {
		return ""
	}

	return symbol
}

// decimals is a helper function that returns the number of decimals for an ERC20 token.
// It performs a static call to the token contract's decimals() function.
// If any error occurs during the static call (e.g. out of gas, contract reverts),
// or if the return value cannot be unpacked, it returns 0 as a safe default.
func (k ERC20Keeper) decimals(ctx context.Context, contractAddr common.Address) (d uint8) {
	ctx, cleanup := prepareStaticCall(ctx, "erc20 decimals")
	defer func() {
		// ignore the panic
		if r := recover(); r != nil {
			d = 0
		}

		cleanup()
	}()

	inputBz, err := k.ERC20ABI.Pack("decimals")
	if err != nil {
		return 0
	}

	retBz, err := k.EVMStaticCall(
		// set the context value to prevent infinite loop
		sdk.UnwrapSDKContext(ctx).WithValue(types.CONTEXT_KEY_LOAD_DECIMALS, true),
		types.NullAddress, contractAddr, inputBz, nil)
	if err != nil {
		return 0
	}

	res, err := k.ERC20ABI.Unpack("decimals", retBz)
	if err != nil {
		return 0
	}

	decimals, ok := res[0].(uint8)
	if !ok {
		return 0
	}

	return decimals
}
