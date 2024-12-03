package backend_test

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"

	minitiaapp "github.com/initia-labs/minievm/app"
	"github.com/initia-labs/minievm/indexer"
	"github.com/initia-labs/minievm/jsonrpc/backend"
	"github.com/initia-labs/minievm/jsonrpc/config"
	"github.com/initia-labs/minievm/tests"
	evmtypes "github.com/initia-labs/minievm/x/evm/types"
)

type testInput struct {
	app      *minitiaapp.MinitiaApp
	indexer  indexer.EVMIndexer
	backend  *backend.JSONRPCBackend
	addrs    []common.Address
	privKeys []*ecdsa.PrivateKey
	cometRPC *tests.MockCometRPC
}

func setupBackend(t *testing.T) testInput {
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

	mockCometRPC := tests.NewMockCometRPC(app.BaseApp)
	clientCtx = clientCtx.WithClient(mockCometRPC)

	backend, err := backend.NewJSONRPCBackend(ctx, app, app.Logger(), svrCtx, clientCtx, cfg)
	require.NoError(t, err)

	return testInput{
		app:      app,
		indexer:  indexer,
		backend:  backend,
		addrs:    addrs,
		privKeys: privKeys,
		cometRPC: mockCometRPC,
	}
}

func Test_FloodingQuery(t *testing.T) {
	input := setupBackend(t)
	app, _, backend, addrs, privKeys := input.app, input.addrs, input.backend, input.addrs, input.privKeys

	tx, _ := tests.GenerateCreateERC20Tx(t, app, privKeys[0])
	_, finalizeRes := tests.ExecuteTxs(t, app, tx)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	events := finalizeRes.TxResults[0].Events
	createEvent := events[len(events)-3]
	require.Equal(t, evmtypes.EventTypeContractCreated, createEvent.GetType())

	contractAddr, err := hexutil.Decode(createEvent.Attributes[0].Value)
	require.NoError(t, err)

	// mint 1_000_000 tokens to the first address
	tx, _ = tests.GenerateMintERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[0], new(big.Int).SetUint64(1_000_000_000_000))
	_, finalizeRes = tests.ExecuteTxs(t, app, tx)
	tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

	time.Sleep(3 * time.Second)

	ctx, cancel := context.WithCancel(context.Background())
	queryFn := func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				_, err := backend.GetBalance(addrs[0], rpc.BlockNumberOrHashWithNumber(-1))
				require.NoError(t, err)

				time.Sleep(5 * time.Millisecond)
			}
		}
	}

	for i := 0; i < 100; i++ {
		go queryFn()
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for i := 0; i < 1000; i++ {
			tx, _ = tests.GenerateTransferERC20Tx(t, app, privKeys[0], common.BytesToAddress(contractAddr), addrs[1], new(big.Int).SetUint64(1_000_000))
			_, finalizeRes = tests.ExecuteTxs(t, app, tx)
			tests.CheckTxResult(t, finalizeRes.TxResults[0], true)

			time.Sleep(5 * time.Millisecond)
		}
		wg.Done()
	}()

	wg.Wait()
	cancel()
}
