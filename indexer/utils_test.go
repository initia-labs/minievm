package indexer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_collJsonVal(t *testing.T) {
	type s1 struct {
		A string `json:"a"`
		B uint64 `json:"b"`
	}

	codec := collJsonVal[s1]{}
	bz, err := codec.Encode(s1{
		A: "a",
		B: 1,
	})
	require.NoError(t, err)
	require.Equal(t, `{"a":"a","b":1}`, string(bz))

	out, err := codec.Decode(bz)
	require.NoError(t, err)
	require.Equal(t, s1{
		A: "a",
		B: 1,
	}, out)
}
