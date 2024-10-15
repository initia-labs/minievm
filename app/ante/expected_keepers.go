package ante

import (
	"context"

	evmtypes "github.com/initia-labs/minievm/x/evm/types"
)

type EVMKeeper interface {
	GetFeeDenom(ctx context.Context) (string, error)
	TxUtils() evmtypes.TxUtils
}
