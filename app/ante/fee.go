package ante

import (
	errorsmod "cosmossdk.io/errors"
	storetypes "cosmossdk.io/store/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

// GasFreeFeeDecorator is a decorator that sets the gas meter to infinite before calling the inner DeductFeeDecorator
// and then resets the gas meter to the original value after the inner DeductFeeDecorator is called.
//
// This gas meter manipulation only happens when the tx contains a fee which is defined as fee denom in x/evm params.
type GasFreeFeeDecorator struct {
	inner ante.DeductFeeDecorator

	// ek is used to get the fee denom from the x/evm params.
	ek EVMKeeper
}

func NewGasFreeFeeDecorator(
	ak ante.AccountKeeper, bk types.BankKeeper,
	fk ante.FeegrantKeeper, ek EVMKeeper,
	tfc ante.TxFeeChecker) GasFreeFeeDecorator {
	return GasFreeFeeDecorator{
		inner: ante.NewDeductFeeDecorator(ak, bk, fk, tfc),
		ek:    ek,
	}
}

// gasLimitForFeeDeduction is the gas limit used for fee deduction.
const gasLimitForFeeDeduction = 1_000_000

func (fd GasFreeFeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	feeTx, ok := tx.(sdk.FeeTx)
	if !ok {
		return ctx, errorsmod.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
	}

	fees := feeTx.GetFee()
	feeDenom, err := fd.ek.GetFeeDenom(ctx.WithGasMeter(storetypes.NewInfiniteGasMeter()))
	if !(err == nil && len(fees) == 1 && fees[0].Denom == feeDenom) {
		return fd.inner.AnteHandle(ctx, tx, simulate, next)
	}

	// If the fee contains only one denom and it is the fee denom, set the gas meter to infinite
	// to avoid gas consumption for fee deduction.
	gasMeter := ctx.GasMeter()
	ctx, err = fd.inner.AnteHandle(ctx.WithGasMeter(storetypes.NewGasMeter(gasLimitForFeeDeduction)), tx, simulate, noopAnteHandler)
	// restore the original gas meter
	ctx = ctx.WithGasMeter(gasMeter)
	if err != nil {
		return ctx, err
	}

	return next(ctx, tx, simulate)
}

func noopAnteHandler(ctx sdk.Context, _ sdk.Tx, _ bool) (sdk.Context, error) {
	return ctx, nil
}
