package filters_test

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	ethfilters "github.com/ethereum/go-ethereum/eth/filters"
	"github.com/ethereum/go-ethereum/rpc"

	minitiaapp "github.com/initia-labs/minievm/app"
	"github.com/initia-labs/minievm/indexer"
	"github.com/initia-labs/minievm/jsonrpc/backend"
	"github.com/initia-labs/minievm/jsonrpc/config"
	"github.com/initia-labs/minievm/jsonrpc/namespaces/eth/filters"
	rpctypes "github.com/initia-labs/minievm/jsonrpc/types"
	"github.com/initia-labs/minievm/tests"
	evmtypes "github.com/initia-labs/minievm/x/evm/types"
)

type testInput struct {
	app       *minitiaapp.MinitiaApp
	indexer   indexer.EVMIndexer
	backend   *backend.JSONRPCBackend
	addrs     []common.Address
	privKeys  []*ecdsa.PrivateKey
	cometRPC  *tests.MockCometRPC
	filterAPI *filters.FilterAPI
}

func setupFilterAPI(t *testing.T) testInput {
	app, addrs, privKeys := tests.CreateApp(t)
	indexer := app.EVMIndexer()

	ctx := context.Background()
	svrCtx := server.NewDefaultContext()
	clientCtx := client.Context{}.WithCodec(app.AppCodec()).
		WithInterfaceRegistry(app.AppCodec().InterfaceRegistry()).
		WithTxConfig(app.TxConfig()).
		WithLegacyAmino(app.LegacyAmino()).
		WithAccountRetriever(authtypes.AccountRetriever{})

	cfg := config.DefaultJSONRPCConfig()
	cfg.Enable = true
	cfg.FilterTimeout = 10 * time.Second

	mockCometRPC := tests.NewMockCometRPC(app.BaseApp)
	clientCtx = clientCtx.WithClient(mockCometRPC)

	backend, err := backend.NewJSONRPCBackend(ctx, app, app.Logger(), svrCtx, clientCtx, cfg)
	require.NoError(t, err)

	filterAPI := filters.NewFilterAPI(ctx, app, backend, app.Logger())

	return testInput{
		app:       app,
		indexer:   indexer,
		backend:   backend,
		addrs:     addrs,
		privKeys:  privKeys,
		cometRPC:  mockCometRPC,
		filterAPI: filterAPI,
	}
}

func Test_NewPendingTransactionFilter_FullTx(t *testing.T) {
	input := setupFilterAPI(t)
	defer input.app.Close()

	// wait indexer to be ready
	time.Sleep(3 * time.Second)

	fullTx := true
	filterID, err := input.filterAPI.NewPendingTransactionFilter(&fullTx)
	require.NoError(t, err)
	require.NotEmpty(t, filterID)

	app, backend, addrs, privKeys := input.app, input.backend, input.addrs, input.privKeys

	ctx, err := app.CreateQueryContext(0, false)
	require.NoError(t, err)

	tx, _ := tests.GenerateCreateERC20Tx(t, app, privKeys[0])
	evmTx, _, err := app.EVMKeeper.TxUtils().ConvertCosmosTxToEthereumTx(ctx, tx)
	require.NoError(t, err)

	txBz, err := evmTx.MarshalBinary()
	require.NoError(t, err)
	txHash1, err := backend.SendRawTransaction(txBz)
	require.NoError(t, err)

	_, finalizeRes := tests.ExecuteTxs(t, app, tx)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	events := finalizeRes.TxResults[0].Events
	createEvent := events[len(events)-3]
	require.Equal(t, evmtypes.EventTypeContractCreated, createEvent.GetType())

	contractAddr, err := hexutil.Decode(createEvent.Attributes[0].Value)
	require.NoError(t, err)

	// mint 1_000_000 tokens to the first address
	tx2, _ := tests.GenerateMintERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[0], new(big.Int).SetUint64(1_000_000_000_000))
	evmTx2, _, err := app.EVMKeeper.TxUtils().ConvertCosmosTxToEthereumTx(ctx, tx2)
	require.NoError(t, err)

	txBz, err = evmTx2.MarshalBinary()
	require.NoError(t, err)
	txHash2, err := backend.SendRawTransaction(txBz)
	require.NoError(t, err)

	_, finalizeRes = tests.ExecuteTxs(t, app, tx2)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	rpcTx1, err := backend.GetTransactionByHash(txHash1)
	require.NoError(t, err)
	rpcTx2, err := backend.GetTransactionByHash(txHash2)
	require.NoError(t, err)

	// wait txs to be indexed
	time.Sleep(1 * time.Second)

	// there should be 2 changes
	changes, err := input.filterAPI.GetFilterChanges(filterID)
	require.NoError(t, err)
	require.Len(t, changes, 2)

	// to compare with pending tx filter, we need to remove block hash, block number, and transaction index
	rpcTx1.BlockHash = nil
	rpcTx1.BlockNumber = nil
	rpcTx1.TransactionIndex = nil
	rpcTx2.BlockHash = nil
	rpcTx2.BlockNumber = nil
	rpcTx2.TransactionIndex = nil

	res := []string{}
	for i := 0; i < 2; i++ {
		rpcTx := changes.([]*rpctypes.RPCTransaction)[i]
		res = append(res, rpcTx.String())
	}

	require.Equal(t, []string{rpcTx1.String(), rpcTx2.String()}, res)

	// uninstall filter
	found := input.filterAPI.UninstallFilter(filterID)
	require.True(t, found)

	// more changes should not be returned
	_, err = input.filterAPI.GetFilterChanges(filterID)
	require.ErrorContains(t, err, "filter not found")
}

