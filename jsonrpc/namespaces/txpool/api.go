package txpool

import (
	"context"

	"cosmossdk.io/log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/initia-labs/minievm/jsonrpc/backend"
	rpctypes "github.com/initia-labs/minievm/jsonrpc/types"
)

var _ TxpoolEthereumAPI = (*TxPoolAPI)(nil)

// TxpoolEthereumAPI is the txpool namespace for the Ethereum JSON-RPC APIs.
// Current it is used for tracking what APIs should be implemented for Ethereum compatibility.
// After fully implementing the Ethereum APIs, this interface can be removed.
type TxpoolEthereumAPI interface {
	Content() (map[string]map[string]map[string]*rpctypes.RPCTransaction, error)
	ContentFrom(addr common.Address) (map[string]map[string]*rpctypes.RPCTransaction, error)
	Inspect() (map[string]map[string]map[string]string, error)
	Status() (map[string]hexutil.Uint, error)
}

// TxPoolAPI is the txpool namespace for the Ethereum JSON-RPC APIs.
type TxPoolAPI struct {
	ctx     context.Context
	logger  log.Logger
	backend *backend.JSONRPCBackend
}

// NewTxPoolAPI creates a new txpool API instance.
func NewTxPoolAPI(ctx context.Context, logger log.Logger, backend *backend.JSONRPCBackend) *TxPoolAPI {
	return &TxPoolAPI{
		ctx:     ctx,
		logger:  logger,
		backend: backend,
	}
}

func (api *TxPoolAPI) Content() (map[string]map[string]map[string]*rpctypes.RPCTransaction, error) {
	return api.backend.TxPoolContent()
}

func (api *TxPoolAPI) ContentFrom(addr common.Address) (map[string]map[string]*rpctypes.RPCTransaction, error) {
	return api.backend.TxPoolContentFrom(addr)
}

func (api *TxPoolAPI) Inspect() (map[string]map[string]map[string]string, error) {
	return api.backend.TxPoolInspect()
}

func (api *TxPoolAPI) Status() (map[string]hexutil.Uint, error) {
	return api.backend.TxPoolStatus()
}
