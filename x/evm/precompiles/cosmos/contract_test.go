package cosmosprecompile_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"

	"cosmossdk.io/core/address"
	"cosmossdk.io/log"
	"cosmossdk.io/store"
	storemetrics "cosmossdk.io/store/metrics"
	db "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authcodec "github.com/cosmos/cosmos-sdk/x/auth/codec"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"

	contracts "github.com/initia-labs/minievm/x/evm/contracts/i_cosmos"
	precompiles "github.com/initia-labs/minievm/x/evm/precompiles/cosmos"
)

func setup() (sdk.Context, codec.Codec, address.Codec) {
	kv := db.NewMemDB()
	cms := store.NewCommitMultiStore(kv, log.NewNopLogger(), storemetrics.NewNoOpMetrics())

	cdc := codec.NewProtoCodec(codectypes.NewInterfaceRegistry())
	ac := authcodec.NewBech32Codec("init")
	return sdk.NewContext(cms, cmtproto.Header{}, false, log.NewNopLogger()), cdc, ac
}

func Test_CosmosPrecompile_ToCosmosAddress(t *testing.T) {
	ctx, cdc, ac := setup()
	cosmosPrecompile, err := precompiles.NewCosmosPrecompile(cdc, ac)
	require.NoError(t, err)

	cosmosPrecompile = cosmosPrecompile.WithContext(ctx).(precompiles.CosmosPrecompile)

	evmAddr := common.HexToAddress("0x1")
	cosmosAddr, err := ac.BytesToString(evmAddr.Bytes())
	require.NoError(t, err)

	abi, err := contracts.ICosmosMetaData.GetAbi()
	require.NoError(t, err)

	// to cosmos address
	inputBz, err := abi.Pack(precompiles.METHOD_TO_COSMOS_ADDRESS, evmAddr)
	require.NoError(t, err)

	// out of gas panic
	require.Panics(t, func() {
		_, _, _ = cosmosPrecompile.ExtendedRun(vm.AccountRef(evmAddr), inputBz, precompiles.TO_COSMOS_ADDRESS_GAS-1, true)
	})

	retBz, _, err := cosmosPrecompile.ExtendedRun(vm.AccountRef(evmAddr), inputBz, precompiles.TO_COSMOS_ADDRESS_GAS, true)
	require.NoError(t, err)

	ret, err := abi.Unpack(precompiles.METHOD_TO_COSMOS_ADDRESS, retBz)
	require.NoError(t, err)
	require.Equal(t, cosmosAddr, ret[0].(string))
}

func Test_CosmosPrecompile_ToEVMAddress(t *testing.T) {
	ctx, cdc, ac := setup()
	cosmosPrecompile, err := precompiles.NewCosmosPrecompile(cdc, ac)
	require.NoError(t, err)

	cosmosPrecompile = cosmosPrecompile.WithContext(ctx).(precompiles.CosmosPrecompile)

	evmAddr := common.HexToAddress("0x1")
	cosmosAddr, err := ac.BytesToString(evmAddr.Bytes())
	require.NoError(t, err)

	abi, err := contracts.ICosmosMetaData.GetAbi()
	require.NoError(t, err)

	// to cosmos address
	inputBz, err := abi.Pack(precompiles.METHOD_TO_EVM_ADDRESS, cosmosAddr)
	require.NoError(t, err)

	// out of gas panic
	require.Panics(t, func() {
		_, _, _ = cosmosPrecompile.ExtendedRun(vm.AccountRef(evmAddr), inputBz, precompiles.TO_EVM_ADDRESS_GAS-1, true)
	})

	retBz, _, err := cosmosPrecompile.ExtendedRun(vm.AccountRef(evmAddr), inputBz, precompiles.TO_EVM_ADDRESS_GAS, true)
	require.NoError(t, err)

	ret, err := abi.Unpack(precompiles.METHOD_TO_EVM_ADDRESS, retBz)
	require.NoError(t, err)
	require.Equal(t, evmAddr, ret[0].(common.Address))
}
