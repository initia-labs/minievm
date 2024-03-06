package main

import (
	"fmt"
	"time"

	tmcfg "github.com/cometbft/cometbft/config"
	serverconfig "github.com/cosmos/cosmos-sdk/server/config"

	evmconfig "github.com/initia-labs/minievm/x/evm/config"

	"github.com/initia-labs/minievm/types"
)

// minitiaAppConfig initia specify app config
type minitiaAppConfig struct {
	serverconfig.Config
	EVMConfig evmconfig.EVMConfig `mapstructure:"evm"`
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

	minitiaAppConfig := minitiaAppConfig{
		Config:    *srvCfg,
		EVMConfig: evmconfig.DefaultEVMConfig(),
	}

	minitiaAppTemplate := serverconfig.DefaultConfigTemplate +
		evmconfig.DefaultConfigTemplate

	return minitiaAppTemplate, minitiaAppConfig
}

// initTendermintConfig helps to override default Tendermint Config values.
// return tmcfg.DefaultConfig if no custom configuration is required for the application.
func initTendermintConfig() *tmcfg.Config {
	cfg := tmcfg.DefaultConfig()

	// empty block configure
	cfg.Consensus.CreateEmptyBlocks = false
	cfg.Consensus.CreateEmptyBlocksInterval = time.Minute

	// block time to 0.5s
	cfg.Consensus.TimeoutPropose = 300 * time.Millisecond
	cfg.Consensus.TimeoutProposeDelta = 500 * time.Millisecond
	cfg.Consensus.TimeoutPrevote = 1000 * time.Millisecond
	cfg.Consensus.TimeoutPrevoteDelta = 500 * time.Millisecond
	cfg.Consensus.TimeoutPrecommit = 1000 * time.Millisecond
	cfg.Consensus.TimeoutPrecommitDelta = 500 * time.Millisecond
	cfg.Consensus.TimeoutCommit = 500 * time.Millisecond

	return cfg
}
