package indexer

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/jellydator/ttlcache/v3"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	corestoretypes "cosmossdk.io/core/store"
	"cosmossdk.io/log"
	snapshot "cosmossdk.io/store/snapshots/types"
	storetypes "cosmossdk.io/store/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/ethereum/go-ethereum/common"
	coretypes "github.com/ethereum/go-ethereum/core/types"

	rpctypes "github.com/initia-labs/minievm/jsonrpc/types"
	evmkeeper "github.com/initia-labs/minievm/x/evm/keeper"
)

// EVMIndexer is an interface to interact with the EVM indexer.
type EVMIndexer interface {
	snapshot.ExtensionSnapshotter
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

	// last indexed height
	GetLastIndexedHeight(ctx context.Context) (uint64, error)

	// mempool
	TxInPending(hash common.Hash) *rpctypes.RPCTransaction
	TxInQueued(hash common.Hash) *rpctypes.RPCTransaction
	PushPendingTx(tx *coretypes.Transaction)
	PushQueuedTx(tx *coretypes.Transaction)
	PendingTxs() []*rpctypes.RPCTransaction
	QueuedTxs() []*rpctypes.RPCTransaction
	NumPendingTxs() int
	NumQueuedTxs() int
	RemovePendingTx(hash common.Hash)
	RemoveQueuedTx(hash common.Hash)

	// bloom
	ReadBloomBits(ctx context.Context, section uint64, index uint32) ([]byte, error)
	PeekBloomBitsNextSection(ctx context.Context) (uint64, error)
	IsBloomIndexingRunning() bool

	// Close stops the indexer process, waits for pending operations to complete,
	// and flushes the store to disk with a timeout. Returns an error if the
	// timeout is reached before operations complete.
	Close() error

	// Wait waits for all the indexing to finish.
	Wait()
}

// EVMIndexerImpl implements EVMIndexer.
type EVMIndexerImpl struct {
	enabled      bool
	retainHeight uint64

	pruningRunning       *atomic.Bool
	bloomIndexingRunning *atomic.Bool
	lastIndexedHeight    *atomic.Uint64

	db       dbm.DB
	logger   log.Logger
	txConfig client.TxConfig
	ac       address.Codec
	appCodec codec.Codec

	store     *CacheStoreWithBatch
	evmKeeper *evmkeeper.Keeper

	schema collections.Schema

	// blocks
	BlockHeaderMap           collections.Map[uint64, coretypes.Header]
	BlockAndIndexToTxHashMap collections.Map[collections.Pair[uint64, uint64], []byte]
	BlockHashToNumberMap     collections.Map[[]byte, uint64]

	// txs
	TxMap                collections.Map[[]byte, rpctypes.RPCTransaction]
	TxReceiptMap         collections.Map[[]byte, coretypes.Receipt]
	TxHashToCosmosTxHash collections.Map[[]byte, []byte]
	CosmosTxHashToTxHash collections.Map[[]byte, []byte]

	// bloom
	BloomBits            collections.Map[collections.Pair[uint64, uint32], []byte]
	BloomBitsNextSection collections.Sequence

	blockChans   []chan *coretypes.Header
	logsChans    []chan []*coretypes.Log
	pendingChans []chan *rpctypes.RPCTransaction

	// pendingTxs is a map to store tx hashes in pending state.
	pendingTxs *ttlcache.Cache[common.Hash, *rpctypes.RPCTransaction]

	// queuedTxs is a map to store tx hashes in queued state.
	queuedTxs *ttlcache.Cache[common.Hash, *rpctypes.RPCTransaction]

	// indexingChan is a channel to receive indexing tasks.
	indexingChan chan *indexingTask

	// indexingWg is a wait group to wait for all the indexing to finish.
	indexingWg *sync.WaitGroup
}

// indexingTask is a task to be indexed.
type indexingTask struct {
	req *abci.RequestFinalizeBlock
	res *abci.ResponseFinalizeBlock

	// state dependent args for extractEthTxInfo
	args *indexingArgs
}

