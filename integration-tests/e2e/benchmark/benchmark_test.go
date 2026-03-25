//go:build benchmark

package benchmark

import (
	"context"
	"encoding/hex"
	"math/big"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/initia-labs/minievm/integration-tests/e2e/benchmark/bench_erc20"
	"github.com/initia-labs/minievm/integration-tests/e2e/benchmark/bench_heavy_state"
	"github.com/initia-labs/minievm/integration-tests/e2e/cluster"
	"github.com/stretchr/testify/require"
)

const (
	clusterReadyTimeout = 120 * time.Second
	mempoolDrainTimeout = 180 * time.Second
	mempoolPollInterval = 500 * time.Millisecond
	warmupSettleTime    = 5 * time.Second
)

func resultsDir(t *testing.T) string {
	t.Helper()
	if d := os.Getenv("BENCHMARK_RESULTS_DIR"); d != "" {
		return d
	}
	return filepath.Join("results")
}

func setupCluster(t *testing.T, ctx context.Context, cfg BenchConfig) *cluster.Cluster {
	t.Helper()

	cl, err := cluster.NewCluster(ctx, t, cluster.ClusterOptions{
		NodeCount:      cfg.NodeCount,
		AccountCount:   cfg.AccountCount,
		ChainID:        "bench-minievm",
		BinaryPath:     os.Getenv("E2E_MINITIAD_BIN"),
		MemIAVL:        cfg.MemIAVL,
		ValidatorCount: cfg.ValidatorCount,
		MaxBlockGas:    cfg.MaxBlockGas,
		NoAllowQueued:  cfg.NoAllowQueued,
	})
	require.NoError(t, err)

	require.NoError(t, cl.Start(ctx))
	t.Cleanup(cl.Close)
	require.NoError(t, cl.WaitForReady(ctx, clusterReadyTimeout))

	return cl
}

func runBenchmarkWithCluster(t *testing.T, ctx context.Context, cl *cluster.Cluster, cfg BenchConfig, loadFn func(ctx context.Context, cl *cluster.Cluster, cfg BenchConfig, metas map[string]cluster.AccountMeta) LoadResult) BenchResult {
	t.Helper()

	metas, err := CollectInitialMetas(ctx, cl)
	require.NoError(t, err)

	Warmup(ctx, cl, metas)
	require.NoError(t, cl.WaitForMempoolEmpty(ctx, 30*time.Second))
	time.Sleep(warmupSettleTime)

	metas, err = CollectInitialMetas(ctx, cl)
	require.NoError(t, err)

	startHeight, err := cl.LatestHeight(ctx, 0)
	require.NoError(t, err)

	poller := NewMempoolPoller(ctx, cl, mempoolPollInterval)

	t.Logf("Starting load: %d accounts x %d txs = %d total", cfg.AccountCount, cfg.TxPerAccount, cfg.TotalTx())
	loadResult := loadFn(ctx, cl, cfg, metas)
	t.Logf("Load complete: %d submitted, %d errors, duration=%.1fs",
		len(loadResult.Submissions), len(loadResult.Errors),
		loadResult.EndTime.Sub(loadResult.StartTime).Seconds())

	drainTimeout := mempoolDrainTimeout + time.Duration(cfg.TotalTx()/20)*time.Second
	endHeight, err := WaitForLoadToSettle(ctx, cl, drainTimeout, cfg.NoAllowQueued)
	require.NoError(t, err)

	peakMempool := poller.Stop()

	result, err := CollectResults(ctx, cl, cfg, loadResult, startHeight, endHeight, peakMempool)
	require.NoError(t, err)

	t.Logf("Results: TPS=%.1f, P50=%.0fms, P95=%.0fms, P99=%.0fms, included=%d/%d, peak_mempool=%d",
		result.TxPerSecond, result.P50LatencyMs, result.P95LatencyMs, result.P99LatencyMs,
		result.TotalIncluded, result.TotalSubmitted, result.PeakMempoolSize)

	require.NoError(t, WriteResult(t, result, resultsDir(t)))

	return result
}

func runBenchmark(t *testing.T, cfg BenchConfig, loadFn func(ctx context.Context, cl *cluster.Cluster, cfg BenchConfig, metas map[string]cluster.AccountMeta) LoadResult) BenchResult {
	t.Helper()
	ctx := context.Background()

	cl := setupCluster(t, ctx, cfg)
	defer cl.Close()

	return runBenchmarkWithCluster(t, ctx, cl, cfg, loadFn)
}

// ---------------------------------------------------------------------------
// EVM exec setup
// ---------------------------------------------------------------------------

type evmExecLoadMode int

const (
	evmExecBurst evmExecLoadMode = iota
	evmExecSequential
)

