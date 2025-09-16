package main

import (
	"context"
	"errors"
	"io"
	"os"
	"path"
	"path/filepath"

	tmcli "github.com/cometbft/cometbft/libs/cli"
	"golang.org/x/sync/errgroup"

	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"cosmossdk.io/log"
	confixcmd "cosmossdk.io/tools/confix/cmd"

	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/config"
	"github.com/cosmos/cosmos-sdk/client/debug"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/pruning"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/client/snapshot"
	"github.com/cosmos/cosmos-sdk/server"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authcmd "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"

	minitiaapp "github.com/initia-labs/minievm/app"
	"github.com/initia-labs/minievm/jsonrpc"
	jsonrpcconfig "github.com/initia-labs/minievm/jsonrpc/config"
	evmconfig "github.com/initia-labs/minievm/x/evm/config"

	"github.com/initia-labs/initia/app/params"
	initiakeyring "github.com/initia-labs/initia/crypto/keyring"

	opchildcli "github.com/initia-labs/OPinit/x/opchild/client/cli"

	cmtcmd "github.com/cometbft/cometbft/cmd/cometbft/commands"
)

// NewRootCmd creates a new root command for initiad. It is called once in the
// main function.
func NewRootCmd() (*cobra.Command, params.EncodingConfig) {
	sdkConfig := sdk.GetConfig()
	sdkConfig.SetCoinType(minitiaapp.CoinType)

	accountPubKeyPrefix := minitiaapp.AccountAddressPrefix + "pub"
	validatorAddressPrefix := minitiaapp.AccountAddressPrefix + "valoper"
	validatorPubKeyPrefix := minitiaapp.AccountAddressPrefix + "valoperpub"
	consNodeAddressPrefix := minitiaapp.AccountAddressPrefix + "valcons"
	consNodePubKeyPrefix := minitiaapp.AccountAddressPrefix + "valconspub"

	sdkConfig.SetBech32PrefixForAccount(minitiaapp.AccountAddressPrefix, accountPubKeyPrefix)
	sdkConfig.SetBech32PrefixForValidator(validatorAddressPrefix, validatorPubKeyPrefix)
	sdkConfig.SetBech32PrefixForConsensusNode(consNodeAddressPrefix, consNodePubKeyPrefix)
	sdkConfig.SetAddressVerifier(minitiaapp.VerifyAddressLen())

	// seal moved to post setup
	// sdkConfig.Seal()

	encodingConfig := minitiaapp.MakeEncodingConfig()
	basicManager := minitiaapp.BasicManager()

	// Get the executable name and configure the viper instance so that environmental
	// variables are checked based off that name. The underscore character is used
	// as a separator
	executableName, err := os.Executable()
	if err != nil {
		panic(err)
	}

	basename := path.Base(executableName)

	// Configure the viper instance
	initClientCtx := client.Context{}.
		WithCodec(encodingConfig.Codec).
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithInput(os.Stdin).
		WithAccountRetriever(types.AccountRetriever{}).
		WithHomeDir(minitiaapp.DefaultNodeHome).
		WithViper(minitiaapp.EnvPrefix).
		WithKeyringOptions(initiakeyring.EthSecp256k1Option())

	rootCmd := &cobra.Command{
		Use:   basename,
		Short: "minitia App",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			// except for launch command, seal the config
			if cmd.Name() != "launch" {
				sdk.GetConfig().Seal()
			}

			// set the default command outputs
			cmd.SetOut(cmd.OutOrStdout())
			cmd.SetErr(cmd.ErrOrStderr())

			// read envs before reading persistent flags
			// TODO - should we handle this for tx flags & query flags?
			initClientCtx, err := readEnv(initClientCtx)
			if err != nil {
				return err
			}

			// read persistent flags if they changed, and override the env configs.
			initClientCtx, err = client.ReadPersistentCommandFlags(initClientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			// unsafe-reset-all is not working without viper set
			viper.Set(tmcli.HomeFlag, initClientCtx.HomeDir)

			initClientCtx, err = config.ReadFromClientConfig(initClientCtx)
			if err != nil {
				return err
			}

			// override the keyring if it's set
			if initClientCtx.Keyring != nil {
				kr, err := initiakeyring.NewKeyring(initClientCtx, initClientCtx.Keyring.Backend())
				if err != nil {
					return err
				}

				initClientCtx = initClientCtx.WithKeyring(kr)
			}

			if err := client.SetCmdClientContext(cmd, initClientCtx); err != nil {
				return err
			}

			initiaappTemplate, initiaappConfig := initAppConfig()
			customTMConfig := initTendermintConfig()

			return server.InterceptConfigsPreRunHandler(cmd, initiaappTemplate, initiaappConfig, customTMConfig)
		},
	}

	initRootCmd(rootCmd, encodingConfig, basicManager)

	// add keyring to autocli opts
	autoCliOpts := minitiaapp.AutoCliOpts()
	initClientCtx, _ = config.ReadFromClientConfig(initClientCtx)
	autoCliOpts.ClientCtx = initClientCtx

	if err := autoCliOpts.EnhanceRootCommand(rootCmd); err != nil {
		panic(err)
	}

	return rootCmd, encodingConfig
}

