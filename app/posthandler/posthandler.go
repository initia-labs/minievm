package posthandler

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
)

// NewPostHandler returns a new sdk.PostHandler that is composed of the sdk.ChainPostDecorators
func NewPostHandler(ak authante.AccountKeeper) sdk.PostHandler {
	return sdk.ChainPostDecorators(NewSequenceIncrementDecorator(ak))
}
