package posthandler

import (
	"context"

	"cosmossdk.io/log"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	evmtypes "github.com/initia-labs/minievm/x/evm/types"
)

// NewPostHandler returns a new sdk.PostHandler that is composed of the sdk.ChainPostDecorators
func NewPostHandler(
	logger log.Logger,
	ek EVMKeeper,
) sdk.PostHandler {
	return sdk.ChainPostDecorators(
		NewGasRefundDecorator(logger, ek),
	)
}

type EVMKeeper interface {
	GasRefundRatio(context.Context) (math.LegacyDec, error)
	ERC20Keeper() evmtypes.IERC20Keeper
	TxUtils() evmtypes.TxUtils
}
