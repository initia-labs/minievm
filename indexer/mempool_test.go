package indexer_test

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/initia-labs/minievm/tests"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/types/mempool"

	coretypes "github.com/ethereum/go-ethereum/core/types"
)

func Test_Mempool_Subscribe(t *testing.T) {
	app, _, privKeys := tests.CreateApp(t)
	indexer := app.EVMIndexer()
	defer app.Close()

	blockChan, logsChan, pendChan := indexer.Subscribe()
	go consumeBlockLogsChan(blockChan, logsChan, 5)

	tx, evmTxHash := tests.GenerateCreateERC20Tx(t, app, privKeys[0])

	ctx, done := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		select {
		case pendingTx := <-pendChan:
			require.NotNil(t, pendingTx)
			require.Equal(t, evmTxHash, pendingTx.Hash)
			wg.Done()
		case <-ctx.Done():
			return
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
	done()
}

func Test_Mempool_TriggerFlushQueuedTxs(t *testing.T) {
	app, addrs, privKeys := tests.CreateApp(t)
	indexer := app.EVMIndexer()
	defer app.Close()

	received := make(map[string]uint64)

	wg := sync.WaitGroup{}
	wg.Add(1)
	indexer.RegisterFlushQueuedTxs(func(senderHex string, accSeq uint64) error {
		received[senderHex] = accSeq
		wg.Done()
		return nil
	})

	tx, _ := tests.GenerateCreateERC20Tx(t, app, privKeys[0])

	noopMempool := &mempool.NoOpMempool{}
	mempool := indexer.MempoolWrapper(noopMempool)

	// insert tx into mempool
	ctx, err := app.CreateQueryContext(0, false)
	require.NoError(t, err)
	seq, err := app.AccountKeeper.GetSequence(ctx, addrs[0].Bytes())
	require.NoError(t, err)
	err = mempool.Insert(ctx, tx)
	require.NoError(t, err)

	wg.Wait()

	require.Len(t, received, 1)
	require.Equal(t, uint64(seq+1), received[addrs[0].Hex()])
}

func consumeBlockLogsChan(blockCh <-chan *coretypes.Header, logChan <-chan []*coretypes.Log, duration int) {
	timer := time.NewTimer(time.Second * time.Duration(duration))

	for {
		select {
		case <-blockCh:
		case <-logChan:
		case <-timer.C:
			return
		}
	}
}
