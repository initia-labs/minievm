package web3

import (
	"context"

	"cosmossdk.io/log"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/initia-labs/minievm/jsonrpc/backend"

	"github.com/ethereum/go-ethereum/crypto"
)

// Web3EthereumAPI is the web3 namespace for the Ethereum JSON-RPC APIs.
// Current it is used for tracking what APIs should be implemented for Ethereum compatibility.
// After fully implementing the Ethereum APIs, this interface can be removed.
type Web3EthereumAPI interface {
	ClientVersion() string
	Sha3(input hexutil.Bytes) hexutil.Bytes
}

type Web3API struct {
	ctx     context.Context
	logger  log.Logger
	backend *backend.JSONRPCBackend
}

// NewWeb3API creates a new net API instance
func NewWeb3API(logger log.Logger, backend *backend.JSONRPCBackend) *Web3API {
	return &Web3API{
		ctx:     context.TODO(),
		logger:  logger,
		backend: backend,
	}
}

// ClientVersion returns the node name
func (s *Web3API) ClientVersion() (string, error) {
	return s.backend.ClientVersion()
}

// Sha3 applies the ethereum sha3 implementation on the input.
// It assumes the input is hex encoded.
func (s *Web3API) Sha3(input hexutil.Bytes) hexutil.Bytes {
	return crypto.Keccak256(input)
}