// setupEvmExecLoad deploys BenchHeavyState and estimates gas.
// Returns a LoadFn closure that uses burst or sequential EVM exec depending on the mode.
func setupEvmExecLoad(t *testing.T, ctx context.Context, cl *cluster.Cluster, mode ...evmExecLoadMode) func(ctx context.Context, cl *cluster.Cluster, cfg BenchConfig, metas map[string]cluster.AccountMeta) LoadResult {
	t.Helper()

	const (
		sharedWrites int64 = 5
		localWrites  int64 = 25
	)

	// 1. Deploy BenchHeavyState contract
	deployerName := cl.AccountNames()[0]
	res := cl.DeployContract(ctx, deployerName, bench_heavy_state.BenchHeavyStateBin, 0)
	require.NoError(t, res.Err)
	require.Equal(t, int64(0), res.Code, "deploy failed: %s", res.RawLog)

	// 2. Wait for inclusion
	require.NoError(t, cl.WaitForMempoolEmpty(ctx, 30*time.Second))
	time.Sleep(3 * time.Second)

	// 3. Extract contract address from tx result
	txResult, err := cl.QueryTxResult(ctx, 0, res.TxHash)
	require.NoError(t, err)
	contractAddr, err := cluster.ExtractContractAddress(txResult)
	require.NoError(t, err)
	t.Logf("BenchHeavyState deployed at: %s", contractAddr)

	// 4. ABI-encode writeMixed(sharedCount, localCount) call
	inputBytes, err := bench_heavy_state.PackWriteMixed(sharedWrites, localWrites)
	require.NoError(t, err)
	inputHex := "0x" + hex.EncodeToString(inputBytes)

	// 5. Estimate gas
	estimatedGas, err := cl.EstimateEvmGas(ctx, deployerName, contractAddr, inputHex, 0)
	require.NoError(t, err)
	t.Logf("Estimated gas for writeMixed(%d shared, %d local): %d", sharedWrites, localWrites, estimatedGas)

	estimatedGas = estimatedGas * 3 / 2
	t.Logf("Adjusted gas (1.5x): %d", estimatedGas)

	// 6. Return load function
	m := evmExecBurst
	if len(mode) > 0 {
		m = mode[0]
	}
	if m == evmExecSequential {
		return EvmExecSequentialLoad(contractAddr, inputHex, estimatedGas)
	}

	return EvmExecBurstLoad(contractAddr, inputHex, estimatedGas)
}

// setupEvmExecCluster deploys BenchHeavyState and returns params for pre-signing.
func setupEvmExecCluster(t *testing.T, ctx context.Context, cl *cluster.Cluster, sharedWrites, localWrites int64) (contractAddr, inputHex string, estimatedGas uint64) {
	t.Helper()

	deployerName := cl.AccountNames()[0]
	res := cl.DeployContract(ctx, deployerName, bench_heavy_state.BenchHeavyStateBin, 0)
	require.NoError(t, res.Err)
	require.Equal(t, int64(0), res.Code, "deploy failed: %s", res.RawLog)
	require.NoError(t, cl.WaitForMempoolEmpty(ctx, 30*time.Second))
	time.Sleep(3 * time.Second)

	txResult, err := cl.QueryTxResult(ctx, 0, res.TxHash)
	require.NoError(t, err)
	contractAddr, err = cluster.ExtractContractAddress(txResult)
	require.NoError(t, err)

	inputBytes, err := bench_heavy_state.PackWriteMixed(sharedWrites, localWrites)
	require.NoError(t, err)
	inputHex = "0x" + hex.EncodeToString(inputBytes)

	estimatedGas, err = cl.EstimateEvmGas(ctx, deployerName, contractAddr, inputHex, 0)
	require.NoError(t, err)

	estimatedGas = estimatedGas * 3 / 2
	t.Logf("Estimated gas for writeMixed(%d shared, %d local): %d (with 1.5x adjustment)", sharedWrites, localWrites, estimatedGas)

	return contractAddr, inputHex, estimatedGas
}

// ---------------------------------------------------------------------------
// Mempool comparison: CList vs. Proxy+Priority
// ---------------------------------------------------------------------------

func TestBenchmarkBaselineSeq(t *testing.T) {
	cfg := BaselineConfig()
	cfg.Label = "clist/iavl/seq"
	runBenchmark(t, cfg, SequentialLoad)
}

func TestBenchmarkBaselineBurst(t *testing.T) {
	cfg := BaselineConfig()
	cfg.Label = "clist/iavl/burst"
	runBenchmark(t, cfg, BurstLoad)
}

func TestBenchmarkSeqComparison(t *testing.T) {
	var results []BenchResult

	baselines := LoadBaselineResultsByLabel(resultsDir(t), "clist/iavl/seq")
	if len(baselines) > 0 {
		t.Logf("Loaded baseline result: %s", baselines[0].Config.Label)
		results = append(results, baselines[0])
	} else {
		t.Log("No baseline results found. Run TestBenchmarkBaselineSeq with pre-proxy binary for full comparison.")
	}

	t.Run("MempoolOnly", func(t *testing.T) {
		cfg := MempoolOnlyConfig()
		cfg.Label = "proxy+priority/iavl/seq"
		result := runBenchmark(t, cfg, SequentialLoad)
		results = append(results, result)
	})

	t.Run("Combined", func(t *testing.T) {
		cfg := CombinedConfig()
		cfg.Label = "proxy+priority/memiavl/seq"
		result := runBenchmark(t, cfg, SequentialLoad)
		results = append(results, result)
	})

	if len(results) >= 2 {
		PrintComparisonTable(t, results)
		PrintImprovementTable(t, results)
	}
}

