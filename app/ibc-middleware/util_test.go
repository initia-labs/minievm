package ibc_middleware

import (
	"testing"

	"github.com/stretchr/testify/require"

	evmtypes "github.com/initia-labs/minievm/x/evm/types"
)

func Test_validateAndParseMemo(t *testing.T) {
	memo := `
	{
		"wasm" : {
			"sender": "init_addr",
			"contract_addr": "contract_addr",
			"input": {}
		}
	}`
	isWasmRouted, msg, err := validateAndParseMemo(memo, "contract_addr")
	require.True(t, isWasmRouted)
	require.NoError(t, err)
	require.Equal(t, evmtypes.MsgCall{
		Sender:       "init_addr",
		ContractAddr: "contract_addr",
		Input:        []byte("{}"),
	}, msg)

	// invalid receiver
	isWasmRouted, _, err = validateAndParseMemo(memo, "invalid_addr")
	require.True(t, isWasmRouted)
	require.Error(t, err)

	isWasmRouted, _, err = validateAndParseMemo("hihi", "invalid_addr")
	require.False(t, isWasmRouted)
	require.NoError(t, err)
}
