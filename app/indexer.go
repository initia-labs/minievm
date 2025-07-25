package app

import (
	storetypes "cosmossdk.io/store/types"

	dbm "github.com/cosmos/cosmos-db"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"

	// local imports
	"github.com/initia-labs/initia/app/params"
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
	encodingConfig params.EncodingConfig,
	indexerDB, kvindexerDB dbm.DB,
) (evmindexer.EVMIndexer, *kvindexerkeeper.Keeper, *kvindexermodule.AppModuleBasic, *storetypes.StreamingManager, error) {
	// setup evm indexer
	evmIndexer, err := evmindexer.NewEVMIndexer(indexerDB, app.ac, app.appCodec, app.Logger(), app.txConfig, app.EVMKeeper)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	// initialize the indexer keeper
	kvindexerConfig, err := kvindexerconfig.NewConfig(appOpts)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	kvIndexerKeeper := kvindexerkeeper.NewKeeper(
		app.appCodec,
		"evm",
		kvindexerDB,
		kvindexerConfig,
		app.ac,
		app.vc,
	)

	smBlock, err := blocksubmodule.NewBlockSubmodule(encodingConfig, kvIndexerKeeper, app.OPChildKeeper)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	smTx, err := tx.NewTxSubmodule(encodingConfig, kvIndexerKeeper)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	smPair, err := pair.NewPairSubmodule(encodingConfig, kvIndexerKeeper, app.IBCKeeper.ChannelKeeper, app.TransferKeeper)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	smNft, err := nft.NewEvmNFTSubmodule(app.ac, encodingConfig, kvIndexerKeeper, app.EVMKeeper, smPair)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	// order of registration matters: smPair must be registered before smNft since smNft depends on smPair
	err = kvIndexerKeeper.RegisterSubmodules(smBlock, smTx, smPair, smNft)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	// Add your implementation here

	kvIndexer, err := kvindexer.NewIndexer(app.GetBaseApp().Logger(), kvIndexerKeeper)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	listeners := []storetypes.ABCIListener{evmIndexer}

	var kvIndexerModule *kvindexermodule.AppModuleBasic
	if kvIndexer != nil {
		if err = kvIndexer.Validate(); err != nil {
			return nil, nil, nil, nil, err
		}

		if err = kvIndexer.Prepare(nil); err != nil {
			return nil, nil, nil, nil, err
		}

		if err = kvIndexerKeeper.Seal(); err != nil {
			return nil, nil, nil, nil, err
		}

		if err = kvIndexer.Start(nil); err != nil {
			return nil, nil, nil, nil, err
		}

		listeners = append(listeners, kvIndexer)

		// set kvindexer module
		m := kvindexermodule.NewAppModuleBasic(kvIndexerKeeper)
		kvIndexerModule = &m
	}

	streamingManager := storetypes.StreamingManager{
		ABCIListeners: listeners,
		StopNodeOnErr: true,
	}

	return evmIndexer, kvIndexerKeeper, kvIndexerModule, &streamingManager, nil
}
