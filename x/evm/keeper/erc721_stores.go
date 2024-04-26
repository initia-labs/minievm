package keeper

import (
	"context"

	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/initia-labs/minievm/x/evm/types"
)

// ERC721StoresKeeper defines the erc721 related store keeper.
// It keeps track of the erc721 contract addresses registered to user's store
// and the erc721 contract addresses registered to the store.
type ERC721StoresKeeper struct {
	*Keeper
}

// NewERC721StoresKeeper creates a new instance of the ERC721StoresKeeper.
func NewERC721StoresKeeper(k *Keeper) types.IERC721StoresKeeper {
	return &ERC721StoresKeeper{k}
}

// IsStoreRegistered checks if the erc721 contract address is registered to user's store.
func (k ERC721StoresKeeper) IsStoreRegistered(ctx context.Context, addr sdk.AccAddress, contractAddr common.Address) (bool, error) {
	return k.ERC721Stores.Has(ctx, collections.Join(addr.Bytes(), contractAddr.Bytes()))
}

// RegisterStore registers the erc721 contract address to user's store.
func (k ERC721StoresKeeper) RegisterStore(ctx context.Context, addr sdk.AccAddress, contractAddr common.Address) error {
	// create account if not exists
	if !k.accountKeeper.HasAccount(ctx, addr) {
		k.accountKeeper.SetAccount(ctx, k.accountKeeper.NewAccountWithAddress(ctx, addr))
	}

	return k.ERC721Stores.Set(ctx, collections.Join(addr.Bytes(), contractAddr.Bytes()))
}

// Register registers the erc721 contract address to the store.
func (k ERC721StoresKeeper) Register(ctx context.Context, contractAddr common.Address) error {
	if found, err := k.ERC721s.Has(ctx, contractAddr.Bytes()); err != nil {
		return err
	} else if found {
		return nil
	}

	return k.ERC721s.Set(ctx, contractAddr.Bytes())
}
