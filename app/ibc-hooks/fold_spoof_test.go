package evm_hooks

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Test_validateAndParseMemo_foldSpoof is the regression test for the
// ERC20Wrapper callback-id spoof. A user-supplied memo key that folds onto the
// reserved "async_callback" key (e.g. the long-s "aſync_callback" or an ASCII
// case variant) must be rejected.
func Test_validateAndParseMemo_foldSpoof(t *testing.T) {
	// A legitimate memo parses and yields the caller's callback id.
	legit := `{"evm":{"async_callback":{"id":7,"contract_address":"0xabc"}}}`
	isEVM, hook, err := validateAndParseMemo(legit)
	require.NoError(t, err)
	require.True(t, isEVM)
	require.NotNil(t, hook.AsyncCallback)
	require.Equal(t, uint64(7), hook.AsyncCallback.Id)

	// Long-s spoof. "aſync_callback" (U+017F) survives a byte-exact merge yet
	// folds onto "async_callback", overriding the trusted id=7 with id=1 under
	// last-key-wins. Must now be rejected.
	spoofLongS := "{\"evm\":{\"async_callback\":{\"id\":7,\"contract_address\":\"0xabc\"}," +
		"\"aſync_callback\":{\"id\":1,\"contract_address\":\"0xdead\"}}}"
	_, _, err = validateAndParseMemo(spoofLongS)
	require.Error(t, err)

	// ASCII case-fold variant of the same attack.
	spoofASCII := `{"evm":{"async_callback":{"id":7,"contract_address":"0xabc"},"Async_callback":{"id":1,"contract_address":"0xdead"}}}`
	_, _, err = validateAndParseMemo(spoofASCII)
	require.Error(t, err)

	// Long-s spoof on the "message" key ("mesſage" folds onto "message"),
	// confirming the guard is not specific to the async_callback key.
	spoofMessage := "{\"evm\":{\"mesſage\":{\"contract_addr\":\"0x1\"}}}"
	_, _, err = validateAndParseMemo(spoofMessage)
	require.Error(t, err)

	// A benign unrelated ASCII key does not fold onto a reserved key and is left
	// untouched (ignored by the decoder as before).
	benign := `{"evm":{"async_callback":{"id":7,"contract_address":"0xabc"},"note":"hello"}}`
	_, _, err = validateAndParseMemo(benign)
	require.NoError(t, err)
}
