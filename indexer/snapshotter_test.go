package indexer_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/initia-labs/minievm/tests"
	evmtypes "github.com/initia-labs/minievm/x/evm/types"
)

func Test_Snapshotter(t *testing.T) {
	app, _, privKeys := tests.CreateApp(t)
	indexer := app.EVMIndexer()
	defer app.Close()

	tx, _ := tests.GenerateCreateERC20Tx(t, app, privKeys[0])
	finalizeReq, finalizeRes := tests.ExecuteTxs(t, app, tx)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	events := finalizeRes.TxResults[0].Events
	createEvent := events[len(events)-3]
	require.Equal(t, evmtypes.EventTypeContractCreated, createEvent.GetType())

	height := uint64(finalizeReq.Height)

	snapshotBz := []byte{}
	payloadWriter := func(payload []byte) error {
		snapshotBz = append(snapshotBz, payload...)
		return nil
	}
	payloadReader := func() ([]byte, error) {
		return snapshotBz, nil
	}

	// write snapshot
	err := indexer.SnapshotExtension(height, payloadWriter)
	require.NoError(t, err)

	// restore snapshot on another app
	app2, _, _ := tests.CreateApp(t)
	indexer2 := app2.EVMIndexer()

	// invalid format
	err = indexer2.RestoreExtension(height, 2, payloadReader)
	require.Error(t, err)

	// restore snapshot
	err = indexer2.RestoreExtension(height, 1, payloadReader)
	require.NoError(t, err)

	// create another snapshot and compare with the previous one
	snapshotBz2 := []byte{}
	payloadWriter2 := func(payload []byte) error {
		snapshotBz2 = append(snapshotBz2, payload...)
		return nil
	}

	err = indexer2.SnapshotExtension(height, payloadWriter2)
	require.NoError(t, err)
	require.Equal(t, snapshotBz, snapshotBz2)
}
