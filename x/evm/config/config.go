package config

import (
	"fmt"
	"time"

	"github.com/spf13/cast"
	"github.com/spf13/cobra"

	servertypes "github.com/cosmos/cosmos-sdk/server/types"
)

const (
	// DefaultContractSimulationGasLimit - default max simulation gas
	DefaultContractSimulationGasLimit = uint64(3_000_000)
	// DefaultIndexerDisable is the default flag to disable indexer
	DefaultIndexerDisable = false
	// DefaultIndexerRetainHeight is the default height to retain indexer data.
	DefaultIndexerRetainHeight = uint64(0)
	// DefaultTracerTimeout is the default tracer timeout.
	DefaultTracerTimeout = 10 * time.Second
	// DefaultIndexerBackfillStartHeight is the default height to start backfilling indexer data.
	DefaultIndexerBackfillStartHeight = uint64(0)

	// SectionSize is the size of the section for bloom indexing
	SectionSize = uint64(4096)
)

const (
	flagContractSimulationGasLimit = "evm.contract-simulation-gas-limit"
	flagIndexerDisable             = "evm.indexer-disable"
	flagIndexerRetainHeight        = "evm.indexer-retain-height"
	flagTracerTimeout              = "evm.tracer-timeout"
	flagIndexerBackfillStartHeight = "evm.indexer-backfill-start-height"
)

// EVMConfig is the extra config required for evm
type EVMConfig struct {
	// ContractSimulationGasLimit is the maximum gas amount can be used in a tx simulation call.
	ContractSimulationGasLimit uint64 `mapstructure:"contract-simulation-gas-limit"`
	// IndexerDisable is the flag to disable indexer
	IndexerDisable bool `mapstructure:"indexer-disable"`
	// IndexerRetainHeight is the height to retain indexer data.
	// If 0, it will retain all data.
	IndexerRetainHeight uint64 `mapstructure:"indexer-retain-height"`
	// TracerTimeout is the timeout for the tracer.
	TracerTimeout time.Duration `mapstructure:"tracer-timeout"`
	// IndexerBackfillStartHeight is the height to start backfilling indexer data.
	// If non-zero, it will start backfilling from this height until last indexed height.
	IndexerBackfillStartHeight uint64 `mapstructure:"indexer-backfill-start-height"`
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
		IndexerRetainHeight:        DefaultIndexerRetainHeight,
		TracerTimeout:              DefaultTracerTimeout,
		IndexerBackfillStartHeight: DefaultIndexerBackfillStartHeight,
	}
}

// GetConfig load config values from the app options
func GetConfig(appOpts servertypes.AppOptions) EVMConfig {
	tracerTimeout := cast.ToDuration(appOpts.Get(flagTracerTimeout))
	if tracerTimeout == 0 {
		tracerTimeout = DefaultTracerTimeout
	}

	return EVMConfig{
		ContractSimulationGasLimit: cast.ToUint64(appOpts.Get(flagContractSimulationGasLimit)),
		IndexerDisable:             cast.ToBool(appOpts.Get(flagIndexerDisable)),
		IndexerRetainHeight:        cast.ToUint64(appOpts.Get(flagIndexerRetainHeight)),
		TracerTimeout:              tracerTimeout,
		IndexerBackfillStartHeight: cast.ToUint64(appOpts.Get(flagIndexerBackfillStartHeight)),
	}
}

// AddConfigFlags implements servertypes.EVMConfigFlags interface.
func AddConfigFlags(startCmd *cobra.Command) {
	startCmd.Flags().Uint64(flagContractSimulationGasLimit, DefaultContractSimulationGasLimit, "Maximum simulation gas amount for evm contract execution")
	startCmd.Flags().Bool(flagIndexerDisable, DefaultIndexerDisable, "Disable evm indexer")
	startCmd.Flags().Uint64(flagIndexerRetainHeight, DefaultIndexerRetainHeight, "Height to retain indexer data")
	startCmd.Flags().Duration(flagTracerTimeout, DefaultTracerTimeout, "Timeout for the tracer")
	startCmd.Flags().Uint64(flagIndexerBackfillStartHeight, DefaultIndexerBackfillStartHeight, "Height to start backfilling indexer data")
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

# IndexerRetainHeight is the height to retain indexer data.
# If 0, it will retain all data.
indexer-retain-height = {{ .EVMConfig.IndexerRetainHeight }}

# TracerTimeout is the timeout for the tracer.
tracer-timeout = "{{ .EVMConfig.TracerTimeout }}"

# IndexerBackfillStartHeight is the height to start backfilling indexer data.
# If non-zero, it will start backfilling from this height until last indexed height.
indexer-backfill-start-height = {{ .EVMConfig.IndexerBackfillStartHeight }}
`
