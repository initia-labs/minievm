package backend_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ClientVersion(t *testing.T) {
	test := setupBackend(t)
	backend, cometRPC := test.backend, test.cometRPC

	cometRPC.ClientVersion = "v0.6.2"
	version, err := backend.ClientVersion()
	require.NoError(t, err)
	require.Equal(t, "v0.6.2", version)
}
