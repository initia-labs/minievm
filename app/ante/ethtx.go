package ante

import sdk "github.com/cosmos/cosmos-sdk/types"

// EthTxDecorator is a decorator that sets ethereum tx to the context.
//
// This decorator only happens when the tx contains a ethereum transaction.
type EthTxDecorator struct {
	ek EVMKeeper
}

func NewEthTxDecorator(ek EVMKeeper) EthTxDecorator {
	return EthTxDecorator{
		ek: ek,
	}
}

const (
	ContextKeyEthTx = iota
	ContextKeyEthTxSender
)

func (d EthTxDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	ethTx, expectedSender, err := d.ek.TxUtils().ConvertCosmosTxToEthereumTx(ctx, tx)
	if err != nil {
		return ctx, err
	} else if ethTx == nil {
		return next(ctx, tx, simulate)
	}

	ctx = ctx.WithValue(ContextKeyEthTx, ethTx)
	ctx = ctx.WithValue(ContextKeyEthTxSender, expectedSender)

	return next(ctx, tx, simulate)
}
