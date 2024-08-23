package types

import (
	"errors"

	"cosmossdk.io/core/address"
)

func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:       DefaultParams(),
		Erc20Factory: nil,
		KeyValues:    []GenesisKeyValue{},
	}
}

// Validate performs basic validation of move genesis data returning an
// error for any failed validation criteria.
func (genState *GenesisState) Validate(ac address.Codec) error {
	if genState.Erc20Factory == nil {
		return errors.New("invalid empty ERC20Factory address")
	}

	return genState.Params.Validate(ac)
}
