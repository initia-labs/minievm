# E2E Benchmark

Performance benchmark for ProxyMempool + PriorityMempool + MemIAVL on MiniEVM. Measures throughput, latency, and mempool behavior
across optimization layers using a multi-node cluster with production-realistic settings.

## Cluster topology

4-node cluster: 1 sequencer + 3 fullnodes on localhost with deterministic port allocation.

- **Fullnode submission**: Benchmark load is submitted to fullnode (non-validator) nodes (indices 1-3), testing gossip propagation to the sequencer.
- **Block interval**: 100ms (sequencing default, `CreateEmptyBlocks = false` thus blocks only created when txs exist)
- **Gas price**: 0GAS
- **Queued tx extension**: All tx submissions include `--allow-queued` flag (required for `ExtensionOptionQueuedTx`)

## Comparison matrix

### 1. Mempool comparison: CList vs Proxy+Priority

Three load patterns: sequential tests give a fair TPS comparison (both mempools handle in-order correctly), burst tests demonstrate CList's tx-drop problem.

**Baselines** (run with v1.2.14 binary):

| Test | Load | Config |
|---|---|---|
| `TestBenchmarkBaselineSeq` | Sequential, bank send | 10 accts x 200 txs |
| `TestBenchmarkBaselineBurst` | Burst, bank send | 10 accts x 200 txs |
| `TestBenchmarkBaselineSeqEvmExec` | Sequential, EVM exec | 100 accts x 50 txs, 30 writes/tx |

**Comparisons** (3-way: CList vs Proxy+IAVL vs Proxy+MemIAVL):

| Test | Load | Purpose |
|---|---|---|
| `TestBenchmarkSeqComparison` | Sequential, bank send | Fair TPS comparison, lightweight workload |
| `TestBenchmarkSeqComparisonEvmExec` | Sequential, EVM exec | Fair TPS comparison under heavy state pressure |
| `TestBenchmarkBurstComparison` | Burst, bank send | Inclusion rate (CList drops txs) |
| `TestBenchmarkBurstComparisonEvmExec` | Burst, EVM exec | Proxy+IAVL vs Proxy+MemIAVL under burst + heavy state (no CList baseline since it would just drop txs) |

### 2. State db comparison: IAVL vs MemIAVL

Both use Proxy+Priority mempool. Isolates state storage impact.

| Test | Workload | Config |
|---|---|---|
| `TestBenchmarkMemIAVLBankSend` | Bank sends | 100 accts x 200 txs |
| `TestBenchmarkMemIAVLEvmExec` | EVM exec (BenchHeavyState::writeMixed) | 100 accts x 50 txs, 30 writes/tx |

### 3. Pre-signed HTTP broadcast (saturated chain)

Bypasses CLI bottleneck. Txs are generated+signed offline, then POSTed via HTTP to `/broadcast_tx_sync`.

| Test | Load | Config |
|---|---|---|
| `TestBenchmarkPreSignedSeqComparison` | Sequential, bank send, HTTP | 20 accts x 100 txs |
| `TestBenchmarkPreSignedBurstComparison` | Burst, bank send, HTTP | 20 accts x 100 txs |
| `TestBenchmarkPreSignedSeqEvmExec` | Sequential, EVM exec, HTTP | 20 accts x 100 txs, 100 writes/tx |
| `TestBenchmarkPreSignedSeqEvmExecStress` | Sequential, EVM exec, HTTP (stress) | 20 accts x 200 txs, 100 writes/tx |

### 4. JSON-RPC benchmarks (`eth_sendRawTransaction`)

Submits transactions via the Ethereum JSON-RPC endpoint, another primary production interface for MiniEVM. 
Separate ETH accounts are generated and funded in genesis.

