package posthandler

import (
	"fmt"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/log"
	"cosmossdk.io/math"
	storetypes "cosmossdk.io/store/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	evmante "github.com/initia-labs/minievm/x/evm/ante"
)

var _ sdk.PostDecorator = &GasRefundDecorator{}

type GasRefundDecorator struct {
	logger log.Logger
	ek     EVMKeeper
}

func NewGasRefundDecorator(logger log.Logger, ek EVMKeeper) sdk.PostDecorator {
	logger = logger.With("module", "gas_refund")
	return &GasRefundDecorator{
		logger, ek,
	}
}

// PostHandle handles the gas refund logic for EVM transactions.
func (erd *GasRefundDecorator) PostHandle(ctx sdk.Context, tx sdk.Tx, simulate, success bool, next sdk.PostHandler) (newCtx sdk.Context, err error) {
	if success && ctx.ExecMode() == sdk.ExecModeFinalize {
		// Conduct gas refund only for the successful EVM transactions
		if ok, err := erd.ek.TxUtils().IsEthereumTx(ctx, tx); err != nil || !ok {
			return next(ctx, tx, simulate, success)
		}

		feeTx, ok := tx.(sdk.FeeTx)
		if !ok {
			return ctx, errorsmod.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
		}

		value := ctx.Value(evmante.ContextKeyGasPrices)
		if value == nil {
			return next(ctx, tx, simulate, success)
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
			return next(ctx, tx, simulate, success)
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

		// conduct gas refund
		erd.safeRefund(ctx, feePayer, coinsRefund)
	}

	return next(ctx, tx, simulate, success)
}

func (erd *GasRefundDecorator) safeRefund(ctx sdk.Context, feePayer sdk.AccAddress, coinsRefund sdk.Coins) {
	defer func() {
		if r := recover(); r != nil {
			switch r := r.(type) {
			case storetypes.ErrorOutOfGas:
				erd.logger.Error("failed to refund gas", "err", r, "feePayer", feePayer, "refundAmount", coinsRefund)
			default:
				panic(r)
			}
		}
	}()

	// prepare context for refund operation
	const gasLimitForRefund = 1_000_000
	cacheCtx, commit := ctx.CacheContext()
	cacheCtx = cacheCtx.WithGasMeter(storetypes.NewGasMeter(gasLimitForRefund))

	feeCollectorAddr := authtypes.NewModuleAddress(authtypes.FeeCollectorName)
	err := erd.ek.ERC20Keeper().SendCoins(cacheCtx, feeCollectorAddr, feePayer, coinsRefund)
	if err != nil {
		erd.logger.Error("failed to refund gas", "err", err)
		return
	}

	commit()
}

const (
	EventTypeGasRefund = "gas_refund"
	AttributeKeyGas    = "gas"
	AttributeKeyCoins  = "coins"
)
