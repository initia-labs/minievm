package indexer_test

import (
	"math/big"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	storetypes "cosmossdk.io/store/types"

	evmtypes "github.com/initia-labs/minievm/x/evm/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func Test_ListenFinalizeBlock(t *testing.T) {
	app, indexer, addrs, privKeys := setupIndexer(t)
	defer app.Close()

	tx, evmTxHash := generateCreateERC20Tx(t, app, privKeys[0])
	finalizeReq, finalizeRes := executeTxs(t, app, tx)
	checkTxResult(t, finalizeRes.TxResults[0], true)

	events := finalizeRes.TxResults[0].Events
	createEvent := events[len(events)-3]
	require.Equal(t, evmtypes.EventTypeContractCreated, createEvent.GetType())

	contractAddr, err := hexutil.Decode(createEvent.Attributes[0].Value)
	require.NoError(t, err)

	// listen finalize block
	ctx, err := app.CreateQueryContext(0, false)
	require.NoError(t, err)

	err = indexer.ListenFinalizeBlock(ctx.WithBlockGasMeter(storetypes.NewInfiniteGasMeter()), *finalizeReq, *finalizeRes)
	require.NoError(t, err)

	// check the tx is indexed
	evmTx, err := indexer.TxByHash(ctx, evmTxHash)
	require.NoError(t, err)
	require.NotNil(t, evmTx)

	// mint 1_000_000 tokens to the first address
	tx, evmTxHash = generateMintERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[0], new(big.Int).SetUint64(1_000_000_000_000))
	finalizeReq, finalizeRes = executeTxs(t, app, tx)
	checkTxResult(t, finalizeRes.TxResults[0], true)

	// listen finalize block
	ctx, err = app.CreateQueryContext(0, false)
	require.NoError(t, err)

	err = indexer.ListenFinalizeBlock(ctx.WithBlockGasMeter(storetypes.NewInfiniteGasMeter()), *finalizeReq, *finalizeRes)
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
	app, indexer, _, privKeys := setupIndexer(t)
	defer app.Close()

	blockChan, logsChan, pendChan := indexer.Subscribe()
	close(pendChan)

	tx, evmTxHash := generateCreateERC20Tx(t, app, privKeys[0])
	finalizeReq, finalizeRes := executeTxs(t, app, tx)
	checkTxResult(t, finalizeRes.TxResults[0], true)

	events := finalizeRes.TxResults[0].Events
	createEvent := events[len(events)-3]
	require.Equal(t, evmtypes.EventTypeContractCreated, createEvent.GetType())

	contractAddr, err := hexutil.Decode(createEvent.Attributes[0].Value)
	require.NoError(t, err)

	// listen finalize block
	ctx, err := app.CreateQueryContext(0, false)
	require.NoError(t, err)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for {
			select {
			case block := <-blockChan:
				require.NotNil(t, block)
				require.Equal(t, finalizeReq.Height, block.Number.Int64())
				wg.Done()
			case logs := <-logsChan:
				require.NotNil(t, logs)

				for _, log := range logs {
					if log.Address == common.BytesToAddress(contractAddr) {
						require.Equal(t, evmTxHash, log.TxHash)
						require.Equal(t, uint64(finalizeReq.Height), log.BlockNumber)
						wg.Done()
					}
				}
			case <-time.After(10 * time.Second):
				t.Error("timeout waiting for pending transaction")
				wg.Done()
			}
		}
	}()

	err = indexer.ListenFinalizeBlock(ctx.WithBlockGasMeter(storetypes.NewInfiniteGasMeter()), *finalizeReq, *finalizeRes)
	require.NoError(t, err)

	wg.Wait()
}

func Test_ListenFinalizeBlock_ContractCreation(t *testing.T) {
	app, indexer, _, privKeys := setupIndexer(t)
	defer app.Close()

	tx, evmTxHash := generateCreateInitiaERC20Tx(t, app, privKeys[0])
	finalizeReq, finalizeRes := executeTxs(t, app, tx)
	checkTxResult(t, finalizeRes.TxResults[0], true)

	events := finalizeRes.TxResults[0].Events
	createEvent := events[len(events)-3]
	require.Equal(t, evmtypes.EventTypeContractCreated, createEvent.GetType())

	contractAddr, err := hexutil.Decode(createEvent.Attributes[0].Value)
	require.NoError(t, err)

	// listen finalize block
	ctx, err := app.CreateQueryContext(0, false)
	require.NoError(t, err)

	err = indexer.ListenFinalizeBlock(ctx.WithBlockGasMeter(storetypes.NewInfiniteGasMeter()), *finalizeReq, *finalizeRes)
	require.NoError(t, err)

	// check the tx is indexed
	receipt, err := indexer.TxReceiptByHash(ctx, evmTxHash)
	require.NoError(t, err)
	require.NotNil(t, receipt)

	// contract creation should have contract address in receipt
	require.Equal(t, contractAddr, receipt.ContractAddress.Bytes())
}
