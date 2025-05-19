package indexer

import (
	"testing"

	dbm "github.com/cosmos/cosmos-db"
	"github.com/stretchr/testify/require"
)

func Test_StoreIO(t *testing.T) {
	db := dbm.NewMemDB()
	store := NewCacheStoreWithBatch(db, 100)

	key := []byte("key")
	value := []byte("value")

	// case 1. key not exists
	ok, err := store.Has(key)
	require.NoError(t, err)
	require.False(t, ok)

	bz, err := store.Get(key)
	require.NoError(t, err)
	require.Nil(t, bz)

	// case 2. set key
	err = store.Set(key, value)
	require.NoError(t, err)

	ok, err = store.Has(key)
	require.NoError(t, err)
	require.True(t, ok)

	bz, err = store.Get(key)
	require.NoError(t, err)
	require.Equal(t, value, bz)

	// case 3. delete key
	err = store.Delete(key)
	require.NoError(t, err)

	ok, err = store.Has(key)
	require.NoError(t, err)
	require.False(t, ok)

	bz, err = store.Get(key)
	require.NoError(t, err)
	require.Nil(t, bz)
}