func Test_NewPendingTransactionFilter(t *testing.T) {
	input := setupFilterAPI(t)
	defer input.app.Close()

	// wait indexer to be ready
	time.Sleep(3 * time.Second)

	fullTx := false
	filterID, err := input.filterAPI.NewPendingTransactionFilter(&fullTx)
	require.NoError(t, err)
	require.NotEmpty(t, filterID)

	app, backend, addrs, privKeys := input.app, input.backend, input.addrs, input.privKeys

	ctx, err := app.CreateQueryContext(0, false)
	require.NoError(t, err)

	tx, _ := tests.GenerateCreateERC20Tx(t, app, privKeys[0])
	evmTx, _, err := app.EVMKeeper.TxUtils().ConvertCosmosTxToEthereumTx(ctx, tx)
	require.NoError(t, err)

	txBz, err := evmTx.MarshalBinary()
	require.NoError(t, err)
	txHash1, err := backend.SendRawTransaction(txBz)
	require.NoError(t, err)

	_, finalizeRes := tests.ExecuteTxs(t, app, tx)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	events := finalizeRes.TxResults[0].Events
	createEvent := events[len(events)-3]
	require.Equal(t, evmtypes.EventTypeContractCreated, createEvent.GetType())

	contractAddr, err := hexutil.Decode(createEvent.Attributes[0].Value)
	require.NoError(t, err)

	// mint 1_000_000 tokens to the first address
	tx2, _ := tests.GenerateMintERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[0], new(big.Int).SetUint64(1_000_000_000_000))
	evmTx2, _, err := app.EVMKeeper.TxUtils().ConvertCosmosTxToEthereumTx(ctx, tx2)
	require.NoError(t, err)

	txBz, err = evmTx2.MarshalBinary()
	require.NoError(t, err)
	txHash2, err := backend.SendRawTransaction(txBz)
	require.NoError(t, err)

	_, finalizeRes = tests.ExecuteTxs(t, app, tx2)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	// wait txs to be indexed
	time.Sleep(1 * time.Second)

	// there should be 2 changes
	changes, err := input.filterAPI.GetFilterChanges(filterID)
	require.NoError(t, err)
	require.Len(t, changes, 2)
	require.Equal(t, []common.Hash{txHash1, txHash2}, changes.([]common.Hash))
}