func TestBenchmarkBurstComparison(t *testing.T) {
	var results []BenchResult

	baselines := LoadBaselineResultsByLabel(resultsDir(t), "clist/iavl/burst")
	if len(baselines) > 0 {
		t.Logf("Loaded baseline result: %s", baselines[0].Config.Label)
		results = append(results, baselines[0])
	} else {
		t.Log("No baseline results found. Run TestBenchmarkBaselineBurst with pre-proxy binary for full comparison.")
	}

	t.Run("MempoolOnly", func(t *testing.T) {
		cfg := MempoolOnlyConfig()
		cfg.Label = "proxy+priority/iavl/burst"
		result := runBenchmark(t, cfg, BurstLoad)
		results = append(results, result)
	})

	t.Run("Combined", func(t *testing.T) {
		cfg := CombinedConfig()
		cfg.Label = "proxy+priority/memiavl/burst"
		result := runBenchmark(t, cfg, BurstLoad)
		results = append(results, result)
	})

	if len(results) >= 2 {
		PrintComparisonTable(t, results)
		PrintImprovementTable(t, results)
	}
}

// ---------------------------------------------------------------------------
// EVM exec tests (CLI-based)
// ---------------------------------------------------------------------------

func TestBenchmarkBaselineSeqEvmExec(t *testing.T) {
	cfg := BaselineConfig()
	cfg.AccountCount = 100
	cfg.TxPerAccount = 50
	cfg.Label = "clist/iavl/seq-evm-exec"

	ctx := context.Background()
	cl := setupCluster(t, ctx, cfg)
	defer cl.Close()

	evmLoadFn := setupEvmExecLoad(t, ctx, cl, evmExecSequential)
	runBenchmarkWithCluster(t, ctx, cl, cfg, evmLoadFn)
}

func TestBenchmarkSeqComparisonEvmExec(t *testing.T) {
	var results []BenchResult

	baselines := LoadBaselineResultsByLabel(resultsDir(t), "clist/iavl/seq-evm-exec")
	if len(baselines) > 0 {
		t.Logf("Loaded baseline result: %s", baselines[0].Config.Label)
		results = append(results, baselines[0])
	} else {
		t.Log("No baseline results found. Run TestBenchmarkBaselineSeqEvmExec with pre-proxy binary for full comparison.")
	}

	t.Run("MempoolOnly", func(t *testing.T) {
		cfg := MempoolOnlyConfig()
		cfg.AccountCount = 100
		cfg.TxPerAccount = 50
		cfg.Label = "proxy+priority/iavl/seq-evm-exec"

		ctx := context.Background()
		cl := setupCluster(t, ctx, cfg)
		defer cl.Close()

		evmLoadFn := setupEvmExecLoad(t, ctx, cl, evmExecSequential)
		result := runBenchmarkWithCluster(t, ctx, cl, cfg, evmLoadFn)
		results = append(results, result)
	})

	t.Run("Combined", func(t *testing.T) {
		cfg := CombinedConfig()
		cfg.AccountCount = 100
		cfg.TxPerAccount = 50
		cfg.Label = "proxy+priority/memiavl/seq-evm-exec"

		ctx := context.Background()
		cl := setupCluster(t, ctx, cfg)
		defer cl.Close()

		evmLoadFn := setupEvmExecLoad(t, ctx, cl, evmExecSequential)
		result := runBenchmarkWithCluster(t, ctx, cl, cfg, evmLoadFn)
		results = append(results, result)
	})

	if len(results) >= 2 {
		PrintComparisonTable(t, results)
		PrintImprovementTable(t, results)
	}
}

func TestBenchmarkBurstComparisonEvmExec(t *testing.T) {
	var results []BenchResult

	t.Run("MempoolOnly", func(t *testing.T) {
		cfg := MempoolOnlyConfig()
		cfg.AccountCount = 100
		cfg.TxPerAccount = 50
		cfg.Label = "proxy+priority/iavl/burst-evm-exec"

		ctx := context.Background()
		cl := setupCluster(t, ctx, cfg)
		defer cl.Close()

		evmLoadFn := setupEvmExecLoad(t, ctx, cl)
		result := runBenchmarkWithCluster(t, ctx, cl, cfg, evmLoadFn)
		results = append(results, result)
	})

	t.Run("Combined", func(t *testing.T) {
		cfg := CombinedConfig()
		cfg.AccountCount = 100
		cfg.TxPerAccount = 50
		cfg.Label = "proxy+priority/memiavl/burst-evm-exec"

		ctx := context.Background()
		cl := setupCluster(t, ctx, cfg)
		defer cl.Close()

		evmLoadFn := setupEvmExecLoad(t, ctx, cl)
		result := runBenchmarkWithCluster(t, ctx, cl, cfg, evmLoadFn)
		results = append(results, result)
	})

	if len(results) >= 2 {
		PrintComparisonTable(t, results)
	}
}

// ---------------------------------------------------------------------------
// Pre-signed HTTP broadcast benchmarks
// ---------------------------------------------------------------------------

