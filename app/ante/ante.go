package ante

import (
	errorsmod "cosmossdk.io/errors"
	txsigning "cosmossdk.io/x/tx/signing"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	ibcante "github.com/cosmos/ibc-go/v8/modules/core/ante"
	ibckeeper "github.com/cosmos/ibc-go/v8/modules/core/keeper"

	opchildante "github.com/initia-labs/OPinit/x/opchild/ante"
	opchildkeeper "github.com/initia-labs/OPinit/x/opchild/keeper"
	"github.com/initia-labs/initia/app/ante/accnum"
	"github.com/initia-labs/initia/app/ante/sigverify"
	evmante "github.com/initia-labs/minievm/x/evm/ante"

	auctionante "github.com/skip-mev/block-sdk/v2/x/auction/ante"
	auctionkeeper "github.com/skip-mev/block-sdk/v2/x/auction/keeper"
)

// HandlerOptions extends the SDK's AnteHandler options by requiring the IBC
// channel keeper.
type HandlerOptions struct {
	ante.HandlerOptions
	Codec         codec.BinaryCodec
	IBCkeeper     *ibckeeper.Keeper
	OPChildKeeper *opchildkeeper.Keeper
	AuctionKeeper *auctionkeeper.Keeper
	EVMKeeper     EVMKeeper

	TxEncoder sdk.TxEncoder
	MevLane   auctionante.MEVLane
}

// NewAnteHandler returns an AnteHandler that checks and increments sequence
// numbers, checks signatures & account numbers, and deducts fees from the first
// signer.
func NewAnteHandler(options HandlerOptions) (sdk.AnteHandler, error) {
	if options.AccountKeeper == nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "account keeper is required for ante builder")
	}
	if options.BankKeeper == nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "bank keeper is required for ante builder")
	}
	if options.SignModeHandler == nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "sign mode handler is required for ante builder")
	}
	if options.EVMKeeper == nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "EVM keeper is required for ante builder")
	}
	if options.IBCkeeper == nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "IBC keeper is required for ante builder")
	}
	if options.OPChildKeeper == nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "OPChild keeper is required for ante builder")
	}
	if options.AuctionKeeper == nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "Auction keeper is required for ante builder")
	}

	sigGasConsumer := options.SigGasConsumer
	if sigGasConsumer == nil {
		sigGasConsumer = sigverify.DefaultSigVerificationGasConsumer
	}

	txFeeChecker := options.TxFeeChecker
	if txFeeChecker == nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "tx fee checker is required for ante builder")
	}

	anteDecorators := []sdk.AnteDecorator{
		accnum.NewAccountNumberDecorator(options.AccountKeeper),
		ante.NewSetUpContextDecorator(), // outermost AnteDecorator. SetUpContext must be called first
		ante.NewExtensionOptionsDecorator(options.ExtensionOptionChecker),
		NewEthTxDecorator(options.EVMKeeper),
		evmante.NewGasPricesDecorator(),
		ante.NewValidateBasicDecorator(),
		ante.NewTxTimeoutHeightDecorator(),
		ante.NewValidateMemoDecorator(options.AccountKeeper),
		NewConsumeTxSizeGasDecorator(options.AccountKeeper),
		NewGasFreeFeeDecorator(options.AccountKeeper, options.BankKeeper, options.FeegrantKeeper, options.EVMKeeper, txFeeChecker),
		// SetPubKeyDecorator must be called before all signature verification decorators
		ante.NewSetPubKeyDecorator(options.AccountKeeper),
		ante.NewValidateSigCountDecorator(options.AccountKeeper),
		NewSigGasConsumeDecorator(options.AccountKeeper, sigGasConsumer),
		NewSigVerificationDecorator(options.AccountKeeper, options.SignModeHandler),
		evmante.NewIncrementSequenceDecorator(options.AccountKeeper),
		ibcante.NewRedundantRelayDecorator(options.IBCkeeper),
		auctionante.NewAuctionDecorator(options.AuctionKeeper, options.TxEncoder, options.MevLane),
		opchildante.NewRedundantBridgeDecorator(options.OPChildKeeper),
	}

	return sdk.ChainAnteDecorators(anteDecorators...), nil
}

func CreateAnteHandlerForOPinit(ak ante.AccountKeeper, signModeHandler *txsigning.HandlerMap) sdk.AnteHandler {
	return sdk.ChainAnteDecorators(
		ante.NewValidateBasicDecorator(),
		ante.NewSetPubKeyDecorator(ak),
		ante.NewValidateSigCountDecorator(ak),
		NewSigGasConsumeDecorator(ak, sigverify.DefaultSigVerificationGasConsumer),
		NewSigVerificationDecorator(ak, signModeHandler),
		evmante.NewIncrementSequenceDecorator(ak),
	)
}
