package keeper

import (
	"context"
	"slices"

	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common"

	"github.com/initia-labs/minievm/x/evm/types"
)

// ERC20StoresKeeper defines the erc20 related store keeper.
// It keeps track of the erc20 contract addresses registered to user's store
// and the erc20 contract addresses registered to the store.
type ERC20StoresKeeper struct {
	*Keeper
}

// NewERC20StoresKeeper creates a new instance of the ERC20StoresKeeper.
func NewERC20StoresKeeper(k *Keeper) types.IERC20StoresKeeper {
	return &ERC20StoresKeeper{k}
}

// IsStoreRegistered checks if the erc20 contract address is registered to user's store.
func (k ERC20StoresKeeper) IsStoreRegistered(ctx context.Context, addr sdk.AccAddress, contractAddr common.Address) (bool, error) {
	return k.ERC20Stores.Has(ctx, collections.Join(addr.Bytes(), contractAddr.Bytes()))
}

// RegisterStore registers the erc20 contract address to user's store.
func (k ERC20StoresKeeper) RegisterStore(ctx context.Context, addr sdk.AccAddress, contractAddr common.Address) error {
	if found, err := k.ERC20s.Has(ctx, contractAddr.Bytes()); err != nil {
		return err
	} else if !found {
		return nil
	}

	// create account if not exists
	if !k.accountKeeper.HasAccount(ctx, addr) {
		k.accountKeeper.SetAccount(ctx, k.accountKeeper.NewAccountWithAddress(ctx, addr))
	}

	return k.ERC20Stores.Set(ctx, collections.Join(addr.Bytes(), contractAddr.Bytes()))
}

// Register registers the erc20 contract address to the store.
func (k ERC20StoresKeeper) Register(ctx context.Context, contractAddr common.Address) error {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return err
	}

	if !params.AllowCustomERC20 {
		return types.ErrCustomERC20NotAllowed
	}

	if found, err := k.ERC20s.Has(ctx, contractAddr.Bytes()); err != nil {
		return err
	} else if found {
		return nil
	}

	// default action is to allow all custom erc20s
	// but if allowedCustomERC20s is set, only allow those
	if len(params.AllowedCustomERC20s) > 0 {
		if idx := slices.IndexFunc(params.AllowedCustomERC20s, func(s string) bool {
			return contractAddr.String() == s
		}); idx == -1 {
			return types.ErrCustomERC20NotAllowed
		}
	}

	return k.ERC20s.Set(ctx, contractAddr.Bytes())
}

// RegisterFromFactory registers the erc20 contract address to the store.
func (k ERC20StoresKeeper) RegisterFromFactory(ctx context.Context, caller, contractAddr common.Address) error {
	if factoryAddr, err := k.GetERC20FactoryAddr(ctx); err != nil {
		return err
	} else if factoryAddr != caller {
		return types.ErrInvalidERC20FactoryAddr
	}

	return k.ERC20s.Set(ctx, contractAddr.Bytes())
}