func NewEVMIndexer(
	db dbm.DB,
	ac address.Codec,
	appCodec codec.Codec,
	logger log.Logger,
	txConfig client.TxConfig,
	evmKeeper *evmkeeper.Keeper,
) (EVMIndexer, error) {
	cfg := evmKeeper.Config()

	store := NewCacheStoreWithBatch(db)
	sb := collections.NewSchemaBuilderFromAccessor(
		func(_ context.Context) corestoretypes.KVStore {
			return store
		},
	)

	logger.Info("EVM Indexer", "enable", !cfg.IndexerDisable)
	indexer := &EVMIndexerImpl{
		enabled:      !cfg.IndexerDisable,
		retainHeight: cfg.IndexerRetainHeight,

		pruningRunning:       &atomic.Bool{},
		bloomIndexingRunning: &atomic.Bool{},
		lastIndexedHeight:    &atomic.Uint64{},

		db:       db,
		store:    store,
		logger:   logger.With("module", "evm-indexer"),
		txConfig: txConfig,
		ac:       ac,
		appCodec: appCodec,

		evmKeeper: evmKeeper,

		TxMap:                    collections.NewMap(sb, prefixTx, "tx", collections.BytesKey, CollJsonVal[rpctypes.RPCTransaction]()),
		TxReceiptMap:             collections.NewMap(sb, prefixTxReceipt, "tx_receipt", collections.BytesKey, CollJsonVal[coretypes.Receipt]()),
		BlockHeaderMap:           collections.NewMap(sb, prefixBlockHeader, "block_header", collections.Uint64Key, CollJsonVal[coretypes.Header]()),
		BlockAndIndexToTxHashMap: collections.NewMap(sb, prefixBlockAndIndexToTxHash, "block_and_index_to_tx_hash", collections.PairKeyCodec(collections.Uint64Key, collections.Uint64Key), collections.BytesValue),
		BlockHashToNumberMap:     collections.NewMap(sb, prefixBlockHashToNumber, "block_hash_to_number", collections.BytesKey, collections.Uint64Value),
		TxHashToCosmosTxHash:     collections.NewMap(sb, prefixTxHashToCosmosTxHash, "tx_hash_to_cosmos_tx_hash", collections.BytesKey, collections.BytesValue),
		CosmosTxHashToTxHash:     collections.NewMap(sb, prefixCosmosTxHashToTxHash, "cosmos_tx_hash_to_tx_hash", collections.BytesKey, collections.BytesValue),
		BloomBits:                collections.NewMap(sb, prefixBloomBits, "bloom_bits", collections.PairKeyCodec(collections.Uint64Key, collections.Uint32Key), collections.BytesValue),
		BloomBitsNextSection:     collections.NewSequence(sb, prefixBloomBitsNextSection, "bloom_bits_next_section"),

		blockChans:   nil,
		logsChans:    nil,
		pendingChans: nil,

		// Use ttlcache to cope with abnormal cases like tx not included in a block
		pendingTxs: ttlcache.New(
			// pending tx lifetime is 1 minutes in indexer
			ttlcache.WithTTL[common.Hash, *rpctypes.RPCTransaction](time.Minute),
		),
		queuedTxs: ttlcache.New(
			// queued tx lifetime is 1 minutes in indexer
			ttlcache.WithTTL[common.Hash, *rpctypes.RPCTransaction](time.Minute),
		),

		// use buffered channel to avoid blocking the main thread
		indexingChan: make(chan *indexingTask, 10),

		// for graceful shutdown
		indexingWg: &sync.WaitGroup{},
	}

	schema, err := sb.Build()
	if err != nil {
		return nil, err
	}
	indexer.schema = schema

	// expire pending tx
	go indexer.pendingTxs.Start()
	go indexer.queuedTxs.Start()

	// start indexing loop
	go indexer.indexingLoop()

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

func (e *EVMIndexerImpl) GetLastIndexedHeight(ctx context.Context) (uint64, error) {
	// if lastIndexedHeight is not set, get the last indexed block header from the store
	lastIndexedHeight := e.lastIndexedHeight.Load()
	if lastIndexedHeight == 0 {
		blockHeader, err := e.BlockHeaderMap.Iterate(ctx, new(collections.Range[uint64]).Descending())
		if err != nil {
			return 0, err
		}

		if blockHeader.Valid() {
			lastIndexedHeight, err = blockHeader.Key()
			if err != nil {
				return 0, err
			}

			e.lastIndexedHeight.Store(lastIndexedHeight)
		}
	}

	return lastIndexedHeight, nil
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

// Close stops the indexer.
func (e *EVMIndexerImpl) Close() error {
	if e.pendingTxs != nil {
		e.pendingTxs.Stop()
	}
	if e.queuedTxs != nil {
		e.queuedTxs.Stop()
	}

	// flush store before stopping
	return e.flushStore()
}

// flushStore flushes the store before stopping the indexer.
// it waits for all the indexing to finish and then flushes the store.
// it also waits for pruning to complete before flushing the store.
func (e *EVMIndexerImpl) flushStore() error {
	e.logger.Info("service stop", "msg", "Stopping EVM indexer service")

	// wait for all the indexing to finish
	e.Wait()

	// wait for pruning to complete before flushing the store
	ticker := time.NewTicker(time.Millisecond * 100)
	timeout := time.NewTimer(time.Second * 30)
	defer ticker.Stop()
	defer timeout.Stop()

	for {
		select {
		case <-ticker.C:
			if !e.pruningRunning.Swap(true) {
				// if pruning is finished, flush the store
				e.store.Write()

				// close the store
				e.logger.Info("Closing evm_index.db")
				if err := e.store.db.Close(); err != nil {
					return err
				}

				return nil
			}
		case <-timeout.C:
			return fmt.Errorf("timeout waiting for EVM indexer to complete before shutdown")
		}
	}
}

// Wait waits for all the indexing to finish.
func (e *EVMIndexerImpl) Wait() {
	e.indexingWg.Wait()
}
