package debug

import (
	"cosmossdk.io/log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/eth/tracers"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/initia-labs/minievm/jsonrpc/backend"
	rpctypes "github.com/initia-labs/minievm/jsonrpc/types"
)

var _ DebugEthereumAPI = (*DebugAPI)(nil)

// DebugEthereumAPI is a collection of debug namespaced APIs.
type DebugEthereumAPI interface {
	TraceBlockByNumber(number rpc.BlockNumber, config *tracers.TraceConfig) ([]*rpctypes.TxTraceResult, error)
	TraceBlockByHash(hash common.Hash, config *tracers.TraceConfig) ([]*rpctypes.TxTraceResult, error)
}

// DebugAPI is the debug namespace for the Ethereum JSON-RPC APIs.
type DebugAPI struct {
	logger  log.Logger
	backend *backend.JSONRPCBackend
}

// NewDebugAPI creates an instance of the public ETH Web3 API.
func NewDebugAPI(logger log.Logger, backend *backend.JSONRPCBackend) *DebugAPI {
	api := &DebugAPI{
		logger:  logger.With("client", "json-rpc"),
		backend: backend,
	}

	return api
}

// *************************************
// *               Trace              *
// *************************************

// GetBlockByNumber returns the block identified by number.
func (api *DebugAPI) TraceBlockByNumber(ethBlockNum rpc.BlockNumber, config *tracers.TraceConfig) ([]*rpctypes.TxTraceResult, error) {
	api.logger.Debug("debug_traceBlockByNumber", "number", ethBlockNum, "config", config)
	return api.backend.TraceBlockByNumber(ethBlockNum, config)
}

// GetBlockByHash returns the block identified by hash.
func (api *DebugAPI) TraceBlockByHash(hash common.Hash, config *tracers.TraceConfig) ([]*rpctypes.TxTraceResult, error) {
	api.logger.Debug("debug_traceBlockByHash", "hash", hash.Hex(), "config", config)
	return api.backend.TraceBlockByHash(hash, config)
}
