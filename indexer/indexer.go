package indexer

import (
	"context"
	"time"

	"github.com/jellydator/ttlcache/v3"

	"cosmossdk.io/collections"
	corestoretypes "cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"cosmossdk.io/store/dbadapter"
	storetypes "cosmossdk.io/store/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/mempool"

	"github.com/ethereum/go-ethereum/common"
	coretypes "github.com/ethereum/go-ethereum/core/types"

	rpctypes "github.com/initia-labs/minievm/jsonrpc/types"
	evmconfig "github.com/initia-labs/minievm/x/evm/config"
	evmkeeper "github.com/initia-labs/minievm/x/evm/keeper"
)

// EVMIndexer is an interface to interact with the EVM indexer.
type EVMIndexer interface {
	storetypes.ABCIListener

	// tx
	TxByHash(ctx context.Context, hash common.Hash) (*rpctypes.RPCTransaction, error)
	IterateBlockTxs(ctx context.Context, blockHeight uint64, cb func(tx *rpctypes.RPCTransaction) (bool, error)) error
	TxHashByBlockAndIndex(ctx context.Context, blockHeight uint64, index uint64) (common.Hash, error)

	// tx receipt
	TxReceiptByHash(ctx context.Context, hash common.Hash) (*coretypes.Receipt, error)
	IterateBlockTxReceipts(ctx context.Context, blockHeight uint64, cb func(tx *coretypes.Receipt) (bool, error)) error

	// block
	BlockHashToNumber(ctx context.Context, hash common.Hash) (uint64, error)
	BlockHeaderByNumber(ctx context.Context, number uint64) (*coretypes.Header, error)

	// cosmos tx hash
	CosmosTxHashByTxHash(ctx context.Context, hash common.Hash) ([]byte, error)
	TxHashByCosmosTxHash(ctx context.Context, hash []byte) (common.Hash, error)

	// event subscription
	Subscribe() (chan *coretypes.Header, chan []*coretypes.Log, chan *rpctypes.RPCTransaction)

	// mempool
	MempoolWrapper(mempool mempool.Mempool) mempool.Mempool
	TxInMempool(hash common.Hash) *rpctypes.RPCTransaction

	// Stop
	Stop()
}

// EVMIndexerImpl implements EVMIndexer.
type EVMIndexerImpl struct {
	db       dbm.DB
	logger   log.Logger
	txConfig client.TxConfig
	appCodec codec.Codec

	store     *CacheStore
	evmKeeper *evmkeeper.Keeper

	schema                   collections.Schema
	TxMap                    collections.Map[[]byte, rpctypes.RPCTransaction]
	TxReceiptMap             collections.Map[[]byte, coretypes.Receipt]
	BlockHeaderMap           collections.Map[uint64, coretypes.Header]
	BlockAndIndexToTxHashMap collections.Map[collections.Pair[uint64, uint64], []byte]
	BlockHashToNumberMap     collections.Map[[]byte, uint64]
	TxHashToCosmosTxHash     collections.Map[[]byte, []byte]
	CosmosTxHashToTxHash     collections.Map[[]byte, []byte]

	blockChans   []chan *coretypes.Header
	logsChans    []chan []*coretypes.Log
	pendingChans []chan *rpctypes.RPCTransaction

	// txPendingMap is a map to store tx hashes in pending state.
	txPendingMap *ttlcache.Cache[common.Hash, *rpctypes.RPCTransaction]
}

func NewEVMIndexer(
	db dbm.DB,
	appCodec codec.Codec,
	logger log.Logger,
	txConfig client.TxConfig,
	evmKeeper *evmkeeper.Keeper,
) (EVMIndexer, error) {
	cfg := evmKeeper.Config()
	if cfg.IndexerCacheSize == 0 {
		cfg.IndexerCacheSize = evmconfig.DefaultIndexerCacheSize
	}

	store := NewCacheStore(dbadapter.Store{DB: db}, cfg.IndexerCacheSize)
	sb := collections.NewSchemaBuilderFromAccessor(
		func(ctx context.Context) corestoretypes.KVStore {
			return store
		},
	)

	indexer := &EVMIndexerImpl{
		db:       db,
		store:    store,
		logger:   logger,
		txConfig: txConfig,
		appCodec: appCodec,

		evmKeeper: evmKeeper,

		TxMap:                    collections.NewMap(sb, prefixTx, "tx", collections.BytesKey, CollJsonVal[rpctypes.RPCTransaction]()),
		TxReceiptMap:             collections.NewMap(sb, prefixTxReceipt, "tx_receipt", collections.BytesKey, CollJsonVal[coretypes.Receipt]()),
		BlockHeaderMap:           collections.NewMap(sb, prefixBlockHeader, "block_header", collections.Uint64Key, CollJsonVal[coretypes.Header]()),
		BlockAndIndexToTxHashMap: collections.NewMap(sb, prefixBlockAndIndexToTxHash, "block_and_index_to_tx_hash", collections.PairKeyCodec(collections.Uint64Key, collections.Uint64Key), collections.BytesValue),
		BlockHashToNumberMap:     collections.NewMap(sb, prefixBlockHashToNumber, "block_hash_to_number", collections.BytesKey, collections.Uint64Value),
		TxHashToCosmosTxHash:     collections.NewMap(sb, prefixTxHashToCosmosTxHash, "tx_hash_to_cosmos_tx_hash", collections.BytesKey, collections.BytesValue),
		CosmosTxHashToTxHash:     collections.NewMap(sb, prefixCosmosTxHashToTxHash, "cosmos_tx_hash_to_tx_hash", collections.BytesKey, collections.BytesValue),

		blockChans:   nil,
		logsChans:    nil,
		pendingChans: nil,

		// Use ttlcache to cope with abnormal cases like tx not included in a block
		txPendingMap: ttlcache.New(
			// pending tx lifetime is 1 minute in indexer
			ttlcache.WithTTL[common.Hash, *rpctypes.RPCTransaction](time.Minute),
		),
	}

	schema, err := sb.Build()
	if err != nil {
		return nil, err
	}
	indexer.schema = schema

	// expire pending tx
	go indexer.txPendingMap.Start()

	return indexer, nil
}

// Subscribe returns channels to receive blocks and logs.
func (e *EVMIndexerImpl) Subscribe() (chan *coretypes.Header, chan []*coretypes.Log, chan *rpctypes.RPCTransaction) {
	blockChan := make(chan *coretypes.Header)
	logsChan := make(chan []*coretypes.Log)
	pendingChan := make(chan *rpctypes.RPCTransaction)

	e.blockChans = append(e.blockChans, blockChan)
	e.logsChans = append(e.logsChans, logsChan)
	e.pendingChans = append(e.pendingChans, pendingChan)
	return blockChan, logsChan, pendingChan
}

// blockEvents is a struct to emit block events.
type blockEvents struct {
	header *coretypes.Header
	logs   [][]*coretypes.Log
}

// blockEventsEmitter emits block events to subscribers.
func (e *EVMIndexerImpl) blockEventsEmitter(blockEvents *blockEvents, done chan struct{}) {
	defer close(done)
	if blockEvents == nil {
		return
	}

	// emit logs first; use unbuffered channel to ensure logs are emitted before block header
	for _, logs := range blockEvents.logs {
		for _, logsChan := range e.logsChans {
			logsChan <- logs
		}
	}

	// emit block header
	for _, blockChan := range e.blockChans {
		blockChan <- blockEvents.header
	}
}

// Stop stops the indexer.
func (e *EVMIndexerImpl) Stop() {
	if e.txPendingMap != nil {
		e.txPendingMap.Stop()
	}
}
