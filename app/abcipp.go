package app

import (
	"cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/x/authz"
	clienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	opchildtypes "github.com/initia-labs/OPinit/x/opchild/types"

	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	sdkmempool "github.com/cosmos/cosmos-sdk/types/mempool"
	cosmosante "github.com/cosmos/cosmos-sdk/x/auth/ante"

	opchildante "github.com/initia-labs/OPinit/x/opchild/ante"
	"github.com/initia-labs/initia/abcipp"
	initiatx "github.com/initia-labs/initia/tx"

	appante "github.com/initia-labs/minievm/app/ante"
)

func (app *MinitiaApp) setupABCIPP(mempoolMaxTxs int, appOpts servertypes.AppOptions) (
	sdkmempool.Mempool,
	sdk.AnteHandler,
	sdk.PrepareProposalHandler,
	sdk.ProcessProposalHandler,
	abcipp.CheckTx,
	error,
) {

	freeFeeChecker := func(ctx sdk.Context, tx sdk.Tx) bool {
		feeTx, ok := tx.(sdk.FeeTx)
		if !ok {
			return false
		}

		whitelist, err := app.OPChildKeeper.FeeWhitelist(ctx)
		if err != nil {
			return false
		}

		payer, err := app.ac.BytesToString(feeTx.FeePayer())
		if err != nil {
			return false
		}

		var granter string
		if feeTx.FeeGranter() != nil {
			granter, err = app.ac.BytesToString(feeTx.FeeGranter())
			if err != nil {
				return false
			}
		}

		for _, addr := range whitelist {
			if addr == payer || addr == granter {
				return true
			}
		}

		return false
	}

	feeChecker := opchildante.NewMempoolFeeChecker(app.OPChildKeeper).CheckTxFeeWithMinGasPrices
	feeCheckerWrapper := func(ctx sdk.Context, tx sdk.Tx) (sdk.Coins, int64, error) {
		if !freeFeeChecker(ctx, tx) {
			return feeChecker(ctx, tx)
		}

		// return fee without fee check
		feeTx, ok := tx.(sdk.FeeTx)
		if !ok {
			return nil, 0, errors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
		}

		return feeTx.GetFee(), 1 /* FIFO */, nil
	}

	handlerOpts := appante.HandlerOptions{
		HandlerOptions: cosmosante.HandlerOptions{
			AccountKeeper:          app.AccountKeeper,
			BankKeeper:             app.BankKeeper,
			FeegrantKeeper:         app.FeeGrantKeeper,
			SignModeHandler:        app.txConfig.SignModeHandler(),
			ExtensionOptionChecker: initiatx.ExtensionOptionChecker,
			TxFeeChecker:           feeCheckerWrapper,
		},
		IBCkeeper:     app.IBCKeeper,
		Codec:         app.appCodec,
		OPChildKeeper: app.OPChildKeeper,
		EVMKeeper:     app.EVMKeeper,
	}

	fullHandler, err := appante.NewAnteHandler(handlerOpts)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	minimalHandler, err := appante.NewMinimalAnteHandler(handlerOpts)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	anteHandler := appante.NewDualAnteHandler(minimalHandler, fullHandler)
	abcippCfg := abcipp.GetConfig(appOpts)

	// system tier: oracle and ibc client update messages
	systemTierMatcher := func(_ sdk.Context, tx sdk.Tx) bool {
		for _, msg := range tx.GetMsgs() {
			switch msg := msg.(type) {
			case *clienttypes.MsgUpdateClient:
			case *opchildtypes.MsgUpdateOracle:
			case *opchildtypes.MsgRelayOracleData:
			case *authz.MsgExec:
				msgs, err := msg.GetMessages()
				if err != nil || len(msgs) != 1 {
					return false
				}
				switch msgs[0].(type) {
				case *opchildtypes.MsgUpdateOracle, *opchildtypes.MsgRelayOracleData:
				default:
					return false
				}
			default:
				return false
			}
		}
		return true
	}

	mempool := abcipp.NewPriorityMempool(
		abcipp.PriorityMempoolConfig{
			MaxTx:              mempoolMaxTxs,
			MaxQueuedPerSender: abcippCfg.MaxQueuedPerSender,
			MaxQueuedTotal:     abcippCfg.MaxQueuedTotal,
			QueuedGapTTL:       abcippCfg.QueuedGapTTL,
			AnteHandler:        fullHandler, // cleaning worker uses full handler
			Tiers: []abcipp.Tier{
				{Name: "system", Matcher: systemTierMatcher},
				{Name: "admin", Matcher: freeFeeChecker},
			},
		}, app.Logger(), app.txConfig.TxEncoder(), app.GetAccountKeeper(),
	)

	// start mempool cleaning worker
	mempool.StartCleaningWorker(app.BaseApp, abcipp.DefaultMempoolCleaningInterval)

	proposalHandler, err := abcipp.NewProposalHandler(
		app.Logger(),
		app.txConfig.TxDecoder(),
		app.txConfig.TxEncoder(),
		mempool,
		fullHandler, // proposal handler uses full handler
	)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	checkTxHandler, err := abcipp.NewCheckTxHandler(
		app.Logger(),
		app.BaseApp,
		mempool,
		app.txConfig.TxDecoder(),
		app.BaseApp.CheckTx,
		feeCheckerWrapper,
	)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	return mempool, anteHandler, proposalHandler.PrepareProposalHandler(), proposalHandler.ProcessProposalHandler(), checkTxHandler.CheckTx, nil
}
