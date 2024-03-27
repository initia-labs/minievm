package keeper

import (
	"context"

	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/initia-labs/minievm/x/evm/types"
)

type ERC20StoresKeeper struct {
	*Keeper
}

func NewERC20StoresKeeper(k *Keeper) types.IERC20StoresKeeper {
	return &ERC20StoresKeeper{k}
}

func (k ERC20StoresKeeper) IsStoreRegistered(ctx context.Context, addr sdk.AccAddress, contractAddr common.Address) (bool, error) {
	return k.ERC20Stores.Has(ctx, collections.Join(addr.Bytes(), contractAddr.Bytes()))
}

func (k ERC20StoresKeeper) RegisterStore(ctx context.Context, addr sdk.AccAddress, contractAddr common.Address) error {
	// create account if not exists
	if !k.accountKeeper.HasAccount(ctx, addr) {
		k.accountKeeper.SetAccount(ctx, k.accountKeeper.NewAccountWithAddress(ctx, addr))
	}

	return k.ERC20Stores.Set(ctx, collections.Join(addr.Bytes(), contractAddr.Bytes()))
}

func (k ERC20StoresKeeper) Register(ctx context.Context, contractAddr common.Address) error {
	if found, err := k.ERC20DenomsByContractAddr.Has(ctx, contractAddr.Bytes()); err != nil {
		return err
	} else if !found {
		// register denom and contract address conversion to the store
		denom, err := types.ContractAddrToDenom(ctx, k, contractAddr)
		if err != nil {
			return err
		}

		if err := k.ERC20DenomsByContractAddr.Set(ctx, contractAddr.Bytes(), denom); err != nil {
			return err
		}

		if err := k.ERC20ContractAddrsByDenom.Set(ctx, denom, contractAddr.Bytes()); err != nil {
			return err
		}
	}

	return k.ERC20s.Set(ctx, contractAddr.Bytes())
}
