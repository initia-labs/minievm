package types

const (
	EventTypeCall   = "call"
	EventTypeCreate = "create"
	EventTypeEVM    = "evm"
	EventTypeSubmsg = "submsg"

	// state db events
	EventTypeContractCreated = "contract_created"

	// erc20 events
	EventTypeERC20Created = "erc20_created"

	// erc721 events
	EventTypeERC721Created = "erc721_created"
	EventTypeERC721Minted  = "erc721_minted"
	EventTypeERC721Burned  = "erc721_burned"

	// NOTE: if there are changes on the list of event types,

	AttributeKeyContract = "contract"
	AttributeKeyAddress  = "address"
	AttributeKeyLog      = "log"
	AttributeKeyData     = "data"
	AttributeKeyRet      = "ret"
	AttributeKeyDenom    = "denom"

	AttributeKeyClassId       = "class_id"
	AttributeKeyTokenId       = "token_id"
	AttributeKeyTokenOriginId = "token_origin_id"

	AttributeKeySuccess = "success"
	AttributeKeyReason  = "reason"
)
