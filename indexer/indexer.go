package indexer

import (
	"context"
	"fmt"
	"io"
	"slices"
	"sync"
	"sync/atomic"
	"time"

	abci "github.com/cometbft/cometbft/abci/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	corestoretypes "cosmossdk.io/core/store"
	"cosmossdk.io/log"
	snapshot "cosmossdk.io/store/snapshots/types"
	storetypes "cosmossdk.io/store/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

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
	TxStartLogIndexByHash(ctx context.Context, hash common.Hash) (uint64, error)
	StoreTxStartLogIndex(ctx context.Context, hash common.Hash, index uint64) error

	// block
	BlockHashToNumber(ctx context.Context, hash common.Hash) (uint64, error)
	BlockHeaderByNumber(ctx context.Context, number uint64) (*coretypes.Header, error)

	// cosmos tx hash
	CosmosTxHashByTxHash(ctx context.Context, hash common.Hash) ([]byte, error)
	TxHashByCosmosTxHash(ctx context.Context, hash []byte) (common.Hash, error)

	// event subscription
	Subscribe() (chan *coretypes.Header, chan []*coretypes.Log, func())

	// mempool cache
	MempoolCache() *MempoolTxCache

	// last indexed height
	GetLastIndexedHeight(ctx context.Context) (uint64, error)

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

	// Initialize sets the client context.
	Initialize(clientCtx client.Context, contextCreator contextCreator, consensusParamsGetter consensusParamsGetter) error
}

// contextCreator creates a new SDK context.
type contextCreator func(height int64, prove bool) (sdk.Context, io.Closer, error)

// consensusParamsGetter gets the consensus params from the SDK context.
type consensusParamsGetter func(ctx sdk.Context) cmtproto.ConsensusParams

// EVMIndexerImpl implements EVMIndexer.
type EVMIndexerImpl struct {
	enabled             bool
	retainHeight        uint64
	backfillStartHeight uint64

	pruningRunning         *atomic.Bool
	pruneRequestedHeight   *atomic.Uint64
	bloomIndexingRunning   *atomic.Bool
	lastIndexedHeight      *atomic.Uint64
	lastPrunedHeight       *atomic.Uint64
	lastBloomIndexedHeight *atomic.Uint64

	db       dbm.DB
	logger   log.Logger
	txConfig client.TxConfig
	ac       address.Codec
	appCodec codec.Codec

	clientCtx             client.Context
	contextCreator        contextCreator
	consensusParamsGetter consensusParamsGetter

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
	TxStartLogIndexMap   collections.Map[[]byte, uint64]

	// bloom
	BloomBits            collections.Map[collections.Pair[uint64, uint32], []byte]
	BloomBitsNextSection collections.Sequence

	subMu sync.RWMutex
	subs  []blockLogsSub

	mempoolCache *MempoolTxCache

	// indexingChan is a channel to receive indexing tasks.
	indexingChan chan *indexingTask

	// indexingWg is a wait group to wait for all the indexing to finish.
	indexingWg *sync.WaitGroup

	// stopped indicates whether the indexer has been
	stopped atomic.Bool
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
		enabled:             !cfg.IndexerDisable,
		retainHeight:        cfg.IndexerRetainHeight,
		backfillStartHeight: cfg.IndexerBackfillStartHeight,

		pruningRunning:         &atomic.Bool{},
		pruneRequestedHeight:   &atomic.Uint64{},
		bloomIndexingRunning:   &atomic.Bool{},
		lastIndexedHeight:      &atomic.Uint64{},
		lastPrunedHeight:       &atomic.Uint64{},
		lastBloomIndexedHeight: &atomic.Uint64{},

		db:       db,
		store:    store,
		logger:   logger.With("module", "evm-indexer"),
		txConfig: txConfig,
		ac:       ac,
		appCodec: appCodec,

		evmKeeper: evmKeeper,

		TxMap:                    collections.NewMap(sb, prefixTx, "tx", collections.BytesKey, CollJsonVal[rpctypes.RPCTransaction]()),
		TxReceiptMap:             collections.NewMap(sb, prefixTxReceipt, "tx_receipt", collections.BytesKey, CollJsonVal[coretypes.Receipt]()),
		TxStartLogIndexMap:       collections.NewMap(sb, prefixTxStartLogIndex, "tx_start_log_index", collections.BytesKey, collections.Uint64Value),
		BlockHeaderMap:           collections.NewMap(sb, prefixBlockHeader, "block_header", collections.Uint64Key, CollJsonVal[coretypes.Header]()),
		BlockAndIndexToTxHashMap: collections.NewMap(sb, prefixBlockAndIndexToTxHash, "block_and_index_to_tx_hash", collections.PairKeyCodec(collections.Uint64Key, collections.Uint64Key), collections.BytesValue),
		BlockHashToNumberMap:     collections.NewMap(sb, prefixBlockHashToNumber, "block_hash_to_number", collections.BytesKey, collections.Uint64Value),
		TxHashToCosmosTxHash:     collections.NewMap(sb, prefixTxHashToCosmosTxHash, "tx_hash_to_cosmos_tx_hash", collections.BytesKey, collections.BytesValue),
		CosmosTxHashToTxHash:     collections.NewMap(sb, prefixCosmosTxHashToTxHash, "cosmos_tx_hash_to_tx_hash", collections.BytesKey, collections.BytesValue),
		BloomBits:                collections.NewMap(sb, prefixBloomBits, "bloom_bits", collections.PairKeyCodec(collections.Uint64Key, collections.Uint32Key), collections.BytesValue),
		BloomBitsNextSection:     collections.NewSequence(sb, prefixBloomBitsNextSection, "bloom_bits_next_section"),

		subs: nil,

		mempoolCache: NewMempoolTxCache(),

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

	return indexer, nil
}