| Test | Load | Config |
|---|---|---|
| `TestBenchmarkJsonRpcErc20Transfer` | Sequential, ERC20 transfer, JSON-RPC | 20 accts x 100 txs |
| `TestBenchmarkJsonRpcSeqEvmExec` | Sequential, EVM contract call, JSON-RPC | 20 accts x 100 txs |
| `TestBenchmarkJsonRpcBurstEvmExec` | Burst, EVM contract call, JSON-RPC | 20 accts x 100 txs |
| `TestBenchmarkJsonRpcQueuePromotion` | Out-of-order nonces, EVM contract call, JSON-RPC | 10 accts x 50 txs |

### 5. Capability demos

| Test | What | Config |
|---|---|---|
| `TestBenchmarkQueuePromotion` | Out-of-order nonce handling, 100% inclusion | 10 accts x 50 txs |
| `TestBenchmarkGossipPropagation` | Gossip across nodes | 5 accts x 50 txs |
| `TestBenchmarkQueuedFlood` | Future-nonce burst (nonce gaps), queued pool stress + promotion cascade | 10 accts x 50 txs |
| `TestBenchmarkQueuedGapEviction` | Gap TTL eviction under sustained load | 10 accts x 50 txs |

## Expected outcomes

1. **Sequential (fair comparison)**: CList and Proxy+Priority both handle in-order nonces correctly, so sequential submission should show similar TPS. This is the control that proves Proxy+Priority doesn't regress on the happy path.
2. **Burst (stress test)**: Proxy+Priority >> CList. Under burst, CList's `reCheckTx` and cache-based dedup cause it to silently drop txs, while Proxy+Priority's queued pool absorbs out-of-order arrivals and achieves 100% inclusion.
3. **Heavy state writes**: MemIAVL > IAVL. Lightweight workloads (bank send) won't show a difference because the state db isn't the bottleneck. Heavy EVM exec with many writes per tx is needed, and the chain must be saturated (pre-signed HTTP or JSON-RPC) so the state db becomes the limiting factor.
4. **Combined (Proxy+Priority+MemIAVL)**: Best overall throughput and latency, the mempool improvement eliminates tx drops, and MemIAVL reduces state commit time under heavy writes.
5. **JSON-RPC vs CLI**: JSON-RPC submission with pre-signed Ethereum transactions should match or exceed pre-signed HTTP broadcast rates, as the JSON-RPC path is the native EVM submission path.

## Results

### 1. Mempool comparison: CList (v1.2.14) vs Proxy+Priority

#### Sequential bank send

| Config | Variant | TPS | vs base | P50ms | vs base | P95ms | vs base | Included | Peak MP |
|---|---|---:|--------:|---:|--------:|---:|--------:|---:|---:|
| clist/iavl/seq | baseline | 55.0 |       - | 1081 |       - | 1915 |       - | 2000/2000 | 105 |
| proxy+priority/iavl/seq | mempool-only | 56.7 |   +3.1% | 248 |  -77.1% | 338 |  -82.3% | 2000/2000 | 11 |
| proxy+priority/memiavl/seq | combined | 57.0 |   +3.6% | 249 |  -77.0% | 339 |  -82.3% | 2000/2000 | 13 |

#### Burst bank send

| Config | Variant | TPS | vs base | P50ms | vs base | P95ms | vs base | Included | Peak MP |
|---|---|---:|--------:|---:|--------:|---:|--------:|---:|---:|
| clist/iavl/burst | baseline | 17.4 |       - | 264 |       - | 2012 |       - | 41/2000 | 11 |
| proxy+priority/iavl/burst | mempool-only | 57.2 | +228.7% | 249 |   -5.7% | 333 |  -83.4% | 2000/2000 | 12 |
| proxy+priority/memiavl/burst | combined | 57.4 | +229.9% | 248 |   -6.1% | 341 |  -83.1% | 2000/2000 | 11 |

#### Sequential EVM exec (BenchHeavyState, 30 unique-key writes/tx)

