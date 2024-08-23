package state_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func Test_SnapshotRevert(t *testing.T) {
	ctx, input := createDefaultTestInput(t)

	stateDB, err := input.EVMKeeper.NewStateDB(ctx)
	require.NoError(t, err)

	_, _, addr := keyPubAddr()
	require.False(t, stateDB.AddressInAccessList(common.BytesToAddress(addr)))

	// take snapshot
	sid := stateDB.Snapshot()

	stateDB.AddAddressToAccessList(common.BytesToAddress(addr))
	require.True(t, stateDB.AddressInAccessList(common.BytesToAddress(addr)))

	// revert to snapshot
	stateDB.RevertToSnapshot(sid)
	require.False(t, stateDB.AddressInAccessList(common.BytesToAddress(addr)))
}
