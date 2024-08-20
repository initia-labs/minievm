package backend

import (
	"context"
	"sync"

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

	queuedTxs *lrucache.Cache[string, []byte]
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

	queuedTxs, err := lrucache.New[string, []byte](cfg.QueuedTransactionCap)
	if err != nil {
		return nil, err
	}
	accMuts, err := lrucache.New[string, *sync.Mutex](cfg.QueuedTransactionCap / 10)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
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
