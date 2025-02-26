package backend

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/lru"
	"github.com/ethereum/go-ethereum/core/bloombits"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	lrucache "github.com/hashicorp/golang-lru/v2"

	"cosmossdk.io/log"
	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/initia-labs/minievm/app"
	"github.com/initia-labs/minievm/jsonrpc/config"
	rpctypes "github.com/initia-labs/minievm/jsonrpc/types"
	evmconfig "github.com/initia-labs/minievm/x/evm/config"
)

type JSONRPCBackend struct {
	app    *app.MinitiaApp
	logger log.Logger

	queuedTxs        *lrucache.Cache[string, txQueueItem]
	queuedTxHashes   *sync.Map
	queuedTxAccounts *sync.Map

	historyCache *lru.Cache[cacheKey, processedFees]

	// per block caches
	headerCache        *lru.Cache[uint64, *coretypes.Header]
	blockTxsCache      *lru.Cache[uint64, []*rpctypes.RPCTransaction]
	blockReceiptsCache *lru.Cache[uint64, []*coretypes.Receipt]
	blockHashCache     *lru.Cache[common.Hash, uint64]
	logsCache          *lru.Cache[uint64, []*coretypes.Log]

	// per tx caches
	txLookupCache *lru.Cache[common.Hash, *rpctypes.RPCTransaction]
	receiptCache  *lru.Cache[common.Hash, *coretypes.Receipt]

	// fee cache
	feeDenom    string
	feeDecimals uint8
	feeMutex    sync.RWMutex

	mut     sync.Mutex // mutex for accMuts
	accMuts map[string]*AccMut

	ctx       context.Context
	svrCtx    *server.Context
	clientCtx client.Context

	cfg           config.JSONRPCConfig
	gasMultiplier math.LegacyDec

	// Channel receiving bloom data retrieval requests
	bloomRequests chan chan *bloombits.Retrieval
}

type txQueueItem struct {
	hash  common.Hash
	bytes []byte

	sender string
	body   *coretypes.Transaction
}

const (
	feeHistoryCacheSize = 2048
	blockCacheLimit     = 256
	txLookupCacheLimit  = 1024
)

// NewJSONRPCBackend creates a new JSONRPCBackend instance
func NewJSONRPCBackend(
	ctx context.Context,
	app *app.MinitiaApp,
	logger log.Logger,
	svrCtx *server.Context,
	clientCtx client.Context,
	cfg config.JSONRPCConfig,
) (*JSONRPCBackend, error) {
	if cfg.QueuedTransactionCap == 0 {
		cfg.QueuedTransactionCap = config.DefaultQueuedTransactionCap
	}
	if cfg.LogCacheSize == 0 {
		cfg.LogCacheSize = config.DefaultLogCacheSize
	}
	if cfg.GasMultiplier == "" {
		cfg.GasMultiplier = config.DefaultGasMultiplier
	}
	if cfg.FilterMaxBlockRange == 0 {
		cfg.FilterMaxBlockRange = config.DefaultFilterMaxBlockRange
	}

	gasMultiplier, err := math.LegacyNewDecFromStr(cfg.GasMultiplier)
	if err != nil {
		return nil, err
	}

	queuedTxHashes := new(sync.Map)
	queuedTxAccounts := new(sync.Map)
	queuedTxs, err := lrucache.NewWithEvict(cfg.QueuedTransactionCap, func(_ string, txCache txQueueItem) {
		queuedTxHashes.Delete(txCache.hash)

		// decrement the reference count of the sender account
		// if the reference count reaches zero, then delete account from the map
		if rc, ok := queuedTxAccounts.Load(txCache.sender); ok {
			if rc := rc.(*atomic.Int64).Add(-1); rc == 0 {
				queuedTxAccounts.Delete(txCache.sender)
			}
		}
	})
	if err != nil {
		return nil, err
	}

	b := &JSONRPCBackend{
		app:    app,
		logger: logger.With("module", "jsonrpc"),

		queuedTxs:        queuedTxs,
		queuedTxHashes:   queuedTxHashes,
		queuedTxAccounts: queuedTxAccounts,

		historyCache: lru.NewCache[cacheKey, processedFees](feeHistoryCacheSize),

		// per block caches
		headerCache:        lru.NewCache[uint64, *coretypes.Header](blockCacheLimit),
		blockTxsCache:      lru.NewCache[uint64, []*rpctypes.RPCTransaction](blockCacheLimit),
		blockReceiptsCache: lru.NewCache[uint64, []*coretypes.Receipt](blockCacheLimit),
		blockHashCache:     lru.NewCache[common.Hash, uint64](blockCacheLimit),
		logsCache:          lru.NewCache[uint64, []*coretypes.Log](cfg.LogCacheSize),

		// per tx caches
		txLookupCache: lru.NewCache[common.Hash, *rpctypes.RPCTransaction](txLookupCacheLimit),
		receiptCache:  lru.NewCache[common.Hash, *coretypes.Receipt](txLookupCacheLimit),

		accMuts: make(map[string]*AccMut),

		ctx:       ctx,
		svrCtx:    svrCtx,
		clientCtx: clientCtx,

		cfg:           cfg,
		gasMultiplier: gasMultiplier,

		bloomRequests: make(chan chan *bloombits.Retrieval),
	}

	// start fee fetcher
	go b.feeFetcher()

	// start queued tx flusher
	go b.queuedTxFlusher()

	// Start the bloom bits servicing goroutines
	b.startBloomHandlers(evmconfig.SectionSize)

	return b, nil
}

