package ante

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
)

// IncrementSequenceDecorator is an AnteDecorator that increments the sequence number
// of all signers in a transaction. It also sets a flag in the context to indicate
// that the sequence has been incremented in the ante handler.
type IncrementSequenceDecorator struct {
	ak authante.AccountKeeper
}

func NewIncrementSequenceDecorator(ak authante.AccountKeeper) IncrementSequenceDecorator {
	return IncrementSequenceDecorator{
		ak: ak,
	}
}

func (isd IncrementSequenceDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	sigTx, ok := tx.(authsigning.SigVerifiableTx)
	if !ok {
		return ctx, errorsmod.Wrap(sdkerrors.ErrTxDecode, "invalid transaction type")
	}

	// increment sequence of all signers
	signers, err := sigTx.GetSigners()
	if err != nil {
		return sdk.Context{}, err
	}

	for _, signer := range signers {
		acc := isd.ak.GetAccount(ctx, signer)
		if err := acc.SetSequence(acc.GetSequence() + 1); err != nil {
			panic(err)
		}

		isd.ak.SetAccount(ctx, acc)
	}

	// set a flag in context to indicate that sequence has been incremented in ante handler
	incremented := true // use pointer to enable revert after first call
	ctx = ctx.WithValue(ContextKeySequenceIncremented, &incremented)
	return next(ctx, tx, simulate)
}
