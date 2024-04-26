package types

const (
	EventTypeCall   = "call"
	EventTypeCreate = "create"
	EventTypeEVM    = "evm"

	// erc20 events
	EventTypeERC20Created = "erc20_created"
	// erc721 events
	EventTypeERC721Created = "erc721_created"
	EventTypeERC721Minted  = "erc721_minted"
	EventTypeERC721Burned  = "erc721_burned"

	AttributeKeyContract = "contract"
	AttributeKeyAddress  = "address"
	AttributeKeyLog      = "log"
	AttributeKeyData     = "data"
	AttributeKeyRet      = "ret"
	AttributeKeyDenom    = "denom"

	AttributeKeyClassId = "class_id"
	AttributeKeyTokenId = "token_id"
)
