package backend

import (
	"context"
	"sync"
	"time"

	lrucache "github.com/hashicorp/golang-lru/v2"

	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server"

	"github.com/initia-labs/minievm/app"
	"github.com/initia-labs/minievm/jsonrpc/config"
)

type JSONRPCBackend struct {
	app    *app.MinitiaApp
	logger log.Logger

	queuedTxs    *lrucache.Cache[string, []byte]
	historyCache *lrucache.Cache[cacheKey, processedFees]

	mut     sync.Mutex // mutex for accMuts
	accMuts map[string]*AccMut

	ctx       context.Context
	svrCtx    *server.Context
	clientCtx client.Context

	cfg config.JSONRPCConfig
}

// NewJSONRPCBackend creates a new JSONRPCBackend instance
func NewJSONRPCBackend(
	app *app.MinitiaApp,
	logger log.Logger,
	svrCtx *server.Context,
	clientCtx client.Context,
	cfg config.JSONRPCConfig,
) (*JSONRPCBackend, error) {
	queuedTxs, err := lrucache.New[string, []byte](cfg.QueuedTransactionCap)
	if err != nil {
		return nil, err
	}
	historyCache, err := lrucache.New[cacheKey, processedFees](2048)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	return &JSONRPCBackend{
		app:    app,
		logger: logger,

		queuedTxs:    queuedTxs,
		historyCache: historyCache,

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
