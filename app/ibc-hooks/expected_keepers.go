package evm_hooks

import "context"

type OPChildKeeper interface {
	GetIBCToL2DenomMap(ctx context.Context, ibcDenom string) (string, error)
	HasIBCToL2DenomMap(ctx context.Context, ibcDenom string) (bool, error)
}
