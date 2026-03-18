package cluster

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	MaxNodeCount    = 10
	defaultBasePort = 26000
	defaultStride   = 20
)

type ClusterOptions struct {
	NodeCount      int
	AccountCount   int
	ChainID        string
	BasePort       int
	PortStride     int
	BinaryPath     string
	MemIAVL        bool
	ValidatorCount int   // 0 or 1 = single validator (current behavior)
	MaxBlockGas    int64 // 0 = no limit; >0 sets block max_gas in genesis
	NoAllowQueued  bool  // true = omit --allow-queued flag (for pre-proxy baseline binaries)
}

type Node struct {
	Index   int
	Name    string
	Home    string
	Ports   NodePorts
	PeerID  string
	LogPath string

	cmd     *exec.Cmd
	logFile *os.File
}

type AccountMeta struct {
	Address       string
	AccountNumber uint64
	Sequence      uint64
}

type TxResult struct {
	Code   int64
	TxHash string
	RawLog string
	Err    error
}

type Cluster struct {
	t     *testing.T
	opts  ClusterOptions
	bin   string
	repo  string
	root  string
	nodes []*Node

	valAddresses []string
	accounts     map[string]string
	ethAccounts  []EthAccount

	mu     sync.Mutex
	closed bool
}

func NewCluster(ctx context.Context, t *testing.T, opts ClusterOptions) (*Cluster, error) {
	t.Helper()

	if opts.NodeCount < 1 || opts.NodeCount > MaxNodeCount {
		return nil, fmt.Errorf("node count must be 1..%d, got %d", MaxNodeCount, opts.NodeCount)
	}
	if opts.AccountCount < 1 {
		opts.AccountCount = 3
	}
	if opts.ChainID == "" {
		opts.ChainID = "testnet"
	}
	if opts.BasePort == 0 {
		opts.BasePort = defaultBasePort
	}
	if opts.PortStride == 0 {
		opts.PortStride = defaultStride
	}

	repoRoot, err := findRepoRoot()
	if err != nil {
		return nil, err
	}

	binPath := opts.BinaryPath
	if binPath == "" {
		if envBin := os.Getenv("E2E_MINITIAD_BIN"); envBin != "" {
			binPath = envBin
		} else {
			binPath = filepath.Join(t.TempDir(), "minitiad")
			if err := buildMinitiad(ctx, repoRoot, binPath); err != nil {
				return nil, err
			}
		}
	}

	c := &Cluster{
		t:        t,
		opts:     opts,
		bin:      binPath,
		repo:     repoRoot,
		root:     t.TempDir(),
		nodes:    make([]*Node, 0, opts.NodeCount),
		accounts: map[string]string{},
	}

	if err := c.initNodes(ctx); err != nil {
		return nil, err
	}
	if err := c.configureNodes(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Cluster) Start(ctx context.Context) error {
	for _, n := range c.nodes {
		if err := c.startNode(ctx, n); err != nil {
			c.Close()
			return err
		}
	}

	return nil
}

func (c *Cluster) Logf(format string, args ...any) {
	if c.t != nil {
		c.t.Logf(format, args...)
	}
}

func (c *Cluster) Close() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.closed {
		return
	}
	c.closed = true

	for _, n := range c.nodes {
		if n.cmd == nil || n.cmd.Process == nil {
			if n.logFile != nil {
				_ = n.logFile.Close()
			}
			continue
		}

		// Kill the entire process group to ensure child processes are also terminated.
		_ = syscall.Kill(-n.cmd.Process.Pid, syscall.SIGTERM)
	}

	deadline := time.Now().Add(10 * time.Second)
	for _, n := range c.nodes {
		if n.cmd == nil {
			if n.logFile != nil {
				_ = n.logFile.Close()
			}
			continue
		}

		done := make(chan error, 1)
		go func(cmd *exec.Cmd) {
			done <- cmd.Wait()
		}(n.cmd)

		select {
		case <-done:
		case <-time.After(time.Until(deadline)):
			if n.cmd.Process != nil {
				_ = syscall.Kill(-n.cmd.Process.Pid, syscall.SIGKILL)
				<-done
			}
		}

		if n.logFile != nil {
			_ = n.logFile.Close()
		}
	}
}

func (c *Cluster) WaitForReady(ctx context.Context, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("wait for network ready timed out: %w", ctx.Err())
		default:
		}

		allHealthy := true
		for _, n := range c.nodes {
			h, _, err := c.nodeStatus(ctx, n)
			if err != nil || !h {
				allHealthy = false
				break
			}
		}
		if !allHealthy {
			time.Sleep(800 * time.Millisecond)
			continue
		}

		h1, err := c.latestHeight(ctx, 0)
		if err != nil {
			time.Sleep(800 * time.Millisecond)
			continue
		}
		time.Sleep(2 * time.Second)
		h2, err := c.latestHeight(ctx, 0)
		if err != nil {
			time.Sleep(800 * time.Millisecond)
			continue
		}
		if h2 > h1 && h2 > 1 {
			return nil
		}
	}
}

