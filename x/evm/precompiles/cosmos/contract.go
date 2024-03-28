package cosmosprecompile

import (
	"context"
	"errors"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"

	"cosmossdk.io/core/address"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/initia-labs/minievm/x/evm/contracts/i_cosmos"
	"github.com/initia-labs/minievm/x/evm/types"
)

var _ vm.ExtendedPrecompiledContract = CosmosPrecompile{}
var _ vm.PrecompiledContract = CosmosPrecompile{}
var _ types.WithContext = CosmosPrecompile{}

type CosmosPrecompile struct {
	*abi.ABI

	ctx context.Context
	cdc codec.Codec
	ac  address.Codec
}

func NewCosmosPrecompile(cdc codec.Codec, ac address.Codec) (CosmosPrecompile, error) {
	abi, err := i_cosmos.ICosmosMetaData.GetAbi()
	if err != nil {
		return CosmosPrecompile{}, err
	}

	return CosmosPrecompile{ABI: abi, cdc: cdc, ac: ac}, nil
}

func (e CosmosPrecompile) WithContext(ctx context.Context) vm.PrecompiledContract {
	e.ctx = ctx
	return e
}

const (
	METHOD_TO_COSMOS_ADDRESS         = "to_cosmos_address"
	METHOD_TO_EVM_ADDRESS            = "to_evm_address"
	METHOD_TO_EXECUTE_COSMOS_MESSAGE = "execute_cosmos_message"
)

// ExtendedRun implements vm.ExtendedPrecompiledContract.
func (e CosmosPrecompile) ExtendedRun(caller vm.ContractRef, input []byte, suppliedGas uint64, readOnly bool) (resBz []byte, usedGas uint64, err error) {
	method, err := e.ABI.MethodById(input)
	if err != nil {
		return nil, 0, err
	}

	args, err := method.Inputs.Unpack(input[4:])
	if err != nil {
		return nil, 0, err
	}

	ctx := sdk.UnwrapSDKContext(e.ctx).WithGasMeter(storetypes.NewGasMeter(suppliedGas))
	switch method.Name {
	case METHOD_TO_COSMOS_ADDRESS:
		var toCosmosAddressArguments ToCosmosAddressArguments
		if err := method.Inputs.Copy(&toCosmosAddressArguments, args); err != nil {
			return nil, 0, err
		}

		addr, err := e.ac.BytesToString(toCosmosAddressArguments.EVMAddress.Bytes())
		if err != nil {
			return nil, 0, err
		}

		resBz, err = method.Outputs.Pack(addr)
		if err != nil {
			return nil, 0, err
		}

		ctx.GasMeter().ConsumeGas(TO_COSMOS_ADDRESS_GAS, "to_cosmos_address")
	case METHOD_TO_EVM_ADDRESS:
		var toEVMAddressArguments ToEVMAddressArguments
		if err := method.Inputs.Copy(&toEVMAddressArguments, args); err != nil {
			return nil, 0, err
		}

		addr, err := e.ac.StringToBytes(toEVMAddressArguments.CosmosAddress)
		if err != nil {
			return nil, 0, err
		}

		// check address length
		if len(addr) != common.AddressLength {
			return nil, 0, types.ErrInvalidAddressLength.Wrap(common.Bytes2Hex(addr))
		}

		resBz, err = method.Outputs.Pack(common.BytesToAddress(addr))
		if err != nil {
			return nil, 0, err
		}

		ctx.GasMeter().ConsumeGas(TO_EVM_ADDRESS_GAS, "to_evm_address")
	// case METHOD_TO_EXECUTE_COSMOS_MESSAGE:

	default:
		return nil, 0, types.ErrUnknownPrecompileMethod.Wrap(method.Name)
	}

	usedGas = ctx.GasMeter().GasConsumedToLimit()
	return resBz, usedGas, nil
}

// RequiredGas implements vm.PrecompiledContract.
func (e CosmosPrecompile) RequiredGas(input []byte) uint64 {
	return 0
}

// Run implements vm.PrecompiledContract.
func (e CosmosPrecompile) Run(input []byte) ([]byte, error) {
	return nil, errors.New("the CosmosPrecompile works exclusively with ExtendedRun")
}
