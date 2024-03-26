package ibc_middleware

import (
	"testing"

	"github.com/stretchr/testify/require"

	evmtypes "github.com/initia-labs/minievm/x/evm/types"
)

func Test_validateAndParseMemo(t *testing.T) {
	memo := `
	{
		"evm" : {
			"sender": "init_addr",
			"contract_addr": "contract_addr",
			"input": ""
		}
	}`
	isEVMRouted, msg, err := validateAndParseMemo(memo, "contract_addr")
	require.True(t, isEVMRouted)
	require.NoError(t, err)
	require.Equal(t, evmtypes.MsgCall{
		Sender:       "init_addr",
		ContractAddr: "contract_addr",
		Input:        "",
	}, msg)

	// invalid receiver
	isEVMRouted, _, err = validateAndParseMemo(memo, "invalid_addr")
	require.True(t, isEVMRouted)
	require.Error(t, err)

	isEVMRouted, _, err = validateAndParseMemo("hihi", "invalid_addr")
	require.False(t, isEVMRouted)
	require.NoError(t, err)
}
