package txpool

import (
	"context"

	"cosmossdk.io/log"
	"github.com/initia-labs/minievm/jsonrpc/backend"
)

var _ TxpoolEthereumAPI = (*TxPoolAPI)(nil)

// TxpoolEthereumAPI is the txpool namespace for the Ethereum JSON-RPC APIs.
// Current it is used for tracking what APIs should be implemented for Ethereum compatibility.
// After fully implementing the Ethereum APIs, this interface can be removed.
type TxpoolEthereumAPI interface {
	// TODO: implement the following apis
	//Content() map[string]map[string]map[string]*rpctypes.RPCTransaction
	//ContentFrom() map[string]map[string]*rpctypes.RPCTransaction
	//Inspect() map[string]map[string]map[string]string
	//Status() map[string]hexutil.Uint
}

// TxPoolAPI is the txpool namespace for the Ethereum JSON-RPC APIs.
type TxPoolAPI struct {
	ctx     context.Context
	logger  log.Logger
	backend *backend.JSONRPCBackend
}

// NewTxPoolAPI creates a new txpool API instance.
func NewTxPoolAPI(logger log.Logger, backend *backend.JSONRPCBackend) *TxPoolAPI {
	return &TxPoolAPI{
		ctx:     context.TODO(),
		logger:  logger,
		backend: backend,
	}
}

// TODO: implement txpool_content
//func (api *TxPoolAPI) Content() map[string]map[string]map[string]*rpctypes.RPCTransaction {
//	return nil
//}

// TODO: implement txpool_contentFrom
//func (api *TxPoolAPI) ContentFrom() map[string]map[string]*rpctypes.RPCTransaction {
//	return nil
//}

// TODO: implement txpool_inspect
//func (api *TxPoolAPI) Inspect() map[string]map[string]map[string]string {
//	return nil
//}

// TODO: implement txpool_status
//func (api *TxPoolAPI) Status() map[string]hexutil.Uint {
//	return nil
//}