func initRootCmd(rootCmd *cobra.Command, encodingConfig params.EncodingConfig, basicManager module.BasicManager) {
	a := &appCreator{}
	// you can get app from a.app in post setup handler

	rootCmd.AddCommand(
		InitCmd(basicManager, minitiaapp.DefaultNodeHome),
		debug.Cmd(),
		confixcmd.ConfigCommand(),
		pruning.Cmd(a.AppCreator(), minitiaapp.DefaultNodeHome),
		snapshot.Cmd(a.AppCreator()),
	)

	server.AddCommandsWithStartCmdOptions(rootCmd, minitiaapp.DefaultNodeHome, a.AppCreator(), a.appExport, server.StartCmdOptions{
		AddFlags: addModuleInitFlags,
		PostSetup: func(svrCtx *server.Context, clientCtx client.Context, ctx context.Context, g *errgroup.Group) error {
			// start jsonrpc server
			if err := jsonrpc.StartJSONRPC(
				ctx, g, a.App().(*minitiaapp.MinitiaApp),
				svrCtx,
				clientCtx,
				jsonrpcconfig.GetConfig(a.appOpts),
				false,
			); err != nil {
				return err
			}

			// start jsonrpc websocket server
			if err := jsonrpc.StartJSONRPC(
				ctx, g, a.App().(*minitiaapp.MinitiaApp),
				svrCtx,
				clientCtx,
				jsonrpcconfig.GetConfig(a.appOpts),
				true,
			); err != nil {
				return err
			}

			// set client context for the evm Indexer
			if err := a.App().(*minitiaapp.MinitiaApp).InitializeIndexer(clientCtx); err != nil {
				return err
			}

			return nil
		},
	})

	// add keybase, auxiliary RPC, query, and tx child commands
	rootCmd.AddCommand(
		server.StatusCommand(),
		genesisCommand(encodingConfig, basicManager),
		queryCommand(),
		txCommand(),
		initiakeyring.OverrideDefaultKeyType(keys.Commands()),
	)

	// add launch commands
	rootCmd.AddCommand(LaunchCommand(a, encodingConfig, basicManager))
	rootCmd.AddCommand(NewMultipleRollbackCmd(a.AppCreator()))
	rootCmd.AddCommand(cmtcmd.FetchGenesisCmd)
}

func addModuleInitFlags(startCmd *cobra.Command) {
	crisis.AddModuleInitFlags(startCmd)
}

func genesisCommand(encodingConfig params.EncodingConfig, basicManager module.BasicManager) *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "genesis",
		Short:                      "Application's genesis-related subcommands",
		DisableFlagParsing:         false,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	ac := encodingConfig.TxConfig.SigningContext().AddressCodec()

	cmd.AddCommand(
		genutilcli.AddGenesisAccountCmd(minitiaapp.DefaultNodeHome, ac),
		opchildcli.AddGenesisValidatorCmd(basicManager, encodingConfig.TxConfig, banktypes.GenesisBalancesIterator{}, minitiaapp.DefaultNodeHome),
		opchildcli.AddFeeWhitelistCmd(minitiaapp.DefaultNodeHome, ac),
		genutilcli.ValidateGenesisCmd(basicManager),
		genutilcli.GenTxCmd(basicManager, encodingConfig.TxConfig, banktypes.GenesisBalancesIterator{}, minitiaapp.DefaultNodeHome, ac),
	)

	return cmd
}

func queryCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "query",
		Aliases:                    []string{"q"},
		Short:                      "Querying subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		rpc.QueryEventForTxCmd(),
		server.QueryBlockCmd(),
		authcmd.QueryTxsByEventsCmd(),
		server.QueryBlocksCmd(),
		authcmd.QueryTxCmd(),
		server.QueryBlockResultsCmd(),
	)

	return cmd
}

func txCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "tx",
		Short:                      "Transactions subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		authcmd.GetSignCommand(),
		authcmd.GetSignBatchCommand(),
		authcmd.GetMultiSignCommand(),
		authcmd.GetMultiSignBatchCmd(),
		authcmd.GetValidateSignaturesCommand(),
		authcmd.GetBroadcastCommand(),
		authcmd.GetEncodeCommand(),
		authcmd.GetDecodeCommand(),
		authcmd.GetSimulateCmd(),
	)

	return cmd
}

