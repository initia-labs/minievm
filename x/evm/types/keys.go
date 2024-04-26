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
	VMStorePrefix                   = []byte{0x21} // prefix for vm
	ERC20sPrefix                    = []byte{0x31} // prefix for erc20 stores
	ERC20StoresPrefix               = []byte{0x32} // prefix for erc20 stores
	ERC20DenomsByContractAddrPrefix = []byte{0x33} // prefix for erc20 denoms
	ERC20ContractAddrsByDenomPrefix = []byte{0x34} // prefix for erc20 denoms

	ERC721sPrefix                      = []byte{0x41} // prefix for erc721 stores
	ERC721StoresPrefix                 = []byte{0x42} // prefix for erc721 stores
	ERC721ClassIdsByContractAddrPrefix = []byte{0x43} // prefix for erc721 denoms
	ERC721ContractAddrsByClassIdPrefix = []byte{0x44} // prefix for erc721 denoms

	ParamsKey = []byte{0x51} // key of parameters for module x/evm
	VMRootKey = []byte{0x61} // key of evm state root
)

// ContextKey type for context key
type ContextKey int

const (
	// CONTEXT_KEY_COSMOS_MESSAGES is a context key for cosmos messages
	CONTEXT_KEY_COSMOS_MESSAGES ContextKey = iota
)
