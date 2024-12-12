package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/initia-labs/minievm/x/evm/types"
)

// convertToEVMAddress converts a cosmos address to an EVM address
// check if the shorthand has been registered and if so, check the
// registered account's origin address is same with given address.
//
// Also we create shorthand account if the address is not registered yet and is a signer.
func (k Keeper) convertToEVMAddress(ctx context.Context, addr sdk.AccAddress, isSigner bool) (common.Address, error) {
	if len(addr) == common.AddressLength {
		return common.BytesToAddress(addr.Bytes()), nil
	}

	accountNumber := uint64(0)
	shorthandAddr := common.BytesToAddress(addr.Bytes())
	if found := k.accountKeeper.HasAccount(ctx, shorthandAddr.Bytes()); found {
		account := k.accountKeeper.GetAccount(ctx, shorthandAddr.Bytes())

		// if the account is empty account, convert it to shorthand account
		if !types.IsEmptyAccount(account) {

			// check if the account is shorthand account, and if so, check if the original address is the same
			if shorthandAccount, isShorthandAccount := account.(types.ShorthandAccountI); isShorthandAccount {
				if originAddr, err := shorthandAccount.GetOriginalAddress(k.ac); err != nil {
					return common.Address{}, err
				} else if originAddr.Equals(addr) {
					return shorthandAddr, nil
				}
			}

			return common.Address{}, types.ErrAddressAlreadyExists.Wrapf("failed to create shorthand account of `%s`: `%s`", addr, shorthandAddr)
		}

		accountNumber = account.GetAccountNumber()
	}

	if isSigner {
		// if account number is not set, get next account number
		if accountNumber == 0 {
			accountNumber = k.accountKeeper.NextAccountNumber(ctx)
		}

		// create shorthand account
		shorthandAccount, err := types.NewShorthandAccountWithAddress(k.ac, addr)
		if err != nil {
			return common.Address{}, err
		}

		// register shorthand account
		shorthandAccount.AccountNumber = accountNumber
		k.accountKeeper.SetAccount(ctx, shorthandAccount)
	}

	return shorthandAddr, nil
}
