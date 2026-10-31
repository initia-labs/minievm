package tests

import (
	"crypto/ecdsa"
	"math/big"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	abcitypes "github.com/cometbft/cometbft/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	minitiaapp "github.com/initia-labs/minievm/app"
)

// asm is a tiny EVM assembler with jump labels, used to build test contracts
// without a solidity toolchain.
type asm struct {
	code   []byte
	labels map[string]int
	fixups map[int]string
}

func newAsm() *asm { return &asm{labels: map[string]int{}, fixups: map[int]string{}} }

func (a *asm) op(b ...byte) *asm { a.code = append(a.code, b...); return a }

func (a *asm) label(name string) *asm {
	a.labels[name] = len(a.code)
	return a.op(0x5b) // JUMPDEST
}

func (a *asm) push2label(name string) *asm {
	a.op(0x61)
	a.fixups[len(a.code)] = name
	return a.op(0, 0)
}

func (a *asm) bytes() []byte {
	for pos, name := range a.fixups {
		dst := a.labels[name]
		a.code[pos] = byte(dst >> 8)
		a.code[pos+1] = byte(dst)
	}
	return a.code
}

// loop runs body n times where n is the calldata word at off. The body must
// keep the counter on top of the stack.
func (a *asm) loop(name string, off byte, body []byte) *asm {
	a.op(0x60, off, 0x35) // PUSH1 off CALLDATALOAD
	a.label("L" + name)
	a.op(0x80, 0x15).push2label("E" + name).op(0x57) // DUP1 ISZERO PUSH2 E JUMPI
	a.op(body...)
	a.op(0x60, 0x01, 0x90, 0x03) // PUSH1 1 SWAP1 SUB
	a.push2label("L" + name).op(0x56)
	a.label("E" + name)
	return a.op(0x50) // POP
}

// loopRuntime takes four calldata words: the number of CALLs to a codeless
// address, TSTOREs to distinct keys, SSTOREs to distinct keys and empty loop
// iterations, run in that order.
func loopRuntime() []byte {
	callBody := []byte{
		0x60, 0x00, 0x60, 0x00, 0x60, 0x00, 0x60, 0x00, 0x60, 0x00, // retSize retOff argSize argOff value
		0x61, 0xde, 0xad, // PUSH2 0xdead (codeless, nonexistent)
		0x5a, 0xf1, 0x50, // GAS CALL POP
	}
	a := newAsm()
	a.loop("call", 0x00, callBody)
	a.loop("tstore", 0x20, []byte{0x80, 0x80, 0x5d}) // DUP1 DUP1 TSTORE
	a.loop("sstore", 0x40, []byte{0x80, 0x80, 0x55}) // DUP1 DUP1 SSTORE
	a.loop("plain", 0x60, nil)
	a.op(0x00) // STOP
	return a.bytes()
}

func loopInput(nCall, nTStore, nSStore, nPlain uint64) []byte {
	return append(append(append(word(nCall), word(nTStore)...), word(nSStore)...), word(nPlain)...)
}

// initCode wraps runtime code in creation code that returns it.
func initCode(rt []byte) []byte {
	n := len(rt)
	pre := []byte{
		0x61, byte(n >> 8), byte(n), // PUSH2 len
		0x60, 0x0e, // PUSH1 offset(=14)
		0x60, 0x00, // PUSH1 0
		0x39,                        // CODECOPY
		0x61, byte(n >> 8), byte(n), // PUSH2 len
		0x60, 0x00, // PUSH1 0
		0xf3, // RETURN
	}
	return append(pre, rt...)
}

func word(v uint64) []byte { return common.BigToHash(new(big.Int).SetUint64(v)).Bytes() }

type ecdsaKey struct {
	k    *ecdsa.PrivateKey
	addr common.Address
}

// deployRuntime deploys the given runtime code from key and returns its address.
func deployRuntime(t *testing.T, app *minitiaapp.MinitiaApp, key *ecdsaKey, rt []byte) common.Address {
	ctx, closer, err := app.CreateQueryContext(0, false)
	require.NoError(t, err)
	nonce, err := app.AccountKeeper.GetSequence(ctx, sdk.AccAddress(key.addr.Bytes()))
	require.NoError(t, err)
	if closer != nil {
		closer.Close()
	}

	tx, _ := GenerateTx(t, app, key.k, nil, initCode(rt), new(big.Int), SetGasLimit(5_000_000))
	_, res := ExecuteTxs(t, app, tx)
	CheckTxResult(t, res.TxResults[0], true)
	return crypto.CreateAddress(key.addr, nonce)
}

func deployLoop(t *testing.T, app *minitiaapp.MinitiaApp, key *ecdsaKey) common.Address {
	return deployRuntime(t, app, key, loopRuntime())
}

type loopResult struct {
	finalize, commit time.Duration
	gasUsed          int64
	code             uint32
	log              string
	allocMB          float64
}

// runLoop executes one loop-contract tx in its own block and reports time and
// memory allocated during FinalizeBlock.
func runLoop(t *testing.T, app *minitiaapp.MinitiaApp, key *ecdsaKey, contract common.Address, gas uint64, nCall, nTStore, nSStore, nPlain uint64) loopResult {
	tx, _ := GenerateTx(t, app, key.k, &contract, loopInput(nCall, nTStore, nSStore, nPlain), new(big.Int), SetGasLimit(gas))
	txBytes, err := app.TxConfig().TxEncoder()(tx)
	require.NoError(t, err)

	runtime.GC()
	var before runtime.MemStats
	runtime.ReadMemStats(&before)

	start := time.Now()
	resBlock, err := app.FinalizeBlock(&abcitypes.RequestFinalizeBlock{Txs: [][]byte{txBytes}, Height: app.LastBlockHeight() + 1})
	require.NoError(t, err)
	fin := time.Since(start)

	var after runtime.MemStats
	runtime.ReadMemStats(&after)

	start = time.Now()
	_, err = app.Commit()
	require.NoError(t, err)
	com := time.Since(start)
	app.EVMIndexer().Wait()

	r := resBlock.TxResults[0]
	return loopResult{fin, com, r.GasUsed, r.Code, r.Log, float64(after.TotalAlloc-before.TotalAlloc) / (1 << 20)}
}

func trunc(s string) string {
	if len(s) > 120 {
		return s[:120]
	}
	return s
}
