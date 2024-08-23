package state_test

import (
	"testing"

	"github.com/holiman/uint256"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/tracing"
	coretypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/initia-labs/minievm/x/evm/state"
	evmtypes "github.com/initia-labs/minievm/x/evm/types"
)

func Test_SnapshotRevert(t *testing.T) {
	sdkCtx, input := createDefaultTestInput(t)

	_, _, addr1 := keyPubAddr()
	_, _, addr2 := keyPubAddr()

	feeDenom, err := input.EVMKeeper.GetFeeDenom(sdkCtx)
	require.NoError(t, err)
	input.Faucet.Fund(sdkCtx, addr2, sdk.NewInt64Coin(feeDenom, 100))

	_, evm, err := input.EVMKeeper.CreateEVM(sdkCtx, evmtypes.StdAddress, nil)
	require.NoError(t, err)
	stateDB := evm.StateDB.(*state.StateDB)

	require.Equal(t, uint64(0), stateDB.GetBalance(common.BytesToAddress(addr1)).Uint64())
	require.Equal(t, uint64(0), stateDB.GetRefund())
	require.Equal(t, evmtypes.Logs{}, stateDB.Logs())
	require.False(t, stateDB.AddressInAccessList(common.BytesToAddress(addr1)))
	require.False(t, stateDB.AddressInAccessList(common.BytesToAddress(addr2)))

	// take snapshot
	sid := stateDB.Snapshot()

	log1 := &coretypes.Log{
		Address: common.BytesToAddress(addr1),
		Topics:  []common.Hash{{1}, {2}, {3}},
		Data:    []byte{0, 1, 2, 3},
	}
	log2 := &coretypes.Log{
		Address: common.BytesToAddress(addr2),
		Topics:  []common.Hash{{4}, {5}, {6}},
		Data:    []byte{4, 5, 6, 7},
	}

	stateDB.AddLog(log1)
	stateDB.AddLog(log2)
	stateDB.AddBalance(common.BytesToAddress(addr1), uint256.NewInt(100), tracing.BalanceIncreaseSelfdestruct)
	stateDB.SubBalance(common.BytesToAddress(addr1), uint256.NewInt(10), tracing.BalanceDecreaseSelfdestructBurn)
	stateDB.AddRefund(100)
	stateDB.SubRefund(10)
	stateDB.AddAddressToAccessList(common.BytesToAddress(addr1))
	stateDB.AddAddressToAccessList(common.BytesToAddress(addr2))

	require.Equal(t, uint64(90), stateDB.GetBalance(common.BytesToAddress(addr1)).Uint64())
	require.Equal(t, uint64(90), stateDB.GetRefund())
	require.Equal(t, evmtypes.NewLogs([]*coretypes.Log{log1, log2}), stateDB.Logs()[:2])
	require.True(t, stateDB.AddressInAccessList(common.BytesToAddress(addr1)))
	require.True(t, stateDB.AddressInAccessList(common.BytesToAddress(addr2)))

	// revert to snapshot
	stateDB.RevertToSnapshot(sid)

	require.Equal(t, uint64(0), stateDB.GetBalance(common.BytesToAddress(addr1)).Uint64())
	require.Equal(t, uint64(0), stateDB.GetRefund())
	require.Equal(t, evmtypes.Logs{}, stateDB.Logs())
	require.False(t, stateDB.AddressInAccessList(common.BytesToAddress(addr1)))
	require.False(t, stateDB.AddressInAccessList(common.BytesToAddress(addr2)))
}
