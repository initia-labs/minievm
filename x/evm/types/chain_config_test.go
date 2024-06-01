package types_test

import (
	"math/big"
	"testing"

	"github.com/initia-labs/minievm/x/evm/types"
	"github.com/stretchr/testify/require"
)

func Test_ConvertChainIdBiDirectional(t *testing.T) {
	chainID := "minievm"
	expectedEthChainID := new(big.Int).SetUint64(13568054635622948241)
	ethChainID := types.ConvertCosmosChainIDToEthereumChainID(chainID)
	require.Equal(t, ethChainID, (expectedEthChainID))
}
