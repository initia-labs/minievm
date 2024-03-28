package types

import (
	"fmt"

	"cosmossdk.io/core/address"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/ethereum/go-ethereum/common"
)

var (
	_ sdk.AccountI = (*ShorthandAccount)(nil)
	_ sdk.AccountI = (*ShorthandAccount)(nil)

	_ authtypes.GenesisAccount = (*ContractAccount)(nil)
	_ authtypes.GenesisAccount = (*ShorthandAccount)(nil)

	_ ShorthandAccountI = (*ShorthandAccount)(nil)
)

// NewContractAccountWithAddress create new contract account with the given address.
func NewContractAccountWithAddress(addr sdk.AccAddress) *ContractAccount {
	return &ContractAccount{
		authtypes.NewBaseAccountWithAddress(addr),
	}
}

// SetPubKey - Implements AccountI
func (ca ContractAccount) SetPubKey(pubKey cryptotypes.PubKey) error {
	return fmt.Errorf("not supported for contract accounts")
}

type ShorthandAccountI interface {
	GetOriginalAddress(ac address.Codec) (sdk.AccAddress, error)
}

// NewShorthandAccountWithAddress create new contract account with the given address.
func NewShorthandAccountWithAddress(ac address.Codec, addr sdk.AccAddress) (*ShorthandAccount, error) {
	originAddr, err := ac.BytesToString(addr.Bytes())
	if err != nil {
		return nil, err
	}

	shorthandAddr := common.BytesToAddress(addr.Bytes())
	return &ShorthandAccount{
		BaseAccount:     authtypes.NewBaseAccountWithAddress(shorthandAddr.Bytes()),
		OriginalAddress: originAddr,
	}, nil
}

// SetPubKey - Implements AccountI
func (sa ShorthandAccount) SetPubKey(pubKey cryptotypes.PubKey) error {
	return fmt.Errorf("not supported for shorthand accounts")
}

func (sa ShorthandAccount) GetOriginalAddress(ac address.Codec) (sdk.AccAddress, error) {
	return ac.StringToBytes(sa.OriginalAddress)
}
