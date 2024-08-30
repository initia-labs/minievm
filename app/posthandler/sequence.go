package posthandler

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	evmtypes "github.com/initia-labs/minievm/x/evm/types"
)

var _ sdk.PostDecorator = &SequenceIncrementDecorator{}

type SequenceIncrementDecorator struct {
	ak authante.AccountKeeper
}

func NewSequenceIncrementDecorator(ak authante.AccountKeeper) sdk.PostDecorator {
	return &SequenceIncrementDecorator{
		ak,
	}
}

// If a transaction fails, we need to revert the sequence decrement for EVM messages that were executed in the ante handler.
// This is necessary because the sequence increment in EVM was reverted due to the failure.
func (sid *SequenceIncrementDecorator) PostHandle(ctx sdk.Context, tx sdk.Tx, simulate, success bool, next sdk.PostHandler) (newCtx sdk.Context, err error) {
	if !success && ctx.ExecMode() == sdk.ExecModeFinalize {
		signerMap := make(map[string]bool)
		for _, msg := range tx.GetMsgs() {
			var caller string
			switch msg := msg.(type) {
			case *evmtypes.MsgCreate:
				caller = msg.Sender
			case *evmtypes.MsgCreate2:
				caller = msg.Sender
			case *evmtypes.MsgCall:
				caller = msg.Sender
			default:
				continue
			}

			if _, ok := signerMap[caller]; ok {
				continue
			}
			signerMap[caller] = true

			callerAccAddr, err := sid.ak.AddressCodec().StringToBytes(caller)
			if err != nil {
				return ctx, err
			}

			acc := sid.ak.GetAccount(ctx, callerAccAddr)
			if err := acc.SetSequence(acc.GetSequence() + 1); err != nil {
				panic(err)
			}

			sid.ak.SetAccount(ctx, acc)
		}
	}

	return next(ctx, tx, simulate, success)
}