func runPreSignedBenchmark(
	t *testing.T, ctx context.Context, cl *cluster.Cluster, cfg BenchConfig,
	preSignFn func(metas map[string]cluster.AccountMeta) []cluster.SignedTx,
	loadFnFactory func([]cluster.SignedTx) func(ctx context.Context, cl *cluster.Cluster, cfg BenchConfig, metas map[string]cluster.AccountMeta) LoadResult,
) BenchResult {
	t.Helper()

	metas, err := CollectInitialMetas(ctx, cl)
	require.NoError(t, err)

	Warmup(ctx, cl, metas)
	require.NoError(t, cl.WaitForMempoolEmpty(ctx, 30*time.Second))
	time.Sleep(warmupSettleTime)

	metas, err = CollectInitialMetas(ctx, cl)
	require.NoError(t, err)

	signedTxs := preSignFn(metas)

	startHeight, err := cl.LatestHeight(ctx, 0)
	require.NoError(t, err)

	poller := NewMempoolPoller(ctx, cl, mempoolPollInterval)

	t.Logf("Starting load: %d accounts x %d txs = %d total (pre-signed HTTP)", cfg.AccountCount, cfg.TxPerAccount, cfg.TotalTx())
	loadFn := loadFnFactory(signedTxs)
	loadResult := loadFn(ctx, cl, cfg, metas)
	t.Logf("Load complete: %d submitted, %d errors, duration=%.1fs",
		len(loadResult.Submissions), len(loadResult.Errors),
		loadResult.EndTime.Sub(loadResult.StartTime).Seconds())

	drainTimeout := mempoolDrainTimeout + time.Duration(cfg.TotalTx()/20)*time.Second
	endHeight, err := WaitForLoadToSettle(ctx, cl, drainTimeout, cfg.NoAllowQueued)
	if err != nil {
		t.Logf("Warning: mempool drain incomplete: %v (collecting partial results)", err)
		// Still collect results with whatever blocks we have
		endHeight, _ = cl.LatestHeight(ctx, 0)
	}

	peakMempool := poller.Stop()

	result, err := CollectResults(ctx, cl, cfg, loadResult, startHeight, endHeight, peakMempool)
	require.NoError(t, err)

	t.Logf("Results: TPS=%.1f, P50=%.0fms, P95=%.0fms, P99=%.0fms, included=%d/%d, peak_mempool=%d",
		result.TxPerSecond, result.P50LatencyMs, result.P95LatencyMs, result.P99LatencyMs,
		result.TotalIncluded, result.TotalSubmitted, result.PeakMempoolSize)

	require.NoError(t, WriteResult(t, result, resultsDir(t)))
	return result
}

func TestBenchmarkPreSignedSeqComparison(t *testing.T) {
	var results []BenchResult

	t.Run("IAVL", func(t *testing.T) {
		cfg := MempoolOnlyConfig()
		cfg.AccountCount = 20
		cfg.TxPerAccount = 100
		cfg.Label = "presigned/iavl/seq"

		ctx := context.Background()
		cl := setupCluster(t, ctx, cfg)
		defer cl.Close()

		result := runPreSignedBenchmark(t, ctx, cl, cfg,
			func(metas map[string]cluster.AccountMeta) []cluster.SignedTx {
				return PreSignBankTxs(ctx, t, cl, cfg, metas)
			}, PreSignedSequentialLoad)
		results = append(results, result)
	})

	t.Run("MemIAVL", func(t *testing.T) {
		cfg := CombinedConfig()
		cfg.AccountCount = 20
		cfg.TxPerAccount = 100
		cfg.Label = "presigned/memiavl/seq"

		ctx := context.Background()
		cl := setupCluster(t, ctx, cfg)
		defer cl.Close()

		result := runPreSignedBenchmark(t, ctx, cl, cfg,
			func(metas map[string]cluster.AccountMeta) []cluster.SignedTx {
				return PreSignBankTxs(ctx, t, cl, cfg, metas)
			}, PreSignedSequentialLoad)
		results = append(results, result)
	})

	if len(results) == 2 {
		PrintComparisonTable(t, results)
	}
}

func TestBenchmarkPreSignedBurstComparison(t *testing.T) {
	var results []BenchResult

	t.Run("IAVL", func(t *testing.T) {
		cfg := MempoolOnlyConfig()
		cfg.AccountCount = 20
		cfg.TxPerAccount = 100
		cfg.Label = "presigned/iavl/burst"

		ctx := context.Background()
		cl := setupCluster(t, ctx, cfg)
		defer cl.Close()

		result := runPreSignedBenchmark(t, ctx, cl, cfg,
			func(metas map[string]cluster.AccountMeta) []cluster.SignedTx {
				return PreSignBankTxs(ctx, t, cl, cfg, metas)
			}, PreSignedBurstLoad)
		results = append(results, result)
	})

	t.Run("MemIAVL", func(t *testing.T) {
		cfg := CombinedConfig()
		cfg.AccountCount = 20
		cfg.TxPerAccount = 100
		cfg.Label = "presigned/memiavl/burst"

		ctx := context.Background()
		cl := setupCluster(t, ctx, cfg)
		defer cl.Close()

		result := runPreSignedBenchmark(t, ctx, cl, cfg,
			func(metas map[string]cluster.AccountMeta) []cluster.SignedTx {
				return PreSignBankTxs(ctx, t, cl, cfg, metas)
			}, PreSignedBurstLoad)
		results = append(results, result)
	})

	if len(results) == 2 {
		PrintComparisonTable(t, results)
	}
}

