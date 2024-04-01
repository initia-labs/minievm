package types

const (
	EventTypeCall   = "call"
	EventTypeCreate = "create"
	EventTypeEVM    = "evm"

	// erc20 events
	EventTypeERC20Created = "erc20_created"

	AttributeKeyContract = "contract"
	AttributeKeyAddress  = "address"
	AttributeKeyLog      = "log"
	AttributeKeyData     = "data"
	AttributeKeyRet      = "ret"
	AttributeKeyDenom    = "denom"
)
