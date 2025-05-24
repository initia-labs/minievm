package indexer_test

import (
	"context"
	"math/big"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"

	comettypes "github.com/cometbft/cometbft/types"

	"cosmossdk.io/collections"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

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

	// 1. Test that Ethereum transactions are properly indexed
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

	// check the tx is indexed
	ih, err := indexer.GetLastIndexedHeight(ctx)
	require.NoError(t, err)
	require.Equal(t, finalizeReq.Height, int64(ih))

	// 2. Test that Cosmos transactions which generate EVM logs are properly indexed
	// This verifies the indexer handles non-Ethereum transactions that still produce EVM events
	feeDenom, err := app.EVMKeeper.GetFeeDenom(ctx)
	require.NoError(t, err)
	tx = tests.GenerateCosmosTx(t, app, privKeys[0], []sdk.Msg{
		&banktypes.MsgSend{
			FromAddress: sdk.AccAddress(addrs[0].Bytes()).String(),
			ToAddress:   sdk.AccAddress(addrs[1].Bytes()).String(),
			Amount:      sdk.NewCoins(sdk.NewCoin(feeDenom, math.NewIntFromBigInt(new(big.Int).SetUint64(1_000_000_000_000)))),
		},
	})

	txBytes, err := app.TxConfig().TxEncoder()(tx)
	require.NoError(t, err)
	cosmosTxHash := comettypes.Tx(txBytes).Hash()

	finalizeReq, finalizeRes = tests.ExecuteTxs(t, app, tx)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	// check the block header is indexed
	header, err = indexer.BlockHeaderByNumber(ctx, uint64(finalizeReq.Height))
	require.NoError(t, err)
	require.NotNil(t, header)
	require.Equal(t, finalizeReq.Height, header.Number.Int64())

	// check the tx is indexed
	evmTxHash, err = indexer.TxHashByCosmosTxHash(ctx, cosmosTxHash)
	require.NoError(t, err)
	require.NotNil(t, evmTxHash)

	evmTx, err = indexer.TxByHash(ctx, evmTxHash)
	require.NoError(t, err)
	require.NotNil(t, evmTx)

	// check the tx is indexed
	ih, err = indexer.GetLastIndexedHeight(ctx)
	require.NoError(t, err)
	require.Equal(t, finalizeReq.Height, int64(ih))

	// 3. Test that Cosmos transactions which do not generate EVM logs are not indexed
	authzMsg, err := authz.NewMsgGrant(sdk.AccAddress(addrs[0].Bytes()), sdk.AccAddress(addrs[1].Bytes()), authz.NewGenericAuthorization("/cosmos.bank.v1beta1.MsgSend"), nil)
	require.NoError(t, err)

	tx = tests.GenerateCosmosTx(t, app, privKeys[0], []sdk.Msg{authzMsg})

	txBytes, err = app.TxConfig().TxEncoder()(tx)
	require.NoError(t, err)
	cosmosTxHash = comettypes.Tx(txBytes).Hash()

	finalizeReq, finalizeRes = tests.ExecuteTxs(t, app, tx)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	// check the block header is indexed
	header, err = indexer.BlockHeaderByNumber(ctx, uint64(finalizeReq.Height))
	require.NoError(t, err)
	require.NotNil(t, header)
	require.Equal(t, finalizeReq.Height, header.Number.Int64())

	// check the tx is indexed
	ih, err = indexer.GetLastIndexedHeight(ctx)
	require.NoError(t, err)
	require.Equal(t, finalizeReq.Height, int64(ih))

	// check the tx is not indexed
	_, err = indexer.TxHashByCosmosTxHash(ctx, cosmosTxHash)
	require.ErrorIs(t, err, collections.ErrNotFound)

	// 4. Test that failed Cosmo transactions are not indexed
	tx = tests.GenerateCosmosTx(t, app, privKeys[0], []sdk.Msg{
		&banktypes.MsgSend{
			FromAddress: sdk.AccAddress(addrs[0].Bytes()).String(),
			ToAddress:   sdk.AccAddress(addrs[1].Bytes()).String(),
			Amount:      sdk.NewCoins(sdk.NewCoin(feeDenom, math.ZeroInt())),
		},
	})

	txBytes, err = app.TxConfig().TxEncoder()(tx)
	require.NoError(t, err)
	cosmosTxHash = comettypes.Tx(txBytes).Hash()

	finalizeReq, finalizeRes = tests.ExecuteTxs(t, app, tx)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], false)

	// check the block header is indexed
	header, err = indexer.BlockHeaderByNumber(ctx, uint64(finalizeReq.Height))
	require.NoError(t, err)
	require.NotNil(t, header)

	// check the tx is indexed
	ih, err = indexer.GetLastIndexedHeight(ctx)
	require.NoError(t, err)
	require.Equal(t, finalizeReq.Height, int64(ih))

	// check the tx is not indexed
	_, err = indexer.TxHashByCosmosTxHash(ctx, cosmosTxHash)
	require.ErrorIs(t, err, collections.ErrNotFound)
}

func Test_ListenFinalizeBlock_Subscribe(t *testing.T) {
	app, _, privKeys := tests.CreateApp(t)
	indexer := app.EVMIndexer()
	defer app.Close()

	blockChan, logsChan, _ := indexer.Subscribe()

	tx, evmTxHash := tests.GenerateCreateERC20Tx(t, app, privKeys[0])

	ctx, done := context.WithCancel(context.Background())
	reqHeight := app.LastBlockHeight() + 1
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for {
			select {
			case block := <-blockChan:
				if block == nil || block.Number.Int64() < reqHeight {
					continue
				}

				require.NotNil(t, block)
				require.Equal(t, reqHeight, block.Number.Int64())
				wg.Done()
			case logs := <-logsChan:
				require.NotNil(t, logs)
				if logs[0].BlockNumber < uint64(reqHeight) {
					continue
				}

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
