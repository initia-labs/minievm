package backend

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server"

	evmindexer "github.com/initia-labs/minievm/indexer"
)

type JSONRPCBackend struct {
	svrCtx     *server.Context
	clientCtx  client.Context
	evmindexer evmindexer.EVMIndexer
}

// NewJSONRPCBackend creates a new JSONRPCBackend instance
func NewJSONRPCBackend(svrCtx *server.Context, clientCtx client.Context, evmindexer evmindexer.EVMIndexer) JSONRPCBackend {
	return JSONRPCBackend{
		svrCtx, clientCtx, evmindexer,
	}
}
