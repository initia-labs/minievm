package ante

import (
	errorsmod "cosmossdk.io/errors"
	storetypes "cosmossdk.io/store/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/cosmos/cosmos-sdk/x/auth/types"

	evmkeeper "github.com/initia-labs/minievm/x/evm/keeper"
)

// feeDeductionGasAmount is a estimated gas amount of fee payment
const feeDeductionGasAmount = 250_000

// GasFreeFeeDecorator is a decorator that sets the gas meter to infinite before calling the inner DeductFeeDecorator
// and then resets the gas meter to the original value after the inner DeductFeeDecorator is called.
//
// This gas meter manipulation only happens when the tx contains a fee which is defined as fee denom in x/evm params.
type GasFreeFeeDecorator struct {
	inner ante.DeductFeeDecorator

	// ek is used to get the fee denom from the x/evm params.
	ek *evmkeeper.Keeper
}

func NewGasFreeFeeDecorator(
	ak ante.AccountKeeper, bk types.BankKeeper,
	fk ante.FeegrantKeeper, ek *evmkeeper.Keeper,
	tfc ante.TxFeeChecker) GasFreeFeeDecorator {
	return GasFreeFeeDecorator{
		inner: ante.NewDeductFeeDecorator(ak, bk, fk, tfc),
		ek:    ek,
	}
}

func (fd GasFreeFeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	feeTx, ok := tx.(sdk.FeeTx)
	if !ok {
		return ctx, errorsmod.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
	}

	fees := feeTx.GetFee()
	feeDenom, err := fd.ek.GetFeeDenom(ctx.WithGasMeter(storetypes.NewInfiniteGasMeter()))
	if !(err == nil && len(fees) == 1 && fees[0].Denom == feeDenom) {
		if simulate && fees.IsZero() {
			// Charge gas for fee deduction simulation
			//
			// At gas simulation normally gas amount is zero, so the gas is not charged in the simulation.
			ctx.GasMeter().ConsumeGas(feeDeductionGasAmount, "fee deduction")
		}

		return fd.inner.AnteHandle(ctx, tx, simulate, next)
	}

	// If the fee contains only one denom and it is the fee denom, set the gas meter to infinite
	// to avoid gas consumption for fee deduction.
	gasMeter := ctx.GasMeter()
	ctx, err = fd.inner.AnteHandle(ctx.WithGasMeter(storetypes.NewInfiniteGasMeter()), tx, simulate, next)
	return ctx.WithGasMeter(gasMeter), err
}
