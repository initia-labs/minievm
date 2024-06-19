package ethapi

import (
	"context"
	"encoding/hex"
	"errors"
	"strings"

	"cosmossdk.io/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/initia-labs/minievm/jsonrpc/backend"
	rpctypes "github.com/initia-labs/minievm/jsonrpc/types"
)

var _ EthEthereumAPI = (*EthAPI)(nil)

// EthEthereumAPI is a collection of eth namespaced APIs.
// Current it is used for tracking what APIs should be implemented for Ethereum compatibility.
// After fully implementing the Ethereum APIs, this interface can be removed.
type EthEthereumAPI interface {
	// Getting Blocks
	//
	// Retrieves information from a particular block in the blockchain.
	BlockNumber() hexutil.Uint64
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
	GetRawTransactionByBlockHashAndIndex(blockHash common.Hash, index hexutil.Uint) hexutil.Bytes

	// eth_getBlockReceipts

	// Writing Transactions
	//
	// Allows developers to both send ETH from one address to another, write data
	// on-chain, and interact with smart contracts.
	SendRawTransaction(data hexutil.Bytes) (common.Hash, error)
	//SendTransaction(args rpctypes.TransactionArgs) (common.Hash, error)
	// eth_sendPrivateTransaction
	// eth_cancel	PrivateTransaction

	// Account Information
	//
	// Returns information regarding an address's stored on-chain data.
	//Accounts() ([]common.Address, error)
	GetBalance(address common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*hexutil.Big, error)
	GetStorageAt(address common.Address, hexKey string, blockNrOrHash rpc.BlockNumberOrHash) (hexutil.Bytes, error)
	GetCode(address common.Address, blockNrOrHash rpc.BlockNumberOrHash) (hexutil.Bytes, error)
	// GetProof(address common.Address, storageKeys []string, blockNrOrHash rpc.BlockNumberOrHash) (*rpc.AccountResult, error)

	// EVM/Smart Contract Execution
	//
	// Allows developers to read data from the blockchain which includes executing
	// smart contracts. However, no data is published to the Ethereum network.
	Call(args rpctypes.TransactionArgs, blockNrOrHash *rpc.BlockNumberOrHash, so *rpctypes.StateOverride, bo *rpctypes.BlockOverrides) (hexutil.Bytes, error)

	// // Chain Information
	// //
	// // Returns information on the Ethereum network and internal settings.
	// ProtocolVersion() hexutil.Uint
	GasPrice() (*hexutil.Big, error)
	EstimateGas(args rpctypes.TransactionArgs, blockNrOptional *rpc.BlockNumberOrHash, so *rpctypes.StateOverride) (hexutil.Uint64, error)
	// FeeHistory(blockCount rpc.DecimalOrHex, lastBlock rpc.BlockNumber, rewardPercentiles []float64) (*rpc.FeeHistoryResult, error)
	MaxPriorityFeePerGas() (*hexutil.Big, error)
	ChainId() *hexutil.Big

	// // Other
	// Syncing() (interface{}, error)
	// Coinbase() (string, error)
	// Sign(address common.Address, data hexutil.Bytes) (hexutil.Bytes, error)
	// GetTransactionLogs(txHash common.Hash) ([]*ethtypes.Log, error)
	// SignTypedData(address common.Address, typedData apitypes.TypedData) (hexutil.Bytes, error)
	// FillTransaction(args evmtypes.TransactionArgs) (*rpc.SignTransactionResult, error)
	// Resend(args evmtypes.TransactionArgs, gasPrice *hexutil.Big, gasLimit *hexutil.Uint64) (common.Hash, error)
	PendingTransactions() ([]*rpctypes.RPCTransaction, error)
	// // eth_signTransaction (on Ethereum.org)
	// // eth_getCompilers (on Ethereum.org)
	// // eth_compileSolidity (on Ethereum.org)
	// // eth_compileLLL (on Ethereum.org)
	// // eth_compileSerpent (on Ethereum.org)
	// // eth_getWork (on Ethereum.org)
	// // eth_submitWork (on Ethereum.org)
	// // eth_submitHashrate (on Ethereum.org)
}

// EthAPI is the txpool namespace for the Ethereum JSON-RPC APIs.
type EthAPI struct {
	ctx     context.Context
	logger  log.Logger
	backend *backend.JSONRPCBackend
}

