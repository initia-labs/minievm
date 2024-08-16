package jsonrpc

import (
	"context"
	"net"
	"net/http"
	"slices"

	"github.com/gorilla/mux"
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
	cosmosns "github.com/initia-labs/minievm/jsonrpc/namespaces/cosmos"
	ethns "github.com/initia-labs/minievm/jsonrpc/namespaces/eth"
	"github.com/initia-labs/minievm/jsonrpc/namespaces/eth/filters"
	netns "github.com/initia-labs/minievm/jsonrpc/namespaces/net"
	txpoolns "github.com/initia-labs/minievm/jsonrpc/namespaces/txpool"
	web3ns "github.com/initia-labs/minievm/jsonrpc/namespaces/web3"
)

// RPC namespaces and API version
const (
	// TODO: implement commented apis in the namespaces for full Ethereum compatibility
	EthNamespace      = "eth"
	NetNamespace      = "net"
	TxPoolNamespace   = "txpool"
	Web3Namespace     = "web3"
	PersonalNamespace = "personal"
	DebugNamespace    = "debug"
	MinerNamespace    = "miner"
	CosmosNamespace   = "cosmos"

	apiVersion = "1.0"
)

func StartJSONRPC(
	ctx context.Context,
	g *errgroup.Group,
	app *app.MinitiaApp,
	svrCtx *server.Context,
	clientCtx client.Context,
	jsonRPCConfig config.JSONRPCConfig,
	isWebSocket bool,
) error {
	if !isWebSocket && !jsonRPCConfig.Enable {
		return nil
	} else if isWebSocket && !jsonRPCConfig.EnableWS {
		return nil
	}

	logger := svrCtx.Logger.With("module", "geth").With("api", "jsonrpc")
	ethlog.SetDefault(ethlog.NewLogger(newLogger(logger)))

	rpcServer := rpc.NewServer()
	rpcServer.SetBatchLimits(jsonRPCConfig.BatchRequestLimit, jsonRPCConfig.BatchResponseMaxSize)
	bkd := backend.NewJSONRPCBackend(app, svrCtx, clientCtx, jsonRPCConfig)
	apis := []rpc.API{
		{
			Namespace: EthNamespace,
			Version:   apiVersion,
			Service:   ethns.NewEthAPI(svrCtx.Logger, bkd),
			Public:    true,
		},
		{
			Namespace: EthNamespace,
			Version:   apiVersion,
			Service:   filters.NewFilterAPI(app, bkd, svrCtx.Logger),
			Public:    true,
		},
		{
			Namespace: NetNamespace,
			Version:   apiVersion,
			Service:   netns.NewNetAPI(svrCtx.Logger, bkd),
			Public:    true,
		},
		{
			Namespace: Web3Namespace,
			Version:   apiVersion,
			Service:   web3ns.NewWeb3API(svrCtx.Logger, bkd),
			Public:    true,
		},
		{
			Namespace: TxPoolNamespace,
			Version:   apiVersion,
			Service:   txpoolns.NewTxPoolAPI(svrCtx.Logger, bkd),
			Public:    true,
		},
		{
			Namespace: CosmosNamespace,
			Version:   apiVersion,
			Service:   cosmosns.NewCosmosAPI(svrCtx.Logger, bkd),
			Public:    true,
		},
	}

	for _, api := range apis {
		if slices.Index(jsonRPCConfig.APIs, api.Namespace) == -1 {
			continue
		}

		if err := rpcServer.RegisterName(api.Namespace, api.Service); err != nil {
			svrCtx.Logger.Error(
				"failed to register service in JSON RPC namespace",
				"namespace", api.Namespace,
				"service", api.Service,
			)
			return err
		}
	}

	router := mux.NewRouter()

	var addr string
	if !isWebSocket {
		addr = jsonRPCConfig.Address
		router.Handle("/", rpcServer).Methods("GET").Methods("POST")
	} else {
		allowedOrigins := []string{}
		if jsonRPCConfig.EnableUnsafeCORS {
			allowedOrigins = []string{"*"}
		}

		addr = jsonRPCConfig.AddressWS
		router.Handle("/", rpcServer.WebsocketHandler(allowedOrigins))
	}

	handlerWithCors := cors.Default()
	if jsonRPCConfig.EnableUnsafeCORS {
		handlerWithCors = cors.AllowAll()
	}

	httpSrv := &http.Server{
		Addr:              addr,
		Handler:           handlerWithCors.Handler(router),
		ReadHeaderTimeout: jsonRPCConfig.HTTPTimeout,
		ReadTimeout:       jsonRPCConfig.HTTPTimeout,
		WriteTimeout:      jsonRPCConfig.HTTPTimeout,
		IdleTimeout:       jsonRPCConfig.HTTPIdleTimeout,
	}

	ln, err := listen(httpSrv.Addr, jsonRPCConfig)
	if err != nil {
		return err
	}

	g.Go(func() error {
		errCh := make(chan error)

		go func() {
			if !isWebSocket {
				svrCtx.Logger.Info("Starting JSON-RPC server", "address", jsonRPCConfig.Address)
			} else {
				svrCtx.Logger.Info("Starting JSON-RPC WebSocket server", "address", jsonRPCConfig.AddressWS)
			}

			errCh <- httpSrv.Serve(ln)
		}()

		// Start a blocking select to wait for an indication to stop the server or that
		// the server failed to start properly.
		select {
		case <-ctx.Done():
			// The calling process canceled or closed the provided context, so we must
			// gracefully stop the gRPC server.
			if !isWebSocket {
				logger.Info("stopping Ethereum JSONRPC server...", "address", jsonRPCConfig.Address)
			} else {
				logger.Info("stopping Ethereum JSONRPC WebSocket server...", "address", jsonRPCConfig.AddressWS)
			}

			return httpSrv.Close()

		case err := <-errCh:
			if err != nil {
				if !isWebSocket {
					logger.Error("failed to start Ethereum JSONRPC server", "err", err)
				} else {
					logger.Error("failed to start Ethereum JSONRPC WebSocket server", "err", err)
				}
			}

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