| Config | Variant | TPS | vs base | P50ms | vs base | P95ms | vs base | Included | Peak MP |
|---|---|---:|--------:|---:|--------:|---:|--------:|---:|---:|
| clist/iavl/seq-evm-exec | baseline | 28.0 |       - | 2145 |       - | 3117 |       - | 2776/5000 | 1000 |
| proxy+priority/iavl/seq-evm-exec | mempool-only | 50.8 |  +81.4% | 2406 |  +12.2% | 3907 |  +25.3% | 5000/5000 | 324 |
| proxy+priority/memiavl/seq-evm-exec | combined | 54.3 |  +93.9% | 2064 |   -3.8% | 2916 |   -6.4% | 5000/5000 | 120 |

#### Burst EVM exec (no CList baseline since CList drops txs under burst)

| Config | Variant | TPS | P50ms | P95ms | P99ms | Included | Peak MP |
|---|---|---:|---:|---:|---:|---:|---:|
| proxy+priority/iavl/burst-evm-exec | mempool-only | 50.7 | 2499 | 3915 | 4576 | 5000/5000 | 216 |
| proxy+priority/memiavl/burst-evm-exec | combined | 54.1 | 2107 | 3078 | 3477 | 5000/5000 | 113 |

With unique-key writes per tx (state tree grows continuously), MemIAVL shows +6.7% TPS and -15.7% P50 latency even through the CLI bottleneck. The full differentiation appears in the saturated pre-signed tests below.

### 2. State db comparison: IAVL vs MemIAVL (CLI-based, Proxy+Priority)

| Config | Workload | TPS | P50ms | P95ms | P99ms | Included | Peak MP |
|---|---|---:|---:|---:|---:|---:|---:|
| memiavl-compare/iavl/bank-send | bank send | 53.1 | 2027 | 3623 | 9122 | 20000/20000 | 489 |
| memiavl-compare/memiavl/bank-send | bank send | 53.5 | 1939 | 2585 | 2871 | 20000/20000 | 49 |
| memiavl-compare/iavl/evm-exec | evm exec | 51.9 | 2500 | 4242 | 5005 | 5000/5000 | 171 |
| memiavl-compare/memiavl/evm-exec | evm exec | 52.6 | 2091 | 3164 | 4120 | 5000/5000 | 127 |

CLI-based tests are bottlenecked by CLI overhead (~55 TPS ceiling), masking IAVL vs MemIAVL throughput differences. MemIAVL shows clear improvement in tail latency (P95: -25.4%, P99: -17.7% for EVM exec) and peak mempool size (-26%). See pre-signed HTTP results below for saturated-chain comparison.

### 3. Pre-signed HTTP broadcast (saturated chain)

#### Bank send (IAVL vs MemIAVL)

| Config | TPS | P50ms | P95ms | P99ms | Included | Peak MP |
|---|---:|---:|---:|---:|---:|---:|
| presigned/iavl/seq | 1375.2 | 251 | 330 | 353 | 2000/2000 | 614 |
| presigned/memiavl/seq | 1606.9 | 281 | 354 | 505 | 2000/2000 | 761 |
| presigned/iavl/burst | 1499.5 | 679 | 1166 | 1173 | 2000/2000 | 1597 |
| presigned/memiavl/burst | 1567.1 | 615 | 1124 | 1129 | 2000/2000 | 1597 |

#### EVM exec (IAVL vs MemIAVL, 100 unique-key writes/tx)

| Config | TPS | P50ms | P95ms | P99ms | Included | Peak MP |
|---|---:|---:|---:|---:|---:|---:|
| presigned/iavl/seq-evm-exec | 258.1 | 2258 | 3674 | 4008 | 2000/2000 | 810 |
| presigned/memiavl/seq-evm-exec | 421.7 | 1554 | 3125 | 3201 | 2000/2000 | 1303 |

#### EVM exec stress (IAVL vs MemIAVL, 100 unique-key writes/tx, 4000 txs)

| Config | TPS | P50ms | P95ms | P99ms | Included | Peak MP |
|---|---:|---:|---:|---:|---:|---:|
| presigned-stress/iavl/seq-evm-exec | 156.9 | 4738 | 14339 | 14464 | 4000/4000 | 1802 |
| presigned-stress/memiavl/seq-evm-exec | 432.5 | 1998 | 3503 | 3593 | 4000/4000 | 1487 |

