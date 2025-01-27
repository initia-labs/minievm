package jsonutils_test

import (
	"testing"

	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/stretchr/testify/require"

	"cosmossdk.io/core/address"
	"cosmossdk.io/log"
	"cosmossdk.io/store"
	storemetrics "cosmossdk.io/store/metrics"
	"cosmossdk.io/x/tx/signing"
	db "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/codec"
	codecaddress "github.com/cosmos/cosmos-sdk/codec/address"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/gogoproto/proto"

	contracts "github.com/initia-labs/minievm/x/evm/contracts/i_jsonutils"
	precompiles "github.com/initia-labs/minievm/x/evm/precompiles/jsonutils"
	"github.com/initia-labs/minievm/x/evm/types"
)

func setup() (sdk.Context, codec.Codec, address.Codec) {
	kv := db.NewMemDB()
	cms := store.NewCommitMultiStore(kv, log.NewNopLogger(), storemetrics.NewNoOpMetrics())

	ctx := sdk.NewContext(cms, cmtproto.Header{}, false, log.NewNopLogger()).WithValue(types.CONTEXT_KEY_EXECUTE_REQUESTS, &[]types.ExecuteRequest{})

	interfaceRegistry, _ := codectypes.NewInterfaceRegistryWithOptions(codectypes.InterfaceRegistryOptions{
		ProtoFiles: proto.HybridResolver,
		SigningOptions: signing.Options{
			AddressCodec:          codecaddress.NewBech32Codec("init"),
			ValidatorAddressCodec: codecaddress.NewBech32Codec("initvaloper"),
		},
	})
	banktypes.RegisterInterfaces(interfaceRegistry)

	cdc := codec.NewProtoCodec(interfaceRegistry)
	ac := codecaddress.NewBech32Codec("init")

	return ctx, cdc, ac
}

func Test_JSONUtilsPrecompile_Merge(t *testing.T) {
	testCases := []struct {
		name        string
		dst         string
		src         string
		expected    string
		expectedErr bool
	}{
		{
			name:     "simple merge",
			dst:      `{"a": 1, "b": 2}`,
			src:      `{"b": 3, "c": 4}`,
			expected: `{"a":1,"b":3,"c":4}`,
		},
		{
			name:     "nested merge",
			dst:      `{"a": 1, "b": {"c": 2}}`,
			src:      `{"b": {"d": 3}, "c": 4}`,
			expected: `{"a":1,"b":{"c":2,"d":3},"c":4}`,
		},
		{
			name:     "nested merge with conflict",
			dst:      `{"a": 1, "b": {"c": 2}}`,
			src:      `{"b": {"c": 3}, "c": 4}`,
			expected: `{"a":1,"b":{"c":3},"c":4}`,
		},
		{
			name:        "invalid dst json",
			dst:         `{"a": 1, "b": 2`,
			src:         `{"b": 3, "c": 4}`,
			expectedErr: true,
		},
		{
			name:        "invalid src json",
			dst:         `{"a": 1, "b": 2}`,
			src:         `{"b": 3, "c": 4`,
			expectedErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx, _, _ := setup()
			stateDB := NewMockStateDB(ctx)
			contract, err := precompiles.NewJSONUtilsPrecompile(stateDB)
			require.NoError(t, err)

			abi, err := contracts.IJsonutilsMetaData.GetAbi()
			require.NoError(t, err)

			bz, err := abi.Pack(precompiles.METHOD_MERGE_JSON, tc.dst, tc.src)
			require.NoError(t, err)

			// out of gas error
			_, _, err = contract.ExtendedRun(nil, bz, precompiles.MERGE_GAS-1, false)
			require.ErrorIs(t, err, vm.ErrOutOfGas)

			// success
			resBz, _, err := contract.ExtendedRun(nil, bz, precompiles.MERGE_GAS+uint64(len(bz)), false)
			if tc.expectedErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			res, err := abi.Unpack(precompiles.METHOD_MERGE_JSON, resBz)
			require.NoError(t, err)
			require.Equal(t, tc.expected, res[0].(string))
		})
	}
}

func Test_JSONUtilsPrecompile_Stringify(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "simple map",
			input:    `{"a": 1, "b": 2}`,
			expected: `"{\"a\": 1, \"b\": 2}"`,
		},
		{
			name:     "nested map",
			input:    `{"a": 1, "b": {"c": 2}}`,
			expected: `"{\"a\": 1, \"b\": {\"c\": 2}}"`,
		},
		{
			name:     "slice",
			input:    `[1,2,3]`,
			expected: `"[1,2,3]"`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx, _, _ := setup()
			stateDB := NewMockStateDB(ctx)
			contract, err := precompiles.NewJSONUtilsPrecompile(stateDB)
			require.NoError(t, err)

			abi, err := contracts.IJsonutilsMetaData.GetAbi()
			require.NoError(t, err)

			bz, err := abi.Pack(precompiles.METHOD_STRINGIFY_JSON, tc.input)
			require.NoError(t, err)

			// out of gas error
			_, _, err = contract.ExtendedRun(nil, bz, precompiles.STRINGIFY_JSON_GAS-1, false)
			require.ErrorIs(t, err, vm.ErrOutOfGas)

			// success
			resBz, _, err := contract.ExtendedRun(nil, bz, precompiles.STRINGIFY_JSON_GAS+uint64(len(bz)), false)
			require.NoError(t, err)
			res, err := abi.Unpack(precompiles.METHOD_STRINGIFY_JSON, resBz)
			require.NoError(t, err)
			require.Equal(t, tc.expected, res[0].(string))
		})
	}
}
