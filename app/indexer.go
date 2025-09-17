package app

import (
	storetypes "cosmossdk.io/store/types"

	dbm "github.com/cosmos/cosmos-db"

	// local imports

	evmindexer "github.com/initia-labs/minievm/indexer"
)

func setupIndexer(
	app *MinitiaApp,
	indexerDB dbm.DB,
) (evmindexer.EVMIndexer, *storetypes.StreamingManager, error) {
	evmIndexer, err := evmindexer.NewEVMIndexer(indexerDB, app.ac, app.appCodec, app.Logger(), app.txConfig, app.EVMKeeper)
	if err != nil {
		return nil, nil, err
	}

	listeners := []storetypes.ABCIListener{evmIndexer}
	streamingManager := storetypes.StreamingManager{
		ABCIListeners: listeners,
		StopNodeOnErr: true,
	}

	return evmIndexer, &streamingManager, nil
}