Under saturated heavy state writes with continuously growing state tree, MemIAVL demonstrates decisive superiority. 
At 2000 txs: **+63.4% TPS** with 100% inclusion for both. 
At 4000 txs (stress): **+175.7% TPS** (432 vs 157) and **-57.8% P50 latency** (1998 vs 4738ms).
IAVL degrades sharply as the state tree grows while MemIAVL maintains consistent throughput. Both achieve 100% inclusion.

### 4. JSON-RPC benchmarks (via `eth_sendRawTransaction`)

Native signing submitted via the Ethereum JSON-RPC endpoint, another primary production interface for MiniEVM.

#### ERC20 transfer (standard token workload, IAVL vs MemIAVL)

| Config | TPS | P50ms | P95ms | P99ms | Included | Peak MP |
|---|---:|---:|---:|---:|---:|---:|
| jsonrpc/iavl/erc20-transfer | 1376.9 | 179 | 409 | 426 | 2000/2000 | 680 |
| jsonrpc/memiavl/erc20-transfer | 1609.0 | 223 | 355 | 370 | 2000/2000 | 602 |

ERC20 transfers (2 storage writes per tx: sender balance, recipient balance) achieve **1609 TPS** with MemIAVL, **+16.9%** over IAVL (1377 TPS). 
Even for this lightweight EVM workload, MemIAVL shows a clear throughput advantage with lower tail latency (P95: -13.2%, P99: -13.1%) and lower peak mempool pressure. 
Both achieve 100% inclusion.

#### Heavy state EVM exec (BenchHeavyState, 30 writes/tx)

| Config | TPS | P50ms | P95ms | P99ms | Included | Peak MP |
|---|---:|---:|---:|---:|---:|---:|
| jsonrpc/iavl/seq-evm-exec | 754.4 | 509 | 869 | 968 | 2000/2000 | 871 |
| jsonrpc/iavl/burst-evm-exec | 716.6 | 799 | 1165 | 1184 | 2000/2000 | 913 |

Heavy state EVM contract calls (~30 writes/tx) achieve ~750 TPS, the higher per-tx storage write count is the primary throughput bottleneck compared to lightweight ERC20 transfers (~1600 TPS with 2 writes/tx).

#### JSON-RPC queue promotion (out-of-order nonces)

| Config | TPS | P50ms | P95ms | P99ms | Included | Peak MP | Notes |
|---|---:|---:|---:|---:|---:|---:|---|
| jsonrpc-queue-promotion | 561.7 | 147 | 263 | 274 | 500/500 | 104 | Out-of-order nonces via `eth_sendRawTransaction`, 100% inclusion |

Confirms that Proxy+Priority mempool's queued pool works correctly for Ethereum transactions submitted via JSON-RPC with out-of-order nonces.

### 5. Capability demos

| Test |   TPS | P50ms | P95ms | P99ms | Included | Peak MP | Notes |
|---|------:|------:|------:|------:|---------:|--------:|---|
| queue-promotion |  57.9 |   249 |   339 |   587 |  500/500 |      13 | Out-of-order nonces, 100% inclusion |
| gossip |  49.0 |   180 |   241 |   249 |  250/250 |      10 | All txs to single node, gossip to validator |
| queued-flood | 687.6 |  6684 | 10445 | 10925 |  500/500 |     490 | Nonce gap burst + promotion cascade |
| queued-gap-eviction |     - |     - |     - |     - |        - |       - | Qualitative: gap TTL eviction confirmed |

## Run

All commands assume `cd integration-tests` first. The full workflow has 3 phases:
baselines first, then current-branch benchmarks, then the comparison tests that
load both result sets. Capability demos / queued tests are standalone and can run
any time.

### Phase 1 Collecting baselines (CList mempool)

Build the pre-proxy binary once, then run the three baseline tests.
Results are written to `e2e/benchmark/results/` as JSON keyed by label.

