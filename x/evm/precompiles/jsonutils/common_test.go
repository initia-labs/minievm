package jsonutils_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/stateless"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/trie/utils"
	"github.com/holiman/uint256"

	"github.com/initia-labs/minievm/x/evm/state"
	evmtypes "github.com/initia-labs/minievm/x/evm/types"
)

var _ evmtypes.StateDB = &MockStateDB{}

type MockStateDB struct {
	ctx        state.Context
	initialCtx state.Context

	// Snapshot stack
	snaps []*state.Snapshot

	evm *vm.EVM
}

func NewMockStateDB(sdkCtx sdk.Context) *MockStateDB {
	ctx := state.NewContext(sdkCtx)
	return &MockStateDB{
		ctx:        ctx,
		initialCtx: ctx,

		evm: &vm.EVM{},
	}
}

// Snapshot implements types.StateDB.
func (m *MockStateDB) Snapshot() int {
	// get a current snapshot id
	sid := len(m.snaps) - 1

	// create a new snapshot
	snap := state.NewSnapshot(m.ctx)
	m.snaps = append(m.snaps, snap)

	// use the new snapshot context
	m.ctx = snap.Context()

	// return the current snapshot id
	return sid
}

// RevertToSnapshot implements types.StateDB.
func (m *MockStateDB) RevertToSnapshot(i int) {
	if i == -1 {
		m.ctx = m.initialCtx
		m.snaps = m.snaps[:0]
		return
	}

	// revert to the snapshot with the given id
	snap := m.snaps[i]
	m.ctx = snap.Context()

	// clear the snapshots after the given id
	m.snaps = m.snaps[:i+1]
}

// Context implements types.StateDB.
func (m *MockStateDB) Context() sdk.Context {
	return m.ctx.Context
}

// EVM implements types.StateDB.
func (m *MockStateDB) EVM() *vm.EVM {
	return m.evm
}

//////////////////////// MOCKED METHODS ////////////////////////

// AddAddressToAccessList implements types.StateDB.
func (m *MockStateDB) AddAddressToAccessList(addr common.Address) {
	panic("unimplemented")
}

// AddBalance implements types.StateDB.
func (m *MockStateDB) AddBalance(common.Address, *uint256.Int, tracing.BalanceChangeReason) {
	panic("unimplemented")
}

// AddLog implements types.StateDB.
func (m *MockStateDB) AddLog(*types.Log) {
	panic("unimplemented")
}

// AddPreimage implements types.StateDB.
func (m *MockStateDB) AddPreimage(common.Hash, []byte) {
	panic("unimplemented")
}

// AddRefund implements types.StateDB.
func (m *MockStateDB) AddRefund(uint64) {
	panic("unimplemented")
}

// AddSlotToAccessList implements types.StateDB.
func (m *MockStateDB) AddSlotToAccessList(addr common.Address, slot common.Hash) {
	panic("unimplemented")
}

// AddressInAccessList implements types.StateDB.
func (m *MockStateDB) AddressInAccessList(addr common.Address) bool {
	panic("unimplemented")
}

// CreateAccount implements types.StateDB.
func (m *MockStateDB) CreateAccount(common.Address) {
	panic("unimplemented")
}

// CreateContract implements types.StateDB.
func (m *MockStateDB) CreateContract(common.Address) {
	panic("unimplemented")
}

// Empty implements types.StateDB.
func (m *MockStateDB) Empty(common.Address) bool {
	panic("unimplemented")
}

// Exist implements types.StateDB.
func (m *MockStateDB) Exist(common.Address) bool {
	panic("unimplemented")
}

// GetBalance implements types.StateDB.
func (m *MockStateDB) GetBalance(common.Address) *uint256.Int {
	panic("unimplemented")
}

// GetCode implements types.StateDB.
func (m *MockStateDB) GetCode(common.Address) []byte {
	panic("unimplemented")
}

// GetCodeHash implements types.StateDB.
func (m *MockStateDB) GetCodeHash(common.Address) common.Hash {
	panic("unimplemented")
}

// GetCodeSize implements types.StateDB.
func (m *MockStateDB) GetCodeSize(common.Address) int {
	panic("unimplemented")
}

// GetCommittedState implements types.StateDB.
func (m *MockStateDB) GetCommittedState(common.Address, common.Hash) common.Hash {
	panic("unimplemented")
}

// GetNonce implements types.StateDB.
func (m *MockStateDB) GetNonce(common.Address) uint64 {
	panic("unimplemented")
}

// GetRefund implements types.StateDB.
func (m *MockStateDB) GetRefund() uint64 {
	panic("unimplemented")
}

// GetState implements types.StateDB.
func (m *MockStateDB) GetState(common.Address, common.Hash) common.Hash {
	panic("unimplemented")
}

// GetStorageRoot implements types.StateDB.
func (m *MockStateDB) GetStorageRoot(addr common.Address) common.Hash {
	panic("unimplemented")
}

// GetTransientState implements types.StateDB.
func (m *MockStateDB) GetTransientState(addr common.Address, key common.Hash) common.Hash {
	panic("unimplemented")
}

// HasSelfDestructed implements types.StateDB.
func (m *MockStateDB) HasSelfDestructed(common.Address) bool {
	panic("unimplemented")
}

// PointCache implements types.StateDB.
func (m *MockStateDB) PointCache() *utils.PointCache {
	panic("unimplemented")
}

// Prepare implements types.StateDB.
func (m *MockStateDB) Prepare(rules params.Rules, sender common.Address, coinbase common.Address, dest *common.Address, precompiles []common.Address, txAccesses types.AccessList) {
	panic("unimplemented")
}

// SelfDestruct implements types.StateDB.
func (m *MockStateDB) SelfDestruct(common.Address) {
	panic("unimplemented")
}

// Selfdestruct6780 implements types.StateDB.
func (m *MockStateDB) Selfdestruct6780(common.Address) {
	panic("unimplemented")
}

// SetCode implements types.StateDB.
func (m *MockStateDB) SetCode(common.Address, []byte) {
	panic("unimplemented")
}

// SetNonce implements types.StateDB.
func (m *MockStateDB) SetNonce(common.Address, uint64) {
	panic("unimplemented")
}

// SetState implements types.StateDB.
func (m *MockStateDB) SetState(common.Address, common.Hash, common.Hash) {
	panic("unimplemented")
}

// SetTransientState implements types.StateDB.
func (m *MockStateDB) SetTransientState(addr common.Address, key common.Hash, value common.Hash) {
	panic("unimplemented")
}

// SlotInAccessList implements types.StateDB.
func (m *MockStateDB) SlotInAccessList(addr common.Address, slot common.Hash) (addressOk bool, slotOk bool) {
	panic("unimplemented")
}

// SubBalance implements types.StateDB.
func (m *MockStateDB) SubBalance(common.Address, *uint256.Int, tracing.BalanceChangeReason) {
	panic("unimplemented")
}

// SubRefund implements types.StateDB.
func (m *MockStateDB) SubRefund(uint64) {
	panic("unimplemented")
}

// Witness implements types.StateDB.
func (m *MockStateDB) Witness() *stateless.Witness {
	panic("unimplemented")
}