// NewEthAPI creates an instance of the public ETH Web3 API.
func NewEthAPI(logger log.Logger, backend *backend.JSONRPCBackend) *EthAPI {
	api := &EthAPI{
		ctx:     context.TODO(),
		logger:  logger.With("client", "json-rpc"),
		backend: backend,
	}

	return api
}

// *************************************
// *               Blocks              *
// *************************************

// BlockNumber returns the current block number.
func (api *EthAPI) BlockNumber() hexutil.Uint64 {
	api.logger.Debug("eth_blockNumber")
	blockNumber, err := api.backend.BlockNumber()
	if err != nil {
		api.logger.Error("eth_blockNumber", "error", err)
		return 0
	}
	return blockNumber
}

// GetBlockByNumber returns the block identified by number.
func (api *EthAPI) GetBlockByNumber(ethBlockNum rpc.BlockNumber, fullTx bool) (map[string]interface{}, error) {
	api.logger.Debug("eth_getBlockByNumber", "number", ethBlockNum, "full", fullTx)
	return api.backend.GetBlockByNumber(ethBlockNum, fullTx)
}

// GetBlockByHash returns the block identified by hash.
func (api *EthAPI) GetBlockByHash(hash common.Hash, fullTx bool) (map[string]interface{}, error) {
	api.logger.Debug("eth_getBlockByHash", "hash", hash.Hex(), "full", fullTx)
	return api.backend.GetBlockByHash(hash, fullTx)
}

// *************************************
// *              Read Txs             *
// *************************************

// GetTransactionByHash returns the transaction identified by hash.
func (api *EthAPI) GetTransactionByHash(hash common.Hash) (*rpctypes.RPCTransaction, error) {
	api.logger.Debug("eth_getTransactionByHash", "hash", hash.Hex())
	return api.backend.GetTransactionByHash(hash)
}

// GetRawTransactionByHash returns the bytes of the transaction for the given tx hash.
func (api *EthAPI) GetRawTransactionByHash(blockHash common.Hash) (hexutil.Bytes, error) {
	return api.backend.GetRawTransactionByHash(blockHash)
}

// GetRawTransactionByBlockHashAndIndex returns the bytes of the transaction for the given block hash and index.
func (api *EthAPI) GetRawTransactionByBlockHashAndIndex(blockHash common.Hash, index hexutil.Uint) hexutil.Bytes {
	rawTx, err := api.backend.GetRawTransactionByBlockHashAndIndex(blockHash, index)
	if err != nil {
		api.logger.Error("eth_getRawTransactionByBlockHashAndIndex", "error", err)
		return nil
	}
	return rawTx
}

// GetTransactionCount returns the number of transactions at the given address up to the given block number.
func (api *EthAPI) GetTransactionCount(address common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*hexutil.Uint64, error) {
	api.logger.Debug("eth_getTransactionCount", "address", address.Hex(), "block number or hash", blockNrOrHash)
	return api.backend.GetTransactionCount(address, blockNrOrHash)
}

// GetTransactionReceipt returns the transaction receipt identified by hash.
func (api *EthAPI) GetTransactionReceipt(hash common.Hash) (map[string]interface{}, error) {
	hexTx := hash.Hex()
	api.logger.Debug("eth_getTransactionReceipt", "hash", hexTx)
	return api.backend.GetTransactionReceipt(hash)
}

// GetBlockTransactionCountByHash returns the number of transactions in the block identified by hash.
func (api *EthAPI) GetBlockTransactionCountByHash(hash common.Hash) (*hexutil.Uint, error) {
	api.logger.Debug("eth_getBlockTransactionCountByHash", "hash", hash.Hex())
	return api.backend.GetBlockTransactionCountByHash(hash)
}

// GetBlockTransactionCountByNumber returns the number of transactions in the block identified by number.
func (api *EthAPI) GetBlockTransactionCountByNumber(blockNum rpc.BlockNumber) (*hexutil.Uint, error) {
	api.logger.Debug("eth_getBlockTransactionCountByNumber", "height", blockNum.Int64())
	return api.backend.GetBlockTransactionCountByNumber(blockNum)
}