```bash
# Build pre-proxy binary
git checkout tags/v1.2.14
go build -o build/minitiad-baseline ./cmd/minitiad
git checkout -   # return to current branch

cd integration-tests

# Sequential bank send baseline
E2E_MINITIAD_BIN="$(pwd)/../build/minitiad-baseline" \
  go test -v -tags benchmark -run TestBenchmarkBaselineSeq -timeout 30m -count=1 ./e2e/benchmark/

# Burst bank send baseline
E2E_MINITIAD_BIN="$(pwd)/../build/minitiad-baseline" \
  go test -v -tags benchmark -run TestBenchmarkBaselineBurst -timeout 30m -count=1 ./e2e/benchmark/

# Sequential EVM exec baseline
E2E_MINITIAD_BIN="$(pwd)/../build/minitiad-baseline" \
  go test -v -tags benchmark -run TestBenchmarkBaselineSeqEvmExec -timeout 60m -count=1 ./e2e/benchmark/
```

### Phase 2 Running current-branch benchmarks

These use the current binary (auto-built or via `E2E_MINITIAD_BIN`).
Each test writes its own result JSON.

```bash
# Build current binary
go build -o ./minitiad ../cmd/minitiad

# State db comparison (IAVL vs MemIAVL)
E2E_MINITIAD_BIN=./minitiad \
  go test -v -tags benchmark -run TestBenchmarkMemIAVLBankSend -timeout 60m -count=1 ./e2e/benchmark/
E2E_MINITIAD_BIN=./minitiad \
  go test -v -tags benchmark -run TestBenchmarkMemIAVLEvmExec -timeout 60m -count=1 ./e2e/benchmark/

# Capability demos (standalone, no baselines needed)
E2E_MINITIAD_BIN=./minitiad \
  go test -v -tags benchmark -run TestBenchmarkQueuePromotion -timeout 30m -count=1 ./e2e/benchmark/
E2E_MINITIAD_BIN=./minitiad \
  go test -v -tags benchmark -run TestBenchmarkGossipPropagation -timeout 30m -count=1 ./e2e/benchmark/

# Queued mempool behavior (standalone, no baselines needed)
E2E_MINITIAD_BIN=./minitiad \
  go test -v -tags benchmark -run TestBenchmarkQueuedFlood -timeout 30m -count=1 ./e2e/benchmark/
E2E_MINITIAD_BIN=./minitiad \
  go test -v -tags benchmark -run TestBenchmarkQueuedGapEviction -timeout 30m -count=1 ./e2e/benchmark/
```

### Pre-signed HTTP broadcast tests (saturated chain)

These use pre-signed txs via HTTP to saturate the chain, bypassing the CLI bottleneck.

```bash
# Sequential bank send (IAVL vs MemIAVL)
E2E_MINITIAD_BIN=./minitiad \
  go test -v -tags benchmark -run TestBenchmarkPreSignedSeqComparison -timeout 20m -count=1 ./e2e/benchmark/

# Burst bank send (IAVL vs MemIAVL)
E2E_MINITIAD_BIN=./minitiad \
  go test -v -tags benchmark -run TestBenchmarkPreSignedBurstComparison -timeout 20m -count=1 ./e2e/benchmark/

# Sequential EVM exec (IAVL vs MemIAVL)
E2E_MINITIAD_BIN=./minitiad \
  go test -v -tags benchmark -run TestBenchmarkPreSignedSeqEvmExec$ -timeout 30m -count=1 ./e2e/benchmark/

# Sequential EVM exec stress (IAVL vs MemIAVL, 4000 txs)
E2E_MINITIAD_BIN=./minitiad \
  go test -v -tags benchmark -run TestBenchmarkPreSignedSeqEvmExecStress -timeout 30m -count=1 ./e2e/benchmark/
```

### JSON-RPC benchmarks (Ethereum native submission)

These use go-ethereum ECDSA-signed transactions submitted via `eth_sendRawTransaction`.

