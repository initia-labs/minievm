package benchmark

import (
	"context"
	"encoding/hex"
	"math/big"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/initia-labs/minievm/integration-tests/e2e/benchmark/bench_erc20"
	"github.com/initia-labs/minievm/integration-tests/e2e/cluster"
)

// TxSubmission records a single submitted transaction.
type TxSubmission struct {
	TxHash     string    `json:"tx_hash"`
	Account    string    `json:"account"`
	Sequence   uint64    `json:"sequence"`
	SubmitTime time.Time `json:"submit_time"`
	ViaNode    int       `json:"via_node"`
	Code       int64     `json:"code,omitempty"`
}

// LoadResult holds the outcome of a load generation run.
type LoadResult struct {
	Submissions []TxSubmission
	Errors      []error
	StartTime   time.Time
	EndTime     time.Time
}

// BurstLoad submits all transactions concurrently across accounts with sequential nonces.
func BurstLoad(ctx context.Context, cl *cluster.Cluster, cfg BenchConfig, metas map[string]cluster.AccountMeta) LoadResult {
	if cfg.NodeCount <= 0 {
		panic("BurstLoad: cfg.NodeCount must be > 0")
	}

	var (
		mu     sync.Mutex
		wg     sync.WaitGroup
		result LoadResult
	)

	result.StartTime = time.Now()

	for _, name := range cl.AccountNames() {
		meta := metas[name]

		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < cfg.TxPerAccount; i++ {
				select {
				case <-ctx.Done():
					return
				default:
				}

				seq := meta.Sequence + uint64(i) //nolint:gosec // i is bounded by TxPerAccount
				viaNode := i % cfg.NodeCount
				if cfg.ValidatorCount > 0 && cfg.ValidatorCount < cfg.NodeCount {
					viaNode = edgeNodeIndex(i, cfg.NodeCount, cfg.ValidatorCount)
				}
				submitTime := time.Now()

				res := cl.SendBankTxWithSequence(
					ctx, name, cl.ValidatorAddress(), "1GAS",
					meta.AccountNumber, seq, cfg.GetGasLimit(), viaNode,
				)

				sub := TxSubmission{
					TxHash:     res.TxHash,
					Account:    name,
					Sequence:   seq,
					SubmitTime: submitTime,
					ViaNode:    viaNode,
					Code:       res.Code,
				}

				mu.Lock()
				result.Submissions = append(result.Submissions, sub)
				if res.Err != nil {
					result.Errors = append(result.Errors, res.Err)
				}
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	result.EndTime = time.Now()

	return result
}

// SequentialLoad submits transactions sequentially per account, pinned to a single node.
func SequentialLoad(ctx context.Context, cl *cluster.Cluster, cfg BenchConfig, metas map[string]cluster.AccountMeta) LoadResult {
	if cfg.NodeCount <= 0 {
		panic("SequentialLoad: cfg.NodeCount must be > 0")
	}

	var (
		mu     sync.Mutex
		wg     sync.WaitGroup
		result LoadResult
	)

	result.StartTime = time.Now()

	for accountIdx, name := range cl.AccountNames() {
		meta := metas[name]
		viaNode := accountIdx % cfg.NodeCount
		if cfg.ValidatorCount > 0 && cfg.ValidatorCount < cfg.NodeCount {
			viaNode = edgeNodeIndex(accountIdx, cfg.NodeCount, cfg.ValidatorCount)
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < cfg.TxPerAccount; i++ {
				select {
				case <-ctx.Done():
					return
				default:
				}

				seq := meta.Sequence + uint64(i) //nolint:gosec // i is bounded by TxPerAccount
				submitTime := time.Now()

				res := cl.SendBankTxWithSequence(
					ctx, name, cl.ValidatorAddress(), "1GAS",
					meta.AccountNumber, seq, cfg.GetGasLimit(), viaNode,
				)

				sub := TxSubmission{
					TxHash:     res.TxHash,
					Account:    name,
					Sequence:   seq,
					SubmitTime: submitTime,
					ViaNode:    viaNode,
					Code:       res.Code,
				}

				mu.Lock()
				result.Submissions = append(result.Submissions, sub)
				if res.Err != nil {
					result.Errors = append(result.Errors, res.Err)
				}
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	result.EndTime = time.Now()

	return result
}

// OutOfOrderLoad submits the first 3 txs per account with out-of-order nonces
// (seq+2, seq+0, seq+1), then the rest sequentially.
func OutOfOrderLoad(ctx context.Context, cl *cluster.Cluster, cfg BenchConfig, metas map[string]cluster.AccountMeta) LoadResult {
	if cfg.NodeCount <= 0 {
		panic("OutOfOrderLoad: cfg.NodeCount must be > 0")
	}
	if cfg.TxPerAccount < 3 {
		panic("OutOfOrderLoad: TxPerAccount must be >= 3 for out-of-order pattern")
	}

	var (
		mu     sync.Mutex
		wg     sync.WaitGroup
		result LoadResult
	)

	result.StartTime = time.Now()

	for _, name := range cl.AccountNames() {
		meta := metas[name]

		wg.Add(1)
		go func() {
			defer wg.Done()
			seqs := sequencePattern(meta.Sequence, cfg.TxPerAccount)
			for i, seq := range seqs {
				select {
				case <-ctx.Done():
					return
				default:
				}

				viaNode := i % cfg.NodeCount
				if cfg.ValidatorCount > 0 && cfg.ValidatorCount < cfg.NodeCount {
					viaNode = edgeNodeIndex(i, cfg.NodeCount, cfg.ValidatorCount)
				}
				submitTime := time.Now()

				res := cl.SendBankTxWithSequence(
					ctx, name, cl.ValidatorAddress(), "1GAS",
					meta.AccountNumber, seq, cfg.GetGasLimit(), viaNode,
				)

				sub := TxSubmission{
					TxHash:     res.TxHash,
					Account:    name,
					Sequence:   seq,
					SubmitTime: submitTime,
					ViaNode:    viaNode,
					Code:       res.Code,
				}

				mu.Lock()
				result.Submissions = append(result.Submissions, sub)
				if res.Err != nil {
					result.Errors = append(result.Errors, res.Err)
				}
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	result.EndTime = time.Now()

	return result
}

// SingleNodeLoad submits all transactions to a single specified node.
func SingleNodeLoad(ctx context.Context, cl *cluster.Cluster, cfg BenchConfig, metas map[string]cluster.AccountMeta, targetNode int) LoadResult {
	if targetNode < 0 || targetNode >= cfg.NodeCount {
		panic("SingleNodeLoad: targetNode out of range")
	}

	var (
		mu     sync.Mutex
		wg     sync.WaitGroup
		result LoadResult
	)

	result.StartTime = time.Now()

	for _, name := range cl.AccountNames() {
		meta := metas[name]

		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < cfg.TxPerAccount; i++ {
				select {
				case <-ctx.Done():
					return
				default:
				}

				seq := meta.Sequence + uint64(i) //nolint:gosec // i is bounded by TxPerAccount
				submitTime := time.Now()

				res := cl.SendBankTxWithSequence(
					ctx, name, cl.ValidatorAddress(), "1GAS",
					meta.AccountNumber, seq, cfg.GetGasLimit(), targetNode,
				)

				sub := TxSubmission{
					TxHash:     res.TxHash,
					Account:    name,
					Sequence:   seq,
					SubmitTime: submitTime,
					ViaNode:    targetNode,
					Code:       res.Code,
				}

				mu.Lock()
				result.Submissions = append(result.Submissions, sub)
				if res.Err != nil {
					result.Errors = append(result.Errors, res.Err)
				}
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	result.EndTime = time.Now()

	return result
}

// EvmExecSequentialLoad returns a LoadFn that submits EVM call transactions
// sequentially per account, pinned to a single node.
func EvmExecSequentialLoad(contractAddr, inputHex string, gasLimit uint64) func(ctx context.Context, cl *cluster.Cluster, cfg BenchConfig, metas map[string]cluster.AccountMeta) LoadResult {
	return func(ctx context.Context, cl *cluster.Cluster, cfg BenchConfig, metas map[string]cluster.AccountMeta) LoadResult {
		if cfg.NodeCount <= 0 {
			panic("EvmExecSequentialLoad: cfg.NodeCount must be > 0")
		}

		var (
			mu     sync.Mutex
			wg     sync.WaitGroup
			result LoadResult
		)

		result.StartTime = time.Now()

		for accountIdx, name := range cl.AccountNames() {
			meta := metas[name]
			viaNode := accountIdx % cfg.NodeCount
			if cfg.ValidatorCount > 0 && cfg.ValidatorCount < cfg.NodeCount {
				viaNode = edgeNodeIndex(accountIdx, cfg.NodeCount, cfg.ValidatorCount)
			}

			wg.Add(1)
			go func() {
				defer wg.Done()
				for i := 0; i < cfg.TxPerAccount; i++ {
					select {
					case <-ctx.Done():
						return
					default:
					}

					seq := meta.Sequence + uint64(i) //nolint:gosec // i is bounded by TxPerAccount
					submitTime := time.Now()

					res := cl.CallContract(
						ctx, name, contractAddr, inputHex,
						meta.AccountNumber, seq, gasLimit, viaNode,
					)

					sub := TxSubmission{
						TxHash:     res.TxHash,
						Account:    name,
						Sequence:   seq,
						SubmitTime: submitTime,
						ViaNode:    viaNode,
						Code:       res.Code,
					}

					mu.Lock()
					result.Submissions = append(result.Submissions, sub)
					if res.Err != nil {
						result.Errors = append(result.Errors, res.Err)
					}
					mu.Unlock()
				}
			}()
		}

		wg.Wait()
		result.EndTime = time.Now()

		return result
	}
}

// EvmExecBurstLoad returns a LoadFn that submits EVM call transactions concurrently.
func EvmExecBurstLoad(contractAddr, inputHex string, gasLimit uint64) func(ctx context.Context, cl *cluster.Cluster, cfg BenchConfig, metas map[string]cluster.AccountMeta) LoadResult {
	return func(ctx context.Context, cl *cluster.Cluster, cfg BenchConfig, metas map[string]cluster.AccountMeta) LoadResult {
		if cfg.NodeCount <= 0 {
			panic("EvmExecBurstLoad: cfg.NodeCount must be > 0")
		}

		var (
			mu     sync.Mutex
			wg     sync.WaitGroup
			result LoadResult
		)

		result.StartTime = time.Now()

		for _, name := range cl.AccountNames() {
			meta := metas[name]

			wg.Add(1)
			go func() {
				defer wg.Done()
				for i := 0; i < cfg.TxPerAccount; i++ {
					select {
					case <-ctx.Done():
						return
					default:
					}

					seq := meta.Sequence + uint64(i) //nolint:gosec // i is bounded by TxPerAccount
					viaNode := i % cfg.NodeCount
					if cfg.ValidatorCount > 0 && cfg.ValidatorCount < cfg.NodeCount {
						viaNode = edgeNodeIndex(i, cfg.NodeCount, cfg.ValidatorCount)
					}
					submitTime := time.Now()

					res := cl.CallContract(
						ctx, name, contractAddr, inputHex,
						meta.AccountNumber, seq, gasLimit, viaNode,
					)

					sub := TxSubmission{
						TxHash:     res.TxHash,
						Account:    name,
						Sequence:   seq,
						SubmitTime: submitTime,
						ViaNode:    viaNode,
						Code:       res.Code,
					}

					mu.Lock()
					result.Submissions = append(result.Submissions, sub)
					if res.Err != nil {
						result.Errors = append(result.Errors, res.Err)
					}
					mu.Unlock()
				}
			}()
		}

		wg.Wait()
		result.EndTime = time.Now()

		return result
	}
}

// QueuedFloodLoad submits txs with nonces [base+1..base+N] (skipping base+0),
// then fills the gap to trigger a promotion cascade.
func QueuedFloodLoad(ctx context.Context, cl *cluster.Cluster, cfg BenchConfig, metas map[string]cluster.AccountMeta) LoadResult {
	if cfg.NodeCount <= 0 {
		panic("QueuedFloodLoad: cfg.NodeCount must be > 0")
	}
	if cfg.TxPerAccount < 2 {
		panic("QueuedFloodLoad: TxPerAccount must be >= 2")
	}

	var (
		mu     sync.Mutex
		wg     sync.WaitGroup
		result LoadResult
	)

	result.StartTime = time.Now()

	// submit future-nonce txs [base+1..base+N-1]
	for _, name := range cl.AccountNames() {
		meta := metas[name]

		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 1; i < cfg.TxPerAccount; i++ {
				select {
				case <-ctx.Done():
					return
				default:
				}

				seq := meta.Sequence + uint64(i) //nolint:gosec // i is bounded by TxPerAccount
				viaNode := i % cfg.NodeCount
				if cfg.ValidatorCount > 0 && cfg.ValidatorCount < cfg.NodeCount {
					viaNode = edgeNodeIndex(i, cfg.NodeCount, cfg.ValidatorCount)
				}
				submitTime := time.Now()

				res := cl.SendBankTxWithSequence(
					ctx, name, cl.ValidatorAddress(), "1GAS",
					meta.AccountNumber, seq, cfg.GetGasLimit(), viaNode,
				)

				sub := TxSubmission{
					TxHash:     res.TxHash,
					Account:    name,
					Sequence:   seq,
					SubmitTime: submitTime,
					ViaNode:    viaNode,
					Code:       res.Code,
				}

				mu.Lock()
				result.Submissions = append(result.Submissions, sub)
				if res.Err != nil {
					result.Errors = append(result.Errors, res.Err)
				}
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	time.Sleep(2 * time.Second)

	// send the gap-filling base+0 tx for each account
	for _, name := range cl.AccountNames() {
		meta := metas[name]
		viaNode := 0
		if cfg.ValidatorCount > 0 && cfg.ValidatorCount < cfg.NodeCount {
			viaNode = edgeNodeIndex(0, cfg.NodeCount, cfg.ValidatorCount)
		}
		submitTime := time.Now()

		res := cl.SendBankTxWithSequence(
			ctx, name, cl.ValidatorAddress(), "1GAS",
			meta.AccountNumber, meta.Sequence, cfg.GetGasLimit(), viaNode,
		)

		sub := TxSubmission{
			TxHash:     res.TxHash,
			Account:    name,
			Sequence:   meta.Sequence,
			SubmitTime: submitTime,
			ViaNode:    viaNode,
			Code:       res.Code,
		}

		mu.Lock()
		result.Submissions = append(result.Submissions, sub)
		if res.Err != nil {
			result.Errors = append(result.Errors, res.Err)
		}
		mu.Unlock()
	}

	result.EndTime = time.Now()

	return result
}

// QueuedGapLoad submits txs with future nonces and never fills the gap.
func QueuedGapLoad(ctx context.Context, cl *cluster.Cluster, cfg BenchConfig, metas map[string]cluster.AccountMeta) LoadResult {
	if cfg.NodeCount <= 0 {
		panic("QueuedGapLoad: cfg.NodeCount must be > 0")
	}

	var (
		mu     sync.Mutex
		wg     sync.WaitGroup
		result LoadResult
	)

	result.StartTime = time.Now()

	for _, name := range cl.AccountNames() {
		meta := metas[name]

		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 1; i <= cfg.TxPerAccount; i++ {
				select {
				case <-ctx.Done():
					return
				default:
				}

				seq := meta.Sequence + uint64(i) //nolint:gosec // i is bounded by TxPerAccount
				viaNode := i % cfg.NodeCount
				if cfg.ValidatorCount > 0 && cfg.ValidatorCount < cfg.NodeCount {
					viaNode = edgeNodeIndex(i, cfg.NodeCount, cfg.ValidatorCount)
				}
				submitTime := time.Now()

				res := cl.SendBankTxWithSequence(
					ctx, name, cl.ValidatorAddress(), "1GAS",
					meta.AccountNumber, seq, cfg.GetGasLimit(), viaNode,
				)

				sub := TxSubmission{
					TxHash:     res.TxHash,
					Account:    name,
					Sequence:   seq,
					SubmitTime: submitTime,
					ViaNode:    viaNode,
					Code:       res.Code,
				}

				mu.Lock()
				result.Submissions = append(result.Submissions, sub)
				if res.Err != nil {
					result.Errors = append(result.Errors, res.Err)
				}
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	result.EndTime = time.Now()

	return result
}

// PreSignBankTxs generates and signs all bank send transactions offline in parallel.
func PreSignBankTxs(ctx context.Context, t *testing.T, cl *cluster.Cluster, cfg BenchConfig, metas map[string]cluster.AccountMeta) []cluster.SignedTx {
	t.Helper()
	total := cfg.TotalTx()
	txs := make([]cluster.SignedTx, 0, total)
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, name := range cl.AccountNames() {
		meta := metas[name]
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < cfg.TxPerAccount; i++ {
				seq := meta.Sequence + uint64(i) //nolint:gosec // i is bounded by TxPerAccount
				signed, err := cl.GenerateSignedBankTx(
					ctx, name, cl.ValidatorAddress(), "1GAS",
					meta.AccountNumber, seq, cfg.GetGasLimit(),
				)
				if err != nil {
					t.Logf("[pre-sign] failed from=%s seq=%d err=%v", name, seq, err)
					continue
				}
				mu.Lock()
				txs = append(txs, signed)
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	t.Logf("Pre-signed %d/%d bank txs", len(txs), total)

	return txs
}

// PreSignEvmCallTxs generates and signs all EVM call transactions offline in parallel.
func PreSignEvmCallTxs(
	ctx context.Context, t *testing.T, cl *cluster.Cluster, cfg BenchConfig, metas map[string]cluster.AccountMeta,
	contractAddr, inputHex string, gasLimit uint64,
) []cluster.SignedTx {
	t.Helper()
	total := cfg.TotalTx()
	txs := make([]cluster.SignedTx, 0, total)
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, name := range cl.AccountNames() {
		meta := metas[name]
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < cfg.TxPerAccount; i++ {
				seq := meta.Sequence + uint64(i) //nolint:gosec // i is bounded by TxPerAccount
				signed, err := cl.GenerateSignedEvmCallTx(
					ctx, name, contractAddr, inputHex,
					meta.AccountNumber, seq, gasLimit,
				)
				if err != nil {
					t.Logf("[pre-sign-evm] failed from=%s seq=%d err=%v", name, seq, err)
					continue
				}
				mu.Lock()
				txs = append(txs, signed)
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	t.Logf("Pre-signed %d/%d evm call txs", len(txs), total)

	return txs
}

// PreSignedBurstLoad broadcasts pre-signed transactions via HTTP as fast as possible.
func PreSignedBurstLoad(signedTxs []cluster.SignedTx) func(ctx context.Context, cl *cluster.Cluster, cfg BenchConfig, metas map[string]cluster.AccountMeta) LoadResult {
	return func(ctx context.Context, cl *cluster.Cluster, cfg BenchConfig, _ map[string]cluster.AccountMeta) LoadResult {
		if cfg.NodeCount <= 0 {
			panic("PreSignedBurstLoad: cfg.NodeCount must be > 0")
		}

		byAccount := make(map[string][]cluster.SignedTx)
		for _, tx := range signedTxs {
			byAccount[tx.Account] = append(byAccount[tx.Account], tx)
		}

		for acct := range byAccount {
			txs := byAccount[acct]
			for i := 1; i < len(txs); i++ {
				for j := i; j > 0 && txs[j].Sequence < txs[j-1].Sequence; j-- {
					txs[j], txs[j-1] = txs[j-1], txs[j]
				}
			}
		}

		var (
			mu     sync.Mutex
			wg     sync.WaitGroup
			result LoadResult
		)

		result.StartTime = time.Now()

		for acct, txs := range byAccount {
			_ = acct
			wg.Add(1)
			go func() {
				defer wg.Done()
				for i, stx := range txs {
					select {
					case <-ctx.Done():
						return
					default:
					}

					viaNode := i % cfg.NodeCount
					if cfg.ValidatorCount > 0 && cfg.ValidatorCount < cfg.NodeCount {
						viaNode = edgeNodeIndex(i, cfg.NodeCount, cfg.ValidatorCount)
					}
					submitTime := time.Now()

					res, err := cl.BroadcastTxSync(ctx, viaNode, stx.TxBase64)

					sub := TxSubmission{
						TxHash:     stx.TxHash,
						Account:    stx.Account,
						Sequence:   stx.Sequence,
						SubmitTime: submitTime,
						ViaNode:    viaNode,
					}
					if err == nil {
						sub.Code = res.Code
						if res.TxHash != "" {
							sub.TxHash = res.TxHash
						}
					}

					mu.Lock()
					result.Submissions = append(result.Submissions, sub)
					if err != nil {
						result.Errors = append(result.Errors, err)
					} else if res.Err != nil {
						result.Errors = append(result.Errors, res.Err)
					}
					mu.Unlock()
				}
			}()
		}

		wg.Wait()
		result.EndTime = time.Now()

		return result
	}
}

// PreSignedSequentialLoad broadcasts pre-signed transactions via HTTP, one at a time
// per account, pinned to a single node.
func PreSignedSequentialLoad(signedTxs []cluster.SignedTx) func(ctx context.Context, cl *cluster.Cluster, cfg BenchConfig, metas map[string]cluster.AccountMeta) LoadResult {
	return func(ctx context.Context, cl *cluster.Cluster, cfg BenchConfig, _ map[string]cluster.AccountMeta) LoadResult {
		if cfg.NodeCount <= 0 {
			panic("PreSignedSequentialLoad: cfg.NodeCount must be > 0")
		}

		byAccount := make(map[string][]cluster.SignedTx)
		for _, tx := range signedTxs {
			byAccount[tx.Account] = append(byAccount[tx.Account], tx)
		}

		for acct := range byAccount {
			txs := byAccount[acct]
			for i := 1; i < len(txs); i++ {
				for j := i; j > 0 && txs[j].Sequence < txs[j-1].Sequence; j-- {
					txs[j], txs[j-1] = txs[j-1], txs[j]
				}
			}
		}

		var (
			mu     sync.Mutex
			wg     sync.WaitGroup
			result LoadResult
		)

		result.StartTime = time.Now()

		accountIdx := 0
		for acct, txs := range byAccount {
			_ = acct
			viaNode := accountIdx % cfg.NodeCount
			if cfg.ValidatorCount > 0 && cfg.ValidatorCount < cfg.NodeCount {
				viaNode = edgeNodeIndex(accountIdx, cfg.NodeCount, cfg.ValidatorCount)
			}

			wg.Add(1)
			go func() {
				defer wg.Done()
				for i, stx := range txs {
					select {
					case <-ctx.Done():
						return
					default:
					}

					submitTime := time.Now()
					res, err := cl.BroadcastTxSync(ctx, viaNode, stx.TxBase64)

					sub := TxSubmission{
						TxHash:     stx.TxHash,
						Account:    stx.Account,
						Sequence:   stx.Sequence,
						SubmitTime: submitTime,
						ViaNode:    viaNode,
					}
					if err == nil {
						sub.Code = res.Code
						if res.TxHash != "" {
							sub.TxHash = res.TxHash
						}
					}

					mu.Lock()
					result.Submissions = append(result.Submissions, sub)
					if err != nil {
						result.Errors = append(result.Errors, err)
					} else if res.Err != nil {
						result.Errors = append(result.Errors, res.Err)
					}
					mu.Unlock()

					if i < len(txs)-1 {
						time.Sleep(5 * time.Millisecond)
					}
				}
			}()
			accountIdx++
		}

		wg.Wait()
		result.EndTime = time.Now()

		return result
	}
}

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

func edgeNodeIndex(i, nodeCount, validatorCount int) int {
	edgeCount := nodeCount - validatorCount
	if edgeCount <= 0 {
		return i % nodeCount
	}
	return validatorCount + (i % edgeCount)
}

func sequencePattern(base uint64, count int) []uint64 {
	seqs := []uint64{base + 2, base, base + 1}
	if count <= 3 {
		return seqs[:count]
	}

	for i := 3; i < count; i++ {
		seqs = append(seqs, base+uint64(i)) //nolint:gosec // i is bounded by count
	}

	return seqs
}

// Warmup sends a small number of transactions to warm up the cluster.
func Warmup(ctx context.Context, cl *cluster.Cluster, metas map[string]cluster.AccountMeta) {
	names := cl.AccountNames()
	for i := 0; i < 5 && i < len(names); i++ {
		name := names[i]
		meta := metas[name]
		viaNode := i % cl.NodeCount()
		cl.SendBankTxWithSequence(
			ctx, name, cl.ValidatorAddress(), "1GAS",
			meta.AccountNumber, meta.Sequence, defaultGasLimit, viaNode,
		)
	}
}

// --------
// JSON-RPC
// --------

// EthNonceMeta holds the starting nonce for an Ethereum account.
type EthNonceMeta struct {
	Account cluster.EthAccount
	Nonce   uint64
}

// CollectEthNonces queries the current nonce for each Ethereum account.
func CollectEthNonces(ctx context.Context, cl *cluster.Cluster) (map[string]EthNonceMeta, error) {
	metas := make(map[string]EthNonceMeta)
	for _, acct := range cl.EthAccounts() {
		nonce, err := cl.EthGetTransactionCount(ctx, 0, acct.Address.Hex())
		if err != nil {
			return nil, err
		}
		metas[acct.Name] = EthNonceMeta{Account: acct, Nonce: nonce}
	}
	return metas, nil
}

// PreSignEthContractCallTxs generates and signs all ETH contract call transactions.
func PreSignEthContractCallTxs(
	ctx context.Context, t *testing.T, cl *cluster.Cluster, cfg BenchConfig,
	nonces map[string]EthNonceMeta, contractAddr common.Address, inputHex string, gasLimit uint64,
) []cluster.SignedEthTx {
	t.Helper()
	total := cfg.TotalTx()
	txs := make([]cluster.SignedEthTx, 0, total)
	var mu sync.Mutex
	var wg sync.WaitGroup

	cleanHex := strings.TrimPrefix(inputHex, "0x")
	inputData, err := hex.DecodeString(cleanHex)
	if err != nil {
		t.Fatalf("decode input hex: %v", err)
	}

	for _, acct := range cl.EthAccounts() {
		meta := nonces[acct.Name]
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < cfg.TxPerAccount; i++ {
				nonce := meta.Nonce + uint64(i) //nolint:gosec
				signed, err := cl.SignEthContractCall(meta.Account, contractAddr, inputData, nonce, gasLimit)
				if err != nil {
					t.Logf("[pre-sign-eth-call] failed from=%s nonce=%d err=%v", meta.Account.Name, nonce, err)
					continue
				}
				mu.Lock()
				txs = append(txs, signed)
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	t.Logf("Pre-signed %d/%d eth contract call txs", len(txs), total)

	return txs
}

// PreSignEthContractCallsOutOfOrder signs ETH contract call txs with the first 3 nonces shuffled
// [base+2, base+0, base+1, base+3, base+4, ...] to test queued tx promotion via JSON-RPC.
func PreSignEthContractCallsOutOfOrder(
	ctx context.Context, t *testing.T, cl *cluster.Cluster, cfg BenchConfig,
	nonces map[string]EthNonceMeta, contractAddr common.Address, inputHex string, gasLimit uint64,
) []cluster.SignedEthTx {
	t.Helper()
	total := cfg.TotalTx()
	txs := make([]cluster.SignedEthTx, 0, total)
	var mu sync.Mutex
	var wg sync.WaitGroup

	cleanHex := strings.TrimPrefix(inputHex, "0x")
	inputData, err := hex.DecodeString(cleanHex)
	if err != nil {
		t.Fatalf("decode input hex: %v", err)
	}

	for _, acct := range cl.EthAccounts() {
		meta := nonces[acct.Name]
		wg.Add(1)
		go func() {
			defer wg.Done()
			seqs := sequencePattern(meta.Nonce, cfg.TxPerAccount)
			for _, nonce := range seqs {
				signed, err := cl.SignEthContractCall(meta.Account, contractAddr, inputData, nonce, gasLimit)
				if err != nil {
					t.Logf("[pre-sign-eth-ooo] failed from=%s nonce=%d err=%v", meta.Account.Name, nonce, err)
					continue
				}
				mu.Lock()
				txs = append(txs, signed)
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	t.Logf("Pre-signed %d/%d eth out-of-order contract call txs", len(txs), total)

	return txs
}

// JsonRpcBurstLoad broadcasts pre-signed Ethereum transactions via JSON-RPC concurrently.
func JsonRpcBurstLoad(signedEthTxs []cluster.SignedEthTx) func(ctx context.Context, cl *cluster.Cluster, cfg BenchConfig, metas map[string]cluster.AccountMeta) LoadResult {
	return func(ctx context.Context, cl *cluster.Cluster, cfg BenchConfig, _ map[string]cluster.AccountMeta) LoadResult {
		if cfg.NodeCount <= 0 {
			panic("JsonRpcBurstLoad: cfg.NodeCount must be > 0")
		}

		byAccount := make(map[string][]cluster.SignedEthTx)
		for _, tx := range signedEthTxs {
			byAccount[tx.Account] = append(byAccount[tx.Account], tx)
		}
		for acct := range byAccount {
			txs := byAccount[acct]
			for i := 1; i < len(txs); i++ {
				for j := i; j > 0 && txs[j].Nonce < txs[j-1].Nonce; j-- {
					txs[j], txs[j-1] = txs[j-1], txs[j]
				}
			}
		}

		var (
			mu     sync.Mutex
			wg     sync.WaitGroup
			result LoadResult
		)

		result.StartTime = time.Now()

		for _, txs := range byAccount {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for i, stx := range txs {
					select {
					case <-ctx.Done():
						return
					default:
					}

					viaNode := i % cfg.NodeCount
					if cfg.ValidatorCount > 0 && cfg.ValidatorCount < cfg.NodeCount {
						viaNode = edgeNodeIndex(i, cfg.NodeCount, cfg.ValidatorCount)
					}
					submitTime := time.Now()

					txHash, err := cl.EthSendRawTransaction(ctx, viaNode, stx.RawHex)

					sub := TxSubmission{
						TxHash:     stx.TxHash,
						Account:    stx.Account,
						Sequence:   stx.Nonce,
						SubmitTime: submitTime,
						ViaNode:    viaNode,
					}
					if err == nil && txHash != "" {
						sub.TxHash = txHash
					}

					mu.Lock()
					result.Submissions = append(result.Submissions, sub)
					if err != nil {
						result.Errors = append(result.Errors, err)
					}
					mu.Unlock()
				}
			}()
		}

		wg.Wait()
		result.EndTime = time.Now()

		return result
	}
}

// JsonRpcPreserveOrderLoad broadcasts pre-signed Ethereum transactions via JSON-RPC
// in the exact order they appear per account (no nonce sorting). This is used to
// test out-of-order nonce submission for queue promotion.
func JsonRpcPreserveOrderLoad(signedEthTxs []cluster.SignedEthTx) func(ctx context.Context, cl *cluster.Cluster, cfg BenchConfig, metas map[string]cluster.AccountMeta) LoadResult {
	return func(ctx context.Context, cl *cluster.Cluster, cfg BenchConfig, _ map[string]cluster.AccountMeta) LoadResult {
		if cfg.NodeCount <= 0 {
			panic("JsonRpcPreserveOrderLoad: cfg.NodeCount must be > 0")
		}

		byAccount := make(map[string][]cluster.SignedEthTx)
		for _, tx := range signedEthTxs {
			byAccount[tx.Account] = append(byAccount[tx.Account], tx)
		}

		var (
			mu     sync.Mutex
			wg     sync.WaitGroup
			result LoadResult
		)

		result.StartTime = time.Now()

		accountIdx := 0
		for _, txs := range byAccount {
			viaNode := accountIdx % cfg.NodeCount
			if cfg.ValidatorCount > 0 && cfg.ValidatorCount < cfg.NodeCount {
				viaNode = edgeNodeIndex(accountIdx, cfg.NodeCount, cfg.ValidatorCount)
			}

			wg.Add(1)
			go func() {
				defer wg.Done()
				for i, stx := range txs {
					select {
					case <-ctx.Done():
						return
					default:
					}

					submitTime := time.Now()
					txHash, err := cl.EthSendRawTransaction(ctx, viaNode, stx.RawHex)

					sub := TxSubmission{
						TxHash:     stx.TxHash,
						Account:    stx.Account,
						Sequence:   stx.Nonce,
						SubmitTime: submitTime,
						ViaNode:    viaNode,
					}
					if err == nil && txHash != "" {
						sub.TxHash = txHash
					}

					mu.Lock()
					result.Submissions = append(result.Submissions, sub)
					if err != nil {
						result.Errors = append(result.Errors, err)
					}
					mu.Unlock()

					if i < len(txs)-1 {
						time.Sleep(5 * time.Millisecond)
					}
				}
			}()
			accountIdx++
		}

		wg.Wait()
		result.EndTime = time.Now()

		return result
	}
}

// JsonRpcSequentialLoad broadcasts pre-signed Ethereum transactions via JSON-RPC,
// sequentially per account, pinned to a single node.
func JsonRpcSequentialLoad(signedEthTxs []cluster.SignedEthTx) func(ctx context.Context, cl *cluster.Cluster, cfg BenchConfig, metas map[string]cluster.AccountMeta) LoadResult {
	return func(ctx context.Context, cl *cluster.Cluster, cfg BenchConfig, _ map[string]cluster.AccountMeta) LoadResult {
		if cfg.NodeCount <= 0 {
			panic("JsonRpcSequentialLoad: cfg.NodeCount must be > 0")
		}

		byAccount := make(map[string][]cluster.SignedEthTx)
		for _, tx := range signedEthTxs {
			byAccount[tx.Account] = append(byAccount[tx.Account], tx)
		}
		for acct := range byAccount {
			txs := byAccount[acct]
			for i := 1; i < len(txs); i++ {
				for j := i; j > 0 && txs[j].Nonce < txs[j-1].Nonce; j-- {
					txs[j], txs[j-1] = txs[j-1], txs[j]
				}
			}
		}

		var (
			mu     sync.Mutex
			wg     sync.WaitGroup
			result LoadResult
		)

		result.StartTime = time.Now()

		accountIdx := 0
		for _, txs := range byAccount {
			viaNode := accountIdx % cfg.NodeCount
			if cfg.ValidatorCount > 0 && cfg.ValidatorCount < cfg.NodeCount {
				viaNode = edgeNodeIndex(accountIdx, cfg.NodeCount, cfg.ValidatorCount)
			}

			wg.Add(1)
			go func() {
				defer wg.Done()
				for i, stx := range txs {
					select {
					case <-ctx.Done():
						return
					default:
					}

					submitTime := time.Now()
					txHash, err := cl.EthSendRawTransaction(ctx, viaNode, stx.RawHex)

					sub := TxSubmission{
						TxHash:     stx.TxHash,
						Account:    stx.Account,
						Sequence:   stx.Nonce,
						SubmitTime: submitTime,
						ViaNode:    viaNode,
					}
					if err == nil && txHash != "" {
						sub.TxHash = txHash
					}

					mu.Lock()
					result.Submissions = append(result.Submissions, sub)
					if err != nil {
						result.Errors = append(result.Errors, err)
					}
					mu.Unlock()

					if i < len(txs)-1 {
						time.Sleep(5 * time.Millisecond)
					}
				}
			}()
			accountIdx++
		}

		wg.Wait()
		result.EndTime = time.Now()

		return result
	}
}

// PreSignEthErc20TransferTxs generates and signs ERC20 transfer txs.
func PreSignEthErc20TransferTxs(
	ctx context.Context, t *testing.T, cl *cluster.Cluster, cfg BenchConfig,
	nonces map[string]EthNonceMeta, contractAddr common.Address, transferData []byte, gasLimit uint64,
) []cluster.SignedEthTx {
	t.Helper()
	total := cfg.TotalTx()
	txs := make([]cluster.SignedEthTx, 0, total)
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, acct := range cl.EthAccounts() {
		meta := nonces[acct.Name]
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < cfg.TxPerAccount; i++ {
				nonce := meta.Nonce + uint64(i) //nolint:gosec
				signed, err := cl.SignEthContractCall(meta.Account, contractAddr, transferData, nonce, gasLimit)
				if err != nil {
					t.Logf("[pre-sign-erc20] failed from=%s nonce=%d err=%v", meta.Account.Name, nonce, err)
					continue
				}
				mu.Lock()
				txs = append(txs, signed)
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	t.Logf("Pre-signed %d/%d erc20 transfer txs", len(txs), total)

	return txs
}

// MintErc20ForAccounts mints ERC20 tokens to each ETH account via JSON-RPC.
func MintErc20ForAccounts(
	ctx context.Context, t *testing.T, cl *cluster.Cluster,
	nonces map[string]EthNonceMeta, contractAddr common.Address, mintAmount *big.Int, gasLimit uint64,
) {
	t.Helper()
	for _, acct := range cl.EthAccounts() {
		meta := nonces[acct.Name]
		mintData, err := bench_erc20.PackMint(meta.Account.Address, mintAmount)
		if err != nil {
			t.Fatalf("pack mint: %v", err)
		}
		signed, err := cl.SignEthContractCall(meta.Account, contractAddr, mintData, meta.Nonce, gasLimit)
		if err != nil {
			t.Fatalf("sign mint tx: %v", err)
		}
		_, err = cl.EthSendRawTransaction(ctx, 0, signed.RawHex)
		if err != nil {
			t.Logf("[mint-erc20] send failed for %s: %v", acct.Name, err)
		}
	}
}

// WarmupEth sends a small number of ETH transfers to warm up the JSON-RPC path.
func WarmupEth(ctx context.Context, t *testing.T, cl *cluster.Cluster, nonces map[string]EthNonceMeta, toAddr common.Address, gasLimit uint64) {
	t.Helper()
	ethAccounts := cl.EthAccounts()
	oneGas := big.NewInt(1)

	for i := 0; i < 5 && i < len(ethAccounts); i++ {
		acct := ethAccounts[i]
		meta := nonces[acct.Name]
		signed, err := cl.SignEthTransfer(meta.Account, toAddr, oneGas, meta.Nonce, gasLimit)
		if err != nil {
			t.Logf("[warmup-eth] sign failed: %v", err)
			continue
		}
		viaNode := i % cl.NodeCount()
		_, err = cl.EthSendRawTransaction(ctx, viaNode, signed.RawHex)
		if err != nil {
			t.Logf("[warmup-eth] send failed: %v", err)
		}
	}
}
