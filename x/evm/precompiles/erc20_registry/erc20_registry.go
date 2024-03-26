package erc20registry

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

var _ vm.ExtendedPrecompiledContract = ERC20Registry{}
var _ vm.PrecompiledContract = ERC20Registry{}
var _ types.WithContext = ERC20Registry{}

type ERC20Registry struct {
	*abi.ABI
	ctx context.Context
	k   types.IERC20StoresKeeper
}

func NewERC20Registry(k types.IERC20StoresKeeper) (ERC20Registry, error) {
	abi, err := i_erc20_registry.IErc20RegistryMetaData.GetAbi()
	if err != nil {
		return ERC20Registry{}, err
	}

	return ERC20Registry{ABI: abi, k: k}, nil
}

func (e ERC20Registry) WithContext(ctx context.Context) vm.PrecompiledContract {
	e.ctx = ctx
	return e
}

const (
	METHOD_REGISTER            = "register_erc20"
	METHOD_REGISTER_STORE      = "register_erc20_store"
	METHOD_IS_STORE_REGISTERED = "is_erc20_store_registered"
)

// ExtendedRun implements vm.ExtendedPrecompiledContract.
func (e ERC20Registry) ExtendedRun(caller vm.ContractRef, input []byte, suppliedGas uint64, readOnly bool) (resBz []byte, usedGas uint64, err error) {
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
	case METHOD_REGISTER:
		if readOnly {
			return nil, 0, types.ErrNonReadOnlyMethod.Wrap(method.Name)
		}

		if err := e.k.Register(ctx, caller.Address()); err != nil {
			return nil, 0, err
		}
	case METHOD_REGISTER_STORE:
		if readOnly {
			return nil, 0, types.ErrNonReadOnlyMethod.Wrap(method.Name)
		}

		var registerArgs RegisterArguments
		if err := method.Inputs.Copy(&registerArgs, args); err != nil {
			return nil, 0, err
		}

		if err := e.k.RegisterStore(ctx, registerArgs.Account.Bytes(), caller.Address()); err != nil {
			return nil, 0, err
		}
	case METHOD_IS_STORE_REGISTERED:
		var isRegisteredArgs IsRegisteredArguments
		if err := method.Inputs.Copy(&isRegisteredArgs, args); err != nil {
			return nil, 0, err
		}

		ok, err := e.k.IsStoreRegistered(ctx, isRegisteredArgs.Account.Bytes(), caller.Address())
		if err != nil {
			return nil, 0, err
		}

		resBz, err = method.Outputs.Pack(ok)
		if err != nil {
			return nil, 0, err
		}
	default:
		return nil, 0, types.ErrUnknownPrecompileMethod.Wrap(method.Name)
	}

	usedGas = ctx.GasMeter().GasConsumedToLimit()
	return resBz, usedGas, nil
}

// RequiredGas implements vm.PrecompiledContract.
func (e ERC20Registry) RequiredGas(input []byte) uint64 {
	return 0
}

// Run implements vm.PrecompiledContract.
func (e ERC20Registry) Run(input []byte) ([]byte, error) {
	return nil, errors.New("the ERC20Registry works exclusively with ExtendedRun")
}
