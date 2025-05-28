package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/tracing"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
)

// Tracing wraps tracing.Hooks to adjust call depth tracking.
// When executing transactions in minievm, a dummy call is created with depth 0.
// This wrapper adds 1 to the depth passed to tracers to account for the dummy call
// and maintain accurate call stack depth reporting.
type Tracing struct {
	tracer    *tracing.Hooks
	vmContext *tracing.VMContext
}

// NewTracing creates a new Tracing
//
//	evm: the EVM instance
//	tracer: the tracer instance
//
// Returns:
// - *Tracing: the new Tracing instance
func NewTracing(evm *vm.EVM, tracer *tracing.Hooks) *Tracing {
	return &Tracing{
		tracer: &tracing.Hooks{
			// VM events
			OnTxStart: tracer.OnTxStart,
			OnTxEnd:   tracer.OnTxEnd,
			OnEnter: func(depth int, typ byte, from common.Address, to common.Address, input []byte, gas uint64, value *big.Int) {
				if tracer.OnEnter != nil {
					tracer.OnEnter(depth+1, typ, from, to, input, gas, value)
				}
			},
			OnExit: func(depth int, output []byte, gasUsed uint64, err error, reverted bool) {
				if tracer.OnExit != nil {
					tracer.OnExit(depth+1, output, gasUsed, err, reverted)
				}
			},
			OnOpcode: func(pc uint64, op byte, gas, cost uint64, scope tracing.OpContext, rData []byte, depth int, err error) {
				if tracer.OnOpcode != nil {
					tracer.OnOpcode(pc, op, gas, cost, scope, rData, depth+1, err)
				}
			},
			OnFault: func(pc uint64, op byte, gas, cost uint64, scope tracing.OpContext, depth int, err error) {
				if tracer.OnFault != nil {
					tracer.OnFault(pc, op, gas, cost, scope, depth+1, err)
				}
			},
			OnGasChange: tracer.OnGasChange,
			// Chain events
			OnBlockchainInit:  tracer.OnBlockchainInit,
			OnClose:           tracer.OnClose,
			OnBlockStart:      tracer.OnBlockStart,
			OnBlockEnd:        tracer.OnBlockEnd,
			OnSkippedBlock:    tracer.OnSkippedBlock,
			OnGenesisBlock:    tracer.OnGenesisBlock,
			OnSystemCallStart: tracer.OnSystemCallStart,
			OnSystemCallEnd:   tracer.OnSystemCallEnd,
			// State events
			OnBalanceChange: tracer.OnBalanceChange,
			OnNonceChange:   tracer.OnNonceChange,
			OnCodeChange:    tracer.OnCodeChange,
			OnStorageChange: tracer.OnStorageChange,
			OnLog:           tracer.OnLog,
		},
		vmContext: tracingContext(evm),
	}
}

func (t *Tracing) Tracer() *tracing.Hooks {
	return t.tracer
}

func (t *Tracing) VMContext() *tracing.VMContext {
	return t.vmContext
}

func tracingContext(evm *vm.EVM) *tracing.VMContext {
	return &tracing.VMContext{
		Coinbase:    evm.Context.Coinbase,
		BlockNumber: evm.Context.BlockNumber,
		Time:        evm.Context.Time,
		Random:      evm.Context.Random,
		GasPrice:    evm.TxContext.GasPrice,
		ChainConfig: evm.ChainConfig(),
		StateDB:     evm.StateDB,
	}
}

func TracingTx(gasLimit uint64) *coretypes.Transaction {
	return coretypes.NewTx(&coretypes.LegacyTx{
		Gas: gasLimit,
	})
}
