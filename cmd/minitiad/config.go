package main

import (
	"fmt"
	"time"

	tmcfg "github.com/cometbft/cometbft/config"
	serverconfig "github.com/cosmos/cosmos-sdk/server/config"

	jsonrpcconfig "github.com/initia-labs/minievm/jsonrpc/config"
	evmconfig "github.com/initia-labs/minievm/x/evm/config"

	"github.com/initia-labs/minievm/types"

	initiastorecfg "github.com/initia-labs/store/config"
)

// minitiaAppConfig initia specify app config
type minitiaAppConfig struct {
	serverconfig.Config
	MemIAVL       initiastorecfg.MemIAVLConfig   `mapstructure:"memiavl"`
	VersionDB     initiastorecfg.VersionDBConfig `mapstructure:"versiondb"`
	EVMConfig     evmconfig.EVMConfig            `mapstructure:"evm"`
	JSONRPCConfig jsonrpcconfig.JSONRPCConfig    `mapstructure:"jsonrpc"`
}

// initAppConfig helps to override default appConfig template and configs.
// return "", nil if no custom configuration is required for the application.
func initAppConfig() (string, interface{}) {
	// Optionally allow the chain developer to overwrite the SDK's default
	// server config.
	srvCfg := serverconfig.DefaultConfig()

	// The SDK's default minimum gas price is set to "" (empty value) inside
	// app.toml. If left empty by validators, the node will halt on startup.
	// However, the chain developer can set a default app.toml value for their
	// validators here.
	//
	// In summary:
	// - if you leave srvCfg.MinGasPrices = "", all validators MUST tweak their
	//   own app.toml config,
	// - if you set srvCfg.MinGasPrices non-empty, validators CAN tweak their
	//   own app.toml to override, or use this default value.
	//
	// In simapp, we set the min gas prices to 0.
	srvCfg.MinGasPrices = fmt.Sprintf("0%s", types.BaseDenom)
	srvCfg.Mempool.MaxTxs = 2000
	srvCfg.QueryGasLimit = 10_000_000
	srvCfg.InterBlockCache = false

	// Enable API and unsafe CORS (CORS allowed from any host)
	srvCfg.API.Enable = true
	srvCfg.API.Swagger = true
	srvCfg.API.EnableUnsafeCORS = true
	srvCfg.API.Address = "tcp://0.0.0.0:1317"

	srvCfg.GRPC.Enable = true
	srvCfg.GRPC.Address = "0.0.0.0:9090"

	evmCfg := evmconfig.DefaultEVMConfig()
	evmCfg.ContractSimulationGasLimit = 10_000_000

	jsonRPCConfig := jsonrpcconfig.DefaultJSONRPCConfig()
	jsonRPCConfig.Address = "0.0.0.0:8545"
	jsonRPCConfig.AddressWS = "0.0.0.0:8546"

	memIAVLCfg := initiastorecfg.DefaultMemIAVLConfig()
	versionDBCfg := initiastorecfg.DefaultVersionDBConfig()

	minitiaAppConfig := minitiaAppConfig{
		Config:        *srvCfg,
		EVMConfig:     evmCfg,
		JSONRPCConfig: jsonRPCConfig,
		MemIAVL:       memIAVLCfg,
		VersionDB:     versionDBCfg,
	}

	minitiaAppConfig.JSONRPCConfig.Address = "0.0.0.0:8545"
	minitiaAppConfig.JSONRPCConfig.AddressWS = "0.0.0.0:8546"

	minitiaAppTemplate := serverconfig.DefaultConfigTemplate +
		evmconfig.DefaultConfigTemplate +
		jsonrpcconfig.DefaultConfigTemplate +
		initiastorecfg.DefaultMemIAVLConfigTemplate +
		initiastorecfg.DefaultVersionDBConfigTemplate

	return minitiaAppTemplate, minitiaAppConfig
}

// initTendermintConfig helps to override default Tendermint Config values.
// return tmcfg.DefaultConfig if no custom configuration is required for the application.
func initTendermintConfig() *tmcfg.Config {
	cfg := tmcfg.DefaultConfig()

	// rpc configure
	cfg.RPC.ListenAddress = "tcp://0.0.0.0:26657"
	cfg.RPC.CORSAllowedOrigins = []string{"*"}

	// performance turning configs
	cfg.P2P.SendRate = 20480000
	cfg.P2P.RecvRate = 20480000
	cfg.P2P.MaxPacketMsgPayloadSize = 1000000 // 1MB
	cfg.P2P.FlushThrottleTimeout = 10 * time.Millisecond
	cfg.Consensus.PeerGossipSleepDuration = 30 * time.Millisecond

	// mempool configs
	cfg.Mempool.Size = 1000
	cfg.Mempool.MaxTxsBytes = 10737418240
	cfg.Mempool.MaxTxBytes = 2048576

	// propose timeout to 100ms to reduce block latency
	cfg.Sequencing.BlockInterval = 100 * time.Millisecond

	// empty block configure
	cfg.Sequencing.CreateEmptyBlocks = false
	cfg.Sequencing.CreateEmptyBlocksInterval = time.Minute

	return cfg
}