// Initialize initializes the EVM indexer.
func (e *EVMIndexerImpl) Initialize(clientCtx client.Context, contextCreator contextCreator, consensusParamsGetter consensusParamsGetter) error {
	e.clientCtx = clientCtx
	e.contextCreator = contextCreator
	e.consensusParamsGetter = consensusParamsGetter

	if e.backfillStartHeight != 0 {
		lastIndexedHeight, err := e.GetLastIndexedHeight(context.Background())
		if err != nil {
			e.logger.Error("failed to get last indexed height", "err", err)
			return err
		}

		if e.backfillStartHeight < lastIndexedHeight {
			err = e.Backfill(e.backfillStartHeight, lastIndexedHeight)
			if err != nil {
				e.logger.Error("failed to backfill", "err", err)
				return err
			}
		}
	}

	e.logger.Info("EVM indexer initialized")

	// start indexing loop
	go e.indexingLoop()

	return nil
}

// MempoolCache returns the in-memory mempool transaction cache.
func (e *EVMIndexerImpl) MempoolCache() *MempoolTxCache {
	return e.mempoolCache
}

// blockLogsSub is a subscriber entry for block and log events.
type blockLogsSub struct {
	blockChan chan *coretypes.Header
	logsChan  chan []*coretypes.Log
	done      chan struct{} // closed by cancel to signal emitter to skip this sub
}

// Subscribe returns channels to receive blocks and logs, and a cancel func to unsubscribe.
// The cancel func is idempotent and safe to call from any goroutine.
func (e *EVMIndexerImpl) Subscribe() (chan *coretypes.Header, chan []*coretypes.Log, func()) {
	sub := blockLogsSub{
		blockChan: make(chan *coretypes.Header),
		logsChan:  make(chan []*coretypes.Log),
		done:      make(chan struct{}),
	}

	e.subMu.Lock()
	e.subs = append(e.subs, sub)
	e.subMu.Unlock()

	var once sync.Once
	cancel := func() {
		once.Do(func() {
			e.subMu.Lock()
			e.subs = slices.DeleteFunc(e.subs, func(ss blockLogsSub) bool {
				return ss.done == sub.done
			})
			e.subMu.Unlock()
			close(sub.done)
		})
	}

	return sub.blockChan, sub.logsChan, cancel
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

func (e *EVMIndexerImpl) GetLastPrunedHeight() uint64 {
	return e.lastPrunedHeight.Load()
}

func (e *EVMIndexerImpl) GetLastBloomIndexedHeight() uint64 {
	return e.lastBloomIndexedHeight.Load()
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

	// snapshot subscriber list under RLock to avoid races with cancel()
	e.subMu.RLock()
	subs := append([]blockLogsSub(nil), e.subs...)
	e.subMu.RUnlock()

	// emit logs first; use unbuffered channel to ensure logs are emitted before block header
	for _, logs := range blockEvents.logs {
		for _, sub := range subs {
			select {
			case sub.logsChan <- logs:
			case <-sub.done:
			}
		}
	}

	// emit block header
	for _, sub := range subs {
		select {
		case sub.blockChan <- blockEvents.header:
		case <-sub.done:
		}
	}
}

// Close stops the indexer.
func (e *EVMIndexerImpl) Close() error {
	if e.stopped.Swap(true) {
		return nil
	}

	e.logger.Info("EVM indexer closing...")

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
