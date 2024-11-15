package cosmos

import (
	"context"

	"cosmossdk.io/log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/initia-labs/minievm/jsonrpc/backend"
)

// CosmosAPI is the cosmos namespace for the Ethereum JSON-RPC APIs.
type CosmosAPI struct {
	ctx     context.Context
	logger  log.Logger
	backend *backend.JSONRPCBackend
}

// NewCosmosAPI creates an instance of the public ETH Web3 API.
func NewCosmosAPI(ctx context.Context, logger log.Logger, backend *backend.JSONRPCBackend) *CosmosAPI {
	api := &CosmosAPI{
		ctx:     ctx,
		logger:  logger.With("client", "json-rpc"),
		backend: backend,
	}

	return api
}

func (api *CosmosAPI) CosmosTxHashByTxHash(hash common.Hash) (common.UnprefixedHash, error) {
	bz, err := api.backend.CosmosTxHashByTxHash(hash)
	if err != nil {
		return common.UnprefixedHash{}, err
	}
	if bz == nil {
		return common.UnprefixedHash{}, nil
	}

	return common.UnprefixedHash(bz), nil
}

func (api *CosmosAPI) TxHashByCosmosTxHash(hash common.UnprefixedHash) (common.Hash, error) {
	return api.backend.TxHashByCosmosTxHash(hash[:])
}
