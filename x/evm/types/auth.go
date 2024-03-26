package types

import (
	"fmt"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

var (
	_ sdk.AccountI             = (*ContractAccount)(nil)
	_ authtypes.GenesisAccount = (*ContractAccount)(nil)
)

// NewContractAccountWithAddress create new contract account with the given address.
func NewContractAccountWithAddress(addr sdk.AccAddress) *ContractAccount {
	return &ContractAccount{
		authtypes.NewBaseAccountWithAddress(addr),
	}
}

// SetPubKey - Implements AccountI
func (ma ContractAccount) SetPubKey(pubKey cryptotypes.PubKey) error {
	return fmt.Errorf("not supported for contract accounts")
}
