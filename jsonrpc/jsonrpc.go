package jsonrpc

import (
	"context"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	ethns "github.com/initia-labs/minievm/jsonrpc/namespaces/eth"
	netns "github.com/initia-labs/minievm/jsonrpc/namespaces/net"
	"github.com/initia-labs/minievm/jsonrpc/namespaces/txpool"
	"github.com/rs/cors"
	"golang.org/x/net/netutil"
	"golang.org/x/sync/errgroup"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server"

	ethlog "github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/initia-labs/minievm/app"
	"github.com/initia-labs/minievm/jsonrpc/backend"
	"github.com/initia-labs/minievm/jsonrpc/config"
)

// RPC namespaces and API version
const (
	// TODO: implement commented apis in the namespaces for full Ethereum compatibility
	EthNamespace    = "eth"
	NetNamespace    = "net"
	TxPoolNamespace = "txpool"
	// TODO: support more namespaces
	Web3Namespace     = "web3"
	PersonalNamespace = "personal"
	DebugNamespace    = "debug"
	MinerNamespace    = "miner"

	apiVersion = "1.0"
)

func StartJSONRPC(
	ctx context.Context,
	g *errgroup.Group,
	app *app.MinitiaApp,
	svrCtx *server.Context,
	clientCtx client.Context,
	jsonRPCConfig config.JSONRPCConfig,
) error {
	logger := svrCtx.Logger.With("module", "geth")
	ethlog.SetDefault(ethlog.NewLogger(newLogger(logger)))

	rpcServer := rpc.NewServer()
	bkd := backend.NewJSONRPCBackend(app, svrCtx, clientCtx)
	apis := []rpc.API{
		{
			Namespace: EthNamespace,
			Version:   apiVersion,
			Service:   ethns.NewEthAPI(svrCtx.Logger, bkd),
			Public:    true,
		},
		{
			Namespace: TxPoolNamespace,
			Version:   apiVersion,
			Service:   txpool.NewTxPoolAPI(svrCtx.Logger, bkd),
			Public:    true,
		},
		{
			Namespace: NetNamespace,
			Version:   apiVersion,
			Service:   netns.NewNetAPI(svrCtx.Logger, bkd),
			Public:    true,
		},
	}

	for _, api := range apis {
		if err := rpcServer.RegisterName(api.Namespace, api.Service); err != nil {
			svrCtx.Logger.Error(
				"failed to register service in JSON RPC namespace",
				"namespace", api.Namespace,
				"service", api.Service,
			)
			return err
		}
	}

	r := mux.NewRouter()
	r.HandleFunc("/", rpcServer.ServeHTTP).Methods("POST")

	handlerWithCors := cors.Default()
	if jsonRPCConfig.EnableUnsafeCORS {
		handlerWithCors = cors.AllowAll()
	}

	httpSrv := &http.Server{
		Addr:              jsonRPCConfig.Address,
		Handler:           handlerWithCors.Handler(r),
		ReadHeaderTimeout: jsonRPCConfig.HTTPTimeout,
		ReadTimeout:       jsonRPCConfig.HTTPTimeout,
		WriteTimeout:      jsonRPCConfig.HTTPTimeout,
		IdleTimeout:       jsonRPCConfig.HTTPIdleTimeout,
	}

	// httpSrv.Serve()
	ln, err := listen(httpSrv.Addr, jsonRPCConfig)
	if err != nil {
		return err
	}

	g.Go(func() error {
		errCh := make(chan error)

		go func() {
			svrCtx.Logger.Info("Starting JSON-RPC server", "address", jsonRPCConfig.Address)
			errCh <- httpSrv.Serve(ln)
		}()

		// Start a blocking select to wait for an indication to stop the server or that
		// the server failed to start properly.
		select {
		case <-ctx.Done():
			// The calling process canceled or closed the provided context, so we must
			// gracefully stop the gRPC server.
			logger.Info("stopping Ethereum JSONRPC server...", "address", jsonRPCConfig.Address)
			return httpSrv.Close()

		case err := <-errCh:
			logger.Error("failed to start Ethereum JSONRPC server", "err", err)
			return err
		}
	})

	return nil
}

// Listen starts a net.Listener on the tcp network on the given address.
// If there is a specified MaxOpenConnections in the config, it will also set the limitListener.
func listen(addr string, jsonRPCConfig config.JSONRPCConfig) (net.Listener, error) {
	if addr == "" {
		addr = ":http"
	}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	if jsonRPCConfig.MaxOpenConnections > 0 {
		ln = netutil.LimitListener(ln, jsonRPCConfig.MaxOpenConnections)
	}
	return ln, err
}
