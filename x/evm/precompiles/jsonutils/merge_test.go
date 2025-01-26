package jsonutils_test

import (
	"testing"

	"github.com/initia-labs/minievm/x/evm/precompiles/jsonutils"
	"github.com/stretchr/testify/require"
)

func Test_MergeJSON(t *testing.T) {

	testCases := []struct {
		name     string
		dst      string
		src      string
		expected string
	}{
		{
			name:     "simple merge",
			dst:      `{"a": 1, "b": 2}`,
			src:      `{"b": 3, "c": 4}`,
			expected: `{"a":1,"b":3,"c":4}`,
		},
		{
			name:     "nested merge",
			dst:      `{"a": 1, "b": {"c": 2}}`,
			src:      `{"b": {"d": 3}, "c": 4}`,
			expected: `{"a":1,"b":{"c":2,"d":3},"c":4}`,
		},
		{
			name:     "nested merge with conflict",
			dst:      `{"a": 1, "b": {"c": 2}}`,
			src:      `{"b": {"c": 3}, "c": 4}`,
			expected: `{"a":1,"b":{"c":3},"c":4}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := jsonutils.MergeJSON(tc.dst, tc.src)
			require.NoError(t, err)
			require.Equal(t, tc.expected, res)
		})
	}
}
