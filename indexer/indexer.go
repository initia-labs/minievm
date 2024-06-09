package indexer

import (
	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"

	"github.com/ethereum/go-ethereum/common"
	coretypes "github.com/ethereum/go-ethereum/core/types"

	rpctypes "github.com/initia-labs/minievm/jsonrpc/types"
	evmkeeper "github.com/initia-labs/minievm/x/evm/keeper"
)

// EVMIndexer is an interface to interact with the EVM indexer.
type EVMIndexer interface {
	storetypes.ABCIListener

	// tx
	TxByHash(hash common.Hash) (*rpctypes.RPCTransaction, error)
	TxByBlockAndIndex(blockHeight uint64, index uint64) (*rpctypes.RPCTransaction, error)

	// block
	BlockHeaderByHash(hash common.Hash) (*coretypes.Header, error)
	BlockHeaderByNumber(number uint64) (*coretypes.Header, error)
}

// EVMIndexerImpl implements EVMIndexer.
type EVMIndexerImpl struct {
	db       dbm.DB
	logger   log.Logger
	txConfig client.TxConfig
	appCodec codec.Codec

	evmKeeper *evmkeeper.Keeper
}

func NewEVMIndexer(
	appOpts servertypes.AppOptions,
	appCodec codec.Codec,
	logger log.Logger,
	txConfig client.TxConfig,
	evmKeeper *evmkeeper.Keeper,
) (EVMIndexer, error) {
	dbDir, dbBackend := getDBConfig(appOpts)
	db, err := dbm.NewDB("eth_index", dbBackend, dbDir)
	if err != nil {
		return nil, err
	}

	return &EVMIndexerImpl{db, logger, txConfig, appCodec, evmKeeper}, nil
}
