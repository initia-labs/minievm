package erc20registryprecompile

import (
	"context"
	"errors"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/vm"

	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/initia-labs/minievm/x/evm/contracts/i_erc20_registry"
	"github.com/initia-labs/minievm/x/evm/types"
)

var _ vm.ExtendedPrecompiledContract = ERC20RegistryPrecompile{}
var _ vm.PrecompiledContract = ERC20RegistryPrecompile{}
var _ types.WithContext = ERC20RegistryPrecompile{}

type ERC20RegistryPrecompile struct {
	*abi.ABI
	ctx context.Context
	k   types.IERC20StoresKeeper
}

func NewERC20RegistryPrecompile(k types.IERC20StoresKeeper) (ERC20RegistryPrecompile, error) {
	abi, err := i_erc20_registry.IErc20RegistryMetaData.GetAbi()
	if err != nil {
		return ERC20RegistryPrecompile{}, err
	}

	return ERC20RegistryPrecompile{ABI: abi, k: k}, nil
}

func (e ERC20RegistryPrecompile) WithContext(ctx context.Context) vm.PrecompiledContract {
	e.ctx = ctx
	return e
}

const (
	METHOD_REGISTER              = "register_erc20"
	METHOD_REGISTER_FROM_FACTORY = "register_erc20_from_factory"
	METHOD_REGISTER_STORE        = "register_erc20_store"
	METHOD_IS_STORE_REGISTERED   = "is_erc20_store_registered"
)

// ExtendedRun implements vm.ExtendedPrecompiledContract.
func (e ERC20RegistryPrecompile) ExtendedRun(caller vm.ContractRef, input []byte, suppliedGas uint64, readOnly bool) (resBz []byte, usedGas uint64, err error) {
	method, err := e.ABI.MethodById(input)
	if err != nil {
		return nil, 0, types.ErrPrecompileFailed.Wrap(err.Error())
	}

	args, err := method.Inputs.Unpack(input[4:])
	if err != nil {
		return nil, 0, types.ErrPrecompileFailed.Wrap(err.Error())
	}

	ctx := sdk.UnwrapSDKContext(e.ctx).WithGasMeter(storetypes.NewGasMeter(suppliedGas))
	ctx.GasMeter().ConsumeGas(storetypes.Gas(len(input))*GAS_PER_BYTE, "input bytes")

	switch method.Name {
	case METHOD_REGISTER:
		ctx.GasMeter().ConsumeGas(REGISTER_GAS, "register_erc20")

		if readOnly {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrNonReadOnlyMethod.Wrap(method.Name)
		}

		if err := e.k.Register(ctx, caller.Address()); err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		resBz, err = method.Outputs.Pack(true)
		if err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}
	case METHOD_REGISTER_FROM_FACTORY:
		ctx.GasMeter().ConsumeGas(REGISTER_FROM_FACTORY_GAS, "register_erc20_from_factory")

		if readOnly {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrNonReadOnlyMethod.Wrap(method.Name)
		}

		var registerArgs RegisterERC20FromFactoryArguments
		if err := method.Inputs.Copy(&registerArgs, args); err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		if err := e.k.RegisterFromFactory(ctx, caller.Address(), registerArgs.ERC20); err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		resBz, err = method.Outputs.Pack(true)
		if err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}
	case METHOD_REGISTER_STORE:
		ctx.GasMeter().ConsumeGas(REGISTER_STORE_GAS, "register_erc20_store")

		if readOnly {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrNonReadOnlyMethod.Wrap(method.Name)
		}

		var registerArgs RegisterStoreArguments
		if err := method.Inputs.Copy(&registerArgs, args); err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		if err := e.k.RegisterStore(ctx, registerArgs.Account.Bytes(), caller.Address()); err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		resBz, err = method.Outputs.Pack(true)
		if err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}
	case METHOD_IS_STORE_REGISTERED:
		ctx.GasMeter().ConsumeGas(IS_STORE_REGISTERED_GAS, "is_erc20_store_registered")

		var isRegisteredArgs IsStoreRegisteredArguments
		if err := method.Inputs.Copy(&isRegisteredArgs, args); err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		ok, err := e.k.IsStoreRegistered(ctx, isRegisteredArgs.Account.Bytes(), caller.Address())
		if err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		resBz, err = method.Outputs.Pack(ok)
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
func (e ERC20RegistryPrecompile) RequiredGas(input []byte) uint64 {
	return 0
}

// Run implements vm.PrecompiledContract.
func (e ERC20RegistryPrecompile) Run(input []byte) ([]byte, error) {
	return nil, errors.New("the ERC20RegistryPrecompile works exclusively with ExtendedRun")
}