func Test_NewBlockFilter(t *testing.T) {
	input := setupFilterAPI(t)
	defer input.app.Close()

	// wait indexer to be ready
	time.Sleep(3 * time.Second)

	filterID, err := input.filterAPI.NewBlockFilter()
	require.NoError(t, err)
	require.NotEmpty(t, filterID)

	app, backend, addrs, privKeys := input.app, input.backend, input.addrs, input.privKeys

	tx, _ := tests.GenerateCreateERC20Tx(t, app, privKeys[0])
	_, finalizeRes := tests.ExecuteTxs(t, app, tx)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	events := finalizeRes.TxResults[0].Events
	createEvent := events[len(events)-3]
	require.Equal(t, evmtypes.EventTypeContractCreated, createEvent.GetType())

	contractAddr, err := hexutil.Decode(createEvent.Attributes[0].Value)
	require.NoError(t, err)

	// mint 1_000_000 tokens to the first address
	tx2, _ := tests.GenerateMintERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[0], new(big.Int).SetUint64(1_000_000_000_000))
	_, finalizeRes = tests.ExecuteTxs(t, app, tx2)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	// wait txs to be indexed
	time.Sleep(1 * time.Second)

	// there should be 2 changes
	changes, err := input.filterAPI.GetFilterChanges(filterID)
	require.NoError(t, err)
	require.Len(t, changes, 2)

	blockHash := changes.([]common.Hash)[0]
	header, err := backend.GetHeaderByHash(blockHash)
	require.NoError(t, err)
	require.Equal(t, app.LastBlockHeight()-1, header.Number.Int64())

	blockHash = changes.([]common.Hash)[1]
	header, err = backend.GetHeaderByHash(blockHash)
	require.NoError(t, err)
	require.Equal(t, app.LastBlockHeight(), header.Number.Int64())
}

func Test_NewFilter(t *testing.T) {
	input := setupFilterAPI(t)
	defer input.app.Close()

	// wait indexer to be ready
	time.Sleep(3 * time.Second)

	app, addrs, privKeys := input.app, input.addrs, input.privKeys

	// invalid block range
	_, err := input.filterAPI.NewFilter(ethfilters.FilterCriteria{
		FromBlock: big.NewInt(100),
		ToBlock:   big.NewInt(10),
	})
	require.Error(t, err)

	// start tracking after 2 blocks
	filterID, err := input.filterAPI.NewFilter(ethfilters.FilterCriteria{
		FromBlock: big.NewInt(app.LastBlockHeight() + 2),
		ToBlock:   big.NewInt(int64(rpc.LatestBlockNumber)),
	})
	require.NoError(t, err)
	require.NotEmpty(t, filterID)

	// track all blocks
	filterID2, err := input.filterAPI.NewFilter(ethfilters.FilterCriteria{})
	require.NoError(t, err)
	require.NotEmpty(t, filterID2)

	// this should not tracked by filter 1 but by filter 2
	tx, txHash1 := tests.GenerateCreateERC20Tx(t, app, privKeys[0])
	_, finalizeRes := tests.ExecuteTxs(t, app, tx)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	events := finalizeRes.TxResults[0].Events
	createEvent := events[len(events)-3]
	require.Equal(t, evmtypes.EventTypeContractCreated, createEvent.GetType())

	contractAddr, err := hexutil.Decode(createEvent.Attributes[0].Value)
	require.NoError(t, err)

	// this should be tracked by both filters
	// mint 1_000_000 tokens to the first address
	tx2, txHash2 := tests.GenerateMintERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[0], new(big.Int).SetUint64(1_000_000_000_000))
	_, finalizeRes = tests.ExecuteTxs(t, app, tx2)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	// wait txs to be indexed
	time.Sleep(1 * time.Second)

	changes, err := input.filterAPI.GetFilterChanges(filterID)
	require.NoError(t, err)
	require.Len(t, changes, 1)
	for _, change := range changes.([]*coretypes.Log) {
		require.Equal(t, txHash2, change.TxHash)
	}

	changes, err = input.filterAPI.GetFilterChanges(filterID2)
	require.NoError(t, err)
	require.Len(t, changes, 3)
	require.Equal(t, txHash1, changes.([]*coretypes.Log)[0].TxHash)
	require.Equal(t, txHash1, changes.([]*coretypes.Log)[1].TxHash)
	require.Equal(t, txHash2, changes.([]*coretypes.Log)[2].TxHash)
}

