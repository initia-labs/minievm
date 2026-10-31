package tests

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Every call frame takes a StateDB snapshot, which is a cache context layer. A
// frame that returns successfully must release its snapshot, otherwise the
// layers stack up for the rest of the tx and Commit has to copy every dirty key
// through all of them.
func Test_SequentialCalls_ReleaseSnapshots(t *testing.T) {
	app, addrs, privKeys := CreateApp(t)
	key := &ecdsaKey{k: privKeys[0], addr: addrs[0]}
	contract := deployLoop(t, app, key)

	r := runLoop(t, app, key, contract, 10_000_000, 3_000, 3_000, 0, 0)
	require.Equal(t, uint32(0), r.code, r.log)
	require.Less(t, r.allocMB, 1024.0, "call frame snapshots are not released (allocated %.0f MB)", r.allocMB)
}