// GetTransactionByBlockHashAndIndex returns the transaction identified by hash and index.
func (api *EthAPI) GetTransactionByBlockHashAndIndex(hash common.Hash, idx hexutil.Uint) (*rpctypes.RPCTransaction, error) {
	api.logger.Debug("eth_getTransactionByBlockHashAndIndex", "hash", hash.Hex(), "index", idx)
	return api.backend.GetTransactionByBlockHashAndIndex(hash, idx)
}

// GetTransactionByBlockNumberAndIndex returns the transaction identified by number and index.
func (api *EthAPI) GetTransactionByBlockNumberAndIndex(blockNum rpc.BlockNumber, idx hexutil.Uint) (*rpctypes.RPCTransaction, error) {
	api.logger.Debug("eth_getTransactionByBlockNumberAndIndex", "number", blockNum, "index", idx)
	return api.backend.GetTransactionByBlockNumberAndIndex(blockNum, idx)
}

// *************************************
// *             Write Txs             *
// *************************************

// SendRawTransaction send a raw Ethereum transaction.
func (api *EthAPI) SendRawTransaction(data hexutil.Bytes) (common.Hash, error) {
	api.logger.Debug("eth_sendRawTransaction", "length", len(data))
	return api.backend.SendRawTransaction(data)
}

// TODO: Implement eth_sendTransaction
//// SendTransaction sends an Ethereum transaction.
//func (e *EthAPI) SendTransaction(args rpctypes.TransactionArgs) (common.Hash, error) {
//	e.logger.Debug("eth_sendTransaction", "args", args.String())
//	return e.backend.SendTransaction(args)
//}

// *************************************
// *         Account Information       *
// *************************************

// TODO: Implement eth_accounts
//// Accounts returns the list of accounts available to this node.
//func (e *EthAPI) Accounts() ([]common.Address, error) {
//	e.logger.Debug("eth_accounts")
//	return e.backend.Accounts()
//}

// GetBalance returns the provided account's balance up to the provided block number.
func (api *EthAPI) GetBalance(address common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*hexutil.Big, error) {
	api.logger.Debug("eth_getBalance", "address", address.String(), "block number or hash", blockNrOrHash)
	return api.backend.GetBalance(address, blockNrOrHash)
}

// GetStorageAt returns the contract storage at the given address, block number, and key.
func (api *EthAPI) GetStorageAt(address common.Address, hexKey string, blockNrOrHash rpc.BlockNumberOrHash) (hexutil.Bytes, error) {
	api.logger.Debug("eth_getStorageAt", "address", address.Hex(), "hexKey", hexKey, "block number or hash", blockNrOrHash)
	key, _, err := decodeHash(hexKey)
	if err != nil {
		return hexutil.Bytes{}, err
	}
	return api.backend.GetStorageAt(address, key, blockNrOrHash)
}

// GetCode returns the contract code at the given address and block number.
func (api *EthAPI) GetCode(address common.Address, blockNrOrHash rpc.BlockNumberOrHash) (hexutil.Bytes, error) {
	api.logger.Debug("eth_getCode", "address", address.Hex(), "block number or hash", blockNrOrHash)
	return api.backend.GetCode(address, blockNrOrHash)
}

// TODO: Implement eth_getProof
//// GetProof returns an account object with proof and any storage proofs
//func (e *EthAPI) GetProof(address common.Address,
//	storageKeys []string,
//	blockNrOrHash rpctypes.BlockNumberOrHash,
//) (*rpctypes.AccountResult, error) {
//	e.logger.Debug("eth_getProof", "address", address.Hex(), "keys", storageKeys, "block number or hash", blockNrOrHash)
//	return e.backend.GetProof(address, storageKeys, blockNrOrHash)
//}

// **********************************************
// *         EVM/Smart Contract Execution       *
// **********************************************

// Call performs a raw contract call.
func (api *EthAPI) Call(
	args rpctypes.TransactionArgs,
	blockNrOrHash *rpc.BlockNumberOrHash,
	so *rpctypes.StateOverride,
	bo *rpctypes.BlockOverrides,
) (hexutil.Bytes, error) {
	api.logger.Debug("eth_call", "args", args.String(), "block number or hash", blockNrOrHash, "state override", so, "block override", bo)
	return api.backend.Call(args, blockNrOrHash, so, bo)
}

