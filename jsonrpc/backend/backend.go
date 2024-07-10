package backend

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server"

	"github.com/initia-labs/minievm/app"
	"github.com/initia-labs/minievm/jsonrpc/config"
)

type JSONRPCBackend struct {
	app       *app.MinitiaApp
	svrCtx    *server.Context
	clientCtx client.Context
	cfg       config.JSONRPCConfig

	ctx context.Context
}

// NewJSONRPCBackend creates a new JSONRPCBackend instance
func NewJSONRPCBackend(
	app *app.MinitiaApp,
	svrCtx *server.Context,
	clientCtx client.Context,
	cfg config.JSONRPCConfig,
) *JSONRPCBackend {
	ctx := context.Background()
	return &JSONRPCBackend{
		app, svrCtx, clientCtx, cfg, ctx,
	}
}
