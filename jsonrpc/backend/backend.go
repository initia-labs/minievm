package backend

import (
	"context"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/lru"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	lrucache "github.com/hashicorp/golang-lru/v2"

	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server"

	"github.com/initia-labs/minievm/app"
	"github.com/initia-labs/minievm/jsonrpc/config"
	rpctypes "github.com/initia-labs/minievm/jsonrpc/types"
)

type JSONRPCBackend struct {
	app    *app.MinitiaApp
	logger log.Logger

	queuedTxs    *lrucache.Cache[string, []byte]
	historyCache *lru.Cache[cacheKey, processedFees]

	headerCache        *lru.Cache[uint64, *coretypes.Header]
	blockTxsCache      *lru.Cache[uint64, []*rpctypes.RPCTransaction]
	blockReceiptsCache *lru.Cache[uint64, []*coretypes.Receipt]
	blockHashCache     *lru.Cache[common.Hash, uint64]

	txLookupCache *lru.Cache[common.Hash, *rpctypes.RPCTransaction]
	receiptCache  *lru.Cache[common.Hash, *coretypes.Receipt]
	logsCache     *lru.Cache[uint64, []*coretypes.Log]

	mut     sync.Mutex // mutex for accMuts
	accMuts map[string]*AccMut

	ctx       context.Context
	svrCtx    *server.Context
	clientCtx client.Context

	cfg config.JSONRPCConfig
}

const (
	feeHistoryCacheSize = 2048
	blockCacheLimit     = 256
	txLookupCacheLimit  = 1024
)

// NewJSONRPCBackend creates a new JSONRPCBackend instance
func NewJSONRPCBackend(
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

	queuedTxs, err := lrucache.New[string, []byte](cfg.QueuedTransactionCap)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	return &JSONRPCBackend{
		app:    app,
		logger: logger,

		queuedTxs:    queuedTxs,
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
		cfg:       cfg,
	}, nil
}

type AccMut struct {
	mut sync.Mutex
	rc  int // reference count
}

// acquireAccMut acquires the mutex for the account with the given senderHex
// and increments the reference count. If the mutex does not exist, it is created.
func (b *JSONRPCBackend) acquireAccMut(senderHex string) {
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
}

// releaseAccMut releases the mutex for the account with the given senderHex
// and decrements the reference count. If the reference count reaches zero,
// the mutex is deleted.
func (b *JSONRPCBackend) releaseAccMut(senderHex string) {
	accMut := b.accMuts[senderHex]
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
