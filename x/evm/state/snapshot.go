package state

type Snapshot struct {
	ctx    Context
	commit func()
}

func NewSnapshot(ctx Context) *Snapshot {
	cacheCtx, commit := ctx.CacheContext()
	return &Snapshot{
		ctx:    cacheCtx,
		commit: commit,
	}
}

func (s *Snapshot) Commit() {
	s.commit()
}

// for mock testing
func (s *Snapshot) Context() Context {
	return s.ctx
}
