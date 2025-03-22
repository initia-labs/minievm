package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
)

// ConsumeTxSizeGasDecorator is a decorator that consumes gas for the tx size.
// It skips gas consumption for eth tx because we are handling this with eth intrinsic gas metering.
type ConsumeTxSizeGasDecorator struct {
	ante.ConsumeTxSizeGasDecorator
}

// NewConsumeTxSizeGasDecorator creates a new ConsumeTxSizeGasDecorator
func NewConsumeTxSizeGasDecorator(ak ante.AccountKeeper) ConsumeTxSizeGasDecorator {
	return ConsumeTxSizeGasDecorator{
		ConsumeTxSizeGasDecorator: ante.NewConsumeGasForTxSizeDecorator(ak),
	}
}

// AnteHandle consumes gas for the tx size.
// It skips gas consumption for eth tx because we are handling this with eth intrinsic gas metering.
func (d ConsumeTxSizeGasDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	ethTx := ctx.Value(ContextKeyEthTx)
	if ethTx != nil {
		// skip gas consumption for eth tx
		// we are handling this with eth intrinsic gas metering
		return next(ctx, tx, simulate)
	}

	return d.ConsumeTxSizeGasDecorator.AnteHandle(ctx, tx, simulate, next)
}

// SigGasConsumeDecorator is a decorator that consumes gas for the signature verification.
// It skips gas consumption for eth tx because we are handling this with eth intrinsic gas metering.
type SigGasConsumeDecorator struct {
	ante.SigGasConsumeDecorator
}

// NewSigGasConsumeDecorator creates a new SigGasConsumeDecorator
func NewSigGasConsumeDecorator(ak ante.AccountKeeper, sigGasConsumer ante.SignatureVerificationGasConsumer) SigGasConsumeDecorator {
	return SigGasConsumeDecorator{ante.NewSigGasConsumeDecorator(ak, sigGasConsumer)}
}

// AnteHandle consumes gas for the signature verification.
// It skips gas consumption for eth tx because we are handling this with eth intrinsic gas metering.
func (d SigGasConsumeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	ethTx := ctx.Value(ContextKeyEthTx)
	if ethTx != nil {
		// skip gas consumption for eth tx
		// we are handling this with eth intrinsic gas metering
		return next(ctx, tx, simulate)
	}

	return d.SigGasConsumeDecorator.AnteHandle(ctx, tx, simulate, next)
}
