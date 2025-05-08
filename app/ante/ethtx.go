package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	evmtypes "github.com/initia-labs/minievm/x/evm/types"
)

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

func (d EthTxDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	ethTx, expectedSender, err := d.ek.TxUtils().ConvertCosmosTxToEthereumTx(ctx, tx)
	if err != nil {
		return ctx, err
	} else if ethTx == nil {
		return next(ctx, tx, simulate)
	}

	ctx = ctx.WithValue(evmtypes.CONTEXT_KEY_ETH_TX, ethTx)
	ctx = ctx.WithValue(evmtypes.CONTEXT_KEY_ETH_TX_SENDER, expectedSender)

	return next(ctx, tx, simulate)
}
