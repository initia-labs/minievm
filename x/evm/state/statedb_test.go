package state_test

import (
	"testing"

	"github.com/holiman/uint256"
	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/tracing"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"

	"github.com/initia-labs/minievm/x/evm/contracts/counter"
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

	// for committed state test
	stateDB.SetState(common.BytesToAddress(addr2), common.BytesToHash([]byte("key")), common.BytesToHash([]byte("value")))
	stateDB.Commit()

	require.Equal(t, uint64(0), stateDB.GetBalance(common.BytesToAddress(addr1)).Uint64())
	require.Equal(t, uint64(0), stateDB.GetRefund())
	require.Equal(t, evmtypes.Logs{}, stateDB.Logs())
	require.False(t, stateDB.AddressInAccessList(common.BytesToAddress(addr1)))
	require.False(t, stateDB.AddressInAccessList(common.BytesToAddress(addr2)))
	stateDB.AddSlotToAccessList(common.BytesToAddress(addr2), common.BytesToHash([]byte("slot1")))
	stateDB.SetTransientState(common.BytesToAddress(addr1), common.BytesToHash([]byte("key")), common.BytesToHash([]byte("value")))
	stateDB.SetState(common.BytesToAddress(addr2), common.BytesToHash([]byte("key")), common.BytesToHash([]byte("value1")))

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
	stateDB.AddSlotToAccessList(common.BytesToAddress(addr1), common.BytesToHash([]byte("slot1")))
	stateDB.AddSlotToAccessList(common.BytesToAddress(addr1), common.BytesToHash([]byte("slot2")))
	stateDB.SetState(common.BytesToAddress(addr1), common.BytesToHash([]byte("key")), common.BytesToHash([]byte("value2")))
	stateDB.SetState(common.BytesToAddress(addr2), common.BytesToHash([]byte("key")), common.BytesToHash([]byte("value2")))

	require.Equal(t, uint64(90), stateDB.GetBalance(common.BytesToAddress(addr1)).Uint64())
	require.Equal(t, uint64(90), stateDB.GetRefund())
	require.Equal(t, evmtypes.NewLogs([]*coretypes.Log{log1, log2}), stateDB.Logs()[:2])
	require.True(t, stateDB.AddressInAccessList(common.BytesToAddress(addr1)))
	require.True(t, stateDB.AddressInAccessList(common.BytesToAddress(addr2)))
	require.Equal(t, common.BytesToHash([]byte("value2")), stateDB.GetState(common.BytesToAddress(addr1), common.BytesToHash([]byte("key"))))
	require.Equal(t, common.BytesToHash([]byte("value2")), stateDB.GetState(common.BytesToAddress(addr2), common.BytesToHash([]byte("key"))))

	// test committed state
	require.Equal(t, common.BytesToHash([]byte("value")), stateDB.GetCommittedState(common.BytesToAddress(addr2), common.BytesToHash([]byte("key"))))

	// take more snapshots
	stateDB.Snapshot()
	stateDB.SubRefund(10)
	stateDB.Snapshot()
	stateDB.SubBalance(common.BytesToAddress(addr1), uint256.NewInt(10), tracing.BalanceDecreaseSelfdestructBurn)

	// revert to snapshot
	stateDB.RevertToSnapshot(sid)

	require.Equal(t, uint64(0), stateDB.GetBalance(common.BytesToAddress(addr1)).Uint64())
	require.Equal(t, uint64(0), stateDB.GetRefund())
	require.Equal(t, evmtypes.Logs{}, stateDB.Logs())
	require.False(t, stateDB.AddressInAccessList(common.BytesToAddress(addr1)))
	require.True(t, stateDB.AddressInAccessList(common.BytesToAddress(addr2)))
	a, s := stateDB.SlotInAccessList(common.BytesToAddress(addr1), common.BytesToHash([]byte("slot1")))
	require.False(t, a || s)
	a, s = stateDB.SlotInAccessList(common.BytesToAddress(addr1), common.BytesToHash([]byte("slot2")))
	require.False(t, a || s)
	a, s = stateDB.SlotInAccessList(common.BytesToAddress(addr2), common.BytesToHash([]byte("slot1")))
	require.True(t, a && s)
	require.Equal(t, common.BytesToHash([]byte("value")), stateDB.GetTransientState(common.BytesToAddress(addr1), common.BytesToHash([]byte("key"))))
	require.Equal(t, common.BytesToHash([]byte("value1")), stateDB.GetState(common.BytesToAddress(addr2), common.BytesToHash([]byte("key"))))
}

