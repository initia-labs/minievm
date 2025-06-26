package types

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_txMetadata_EncodingDecoding(t *testing.T) {
	meta := txMetadata{
		Type:      2,
		GasFeeCap: big.NewInt(0),
		GasTipCap: big.NewInt(0),
		GasLimit:  100,
	}

	bz, err := json.Marshal(meta)
	require.NoError(t, err)

	var meta2 txMetadata
	err = json.Unmarshal(bz, &meta2)
	require.NoError(t, err)
	require.Equal(t, meta, meta2)

	require.True(t, meta2.GasFeeCap.Uint64() == 0)
	require.True(t, meta2.GasTipCap.Uint64() == 0)
}
