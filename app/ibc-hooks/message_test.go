package evm_hooks_test

import (
	"encoding/json"
	"testing"

	evmhooks "github.com/initia-labs/minievm/app/ibc-hooks"
	"github.com/stretchr/testify/require"
)

func Test_Unmarshal_AsyncCallback(t *testing.T) {
	var callback evmhooks.AsyncCallback
	err := json.Unmarshal([]byte(`{
		"id": 99,
		"contract_address": "0x1"
	}`), &callback)
	require.NoError(t, err)
	require.Equal(t, evmhooks.AsyncCallback{
		Id:              99,
		ContractAddress: "0x1",
	}, callback)

	var callbackStringID evmhooks.AsyncCallback
	err = json.Unmarshal([]byte(`{
		"id": "99",
		"contract_address": "0x1"
	}`), &callbackStringID)
	require.NoError(t, err)
	require.Equal(t, callback, callbackStringID)
}
