package jsonutils_test

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"

	abiapi "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/vm"

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

func mustMarshalJSON(t *testing.T, obj interface{}) []byte {
	bz, err := json.Marshal(obj)
	require.NoError(t, err)
	return bz
}

func Test_JSONUtilsPrecompile_UnmarshalToObject(t *testing.T) {
	testCases := []struct {
		name        string
		input       string
		expected    contracts.IJSONUtilsJSONObject
		expectedErr bool
	}{
		{
			name:     "simple map",
			input:    `{"a": 1, "b": 2}`,
			expected: contracts.IJSONUtilsJSONObject{Elements: []contracts.IJSONUtilsJSONElement{{Key: "a", Value: mustMarshalJSON(t, 1)}, {Key: "b", Value: mustMarshalJSON(t, 2)}}},
		},
		{
			name:     "nested map",
			input:    `{"a": 1, "b": {"c": 2}}`,
			expected: contracts.IJSONUtilsJSONObject{Elements: []contracts.IJSONUtilsJSONElement{{Key: "a", Value: mustMarshalJSON(t, 1)}, {Key: "b", Value: mustMarshalJSON(t, map[string]int{"c": 2})}}},
		},
		{
			name:        "invalid json",
			input:       `{"a": 1, "b": 2`,
			expectedErr: true,
		},
		{
			name:        "slice",
			input:       `[1,2,3]`,
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

			bz, err := abi.Pack(precompiles.METHOD_UNMARSHAL_TO_OBJECT, []byte(tc.input))
			require.NoError(t, err)

			// out of gas error
			_, _, err = contract.ExtendedRun(nil, bz, precompiles.UNMARSHAL_JSON_GAS-1, false)
			require.ErrorIs(t, err, vm.ErrOutOfGas)

			resBz, _, err := contract.ExtendedRun(nil, bz, precompiles.UNMARSHAL_JSON_GAS+uint64(len(bz)), false)
			if tc.expectedErr {
				require.Error(t, err)
				return
			}

			// success
			require.NoError(t, err)

			resValues, err := abi.Unpack(precompiles.METHOD_UNMARSHAL_TO_OBJECT, resBz)
			require.NoError(t, err)

			res := *abiapi.ConvertType(resValues[0], new(contracts.IJSONUtilsJSONObject)).(*contracts.IJSONUtilsJSONObject)
			require.Equal(t, tc.expected, res)
		})
	}
}

func Test_JSONUtilsPrecompile_UnmarshalToString(t *testing.T) {
	testCases := []struct {
		name        string
		input       string
		expected    string
		expectedErr bool
	}{
		{
			name:     "string",
			input:    `"abc"`,
			expected: "abc",
		},
		{
			name:        "number",
			input:       `123`,
			expectedErr: true,
		},
		{
			name:        "map",
			input:       `{"a": 1, "b": 2}`,
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

			bz, err := abi.Pack(precompiles.METHOD_UNMARSHAL_TO_STRING, []byte(tc.input))
			require.NoError(t, err)

			// out of gas error
			_, _, err = contract.ExtendedRun(nil, bz, precompiles.UNMARSHAL_JSON_GAS-1, false)
			require.ErrorIs(t, err, vm.ErrOutOfGas)

			// success
			resBz, _, err := contract.ExtendedRun(nil, bz, precompiles.UNMARSHAL_JSON_GAS+uint64(len(bz)), false)
			if tc.expectedErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			resValues, err := abi.Unpack(precompiles.METHOD_UNMARSHAL_TO_STRING, resBz)
			require.NoError(t, err)

			res := *abiapi.ConvertType(resValues[0], new(string)).(*string)
			require.Equal(t, tc.expected, res)
		})
	}
}

func Test_JSONUtilsPrecompile_UnmarshalToUint(t *testing.T) {
	testCases := []struct {
		name        string
		input       string
		expected    *big.Int
		expectedErr bool
	}{
		{
			name:     "string number",
			input:    `"123"`,
			expected: new(big.Int).SetUint64(123),
		},
		{
			name:     "number",
			input:    `123`,
			expected: new(big.Int).SetUint64(123),
		},
		{
			name:        "map",
			input:       `{"a": 1, "b": 2}`,
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

			bz, err := abi.Pack(precompiles.METHOD_UNMARSHAL_TO_UINT, []byte(tc.input))
			require.NoError(t, err)

			// out of gas error
			_, _, err = contract.ExtendedRun(nil, bz, precompiles.UNMARSHAL_JSON_GAS-1, false)
			require.ErrorIs(t, err, vm.ErrOutOfGas)

			// success
			resBz, _, err := contract.ExtendedRun(nil, bz, precompiles.UNMARSHAL_JSON_GAS+uint64(len(bz)), false)
			if tc.expectedErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			resValues, err := abi.Unpack(precompiles.METHOD_UNMARSHAL_TO_UINT, resBz)
			require.NoError(t, err)

			res := abiapi.ConvertType(resValues[0], new(big.Int)).(*big.Int)
			require.Equal(t, tc.expected, res)
		})
	}
}

