package cosmosprecompile

import (
	storetypes "cosmossdk.io/store/types"

	"github.com/ethereum/go-ethereum/common"
)

// IsBlockedAddressArguments is the arguments for the is_blocked_address method.
type IsBlockedAddressArguments struct {
	Account common.Address `abi:"account"`
}

// IsModuleAddressArguments is the arguments for the is_module_address method.
type IsModuleAddressArguments struct {
	Account common.Address `abi:"account"`
}

// IsAuthorityAddressArguments is the arguments for the is_authority_address method.
type IsAuthorityAddressArguments struct {
	Account common.Address `abi:"account"`
}

// ToCosmosAddressArguments is the arguments for the to_cosmos_address method.
type ToCosmosAddressArguments struct {
	EVMAddress common.Address `abi:"evm_address"`
}

// ToEVMAddressArguments is the arguments for the to_evm_address method.
type ToEVMAddressArguments struct {
	CosmosAddress string `abi:"cosmos_address"`
}

type ExecuteCosmos struct {
	Msg     string         `abi:"msg"`
	Options ExecuteOptions `abi:"options"`
}

type ExecuteOptions struct {
	AllowFailure bool   `json:"allow_failure"`
	CallbackId   uint64 `json:"callback_id"`
}

// QueryCosmosArguments is the arguments for the query_cosmos method.
type QueryCosmosArguments struct {
	Path string `abi:"path"`
	Req  string `abi:"req"`
}

// ToDenomArguments is the arguments for the to_denom method.
type ToDenomArguments struct {
	ERC20Address common.Address `abi:"erc20_address"`
}

// ToERC20Arguments is the arguments for the to_erc20 method.
type ToERC20Arguments struct {
	Denom string `abi:"denom"`
}

const (
	IS_BLOCKED_ADDRESS_GAS   storetypes.Gas = 100
	IS_MODULE_ADDRESS_GAS    storetypes.Gas = 200
	IS_AUTHORITY_ADDRESS_GAS storetypes.Gas = 200

	TO_COSMOS_ADDRESS_GAS storetypes.Gas = 200
	TO_EVM_ADDRESS_GAS    storetypes.Gas = 200

	TO_DENOM_GAS storetypes.Gas = 100
	TO_ERC20_GAS storetypes.Gas = 100

	QUERY_COSMOS_GAS   storetypes.Gas = 200
	EXECUTE_COSMOS_GAS storetypes.Gas = 200

	GAS_PER_BYTE storetypes.Gas = 1
)

const (
	METHOD_IS_BLOCKED_ADDRESS   = "is_blocked_address"
	METHOD_IS_MODULE_ADDRESS    = "is_module_address"
	METHOD_IS_AUTHORITY_ADDRESS = "is_authority_address"

	METHOD_TO_COSMOS_ADDRESS = "to_cosmos_address"
	METHOD_TO_EVM_ADDRESS    = "to_evm_address"

	METHOD_QUERY_COSMOS                = "query_cosmos"
	METHOD_EXECUTE_COSMOS              = "execute_cosmos"
	METHOD_EXECUTE_COSMOS_WITH_OPTIONS = "execute_cosmos_with_options"

	METHOD_TO_DENOM = "to_denom"
	METHOD_TO_ERC20 = "to_erc20"
)
