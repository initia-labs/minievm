package posthandler

import (
	"context"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"

	evmtypes "github.com/initia-labs/minievm/x/evm/types"
)

// NewPostHandler returns a new sdk.PostHandler that is composed of the sdk.ChainPostDecorators
func NewPostHandler(
	ak authante.AccountKeeper,
	ek EVMKeeper,
) sdk.PostHandler {
	return sdk.ChainPostDecorators(
		NewGasRefundDecorator(ek),
	)
}

type EVMKeeper interface {
	GasRefundRatio(context.Context) (math.LegacyDec, error)
	ERC20Keeper() evmtypes.IERC20Keeper
	TxUtils() evmtypes.TxUtils
}