type appCreator struct {
	app     servertypes.Application
	appOpts servertypes.AppOptions
}

func (a *appCreator) AppCreator() servertypes.AppCreator {
	return func(logger log.Logger, db dbm.DB, traceStore io.Writer, appOpts servertypes.AppOptions) servertypes.Application {
		baseappOptions := server.DefaultBaseappOptions(appOpts)

		// create EVM indexer db
		dbDir, dbBackend := getDBConfig(appOpts)
		indexerDB, err := dbm.NewDB("eth_index", dbBackend, dbDir)
		if err != nil {
			panic(err)
		}
		evmConfig := evmconfig.GetConfig(appOpts)
		if err := evmConfig.Validate(); err != nil {
			panic(err)
		}
		app := minitiaapp.NewMinitiaApp(
			logger, db, indexerDB, traceStore, true,
			evmConfig, appOpts, baseappOptions...,
		)

		// store app in creator
		a.app = app
		a.appOpts = appOpts

		return app
	}
}

func (a *appCreator) App() servertypes.Application {
	return a.app
}

func (a appCreator) appExport(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	height int64,
	forZeroHeight bool,
	_ []string,
	appOpts servertypes.AppOptions,
	modulesToExport []string,
) (servertypes.ExportedApp, error) {
	homePath, ok := appOpts.Get(flags.FlagHome).(string)
	if !ok || homePath == "" {
		return servertypes.ExportedApp{}, errors.New("application home not set")
	}

	var initiaApp *minitiaapp.MinitiaApp
	if height != -1 {
		initiaApp = minitiaapp.NewMinitiaApp(logger, db, dbm.NewMemDB(), traceStore, false, evmconfig.DefaultEVMConfig(), appOpts)

		if err := initiaApp.LoadHeight(height); err != nil {
			return servertypes.ExportedApp{}, err
		}
	} else {
		initiaApp = minitiaapp.NewMinitiaApp(logger, db, dbm.NewMemDB(), traceStore, true, evmconfig.DefaultEVMConfig(), appOpts)
	}

	return initiaApp.ExportAppStateAndValidators(forZeroHeight, modulesToExport)
}

func readEnv(clientCtx client.Context) (client.Context, error) {
	if outputFormat := clientCtx.Viper.GetString(tmcli.OutputFlag); outputFormat != "" {
		clientCtx = clientCtx.WithOutputFormat(outputFormat)
	}

	if homeDir := clientCtx.Viper.GetString(flags.FlagHome); homeDir != "" {
		clientCtx = clientCtx.WithHomeDir(homeDir)
	}

	if clientCtx.Viper.GetBool(flags.FlagDryRun) {
		clientCtx = clientCtx.WithSimulation(true)
	}

	if keyringDir := clientCtx.Viper.GetString(flags.FlagKeyringDir); keyringDir != "" {
		clientCtx = clientCtx.WithKeyringDir(clientCtx.Viper.GetString(flags.FlagKeyringDir))
	}

	if chainID := clientCtx.Viper.GetString(flags.FlagChainID); chainID != "" {
		clientCtx = clientCtx.WithChainID(chainID)
	}

	if keyringBackend := clientCtx.Viper.GetString(flags.FlagKeyringBackend); keyringBackend != "" {
		kr, err := client.NewKeyringFromBackend(clientCtx, keyringBackend)
		if err != nil {
			return clientCtx, err
		}

		clientCtx = clientCtx.WithKeyring(kr)
	}

	if nodeURI := clientCtx.Viper.GetString(flags.FlagNode); nodeURI != "" {
		clientCtx = clientCtx.WithNodeURI(nodeURI)

		client, err := client.NewClientFromNode(nodeURI)
		if err != nil {
			return clientCtx, err
		}

		clientCtx = clientCtx.WithClient(client)
	}

	return clientCtx, nil
}

// getDBConfig returns the database configuration for the EVM indexer
func getDBConfig(appOpts servertypes.AppOptions) (string, dbm.BackendType) {
	rootDir := cast.ToString(appOpts.Get("home"))
	dbDir := cast.ToString(appOpts.Get("db_dir"))
	dbBackend := server.GetAppDBBackend(appOpts)

	return rootify(dbDir, rootDir), dbBackend
}

// helper function to make config creation independent of root dir
func rootify(path, root string) string {
	if filepath.IsAbs(path) {
		return path
	}
	return filepath.Join(root, path)
}
