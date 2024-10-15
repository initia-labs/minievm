package posthandler

import (
	"fmt"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	storetypes "cosmossdk.io/store/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	evmante "github.com/initia-labs/minievm/x/evm/ante"
)

var _ sdk.PostDecorator = &GasRefundDecorator{}

type GasRefundDecorator struct {
	ek EVMKeeper
}

func NewGasRefundDecorator(ek EVMKeeper) sdk.PostDecorator {
	return &GasRefundDecorator{
		ek,
	}
}

// PostHandle handles the gas refund logic for EVM transactions.
func (erd *GasRefundDecorator) PostHandle(ctx sdk.Context, tx sdk.Tx, simulate, success bool, next sdk.PostHandler) (newCtx sdk.Context, err error) {
	if success && ctx.ExecMode() == sdk.ExecModeFinalize {
		feeTx, ok := tx.(sdk.FeeTx)
		if !ok {
			return ctx, errorsmod.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
		}

		value := ctx.Value(evmante.ContextKeyGasPrices)
		if value == nil {
			return ctx, nil
		}
		gasRefundRatio, err := erd.ek.GasRefundRatio(ctx)
		if err != nil {
			return ctx, err
		}

		gasPrices := value.(sdk.DecCoins)
		gasLeft := ctx.GasMeter().Limit() - ctx.GasMeter().GasConsumedToLimit()
		gasRefund := gasRefundRatio.MulInt64(int64(gasLeft)).TruncateInt().Uint64()

		// gas used for refund operation
		coinsRefund, _ := gasPrices.MulDec(math.LegacyNewDecFromInt(math.NewIntFromUint64(gasRefund))).TruncateDecimal()
		if coinsRefund.Empty() || coinsRefund.IsZero() {
			return ctx, nil
		}

		feePayer := feeTx.FeePayer()
		if feeGranter := feeTx.FeeGranter(); feeGranter != nil {
			feePayer = feeGranter
		}

		// emit gas refund event
		ctx.EventManager().EmitEvent(sdk.NewEvent(
			EventTypeGasRefund,
			sdk.NewAttribute(AttributeKeyGas, fmt.Sprintf("%d", gasRefund)),
			sdk.NewAttribute(AttributeKeyCoins, coinsRefund.String()),
		))

		// TODO - should we charge gas for refund?
		//
		// for now, we use infinite gas meter to prevent out of gas error or inconsistency between
		// used gas and refunded gas.
		feeCollectorAddr := authtypes.NewModuleAddress(authtypes.FeeCollectorName)
		err = erd.ek.ERC20Keeper().SendCoins(ctx.WithGasMeter(storetypes.NewInfiniteGasMeter()), feeCollectorAddr, feePayer, coinsRefund)
		if err != nil {
			return ctx, err
		}
	}

	return next(ctx, tx, simulate, success)
}

const (
	EventTypeGasRefund = "gas_refund"
	AttributeKeyGas    = "gas"
	AttributeKeyCoins  = "coins"
)
