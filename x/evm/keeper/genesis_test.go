package keeper_test

import (
	"testing"

	"github.com/initia-labs/minievm/x/evm/types"
	"github.com/stretchr/testify/require"
)

func Test_Genesis(t *testing.T) {
	ctx, input := createTestInput(t, false, false)

	genState := types.DefaultGenesis()
	genState.Erc20Factory = []byte{5, 6, 7, 8}
	genState.ERC20s = [][]byte{
		{1, 2, 3},
		{4, 5, 6},
	}
	genState.Erc20Stores = []types.GenesisERC20Stores{
		{
			Address: []byte{1, 2, 3},
			Stores:  [][]byte{{1, 2, 3, 4, 5, 6, 7}, {8, 9, 1, 2, 3, 4}},
		},
		{
			Address: []byte{4, 2, 3},
			Stores:  [][]byte{{1, 2, 3, 4, 5, 6, 7}, {8, 9, 1, 2, 3, 4}},
		},
	}
	genState.KeyValues = []types.GenesisKeyValue{
		{
			Key:   []byte{1, 2, 3, 4},
			Value: []byte{1, 2, 3, 5, 32},
		},
		{
			Key:   []byte{5, 6, 7, 8},
			Value: []byte{1, 2, 3, 1, 32},
		},
		{
			Key:   []byte{9, 10, 11, 12, 13},
			Value: []byte{1, 2, 3, 5, 32, 4},
		},
	}
	genState.DenomTraces = []types.GenesisDenomTrace{
		{
			Denom:           "denom1",
			ContractAddress: []byte{1, 2, 3, 4, 5, 6, 7},
		},
		{
			Denom:           "denom2",
			ContractAddress: []byte{8, 9, 1, 2, 3, 4},
		},
	}
	genState.ClassTraces = []types.GenesisClassTrace{
		{
			ClassId:         "class1",
			ContractAddress: []byte{1, 2, 3, 4, 5, 6, 7},
			Uri:             "uri1",
		},
		{
			ClassId:         "class2",
			ContractAddress: []byte{8, 9, 1, 2, 3, 4},
			Uri:             "uri2",
		},
	}
	genState.EVMBlockHashes = []types.GenesisEVMBlockHash{
		{
			Height: 1,
			Hash:   []byte{1, 2, 3, 4, 5, 6, 7},
		},
		{
			Height: 2,
			Hash:   []byte{8, 9, 1, 2, 3, 4},
		},
	}
	err := input.EVMKeeper.InitGenesis(ctx, *genState)
	require.NoError(t, err)

	_genState := input.EVMKeeper.ExportGenesis(ctx)
	require.Equal(t, genState, _genState)
}
