package net

import (
	"context"

	"cosmossdk.io/log"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/initia-labs/minievm/jsonrpc/backend"
)

// NetEthereumAPI is the net namespace for the Ethereum JSON-RPC APIs.
// Current it is used for tracking what APIs should be implemented for Ethereum compatibility.
// After fully implementing the Ethereum APIs, this interface can be removed.
type NetEthereumAPI interface {
	Listening() bool
	PeerCount() hexutil.Uint
	Version() string
}

type NetAPI struct {
	ctx     context.Context
	logger  log.Logger
	backend *backend.JSONRPCBackend
}

// NewNetAPI creates a new net API instance
func NewNetAPI(logger log.Logger, backend *backend.JSONRPCBackend) *NetAPI {
	return &NetAPI{
		ctx:     context.TODO(),
		logger:  logger,
		backend: backend,
	}
}

// TODO: implement net_listening
//func (api *NetAPI) Listening() bool {
//	return true
//}

// TODO: implement net_peerCount
//func (api *NetAPI) PeerCount() hexutil.Uint {
//	return hexutil.Uint(0)
//}

func (api *NetAPI) Version() string {
	v, err := api.backend.Version()
	if err != nil {
		api.logger.Error("failed to get version", "err", err)
		return "1"
	}
	return v
}
