package types

const (
	// ModuleName is the name of the move module
	ModuleName = "evm"

	// StoreKey is the string store representation
	StoreKey = ModuleName

	// TStoreKey is the string transient store representation
	TStoreKey = "transient_" + ModuleName

	// QuerierRoute is the querier route for the move module
	QuerierRoute = ModuleName

	// RouterKey is the msg router key for the move module
	RouterKey = ModuleName
)

// Keys for move store
// Items are stored with the following key: values
var (
	VMStorePrefix = []byte{0x21} // prefix for vm

	ERC20sPrefix                    = []byte{0x31} // prefix for erc20 stores
	ERC20StoresPrefix               = []byte{0x32} // prefix for erc20 stores
	ERC20DenomsByContractAddrPrefix = []byte{0x33} // prefix for erc20 denoms
	ERC20ContractAddrsByDenomPrefix = []byte{0x34} // prefix for erc20 denoms

	ERC721ClassURIPrefix               = []byte{0x41} // prefix for erc721 class uris
	ERC721ClassIdsByContractAddrPrefix = []byte{0x42} // prefix for erc721 class ids
	ERC721ContractAddrsByClassIdPrefix = []byte{0x43} // prefix for erc721 contract addresses

	EVMBlockHashPrefix = []byte{0x71} // prefix for evm block hashes

	ParamsKey            = []byte{0x51} // key of parameters for module x/evm
	ERC20FactoryAddrKey  = []byte{0x61} // key of erc20 factory contract address
	ERC20WrapperAddrKey  = []byte{0x62} // key of erc20 wrapper contract address
	ConnectOracleAddrKey = []byte{0x63} // key of connect oracle contract address
)

// ContextKey type for context key
type ContextKey int

const (
	// CONTEXT_KEY_EXECUTE_REQUESTS is a context key for execute requests
	CONTEXT_KEY_EXECUTE_REQUESTS ContextKey = iota

	// CONTEXT_KEY_PARENT_EXECUTE_REQUESTS is a context key for parent execute requests
	CONTEXT_KEY_PARENT_EXECUTE_REQUESTS

	// CONTEXT_KEY_RECURSIVE_DEPTH is a context key for recursive depth
	CONTEXT_KEY_RECURSIVE_DEPTH

	// CONTEXT_KEY_LOAD_DECIMALS is a context key for load decimals
	CONTEXT_KEY_LOAD_DECIMALS

	// CONTEXT_KEY_TRACING is a context key for tracing
	CONTEXT_KEY_TRACING

	// CONTEXT_KEY_ETH_TX is a context key for ethereum tx
	CONTEXT_KEY_ETH_TX

	// CONTEXT_KEY_ETH_TX_SENDER is a context key for ethereum tx sender
	CONTEXT_KEY_ETH_TX_SENDER

	// CONTEXT_KEY_GAS_PRICES is a context key for gas prices
	CONTEXT_KEY_GAS_PRICES

	// CONTEXT_KEY_SEQUENCE_INCREMENTED is a context key for sequence incremented
	CONTEXT_KEY_SEQUENCE_INCREMENTED

	// CONTEXT_TRACE_EVM is a context key for trace evm
	CONTEXT_KEY_TRACE_EVM
)
