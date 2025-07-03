package types

import (
	"testing"

	"cosmossdk.io/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/require"
)

func Test_CalGasUsed(t *testing.T) {
	// Test basic gas calculation
	gasUsed := CalGasUsed(100, 50, 10)
	require.Equal(t, uint64(40), gasUsed)

	// Test with refund less than calculated refund
	gasUsed = CalGasUsed(100, 50, 5)
	require.Equal(t, uint64(45), gasUsed)

	// Test with zero gas remaining
	gasUsed = CalGasUsed(100, 0, 10)
	expectedRefund := 100 / params.RefundQuotientEIP3529
	if expectedRefund > 10 {
		expectedRefund = 10
	}
	require.Equal(t, uint64(100-expectedRefund), gasUsed)

	// Test with zero gas refunded
	gasUsed = CalGasUsed(100, 50, 0)
	require.Equal(t, uint64(50), gasUsed)

	// Test with large refund that gets capped
	gasUsed = CalGasUsed(1000, 100, 1000) // refund would be 900/5=180, capped at 1000, so uses 180
	expectedRefund = 900 / params.RefundQuotientEIP3529
	if expectedRefund > 1000 {
		expectedRefund = 1000
	}
	require.Equal(t, uint64(900-expectedRefund), gasUsed)

	// Test edge case: gasBalance equals gasRemaining
	gasUsed = CalGasUsed(100, 100, 10)
	require.Equal(t, uint64(0), gasUsed)

	// Test edge case: gasRemaining greater than gasBalance (shouldn't happen in practice)
	gasUsed = CalGasUsed(100, 150, 10)
	// This would result in underflow, but uint64 wraps around
	// POTENTIAL BUG: The function doesn't validate that gasRemaining <= gasBalance
	// This could lead to unexpected behavior in edge cases
	// 100 - 150 underflows to 18446744073709551566
	expectedGasUsed := uint64(18446744073709551566)
	expectedRefund = expectedGasUsed / params.RefundQuotientEIP3529
	if expectedRefund > 10 {
		expectedRefund = 10
	}
	require.Equal(t, expectedGasUsed-expectedRefund, gasUsed)
}

func Test_BlockGasLimit(t *testing.T) {
	// Test with basic context (no block gas meter)
	ctx := sdk.NewContext(nil, tmproto.Header{}, false, log.NewNopLogger())
	limit := BlockGasLimit(ctx)
	require.Equal(t, uint64(0), limit)

	// Note: Testing BlockGasLimit with different scenarios requires more complex setup
	// that would involve mocking the SDK's internal gas meter and consensus params.
	// The function is primarily tested through integration tests.
}
