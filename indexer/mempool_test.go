package indexer_test

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/types/mempool"
)

func Test_Mempool_Subscribe(t *testing.T) {
	app, indexer, _, privKeys := setupIndexer(t)
	defer app.Close()

	blockChan, logsChan, pendChan := indexer.Subscribe()
	close(blockChan)
	close(logsChan)

	tx, evmTxHash := generateCreateERC20Tx(t, app, privKeys[0])

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		pendingTx := <-pendChan
		require.NotNil(t, pendingTx)
		require.Equal(t, evmTxHash, pendingTx.Hash)
		wg.Done()
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