func Test_SimpleSnapshotRevert(t *testing.T) {
	sdkCtx, input := createDefaultTestInput(t)

	_, evm, err := input.EVMKeeper.CreateEVM(sdkCtx, evmtypes.StdAddress, nil)
	require.NoError(t, err)
	stateDB := evm.StateDB.(*state.StateDB)

	// take snapshot
	sid := stateDB.Snapshot()
	stateDB.AddRefund(100)
	stateDB.SubRefund(10)

	// revert to snapshot
	stateDB.RevertToSnapshot(sid)

	require.Equal(t, uint64(0), stateDB.GetRefund())
}

func Test_GetStroageRoot_NonEmptyState(t *testing.T) {
	sdkCtx, input := createDefaultTestInput(t)

	_, _, addr1 := keyPubAddr()
	_, evm, err := input.EVMKeeper.CreateEVM(sdkCtx, evmtypes.StdAddress, nil)
	require.NoError(t, err)
	stateDB := evm.StateDB.(*state.StateDB)

	require.Equal(t, common.Hash{}, stateDB.GetStorageRoot(common.Address(addr1.Bytes())))

	stateDB.SetState(common.BytesToAddress(addr1), common.BytesToHash([]byte("key")), common.BytesToHash([]byte("value")))
	require.NotEqual(t, common.Hash{}, stateDB.GetStorageRoot(common.Address(addr1.Bytes())))
}

func Test_GetStorageRoot_NonEmptyCosmosAccount(t *testing.T) {
	sdkCtx, input := createDefaultTestInput(t)

	_, _, addr1 := keyPubAddr()
	_, evm, err := input.EVMKeeper.CreateEVM(sdkCtx, evmtypes.StdAddress, nil)
	require.NoError(t, err)
	stateDB := evm.StateDB.(*state.StateDB)

	require.Equal(t, common.Hash{}, stateDB.GetStorageRoot(common.Address(addr1.Bytes())))

	acc := input.AccountKeeper.NewAccountWithAddress(sdkCtx, addr1)
	err = acc.SetPubKey(&secp256k1.PubKey{})
	require.NoError(t, err)
	input.AccountKeeper.SetAccount(sdkCtx, acc)

	require.NotEqual(t, common.Hash{}, stateDB.GetStorageRoot(common.Address(addr1.Bytes())))
}

func Test_SelfDestruct(t *testing.T) {
	sdkCtx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()

	// fund addr
	input.Faucet.Fund(sdkCtx, addr, sdk.NewInt64Coin(sdk.DefaultBondDenom, 1000000000))

	// deploy contract
	counterBz, err := hexutil.Decode(counter.CounterBin)
	require.NoError(t, err)

	caller := common.BytesToAddress(addr.Bytes())
	_, contractAddr, _, err := input.EVMKeeper.EVMCreate(sdkCtx, caller, counterBz, nil, nil)
	require.NoError(t, err)

	_, contractAddr2, _, err := input.EVMKeeper.EVMCreate(sdkCtx, caller, counterBz, nil, nil)
	require.NoError(t, err)

	// increase counter
	parsed, err := counter.CounterMetaData.GetAbi()
	require.NoError(t, err)

	inputBz, err := parsed.Pack("increase")
	require.NoError(t, err)

	// call with value
	res, logs, err := input.EVMKeeper.EVMCall(sdkCtx, caller, contractAddr, inputBz, uint256.NewInt(100), nil)
	require.NoError(t, err)
	require.Empty(t, res)
	require.NotEmpty(t, logs)

	res, logs, err = input.EVMKeeper.EVMCall(sdkCtx, caller, contractAddr2, inputBz, uint256.NewInt(100), nil)
	require.NoError(t, err)
	require.Empty(t, res)
	require.NotEmpty(t, logs)

	// check destruct
	_, evm, err := input.EVMKeeper.CreateEVM(sdkCtx, evmtypes.StdAddress, nil)
	require.NoError(t, err)
	stateDB := evm.StateDB.(*state.StateDB)

	// check balance
	require.Equal(t, uint64(100), stateDB.GetBalance(contractAddr).Uint64())

	// self destruct
	require.False(t, stateDB.HasSelfDestructed(contractAddr))
	stateDB.SelfDestruct(contractAddr)
	require.True(t, stateDB.HasSelfDestructed(contractAddr))

	// check balance
	require.Equal(t, uint64(0), stateDB.GetBalance(contractAddr).Uint64())

	// should clear only self destructed account
	err = stateDB.Commit()
	require.NoError(t, err)

	// should have empty storage root
	require.Equal(t, common.Hash{}, stateDB.GetStorageRoot(contractAddr))
	require.NotEqual(t, common.Hash{}, stateDB.GetStorageRoot(contractAddr2))

	// should clear cosmos account
	require.Nil(t, input.AccountKeeper.GetAccount(sdkCtx, sdk.AccAddress(contractAddr.Bytes())))
	require.NotNil(t, input.AccountKeeper.GetAccount(sdkCtx, sdk.AccAddress(contractAddr2.Bytes())))
}

