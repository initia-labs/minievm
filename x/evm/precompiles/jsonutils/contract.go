package jsonutils

import (
	"encoding/json"
	"math/big"
	"slices"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/pkg/errors"

	"cosmossdk.io/math"
	storetypes "cosmossdk.io/store/types"

	"github.com/initia-labs/minievm/x/evm/contracts/i_jsonutils"
	"github.com/initia-labs/minievm/x/evm/types"
)

var _ vm.ExtendedPrecompiledContract = &JSONUtilsPrecompile{}
var _ vm.PrecompiledContract = &JSONUtilsPrecompile{}

var jsonUtilsABI *abi.ABI

func init() {
	var err error
	jsonUtilsABI, err = i_jsonutils.IJsonutilsMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
}

type JSONUtilsPrecompile struct {
	*abi.ABI
	stateDB types.StateDB
}

func NewJSONUtilsPrecompile(stateDB types.StateDB) (*JSONUtilsPrecompile, error) {
	return &JSONUtilsPrecompile{stateDB: stateDB, ABI: jsonUtilsABI}, nil
}

const (
	METHOD_MERGE_JSON            = "merge_json"
	METHOD_STRINGIFY_JSON        = "stringify_json"
	METHOD_UNMARSHAL_TO_OBJECT   = "unmarshal_to_object"
	METHOD_UNMARSHAL_TO_ARRAY    = "unmarshal_to_array"
	METHOD_UNMARSHAL_TO_STRING   = "unmarshal_to_string"
	METHOD_UNMARSHAL_TO_UINT     = "unmarshal_to_uint"
	METHOD_UNMARSHAL_TO_BOOL     = "unmarshal_to_bool"
	METHOD_UNMARSHAL_ISO_TO_UNIX = "unmarshal_iso_to_unix"
)