///////////////////////////////////////////////////////////////////////////////
///                           Event Logs													          ///
///////////////////////////////////////////////////////////////////////////////
// FILTER API at ./filters/api.go

///////////////////////////////////////////////////////////////////////////////
///                           Chain Information										          ///
///////////////////////////////////////////////////////////////////////////////

// TODO: Implement eth_protocolVersion
//// ProtocolVersion returns the supported Ethereum protocol version.
//func (e *EthAPI) ProtocolVersion() hexutil.Uint {
//	e.logger.Debug("eth_protocolVersion")
//	return hexutil.Uint(ethermint.ProtocolVersion)
//}

// GasPrice returns the current gas price based on Ethermint's gas price oracle.
func (api *EthAPI) GasPrice() (*hexutil.Big, error) {
	api.logger.Debug("eth_gasPrice")
	return api.backend.GasPrice()
}

// EstimateGas returns an estimate of gas usage for the given smart contract call.
func (api *EthAPI) EstimateGas(
	args rpctypes.TransactionArgs,
	blockNrOptional *rpc.BlockNumberOrHash,
	so *rpctypes.StateOverride,
) (hexutil.Uint64, error) {
	api.logger.Debug("eth_estimateGas")
	return api.backend.EstimateGas(args, blockNrOptional, so)
}

//// TODO: Implement eth_feeHistory
//// FeeHistory returns the fee history for the last blockCount blocks.
//func (e *EthAPI) FeeHistory(blockCount rpc.DecimalOrHex,
//	lastBlock rpc.BlockNumber,
//	rewardPercentiles []float64,
//) (*rpctypes.FeeHistoryResult, error) {
//	e.logger.Debug("eth_feeHistory")
//	return e.backend.FeeHistory(blockCount, lastBlock, rewardPercentiles)
//}

// MaxPriorityFeePerGas returns a suggestion for a gas tip cap for dynamic fee transactions.
func (api *EthAPI) MaxPriorityFeePerGas() (*hexutil.Big, error) {
	api.logger.Debug("eth_maxPriorityFeePerGas")
	return api.backend.MaxPriorityFeePerGas()
}

// ChainId is the EIP-155 replay-protection chain id for the current ethereum chain config.
func (api *EthAPI) ChainId() *hexutil.Big { //nolint
	api.logger.Debug("eth_chainId")
	chainId, err := api.backend.ChainID()
	if err != nil {
		return nil
	}
	return (*hexutil.Big)(chainId)
}

///////////////////////////////////////////////////////////////////////////////
///                           Uncles															          ///
///////////////////////////////////////////////////////////////////////////////

// *************************************
// *               Uncles              *
// *************************************

// GetUncleByBlockHashAndIndex returns the uncle identified by hash and index. Always returns nil.
func (api *EthAPI) GetUncleByBlockHashAndIndex(_ common.Hash, _ hexutil.Uint) map[string]interface{} {
	return nil
}

// GetUncleByBlockNumberAndIndex returns the uncle identified by number and index. Always returns nil.
func (api *EthAPI) GetUncleByBlockNumberAndIndex(_, _ hexutil.Uint) map[string]interface{} {
	return nil
}

// GetUncleCountByBlockHash returns the number of uncles in the block identified by hash. Always zero.
func (api *EthAPI) GetUncleCountByBlockHash(_ common.Hash) hexutil.Uint {
	return 0
}

// GetUncleCountByBlockNumber returns the number of uncles in the block identified by number. Always zero.
func (api *EthAPI) GetUncleCountByBlockNumber(_ rpc.BlockNumber) hexutil.Uint {
	return 0
}

// *************************************
// *             pow(leagcy)           *
// *************************************

// Hashrate returns the current node's hashrate. Always 0.
func (api *EthAPI) Hashrate() hexutil.Uint64 {
	api.logger.Debug("eth_hashrate")
	return 0
}

// Mining returns whether or not this node is currently mining. Always false.
func (api *EthAPI) Mining() bool {
	api.logger.Debug("eth_mining")
	return false
}

// *************************************
// *               others              *
// *************************************

