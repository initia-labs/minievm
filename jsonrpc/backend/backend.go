package backend

import (
	"context"

	bigcache "github.com/allegro/bigcache/v3"

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

	return &JSONRPCBackend{
		app, logger, queuedTxs, ctx, svrCtx, clientCtx, cfg,
	}, nil
}