func (b *JSONRPCBackend) feeInfo() (string, uint8, error) {
	b.feeMutex.RLock()
	defer b.feeMutex.RUnlock()

	if b.feeDenom == "" {
		return "", 0, NewInternalError("jsonrpc is not ready")
	}

	return b.feeDenom, b.feeDecimals, nil
}

func (b *JSONRPCBackend) feeFetcher() {
	fetcher := func() (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("feeFetcher panic: %v", r)
			}
		}()

		if b.app.LastBlockHeight() <= 1 {
			return nil
		}

		queryCtx, err := b.getQueryCtx()
		if err != nil {
			return err
		}

		params, err := b.app.EVMKeeper.Params.Get(queryCtx)
		if err != nil {
			return err
		}

		feeDenom := params.FeeDenom
		decimals, err := b.app.EVMKeeper.ERC20Keeper().GetDecimals(queryCtx, feeDenom)
		if err != nil {
			return err
		}

		b.feeMutex.Lock()
		b.feeDenom = feeDenom
		b.feeDecimals = decimals
		b.feeMutex.Unlock()

		return nil
	}

	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := fetcher(); err != nil {
				b.logger.Error("failed to fetch fee", "err", err)
			}
		case <-b.ctx.Done():
			return
		}
	}
}

func (b *JSONRPCBackend) queuedTxFlusher() {
	flushRunning := &sync.Map{}
	workerPool := make(chan struct{}, 16)

	flusher := func() (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("queuedTxFlusher panic: %v", r)
			}
		}()

		if b.app.LastBlockHeight() <= 1 {
			return nil
		}

		// load all accounts in the queued txs
		var accounts []string
		b.queuedTxAccounts.Range(func(key, value any) bool {
			senderHex := key.(string)
			accounts = append(accounts, senderHex)

			return true
		})

		checkCtx := b.app.GetContextForCheckTx(nil)
		for _, senderHex := range accounts {
			select {
			case <-b.ctx.Done():
				return nil
			case workerPool <- struct{}{}: // Acquire worker slot
			default:
				// Skip if worker pool is full
				b.logger.Debug("skipping flush due to worker pool full", "sender", senderHex)
				continue
			}

			// trigger the flush for each sender
			go func(senderHex string) {
				defer func() { <-workerPool }() // Release worker slot

				accSeq := uint64(0)
				sender := sdk.AccAddress(common.HexToAddress(senderHex).Bytes())
				if acc := b.app.AccountKeeper.GetAccount(checkCtx, sender); acc != nil {
					accSeq = acc.GetSequence()
				}

				running, _ := flushRunning.LoadOrStore(senderHex, &atomic.Bool{})
				if running.(*atomic.Bool).Swap(true) {
					return
				}

				if err := b.flushQueuedTxs(senderHex, accSeq); err != nil {
					b.logger.Error("failed to flush queued txs", "err", err)
				}

				running.(*atomic.Bool).Store(false)
			}(senderHex)
		}

		return nil
	}

	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := flusher(); err != nil {
				b.logger.Error("failed to flush queued txs", "err", err)
			}
		case <-b.ctx.Done():
			return
		}
	}
}

type AccMut struct {
	mut sync.Mutex
	rc  int // reference count
}

// acquireAccMut acquires the mutex for the account with the given senderHex
// and increments the reference count. If the mutex does not exist, it is created.
func (b *JSONRPCBackend) acquireAccMut(senderHex string) *AccMut {
	// critical section for rc and create
	b.mut.Lock()
	accMut, ok := b.accMuts[senderHex]
	if !ok {
		accMut = &AccMut{rc: 0}
		b.accMuts[senderHex] = accMut
	}
	accMut.rc++
	b.mut.Unlock()
	// critical section end

	accMut.mut.Lock()
	return accMut
}

// releaseAccMut releases the mutex for the account with the given senderHex
// and decrements the reference count. If the reference count reaches zero,
// the mutex is deleted.
func (b *JSONRPCBackend) releaseAccMut(senderHex string, accMut *AccMut) {
	accMut.mut.Unlock()

	// critical section for rc and delete
	b.mut.Lock()
	accMut.rc--
	if accMut.rc == 0 {
		delete(b.accMuts, senderHex)
	}
	b.mut.Unlock()
	// critical section end
}

func (b *JSONRPCBackend) FilterTimeout() time.Duration {
	return b.cfg.FilterTimeout
}

func (b *JSONRPCBackend) FilterMaxBlockRange() int {
	return b.cfg.FilterMaxBlockRange
}
