package config

import (
	"fmt"

	"github.com/spf13/cast"
	"github.com/spf13/cobra"

	servertypes "github.com/cosmos/cosmos-sdk/server/types"
)

const (
	// DefaultContractSimulationGasLimit - default max simulation gas
	DefaultContractSimulationGasLimit = uint64(3_000_000)
	// DefaultIndexerDisable is the default flag to disable indexer
	DefaultIndexerDisable = false
	// DefaultIndexerCacheSize is the default maximum size (MiB) of the cache.
	DefaultIndexerCacheSize = 100
	// DefaultIndexerRetainHeight is the default height to retain indexer data.
	DefaultIndexerRetainHeight = uint64(0)

	// SectionSize is the size of the section for bloom indexing
	SectionSize = uint64(4096)
)

const (
	flagContractSimulationGasLimit = "evm.contract-simulation-gas-limit"
	flagIndexerDisable             = "evm.indexer-disable"
	flagIndexerCacheSize           = "evm.indexer-cache-size"
	flagIndexerRetainHeight        = "evm.indexer-retain-height"
)

// EVMConfig is the extra config required for evm
type EVMConfig struct {
	// ContractSimulationGasLimit is the maximum gas amount can be used in a tx simulation call.
	ContractSimulationGasLimit uint64 `mapstructure:"contract-simulation-gas-limit"`
	// IndexerDisable is the flag to disable indexer
	IndexerDisable bool `mapstructure:"indexer-disable"`
	// IndexerCacheSize is the maximum size (MiB) of the cache.
	IndexerCacheSize int `mapstructure:"indexer-cache-size"`
	// IndexerRetainHeight is the height to retain indexer data.
	// If 0, it will retain all data.
	IndexerRetainHeight uint64 `mapstructure:"indexer-retain-height"`
}

func (c EVMConfig) Validate() error {
	if c.IndexerRetainHeight%SectionSize != 0 {
		return fmt.Errorf("indexer-retain-height must be a multiple of %d", SectionSize)
	}

	return nil
}

// DefaultEVMConfig returns the default settings for EVMConfig
func DefaultEVMConfig() EVMConfig {
	return EVMConfig{
		ContractSimulationGasLimit: DefaultContractSimulationGasLimit,
		IndexerDisable:             DefaultIndexerDisable,
		IndexerCacheSize:           DefaultIndexerCacheSize,
		IndexerRetainHeight:        DefaultIndexerRetainHeight,
	}
}

// GetConfig load config values from the app options
func GetConfig(appOpts servertypes.AppOptions) EVMConfig {
	return EVMConfig{
		ContractSimulationGasLimit: cast.ToUint64(appOpts.Get(flagContractSimulationGasLimit)),
		IndexerDisable:             cast.ToBool(appOpts.Get(flagIndexerDisable)),
		IndexerCacheSize:           cast.ToInt(appOpts.Get(flagIndexerCacheSize)),
		IndexerRetainHeight:        cast.ToUint64(appOpts.Get(flagIndexerRetainHeight)),
	}
}

// AddConfigFlags implements servertypes.EVMConfigFlags interface.
func AddConfigFlags(startCmd *cobra.Command) {
	startCmd.Flags().Uint64(flagContractSimulationGasLimit, DefaultContractSimulationGasLimit, "Maximum simulation gas amount for evm contract execution")
	startCmd.Flags().Bool(flagIndexerDisable, DefaultIndexerDisable, "Disable evm indexer")
	startCmd.Flags().Int(flagIndexerCacheSize, DefaultIndexerCacheSize, "Maximum size (MiB) of the indexer cache")
	startCmd.Flags().Uint64(flagIndexerRetainHeight, DefaultIndexerRetainHeight, "Height to retain indexer data")
}

// DefaultConfigTemplate default config template for evm
const DefaultConfigTemplate = `
###############################################################################
###                         EVM                                             ###
###############################################################################

[evm]

# The maximum gas amount can be used in a tx simulation call.
contract-simulation-gas-limit = "{{ .EVMConfig.ContractSimulationGasLimit }}"

# IndexerDisable is the flag to disable indexer. If true, evm jsonrpc queries will return 
# empty results for block, tx, and receipt queries.
indexer-disable = {{ .EVMConfig.IndexerDisable }}

# IndexerCacheSize is the maximum size (MiB) of the cache for evm indexer.
indexer-cache-size = {{ .EVMConfig.IndexerCacheSize }}

# IndexerRetainHeight is the height to retain indexer data.
# If 0, it will retain all data.
indexer-retain-height = {{ .EVMConfig.IndexerRetainHeight }}
`
