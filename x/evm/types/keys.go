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
	VMStorePrefix               = []byte{0x21} // prefix for vm
	TransientVMStorePrefix      = []byte{0x22} // prefix for transient vm store
	TransientCreatedPrefix      = []byte{0x23} // prefix for transient created accounts
	TransientSelfDestructPrefix = []byte{0x24} // prefix for transient self destruct accounts
	TransientLogsPrefix         = []byte{0x25} // prefix for transient logs
	TransientLogSizePrefix      = []byte{0x26} // prefix for transient log size
	TransientAccessListPrefix   = []byte{0x27} // prefix for transient access list
	TransientRefundPrefix       = []byte{0x28} // prefix for transient refund

	ERC20sPrefix                    = []byte{0x31} // prefix for erc20 stores
	ERC20StoresPrefix               = []byte{0x32} // prefix for erc20 stores
	ERC20DenomsByContractAddrPrefix = []byte{0x33} // prefix for erc20 denoms
	ERC20ContractAddrsByDenomPrefix = []byte{0x34} // prefix for erc20 denoms

	ERC721ClassURIPrefix               = []byte{0x41} // prefix for erc721 class uris
	ERC721ClassIdsByContractAddrPrefix = []byte{0x42} // prefix for erc721 class ids
	ERC721ContractAddrsByClassIdPrefix = []byte{0x43} // prefix for erc721 contract addresses

	EVMBlockHashPrefix = []byte{0x71} // prefix for evm block hashes

	ParamsKey           = []byte{0x51} // key of parameters for module x/evm
	ERC20FactoryAddrKey = []byte{0x61} // key of erc20 factory contract address
	ERC20WrapperAddrKey = []byte{0x62} // key of erc20 wrapper contract address
)

// ContextKey type for context key
type ContextKey int

const (
	// CONTEXT_KEY_EXECUTE_REQUESTS is a context key for execute requests
	CONTEXT_KEY_EXECUTE_REQUESTS ContextKey = iota

	// CONTEXT_KEY_RECURSIVE_DEPTH is a context key for recursive depth
	CONTEXT_KEY_RECURSIVE_DEPTH
)
