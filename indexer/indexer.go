package indexer

import (
	"context"

	"cosmossdk.io/collections"
	corestoretypes "cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"cosmossdk.io/store/dbadapter"
	storetypes "cosmossdk.io/store/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"

	"github.com/ethereum/go-ethereum/common"
	coretypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/initia-labs/kvindexer/store"
	rpctypes "github.com/initia-labs/minievm/jsonrpc/types"
	evmkeeper "github.com/initia-labs/minievm/x/evm/keeper"
)

// EVMIndexer is an interface to interact with the EVM indexer.
type EVMIndexer interface {
	storetypes.ABCIListener

	// tx
	TxByHash(ctx context.Context, hash common.Hash) (*rpctypes.RPCTransaction, error)
	TxByBlockAndIndex(ctx context.Context, blockHeight uint64, index uint64) (*rpctypes.RPCTransaction, error)
	IterateBlockTxs(ctx context.Context, blockHeight uint64, cb func(tx *rpctypes.RPCTransaction) (bool, error)) error

	// tx receipt
	TxReceiptByHash(ctx context.Context, hash common.Hash) (*coretypes.Receipt, error)

	// block
	BlockHashToNumber(ctx context.Context, hash common.Hash) (uint64, error)
	BlockHeaderByHash(ctx context.Context, hash common.Hash) (*coretypes.Header, error)
	BlockHeaderByNumber(ctx context.Context, number uint64) (*coretypes.Header, error)
}

// EVMIndexerImpl implements EVMIndexer.
type EVMIndexerImpl struct {
	db       dbm.DB
	logger   log.Logger
	txConfig client.TxConfig
	appCodec codec.Codec

	evmKeeper *evmkeeper.Keeper
	store     *store.CacheStore

	schema                   collections.Schema
	TxMap                    collections.Map[[]byte, rpctypes.RPCTransaction]
	TxReceiptMap             collections.Map[[]byte, coretypes.Receipt]
	BlockHeaderMap           collections.Map[uint64, coretypes.Header]
	BlockAndIndexToTxHashMap collections.Map[collections.Pair[uint64, uint64], []byte]
	BlockHashToNumberMap     collections.Map[[]byte, uint64]
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

	// TODO make cache size configurable
	store := store.NewCacheStore(dbadapter.Store{DB: db}, 100)
	sb := collections.NewSchemaBuilderFromAccessor(
		func(ctx context.Context) corestoretypes.KVStore {
			return store
		},
	)

	indexer := &EVMIndexerImpl{
		db:        db,
		store:     store,
		logger:    logger,
		txConfig:  txConfig,
		appCodec:  appCodec,
		evmKeeper: evmKeeper,

		TxMap:                    collections.NewMap(sb, prefixTx, "tx", collections.BytesKey, CollJsonVal[rpctypes.RPCTransaction]()),
		TxReceiptMap:             collections.NewMap(sb, prefixTxReceipt, "tx_receipt", collections.BytesKey, CollJsonVal[coretypes.Receipt]()),
		BlockHeaderMap:           collections.NewMap(sb, prefixBlockHeader, "block_header", collections.Uint64Key, CollJsonVal[coretypes.Header]()),
		BlockAndIndexToTxHashMap: collections.NewMap(sb, prefixBlockAndIndexToTxHash, "block_and_index_to_tx_hash", collections.PairKeyCodec(collections.Uint64Key, collections.Uint64Key), collections.BytesValue),
		BlockHashToNumberMap:     collections.NewMap(sb, prefixBlockHashToNumber, "block_hash_to_number", collections.BytesKey, collections.Uint64Value),
	}

	schema, err := sb.Build()
	if err != nil {
		return nil, err
	}
	indexer.schema = schema

	return indexer, nil
}
