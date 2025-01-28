package backend

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/lru"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	lrucache "github.com/hashicorp/golang-lru/v2"

	"cosmossdk.io/log"
	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server"

	"github.com/initia-labs/minievm/app"
	"github.com/initia-labs/minievm/jsonrpc/config"
	rpctypes "github.com/initia-labs/minievm/jsonrpc/types"
)

type JSONRPCBackend struct {
	app    *app.MinitiaApp
	logger log.Logger

	queuedTxs      *lrucache.Cache[string, txQueueItem]
	queuedTxHashes *sync.Map

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
	queuedTxs, err := lrucache.NewWithEvict(cfg.QueuedTransactionCap, func(_ string, txCache txQueueItem) {
		queuedTxHashes.Delete(txCache.hash)
	})
	if err != nil {
		return nil, err
	}

	b := &JSONRPCBackend{
		app:    app,
		logger: logger,

		queuedTxs:      queuedTxs,
		queuedTxHashes: queuedTxHashes,

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
	}

	// start fee fetcher
	go b.feeFetcher()

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
