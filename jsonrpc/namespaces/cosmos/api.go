package cosmos

import (
	"context"

	"cosmossdk.io/log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/initia-labs/minievm/jsonrpc/backend"
)

// CosmosAPI is the cosmos namespace for the Ethereum JSON-RPC APIs.
type CosmosAPI struct {
	ctx     context.Context
	logger  log.Logger
	backend *backend.JSONRPCBackend
}

// NewCosmosAPI creates an instance of the public ETH Web3 API.
func NewCosmosAPI(logger log.Logger, backend *backend.JSONRPCBackend) *CosmosAPI {
	api := &CosmosAPI{
		ctx:     context.TODO(),
		logger:  logger.With("client", "json-rpc"),
		backend: backend,
	}

	return api
}

func (api *CosmosAPI) CosmosTxHashByTxHash(hash common.Hash) (hexutil.Bytes, error) {
	return api.backend.CosmosTxHashByTxHash(hash)
}

func (api *CosmosAPI) TxHashByCosmosTxHash(hash hexutil.Bytes) (common.Hash, error) {
	return api.backend.TxHashByCosmosTxHash(hash)
}
