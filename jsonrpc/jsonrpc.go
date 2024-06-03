package jsonrpc

import (
	"context"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"golang.org/x/net/netutil"
	"golang.org/x/sync/errgroup"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server"

	ethlog "github.com/ethereum/go-ethereum/log"
	ethrpc "github.com/ethereum/go-ethereum/rpc"
)

func StartJSONRPC(
	ctx context.Context,
	g *errgroup.Group,
	svrCtx *server.Context,
	clientCtx client.Context,
	jsonRPCConfig *JSONRPCConfig,
) error {
	logger := svrCtx.Logger.With("module", "geth")
	ethlog.SetDefault(ethlog.NewLogger(newLogger(logger)))

	rpcServer := ethrpc.NewServer()
	apis := jsonrpc.GetRPCAPIs(svrCtx, clientCtx, jsonRPCConfig.APIs)

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
func listen(addr string, jsonRPCConfig *JSONRPCConfig) (net.Listener, error) {
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