```bash
# ERC20 transfer benchmark (IAVL vs MemIAVL)
E2E_MINITIAD_BIN=./minitiad \
  go test -v -tags benchmark -run TestBenchmarkJsonRpcErc20Transfer -timeout 30m -count=1 ./e2e/benchmark/

# Sequential EVM contract calls (heavy state)
E2E_MINITIAD_BIN=./minitiad \
  go test -v -tags benchmark -run TestBenchmarkJsonRpcSeqEvmExec -timeout 30m -count=1 ./e2e/benchmark/

# Burst EVM contract calls (heavy state)
E2E_MINITIAD_BIN=./minitiad \
  go test -v -tags benchmark -run TestBenchmarkJsonRpcBurstEvmExec -timeout 30m -count=1 ./e2e/benchmark/

# Queue promotion (out-of-order nonces via JSON-RPC)
E2E_MINITIAD_BIN=./minitiad \
  go test -v -tags benchmark -run TestBenchmarkJsonRpcQueuePromotion -timeout 30m -count=1 ./e2e/benchmark/
```

### Phase 3 Comparison tests (baseline vs current)

These load baseline JSONs from `e2e/benchmark/results/` by label and run Proxy+IAVL
and Proxy+MemIAVL variants, then print a side-by-side comparison table with deltas.

```bash
# Sequential bank send: CList vs Proxy+IAVL vs Proxy+MemIAVL
E2E_MINITIAD_BIN=./minitiad \
  go test -v -tags benchmark -run TestBenchmarkSeqComparison -timeout 30m -count=1 ./e2e/benchmark/

# Sequential EVM exec: CList vs Proxy+IAVL vs Proxy+MemIAVL
E2E_MINITIAD_BIN=./minitiad \
  go test -v -tags benchmark -run TestBenchmarkSeqComparisonEvmExec -timeout 60m -count=1 ./e2e/benchmark/

# Burst bank send: CList vs Proxy+IAVL vs Proxy+MemIAVL
E2E_MINITIAD_BIN=./minitiad \
  go test -v -tags benchmark -run TestBenchmarkBurstComparison -timeout 30m -count=1 ./e2e/benchmark/

# Burst EVM exec: Proxy+IAVL vs Proxy+MemIAVL (no CList since it drops txs under burst)
E2E_MINITIAD_BIN=./minitiad \
  go test -v -tags benchmark -run TestBenchmarkBurstComparisonEvmExec -timeout 60m -count=1 ./e2e/benchmark/
```

Each Phase 3 test prints a comparison table like:

```
Config                    | Variant      |     TPS | vs base |   P50ms | vs base |   P95ms | vs base | Peak Mempool
clist/iavl/seq            | baseline     |   120.5 |       - |    2500 |       - |    4800 |       - |         1950
proxy+priority/iavl/seq   | mempool-only |   245.3 | +103.6% |    1823 |  -27.1% |    3412 |  -28.9% |         1847
proxy+priority/memiavl/seq| combined     |   312.7 | +159.5% |    1401 |  -44.0% |    2845 |  -40.7% |         1823
```

## Configuration

### Ground Rules

1. Baseline requires a separate binary built from v1.2.14 (pre-proxy CometBFT, pre-ABCI++ changes).
2. Run baseline and current benchmarks on the same machine.
3. Warmup runs before every measured load (5 txs, metadata re-queried after).
4. TPS is derived from block timestamps, not submission wall clock.
5. Latency = `block_time - submit_time` (covers mempool wait, gossip, proposal, execution).
6. Load is submitted to edge nodes (non-validator) to test realistic gossip propagation.

### Configurable mempool limits

These can be tuned in `app.toml` under `[abcipp]` (defaults shown):

| Parameter | Default | Description |
|---|---|---|
| `max-queued-per-sender` | 64 | Max queued txs per sender |
| `max-queued-total` | 1024 | Max queued txs globally |
| `queued-gap-ttl` | 60s | TTL for stalled senders missing head nonce |

### Environment variables

| Variable | Default | Description |
|---|---|---|
| `E2E_MINITIAD_BIN` | (auto-build) | Path to prebuilt `minitiad` binary |
| `BENCHMARK_RESULTS_DIR` | `results/` | Output directory for JSON results |

