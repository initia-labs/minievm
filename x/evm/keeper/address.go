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
		account := k.accountKeeper.GetAccount(ctx, shorthandAddr.Bytes())

		// if the account is empty account, convert it to shorthand account
		if types.IsEmptyAccount(account) {
			shorthandAccount, err := types.NewShorthandAccountWithAddress(k.ac, addr)
			if err != nil {
				return common.Address{}, err
			}

			shorthandAccount.AccountNumber = account.GetAccountNumber()
			k.accountKeeper.SetAccount(ctx, shorthandAccount)

			return shorthandAddr, nil
		}

		// check if the account is shorthand account, and if so, check if the original address is the same
		shorthandAccount, isShorthandAccount := account.(types.ShorthandAccountI)
		if isShorthandAccount {
			if originAddr, err := shorthandAccount.GetOriginalAddress(k.ac); err != nil {
				return common.Address{}, err
			} else if originAddr.Equals(addr) {
				return shorthandAddr, nil
			}
		}

		return common.Address{}, types.ErrAddressAlreadyExists.Wrapf("failed to create shorthand account of `%s`: `%s`", addr, shorthandAddr)
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