func Test_Selfdestruct6780_InDifferentTx(t *testing.T) {
	sdkCtx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()

	// fund addr
	input.Faucet.Fund(sdkCtx, addr, sdk.NewInt64Coin(sdk.DefaultBondDenom, 1000000000))

	// deploy contract
	counterBz, err := hexutil.Decode(counter.CounterBin)
	require.NoError(t, err)

	caller := common.BytesToAddress(addr.Bytes())
	_, contractAddr, _, err := input.EVMKeeper.EVMCreate(sdkCtx, caller, counterBz, nil, nil)
	require.NoError(t, err)

	// increase counter
	parsed, err := counter.CounterMetaData.GetAbi()
	require.NoError(t, err)

	inputBz, err := parsed.Pack("increase")
	require.NoError(t, err)

	// call with value
	res, logs, err := input.EVMKeeper.EVMCall(sdkCtx, caller, contractAddr, inputBz, uint256.NewInt(100), nil)
	require.NoError(t, err)
	require.Empty(t, res)
	require.NotEmpty(t, logs)

	// check destruct
	_, evm, err := input.EVMKeeper.CreateEVM(sdkCtx, evmtypes.StdAddress, nil)
	require.NoError(t, err)
	stateDB := evm.StateDB.(*state.StateDB)

	// check balance
	require.Equal(t, uint64(100), stateDB.GetBalance(contractAddr).Uint64())

	// self destruct
	require.False(t, stateDB.HasSelfDestructed(contractAddr))
	stateDB.Selfdestruct6780(contractAddr)
	require.False(t, stateDB.HasSelfDestructed(contractAddr))

	// check balance
	require.Equal(t, uint64(100), stateDB.GetBalance(contractAddr).Uint64())

	// should clear only self destructed account
	err = stateDB.Commit()
	require.NoError(t, err)

	// should have empty storage root
	require.NotEqual(t, common.Hash{}, stateDB.GetStorageRoot(contractAddr))

	// should clear cosmos account
	require.NotNil(t, input.AccountKeeper.GetAccount(sdkCtx, sdk.AccAddress(contractAddr.Bytes())))
}

func Test_Selfdestruct6780_InSameTx(t *testing.T) {
	sdkCtx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()

	// fund addr
	input.Faucet.Fund(sdkCtx, addr, sdk.NewInt64Coin(sdk.DefaultBondDenom, 1000000000))

	caller := common.BytesToAddress(addr.Bytes())
	_, evm, err := input.EVMKeeper.CreateEVM(sdkCtx, caller, nil)
	require.NoError(t, err)
	stateDB := evm.StateDB.(*state.StateDB)

	// deploy contract
	counterBz, err := hexutil.Decode(counter.CounterBin)
	require.NoError(t, err)

	_, contractAddr, _, err := evm.Create(vm.AccountRef(caller), counterBz, 1_000_000, uint256.NewInt(0))
	require.NoError(t, err)

	// increase counter
	parsed, err := counter.CounterMetaData.GetAbi()
	require.NoError(t, err)

	inputBz, err := parsed.Pack("increase")
	require.NoError(t, err)

	// call with value
	res, logs, err := evm.Call(vm.AccountRef(caller), contractAddr, inputBz, 1_000_000, uint256.NewInt(100))
	require.NoError(t, err)
	require.Empty(t, res)
	require.NotEmpty(t, logs)

	// check balance
	require.Equal(t, uint64(100), stateDB.GetBalance(contractAddr).Uint64())

	// self destruct
	require.False(t, stateDB.HasSelfDestructed(contractAddr))
	stateDB.Selfdestruct6780(contractAddr)
	require.True(t, stateDB.HasSelfDestructed(contractAddr))

	// check balance
	require.Equal(t, uint64(0), stateDB.GetBalance(contractAddr).Uint64())

	// should clear only self destructed account
	err = stateDB.Commit()
	require.NoError(t, err)

	// should have empty storage root
	require.Equal(t, common.Hash{}, stateDB.GetStorageRoot(contractAddr))

	// should clear cosmos account
	require.Nil(t, input.AccountKeeper.GetAccount(sdkCtx, sdk.AccAddress(contractAddr.Bytes())))
}

