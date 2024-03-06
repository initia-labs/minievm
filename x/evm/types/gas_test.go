package types_test

import (
	"testing"

	"github.com/initia-labs/minievm/x/evm/types"
	"github.com/stretchr/testify/require"
)

func Test_CalGasUsed(t *testing.T) {
	gasUsed := types.CalGasUsed(100, 50, 10)
	require.Equal(t, uint64(40), gasUsed)

	gasUsed = types.CalGasUsed(100, 50, 5)
	require.Equal(t, uint64(45), gasUsed)
}
