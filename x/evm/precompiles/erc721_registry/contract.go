package erc721registryprecompile

import (
	"context"
	"errors"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/vm"

	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/initia-labs/minievm/x/evm/contracts/i_erc721_registry"
	"github.com/initia-labs/minievm/x/evm/types"
)

var _ vm.ExtendedPrecompiledContract = ERC721RegistryPrecompile{}
var _ vm.PrecompiledContract = ERC721RegistryPrecompile{}
var _ types.WithContext = ERC721RegistryPrecompile{}

type ERC721RegistryPrecompile struct {
	*abi.ABI
	ctx context.Context
	k   types.IERC721StoresKeeper
}

func NewERC721RegistryPrecompile(k types.IERC721StoresKeeper) (ERC721RegistryPrecompile, error) {
	abi, err := i_erc721_registry.IErc721RegistryMetaData.GetAbi()
	if err != nil {
		return ERC721RegistryPrecompile{}, err
	}

	return ERC721RegistryPrecompile{ABI: abi, k: k}, nil
}

func (e ERC721RegistryPrecompile) WithContext(ctx context.Context) vm.PrecompiledContract {
	e.ctx = ctx
	return e
}

const (
	METHOD_REGISTER            = "register_erc721"
	METHOD_REGISTER_STORE      = "register_erc721_store"
	METHOD_IS_STORE_REGISTERED = "is_erc721_store_registered"
)

// ExtendedRun implements vm.ExtendedPrecompiledContract.
func (e ERC721RegistryPrecompile) ExtendedRun(caller vm.ContractRef, input []byte, suppliedGas uint64, readOnly bool) (resBz []byte, usedGas uint64, err error) {
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
		ctx.GasMeter().ConsumeGas(REGISTER_GAS, "register_erc721")

		if readOnly {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrNonReadOnlyMethod.Wrap(method.Name)
		}

		if err := e.k.Register(ctx, caller.Address()); err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}
	case METHOD_REGISTER_STORE:
		ctx.GasMeter().ConsumeGas(REGISTER_STORE_GAS, "register_erc721_store")

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
	case METHOD_IS_STORE_REGISTERED:
		ctx.GasMeter().ConsumeGas(IS_STORE_REGISTERED_GAS, "is_erc721_store_registered")

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
func (e ERC721RegistryPrecompile) RequiredGas(input []byte) uint64 {
	return 0
}

// Run implements vm.PrecompiledContract.
func (e ERC721RegistryPrecompile) Run(input []byte) ([]byte, error) {
	return nil, errors.New("the erc721RegistryPrecompile works exclusively with ExtendedRun")
}