func Test_Empty(t *testing.T) {
	ctx, input := createDefaultTestInput(t)

	_, _, addr := keyPubAddr()

	// get state db
	_, evm, err := input.EVMKeeper.CreateEVM(ctx, evmtypes.StdAddress, nil)
	require.NoError(t, err)

	stateDB := evm.StateDB.(*state.StateDB)
	stateDB.Snapshot()

	require.True(t, stateDB.Empty(common.BytesToAddress(addr.Bytes())))

	stateDB.SetNonce(common.BytesToAddress(addr.Bytes()), 1)
	require.False(t, stateDB.Empty(common.BytesToAddress(addr.Bytes())))

	stateDB.SetNonce(common.BytesToAddress(addr.Bytes()), 0)
	require.True(t, stateDB.Empty(common.BytesToAddress(addr.Bytes())))

	stateDB.SetCode(common.BytesToAddress(addr.Bytes()), []byte{1, 2, 3})
	require.False(t, stateDB.Empty(common.BytesToAddress(addr.Bytes())))

	// fund account
	input.Faucet.Fund(ctx, addr, sdk.NewInt64Coin(sdk.DefaultBondDenom, 1000000000))

	// get state db again
	_, evm, err = input.EVMKeeper.CreateEVM(ctx, evmtypes.StdAddress, nil)
	require.NoError(t, err)

	stateDB = evm.StateDB.(*state.StateDB)
	require.False(t, stateDB.Empty(common.BytesToAddress(addr.Bytes())))
}

func Test_ContractAddrConfilct_DueToCosmosAccount(t *testing.T) {
	ctx, input := createDefaultTestInput(t)

	_, _, addr := keyPubAddr()

	contractAddr, err := input.EVMKeeper.NextContractAddress(ctx, common.BytesToAddress(addr.Bytes()))
	require.NoError(t, err)

	// fund addr
	input.Faucet.Fund(ctx, contractAddr.Bytes(), sdk.NewInt64Coin(sdk.DefaultBondDenom, 1000000000))

	// set pubkey
	contractAcc := input.AccountKeeper.GetAccount(ctx, sdk.AccAddress(contractAddr.Bytes()))
	require.NoError(t, contractAcc.SetPubKey(&secp256k1.PubKey{}))
	input.AccountKeeper.SetAccount(ctx, contractAcc)

	// deploy contract
	counterBz, err := hexutil.Decode(counter.CounterBin)
	require.NoError(t, err)

	caller := common.BytesToAddress(addr.Bytes())
	_, _, _, err = input.EVMKeeper.EVMCreate(ctx, caller, counterBz, nil, nil)
	require.ErrorContains(t, err, vm.ErrContractAddressCollision.Error())
}

func Test_CreateContract_OverrideEmptyAccount(t *testing.T) {
	ctx, input := createDefaultTestInput(t)

	_, _, addr := keyPubAddr()

	contractAddr, err := input.EVMKeeper.NextContractAddress(ctx, common.BytesToAddress(addr.Bytes()))
	require.NoError(t, err)

	// fund addr
	input.Faucet.Fund(ctx, contractAddr.Bytes(), sdk.NewInt64Coin(sdk.DefaultBondDenom, 1000000000))

	// deploy contract
	counterBz, err := hexutil.Decode(counter.CounterBin)
	require.NoError(t, err)

	caller := common.BytesToAddress(addr.Bytes())
	_, _, _, err = input.EVMKeeper.EVMCreate(ctx, caller, counterBz, nil, nil)
	require.NoError(t, err)

	contractAcc := input.AccountKeeper.GetAccount(ctx, sdk.AccAddress(contractAddr.Bytes()))
	_, ok := contractAcc.(*evmtypes.ContractAccount)
	require.True(t, ok)
}
