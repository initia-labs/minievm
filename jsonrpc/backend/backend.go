package backend

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server"
)

type JSONRPCBackend struct {
	svrCtx    *server.Context
	clientCtx client.Context
}

// NewJSONRPCBackend creates a new JSONRPCBackend instance
func NewJSONRPCBackend(svrCtx *server.Context, clientCtx client.Context) JSONRPCBackend {
	return JSONRPCBackend{
		svrCtx, clientCtx,
	}
}
