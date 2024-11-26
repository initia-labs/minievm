package indexer_test

import (
	"sync"
	"testing"
	"time"

	"github.com/initia-labs/minievm/tests"
	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/types/mempool"
)

func Test_Mempool_Subscribe(t *testing.T) {
	app, _, privKeys := tests.CreateApp(t)
	indexer := app.EVMIndexer()
	defer app.Close()

	blockChan, logsChan, pendChan := indexer.Subscribe()
	close(blockChan)
	close(logsChan)

	tx, evmTxHash := tests.GenerateCreateERC20Tx(t, app, privKeys[0])

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		select {
		case pendingTx := <-pendChan:
			require.NotNil(t, pendingTx)
			require.Equal(t, evmTxHash, pendingTx.Hash)
			wg.Done()
		case <-time.After(5 * time.Second):
			t.Error("timeout waiting for pending transaction")
			wg.Done()
		}
	}()

	noopMempool := &mempool.NoOpMempool{}
	mempool := indexer.MempoolWrapper(noopMempool)

	// insert tx into mempool
	ctx, err := app.CreateQueryContext(0, false)
	require.NoError(t, err)
	err = mempool.Insert(ctx, tx)
	require.NoError(t, err)

	rpcTx := indexer.TxInMempool(evmTxHash)
	require.Equal(t, evmTxHash, rpcTx.Hash)

	wg.Wait()
}
