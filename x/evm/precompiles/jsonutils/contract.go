package jsonutils

import (
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/pkg/errors"

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
	METHOD_MERGE_JSON     = "merge_json"
	METHOD_STRINGIFY_JSON = "stringify_json"
)

// ExtendedRun implements vm.ExtendedPrecompiledContract.
func (e *JSONUtilsPrecompile) ExtendedRun(caller vm.ContractRef, input []byte, suppliedGas uint64, readOnly bool) (resBz []byte, usedGas uint64, err error) {
	snapshot := e.stateDB.Snapshot()
	ctx := e.stateDB.ContextOfSnapshot(snapshot).WithGasMeter(storetypes.NewGasMeter(suppliedGas))

	defer func() {
		if r := recover(); r != nil {
			switch r.(type) {
			case storetypes.ErrorOutOfGas:
				// convert cosmos out of gas error to EVM out of gas error
				usedGas = suppliedGas
				err = vm.ErrOutOfGas
			default:
				panic(r)
			}
		}
		if err != nil {
			// convert cosmos error to EVM error
			if err != vm.ErrOutOfGas {
				resBz = types.NewRevertReason(err)
				err = vm.ErrExecutionReverted
			}

			// revert the stateDB to the snapshot
			e.stateDB.RevertToSnapshot(snapshot)
		}
	}()

	method, err := e.ABI.MethodById(input)
	if err != nil {
		return nil, 0, types.ErrPrecompileFailed.Wrap(err.Error())
	}

	args, err := method.Inputs.Unpack(input[4:])
	if err != nil {
		return nil, 0, types.ErrPrecompileFailed.Wrap(err.Error())
	}

	// charge input gas
	ctx.GasMeter().ConsumeGas(storetypes.Gas(len(input))*GAS_PER_BYTE, "input bytes")

	switch method.Name {
	case METHOD_MERGE_JSON:
		ctx.GasMeter().ConsumeGas(MERGE_GAS, "merge_json")

		var MergeJSONArguments MergeJSONArguments
		if err := method.Inputs.Copy(&MergeJSONArguments, args); err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		resJSON, err := MergeJSON(MergeJSONArguments.DstJSON, MergeJSONArguments.SrcJSON)
		if err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		resBz, err = method.Outputs.Pack(resJSON)
		if err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}
	case METHOD_STRINGIFY_JSON:
		ctx.GasMeter().ConsumeGas(STRINGIFY_JSON_GAS, "stringify_json")

		var StringifyArguments StringifyArguments
		if err := method.Inputs.Copy(&StringifyArguments, args); err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		resJSON, err := json.Marshal(StringifyArguments.JSON)
		if err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		resBz, err = method.Outputs.Pack(string(resJSON))
		if err != nil {
			fmt.Println("err", err)
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
