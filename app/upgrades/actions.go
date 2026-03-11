package upgrades

import (
	"context"

	"github.com/ethereum/go-ethereum/common/hexutil"

	opchildtypes "github.com/initia-labs/OPinit/x/opchild/types"

	"github.com/initia-labs/minievm/app/upgrades/contracts/erc20_factory"
	"github.com/initia-labs/minievm/app/upgrades/contracts/erc20_wrapper"
)

// BindOPInitPort binds the OPinit IBC port if not already bound.
func BindOPInitPort(ctx context.Context, app MinitiaApp) error {
	bound, err := app.GetOPChildKeeper().IsBound(ctx, opchildtypes.PortID)
	if err != nil {
		return err
	}
	if !bound {
		if err := app.GetOPChildKeeper().BindPort(ctx, opchildtypes.PortID); err != nil {
			return err
		}
	}
	return nil
}

// NormalizeEVMParams normalizes EVM param addresses via address codec.
func NormalizeEVMParams(ctx context.Context, app MinitiaApp) error {
	params, err := app.GetEVMKeeper().Params.Get(ctx)
	if err != nil {
		return err
	}

	err = params.NormalizeAddresses(app.GetAccountKeeper().AddressCodec())
	if err != nil {
		return err
	}

	return app.GetEVMKeeper().Params.Set(ctx, params)
}

// UpdateERC20WrapperContract replaces ERC20Wrapper bytecode from compiled bindings.
func UpdateERC20WrapperContract(ctx context.Context, app MinitiaApp) error {
	wrapperAddr, err := app.GetEVMKeeper().GetERC20WrapperAddr(ctx)
	if err != nil {
		return err
	}
	wrapperRuntimeCode, err := hexutil.Decode(erc20_wrapper.Erc20WrapperBin)
	if err != nil {
		return err
	}
	wrapperCodeHash := CodeHash(wrapperRuntimeCode)
	return ReplaceCodeAndCodeHash(ctx, app, wrapperAddr.Bytes(), wrapperRuntimeCode, wrapperCodeHash)
}

// UpdateERC20FactoryContract replaces ERC20Factory bytecode from compiled bindings.
func UpdateERC20FactoryContract(ctx context.Context, app MinitiaApp) error {
	factoryAddr, err := app.GetEVMKeeper().GetERC20FactoryAddr(ctx)
	if err != nil {
		return err
	}
	factoryRuntimeCode, err := hexutil.Decode(erc20_factory.Erc20FactoryBin)
	if err != nil {
		return err
	}
	factoryCodeHash := CodeHash(factoryRuntimeCode)
	return ReplaceCodeAndCodeHash(ctx, app, factoryAddr.Bytes(), factoryRuntimeCode, factoryCodeHash)
}
