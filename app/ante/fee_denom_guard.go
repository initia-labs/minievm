package ante

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	storetypes "cosmossdk.io/store/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	evmtypes "github.com/initia-labs/minievm/x/evm/types"
)

// MinGasPricesKeeper returns the chain-level minimum gas prices.
type MinGasPricesKeeper interface {
	MinGasPrices(ctx context.Context) (sdk.DecCoins, error)
}

// FeeDenomGuardDecorator rejects fee coins whose deduction would run untrusted
// EVM code.
//
// An `evm/<address>` denom maps straight to a contract address, so deducting a
// fee in it calls that contract's sudoTransfer. Fee deduction runs before
// signature verification, and a rejected tx pays nothing, so any contract could
// burn the whole tx gas limit in CheckTx (or in simulation) for free. Only the
// EVM fee denom and denoms with a configured min gas price may be paid as
// `evm/` fee coins. Other denoms resolve to factory-created ERC20s.
//
// The guard only runs in CheckTx and simulation, the two untrusted entry points.
// Proposal building and block execution are unchanged: a tx only gets there after
// passing CheckTx, and node-local min gas prices are not available in those modes.
type FeeDenomGuardDecorator struct {
	ek  EVMKeeper
	mgp MinGasPricesKeeper
}

// NewFeeDenomGuardDecorator creates a new FeeDenomGuardDecorator
func NewFeeDenomGuardDecorator(ek EVMKeeper, mgp MinGasPricesKeeper) FeeDenomGuardDecorator {
	return FeeDenomGuardDecorator{ek: ek, mgp: mgp}
}

func (d FeeDenomGuardDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	if !ctx.IsCheckTx() && !simulate {
		return next(ctx, tx, simulate)
	}

	feeTx, ok := tx.(sdk.FeeTx)
	if !ok {
		return ctx, errorsmod.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
	}

	var allowed map[string]bool
	for _, coin := range feeTx.GetFee() {
		if !evmtypes.IsERC20Denom(coin.Denom) {
			continue
		}

		if allowed == nil {
			var err error
			if allowed, err = d.allowedERC20FeeDenoms(ctx); err != nil {
				return ctx, err
			}
		}

		if !allowed[coin.Denom] {
			return ctx, errorsmod.Wrapf(sdkerrors.ErrInvalidCoins, "fee denom %s is not allowed", coin.Denom)
		}
	}

	return next(ctx, tx, simulate)
}

// allowedERC20FeeDenoms returns the fee denom and every denom with a min gas
// price, from chain params and from the local node config.
func (d FeeDenomGuardDecorator) allowedERC20FeeDenoms(ctx sdk.Context) (map[string]bool, error) {
	// reads must not consume tx gas, so simulation estimates stay unchanged
	readCtx := ctx.WithGasMeter(storetypes.NewInfiniteGasMeter())

	allowed := map[string]bool{}
	feeDenom, err := d.ek.GetFeeDenom(readCtx)
	if err != nil {
		return nil, err
	}
	allowed[feeDenom] = true

	minGasPrices, err := d.mgp.MinGasPrices(readCtx)
	if err != nil {
		return nil, err
	}
	for _, gp := range minGasPrices {
		allowed[gp.Denom] = true
	}
	for _, gp := range ctx.MinGasPrices() {
		allowed[gp.Denom] = true
	}

	return allowed, nil
}
