package types

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/stretchr/testify/require"
)

func TestNewTracing(t *testing.T) {
	t.Run("create_new_tracing", func(t *testing.T) {
		// Create a mock EVM
		evm := &vm.EVM{
			Context: vm.BlockContext{
				Coinbase:    common.HexToAddress("0x1234567890123456789012345678901234567890"),
				BlockNumber: big.NewInt(100),
				Time:        1234567890,
			},
		}

		// Create a mock tracer
		mockTracer := &tracing.Hooks{}

		// Create new tracing
		tracing := NewTracing(evm, mockTracer)

		// Verify the tracing was created
		require.NotNil(t, tracing)
		require.NotNil(t, tracing.tracer)
		require.NotNil(t, tracing.vmContext)

		// Verify VM context was set correctly
		require.Equal(t, evm.Context.Coinbase, tracing.vmContext.Coinbase)
		require.Equal(t, evm.Context.BlockNumber, tracing.vmContext.BlockNumber)
		require.Equal(t, evm.Context.Time, tracing.vmContext.Time)
		require.Equal(t, evm.Context.BaseFee, tracing.vmContext.BaseFee)
	})

	t.Run("create_tracing_with_nil_tracer", func(t *testing.T) {
		evm := &vm.EVM{
			Context: vm.BlockContext{
				Coinbase:    common.HexToAddress("0x1234567890123456789012345678901234567890"),
				BlockNumber: big.NewInt(100),
				Time:        1234567890,
			},
		}

		require.Panics(t, func() {
			_ = NewTracing(evm, nil)
		})
	})
}

func TestTracing_Methods(t *testing.T) {
	evm := &vm.EVM{
		Context: vm.BlockContext{
			Coinbase:    common.HexToAddress("0x1234567890123456789012345678901234567890"),
			BlockNumber: big.NewInt(100),
			Time:        1234567890,
		},
		TxContext: vm.TxContext{
			GasPrice: big.NewInt(20000000000),
		},
	}

	mockTracer := &tracing.Hooks{}
	tracing := NewTracing(evm, mockTracer)

	t.Run("tracer_method", func(t *testing.T) {
		tracer := tracing.Tracer()
		require.NotNil(t, tracer)
		require.Equal(t, tracing.tracer, tracer)
	})

	t.Run("vm_context_method", func(t *testing.T) {
		vmContext := tracing.VMContext()
		require.NotNil(t, vmContext)
		require.Equal(t, tracing.vmContext, vmContext)
	})
}

func TestTracingContext(t *testing.T) {
	t.Run("create_tracing_context", func(t *testing.T) {
		evm := &vm.EVM{
			Context: vm.BlockContext{
				Coinbase:    common.HexToAddress("0x1234567890123456789012345678901234567890"),
				BlockNumber: big.NewInt(100),
				Time:        1234567890,
			},
		}

		vmContext := tracingContext(evm)

		require.NotNil(t, vmContext)
		require.Equal(t, evm.Context.Coinbase, vmContext.Coinbase)
		require.Equal(t, evm.Context.BlockNumber, vmContext.BlockNumber)
		require.Equal(t, evm.Context.Time, vmContext.Time)
		require.Equal(t, evm.Context.BaseFee, vmContext.BaseFee)
	})
}

func TestTracingTx(t *testing.T) {
	t.Run("create_tracing_transaction", func(t *testing.T) {
		gasLimit := uint64(21000)
		tx := TracingTx(gasLimit)

		require.NotNil(t, tx)
		require.Equal(t, gasLimit, tx.Gas())
	})

	t.Run("create_tracing_transaction_with_zero_gas", func(t *testing.T) {
		gasLimit := uint64(0)
		tx := TracingTx(gasLimit)

		require.NotNil(t, tx)
		require.Equal(t, gasLimit, tx.Gas())
	})

	t.Run("create_tracing_transaction_with_high_gas", func(t *testing.T) {
		gasLimit := uint64(1000000)
		tx := TracingTx(gasLimit)

		require.NotNil(t, tx)
		require.Equal(t, gasLimit, tx.Gas())
	})
}

func TestTracing_DepthAdjustment(t *testing.T) {
	t.Run("depth_adjustment_in_on_enter", func(t *testing.T) {
		evm := &vm.EVM{
			Context: vm.BlockContext{
				Coinbase:    common.HexToAddress("0x1234567890123456789012345678901234567890"),
				BlockNumber: big.NewInt(100),
				Time:        1234567890,
			},
		}

		var capturedDepth int
		mockTracer := &tracing.Hooks{
			OnEnter: func(depth int, typ byte, from common.Address, to common.Address, input []byte, gas uint64, value *big.Int) {
				capturedDepth = depth
			},
		}

		tracing := NewTracing(evm, mockTracer)

		// Call OnEnter with depth 0
		tracing.tracer.OnEnter(0, 0, common.Address{}, common.Address{}, []byte{}, 0, big.NewInt(0))

		// Verify depth was adjusted (0 + 1 = 1)
		require.Equal(t, 1, capturedDepth)
	})

	t.Run("depth_adjustment_in_on_exit", func(t *testing.T) {
		evm := &vm.EVM{
			Context: vm.BlockContext{
				Coinbase:    common.HexToAddress("0x1234567890123456789012345678901234567890"),
				BlockNumber: big.NewInt(100),
				Time:        1234567890,
			},
		}

		var capturedDepth int
		mockTracer := &tracing.Hooks{
			OnExit: func(depth int, output []byte, gasUsed uint64, err error, reverted bool) {
				capturedDepth = depth
			},
		}

		tracing := NewTracing(evm, mockTracer)

		// Call OnExit with depth 2
		tracing.tracer.OnExit(2, []byte{}, 0, nil, false)

		// Verify depth was adjusted (2 + 1 = 3)
		require.Equal(t, 3, capturedDepth)
	})
}