func Test_GetLogs(t *testing.T) {
	input := setupFilterAPI(t)
	defer input.app.Close()

	// wait indexer to be ready
	time.Sleep(3 * time.Second)

	app, addrs, privKeys := input.app, input.addrs, input.privKeys

	// invalid block range
	_, err := input.filterAPI.NewFilter(ethfilters.FilterCriteria{
		FromBlock: big.NewInt(100),
		ToBlock:   big.NewInt(10),
	})
	require.Error(t, err)

	tx, txHash1 := tests.GenerateCreateERC20Tx(t, app, privKeys[0])
	_, finalizeRes := tests.ExecuteTxs(t, app, tx)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	events := finalizeRes.TxResults[0].Events
	createEvent := events[len(events)-3]
	require.Equal(t, evmtypes.EventTypeContractCreated, createEvent.GetType())

	contractAddr, err := hexutil.Decode(createEvent.Attributes[0].Value)
	require.NoError(t, err)

	// mint 1_000_000 tokens to the first address
	tx2, txHash2 := tests.GenerateMintERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[0], new(big.Int).SetUint64(1_000_000_000_000))
	_, finalizeRes = tests.ExecuteTxs(t, app, tx2)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	// wait txs to be indexed
	time.Sleep(1 * time.Second)

	logs, err := input.filterAPI.GetLogs(context.Background(), ethfilters.FilterCriteria{})
	require.NoError(t, err)
	require.NotEmpty(t, logs)
	for _, log := range logs {
		require.Equal(t, txHash2, log.TxHash)
	}

	logs, err = input.filterAPI.GetLogs(context.Background(), ethfilters.FilterCriteria{
		FromBlock: big.NewInt(app.LastBlockHeight() - 1),
		ToBlock:   big.NewInt(app.LastBlockHeight() - 1),
	})
	require.NoError(t, err)
	require.NotEmpty(t, logs)
	for _, log := range logs {
		require.Equal(t, txHash1, log.TxHash)
	}

	// by block hash
	header, err := input.backend.GetHeaderByNumber(rpc.BlockNumber(app.LastBlockHeight()))
	require.NoError(t, err)
	blockHash := header.Hash()
	logs, err = input.filterAPI.GetLogs(context.Background(), ethfilters.FilterCriteria{
		BlockHash: &blockHash,
	})
	require.NoError(t, err)
	require.NotEmpty(t, logs)
	for _, log := range logs {
		require.Equal(t, txHash2, log.TxHash)
	}

	header, err = input.backend.GetHeaderByNumber(rpc.BlockNumber(app.LastBlockHeight() - 1))
	require.NoError(t, err)
	blockHash = header.Hash()
	logs, err = input.filterAPI.GetLogs(context.Background(), ethfilters.FilterCriteria{
		BlockHash: &blockHash,
	})
	require.NoError(t, err)
	require.NotEmpty(t, logs)
	for _, log := range logs {
		require.Equal(t, txHash1, log.TxHash)
	}
}

func Test_GetFilterLogs(t *testing.T) {
	input := setupFilterAPI(t)
	defer input.app.Close()

	// wait indexer to be ready
	time.Sleep(3 * time.Second)

	app, addrs, privKeys := input.app, input.addrs, input.privKeys

	// start tracking after 2 blocks
	filterID, err := input.filterAPI.NewFilter(ethfilters.FilterCriteria{
		FromBlock: big.NewInt(app.LastBlockHeight() + 2),
		ToBlock:   big.NewInt(int64(rpc.LatestBlockNumber)),
	})
	require.NoError(t, err)
	require.NotEmpty(t, filterID)

	// track all blocks from the last block
	filterID2, err := input.filterAPI.NewFilter(ethfilters.FilterCriteria{
		FromBlock: big.NewInt(app.LastBlockHeight()),
	})
	require.NoError(t, err)
	require.NotEmpty(t, filterID2)

	tx, txHash1 := tests.GenerateCreateERC20Tx(t, app, privKeys[0])
	_, finalizeRes := tests.ExecuteTxs(t, app, tx)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	events := finalizeRes.TxResults[0].Events
	createEvent := events[len(events)-3]
	require.Equal(t, evmtypes.EventTypeContractCreated, createEvent.GetType())

	contractAddr, err := hexutil.Decode(createEvent.Attributes[0].Value)
	require.NoError(t, err)

	// mint 1_000_000 tokens to the first address
	tx2, txHash2 := tests.GenerateMintERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[0], new(big.Int).SetUint64(1_000_000_000_000))
	_, finalizeRes = tests.ExecuteTxs(t, app, tx2)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	// wait txs to be indexed
	time.Sleep(1 * time.Second)

	// there should be 1 changes
	logs, err := input.filterAPI.GetFilterLogs(context.Background(), filterID)
	require.NoError(t, err)
	require.Len(t, logs, 1)
	for _, change := range logs {
		require.Equal(t, txHash2, change.TxHash)
	}

	// there should be 3 changes
	logs, err = input.filterAPI.GetFilterLogs(context.Background(), filterID2)
	require.NoError(t, err)
	require.Len(t, logs, 3)
	require.Equal(t, txHash1, logs[0].TxHash)
	require.Equal(t, txHash1, logs[1].TxHash)
	require.Equal(t, txHash2, logs[2].TxHash)
}
