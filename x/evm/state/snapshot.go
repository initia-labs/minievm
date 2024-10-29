package state

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Snapshot struct {
	ctx    sdk.Context
	commit func()
}

func NewSnapshot(ctx context.Context) *Snapshot {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	cacheCtx, commit := sdkCtx.CacheContext()
	return &Snapshot{
		ctx:    cacheCtx,
		commit: commit,
	}
}

func (s *Snapshot) Commit() {
	s.commit()
}

// for mock testing
func (s *Snapshot) Context() sdk.Context {
	return s.ctx
}
