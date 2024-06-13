package ethapi

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/initia-labs/minievm/jsonrpc/backend"
	rpctypes "github.com/initia-labs/minievm/jsonrpc/types"
)

var _ EthereumAPI = (*backend.JSONRPCBackend)(nil)

type EthereumAPI interface {
	// Getting Blocks
	//
	// Retrieves information from a particular block in the blockchain.
	BlockNumber() (hexutil.Uint64, error)
	GetBlockByNumber(ethBlockNum rpc.BlockNumber, fullTx bool) (map[string]interface{}, error)
	GetBlockByHash(hash common.Hash, fullTx bool) (map[string]interface{}, error)

	// Reading Transactions
	//
	// Retrieves information on the state data for addresses regardless of whether
	// it is a user or a smart contract.
	GetTransactionByHash(hash common.Hash) (*rpctypes.RPCTransaction, error)
	GetTransactionCount(address common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*hexutil.Uint64, error)
	GetTransactionReceipt(hash common.Hash) (map[string]interface{}, error)
	GetTransactionByBlockHashAndIndex(hash common.Hash, idx hexutil.Uint) (*rpctypes.RPCTransaction, error)
	GetTransactionByBlockNumberAndIndex(blockNum rpc.BlockNumber, idx hexutil.Uint) (*rpctypes.RPCTransaction, error)
	GetBlockTransactionCountByHash(hash common.Hash) (*hexutil.Uint, error)
	GetBlockTransactionCountByNumber(blockNum rpc.BlockNumber) (*hexutil.Uint, error)
	GetRawTransactionByHash(hash common.Hash) (hexutil.Bytes, error)
	GetRawTransactionByBlockHashAndIndex(blockHash common.Hash, index hexutil.Uint) (hexutil.Bytes, error)

	// eth_getBlockReceipts

	// // Writing Transactions
	// //
	// // Allows developers to both send ETH from one address to another, write data
	// // on-chain, and interact with smart contracts.
	// SendRawTransaction(data hexutil.Bytes) (common.Hash, error)
	// SendTransaction(args evmtypes.TransactionArgs) (common.Hash, error)
	// // eth_sendPrivateTransaction
	// // eth_cancel	PrivateTransaction

	// // Account Information
	// //
	// // Returns information regarding an address's stored on-chain data.
	// Accounts() ([]common.Address, error)
	// GetBalance(address common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*hexutil.Big, error)
	// GetStorageAt(address common.Address, key string, blockNrOrHash rpc.BlockNumberOrHash) (hexutil.Bytes, error)
	// GetCode(address common.Address, blockNrOrHash rpc.BlockNumberOrHash) (hexutil.Bytes, error)
	// GetProof(address common.Address, storageKeys []string, blockNrOrHash rpc.BlockNumberOrHash) (*rpc.AccountResult, error)

	// // EVM/Smart Contract Execution
	// //
	// // Allows developers to read data from the blockchain which includes executing
	// // smart contracts. However, no data is published to the Ethereum network.
	// Call(args evmtypes.TransactionArgs, blockNrOrHash rpc.BlockNumberOrHash, _ *rpc.StateOverride) (hexutil.Bytes, error)

	// // Chain Information
	// //
	// // Returns information on the Ethereum network and internal settings.
	// ProtocolVersion() hexutil.Uint
	// GasPrice() (*hexutil.Big, error)
	// EstimateGas(args evmtypes.TransactionArgs, blockNrOptional *rpc.BlockNumber) (hexutil.Uint64, error)
	// FeeHistory(blockCount rpc.DecimalOrHex, lastBlock rpc.BlockNumber, rewardPercentiles []float64) (*rpc.FeeHistoryResult, error)
	// MaxPriorityFeePerGas() (*hexutil.Big, error)
	// ChainId() (*hexutil.Big, error)

	// // Other
	// Syncing() (interface{}, error)
	// Coinbase() (string, error)
	// Sign(address common.Address, data hexutil.Bytes) (hexutil.Bytes, error)
	// GetTransactionLogs(txHash common.Hash) ([]*ethtypes.Log, error)
	// SignTypedData(address common.Address, typedData apitypes.TypedData) (hexutil.Bytes, error)
	// FillTransaction(args evmtypes.TransactionArgs) (*rpc.SignTransactionResult, error)
	// Resend(args evmtypes.TransactionArgs, gasPrice *hexutil.Big, gasLimit *hexutil.Uint64) (common.Hash, error)
	GetPendingTransactions() ([]*rpctypes.RPCTransaction, error)
	// // eth_signTransaction (on Ethereum.org)
	// // eth_getCompilers (on Ethereum.org)
	// // eth_compileSolidity (on Ethereum.org)
	// // eth_compileLLL (on Ethereum.org)
	// // eth_compileSerpent (on Ethereum.org)
	// // eth_getWork (on Ethereum.org)
	// // eth_submitWork (on Ethereum.org)
	// // eth_submitHashrate (on Ethereum.org)
}
