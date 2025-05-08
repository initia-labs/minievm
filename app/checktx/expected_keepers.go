package checktx

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type NonceGetter interface {
	GetSequence(ctx context.Context, addr sdk.AccAddress) (uint64, error)
}

type BalanceGetter interface {
	GetBalance(ctx context.Context, addr sdk.AccAddress, denom string) sdk.Coin
}

type FeeDenomGetter interface {
	GetFeeDenom(ctx context.Context) (string, error)
}

type ContextGetter interface {
	GetContextForCheckTx(tx []byte) sdk.Context
}
