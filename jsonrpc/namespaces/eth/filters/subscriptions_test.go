package filters_test

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	coretypes "github.com/ethereum/go-ethereum/core/types"
	ethfilters "github.com/ethereum/go-ethereum/eth/filters"
	"github.com/ethereum/go-ethereum/rpc"

	rpctypes "github.com/initia-labs/minievm/jsonrpc/types"
	"github.com/initia-labs/minievm/tests"
	evmtypes "github.com/initia-labs/minievm/x/evm/types"
)

func Test_NewPendingTransactions_FullTx(t *testing.T) {
	input := tests.CreateAppWithJSONRPC(t)
	defer input.App.Close()

	// wait indexer to be ready
	time.Sleep(3 * time.Second)

	client, err := rpc.Dial("ws://" + input.AddressWS)
	require.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ch := make(chan *rpctypes.RPCTransaction)
	sub, err := client.EthSubscribe(ctx, ch, "newPendingTransactions", true)
	require.NoError(t, err)
	require.NotEmpty(t, sub)

	app, backend, privKeys := input.App, input.Backend, input.PrivKeys

	queryCtx, err := app.CreateQueryContext(0, false)
	require.NoError(t, err)

	tx, txHash1 := tests.GenerateCreateERC20Tx(t, app, privKeys[0])
	evmTx, _, err := app.EVMKeeper.TxUtils().ConvertCosmosTxToEthereumTx(queryCtx, tx)
	require.NoError(t, err)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case tx := <-ch:
				require.NotNil(t, tx)
				require.Equal(t, txHash1, tx.Hash)
				wg.Done()
			}
		}
	}()

	txBz, err := evmTx.MarshalBinary()
	require.NoError(t, err)

	_, err = backend.SendRawTransaction(txBz)
	require.NoError(t, err)

	wg.Wait()
	sub.Unsubscribe()
}

func Test_NewHeads(t *testing.T) {
	input := tests.CreateAppWithJSONRPC(t)
	defer input.App.Close()

	// wait indexer to be ready
	time.Sleep(3 * time.Second)

	client, err := rpc.Dial("ws://" + input.AddressWS)
	require.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ch := make(chan *coretypes.Header)
	sub, err := client.EthSubscribe(ctx, ch, "newHeads")
	require.NoError(t, err)
	require.NotEmpty(t, sub)

	app, backend, privKeys := input.App, input.Backend, input.PrivKeys

	wg := sync.WaitGroup{}  // wait to receive header from channel
	wg2 := sync.WaitGroup{} // wait for header to be set
	wg.Add(1)
	wg2.Add(1)
	var expectedHeader *coretypes.Header
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case header := <-ch:
				wg2.Wait()
				require.NotNil(t, header)
				require.Equal(t, expectedHeader, header)
				wg.Done()
			}
		}
	}()

	tx, _ := tests.GenerateCreateERC20Tx(t, app, privKeys[0])
	_, finalizeRes := tests.ExecuteTxs(t, app, tx)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)
	header, err := backend.GetHeaderByNumber(rpc.BlockNumber(app.LastBlockHeight()))
	require.NoError(t, err)
	expectedHeader = header
	wg2.Done()
	wg.Wait()
	sub.Unsubscribe()
}

func Test_Logs(t *testing.T) {
	input := tests.CreateAppWithJSONRPC(t)
	defer input.App.Close()

	// wait indexer to be ready
	time.Sleep(3 * time.Second)

	client, err := rpc.Dial("ws://" + input.AddressWS)
	require.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ch := make(chan *coretypes.Log)
	sub, err := client.EthSubscribe(ctx, ch, "logs", ethfilters.FilterCriteria{})
	require.NoError(t, err)
	require.NotEmpty(t, sub)

	app, backend, privKeys := input.App, input.Backend, input.PrivKeys

	wg := sync.WaitGroup{}  // wait to receive logs from channel
	wg2 := sync.WaitGroup{} // wait for logs to be set
	wg.Add(1)
	wg2.Add(1)
	var expectedLogs []*coretypes.Log
	var logs []*coretypes.Log
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case log := <-ch:
				wg2.Wait()
				logs = append(logs, log)
				if len(logs) == len(expectedLogs) {
					require.Equal(t, expectedLogs, logs)
					wg.Done()
				}
			}
		}
	}()

	tx, _ := tests.GenerateCreateERC20Tx(t, app, privKeys[0])
	_, finalizeRes := tests.ExecuteTxs(t, app, tx)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	events := finalizeRes.TxResults[0].Events
	createEvent := events[len(events)-3]
	require.Equal(t, evmtypes.EventTypeContractCreated, createEvent.GetType())

	expectedLogs, err = backend.GetLogsByHeight(uint64(app.LastBlockHeight()))
	require.NoError(t, err)

	wg2.Done()
	wg.Wait()
	sub.Unsubscribe()
}
