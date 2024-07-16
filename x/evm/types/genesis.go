package types

import (
	"errors"

	"cosmossdk.io/core/address"
	"github.com/ethereum/go-ethereum/common"
	coretypes "github.com/ethereum/go-ethereum/core/types"
)

func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:       DefaultParams(),
		StateRoot:    coretypes.EmptyRootHash[:],
		Erc20Factory: nil,
		KeyValues:    []GenesisKeyValue{},
	}
}

// Validate performs basic validation of move genesis data returning an
// error for any failed validation criteria.
func (genState *GenesisState) Validate(ac address.Codec) error {
	if len(genState.StateRoot) != 32 {
		return errors.New("invalid StateRoot hash length")
	}

	if common.BytesToHash(genState.StateRoot) != coretypes.EmptyRootHash && genState.Erc20Factory == nil {
		return errors.New("invalid empty ERC20Factory address")
	}

	return genState.Params.Validate(ac)
}
