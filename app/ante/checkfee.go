package ante

import (
	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
)

// CheckFeeDecorator validates that the tx meets minimum fee requirements and
// sets the priority on the context. Unlike DeductFeeDecorator, it does not
// deduct fees from the sender's account. That happens in the full handler
// during PrepareProposal/FinalizeBlock.
type CheckFeeDecorator struct {
	feeChecker ante.TxFeeChecker
}

// NewCheckFeeDecorator returns a CheckFeeDecorator using the given fee checker.
func NewCheckFeeDecorator(feeChecker ante.TxFeeChecker) CheckFeeDecorator {
	return CheckFeeDecorator{feeChecker: feeChecker}
}

func (cfd CheckFeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	if _, ok := tx.(sdk.FeeTx); !ok {
		return ctx, errorsmod.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
	}

	if !simulate {
		_, priority, err := cfd.feeChecker(ctx, tx)
		if err != nil {
			return ctx, err
		}
		ctx = ctx.WithPriority(priority)
	}

	return next(ctx, tx, simulate)
}
