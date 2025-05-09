package main

import (
	"encoding/json"

	"cosmossdk.io/math"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/initia-labs/OPinit/contrib/launchtools"
	"github.com/initia-labs/OPinit/contrib/launchtools/steps"
	"github.com/initia-labs/initia/app"
	"github.com/initia-labs/initia/app/params"
	minitiaapp "github.com/initia-labs/minievm/app"
	"github.com/initia-labs/minievm/types"

	ophosttypes "github.com/initia-labs/OPinit/x/ophost/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultLaunchStepFactories is a list of default launch step factories.
var DefaultLaunchStepFactories = []launchtools.LauncherStepFuncFactory[*launchtools.Config]{
	steps.InitializeConfig,
	steps.InitializeRPCHelpers,

	// Initialize genesis
	steps.InitializeGenesis,

	// Add system keys to the keyring
	steps.InitializeKeyring,

	// Run the app
	steps.RunApp,

	// Establish IBC channels for fungible and NFT transfer
	steps.EstablishIBCChannelsWithNFTTransfer(func() (string, string, string) {
		return "nft-transfer", "nft-transfer", "ics721-1"
	}),

	// Create OP Bridge, using open channel states
	steps.InitializeOpBridge,

	// Set bridge info and update clients
	steps.SetBridgeInfo,

	// Get the L1 and L2 heights
	steps.GetL1Height,
	steps.GetL2Height,

	InitiateDepositStep, // Initiate deposit to L2 with 0 amount to create INIT denom first

	// Cleanup
	steps.StopApp,
}

func LaunchCommand(ac *appCreator, enc params.EncodingConfig, mbm module.BasicManager) *cobra.Command {
	return launchtools.LaunchCmd(
		ac,
		func(denom string) map[string]json.RawMessage {
			// default denom in OPinit is "umin"
			if denom == "umin" {
				// convert to "GAS"
				denom = types.BaseDenom
			}

			return minitiaapp.NewDefaultGenesisState(enc.Codec, mbm, denom)
		},
		DefaultLaunchStepFactories,
	)
}

func InitiateDepositStep(config *launchtools.Config) launchtools.LauncherStepFunc {
	return func(ctx launchtools.Launcher) error {
		ctx.Logger().Info("Initiating token deposit...")

		initiateTokenDepositMsg := ophosttypes.NewMsgInitiateTokenDeposit(
			config.SystemKeys.BridgeExecutor.L1Address,
			*ctx.GetBridgeId(),
			config.SystemKeys.BridgeExecutor.L2Address,
			sdk.NewCoin(app.BondDenom, math.NewInt(0)),
			nil,
		)

		ctx.Logger().Info("broadcasting tx to L1...",
			"from-address", config.SystemKeys.BridgeExecutor.L1Address,
		)

		// already validated in config.go
		gasPrices, _ := sdk.ParseDecCoins(config.L1Config.GasPrices)
		gasFees, _ := gasPrices.MulDec(math.LegacyNewDecFromInt(math.NewInt(200000))).TruncateDecimal()

		// send createOpBridgeMessage to host (L1)
		res, err := ctx.GetRPCHelperL1().BroadcastTxAndWait(
			config.SystemKeys.BridgeExecutor.L1Address,
			config.SystemKeys.BridgeExecutor.Mnemonic,
			200000,
			gasFees,
			initiateTokenDepositMsg,
		)
		if err != nil {
			return errors.Wrap(err, "failed to broadcast tx")
		}

		// if transaction failed, return error
		if res.TxResult.Code != 0 {
			ctx.Logger().Error("tx failed", "code", res.TxResult.Code, "log", res.TxResult.Log)
			return errors.Errorf("tx failed with code %d", res.TxResult.Code)
		}

		ctx.Logger().Info("token deposit initiated")
		return nil
	}
}