func (c *Cluster) WaitForMempoolEmpty(ctx context.Context, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("wait for mempool empty timed out: %w", ctx.Err())
		default:
		}

		allEmpty := true
		for i := range c.nodes {
			n, err := c.unconfirmedTxCount(ctx, i)
			if err != nil || n != 0 {
				allEmpty = false
				break
			}
		}
		if allEmpty {
			return nil
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func (c *Cluster) NodeCount() int {
	return len(c.nodes)
}

// NodeRPCPort returns the RPC port for the given node index.
func (c *Cluster) NodeRPCPort(index int) (int, error) {
	n, err := c.getNode(index)
	if err != nil {
		return 0, err
	}

	return n.Ports.RPC, nil
}

// NodeJSONRPCPort returns the JSON-RPC port for the given node index.
func (c *Cluster) NodeJSONRPCPort(index int) (int, error) {
	n, err := c.getNode(index)
	if err != nil {
		return 0, err
	}

	return n.Ports.JSONRPC, nil
}

// LatestHeight returns the latest block height from the given node.
func (c *Cluster) LatestHeight(ctx context.Context, nodeIndex int) (int64, error) {
	return c.latestHeight(ctx, nodeIndex)
}

// UnconfirmedTxCount returns the number of unconfirmed transactions in the given node's mempool.
func (c *Cluster) UnconfirmedTxCount(ctx context.Context, nodeIndex int) (int64, error) {
	return c.unconfirmedTxCount(ctx, nodeIndex)
}

// BlockResult holds the data extracted from a block query.
type BlockResult struct {
	TxHashes  []string
	BlockTime time.Time
}

// QueryBlock queries a specific block by height from the given node and returns tx hashes and block time.
func (c *Cluster) QueryBlock(ctx context.Context, nodeIndex int, height int64) (BlockResult, error) {
	n, err := c.getNode(nodeIndex)
	if err != nil {
		return BlockResult{}, err
	}

	url := fmt.Sprintf("http://127.0.0.1:%d/block?height=%d", n.Ports.RPC, height)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return BlockResult{}, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return BlockResult{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return BlockResult{}, fmt.Errorf("block query status code %d", resp.StatusCode)
	}

	var decoded struct {
		Result struct {
			Block struct {
				Header struct {
					Time string `json:"time"`
				} `json:"header"`
				Data struct {
					Txs []string `json:"txs"`
				} `json:"data"`
			} `json:"block"`
		} `json:"result"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&decoded); err != nil {
		return BlockResult{}, fmt.Errorf("failed to decode block response: %w", err)
	}

	blockTime, err := time.Parse(time.RFC3339Nano, decoded.Result.Block.Header.Time)
	if err != nil {
		return BlockResult{}, fmt.Errorf("failed to parse block time %q: %w", decoded.Result.Block.Header.Time, err)
	}

	txHashes := make([]string, 0, len(decoded.Result.Block.Data.Txs))
	for idx, txBase64 := range decoded.Result.Block.Data.Txs {
		txBytes, decErr := base64Decode(txBase64)
		if decErr != nil {
			return BlockResult{}, fmt.Errorf("failed to decode tx at height=%d index=%d: %w", height, idx, decErr)
		}
		hash := sha256Hash(txBytes)
		txHashes = append(txHashes, strings.ToUpper(hash))
	}

	return BlockResult{
		TxHashes:  txHashes,
		BlockTime: blockTime,
	}, nil
}

func (c *Cluster) AccountNames() []string {
	names := make([]string, 0, len(c.accounts))
	for i := 1; i <= c.opts.AccountCount; i++ {
		names = append(names, fmt.Sprintf("acc%d", i))
	}

	return names
}

func (c *Cluster) ValidatorAddress() string {
	if len(c.valAddresses) == 0 {
		return ""
	}

	return c.valAddresses[0]
}

func (c *Cluster) ValidatorAddresses() []string {
	return c.valAddresses
}

func (c *Cluster) AccountAddress(name string) (string, error) {
	addr, ok := c.accounts[name]
	if !ok {
		return "", fmt.Errorf("unknown account: %s", name)
	}

	return addr, nil
}

func (c *Cluster) RepoPath(parts ...string) string {
	all := make([]string, 0, len(parts)+1)
	all = append(all, c.repo)
	all = append(all, parts...)

	return filepath.Join(all...)
}

func (c *Cluster) QueryAccountMeta(ctx context.Context, nodeIndex int, address string) (AccountMeta, error) {
	node, err := c.getNode(nodeIndex)
	if err != nil {
		return AccountMeta{}, err
	}

	url := fmt.Sprintf("http://127.0.0.1:%d/cosmos/auth/v1beta1/accounts/%s", node.Ports.API, address)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return AccountMeta{}, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return AccountMeta{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return AccountMeta{}, fmt.Errorf("account query status code %d", resp.StatusCode)
	}

	var decoded map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&decoded); err != nil {
		return AccountMeta{}, fmt.Errorf("failed to parse account query output: %w", err)
	}

	accountAny, ok := decoded["account"]
	if !ok {
		return AccountMeta{}, errors.New("missing account field")
	}

	accountNumber, ok := findUintField(accountAny, "account_number")
	if !ok {
		accountNumber, ok = findUintField(accountAny, "accountNumber")
	}
	if !ok {
		return AccountMeta{}, errors.New("account_number not found")
	}
	sequence, ok := findUintField(accountAny, "sequence")
	if !ok {
		return AccountMeta{}, errors.New("sequence not found")
	}

	return AccountMeta{
		Address:       address,
		AccountNumber: accountNumber,
		Sequence:      sequence,
	}, nil
}

// ---------------------------------------------------------------------------
// Bank send
// ---------------------------------------------------------------------------

func (c *Cluster) SendBankTxWithSequence(ctx context.Context, fromName, toAddress, amount string, accountNumber, sequence, gasLimit uint64, viaNode int) TxResult {
	node, err := c.getNode(viaNode)
	if err != nil {
		return TxResult{Err: err}
	}
	c.t.Logf(
		"[send] from=%s to=%s amount=%s account_number=%d sequence=%d via_node=%d rpc_port=%d",
		fromName, toAddress, amount, accountNumber, sequence, viaNode, node.Ports.RPC,
	)

	args := []string{
		"tx", "bank", "send", fromName, toAddress, amount,
		"--chain-id", c.opts.ChainID,
		"--node", fmt.Sprintf("http://127.0.0.1:%d", node.Ports.RPC),
		"--home", c.nodes[0].Home,
		"--keyring-backend", "test",
		"--gas-prices", "0GAS",
		"--gas", strconv.FormatUint(gasLimit, 10),
		"--offline",
	}
	if !c.opts.NoAllowQueued {
		args = append(args, "--allow-queued")
	}
	args = append(args,
		"--broadcast-mode", "sync",
		"--account-number", strconv.FormatUint(accountNumber, 10),
		"--sequence", strconv.FormatUint(sequence, 10),
		"--yes",
		"--output", "json",
	)

	out, err := c.exec(ctx, args...)
	if err != nil {
		c.t.Logf(
			"[send] failed from=%s sequence=%d err=%v",
			fromName, sequence, err,
		)
		return TxResult{Err: err}
	}

	res, err := parseTxResultFromOutput(out)
	if err != nil {
		c.t.Logf("[send] parse-failed from=%s sequence=%d output=%s", fromName, sequence, strings.TrimSpace(string(out)))
		return TxResult{Err: err}
	}
	c.t.Logf(
		"[send] result from=%s sequence=%d code=%d txhash=%s raw_log=%q",
		fromName, sequence, res.Code, res.TxHash, res.RawLog,
	)

	return res
}

// ---------------------------------------------------------------------------
// EVM methods
// ---------------------------------------------------------------------------

// DeployContract deploys an EVM contract from a hex bytecode string.
// It writes the bytecode to a temp file and calls `tx evm create <file>`.
// Returns TxResult with the transaction hash.
func (c *Cluster) DeployContract(ctx context.Context, fromName, bytecodeHex string, viaNode int) TxResult {
	node, err := c.getNode(viaNode)
	if err != nil {
		return TxResult{Err: err}
	}

	// Write bytecode to temp file (tx evm create expects a bin file path)
	tmpFile, err := os.CreateTemp("", "contract-*.bin")
	if err != nil {
		return TxResult{Err: err}
	}
	defer os.Remove(tmpFile.Name())

	// Strip 0x prefix if present
	hexStr := strings.TrimPrefix(bytecodeHex, "0x")
	if _, err := tmpFile.WriteString(hexStr); err != nil {
		tmpFile.Close()
		return TxResult{Err: err}
	}
	tmpFile.Close()

	args := []string{
		"tx", "evm", "create", tmpFile.Name(),
		"--from", fromName,
		"--chain-id", c.opts.ChainID,
		"--node", fmt.Sprintf("http://127.0.0.1:%d", node.Ports.RPC),
		"--home", c.nodes[0].Home,
		"--keyring-backend", "test",
		"--gas-prices", "0GAS",
		"--gas", "auto",
		"--gas-adjustment", "1.5",
	}
	if !c.opts.NoAllowQueued {
		args = append(args, "--allow-queued")
	}
	args = append(args,
		"--broadcast-mode", "sync",
		"--yes",
		"--output", "json",
	)

	out, err := c.exec(ctx, args...)
	if err != nil {
		return TxResult{Err: err}
	}
	res, err := parseTxResultFromOutput(out)
	if err != nil {
		return TxResult{Err: err}
	}
	c.t.Logf("[evm-deploy] from=%s code=%d txhash=%s", fromName, res.Code, res.TxHash)

	return res
}

// CallContract calls an EVM contract with the given input hex.
func (c *Cluster) CallContract(ctx context.Context, fromName, contractAddr, inputHex string, accountNumber, sequence, gasLimit uint64, viaNode int) TxResult {
	node, err := c.getNode(viaNode)
	if err != nil {
		return TxResult{Err: err}
	}

	args := []string{
		"tx", "evm", "call", contractAddr, inputHex,
		"--from", fromName,
		"--chain-id", c.opts.ChainID,
		"--node", fmt.Sprintf("http://127.0.0.1:%d", node.Ports.RPC),
		"--home", c.nodes[0].Home,
		"--keyring-backend", "test",
		"--gas-prices", "0GAS",
		"--gas", strconv.FormatUint(gasLimit, 10),
		"--offline",
	}
	if !c.opts.NoAllowQueued {
		args = append(args, "--allow-queued")
	}
	args = append(args,
		"--account-number", strconv.FormatUint(accountNumber, 10),
		"--sequence", strconv.FormatUint(sequence, 10),
		"--broadcast-mode", "sync",
		"--yes",
		"--output", "json",
	)

	out, err := c.exec(ctx, args...)
	if err != nil {
		return TxResult{Err: err}
	}
	res, err := parseTxResultFromOutput(out)
	if err != nil {
		return TxResult{Err: err}
	}
	c.t.Logf("[evm-call] from=%s seq=%d gas=%d contract=%s code=%d txhash=%s",
		fromName, sequence, gasLimit, contractAddr, res.Code, res.TxHash)

	return res
}

// EstimateEvmGas estimates gas for an EVM call using --gas auto --generate-only.
func (c *Cluster) EstimateEvmGas(ctx context.Context, fromName, contractAddr, inputHex string, viaNode int) (uint64, error) {
	node, err := c.getNode(viaNode)
	if err != nil {
		return 0, err
	}

	out, err := c.exec(ctx,
		"tx", "evm", "call", contractAddr, inputHex,
		"--from", fromName,
		"--chain-id", c.opts.ChainID,
		"--node", fmt.Sprintf("http://127.0.0.1:%d", node.Ports.RPC),
		"--home", c.nodes[0].Home,
		"--keyring-backend", "test",
		"--gas-prices", "0GAS",
		"--gas", "auto",
		"--gas-adjustment", "1.2",
		"--generate-only",
		"--yes",
		"--output", "json",
	)
	if err != nil {
		return 0, err
	}
	gas, err := parseEstimatedGas(out)
	if err != nil {
		return 0, err
	}
	gas += 200_000 // buffer for fee payment
	c.t.Logf("[evm-estimate] from=%s contract=%s gas=%d", fromName, contractAddr, gas)

	return gas, nil
}

// GenerateSignedEvmCallTx generates a signed, base64-encoded EVM call transaction offline.
func (c *Cluster) GenerateSignedEvmCallTx(
	ctx context.Context,
	fromName, contractAddr, inputHex string,
	accountNumber, sequence, gasLimit uint64,
) (SignedTx, error) {
	genArgs := []string{
		"tx", "evm", "call", contractAddr, inputHex,
		"--from", fromName,
		"--chain-id", c.opts.ChainID,
		"--home", c.nodes[0].Home,
		"--keyring-backend", "test",
		"--gas-prices", "0GAS",
		"--gas", strconv.FormatUint(gasLimit, 10),
		"--account-number", strconv.FormatUint(accountNumber, 10),
		"--sequence", strconv.FormatUint(sequence, 10),
		"--generate-only",
	}
	if !c.opts.NoAllowQueued {
		genArgs = append(genArgs, "--allow-queued")
	}

	return c.signAndEncode(ctx, fromName, accountNumber, sequence, genArgs)
}

// QueryTxResult queries a transaction by hash via the REST API and returns the parsed response.
func (c *Cluster) QueryTxResult(ctx context.Context, nodeIndex int, txHash string) (map[string]any, error) {
	node, err := c.getNode(nodeIndex)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("http://127.0.0.1:%d/cosmos/tx/v1beta1/txs/%s", node.Ports.API, txHash)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("tx query status code %d for hash %s", resp.StatusCode, txHash)
	}

	var decoded map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&decoded); err != nil {
		return nil, fmt.Errorf("failed to parse tx query output: %w", err)
	}

	return decoded, nil
}

// ExtractContractAddress extracts the contract address from a deploy tx result's events.
func ExtractContractAddress(txResult map[string]any) (string, error) {
	txResp, ok := txResult["tx_response"].(map[string]any)
	if !ok {
		txResp = txResult
	}

	events, _ := txResp["events"].([]any)
	for _, ev := range events {
		event, ok := ev.(map[string]any)
		if !ok {
			continue
		}
		attrs, _ := event["attributes"].([]any)
		for _, a := range attrs {
			attr, ok := a.(map[string]any)
			if !ok {
				continue
			}
			key, _ := attr["key"].(string)
			if key == "contract_address" || key == "contract-address" || key == "contract" {
				val, _ := attr["value"].(string)
				if val != "" {
					return val, nil
				}
			}
		}
	}

	logs, _ := txResp["logs"].([]any)
	for _, l := range logs {
		logEntry, ok := l.(map[string]any)
		if !ok {
			continue
		}
		logEvents, _ := logEntry["events"].([]any)
		for _, ev := range logEvents {
			event, ok := ev.(map[string]any)
			if !ok {
				continue
			}
			attrs, _ := event["attributes"].([]any)
			for _, a := range attrs {
				attr, ok := a.(map[string]any)
				if !ok {
					continue
				}
				key, _ := attr["key"].(string)
				if key == "contract_address" || key == "contract-address" || key == "contract" {
					val, _ := attr["value"].(string)
					if val != "" {
						return val, nil
					}
				}
			}
		}
	}

	return "", errors.New("contract_address not found in tx events")
}

// ---------------------------------------------------------------------------
// JSON-RPC methods
// ---------------------------------------------------------------------------

// EthSendRawTransaction sends a signed Ethereum transaction via JSON-RPC.
func (c *Cluster) EthSendRawTransaction(ctx context.Context, nodeIndex int, signedTxHex string) (string, error) {
	node, err := c.getNode(nodeIndex)
	if err != nil {
		return "", err
	}

	if !strings.HasPrefix(signedTxHex, "0x") {
		signedTxHex = "0x" + signedTxHex
	}

	reqBody := fmt.Sprintf(`{"jsonrpc":"2.0","id":1,"method":"eth_sendRawTransaction","params":["%s"]}`, signedTxHex)

	return c.ethJSONRPCCall(ctx, node.Ports.JSONRPC, reqBody, "result")
}

// EthGetTransactionCount returns the nonce for an address via JSON-RPC.
func (c *Cluster) EthGetTransactionCount(ctx context.Context, nodeIndex int, address string) (uint64, error) {
	node, err := c.getNode(nodeIndex)
	if err != nil {
		return 0, err
	}

	reqBody := fmt.Sprintf(`{"jsonrpc":"2.0","id":1,"method":"eth_getTransactionCount","params":["%s","latest"]}`, address)
	result, err := c.ethJSONRPCCall(ctx, node.Ports.JSONRPC, reqBody, "result")
	if err != nil {
		return 0, err
	}

	return strconv.ParseUint(strings.TrimPrefix(result, "0x"), 16, 64)
}

// EthEstimateGas estimates gas for a transaction via JSON-RPC.
func (c *Cluster) EthEstimateGas(ctx context.Context, nodeIndex int, from, to, inputHex string) (uint64, error) {
	node, err := c.getNode(nodeIndex)
	if err != nil {
		return 0, err
	}

	txObj := fmt.Sprintf(`{"from":"%s","to":"%s","data":"%s"}`, from, to, inputHex)
	reqBody := fmt.Sprintf(`{"jsonrpc":"2.0","id":1,"method":"eth_estimateGas","params":[%s]}`, txObj)
	result, err := c.ethJSONRPCCall(ctx, node.Ports.JSONRPC, reqBody, "result")
	if err != nil {
		return 0, err
	}

	return strconv.ParseUint(strings.TrimPrefix(result, "0x"), 16, 64)
}

// EthGetTransactionReceipt returns the receipt for a transaction via JSON-RPC.
func (c *Cluster) EthGetTransactionReceipt(ctx context.Context, nodeIndex int, txHash string) (map[string]any, error) {
	node, err := c.getNode(nodeIndex)
	if err != nil {
		return nil, err
	}

	reqBody := fmt.Sprintf(`{"jsonrpc":"2.0","id":1,"method":"eth_getTransactionReceipt","params":["%s"]}`, txHash)
	url := fmt.Sprintf("http://127.0.0.1:%d", node.Ports.JSONRPC)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, strings.NewReader(reqBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var rpcResp struct {
		Result map[string]any `json:"result"`
		Error  *struct {
			Message string `json:"message"`
		} `json:"error"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&rpcResp); err != nil {
		return nil, err
	}
	if rpcResp.Error != nil {
		return nil, fmt.Errorf("eth_getTransactionReceipt error: %s", rpcResp.Error.Message)
	}

	return rpcResp.Result, nil
}

func (c *Cluster) ethJSONRPCCall(ctx context.Context, port int, reqBody, resultKey string) (string, error) {
	url := fmt.Sprintf("http://127.0.0.1:%d", port)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, strings.NewReader(reqBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var rpcResp map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&rpcResp); err != nil {
		return "", err
	}

	if errObj, ok := rpcResp["error"]; ok && errObj != nil {
		if errMap, ok := errObj.(map[string]any); ok {
			msg, _ := errMap["message"].(string)
			return "", fmt.Errorf("JSON-RPC error: %s", msg)
		}
	}
	result, _ := rpcResp[resultKey].(string)

	return result, nil
}

// ---------------------------------------------------------------------------
// Pre-signed transaction support
// ---------------------------------------------------------------------------

// SignedTx holds a pre-signed, encoded transaction ready for HTTP broadcast.
type SignedTx struct {
	Account       string
	Sequence      uint64
	TxBase64      string
	TxHash        string
	AccountNumber uint64
}

// GenerateSignedBankTx generates a signed, base64-encoded bank send transaction offline.
func (c *Cluster) GenerateSignedBankTx(ctx context.Context, fromName, toAddress, amount string, accountNumber, sequence, gasLimit uint64) (SignedTx, error) {
	genArgs := []string{
		"tx", "bank", "send", fromName, toAddress, amount,
		"--chain-id", c.opts.ChainID,
		"--home", c.nodes[0].Home,
		"--keyring-backend", "test",
		"--gas-prices", "0GAS",
		"--gas", strconv.FormatUint(gasLimit, 10),
		"--account-number", strconv.FormatUint(accountNumber, 10),
		"--sequence", strconv.FormatUint(sequence, 10),
		"--generate-only",
	}
	if !c.opts.NoAllowQueued {
		genArgs = append(genArgs, "--allow-queued")
	}

	return c.signAndEncode(ctx, fromName, accountNumber, sequence, genArgs)
}

// signAndEncode is a shared helper: generate-only → sign → encode → SignedTx.
func (c *Cluster) signAndEncode(ctx context.Context, fromName string, accountNumber, sequence uint64, genArgs []string) (SignedTx, error) {
	unsignedJSON, err := c.exec(ctx, genArgs...)
	if err != nil {
		return SignedTx{}, fmt.Errorf("generate-only: %w", err)
	}

	tmpFile, err := os.CreateTemp("", "unsigned-tx-*.json")
	if err != nil {
		return SignedTx{}, err
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.Write(unsignedJSON); err != nil {
		tmpFile.Close()
		return SignedTx{}, err
	}
	tmpFile.Close()

	signArgs := []string{
		"tx", "sign", tmpFile.Name(),
		"--from", fromName,
		"--chain-id", c.opts.ChainID,
		"--home", c.nodes[0].Home,
		"--keyring-backend", "test",
		"--offline",
		"--account-number", strconv.FormatUint(accountNumber, 10),
		"--sequence", strconv.FormatUint(sequence, 10),
	}

	signedJSON, err := c.exec(ctx, signArgs...)
	if err != nil {
		return SignedTx{}, fmt.Errorf("sign: %w", err)
	}

	signedFile, err := os.CreateTemp("", "signed-tx-*.json")
	if err != nil {
		return SignedTx{}, err
	}
	defer os.Remove(signedFile.Name())

	if _, err := signedFile.Write(signedJSON); err != nil {
		signedFile.Close()
		return SignedTx{}, err
	}
	signedFile.Close()

	encodeArgs := []string{
		"tx", "encode", signedFile.Name(),
		"--home", c.nodes[0].Home,
	}

	encodedOut, err := c.exec(ctx, encodeArgs...)
	if err != nil {
		return SignedTx{}, fmt.Errorf("encode: %w", err)
	}

	txBase64 := strings.TrimSpace(string(encodedOut))
	txBytes, err := base64.StdEncoding.DecodeString(txBase64)
	if err != nil {
		return SignedTx{}, fmt.Errorf("decode base64 for hash: %w", err)
	}
	hash := sha256.Sum256(txBytes)

	return SignedTx{
		Account:       fromName,
		Sequence:      sequence,
		TxBase64:      txBase64,
		TxHash:        strings.ToUpper(hex.EncodeToString(hash[:])),
		AccountNumber: accountNumber,
	}, nil
}

// BroadcastTxSync broadcasts a pre-signed transaction via HTTP to the given node /broadcast_tx_sync endpoint.
func (c *Cluster) BroadcastTxSync(ctx context.Context, nodeIndex int, txBase64 string) (TxResult, error) {
	node, err := c.getNode(nodeIndex)
	if err != nil {
		return TxResult{}, err
	}

	txBytes, err := base64.StdEncoding.DecodeString(txBase64)
	if err != nil {
		return TxResult{}, fmt.Errorf("decode tx base64: %w", err)
	}

	reqBody := fmt.Sprintf(`{"jsonrpc":"2.0","id":1,"method":"broadcast_tx_sync","params":{"tx":"%s"}}`, base64.StdEncoding.EncodeToString(txBytes))
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("http://127.0.0.1:%d", node.Ports.RPC), strings.NewReader(reqBody))
	if err != nil {
		return TxResult{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return TxResult{}, err
	}
	defer resp.Body.Close()

	var rpcResp struct {
		Result struct {
			Code int64  `json:"code"`
			Hash string `json:"hash"`
			Log  string `json:"log"`
		} `json:"result"`
		Error *struct {
			Message string `json:"message"`
		} `json:"error"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&rpcResp); err != nil {
		return TxResult{}, fmt.Errorf("decode broadcast response: %w", err)
	}

	if rpcResp.Error != nil {
		return TxResult{Err: fmt.Errorf("broadcast RPC error: %s", rpcResp.Error.Message)}, nil
	}

	return TxResult{
		Code:   rpcResp.Result.Code,
		TxHash: rpcResp.Result.Hash,
		RawLog: rpcResp.Result.Log,
	}, nil
}

// ---------------------------------------------------------------------------
// Node init / config
// ---------------------------------------------------------------------------

func (c *Cluster) initNodes(ctx context.Context) error {
	for i := 0; i < c.opts.NodeCount; i++ {
		ports, err := allocatePorts(i, c.opts.BasePort, c.opts.PortStride)
		if err != nil {
			return err
		}
		n := &Node{
			Index: i,
			Name:  fmt.Sprintf("node%d", i),
			Home:  filepath.Join(c.root, fmt.Sprintf("node%d", i)),
			Ports: ports,
		}
		if _, err := c.exec(ctx, "init", n.Name, "--home", n.Home, "--chain-id", c.opts.ChainID); err != nil {
			return err
		}
		c.nodes = append(c.nodes, n)
	}

	// resolve node ids early, needed for gentx memos in multi-validator setup.
	for _, n := range c.nodes {
		out, err := c.exec(ctx, "comet", "show-node-id", "--home", n.Home)
		if err != nil {
			return err
		}
		n.PeerID = strings.TrimSpace(string(out))
	}

	baseHome := c.nodes[0].Home

	valCount := 1
	if c.opts.ValidatorCount > 1 {
		valCount = c.opts.ValidatorCount
		if valCount > c.opts.NodeCount {
			valCount = c.opts.NodeCount
		}
	}

	if valCount <= 1 {
		return c.initSingleValidator(ctx, baseHome)
	}

	return c.initMultiValidator(ctx, baseHome, valCount)
}

func (c *Cluster) initSingleValidator(ctx context.Context, baseHome string) error {
	if _, err := c.exec(ctx, "keys", "add", "val", "--keyring-backend", "test", "--home", baseHome); err != nil {
		return err
	}
	valAddr, err := c.keyAddress(ctx, "val")
	if err != nil {
		return err
	}
	c.valAddresses = []string{valAddr}

	if err := c.addAccountKeys(ctx, baseHome); err != nil {
		return err
	}

	if err := c.generateEthAccounts(); err != nil {
		return err
	}

	if _, err := c.exec(ctx,
		"genesis", "add-genesis-account", "val", "1000000000000000GAS",
		"--home", baseHome, "--keyring-backend", "test",
	); err != nil {
		return err
	}

	if err := c.addAccountGenesisAccounts(ctx, baseHome); err != nil {
		return err
	}

	if err := c.addEthGenesisAccounts(ctx, baseHome); err != nil {
		return err
	}

	genesisPath := filepath.Join(baseHome, "config", "genesis.json")
	valPubKeyJSON, err := c.exec(ctx, "comet", "show-validator", "--home", baseHome)
	if err != nil {
		return fmt.Errorf("show-validator: %w", err)
	}

	if err := patchGenesisOpchild(genesisPath, c.valAddresses[0], []validatorInfo{
		{moniker: "node0", pubKeyJSON: strings.TrimSpace(string(valPubKeyJSON))},
	}); err != nil {
		return fmt.Errorf("patch genesis opchild: %w", err)
	}

	return c.distributeGenesis(baseHome)
}

func (c *Cluster) initMultiValidator(ctx context.Context, baseHome string, valCount int) error {
	for i := 0; i < valCount; i++ {
		name := fmt.Sprintf("val%d", i)
		if _, err := c.exec(ctx, "keys", "add", name, "--keyring-backend", "test", "--home", baseHome); err != nil {
			return err
		}
		addr, err := c.keyAddress(ctx, name)
		if err != nil {
			return err
		}
		c.valAddresses = append(c.valAddresses, addr)
	}

	if err := c.addAccountKeys(ctx, baseHome); err != nil {
		return err
	}

	if err := c.generateEthAccounts(); err != nil {
		return err
	}

	for i := 0; i < valCount; i++ {
		name := fmt.Sprintf("val%d", i)
		if _, err := c.exec(ctx,
			"genesis", "add-genesis-account", name, "1000000000000000GAS",
			"--home", baseHome, "--keyring-backend", "test",
		); err != nil {
			return err
		}
	}

	if err := c.addAccountGenesisAccounts(ctx, baseHome); err != nil {
		return err
	}

	if err := c.addEthGenesisAccounts(ctx, baseHome); err != nil {
		return err
	}

	var vals []validatorInfo
	for i := 0; i < valCount; i++ {
		nodeHome := c.nodes[i].Home
		out, err := c.exec(ctx, "comet", "show-validator", "--home", nodeHome)
		if err != nil {
			return fmt.Errorf("show-validator node%d: %w", i, err)
		}
		vals = append(vals, validatorInfo{
			moniker:    fmt.Sprintf("node%d", i),
			pubKeyJSON: strings.TrimSpace(string(out)),
		})
	}

	baseGenesisPath := filepath.Join(baseHome, "config", "genesis.json")
	if err := patchGenesisOpchild(baseGenesisPath, c.valAddresses[0], vals); err != nil {
		return fmt.Errorf("patch genesis opchild: %w", err)
	}

	if err := c.distributeGenesis(baseHome); err != nil {
		return err
	}

	baseKeyringDir := filepath.Join(baseHome, "keyring-test")
	for i := 1; i < len(c.nodes); i++ {
		nodeKeyringDir := filepath.Join(c.nodes[i].Home, "keyring-test")
		if err := copyDir(baseKeyringDir, nodeKeyringDir); err != nil {
			return fmt.Errorf("distribute keyring to node%d: %w", i, err)
		}
	}

	return nil
}

func (c *Cluster) addAccountKeys(ctx context.Context, baseHome string) error {
	for i := 1; i <= c.opts.AccountCount; i++ {
		name := fmt.Sprintf("acc%d", i)
		if _, err := c.exec(ctx, "keys", "add", name, "--keyring-backend", "test", "--home", baseHome); err != nil {
			return err
		}
		addr, err := c.keyAddress(ctx, name)
		if err != nil {
			return err
		}
		c.accounts[name] = addr
	}

	return nil
}

func (c *Cluster) addAccountGenesisAccounts(ctx context.Context, baseHome string) error {
	for i := 1; i <= c.opts.AccountCount; i++ {
		name := fmt.Sprintf("acc%d", i)
		if _, err := c.exec(ctx,
			"genesis", "add-genesis-account", name, "1000000000000000GAS",
			"--home", baseHome, "--keyring-backend", "test",
		); err != nil {
			return err
		}
	}
	return nil
}

type validatorInfo struct {
	moniker    string
	pubKeyJSON string
}

// patchGenesisOpchild configures the opchild module in genesis:
//   - sets admin address
//   - adds validators to opchild.validators
func patchGenesisOpchild(genesisPath, adminAddr string, validators []validatorInfo) error {
	bz, err := os.ReadFile(genesisPath)
	if err != nil {
		return err
	}
	var genesis map[string]interface{}
	if err := json.Unmarshal(bz, &genesis); err != nil {
		return fmt.Errorf("unmarshal genesis: %w", err)
	}

	appState, ok := genesis["app_state"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("genesis missing app_state")
	}
	opchild, ok := appState["opchild"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("genesis missing opchild module")
	}
	params, ok := opchild["params"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("genesis missing opchild.params")
	}

	// set admin
	params["admin"] = adminAddr

	// clear bridge_executors
	if executors, ok := params["bridge_executors"].([]interface{}); ok {
		cleaned := make([]interface{}, 0, len(executors))
		for _, e := range executors {
			if s, ok := e.(string); ok && s == "" {
				continue
			}
			cleaned = append(cleaned, e)
		}
		params["bridge_executors"] = cleaned
	}

	// add validators to opchild.validators[]
	valList := make([]interface{}, 0, len(validators))
	for _, v := range validators {
		var pubKey map[string]interface{}
		if err := json.Unmarshal([]byte(v.pubKeyJSON), &pubKey); err != nil {
			return fmt.Errorf("parse validator pubkey: %w", err)
		}

		keyB64, _ := pubKey["key"].(string)
		keyBytes, err := base64.StdEncoding.DecodeString(keyB64)
		if err != nil {
			return fmt.Errorf("decode pubkey base64: %w", err)
		}

		addrHash := sha256.Sum256(keyBytes)
		valOperAddr, err := sdk.Bech32ifyAddressBytes("initvaloper", addrHash[:20])
		if err != nil {
			return fmt.Errorf("bech32 encode valoper: %w", err)
		}

		valEntry := map[string]interface{}{
			"moniker":          v.moniker,
			"operator_address": valOperAddr,
			"consensus_pubkey": pubKey,
			"cons_power":       "1",
		}
		valList = append(valList, valEntry)
	}
	opchild["validators"] = valList

	out, err := json.MarshalIndent(genesis, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal genesis: %w", err)
	}

	return os.WriteFile(genesisPath, out, 0o600)
}

func patchGenesisBlockGas(genesisPath string, maxGas int64) error {
	bz, err := os.ReadFile(genesisPath)
	if err != nil {
		return err
	}
	var genesis map[string]interface{}
	if err := json.Unmarshal(bz, &genesis); err != nil {
		return fmt.Errorf("unmarshal genesis: %w", err)
	}

	gasStr := strconv.FormatInt(maxGas, 10)

	patched := false
	if cp, ok := genesis["consensus"].(map[string]interface{}); ok {
		if params, ok := cp["params"].(map[string]interface{}); ok {
			if block, ok := params["block"].(map[string]interface{}); ok {
				block["max_gas"] = gasStr
				patched = true
			}
		}
	}
	if cp, ok := genesis["consensus_params"].(map[string]interface{}); ok {
		if block, ok := cp["block"].(map[string]interface{}); ok {
			block["max_gas"] = gasStr
			patched = true
		}
	}
	if !patched {
		return fmt.Errorf("genesis has no consensus_params.block or consensus.params.block to patch max_gas")
	}

	out, err := json.MarshalIndent(genesis, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal genesis: %w", err)
	}

	return os.WriteFile(genesisPath, out, 0o600)
}

func (c *Cluster) distributeGenesis(baseHome string) error {
	baseGenesis := filepath.Join(baseHome, "config", "genesis.json")

	if c.opts.MaxBlockGas > 0 {
		if err := patchGenesisBlockGas(baseGenesis, c.opts.MaxBlockGas); err != nil {
			return fmt.Errorf("patch genesis max_gas: %w", err)
		}
	}

	for i := 1; i < len(c.nodes); i++ {
		n := c.nodes[i]
		if err := copyFile(baseGenesis, filepath.Join(n.Home, "config", "genesis.json")); err != nil {
			return err
		}
	}

	return nil
}

func (c *Cluster) configureNodes() error {
	for _, n := range c.nodes {
		cfgPath := filepath.Join(n.Home, "config", "config.toml")
		appPath := filepath.Join(n.Home, "config", "app.toml")

		if err := setTOMLValue(cfgPath, "rpc", "laddr", fmt.Sprintf("\"tcp://127.0.0.1:%d\"", n.Ports.RPC)); err != nil {
			return err
		}
		if err := setTOMLValue(cfgPath, "p2p", "laddr", fmt.Sprintf("\"tcp://127.0.0.1:%d\"", n.Ports.P2P)); err != nil {
			return err
		}
		if err := setTOMLValue(cfgPath, "p2p", "allow_duplicate_ip", "true"); err != nil {
			return err
		}
		if err := setTOMLValue(cfgPath, "p2p", "addr_book_strict", "false"); err != nil {
			return err
		}

		if err := setTOMLValue(appPath, "api", "enable", "true"); err != nil {
			return err
		}
		if err := setTOMLValue(appPath, "api", "swagger", "true"); err != nil {
			return err
		}
		if err := setTOMLValue(appPath, "api", "address", fmt.Sprintf("\"tcp://127.0.0.1:%d\"", n.Ports.API)); err != nil {
			return err
		}
		if err := setTOMLValue(appPath, "grpc", "address", fmt.Sprintf("\"127.0.0.1:%d\"", n.Ports.GRPC)); err != nil {
			return err
		}

		if err := setTOMLValue(appPath, "json-rpc", "enable", "true"); err != nil {
			return err
		}
		if err := setTOMLValue(appPath, "json-rpc", "address", fmt.Sprintf("\"127.0.0.1:%d\"", n.Ports.JSONRPC)); err != nil {
			return err
		}
		if err := setTOMLValue(appPath, "json-rpc", "address-ws", fmt.Sprintf("\"127.0.0.1:%d\"", n.Ports.WS)); err != nil {
			return err
		}

		memiavlValue := "false"
		if c.opts.MemIAVL {
			memiavlValue = "true"
		}
		if err := setTOMLValue(appPath, "memiavl", "enable", memiavlValue); err != nil {
			return err
		}
	}

	for _, n := range c.nodes {
		peers := make([]string, 0, len(c.nodes)-1)
		for _, other := range c.nodes {
			if other.Index == n.Index {
				continue
			}
			peers = append(peers, fmt.Sprintf("%s@127.0.0.1:%d", other.PeerID, other.Ports.P2P))
		}
		if len(peers) == 0 {
			continue
		}
		cfgPath := filepath.Join(n.Home, "config", "config.toml")
		if err := setTOMLValue(cfgPath, "p2p", "persistent_peers", fmt.Sprintf("\"%s\"", strings.Join(peers, ","))); err != nil {
			return err
		}
	}

	return nil
}

func (c *Cluster) startNode(ctx context.Context, n *Node) error {
	logPath := filepath.Join(n.Home, "node.log")
	f, err := os.Create(logPath)
	if err != nil {
		return err
	}

	//nolint:gosec
	cmd := exec.CommandContext(ctx, c.bin, "start", "--home", n.Home)
	cmd.Stdout = f
	cmd.Stderr = f
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	if err := cmd.Start(); err != nil {
		_ = f.Close()
		return err
	}

	n.LogPath = logPath
	n.logFile = f
	n.cmd = cmd

	return nil
}

// ---------------------------------------------------------------------------
// Internal helpers
// ---------------------------------------------------------------------------

func (c *Cluster) nodeStatus(ctx context.Context, n *Node) (bool, int64, error) {
	url := fmt.Sprintf("http://127.0.0.1:%d/status", n.Ports.RPC)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return false, 0, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, 0, fmt.Errorf("status code %d", resp.StatusCode)
	}

	var decoded struct {
		Result struct {
			SyncInfo struct {
				LatestBlockHeight string `json:"latest_block_height"`
				CatchingUp        bool   `json:"catching_up"`
			} `json:"sync_info"`
		} `json:"result"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&decoded); err != nil {
		return false, 0, err
	}

	h, err := strconv.ParseInt(decoded.Result.SyncInfo.LatestBlockHeight, 10, 64)
	if err != nil {
		return false, 0, err
	}

	return !decoded.Result.SyncInfo.CatchingUp, h, nil
}

func (c *Cluster) latestHeight(ctx context.Context, nodeIndex int) (int64, error) {
	n, err := c.getNode(nodeIndex)
	if err != nil {
		return 0, err
	}
	_, h, err := c.nodeStatus(ctx, n)

	return h, err
}

func (c *Cluster) unconfirmedTxCount(ctx context.Context, nodeIndex int) (int64, error) {
	n, err := c.getNode(nodeIndex)
	if err != nil {
		return 0, err
	}
	url := fmt.Sprintf("http://127.0.0.1:%d/num_unconfirmed_txs", n.Ports.RPC)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return 0, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("status code %d", resp.StatusCode)
	}

	var decoded struct {
		Result struct {
			Total string `json:"total"`
		} `json:"result"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&decoded); err != nil {
		return 0, err
	}

	return strconv.ParseInt(decoded.Result.Total, 10, 64)
}

func (c *Cluster) getNode(index int) (*Node, error) {
	if index < 0 || index >= len(c.nodes) {
		return nil, fmt.Errorf("invalid node index %d", index)
	}

	return c.nodes[index], nil
}

func (c *Cluster) keyAddress(ctx context.Context, name string) (string, error) {
	out, err := c.exec(ctx,
		"keys", "show", name,
		"-a",
		"--keyring-backend", "test",
		"--home", c.nodes[0].Home,
	)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}

func (c *Cluster) exec(ctx context.Context, args ...string) ([]byte, error) {
	//nolint:gosec
	cmd := exec.CommandContext(ctx, c.bin, args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		combined := stdout.String() + stderr.String()
		return nil, fmt.Errorf("%s %s failed: %w\n%s", c.bin, strings.Join(args, " "), err, combined)
	}

	return stdout.Bytes(), nil
}

func buildMinitiad(ctx context.Context, repoRoot, outPath string) error {
	cmd := exec.CommandContext(ctx, "go", "build", "-o", outPath, "./cmd/minitiad")
	cmd.Dir = repoRoot

	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("go build failed: %w\n%s", err, string(out))
	}

	return nil
}

func findRepoRoot() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	current := wd
	for {
		if _, err := os.Stat(filepath.Join(current, "go.mod")); err == nil {
			modData, readErr := os.ReadFile(filepath.Join(current, "go.mod"))
			if readErr == nil && strings.Contains(string(modData), "github.com/initia-labs/minievm\n") {
				return current, nil
			}
			if readErr == nil && !strings.Contains(string(modData), "/integration-tests") {
				return current, nil
			}
		}
		next := filepath.Dir(current)
		if next == current {
			break
		}
		current = next
	}

	current = wd
	for {
		if _, err := os.Stat(filepath.Join(current, "go.mod")); err == nil {
			return current, nil
		}
		next := filepath.Dir(current)
		if next == current {
			break
		}
		current = next
	}

	return "", errors.New("go.mod not found from current directory")
}

func copyFile(src, dst string) error {
	bz, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	return os.WriteFile(dst, bz, 0o600)
}

func copyDir(srcDir, dstDir string) error {
	if err := os.MkdirAll(dstDir, 0o700); err != nil {
		return err
	}
	entries, err := os.ReadDir(srcDir)
	if err != nil {
		return err
	}
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		if err := copyFile(filepath.Join(srcDir, e.Name()), filepath.Join(dstDir, e.Name())); err != nil {
			return err
		}
	}

	return nil
}

func parseTxResultFromOutput(out []byte) (TxResult, error) {
	var txResp map[string]any
	if err := json.Unmarshal(out, &txResp); err != nil {
		jsonOut, extractErr := extractJSONObject(out)
		if extractErr != nil {
			return TxResult{}, fmt.Errorf("failed to parse tx response: %w", err)
		}
		if err := json.Unmarshal(jsonOut, &txResp); err != nil {
			return TxResult{}, fmt.Errorf("failed to parse extracted tx response: %w", err)
		}
	}

	code, _ := findIntField(txResp, "code")
	txHash, _ := txResp["txhash"].(string)
	rawLog, _ := txResp["raw_log"].(string)

	return TxResult{
		Code:   code,
		TxHash: txHash,
		RawLog: rawLog,
	}, nil
}

func findUintField(v any, key string) (uint64, bool) {
	switch vv := v.(type) {
	case map[string]any:
		if raw, ok := vv[key]; ok {
			switch x := raw.(type) {
			case string:
				n, err := strconv.ParseUint(x, 10, 64)
				if err == nil {
					return n, true
				}
			case float64:
				return uint64(x), true
			}
		}
		for _, child := range vv {
			if n, ok := findUintField(child, key); ok {
				return n, true
			}
		}
	case []any:
		for _, child := range vv {
			if n, ok := findUintField(child, key); ok {
				return n, true
			}
		}
	}

	return 0, false
}

func findIntField(v any, key string) (int64, bool) {
	switch vv := v.(type) {
	case map[string]any:
		if raw, ok := vv[key]; ok {
			switch x := raw.(type) {
			case string:
				n, err := strconv.ParseInt(x, 10, 64)
				if err == nil {
					return n, true
				}
			case float64:
				return int64(x), true
			}
		}
		for _, child := range vv {
			if n, ok := findIntField(child, key); ok {
				return n, true
			}
		}
	case []any:
		for _, child := range vv {
			if n, ok := findIntField(child, key); ok {
				return n, true
			}
		}
	}

	return 0, false
}

func extractJSONObject(out []byte) ([]byte, error) {
	s := strings.TrimSpace(string(out))
	start := strings.IndexByte(s, '{')
	end := strings.LastIndexByte(s, '}')
	if start == -1 || end == -1 || end <= start {
		return nil, errors.New("json object not found in output")
	}

	return []byte(s[start : end+1]), nil
}

func base64Decode(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

func sha256Hash(data []byte) string {
	h := sha256.Sum256(data)

	return hex.EncodeToString(h[:])
}

func parseEstimatedGas(out []byte) (uint64, error) {
	var txResp map[string]any
	if err := json.Unmarshal(out, &txResp); err == nil {
		if n, ok := findUintField(txResp, "gas_limit"); ok && n > 0 {
			return n, nil
		}
		if n, ok := findUintField(txResp, "gasLimit"); ok && n > 0 {
			return n, nil
		}
		if n, ok := findUintField(txResp, "gas_wanted"); ok && n > 0 {
			return n, nil
		}
		if n, ok := findUintField(txResp, "gasWanted"); ok && n > 0 {
			return n, nil
		}
	}

	re := regexp.MustCompile(`gas estimate:\s*([0-9]+)`)
	m := re.FindSubmatch(out)
	if len(m) == 2 {
		n, err := strconv.ParseUint(string(m[1]), 10, 64)
		if err == nil && n > 0 {
			return n, nil
		}
	}

	return 0, fmt.Errorf("failed to parse estimated gas from output: %s", strings.TrimSpace(string(out)))
}
