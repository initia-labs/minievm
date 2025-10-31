package state

import "github.com/initia-labs/minievm/x/evm/types"

type Snapshot struct {
	ctx    Context
	commit func()
}

func NewSnapshot(ctx Context) *Snapshot {
	// snapshot maintains a list of execute requests
	// and keep track of the parent execute requests
	ctx.Context = ctx.Context.
		WithValue(
			types.CONTEXT_KEY_PARENT_EXECUTE_REQUESTS,
			ctx.Context.Value(types.CONTEXT_KEY_EXECUTE_REQUESTS)).
		WithValue(types.CONTEXT_KEY_EXECUTE_REQUESTS, &[]types.ExecuteRequest{})

	// create cache context
	cacheCtx, commit := ctx.CacheContext()
	return &Snapshot{
		ctx:    cacheCtx,
		commit: commit,
	}
}

func (s *Snapshot) Commit() {
	// propagate execute requests to parent
	parentExecuteRequests := s.ctx.Value(types.CONTEXT_KEY_PARENT_EXECUTE_REQUESTS).(*[]types.ExecuteRequest)
	executeRequests := s.ctx.Value(types.CONTEXT_KEY_EXECUTE_REQUESTS).(*[]types.ExecuteRequest)
	*parentExecuteRequests = append(*parentExecuteRequests, *executeRequests...)

	// commit cache
	s.commit()
}

// for mock testing
func (s *Snapshot) Context() Context {
	return s.ctx
}
