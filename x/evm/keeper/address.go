package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/initia-labs/minievm/x/evm/types"
)

// convertToEVMAddress converts a cosmos address to an EVM address
// check if the shorthand has been registered and if so return error
// else register the shorthand address as an account.
func (k Keeper) convertToEVMAddress(ctx context.Context, addr sdk.AccAddress) (common.Address, error) {
	if len(addr) == common.AddressLength {
		return common.BytesToAddress(addr.Bytes()), nil
	}

	shorthandAddr := common.BytesToAddress(addr.Bytes())
	if found := k.accountKeeper.HasAccount(ctx, shorthandAddr.Bytes()); found {
		existingAccount := k.accountKeeper.GetAccount(ctx, shorthandAddr.Bytes())
		shorthandAccount, isShorthandAccount := existingAccount.(types.ShorthandAccountI)
		if !isShorthandAccount {
			return common.Address{}, types.ErrAddressAlreadyExists.Wrapf("failed to create shorthand account: %s", shorthandAddr)
		}

		if originAddr, err := shorthandAccount.GetOriginalAddress(k.ac); err != nil {
			return common.Address{}, err
		} else if !originAddr.Equals(addr) {
			return common.Address{}, types.ErrAddressAlreadyExists.Wrapf("failed to create shorthand account: %s", shorthandAddr)
		}

		return shorthandAddr, nil
	}

	// create shorthand account
	shorthandAccount, err := types.NewShorthandAccountWithAddress(k.ac, addr)
	if err != nil {
		return common.Address{}, err
	}

	// register shorthand account
	shorthandAccount.AccountNumber = k.accountKeeper.NextAccountNumber(ctx)
	k.accountKeeper.SetAccount(ctx, shorthandAccount)

	return shorthandAddr, nil
}
