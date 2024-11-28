package indexer_test

import (
	"context"
	"math/big"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/initia-labs/minievm/tests"
	evmtypes "github.com/initia-labs/minievm/x/evm/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func Test_ListenFinalizeBlock(t *testing.T) {
	app, addrs, privKeys := tests.CreateApp(t)
	indexer := app.EVMIndexer()
	defer app.Close()

	tx, evmTxHash := tests.GenerateCreateERC20Tx(t, app, privKeys[0])
	_, finalizeRes := tests.ExecuteTxs(t, app, tx)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	events := finalizeRes.TxResults[0].Events
	createEvent := events[len(events)-3]
	require.Equal(t, evmtypes.EventTypeContractCreated, createEvent.GetType())

	contractAddr, err := hexutil.Decode(createEvent.Attributes[0].Value)
	require.NoError(t, err)

	// listen finalize block
	ctx, err := app.CreateQueryContext(0, false)
	require.NoError(t, err)

	// check the tx is indexed
	evmTx, err := indexer.TxByHash(ctx, evmTxHash)
	require.NoError(t, err)
	require.NotNil(t, evmTx)

	// mint 1_000_000 tokens to the first address
	tx, evmTxHash = tests.GenerateMintERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[0], new(big.Int).SetUint64(1_000_000_000_000))
	finalizeReq, finalizeRes := tests.ExecuteTxs(t, app, tx)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	// listen finalize block
	ctx, err = app.CreateQueryContext(0, false)
	require.NoError(t, err)

	// check the tx is indexed
	evmTx, err = indexer.TxByHash(ctx, evmTxHash)
	require.NoError(t, err)
	require.NotNil(t, evmTx)

	// check the block header is indexed
	header, err := indexer.BlockHeaderByNumber(ctx, uint64(finalizeReq.Height))
	require.NoError(t, err)
	require.NotNil(t, header)
	require.Equal(t, finalizeReq.Height, header.Number.Int64())

}

func Test_ListenFinalizeBlock_Subscribe(t *testing.T) {
	app, _, privKeys := tests.CreateApp(t)
	indexer := app.EVMIndexer()
	defer app.Close()

	blockChan, logsChan, pendChan := indexer.Subscribe()
	defer close(blockChan)
	defer close(logsChan)
	defer close(pendChan)

	tx, evmTxHash := tests.GenerateCreateERC20Tx(t, app, privKeys[0])

	ctx, done := context.WithCancel(context.Background())
	reqHeight := app.LastBlockHeight() + 1
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for {
			select {
			case block := <-blockChan:
				require.NotNil(t, block)
				require.Equal(t, reqHeight, block.Number.Int64())
				wg.Done()
			case logs := <-logsChan:
				require.NotNil(t, logs)

				for _, log := range logs {
					require.Equal(t, evmTxHash, log.TxHash)
					require.Equal(t, uint64(reqHeight), log.BlockNumber)
				}

				wg.Done()
			case <-ctx.Done():
				return
			}
		}
	}()

	finalizeReq, finalizeRes := tests.ExecuteTxs(t, app, tx)
	require.Equal(t, reqHeight, finalizeReq.Height)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	events := finalizeRes.TxResults[0].Events
	createEvent := events[len(events)-3]
	require.Equal(t, evmtypes.EventTypeContractCreated, createEvent.GetType())

	wg.Wait()
	done()
}

func Test_ListenFinalizeBlock_ContractCreation(t *testing.T) {
	app, _, privKeys := tests.CreateApp(t)
	indexer := app.EVMIndexer()
	defer app.Close()

	tx, evmTxHash := tests.GenerateCreateInitiaERC20Tx(t, app, privKeys[0])
	_, finalizeRes := tests.ExecuteTxs(t, app, tx)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	events := finalizeRes.TxResults[0].Events
	createEvent := events[len(events)-3]
	require.Equal(t, evmtypes.EventTypeContractCreated, createEvent.GetType())

	contractAddr, err := hexutil.Decode(createEvent.Attributes[0].Value)
	require.NoError(t, err)

	// check the tx is indexed
	ctx, err := app.CreateQueryContext(0, false)
	require.NoError(t, err)

	receipt, err := indexer.TxReceiptByHash(ctx, evmTxHash)
	require.NoError(t, err)
	require.NotNil(t, receipt)

	// contract creation should have contract address in receipt
	require.Equal(t, contractAddr, receipt.ContractAddress.Bytes())
}