func TestBenchmarkPreSignedSeqEvmExec(t *testing.T) {
	var results []BenchResult

	const (
		sharedWrites int64 = 20
		localWrites  int64 = 80
	)

	t.Run("IAVL", func(t *testing.T) {
		cfg := MempoolOnlyConfig()
		cfg.AccountCount = 20
		cfg.TxPerAccount = 100
		cfg.Label = "presigned/iavl/seq-evm-exec"

		ctx := context.Background()
		cl := setupCluster(t, ctx, cfg)
		defer cl.Close()

		contractAddr, inputHex, gas := setupEvmExecCluster(t, ctx, cl, sharedWrites, localWrites)

		result := runPreSignedBenchmark(t, ctx, cl, cfg,
			func(metas map[string]cluster.AccountMeta) []cluster.SignedTx {
				return PreSignEvmCallTxs(ctx, t, cl, cfg, metas, contractAddr, inputHex, gas)
			}, PreSignedSequentialLoad)
		results = append(results, result)
	})

	t.Run("MemIAVL", func(t *testing.T) {
		cfg := CombinedConfig()
		cfg.AccountCount = 20
		cfg.TxPerAccount = 100
		cfg.Label = "presigned/memiavl/seq-evm-exec"

		ctx := context.Background()
		cl := setupCluster(t, ctx, cfg)
		defer cl.Close()

		contractAddr, inputHex, gas := setupEvmExecCluster(t, ctx, cl, sharedWrites, localWrites)

		result := runPreSignedBenchmark(t, ctx, cl, cfg,
			func(metas map[string]cluster.AccountMeta) []cluster.SignedTx {
				return PreSignEvmCallTxs(ctx, t, cl, cfg, metas, contractAddr, inputHex, gas)
			}, PreSignedSequentialLoad)
		results = append(results, result)
	})

	if len(results) == 2 {
		PrintComparisonTable(t, results)
	}
}

func TestBenchmarkPreSignedSeqEvmExecStress(t *testing.T) {
	var results []BenchResult

	const (
		sharedWrites int64 = 20
		localWrites  int64 = 80
	)

	t.Run("IAVL", func(t *testing.T) {
		cfg := MempoolOnlyConfig()
		cfg.AccountCount = 20
		cfg.TxPerAccount = 200
		cfg.Label = "presigned-stress/iavl/seq-evm-exec"

		ctx := context.Background()
		cl := setupCluster(t, ctx, cfg)
		defer cl.Close()

		contractAddr, inputHex, gas := setupEvmExecCluster(t, ctx, cl, sharedWrites, localWrites)

		result := runPreSignedBenchmark(t, ctx, cl, cfg,
			func(metas map[string]cluster.AccountMeta) []cluster.SignedTx {
				return PreSignEvmCallTxs(ctx, t, cl, cfg, metas, contractAddr, inputHex, gas)
			}, PreSignedSequentialLoad)
		results = append(results, result)
	})

	t.Run("MemIAVL", func(t *testing.T) {
		cfg := CombinedConfig()
		cfg.AccountCount = 20
		cfg.TxPerAccount = 200
		cfg.Label = "presigned-stress/memiavl/seq-evm-exec"

		ctx := context.Background()
		cl := setupCluster(t, ctx, cfg)
		defer cl.Close()

		contractAddr, inputHex, gas := setupEvmExecCluster(t, ctx, cl, sharedWrites, localWrites)

		result := runPreSignedBenchmark(t, ctx, cl, cfg,
			func(metas map[string]cluster.AccountMeta) []cluster.SignedTx {
				return PreSignEvmCallTxs(ctx, t, cl, cfg, metas, contractAddr, inputHex, gas)
			}, PreSignedSequentialLoad)
		results = append(results, result)
	})

	if len(results) == 2 {
		PrintComparisonTable(t, results)
	}
}

// ---------------------------------------------------------------------------
// State DB comparison: IAVL vs MemIAVL
// ---------------------------------------------------------------------------

func TestBenchmarkMemIAVLBankSend(t *testing.T) {
	var results []BenchResult

	t.Run("IAVL", func(t *testing.T) {
		cfg := MempoolOnlyConfig()
		cfg.AccountCount = 100
		cfg.TxPerAccount = 200
		cfg.Label = "memiavl-compare/iavl/bank-send"
		result := runBenchmark(t, cfg, BurstLoad)
		results = append(results, result)
	})

	t.Run("MemIAVL", func(t *testing.T) {
		cfg := CombinedConfig()
		cfg.AccountCount = 100
		cfg.TxPerAccount = 200
		cfg.Label = "memiavl-compare/memiavl/bank-send"
		result := runBenchmark(t, cfg, BurstLoad)
		results = append(results, result)
	})

	if len(results) == 2 {
		PrintComparisonTable(t, results)
	}
}

func TestBenchmarkMemIAVLEvmExec(t *testing.T) {
	var results []BenchResult

	t.Run("IAVL", func(t *testing.T) {
		cfg := MempoolOnlyConfig()
		cfg.AccountCount = 100
		cfg.TxPerAccount = 50
		cfg.Label = "memiavl-compare/iavl/evm-exec"

		ctx := context.Background()
		cl := setupCluster(t, ctx, cfg)
		defer cl.Close()

		evmLoadFn := setupEvmExecLoad(t, ctx, cl)
		result := runBenchmarkWithCluster(t, ctx, cl, cfg, evmLoadFn)
		results = append(results, result)
	})

	t.Run("MemIAVL", func(t *testing.T) {
		cfg := CombinedConfig()
		cfg.AccountCount = 100
		cfg.TxPerAccount = 50
		cfg.Label = "memiavl-compare/memiavl/evm-exec"

		ctx := context.Background()
		cl := setupCluster(t, ctx, cfg)
		defer cl.Close()

		evmLoadFn := setupEvmExecLoad(t, ctx, cl)
		result := runBenchmarkWithCluster(t, ctx, cl, cfg, evmLoadFn)
		results = append(results, result)
	})

	if len(results) == 2 {
		PrintComparisonTable(t, results)
	}
}