// TODO: Implement eth_syncing
//// Syncing returns false in case the node is currently not syncing with the network. It can be up to date or has not
//// yet received the latest block headers from its pears. In case it is synchronizing:
//// - startingBlock: block number this node started to synchronize from
//// - currentBlock:  block number this node is currently importing
//// - highestBlock:  block number of the highest block header this node has received from peers
//// - pulledStates:  number of state entries processed until now
//// - knownStates:   number of known state entries that still need to be pulled
//func (e *EthAPI) Syncing() (interface{}, error) {
//	e.logger.Debug("eth_syncing")
//	return e.backend.Syncing()
//}

// TODO: Implement eth_coinbase
//// Coinbase is the address that staking rewards will be send to (alias for Etherbase).
//func (e *EthAPI) Coinbase() (string, error) {
//	e.logger.Debug("eth_coinbase")
//
//	coinbase, err := e.backend.GetCoinbase()
//	if err != nil {
//		return "", err
//	}
//	ethAddr := common.BytesToAddress(coinbase.Bytes())
//	return ethAddr.Hex(), nil
//}

// TODO: Implement eth_sign
//// Sign signs the provided data using the private key of address via Geth's signature standard.
//func (e *EthAPI) Sign(address common.Address, data hexutil.Bytes) (hexutil.Bytes, error) {
//	e.logger.Debug("eth_sign", "address", address.Hex(), "data", common.Bytes2Hex(data))
//	return e.backend.Sign(address, data)
//}

// TODO: Implement eth_signTypedData
//// SignTypedData signs EIP-712 conformant typed data
//func (e *EthAPI) SignTypedData(address common.Address, typedData apitypes.TypedData) (hexutil.Bytes, error) {
//	e.logger.Debug("eth_signTypedData", "address", address.Hex(), "data", typedData)
//	return e.backend.SignTypedData(address, typedData)
//}

// TODO: Implement eth_fillTransaction
//// FillTransaction fills the defaults (nonce, gas, gasPrice or 1559 fields)
//// on a given unsigned transaction, and returns it to the caller for further
//// processing (signing + broadcast).
//func (e *EthAPI) FillTransaction(args rpctypes.TransactionArgs) (ethapi.SignTransactionResult, error) {
//	// Set some sanity defaults and terminate on failure
//	args, err := e.backend.SetTxDefaults(args)
//	if err != nil {
//		return nil, err
//	}
//
//	// Assemble the transaction and obtain rlp
//	tx := args.ToTransaction().AsTransaction()
//
//	data, err := tx.MarshalBinary()
//	if err != nil {
//		return nil, err
//	}
//
//	return &ethapi.SignTransactionResult{
//		Raw: data,
//		Tx:  tx,
//	}, nil
//}

// TODO: Implement eth_resend
//// Resend accepts an existing transaction and a new gas price and limit. It will remove
//// the given transaction from the pool and reinsert it with the new gas price and limit.
//func (e *EthAPI) Resend(_ context.Context,
//	args rpctypes.TransactionArgs,
//	gasPrice *hexutil.Big,
//	gasLimit *hexutil.Uint64,
//) (common.Hash, error) {
//	e.logger.Debug("eth_resend", "args", args.String())
//	return e.backend.Resend(args, gasPrice, gasLimit)
//}

// PendingTransactions returns the transactions that are in the transaction pool
// and have a from address that is one of the accounts this node manages.
func (api *EthAPI) PendingTransactions() ([]*rpctypes.RPCTransaction, error) {
	api.logger.Debug("eth_pendingTransactions")
	return api.backend.PendingTransactions()
}

// decodeHash parses a hex-encoded 32-byte hash. The input may optionally
// be prefixed by 0x and can have a byte length up to 32.
func decodeHash(s string) (h common.Hash, inputLength int, err error) {
	if strings.HasPrefix(s, "0x") || strings.HasPrefix(s, "0X") {
		s = s[2:]
	}
	if (len(s) & 1) > 0 {
		s = "0" + s
	}
	b, err := hex.DecodeString(s)
	if err != nil {
		return common.Hash{}, 0, errors.New("hex string invalid")
	}
	if len(b) > 32 {
		return common.Hash{}, len(b), errors.New("hex string too long, want at most 32 bytes")
	}
	return common.BytesToHash(b), len(b), nil
}
