package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/initia-labs/minievm/x/evm/contracts/erc20_wrapper"
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
	genState.Erc20Wrapper = []byte{5, 6, 7, 8}
	genState.ConnectOracle = []byte{9, 10, 11, 12}
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

func Test_Initialize(t *testing.T) {
	ctx, input := createTestInput(t, false, true)
	wrapperAddr, err := input.EVMKeeper.GetERC20WrapperAddr(ctx)
	require.NoError(t, err)

	caller := common.HexToAddress("0x0")
	abi, err := erc20_wrapper.Erc20WrapperMetaData.GetAbi()
	require.NoError(t, err)

	viewArg, err := abi.Pack("factory")
	require.NoError(t, err)

	factoryAddrBytes, err := input.EVMKeeper.EVMStaticCall(ctx, caller, wrapperAddr, viewArg, nil)
	require.NoError(t, err)

	factoryAddr := common.BytesToAddress(factoryAddrBytes)
	expectedFactoryAddr, err := input.EVMKeeper.GetERC20FactoryAddr(ctx)
	require.NoError(t, err)

	require.Equal(t, expectedFactoryAddr, factoryAddr)
}

func Test_DeployERC20Factory(t *testing.T) {
	ctx, input := createTestInput(t, false, false)

	// set random factory address and check if it is set correctly
	err := input.EVMKeeper.ERC20FactoryAddr.Set(ctx, common.HexToAddress("0x123").Bytes())
	require.NoError(t, err)

	// set params
	input.EVMKeeper.Params.Set(ctx, types.DefaultParams())

	factoryAddr, err := input.EVMKeeper.GetERC20FactoryAddr(ctx)
	require.NoError(t, err)

	require.Equal(t, common.HexToAddress("0x123"), factoryAddr)

	// deploy wrapper contract and check if the factory address is set correctly
	err = input.EVMKeeper.DeployERC20Wrapper(ctx)
	require.NoError(t, err)

	expectedFactoryAddr, err := input.EVMKeeper.GetERC20FactoryAddr(ctx)
	require.NoError(t, err)

	require.Equal(t, expectedFactoryAddr, queryFactoryAddressFromWrapper(t, ctx, input))

	// deploy factory contract again and check if the factory address is set correctly
	err = input.EVMKeeper.DeployERC20Factory(ctx)
	require.NoError(t, err)

	expectedFactoryAddr, err = input.EVMKeeper.GetERC20FactoryAddr(ctx)
	require.NoError(t, err)

	require.Equal(t, expectedFactoryAddr, queryFactoryAddressFromWrapper(t, ctx, input))
	require.NotEqual(t, factoryAddr, expectedFactoryAddr)
}

func queryFactoryAddressFromWrapper(t *testing.T, ctx sdk.Context, input TestKeepers) common.Address {
	wrapperAddr, err := input.EVMKeeper.GetERC20WrapperAddr(ctx)
	require.NoError(t, err)

	caller := common.HexToAddress("0x0")
	abi, err := erc20_wrapper.Erc20WrapperMetaData.GetAbi()
	require.NoError(t, err)

	viewArg, err := abi.Pack("factory")
	require.NoError(t, err)

	factoryAddrBytes, err := input.EVMKeeper.EVMStaticCall(ctx, caller, wrapperAddr, viewArg, nil)
	require.NoError(t, err)

	factoryAddr := common.BytesToAddress(factoryAddrBytes)
	return factoryAddr
}