// ---------------------------------------------------------------------------
// Capability demos
// ---------------------------------------------------------------------------

func TestBenchmarkQueuePromotion(t *testing.T) {
	cfg := MempoolOnlyConfig()
	cfg.TxPerAccount = 50
	cfg.Label = "queue-promotion/mempool-only"
	result := runBenchmark(t, cfg, OutOfOrderLoad)

	require.Equal(t, result.TotalSubmitted, result.TotalIncluded,
		"not all out-of-order transactions were included: submitted=%d included=%d",
		result.TotalSubmitted, result.TotalIncluded)
}

func TestBenchmarkQueuedFlood(t *testing.T) {
	cfg := MempoolOnlyConfig()
	cfg.TxPerAccount = 50
	cfg.Label = "queued-flood/mempool-only"
	result := runBenchmark(t, cfg, QueuedFloodLoad)

	require.Equal(t, result.TotalSubmitted, result.TotalIncluded,
		"not all queued-flood transactions were included: submitted=%d included=%d",
		result.TotalSubmitted, result.TotalIncluded)
}

func TestBenchmarkQueuedGapEviction(t *testing.T) {
	cfg := MempoolOnlyConfig()
	cfg.TxPerAccount = 50
	cfg.Label = "queued-gap-eviction/mempool-only"

	ctx := context.Background()
	cl := setupCluster(t, ctx, cfg)
	defer cl.Close()

	metas, err := CollectInitialMetas(ctx, cl)
	require.NoError(t, err)

	Warmup(ctx, cl, metas)
	require.NoError(t, cl.WaitForMempoolEmpty(ctx, 30*time.Second))
	time.Sleep(warmupSettleTime)

	metas, err = CollectInitialMetas(ctx, cl)
	require.NoError(t, err)

	// start mempool poller before load to capture peak queued size
	poller := NewMempoolPoller(ctx, cl, mempoolPollInterval)

	loadResult := QueuedGapLoad(ctx, cl, cfg, metas)
	t.Logf("Submitted %d future-nonce txs (no gap fill), %d errors",
		len(loadResult.Submissions), len(loadResult.Errors))

	t.Log("Waiting for gap TTL eviction (60s + 30s buffer)...")
	time.Sleep(90 * time.Second)

	err = cl.WaitForMempoolEmpty(ctx, 30*time.Second)
	peakMempool := poller.Stop()

	t.Logf("Gap eviction test: peak_mempool=%d, mempool_drained=%v",
		peakMempool, err == nil)

	require.NoError(t, err, "mempool should be empty after gap TTL eviction")
	require.Greater(t, peakMempool, 0, "should have observed queued txs in mempool")
}

func TestBenchmarkGossipPropagation(t *testing.T) {
	cfg := MempoolOnlyConfig()
	cfg.AccountCount = 5
	cfg.TxPerAccount = 50
	cfg.Label = "gossip/mempool-only"

	ctx := context.Background()
	cl := setupCluster(t, ctx, cfg)
	defer cl.Close()

	metas, err := CollectInitialMetas(ctx, cl)
	require.NoError(t, err)

	Warmup(ctx, cl, metas)
	require.NoError(t, cl.WaitForMempoolEmpty(ctx, 30*time.Second))
	time.Sleep(warmupSettleTime)

	metas, err = CollectInitialMetas(ctx, cl)
	require.NoError(t, err)

	startHeight, err := cl.LatestHeight(ctx, 0)
	require.NoError(t, err)

	poller := NewMempoolPoller(ctx, cl, mempoolPollInterval)

	loadResult := SingleNodeLoad(ctx, cl, cfg, metas, 0)
	t.Logf("Submitted %d txs to node 0", len(loadResult.Submissions))

	endHeight, err := WaitForLoadToSettle(ctx, cl, mempoolDrainTimeout, false)
	require.NoError(t, err)

	peakMempool := poller.Stop()
	t.Logf("Cluster peak mempool size: %d", peakMempool)

	result, err := CollectResults(ctx, cl, cfg, loadResult, startHeight, endHeight, peakMempool)
	require.NoError(t, err)

	t.Logf("Gossip test: TPS=%.1f, included=%d/%d",
		result.TxPerSecond, result.TotalIncluded, result.TotalSubmitted)
	require.NoError(t, WriteResult(t, result, resultsDir(t)))
}

// ---------------------------------------------------------------------------
// JSON-RPC benchmarks (eth_sendRawTransaction)
// ---------------------------------------------------------------------------

const defaultEthGasLimit uint64 = 100_000

