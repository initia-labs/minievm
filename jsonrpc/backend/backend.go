package backend

import (
	"context"
	"sync"

	bigcache "github.com/allegro/bigcache/v3"
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

	queuedTxs *bigcache.BigCache
	accMuts   *lrucache.Cache[string, *sync.Mutex]

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

	if cfg.QueuedTransactionCap == 0 {
		cfg.QueuedTransactionCap = config.DefaultQueuedTransactionCap
	}
	if cfg.QueuedTransactionTTL == 0 {
		cfg.QueuedTransactionTTL = config.DefaultQueuedTransactionTTL
	}

	cacheConfig := bigcache.DefaultConfig(cfg.QueuedTransactionTTL)
	cacheConfig.HardMaxCacheSize = cfg.QueuedTransactionCap

	ctx := context.Background()

	queuedTxs, err := bigcache.New(ctx, cacheConfig)
	if err != nil {
		return nil, err
	}

	// support concurrent 100 accounts mutex
	accMuts, err := lrucache.New[string, *sync.Mutex](100)
	if err != nil {
		return nil, err
	}

	return &JSONRPCBackend{
		app:    app,
		logger: logger,

		queuedTxs: queuedTxs,
		accMuts:   accMuts,

		ctx:       ctx,
		svrCtx:    svrCtx,
		clientCtx: clientCtx,
		cfg:       cfg,
	}, nil
}
