package app

import (
	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/mempool"
	cosmosante "github.com/cosmos/cosmos-sdk/x/auth/ante"

	// OPinit imports
	opchildante "github.com/initia-labs/OPinit/x/opchild/ante"
	opchildlanes "github.com/initia-labs/OPinit/x/opchild/lanes"

	// initia imports
	initialanes "github.com/initia-labs/initia/app/lanes"

	// skip-mev imports
	blockabci "github.com/skip-mev/block-sdk/v2/abci"
	blockchecktx "github.com/skip-mev/block-sdk/v2/abci/checktx"
	signer_extraction "github.com/skip-mev/block-sdk/v2/adapters/signer_extraction_adapter"
	"github.com/skip-mev/block-sdk/v2/block"
	blockbase "github.com/skip-mev/block-sdk/v2/block/base"
	mevlane "github.com/skip-mev/block-sdk/v2/lanes/mev"

	// local imports
	appante "github.com/initia-labs/minievm/app/ante"
	"github.com/initia-labs/minievm/app/checktx"
)

func setupBlockSDK(
	app *MinitiaApp,
	mempoolMaxTxs int,
) (
	mempool.Mempool,
	sdk.AnteHandler,
	blockchecktx.CheckTx,
	sdk.PrepareProposalHandler,
	sdk.ProcessProposalHandler,
	error,
) {

	// initialize and set the InitiaApp mempool. The current mempool will be the
	// x/auction module's mempool which will extract the top bid from the current block's auction
	// and insert the txs at the top of the block spots.
	signerExtractor := signer_extraction.NewDefaultAdapter()

	systemLane := initialanes.NewSystemLane(blockbase.LaneConfig{
		Logger:          app.Logger(),
		TxEncoder:       app.txConfig.TxEncoder(),
		TxDecoder:       app.txConfig.TxDecoder(),
		MaxBlockSpace:   math.LegacyMustNewDecFromStr("0.1"),
		MaxTxs:          1,
		SignerExtractor: signerExtractor,
	}, opchildlanes.SystemLaneMatchHandler())

	factory := mevlane.NewDefaultAuctionFactory(app.txConfig.TxDecoder(), signerExtractor)
	mevLane := initialanes.NewMEVLane(blockbase.LaneConfig{
		Logger:          app.Logger(),
		TxEncoder:       app.txConfig.TxEncoder(),
		TxDecoder:       app.txConfig.TxDecoder(),
		MaxBlockSpace:   math.LegacyMustNewDecFromStr("0.1"),
		MaxTxs:          100,
		SignerExtractor: signerExtractor,
	}, factory, factory.MatchHandler())

	freeLane := initialanes.NewFreeLane(blockbase.LaneConfig{
		Logger:          app.Logger(),
		TxEncoder:       app.txConfig.TxEncoder(),
		TxDecoder:       app.txConfig.TxDecoder(),
		MaxBlockSpace:   math.LegacyMustNewDecFromStr("0.1"),
		MaxTxs:          100,
		SignerExtractor: signerExtractor,
	}, opchildlanes.NewFreeLaneMatchHandler(app.ac, app.OPChildKeeper).MatchHandler())

	defaultLane := initialanes.NewDefaultLane(blockbase.LaneConfig{
		Logger:          app.Logger(),
		TxEncoder:       app.txConfig.TxEncoder(),
		TxDecoder:       app.txConfig.TxDecoder(),
		MaxBlockSpace:   math.LegacyMustNewDecFromStr("0.7"),
		MaxTxs:          mempoolMaxTxs,
		SignerExtractor: signerExtractor,
	})

	lanes := []block.Lane{systemLane, mevLane, freeLane, defaultLane}
	mempool, err := block.NewLanedMempool(app.Logger(), lanes)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	// create a custom fee checker for free lane
	txFeeChecker := opchildante.NewMempoolFeeChecker(app.OPChildKeeper).CheckTxFeeWithMinGasPrices
	freeLaneFeeChecker := func(ctx sdk.Context, tx sdk.Tx) (sdk.Coins, int64, error) {
		if !freeLane.Match(ctx, tx) {
			return txFeeChecker(ctx, tx)
		}

		// return fee without fee check
		feeTx, ok := tx.(sdk.FeeTx)
		if !ok {
			return nil, 0, errorsmod.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
		}

		return feeTx.GetFee(), 1 /* FIFO */, nil
	}

	anteHandler, err := appante.NewAnteHandler(
		appante.HandlerOptions{
			HandlerOptions: cosmosante.HandlerOptions{
				AccountKeeper:   app.AccountKeeper,
				BankKeeper:      app.BankKeeper,
				FeegrantKeeper:  app.FeeGrantKeeper,
				SignModeHandler: app.txConfig.SignModeHandler(),
				TxFeeChecker:    freeLaneFeeChecker,
			},
			IBCkeeper:     app.IBCKeeper,
			Codec:         app.appCodec,
			OPChildKeeper: app.OPChildKeeper,
			TxEncoder:     app.txConfig.TxEncoder(),
			AuctionKeeper: app.AuctionKeeper,
			MevLane:       mevLane,
			EVMKeeper:     app.EVMKeeper,
		},
	)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	// set ante handler to lanes
	opt := []blockbase.LaneOption{
		blockbase.WithAnteHandler(anteHandler),
	}
	systemLane.(*blockbase.BaseLane).WithOptions(
		opt...,
	)
	mevLane.WithOptions(
		opt...,
	)
	freeLane.(*blockbase.BaseLane).WithOptions(
		opt...,
	)
	defaultLane.(*blockbase.BaseLane).WithOptions(
		opt...,
	)

	mevCheckTx := blockchecktx.NewMEVCheckTxHandler(
		app.BaseApp,
		app.txConfig.TxDecoder(),
		mevLane,
		anteHandler,
		app.BaseApp.CheckTx,
	)
	checkTxHandler := blockchecktx.NewMempoolParityCheckTx(
		app.Logger(),
		mempool,
		app.txConfig.TxDecoder(),
		mevCheckTx.CheckTx(),
		app.BaseApp,
	)
	checkTx := checkTxHandler.CheckTx()

	proposalHandler := blockabci.New(
		app.Logger(),
		app.txConfig.TxDecoder(),
		app.txConfig.TxEncoder(),
		mempool,
		true,
	)

	prepareProposalHandler := proposalHandler.PrepareProposalHandler()
	processProposalHandler := proposalHandler.ProcessProposalHandler()

	// wrap checkTx for unordered txs
	app.checkTxWrapper = checktx.NewCheckTxWrapper(
		app.Logger(),
		app.txConfig,
		app,
		app.AccountKeeper,
		app.BankKeeper,
		app.EVMKeeper,
		app.EVMIndexer(),
		app.EVMKeeper.TxUtils(),
		checkTx,
		freeLaneFeeChecker,
	)

	return mempool, anteHandler, app.checkTxWrapper.CheckTx(), prepareProposalHandler, processProposalHandler, nil
}