func runJsonRpcBenchmark(
	t *testing.T, ctx context.Context, cl *cluster.Cluster, cfg BenchConfig,
	preSignFn func(nonces map[string]EthNonceMeta) []cluster.SignedEthTx,
	loadFnFactory func([]cluster.SignedEthTx) func(ctx context.Context, cl *cluster.Cluster, cfg BenchConfig, metas map[string]cluster.AccountMeta) LoadResult,
) BenchResult {
	t.Helper()

	cosmosMetas, err := CollectInitialMetas(ctx, cl)
	require.NoError(t, err)

	valEthAddr, err := cl.ValidatorEthAddress()
	require.NoError(t, err)

	nonces, err := CollectEthNonces(ctx, cl)
	require.NoError(t, err)

	WarmupEth(ctx, t, cl, nonces, valEthAddr, defaultEthGasLimit)
	require.NoError(t, cl.WaitForMempoolEmpty(ctx, 30*time.Second))
	time.Sleep(warmupSettleTime)

	nonces, err = CollectEthNonces(ctx, cl)
	require.NoError(t, err)

	signedTxs := preSignFn(nonces)

	startHeight, err := cl.LatestHeight(ctx, 0)
	require.NoError(t, err)

	poller := NewMempoolPoller(ctx, cl, mempoolPollInterval)

	t.Logf("Starting JSON-RPC load: %d accounts x %d txs = %d total", cfg.AccountCount, cfg.TxPerAccount, cfg.TotalTx())
	loadFn := loadFnFactory(signedTxs)
	loadResult := loadFn(ctx, cl, cfg, cosmosMetas)
	t.Logf("Load complete: %d submitted, %d errors, duration=%.1fs",
		len(loadResult.Submissions), len(loadResult.Errors),
		loadResult.EndTime.Sub(loadResult.StartTime).Seconds())

	drainTimeout := mempoolDrainTimeout + time.Duration(cfg.TotalTx()/20)*time.Second
	endHeight, err := WaitForLoadToSettle(ctx, cl, drainTimeout, cfg.NoAllowQueued)
	require.NoError(t, err)

	peakMempool := poller.Stop()

	result, err := CollectResultsEth(ctx, cl, cfg, loadResult, startHeight, endHeight, peakMempool)
	require.NoError(t, err)

	t.Logf("Results: TPS=%.1f, P50=%.0fms, P95=%.0fms, P99=%.0fms, included=%d/%d, peak_mempool=%d",
		result.TxPerSecond, result.P50LatencyMs, result.P95LatencyMs, result.P99LatencyMs,
		result.TotalIncluded, result.TotalSubmitted, result.PeakMempoolSize)

	require.NoError(t, WriteResult(t, result, resultsDir(t)))
	return result
}

func TestBenchmarkJsonRpcSeqEvmExec(t *testing.T) {
	const (
		sharedWrites int64 = 5
		localWrites  int64 = 25
	)

	cfg := MempoolOnlyConfig()
	cfg.AccountCount = 20
	cfg.TxPerAccount = 100
	cfg.Label = "jsonrpc/iavl/seq-evm-exec"

	ctx := context.Background()
	cl := setupCluster(t, ctx, cfg)
	defer cl.Close()

	contractAddrHex, inputHex, _ := setupEvmExecCluster(t, ctx, cl, sharedWrites, localWrites)
	contractAddr := common.HexToAddress(contractAddrHex)

	ethAccounts := cl.EthAccounts()
	require.Greater(t, len(ethAccounts), 0)
	gas, err := cl.EthEstimateGas(ctx, 0, ethAccounts[0].Address.Hex(), contractAddr.Hex(), inputHex)
	require.NoError(t, err)
	gas = gas * 3 / 2 // 1.5x adjustment
	t.Logf("JSON-RPC estimated gas: %d (with 1.5x adjustment)", gas)

	runJsonRpcBenchmark(t, ctx, cl, cfg,
		func(nonces map[string]EthNonceMeta) []cluster.SignedEthTx {
			return PreSignEthContractCallTxs(ctx, t, cl, cfg, nonces, contractAddr, inputHex, gas)
		}, JsonRpcSequentialLoad)
}

func TestBenchmarkJsonRpcBurstEvmExec(t *testing.T) {
	const (
		sharedWrites int64 = 5
		localWrites  int64 = 25
	)

	cfg := MempoolOnlyConfig()
	cfg.AccountCount = 20
	cfg.TxPerAccount = 100
	cfg.Label = "jsonrpc/iavl/burst-evm-exec"

	ctx := context.Background()
	cl := setupCluster(t, ctx, cfg)
	defer cl.Close()

	contractAddrHex, inputHex, _ := setupEvmExecCluster(t, ctx, cl, sharedWrites, localWrites)
	contractAddr := common.HexToAddress(contractAddrHex)

	ethAccounts := cl.EthAccounts()
	require.Greater(t, len(ethAccounts), 0)
	gas, err := cl.EthEstimateGas(ctx, 0, ethAccounts[0].Address.Hex(), contractAddr.Hex(), inputHex)
	require.NoError(t, err)
	gas = gas * 3 / 2
	t.Logf("JSON-RPC estimated gas: %d (with 1.5x adjustment)", gas)

	runJsonRpcBenchmark(t, ctx, cl, cfg,
		func(nonces map[string]EthNonceMeta) []cluster.SignedEthTx {
			return PreSignEthContractCallTxs(ctx, t, cl, cfg, nonces, contractAddr, inputHex, gas)
		}, JsonRpcBurstLoad)
}

