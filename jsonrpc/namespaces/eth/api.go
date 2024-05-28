package ethapi

import "github.com/ethereum/go-ethereum/common/hexutil"

// The Ethereum API allows applications to connect to an Evmos node that is
// part of the Evmos blockchain. Developers can interact with on-chain EVM data
// and send different types of transactions to the network by utilizing the
// endpoints provided by the API. The API follows a JSON-RPC standard. If not
// otherwise specified, the interface is derived from the Alchemy Ethereum API:
// https://docs.alchemy.com/alchemy/apis/ethereum
type EthereumAPI interface {
	// Getting Blocks
	//
	// Retrieves information from a particular block in the blockchain.
	BlockNumber() (hexutil.Uint64, error)
	GetBlockByNumber(ethBlockNum rpctypes.BlockNumber, fullTx bool) (map[string]interface{}, error)
	// GetBlockByHash(hash common.Hash, fullTx bool) (map[string]interface{}, error)
	// GetBlockTransactionCountByHash(hash common.Hash) *hexutil.Uint
	// GetBlockTransactionCountByNumber(blockNum rpctypes.BlockNumber) *hexutil.Uint

	// // Reading Transactions
	// //
	// // Retrieves information on the state data for addresses regardless of whether
	// // it is a user or a smart contract.
	// GetTransactionByHash(hash common.Hash) (*rpctypes.RPCTransaction, error)
	// GetTransactionCount(address common.Address, blockNrOrHash rpctypes.BlockNumberOrHash) (*hexutil.Uint64, error)
	// GetTransactionReceipt(hash common.Hash) (map[string]interface{}, error)
	// GetTransactionByBlockHashAndIndex(hash common.Hash, idx hexutil.Uint) (*rpctypes.RPCTransaction, error)
	// GetTransactionByBlockNumberAndIndex(blockNum rpctypes.BlockNumber, idx hexutil.Uint) (*rpctypes.RPCTransaction, error)
	// // eth_getBlockReceipts

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
	// GetBalance(address common.Address, blockNrOrHash rpctypes.BlockNumberOrHash) (*hexutil.Big, error)
	// GetStorageAt(address common.Address, key string, blockNrOrHash rpctypes.BlockNumberOrHash) (hexutil.Bytes, error)
	// GetCode(address common.Address, blockNrOrHash rpctypes.BlockNumberOrHash) (hexutil.Bytes, error)
	// GetProof(address common.Address, storageKeys []string, blockNrOrHash rpctypes.BlockNumberOrHash) (*rpctypes.AccountResult, error)

	// // EVM/Smart Contract Execution
	// //
	// // Allows developers to read data from the blockchain which includes executing
	// // smart contracts. However, no data is published to the Ethereum network.
	// Call(args evmtypes.TransactionArgs, blockNrOrHash rpctypes.BlockNumberOrHash, _ *rpctypes.StateOverride) (hexutil.Bytes, error)

	// // Chain Information
	// //
	// // Returns information on the Ethereum network and internal settings.
	// ProtocolVersion() hexutil.Uint
	// GasPrice() (*hexutil.Big, error)
	// EstimateGas(args evmtypes.TransactionArgs, blockNrOptional *rpctypes.BlockNumber) (hexutil.Uint64, error)
	// FeeHistory(blockCount rpc.DecimalOrHex, lastBlock rpc.BlockNumber, rewardPercentiles []float64) (*rpctypes.FeeHistoryResult, error)
	// MaxPriorityFeePerGas() (*hexutil.Big, error)
	// ChainId() (*hexutil.Big, error)

	// // Getting Uncles
	// //
	// // Returns information on uncle blocks are which are network rejected blocks and replaced by a canonical block instead.
	// GetUncleByBlockHashAndIndex(hash common.Hash, idx hexutil.Uint) map[string]interface{}
	// GetUncleByBlockNumberAndIndex(number, idx hexutil.Uint) map[string]interface{}
	// GetUncleCountByBlockHash(hash common.Hash) hexutil.Uint
	// GetUncleCountByBlockNumber(blockNum rpctypes.BlockNumber) hexutil.Uint

	// // Proof of Work
	// Hashrate() hexutil.Uint64
	// Mining() bool

	// // Other
	// Syncing() (interface{}, error)
	// Coinbase() (string, error)
	// Sign(address common.Address, data hexutil.Bytes) (hexutil.Bytes, error)
	// GetTransactionLogs(txHash common.Hash) ([]*ethtypes.Log, error)
	// SignTypedData(address common.Address, typedData apitypes.TypedData) (hexutil.Bytes, error)
	// FillTransaction(args evmtypes.TransactionArgs) (*rpctypes.SignTransactionResult, error)
	// Resend(ctx context.Context, args evmtypes.TransactionArgs, gasPrice *hexutil.Big, gasLimit *hexutil.Uint64) (common.Hash, error)
	// GetPendingTransactions() ([]*rpctypes.RPCTransaction, error)
	// // eth_signTransaction (on Ethereum.org)
	// // eth_getCompilers (on Ethereum.org)
	// // eth_compileSolidity (on Ethereum.org)
	// // eth_compileLLL (on Ethereum.org)
	// // eth_compileSerpent (on Ethereum.org)
	// // eth_getWork (on Ethereum.org)
	// // eth_submitWork (on Ethereum.org)
	// // eth_submitHashrate (on Ethereum.org)
}
