package ante

import (
	errorsmod "cosmossdk.io/errors"

	ibcante "github.com/cosmos/ibc-go/v8/modules/core/ante"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"

	opchildante "github.com/initia-labs/OPinit/x/opchild/ante"

	"github.com/initia-labs/initia/app/ante/sigverify"
)

// NewMinimalAnteHandler returns a reduced AnteHandler chain for CheckTx mode.
// It validates signatures, format, gas limits, and fees (for priority) but
// does not deduct fees or increment sequences since those are handled by the
// full handler during PrepareProposal/FinalizeBlock.
func NewMinimalAnteHandler(options HandlerOptions) (sdk.AnteHandler, error) {
	if options.AccountKeeper == nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "account keeper is required for minimal ante handler")
	}
	if options.SignModeHandler == nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "sign mode handler is required for minimal ante handler")
	}
	if options.EVMKeeper == nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "EVM keeper is required for minimal ante handler")
	}
	if options.OPChildKeeper == nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "opchild keeper is required for minimal ante handler")
	}
	if options.IBCkeeper == nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "IBC keeper is required for minimal ante handler")
	}

	sigGasConsumer := options.SigGasConsumer
	if sigGasConsumer == nil {
		sigGasConsumer = sigverify.DefaultSigVerificationGasConsumer
	}

	txFeeChecker := options.TxFeeChecker
	if txFeeChecker == nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "tx fee checker is required for minimal ante handler")
	}

	anteDecorators := []sdk.AnteDecorator{
		ante.NewSetUpContextDecorator(),
		ante.NewExtensionOptionsDecorator(options.ExtensionOptionChecker),
		NewEthTxDecorator(options.EVMKeeper),
		NewGasPricesDecorator(),
		ante.NewValidateBasicDecorator(),
		ante.NewTxTimeoutHeightDecorator(),
		ante.NewValidateMemoDecorator(options.AccountKeeper),
		NewConsumeTxSizeGasDecorator(options.AccountKeeper),
		NewCheckFeeDecorator(txFeeChecker), // validate fee + set priority, no deduction
		ante.NewSetPubKeyDecorator(options.AccountKeeper),
		ante.NewValidateSigCountDecorator(options.AccountKeeper),
		NewSigGasConsumeDecorator(options.AccountKeeper, sigGasConsumer),
		NewSigVerificationDecorator(options.AccountKeeper, options.SignModeHandler),
		// no IncrementSequenceDecorator here since mempool tracks nonces
		ibcante.NewRedundantRelayDecorator(options.IBCkeeper),
		opchildante.NewRedundantBridgeDecorator(options.OPChildKeeper),
	}

	return sdk.ChainAnteDecorators(anteDecorators...), nil
}