func TestBenchmarkJsonRpcQueuePromotion(t *testing.T) {
	const (
		sharedWrites int64 = 5
		localWrites  int64 = 25
	)

	cfg := MempoolOnlyConfig()
	cfg.AccountCount = 10
	cfg.TxPerAccount = 50
	cfg.Label = "jsonrpc-queue-promotion/mempool-only"

	ctx := context.Background()
	cl := setupCluster(t, ctx, cfg)
	defer cl.Close()

	contractAddrHex, inputHex, _ := setupEvmExecCluster(t, ctx, cl, sharedWrites, localWrites)
	contractAddr := common.HexToAddress(contractAddrHex)

	ethAccounts := cl.EthAccounts()
	require.Greater(t, len(ethAccounts), 0)
	gas, err := cl.EthEstimateGas(ctx, 0, ethAccounts[0].Address.Hex(), contractAddr.Hex(), inputHex)
	require.NoError(t, err)
	gas = gas * 3 / 2
	t.Logf("JSON-RPC estimated gas: %d (with 1.5x adjustment)", gas)

	result := runJsonRpcBenchmark(t, ctx, cl, cfg,
		func(nonces map[string]EthNonceMeta) []cluster.SignedEthTx {
			return PreSignEthContractCallsOutOfOrder(ctx, t, cl, cfg, nonces, contractAddr, inputHex, gas)
		}, JsonRpcPreserveOrderLoad)

	require.Equal(t, result.TotalSubmitted, result.TotalIncluded,
		"not all out-of-order JSON-RPC transactions were included: submitted=%d included=%d",
		result.TotalSubmitted, result.TotalIncluded)
}

// ---------------------------------------------------------------------------
// ERC20 transfer benchmark (JSON-RPC)
// ---------------------------------------------------------------------------

func setupErc20Benchmark(t *testing.T, ctx context.Context, cl *cluster.Cluster, cfg BenchConfig) BenchResult {
	t.Helper()

	// 1. Deploy BenchERC20 via CLI
	deployerName := cl.AccountNames()[0]
	res := cl.DeployContract(ctx, deployerName, bench_erc20.BenchErc20Bin, 0)
	require.NoError(t, res.Err)
	require.Equal(t, int64(0), res.Code, "deploy failed: %s", res.RawLog)
	require.NoError(t, cl.WaitForMempoolEmpty(ctx, 30*time.Second))
	time.Sleep(3 * time.Second)

	txResult, err := cl.QueryTxResult(ctx, 0, res.TxHash)
	require.NoError(t, err)
	contractAddrHex, err := cluster.ExtractContractAddress(txResult)
	require.NoError(t, err)
	contractAddr := common.HexToAddress(contractAddrHex)
	t.Logf("BenchERC20 deployed at: %s", contractAddr.Hex())

	// 2. Collect ETH nonces
	nonces, err := CollectEthNonces(ctx, cl)
	require.NoError(t, err)

	// 3. Mint tokens for each account (enough for all transfers)
	ethAccounts := cl.EthAccounts()
	mintAmount := new(big.Int).Mul(big.NewInt(int64(cfg.TxPerAccount)), big.NewInt(1e18)) //nolint:gosec
	mintData, err := bench_erc20.PackMint(ethAccounts[0].Address, mintAmount)
	require.NoError(t, err)
	mintGas, err := cl.EthEstimateGas(ctx, 0, ethAccounts[0].Address.Hex(), contractAddr.Hex(), "0x"+hex.EncodeToString(mintData))
	require.NoError(t, err)
	mintGas = mintGas * 3 / 2
	t.Logf("Mint gas: %d (with 1.5x adjustment)", mintGas)

	MintErc20ForAccounts(ctx, t, cl, nonces, contractAddr, mintAmount, mintGas)
	require.NoError(t, cl.WaitForMempoolEmpty(ctx, 30*time.Second))
	time.Sleep(warmupSettleTime)

	// 4. Build transfer calldata
	recipient := ethAccounts[len(ethAccounts)-1].Address
	transferAmount := big.NewInt(1e18)
	transferData, err := bench_erc20.PackTransfer(recipient, transferAmount)
	require.NoError(t, err)

	transferGas, err := cl.EthEstimateGas(ctx, 0, ethAccounts[0].Address.Hex(), contractAddr.Hex(), "0x"+hex.EncodeToString(transferData))
	require.NoError(t, err)
	transferGas = transferGas * 3 / 2
	t.Logf("Transfer gas: %d (with 1.5x adjustment)", transferGas)

	// 5. Run benchmark
	return runJsonRpcBenchmark(t, ctx, cl, cfg,
		func(nonces map[string]EthNonceMeta) []cluster.SignedEthTx {
			return PreSignEthErc20TransferTxs(ctx, t, cl, cfg, nonces, contractAddr, transferData, transferGas)
		}, JsonRpcSequentialLoad)
}

func TestBenchmarkJsonRpcErc20Transfer(t *testing.T) {
	var results []BenchResult

	t.Run("IAVL", func(t *testing.T) {
		cfg := MempoolOnlyConfig()
		cfg.AccountCount = 20
		cfg.TxPerAccount = 100
		cfg.Label = "jsonrpc/iavl/erc20-transfer"

		ctx := context.Background()
		cl := setupCluster(t, ctx, cfg)
		defer cl.Close()

		result := setupErc20Benchmark(t, ctx, cl, cfg)
		results = append(results, result)
	})

	t.Run("MemIAVL", func(t *testing.T) {
		cfg := CombinedConfig()
		cfg.AccountCount = 20
		cfg.TxPerAccount = 100
		cfg.Label = "jsonrpc/memiavl/erc20-transfer"

		ctx := context.Background()
		cl := setupCluster(t, ctx, cfg)
		defer cl.Close()

		result := setupErc20Benchmark(t, ctx, cl, cfg)
		results = append(results, result)
	})

	if len(results) == 2 {
		PrintComparisonTable(t, results)
	}
}
