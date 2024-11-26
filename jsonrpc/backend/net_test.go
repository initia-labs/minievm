package backend_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Version(t *testing.T) {
	input := setupBackend(t)
	_, _, backend, _, _ := input.app, input.addrs, input.backend, input.addrs, input.privKeys

	version, err := backend.Version()
	require.NoError(t, err)
	require.Equal(t, "1", version)
}

func Test_PeerCount(t *testing.T) {
	input := setupBackend(t)
	backend, cometRPC := input.backend, input.cometRPC

	cometRPC.NPeers = 10
	peerCount, err := backend.PeerCount()
	require.NoError(t, err)
	require.Equal(t, uint(10), uint(peerCount))
}

func Test_Listening(t *testing.T) {
	input := setupBackend(t)
	backend, cometRPC := input.backend, input.cometRPC

	cometRPC.Listening = true
	listening, err := backend.Listening()
	require.NoError(t, err)
	require.True(t, listening)
}
