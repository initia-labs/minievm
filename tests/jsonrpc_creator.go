package tests

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/ethereum/go-ethereum/common"

	minitiaapp "github.com/initia-labs/minievm/app"
	"github.com/initia-labs/minievm/indexer"
	"github.com/initia-labs/minievm/jsonrpc"
	"github.com/initia-labs/minievm/jsonrpc/backend"
	"github.com/initia-labs/minievm/jsonrpc/config"
)

type TestInput struct {
	App       *minitiaapp.MinitiaApp
	Indexer   indexer.EVMIndexer
	Backend   *backend.JSONRPCBackend
	Addrs     []common.Address
	PrivKeys  []*ecdsa.PrivateKey
	CometRPC  *MockCometRPC
	Address   string
	AddressWS string
}

// getFreePort asks the kernel for a free open port that is ready to use.
func getFreePort(t *testing.T) (port int) {
	a, err := net.ResolveTCPAddr("tcp", "localhost:0")
	require.NoError(t, err)

	l, err := net.ListenTCP("tcp", a)
	require.NoError(t, err)

	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port
}

func CreateAppWithJSONRPC(t *testing.T) TestInput {
	app, addrs, privKeys := CreateApp(t)
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
	cfg.FilterTimeout = 3 * time.Second
	cfg.Address = fmt.Sprintf("localhost:%d", getFreePort(t))
	cfg.AddressWS = fmt.Sprintf("localhost:%d", getFreePort(t))

	mockCometRPC := NewMockCometRPC(app.BaseApp)
	clientCtx = clientCtx.WithClient(mockCometRPC)

	backend, err := backend.NewJSONRPCBackend(ctx, app, app.Logger(), svrCtx, clientCtx, cfg)
	require.NoError(t, err)

	g, ctx := errgroup.WithContext(ctx)
	err = jsonrpc.StartJSONRPC(ctx, g, app, svrCtx, clientCtx, cfg, false)
	require.NoError(t, err)
	err = jsonrpc.StartJSONRPC(ctx, g, app, svrCtx, clientCtx, cfg, true)
	require.NoError(t, err)

	return TestInput{
		App:       app,
		Indexer:   indexer,
		Backend:   backend,
		Addrs:     addrs,
		PrivKeys:  privKeys,
		CometRPC:  mockCometRPC,
		Address:   cfg.Address,
		AddressWS: cfg.AddressWS,
	}
}