// ExtendedRun implements vm.ExtendedPrecompiledContract.
func (e *JSONUtilsPrecompile) ExtendedRun(caller common.Address, input []byte, suppliedGas uint64, readOnly bool) (resBz []byte, usedGas uint64, err error) {
	snapshot := e.stateDB.Snapshot()
	ctx := e.stateDB.Context().WithGasMeter(storetypes.NewGasMeter(suppliedGas))

	defer func() {
		if r := recover(); r != nil {
			switch r.(type) {
			case storetypes.ErrorOutOfGas:
				// set the used gas to the supplied gas
				usedGas = suppliedGas

				// convert cosmos out of gas error to normal error
				err = errors.New("out of gas in precompile")
			default:
				panic(r)
			}
		}

		if err != nil {
			// convert cosmos error to EVM error
			resBz = types.NewRevertReason(err)
			err = vm.ErrExecutionReverted

			// revert the stateDB to the snapshot
			e.stateDB.RevertToSnapshot(snapshot)
		}
	}()

	method, err := e.MethodById(input)
	if err != nil {
		return nil, 0, types.ErrPrecompileFailed.Wrap(err.Error())
	}

	args, err := method.Inputs.Unpack(input[4:])
	if err != nil {
		return nil, 0, types.ErrPrecompileFailed.Wrap(err.Error())
	}

	// charge input gas
	ctx.GasMeter().ConsumeGas(storetypes.Gas(len(input))*GAS_PER_BYTE, "input bytes")

	// check input size
	if len(input) > MaxJSONSize {
		return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap("input size exceeds the limit")
	}

	switch method.Name {
	case METHOD_MERGE_JSON:
		ctx.GasMeter().ConsumeGas(MERGE_GAS, "merge_json")

		var mergeJSONArguments MergeJSONArguments
		if err := method.Inputs.Copy(&mergeJSONArguments, args); err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		resJSON, err := MergeJSON(mergeJSONArguments.DstJSON, mergeJSONArguments.SrcJSON)
		if err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		resBz, err = method.Outputs.Pack(resJSON)
		if err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}
	case METHOD_STRINGIFY_JSON:
		ctx.GasMeter().ConsumeGas(STRINGIFY_JSON_GAS, "stringify_json")

		var stringifyArguments StringifyArguments
		if err := method.Inputs.Copy(&stringifyArguments, args); err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		resJSON, err := json.Marshal(stringifyArguments.JSON)
		if err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		resBz, err = method.Outputs.Pack(string(resJSON))
		if err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}
	case METHOD_UNMARSHAL_TO_OBJECT:
		ctx.GasMeter().ConsumeGas(UNMARSHAL_JSON_GAS, "unmarshal_to_object")

		var unmarshalJSONArguments UnmarshalJSONArguments
		if err := method.Inputs.Copy(&unmarshalJSONArguments, args); err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		var m map[string]interface{}
		if err := json.Unmarshal(unmarshalJSONArguments.JSONBytes, &m); err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		jsonElements := make([]i_jsonutils.IJSONUtilsJSONElement, len(m))

		i := 0
		for key, val := range m {
			// unmarshal already succeeded, so no need to check error
			valBz, _ := json.Marshal(val)
			jsonElements[i] = i_jsonutils.IJSONUtilsJSONElement{
				Key:   key,
				Value: valBz,
			}
			i++
		}

		// charge sort cost
		ctx.GasMeter().ConsumeGas(storetypes.Gas(len(jsonElements))*GAS_PER_SORT_ITEM, "sort items")

		// sort by key
		slices.SortFunc(jsonElements, func(a, b i_jsonutils.IJSONUtilsJSONElement) int {
			return strings.Compare(a.Key, b.Key)
		})

		resBz, err = method.Outputs.Pack(i_jsonutils.IJSONUtilsJSONObject{Elements: jsonElements})
		if err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

	case METHOD_UNMARSHAL_TO_ARRAY:
		ctx.GasMeter().ConsumeGas(UNMARSHAL_JSON_GAS, "unmarshal_to_array")

		var unmarshalJSONArguments UnmarshalJSONArguments
		if err := method.Inputs.Copy(&unmarshalJSONArguments, args); err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		var a []interface{}
		if err := json.Unmarshal(unmarshalJSONArguments.JSONBytes, &a); err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		bytesArray := make([][]byte, len(a))
		for i, val := range a {
			// unmarshal already succeeded, so no need to check error
			valBz, _ := json.Marshal(val)
			bytesArray[i] = valBz
		}

		resBz, err = method.Outputs.Pack(bytesArray)
		if err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

	case METHOD_UNMARSHAL_TO_STRING:
		ctx.GasMeter().ConsumeGas(UNMARSHAL_JSON_GAS, "unmarshal_to_string")

		var unmarshalJSONArguments UnmarshalJSONArguments
		if err := method.Inputs.Copy(&unmarshalJSONArguments, args); err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		var s string
		if err := json.Unmarshal(unmarshalJSONArguments.JSONBytes, &s); err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		resBz, err = method.Outputs.Pack(s)
		if err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

	case METHOD_UNMARSHAL_TO_UINT:
		ctx.GasMeter().ConsumeGas(UNMARSHAL_JSON_GAS, "unmarshal_to_uint")

		var unmarshalJSONArguments UnmarshalJSONArguments
		if err := method.Inputs.Copy(&unmarshalJSONArguments, args); err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		// try string first
		var n math.Uint
		if err := json.Unmarshal(unmarshalJSONArguments.JSONBytes, &n); err != nil {
			// try number
			var n2 uint64
			if err := json.Unmarshal(unmarshalJSONArguments.JSONBytes, &n2); err != nil {
				return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
			}

			n = math.NewUint(n2)
		}

		resBz, err = method.Outputs.Pack(n.BigInt())
		if err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

	case METHOD_UNMARSHAL_TO_BOOL:
		ctx.GasMeter().ConsumeGas(UNMARSHAL_JSON_GAS, "unmarshal_to_bool")

		var unmarshalJSONArguments UnmarshalJSONArguments
		if err := method.Inputs.Copy(&unmarshalJSONArguments, args); err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		var b bool
		if err := json.Unmarshal(unmarshalJSONArguments.JSONBytes, &b); err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		resBz, err = method.Outputs.Pack(b)
		if err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

	case METHOD_UNMARSHAL_ISO_TO_UNIX:
		ctx.GasMeter().ConsumeGas(UNMARSHAL_JSON_GAS, "unmarshal_iso_to_unix")

		var unmarshalJSONArguments UnmarshalJSONArguments
		if err := method.Inputs.Copy(&unmarshalJSONArguments, args); err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		var t string
		if err := json.Unmarshal(unmarshalJSONArguments.JSONBytes, &t); err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		tt, err := time.Parse(time.RFC3339Nano, t)
		if err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		resBz, err = method.Outputs.Pack(big.NewInt(tt.UnixNano()))
		if err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

	default:
		return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrUnknownPrecompileMethod.Wrap(method.Name)
	}

	usedGas = ctx.GasMeter().GasConsumedToLimit()
	return resBz, usedGas, nil
}

// RequiredGas implements vm.PrecompiledContract.
func (e *JSONUtilsPrecompile) RequiredGas(input []byte) uint64 {
	return 0
}

// Run implements vm.PrecompiledContract.
func (e *JSONUtilsPrecompile) Run(input []byte) ([]byte, error) {
	return nil, errors.New("the JSONUtilsPrecompile works exclusively with ExtendedRun")
}
