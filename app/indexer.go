package app

import (
	storetypes "cosmossdk.io/store/types"

	dbm "github.com/cosmos/cosmos-db"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"

	// local imports
	evmindexer "github.com/initia-labs/minievm/indexer"

	// kvindexer
	kvindexer "github.com/initia-labs/kvindexer"
	kvindexerconfig "github.com/initia-labs/kvindexer/config"
	blocksubmodule "github.com/initia-labs/kvindexer/submodules/block"
	nft "github.com/initia-labs/kvindexer/submodules/evm-nft"
	tx "github.com/initia-labs/kvindexer/submodules/evm-tx"
	"github.com/initia-labs/kvindexer/submodules/pair"
	kvindexermodule "github.com/initia-labs/kvindexer/x/kvindexer"
	kvindexerkeeper "github.com/initia-labs/kvindexer/x/kvindexer/keeper"
)

func setupIndexer(
	app *MinitiaApp,
	appOpts servertypes.AppOptions,
	indexerDB, kvindexerDB dbm.DB,
) (evmindexer.EVMIndexer, *kvindexerkeeper.Keeper, *kvindexermodule.AppModuleBasic, *storetypes.StreamingManager, error) {
	// Initialize EVM Indexer
	evmIndexer, err := evmindexer.NewEVMIndexer(indexerDB, app.appCodec, app.Logger(), app.txConfig, app.EVMKeeper)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	// Configure kvindexer
	kvindexerConfig, err := kvindexerconfig.NewConfig(appOpts)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	// Initialize kvIndexerKeeper
	kvIndexerKeeper := kvindexerkeeper.NewKeeper(
		app.appCodec,
		"evm",
		kvindexerDB,
		kvindexerConfig,
		app.ac,
		app.vc,
	)

	// Setup submodules
	subModules := []interface {
		Setup() error
	}{
		blocksubmodule.NewBlockSubmodule(app.appCodec, kvIndexerKeeper, app.OPChildKeeper),
		tx.NewTxSubmodule(app.appCodec, kvIndexerKeeper),
		pair.NewPairSubmodule(app.appCodec, kvIndexerKeeper, app.IBCKeeper.ChannelKeeper, app.TransferKeeper),
		nft.NewEvmNFTSubmodule(app.ac, app.appCodec, kvIndexerKeeper, app.EVMKeeper, nil),
	}

	for _, subModule := range subModules {
		if err := subModule.Setup(); err != nil {
			return nil, nil, nil, nil, err
		}
	}

	// Register submodules with kvIndexerKeeper
	if err := kvIndexerKeeper.RegisterSubmodules(subModules...); err != nil {
		return nil, nil, nil, nil, err
	}

	// Create kvIndexer
	kvIndexer, err := kvindexer.NewIndexer(app.GetBaseApp().Logger(), kvIndexerKeeper)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	if err := kvIndexer.Validate(); err != nil {
		return nil, nil, nil, nil, err
	}

	if err := kvIndexer.Prepare(nil); err != nil {
		return nil, nil, nil, nil, err
	}

	if err := kvIndexerKeeper.Seal(); err != nil {
		return nil, nil, nil, nil, err
	}

	if err := kvIndexer.Start(nil); err != nil {
		return nil, nil, nil, nil, err
	}

	listeners := []storetypes.ABCIListener{evmIndexer, kvIndexer}

	kvIndexerModule := kvindexermodule.NewAppModuleBasic(kvIndexerKeeper)

	streamingManager := &storetypes.StreamingManager{
		ABCIListeners: listeners,
		StopNodeOnErr: true,
	}

	return evmIndexer, kvIndexerKeeper, &kvIndexerModule, streamingManager, nil
}
