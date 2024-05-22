package keeper

import (
	"context"
	"math/big"
	"strings"

	"cosmossdk.io/collections"
	moderrors "cosmossdk.io/errors"
	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/initia-labs/minievm/x/evm/contracts/erc20"
	"github.com/initia-labs/minievm/x/evm/contracts/erc20_factory"
	"github.com/initia-labs/minievm/x/evm/types"
)

type ERC20Keeper struct {
	*Keeper
	ERC20Bin        []byte
	ERC20ABI        *abi.ABI
	ERC20FactoryABI *abi.ABI
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

	erc20Bin, err := hexutil.Decode(erc20.Erc20Bin)
	if err != nil {
		return ERC20Keeper{}, err
	}

	return &ERC20Keeper{k, erc20Bin, erc20ABI, factoryABI}, nil
}

// BurnCoins implements IERC20Keeper.
func (k ERC20Keeper) BurnCoins(ctx context.Context, addr sdk.AccAddress, amount sdk.Coins) error {
	evmAddr, err := k.convertToEVMAddress(ctx, addr, false)
	if err != nil {
		return err
	}

	for _, coin := range amount {
		// if a coin is not generated from 0x1, then send the coin to community pool
		// because we don't have burn capability.
		if types.IsERC20Denom(coin.Denom) {
			if err := k.communityPoolKeeper.FundCommunityPool(ctx, amount, evmAddr.Bytes()); err != nil {
				return err
			}

			continue
		}

		// burn coins
		contractAddr, err := types.DenomToContractAddr(ctx, k, coin.Denom)
		if err != nil {
			return err
		}

		inputBz, err := k.ERC20ABI.Pack("burn", evmAddr, coin.Amount.BigInt())
		if err != nil {
			return types.ErrFailedToPackABI.Wrap(err.Error())
		}

		// ignore the return values
		_, _, err = k.EVMCall(ctx, types.StdAddress, contractAddr, inputBz)
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
	if err != nil {
		return math.ZeroInt(), err
	}

	return k.balanceOf(
		ctx,
		evmAddr,
		contractAddr,
	)
}

// GetMetadata implements IERC20Keeper.
func (k ERC20Keeper) GetMetadata(ctx context.Context, denom string) (banktypes.Metadata, error) {
	contractAddr, err := types.DenomToContractAddr(ctx, k, denom)
	if err != nil {
		return banktypes.Metadata{}, err
	}

	name, err := k.name(ctx, contractAddr)
	if err != nil {
		return banktypes.Metadata{}, err
	}

	symbol, err := k.symbol(ctx, contractAddr)
	if err != nil {
		return banktypes.Metadata{}, err
	}

	decimals, err := k.decimals(ctx, contractAddr)
	if err != nil {
		return banktypes.Metadata{}, err
	}

	denomUnits := []*banktypes.DenomUnit{
		{
			Denom:    denom,
			Exponent: 0,
		},
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

// GetPaginatedBalances implements IERC20Keeper.
func (k ERC20Keeper) GetPaginatedBalances(ctx context.Context, pageReq *query.PageRequest, addr sdk.AccAddress) (sdk.Coins, *query.PageResponse, error) {
	evmAddr, err := k.convertToEVMAddress(ctx, addr, false)
	if err != nil {
		return nil, nil, err
	}

	coins, res, err := query.CollectionPaginate(ctx, k.ERC20Stores, pageReq, func(key collections.Pair[[]byte, []byte], _ collections.NoValue) (sdk.Coin, error) {
		balance, err := k.balanceOf(ctx, common.BytesToAddress(key.K1()), common.BytesToAddress(key.K2()))
		if err != nil {
			balance = math.ZeroInt()
		}

		denom, err := types.ContractAddrToDenom(ctx, k, common.BytesToAddress(key.K2()))
		if err != nil {
			balance = math.ZeroInt()
		}

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
		supply, err := k.totalSupply(ctx, common.BytesToAddress(contractAddr))
		if err != nil {
			supply = math.ZeroInt()
		}

		denom, err := types.ContractAddrToDenom(ctx, k, common.BytesToAddress(contractAddr))
		if err != nil {
			supply = math.ZeroInt()
		}

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

	return k.totalSupply(ctx, contractAddr)
}

// HasSupply implements IERC20Keeper.
func (k ERC20Keeper) HasSupply(ctx context.Context, denom string) (bool, error) {
	contractAddr, err := types.DenomToContractAddr(ctx, k, denom)
	if err != nil {
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
		balance, err := k.balanceOf(ctx, common.BytesToAddress(key.K1()), common.BytesToAddress(key.K2()))
		if err != nil {
			balance = math.ZeroInt()
		}

		denom, err := types.ContractAddrToDenom(ctx, k, common.BytesToAddress(key.K2()))
		if err != nil {
			balance = math.ZeroInt()
		}

		// not return zero balance
		if balance.IsZero() {
			return false, nil
		}

		return cb(sdk.NewCoin(denom, balance))
	})
}

// IterateSupply implements IERC20Keeper.
func (k ERC20Keeper) IterateSupply(ctx context.Context, cb func(supply sdk.Coin) (bool, error)) error {
	return k.ERC20s.Walk(ctx, nil, func(contractAddr []byte) (stop bool, err error) {
		supply, err := k.totalSupply(ctx, common.BytesToAddress(contractAddr))
		if err != nil {
			supply = math.ZeroInt()
		}

		denom, err := types.ContractAddrToDenom(ctx, k, common.BytesToAddress(contractAddr))
		if err != nil {
			supply = math.ZeroInt()
		}

		return cb(sdk.NewCoin(denom, supply))
	})
}

// MintCoins implements IERC20Keeper.
func (k ERC20Keeper) MintCoins(ctx context.Context, addr sdk.AccAddress, amount sdk.Coins) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
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
			contractAddr, err := k.nextContractAddress(ctx, types.ERC20FactoryAddress())
			if err != nil {
				return err
			}

			if err := k.ERC20DenomsByContractAddr.Set(ctx, contractAddr.Bytes(), denom); err != nil {
				return err
			}

			if err := k.ERC20ContractAddrsByDenom.Set(ctx, denom, contractAddr.Bytes()); err != nil {
				return err
			}

			inputBz, err := k.ERC20FactoryABI.Pack("createERC20", denom, denom, uint8(0))
			if err != nil {
				return types.ErrFailedToPackABI.Wrap(err.Error())
			}

			ret, _, err := k.EVMCall(ctx, types.StdAddress, types.ERC20FactoryAddress(), inputBz)
			if err != nil {
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
		}

		// mint coin
		contractAddr, err := types.DenomToContractAddr(ctx, k, denom)
		if err != nil {
			return err
		}
		inputBz, err := k.ERC20ABI.Pack("mint", evmAddr, coin.Amount.BigInt())
		if err != nil {
			return types.ErrFailedToPackABI.Wrap(err.Error())
		}

		// ignore the return values
		_, _, err = k.EVMCall(ctx, types.StdAddress, contractAddr, inputBz)
		if err != nil {
			return err
		}
	}

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

		inputBz, err := k.ERC20ABI.Pack("transfer", evmToAddr, coin.Amount.BigInt())
		if err != nil {
			return types.ErrFailedToPackABI.Wrap(err.Error())
		}

		// ignore the return values
		_, _, err = k.EVMCall(ctx, evmFromAddr, contractAddr, inputBz)
		if err != nil {
			return err
		}
	}

	return nil
}

func (k ERC20Keeper) balanceOf(ctx context.Context, addr, contractAddr common.Address) (math.Int, error) {
	inputBz, err := k.ERC20ABI.Pack("balanceOf", addr)
	if err != nil {
		return math.ZeroInt(), types.ErrFailedToPackABI.Wrap(err.Error())
	}

	retBz, err := k.EVMStaticCall(ctx, types.NullAddress, contractAddr, inputBz)
	if err != nil {
		return math.ZeroInt(), err
	}

	res, err := k.ERC20ABI.Unpack("balanceOf", retBz)
	if err != nil {
		return math.ZeroInt(), types.ErrFailedToUnpackABI.Wrap(err.Error())
	}

	balance, ok := res[0].(*big.Int)
	if !ok {
		return math.ZeroInt(), types.ErrFailedToDecodeOutput
	}

	return math.NewIntFromBigInt(balance), nil
}

func (k ERC20Keeper) totalSupply(ctx context.Context, contractAddr common.Address) (math.Int, error) {
	inputBz, err := k.ERC20ABI.Pack("totalSupply")
	if err != nil {
		return math.ZeroInt(), types.ErrFailedToPackABI.Wrap(err.Error())
	}

	retBz, err := k.EVMStaticCall(ctx, types.NullAddress, contractAddr, inputBz)
	if err != nil {
		return math.ZeroInt(), err
	}

	res, err := k.ERC20ABI.Unpack("totalSupply", retBz)
	if err != nil {
		return math.ZeroInt(), types.ErrFailedToUnpackABI.Wrap(err.Error())
	}

	totalSupply, ok := res[0].(*big.Int)
	if !ok {
		return math.ZeroInt(), types.ErrFailedToDecodeOutput
	}

	return math.NewIntFromBigInt(totalSupply), nil
}

func (k ERC20Keeper) name(ctx context.Context, contractAddr common.Address) (string, error) {
	inputBz, err := k.ERC20ABI.Pack("name")
	if err != nil {
		return "", types.ErrFailedToPackABI.Wrap(err.Error())
	}

	retBz, err := k.EVMStaticCall(ctx, types.NullAddress, contractAddr, inputBz)
	if err != nil {
		return "", err
	}

	res, err := k.ERC20ABI.Unpack("name", retBz)
	if err != nil {
		return "", types.ErrFailedToUnpackABI.Wrap(err.Error())
	}

	name, ok := res[0].(string)
	if !ok {
		return name, types.ErrFailedToDecodeOutput
	}

	return name, nil
}

func (k ERC20Keeper) symbol(ctx context.Context, contractAddr common.Address) (string, error) {
	inputBz, err := k.ERC20ABI.Pack("symbol")
	if err != nil {
		return "", types.ErrFailedToPackABI.Wrap(err.Error())
	}

	retBz, err := k.EVMStaticCall(ctx, types.NullAddress, contractAddr, inputBz)
	if err != nil {
		return "", err
	}

	res, err := k.ERC20ABI.Unpack("symbol", retBz)
	if err != nil {
		return "", types.ErrFailedToUnpackABI.Wrap(err.Error())
	}

	symbol, ok := res[0].(string)
	if !ok {
		return symbol, types.ErrFailedToDecodeOutput
	}

	return symbol, nil
}

func (k ERC20Keeper) decimals(ctx context.Context, contractAddr common.Address) (uint8, error) {
	inputBz, err := k.ERC20ABI.Pack("decimals")
	if err != nil {
		return 0, types.ErrFailedToPackABI.Wrap(err.Error())
	}

	retBz, err := k.EVMStaticCall(ctx, types.NullAddress, contractAddr, inputBz)
	if err != nil {
		return 0, err
	}

	res, err := k.ERC20ABI.Unpack("decimals", retBz)
	if err != nil {
		return 0, types.ErrFailedToUnpackABI.Wrap(err.Error())
	}

	decimals, ok := res[0].(uint8)
	if !ok {
		return decimals, types.ErrFailedToDecodeOutput
	}

	return decimals, nil
}
