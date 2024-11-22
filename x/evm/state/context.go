package state

import (
	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Context struct {
	sdk.Context

	memStore    storetypes.MultiStore
	memStoreKey storetypes.StoreKey
}

func NewContext(ctx sdk.Context) Context {
	memStore, memStoreKey := newMemStore()
	return Context{
		Context:     ctx,
		memStore:    memStore,
		memStoreKey: memStoreKey,
	}
}

func (c Context) WithSDKContext(sdkCtx sdk.Context) Context {
	c.Context = sdkCtx
	return c
}

func (c Context) WithMemStore(memStore storetypes.MultiStore) Context {
	c.memStore = memStore
	return c
}

func (c Context) CacheContext() (cc Context, writeCache func()) {
	cacheCtx, commit := c.Context.CacheContext()
	cacheMemStore := c.memStore.CacheMultiStore()

	cc = c.WithSDKContext(cacheCtx).WithMemStore(cacheMemStore)
	writeCache = func() {
		commit()
		cacheMemStore.Write()
	}

	return cc, writeCache
}
