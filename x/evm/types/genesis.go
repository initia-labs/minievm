package types

import (
	bytes "bytes"
	"errors"

	"cosmossdk.io/core/address"
	coretypes "github.com/ethereum/go-ethereum/core/types"
)

func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:    DefaultParams(),
		StateRoot: coretypes.EmptyRootHash[:],
		KeyValues: []GenesisKeyValue{},
	}
}

// Validate performs basic validation of move genesis data returning an
// error for any failed validation criteria.
func (genState *GenesisState) Validate(ac address.Codec) error {
	if len(genState.StateRoot) != 32 {
		return errors.New("invalid StateRoot hash length")
	}

	return genState.Params.Validate(ac)
}

func (genState GenesisState) IsExported() bool {
	return !bytes.Equal(genState.StateRoot, coretypes.EmptyRootHash[:])
}
