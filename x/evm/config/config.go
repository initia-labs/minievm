package config

import (
	"github.com/spf13/cast"
	"github.com/spf13/cobra"

	servertypes "github.com/cosmos/cosmos-sdk/server/types"
)

const (
	// DefaultContractSimulationGasLimit - default max simulation gas
	DefaultContractSimulationGasLimit = uint64(3_000_000)
	// DefaultIndexerCacheSize is the default maximum size (MiB) of the cache.
	DefaultIndexerCacheSize = 100

	// SectionSize is the size of the section for bloom indexing
	SectionSize = uint64(4096)
)

const (
	flagContractSimulationGasLimit = "evm.contract-simulation-gas-limit"
	flagIndexerCacheSize           = "evm.indexer-cache-size"
)

// EVMConfig is the extra config required for evm
type EVMConfig struct {
	// ContractSimulationGasLimit is the maximum gas amount can be used in a tx simulation call.
	ContractSimulationGasLimit uint64 `mapstructure:"contract-simulation-gas-limit"`
	// IndexerCacheSize is the maximum size (MiB) of the cache.
	IndexerCacheSize int `mapstructure:"indexer-cache-size"`
}

func (c EVMConfig) Validate() error {
	return nil
}

// DefaultEVMConfig returns the default settings for EVMConfig
func DefaultEVMConfig() EVMConfig {
	return EVMConfig{
		ContractSimulationGasLimit: DefaultContractSimulationGasLimit,
		IndexerCacheSize:           DefaultIndexerCacheSize,
	}
}

// GetConfig load config values from the app options
func GetConfig(appOpts servertypes.AppOptions) EVMConfig {
	return EVMConfig{
		ContractSimulationGasLimit: cast.ToUint64(appOpts.Get(flagContractSimulationGasLimit)),
		IndexerCacheSize:           cast.ToInt(appOpts.Get(flagIndexerCacheSize)),
	}
}

// AddConfigFlags implements servertypes.EVMConfigFlags interface.
func AddConfigFlags(startCmd *cobra.Command) {
	startCmd.Flags().Uint64(flagContractSimulationGasLimit, DefaultContractSimulationGasLimit, "Maximum simulation gas amount for evm contract execution")
	startCmd.Flags().Int(flagIndexerCacheSize, DefaultIndexerCacheSize, "Maximum size (MiB) of the indexer cache")
}

// DefaultConfigTemplate default config template for evm
const DefaultConfigTemplate = `
###############################################################################
###                         EVM                                             ###
###############################################################################

[evm]

# The maximum gas amount can be used in a tx simulation call.
contract-simulation-gas-limit = "{{ .EVMConfig.ContractSimulationGasLimit }}"

# IndexerCacheSize is the maximum size (MiB) of the cache for evm indexer.
indexer-cache-size = {{ .EVMConfig.IndexerCacheSize }}
`
