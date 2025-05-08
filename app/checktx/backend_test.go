package checktx_test

import (
	"context"
	"crypto/ecdsa"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/ethereum/go-ethereum/common"

	minitiaapp "github.com/initia-labs/minievm/app"
	"github.com/initia-labs/minievm/indexer"
	"github.com/initia-labs/minievm/jsonrpc/backend"
	"github.com/initia-labs/minievm/jsonrpc/config"
	"github.com/initia-labs/minievm/tests"
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

	mockCometRPC := tests.NewMockCometRPC(app)
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