func Test_JSONUtilsPrecompile_UnmarshalToBool(t *testing.T) {
	testCases := []struct {
		name        string
		input       string
		expected    bool
		expectedErr bool
	}{
		{
			name:     "true",
			input:    `true`,
			expected: true,
		},
		{
			name:        "false",
			input:       `false`,
			expectedErr: false,
		},
		{
			name:        "map",
			input:       `{"a": 1, "b": 2}`,
			expectedErr: true,
		},
		{
			name:        "number",
			input:       `123`,
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

			bz, err := abi.Pack(precompiles.METHOD_UNMARSHAL_TO_BOOL, []byte(tc.input))
			require.NoError(t, err)

			// out of gas error
			_, _, err = contract.ExtendedRun(nil, bz, precompiles.UNMARSHAL_JSON_GAS-1, false)
			require.ErrorIs(t, err, vm.ErrOutOfGas)

			// success
			resBz, _, err := contract.ExtendedRun(nil, bz, precompiles.UNMARSHAL_JSON_GAS+uint64(len(bz)), false)
			if tc.expectedErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			resValues, err := abi.Unpack(precompiles.METHOD_UNMARSHAL_TO_BOOL, resBz)
			require.NoError(t, err)

			res := *abiapi.ConvertType(resValues[0], new(bool)).(*bool)
			require.Equal(t, tc.expected, res)
		})
	}
}

func Test_JSONUtilsPrecompile_UnmarshalToArray(t *testing.T) {
	testCases := []struct {
		name        string
		input       string
		expected    [][]byte
		expectedErr bool
	}{
		{
			name:     "simple array",
			input:    `[1, 2, 3]`,
			expected: [][]byte{mustMarshalJSON(t, 1), mustMarshalJSON(t, 2), mustMarshalJSON(t, 3)},
		},
		{
			name:     "string array",
			input:    `["a", "b", "c"]`,
			expected: [][]byte{mustMarshalJSON(t, "a"), mustMarshalJSON(t, "b"), mustMarshalJSON(t, "c")},
		},
		{
			name:     "object array",
			input:    `[{"a": 1}, {"b": 2}]`,
			expected: [][]byte{mustMarshalJSON(t, map[string]int{"a": 1}), mustMarshalJSON(t, map[string]int{"b": 2})},
		},
		{
			name:     "array of arrays",
			input:    `[[1, 2], [3, 4]]`,
			expected: [][]byte{mustMarshalJSON(t, []int{1, 2}), mustMarshalJSON(t, []int{3, 4})},
		},
		{
			name:        "invalid json",
			input:       `[1, 2, 3`,
			expectedErr: true,
		},
		{
			name:        "object",
			input:       `{"a": 1, "b": 2}`,
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

			bz, err := abi.Pack(precompiles.METHOD_UNMARSHAL_TO_ARRAY, []byte(tc.input))
			require.NoError(t, err)

			// out of gas error
			_, _, err = contract.ExtendedRun(nil, bz, precompiles.UNMARSHAL_JSON_GAS-1, false)
			require.ErrorIs(t, err, vm.ErrOutOfGas)

			// success
			resBz, _, err := contract.ExtendedRun(nil, bz, precompiles.UNMARSHAL_JSON_GAS+uint64(len(bz)), false)
			if tc.expectedErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			resValues, err := abi.Unpack(precompiles.METHOD_UNMARSHAL_TO_ARRAY, resBz)
			require.NoError(t, err)

			res := *abiapi.ConvertType(resValues[0], new([][]byte)).(*[][]byte)
			require.Equal(t, tc.expected, res)
		})
	}
}

func Test_JSONUtilsPrecompile_UnmarshalISOToUnix(t *testing.T) {
	testCases := []struct {
		name        string
		input       string
		expected    *big.Int
		expectedErr bool
	}{
		{
			name:     "valid iso",
			input:    `"2025-02-05T07:44:14.748093393Z"`,
			expected: big.NewInt(1738741454748093393),
		},
		{
			name:        "invalid iso",
			input:       `"2025-02-05T07:44:14.748093393"`,
			expectedErr: true,
		},
		{
			name:        "invalid json",
			input:       `"2025-02-05T07:44:14.748093393Z`,
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

			bz, err := abi.Pack(precompiles.METHOD_UNMARSHAL_ISO_TO_UNIX, []byte(tc.input))
			require.NoError(t, err)

			// out of gas error
			_, _, err = contract.ExtendedRun(nil, bz, precompiles.UNMARSHAL_JSON_GAS-1, false)
			require.ErrorIs(t, err, vm.ErrOutOfGas)

			// success
			resBz, _, err := contract.ExtendedRun(nil, bz, precompiles.UNMARSHAL_JSON_GAS+uint64(len(bz)), false)
			if tc.expectedErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			resValues, err := abi.Unpack(precompiles.METHOD_UNMARSHAL_ISO_TO_UNIX, resBz)
			require.NoError(t, err)

			res := abiapi.ConvertType(resValues[0], new(big.Int)).(*big.Int)
			require.Equal(t, tc.expected, res)
		})
	}
}
