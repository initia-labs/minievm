package benchmark

import "time"

// Variant identifies which optimization layer is being benchmarked.
type Variant string

const (
	VariantBaseline    Variant = "baseline"     // CListMempool + IAVL
	VariantMempoolOnly Variant = "mempool-only" // ProxyMempool+PriorityMempool + IAVL
	VariantCombined    Variant = "combined"     // ProxyMempool+PriorityMempool + MemIAVL
)

const (
	defaultGasLimit    uint64 = 500_000
	defaultMaxBlockGas int64  = 200_000_000
)

// BenchConfig defines the parameters for a benchmark run.
type BenchConfig struct {
	MemIAVL        bool          `json:"memiavl"`
	NodeCount      int           `json:"node_count"`
	AccountCount   int           `json:"account_count"`
	TxPerAccount   int           `json:"tx_per_account"`
	GasLimit       uint64        `json:"gas_limit"`
	Label          string        `json:"label"`
	Variant        Variant       `json:"variant"`
	TimeoutCommit  time.Duration `json:"timeout_commit_ms,omitempty"`
	ValidatorCount int           `json:"validator_count,omitempty"`
	MaxBlockGas    int64         `json:"max_block_gas,omitempty"`
	NoAllowQueued  bool          `json:"no_allow_queued,omitempty"`
}

// GetGasLimit returns the configured gas limit, falling back to defaultGasLimit.
func (c BenchConfig) GetGasLimit() uint64 {
	if c.GasLimit == 0 {
		return defaultGasLimit
	}
	return c.GasLimit
}

// MempoolOnlyConfig returns a benchmark configuration for the mempool-only improvement layer.
func MempoolOnlyConfig() BenchConfig {
	return BenchConfig{
		MemIAVL:        false,
		NodeCount:      4,
		AccountCount:   10,
		TxPerAccount:   200,
		Label:          "proxy+priority/iavl",
		Variant:        VariantMempoolOnly,
		ValidatorCount: 1,
		MaxBlockGas:    defaultMaxBlockGas,
	}
}

// CombinedConfig returns a benchmark configuration for the combined improvement layer.
func CombinedConfig() BenchConfig {
	return BenchConfig{
		MemIAVL:        true,
		NodeCount:      4,
		AccountCount:   10,
		TxPerAccount:   200,
		Label:          "proxy+priority/memiavl",
		Variant:        VariantCombined,
		ValidatorCount: 1,
		MaxBlockGas:    defaultMaxBlockGas,
	}
}

// BaselineConfig returns a benchmark configuration labeled as baseline.
func BaselineConfig() BenchConfig {
	return BenchConfig{
		MemIAVL:        false,
		NodeCount:      4,
		AccountCount:   10,
		TxPerAccount:   200,
		Label:          "clist/iavl",
		Variant:        VariantBaseline,
		ValidatorCount: 1,
		MaxBlockGas:    defaultMaxBlockGas,
		NoAllowQueued:  true,
	}
}

// TotalTx returns the total number of transactions for this config.
func (c BenchConfig) TotalTx() int {
	return c.AccountCount * c.TxPerAccount
}