## Structure

```
benchmark/
  config.go          Variant definitions, BenchConfig, preset constructors
  load.go            Load generators
  collector.go       MempoolPoller, CollectResults, latency aggregation
  report.go          JSON output, comparison tables, delta calculations, LoadBaselineResultsByLabel
  benchmark_test.go  Test suite (build-tagged `benchmark`)
  bench_heavy_state/ Solidity contract + Go binding (BenchHeavyState)
  bench_erc20/       Solidity contract + Go binding (BenchERC20, minimal ERC20 for transfer benchmarks)
  results/           JSON output directory
```

### Load generators

All load generators route transactions to fullnodes when `ValidatorCount > 0`.

- **BurstLoad**: All accounts submit concurrently with sequential nonces, round-robin across fullnodes.
- **SequentialLoad**: Accounts run concurrently, but each account sends txs one-at-a-time. Each account pinned to a single fullnode.
- **OutOfOrderLoad**: First 3 txs per account use `[seq+2, seq+0, seq+1]` to test queue promotion.
- **SingleNodeLoad**: All txs to a single node for gossip propagation measurement.
- **EvmExecBurstLoad**: Like BurstLoad but calls `CallContract` (EVM `writeMixed`) instead of bank sends.
- **EvmExecSequentialLoad**: Like SequentialLoad but calls `CallContract`. Each account pinned to a single fullnode.
- **QueuedFloodLoad**: Sends txs with nonces `[base+1..base+N]` (skipping `base+0`), then after all are submitted, sends the gap-filling `base+0` tx to trigger promotion cascade.
- **PreSignedBurstLoad**: Broadcasts pre-signed Cosmos txs via HTTP POST to `/broadcast_tx_sync`. All accounts concurrent, round-robin across fullnodes.
- **PreSignedSequentialLoad**: Broadcasts pre-signed Cosmos txs via HTTP POST. Each account pinned to a single fullnode, txs sent sequentially per account.
- **JsonRpcBurstLoad**: Broadcasts pre-signed Ethereum txs via `eth_sendRawTransaction`. All accounts concurrent, round-robin across fullnode JSON-RPC endpoints.
- **JsonRpcSequentialLoad**: Broadcasts pre-signed Ethereum txs via `eth_sendRawTransaction`. Each account pinned to a single fullnode, txs sent sequentially per account.

### Metrics

| Metric | Source |
|---|---|
| **TPS** | `included_tx_count / block_time_span` |
| **Latency** (avg, p50, p95, p99, max) | `block_timestamp - submit_timestamp` per tx |
| **Peak mempool size** | Goroutine polling `/num_unconfirmed_txs` every 500ms |
| **Per-block tx count** | CometBFT RPC `/block?height=N` |

## EVM exec workload: BenchHeavyState

The EVM exec tests deploy the `BenchHeavyState` Solidity contract at runtime. Each tx calls `writeMixed(sharedCount, localCount)` which performs:

- **shared writes** to a global `sharedState` mapping.
- **local writes** to the caller's own `localState` mapping.

Each call writes to **unique keys** using a per-sender nonce, so the state tree grows continuously. This creates IAVL rebalancing pressure that MemIAVL handles more efficiently.

CLI-based tests use `writeMixed(5, 25)` = 30 writes/tx. Pre-signed HTTP tests use `writeMixed(20, 80)` = 100 writes/tx.

## Regenerating the Solidity bindings

```bash
# BenchHeavyState
cd bench_heavy_state
solc --abi --bin BenchHeavyState.sol -o build/
abigen --abi build/BenchHeavyState.abi --bin build/BenchHeavyState.bin \
  --pkg bench_heavy_state --out BenchHeavyState.go
rm -rf build/

# BenchERC20
cd ../bench_erc20
solc --abi --bin BenchERC20.sol -o build/
abigen --abi build/BenchERC20.abi --bin build/BenchERC20.bin \
  --pkg bench_erc20 --out BenchERC20.go
rm -rf build/
```
