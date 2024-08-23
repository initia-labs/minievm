package state

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"

	"github.com/holiman/uint256"

	"cosmossdk.io/collections"
	"cosmossdk.io/log"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"

	evmtypes "github.com/initia-labs/minievm/x/evm/types"
)

type callableEVM interface {
	Call(vm.ContractRef, common.Address, []byte, uint64, *uint256.Int) ([]byte, uint64, error)
	StaticCall(vm.ContractRef, common.Address, []byte, uint64) ([]byte, uint64, error)
}

var _ vm.StateDB = &StateDB{}

type StateDB struct {
	ctx    context.Context
	logger log.Logger

	vmStore               collections.Map[[]byte, []byte]
	transientVMStore      collections.Map[collections.Pair[uint64, []byte], []byte]
	transientCreated      collections.KeySet[collections.Pair[uint64, []byte]]
	transientSelfDestruct collections.KeySet[collections.Pair[uint64, []byte]]
	transientLogs         collections.Map[collections.Pair[uint64, uint64], evmtypes.Log]
	transientLogSize      collections.Map[uint64, uint64]
	transientAccessList   collections.KeySet[collections.Pair[uint64, []byte]]
	transientRefund       collections.Map[uint64, uint64]
	execIndex             uint64

	evm             callableEVM
	erc20ABI        *abi.ABI
	feeContractAddr common.Address // feeDenom contract address

	// Snapshot stack
	snaps []*Snapshot
}

func NewStateDB(
	ctx context.Context,
	logger log.Logger,
	// store params
	vmStore collections.Map[[]byte, []byte],
	transientVMStore collections.Map[collections.Pair[uint64, []byte], []byte],
	transientCreated collections.KeySet[collections.Pair[uint64, []byte]],
	transientSelfDestruct collections.KeySet[collections.Pair[uint64, []byte]],
	transientLogs collections.Map[collections.Pair[uint64, uint64], evmtypes.Log],
	transientLogSize collections.Map[uint64, uint64],
	transientAccessList collections.KeySet[collections.Pair[uint64, []byte]],
	transientRefund collections.Map[uint64, uint64],
	transientExecIndexStore collections.Sequence,
	// erc20 params
	evm callableEVM,
	erc20ABI *abi.ABI,
	feeContractAddr common.Address,
) (*StateDB, error) {
	execIndex, err := transientExecIndexStore.Next(ctx)
	if err != nil {
		return nil, err
	}
	err = transientLogSize.Set(ctx, execIndex, 0)
	if err != nil {
		return nil, err
	}
	err = transientRefund.Set(ctx, execIndex, 0)
	if err != nil {
		return nil, err
	}

	s := &StateDB{
		ctx:    ctx,
		logger: logger,

		vmStore:               vmStore,
		transientVMStore:      transientVMStore,
		transientCreated:      transientCreated,
		transientSelfDestruct: transientSelfDestruct,
		transientLogs:         transientLogs,
		transientLogSize:      transientLogSize,
		transientAccessList:   transientAccessList,
		transientRefund:       transientRefund,
		execIndex:             execIndex,

		evm:             evm,
		erc20ABI:        erc20ABI,
		feeContractAddr: feeContractAddr,
	}

	// take snapshot for the initial state
	s.Snapshot()

	return s, nil
}

// AddBalance mint coins to the recipient
func (s *StateDB) AddBalance(addr common.Address, amount *uint256.Int, _ tracing.BalanceChangeReason) {
	if amount.IsZero() {
		return
	}

	inputBz, err := s.erc20ABI.Pack("sudoMint", addr, amount.ToBig())
	if err != nil {
		panic(err)
	}

	_, _, err = s.evm.Call(vm.AccountRef(addr), s.feeContractAddr, inputBz, 100000, uint256.NewInt(0))
	if err != nil {
		s.logger.Warn("failed to mint token", "error", err)
		panic(err)
	}
}

// SubBalance burn coins from the account with addr
func (s *StateDB) SubBalance(addr common.Address, amount *uint256.Int, _ tracing.BalanceChangeReason) {
	if amount.IsZero() {
		return
	}

	inputBz, err := s.erc20ABI.Pack("sudoBurn", addr, amount.ToBig())
	if err != nil {
		panic(err)
	}

	_, _, err = s.evm.Call(vm.AccountRef(addr), s.feeContractAddr, inputBz, 100000, uint256.NewInt(0))
	if err != nil {
		s.logger.Warn("failed to burn token", "error", err)
		panic(err)
	}
}

// GetBalance returns the erc20 balance of the account with addr
func (s *StateDB) GetBalance(addr common.Address) *uint256.Int {
	inputBz, err := s.erc20ABI.Pack("balanceOf", addr)
	if err != nil {
		panic(err)
	}

	retBz, _, err := s.evm.StaticCall(vm.AccountRef(evmtypes.NullAddress), s.feeContractAddr, inputBz, 100000)
	if err != nil {
		s.logger.Warn("failed to check balance", "error", err)
		panic(err)
	}

	res, err := s.erc20ABI.Unpack("balanceOf", retBz)
	if err != nil {
		panic(err)
	}

	balance, ok := res[0].(*big.Int)
	if !ok {
		panic(fmt.Sprintf("failed to convert balance to *big.Int: %v", res[0]))
	}

	return uint256.MustFromBig(balance)
}

// AddRefund implements vm.StateDB.
func (s *StateDB) AddRefund(gas uint64) {
	refund, err := s.transientRefund.Get(s.ctx, s.execIndex)
	if err != nil {
		panic(err)
	}

	err = s.transientRefund.Set(s.ctx, s.execIndex, refund+gas)
	if err != nil {
		panic(err)
	}
}

// SubRefund implements vm.StateDB.
func (s *StateDB) SubRefund(gas uint64) {
	refund, err := s.transientRefund.Get(s.ctx, s.execIndex)
	if err != nil {
		panic(err)
	}

	if gas > refund {
		panic(fmt.Sprintf("Refund counter below zero (gas: %d > refund: %d)", gas, refund))
	}

	err = s.transientRefund.Set(s.ctx, s.execIndex, refund-gas)
	if err != nil {
		panic(err)
	}
}

// AddAddressToAccessList adds the given address to the access list
func (s *StateDB) AddAddressToAccessList(addr common.Address) {
	err := s.transientAccessList.Set(s.ctx, collections.Join(s.execIndex, addr.Bytes()))
	if err != nil {
		panic(err)
	}

}

// AddSlotToAccessList adds the given (address, slot)-tuple to the access list
func (s *StateDB) AddSlotToAccessList(addr common.Address, slot common.Hash) {
	if !s.AddressInAccessList(addr) {
		s.AddAddressToAccessList(addr)
	}

	err := s.transientAccessList.Set(s.ctx, collections.Join(s.execIndex, append(addr.Bytes(), slot[:]...)))
	if err != nil {
		panic(err)
	}
}

// AddressInAccessList returns true if the given address is in the access list
func (s *StateDB) AddressInAccessList(addr common.Address) bool {
	ok, err := s.transientAccessList.Has(s.ctx, collections.Join(s.execIndex, addr.Bytes()))
	if err != nil {
		panic(err)
	}

	return ok
}

// SlotInAccessList returns true if the given (address, slot)-tuple is in the access list
func (s *StateDB) SlotInAccessList(addr common.Address, slot common.Hash) (addressOk bool, slotOk bool) {
	ok, err := s.transientAccessList.Has(s.ctx, collections.Join(s.execIndex, addr.Bytes()))
	if err != nil {
		panic(err)
	} else if !ok {
		return false, false
	}

	ok, err = s.transientAccessList.Has(s.ctx, collections.Join(s.execIndex, append(addr.Bytes(), slot[:]...)))
	if err != nil {
		panic(err)
	}

	return true, ok
}

// CreateAccount set the nonce of the account with addr to 0
func (s *StateDB) CreateAccount(addr common.Address) {
	if err := s.vmStore.Set(s.ctx, accountKey(addr), EmptyStateAccount().Marshal()); err != nil {
		panic(err)
	}

	if err := s.transientCreated.Set(s.ctx, collections.Join(s.execIndex, addr.Bytes())); err != nil {
		panic(err)
	}
}

// Empty returns empty according to the EIP161 specification (balance = nonce = code = 0)
func (s *StateDB) Empty(addr common.Address) bool {
	accBz, err := s.vmStore.Get(s.ctx, accountKey(addr))
	if err != nil && errors.Is(err, collections.ErrNotFound) {
		return true
	} else if err != nil {
		panic(err)
	}

	// check if the account has non-zero nonce or code hash
	sa := EmptyStateAccount().Unmarshal(accBz)
	if !sa.IsEmpty() {
		return false
	}

	// check if the account has non-zero balance
	if balance := s.GetBalance(addr); balance.Sign() != 0 {
		return false
	}

	return true
}

// Exist reports whether the given account address exists in the state.
// Notably this also returns true for self-destructed accounts.
func (s *StateDB) Exist(addr common.Address) bool {
	ok, err := s.vmStore.Has(s.ctx, accountKey(addr))
	if err != nil {
		panic(err)
	}

	return ok
}

func (s *StateDB) getStateAccount(addr common.Address) *StateAccount {
	acc, err := s.vmStore.Get(s.ctx, accountKey(addr))
	if err != nil && errors.Is(err, collections.ErrNotFound) {
		return nil
	} else if err != nil {
		panic(err)
	}

	return EmptyStateAccount().Unmarshal(acc)
}

// GetCode returns the code of the account with addr
func (s *StateDB) GetCode(addr common.Address) []byte {
	sa := s.getStateAccount(addr)
	if sa == nil {
		return nil
	}

	code, err := s.vmStore.Get(s.ctx, codeKey(addr, sa.CodeHash))
	if err != nil && errors.Is(err, collections.ErrNotFound) {
		return nil
	} else if err != nil {
		panic(err)
	}

	return code
}

// SetCode store the code of the account with addr
func (s *StateDB) SetCode(addr common.Address, code []byte) {
	sa := s.getStateAccount(addr)
	if sa == nil {
		sa = EmptyStateAccount()
	}

	// set the code hash in the state account
	sa.CodeHash = crypto.Keccak256Hash(code).Bytes()
	if err := s.vmStore.Set(s.ctx, accountKey(addr), sa.Marshal()); err != nil {
		panic(err)
	}

	// set the code in the store
	if err := s.vmStore.Set(s.ctx, codeKey(addr, sa.CodeHash), code); err != nil {
		panic(err)
	}

	// set the code size in the store
	if err := s.vmStore.Set(s.ctx, codeSizeKey(addr, sa.CodeHash), uint64ToBytes(uint64(len(code)))); err != nil {
		panic(err)
	}
}

// GetCodeHash returns the code hash of the account with addr
func (s *StateDB) GetCodeHash(addr common.Address) common.Hash {
	sa := s.getStateAccount(addr)
	if sa == nil {
		return common.Hash{}
	}

	return common.BytesToHash(sa.CodeHash)
}

// GetCodeSize returns the code size of the account with addr
func (s *StateDB) GetCodeSize(addr common.Address) int {
	sa := s.getStateAccount(addr)
	if sa == nil {
		return 0
	}

	codeSize, err := s.vmStore.Get(s.ctx, codeSizeKey(addr, sa.CodeHash))
	if err != nil && errors.Is(err, collections.ErrNotFound) {
		return 0
	} else if err != nil {
		panic(err)
	}

	return int(binary.BigEndian.Uint64(codeSize))
}

// GetCommittedState returns the committed state of the account with addr
func (s *StateDB) GetCommittedState(addr common.Address, state common.Hash) common.Hash {
	originCtx := s.ctx

	// use initial context to get the committed state
	s.ctx = s.snaps[0].ctx
	defer func() { s.ctx = originCtx }()

	return s.GetState(addr, state)
}

// GetNonce returns the nonce of the account with addr
func (s *StateDB) GetNonce(addr common.Address) uint64 {
	sa := s.getStateAccount(addr)
	if sa != nil {
		return sa.Nonce
	}

	return 0
}

// SetNonce sets the nonce of the account with addr
func (s *StateDB) SetNonce(addr common.Address, nonce uint64) {
	sa := s.getStateAccount(addr)
	if sa == nil {
		sa = EmptyStateAccount()
	}

	sa.Nonce = nonce
	if err := s.vmStore.Set(s.ctx, accountKey(addr), sa.Marshal()); err != nil {
		panic(err)
	}
}

// GetRefund returns the refund
func (s *StateDB) GetRefund() uint64 {
	refund, err := s.transientRefund.Get(s.ctx, s.execIndex)
	if err != nil {
		panic(err)
	}

	return refund
}

// GetState returns the state of the account with addr and slot
func (s *StateDB) GetState(addr common.Address, slot common.Hash) common.Hash {
	sa := s.getStateAccount(addr)
	if sa != nil {
		state, err := s.vmStore.Get(s.ctx, stateKey(addr, slot))
		if err != nil && errors.Is(err, collections.ErrNotFound) {
			return common.Hash{}
		} else if err != nil {
			panic(err)
		}

		return common.BytesToHash(state)
	}

	return common.Hash{}
}

// HasSelfDestructed return true if the account with addr has self-destructed
func (s *StateDB) HasSelfDestructed(addr common.Address) bool {
	sa := s.getStateAccount(addr)
	if sa != nil {
		ok, err := s.transientSelfDestruct.Has(s.ctx, collections.Join(s.execIndex, addr.Bytes()))
		if err != nil {
			panic(err)
		}

		return ok
	}

	return false
}

// SelfDestruct marks the given account as selfdestructed.
// This clears the account balance.
//
// The account's state object is still available until the state is committed,
// getStateObject will return a non-nil account after SelfDestruct.
func (s *StateDB) SelfDestruct(addr common.Address) {
	sa := s.getStateAccount(addr)
	if sa == nil {
		return
	}

	// mark the account as self-destructed
	if err := s.transientSelfDestruct.Set(s.ctx, collections.Join(s.execIndex, addr.Bytes())); err != nil {
		panic(err)
	}

	// clear the balance of the account
	s.SubBalance(addr, s.GetBalance(addr), tracing.BalanceDecreaseSelfdestructBurn)
}

// Selfdestruct6780 calls selfdestruct and clears the account balance if the account is created in the same transaction.
func (s *StateDB) Selfdestruct6780(addr common.Address) {
	sa := s.getStateAccount(addr)
	if sa == nil {
		return
	}

	ok, err := s.transientCreated.Has(s.ctx, collections.Join(s.execIndex, addr.Bytes()))
	if err != nil {
		panic(err)
	}
	if ok {
		s.SelfDestruct(addr)
	}
}

// SetState implements vm.StateDB.
func (s *StateDB) SetState(addr common.Address, slot common.Hash, value common.Hash) {
	sa := s.getStateAccount(addr)
	if sa == nil {
		sa = EmptyStateAccount()
		s.vmStore.Set(s.ctx, accountKey(addr), sa.Marshal())
	}

	if err := s.vmStore.Set(s.ctx, stateKey(addr, slot), value[:]); err != nil {
		panic(err)
	}
}

// SetTransientState sets transient storage for a given account.
func (s *StateDB) SetTransientState(addr common.Address, key, value common.Hash) {
	prev := s.GetTransientState(addr, key)
	if prev == value {
		return
	}

	s.transientVMStore.Set(s.ctx, collections.Join(s.execIndex, key[:]), value[:])
}

// GetTransientState gets transient storage for a given account.
func (s *StateDB) GetTransientState(addr common.Address, key common.Hash) common.Hash {
	data, err := s.transientVMStore.Get(s.ctx, collections.Join(s.execIndex, key[:]))
	if err != nil && errors.Is(err, collections.ErrNotFound) {
		return common.Hash{}
	} else if err != nil {
		panic(err)
	}

	return common.BytesToHash(data)
}

// Snapshot creates new snapshot(cache context) and return the snapshot id
func (s *StateDB) Snapshot() int {
	snap := NewSnapshot(s.ctx)
	s.snaps = append(s.snaps, snap)
	s.ctx = snap.ctx
	return len(s.snaps) - 2
}

// RevertToSnapshot reverts the state to the snapshot with the given id
func (s *StateDB) RevertToSnapshot(i int) {
	snap := s.snaps[i]
	s.ctx = snap.ctx
	s.snaps = s.snaps[:i]
}

// Prepare handles the preparatory steps for executing a state transition with.
// This method must be invoked before state transition.
//
// Berlin fork:
// - Add sender to access list (2929)
// - Add destination to access list (2929)
// - Add precompiles to access list (2929)
// - Add the contents of the optional tx access list (2930)
//
// Potential EIPs:
// - Reset access list (Berlin)
// - Add coinbase to access list (EIP-3651)
// - Reset transient storage (EIP-1153)
func (s *StateDB) Prepare(rules params.Rules, sender common.Address, coinbase common.Address, dst *common.Address, precompiles []common.Address, list types.AccessList) {
	if rules.IsBerlin {
		// Clear out any leftover from previous executions
		s.AddAddressToAccessList(sender)

		if dst != nil {
			s.AddAddressToAccessList(*dst)
			// If it's a create-tx, the destination will be added inside evm.create
		}
		for _, addr := range precompiles {
			s.AddAddressToAccessList(addr)
		}
		for _, el := range list {
			s.AddAddressToAccessList(el.Address)
			for _, key := range el.StorageKeys {
				s.AddSlotToAccessList(el.Address, key)
			}
		}
		if rules.IsShanghai { // EIP-3651: warm coinbase
			s.AddAddressToAccessList(coinbase)
		}
	}

	// take snapshot for the committed state
	s.Snapshot()
}

func (s *StateDB) Commit() error {
	// commit all changes
	for i := range s.snaps {
		s.snaps[len(s.snaps)-i-1].Commit()
	}

	// use the initial context
	s.ctx = s.snaps[0].ctx

	// clear destructed accounts
	err := s.transientSelfDestruct.Walk(s.ctx, collections.NewPrefixedPairRange[uint64, []byte](s.execIndex), func(key collections.Pair[uint64, []byte]) (stop bool, err error) {
		err = s.vmStore.Clear(s.ctx, new(collections.Range[[]byte]).Prefix(key.K2()))
		return false, err
	})
	if err != nil {
		return err
	}

	return nil
}

// AddLog implements vm.StateDB.
func (s *StateDB) AddLog(log *types.Log) {
	logSize, err := s.transientLogSize.Get(s.ctx, s.execIndex)
	if err != nil {
		panic(err)
	}

	err = s.transientLogSize.Set(s.ctx, s.execIndex, logSize+1)
	if err != nil {
		panic(err)
	}

	err = s.transientLogs.Set(s.ctx, collections.Join(s.execIndex, logSize), evmtypes.NewLog(log))
	if err != nil {
		panic(err)
	}
}

func (s *StateDB) Logs() []evmtypes.Log {
	logSize, err := s.transientLogSize.Get(s.ctx, s.execIndex)
	if err != nil {
		panic(err)
	} else if logSize == 0 {
		return []evmtypes.Log{}
	}

	logs := make([]evmtypes.Log, logSize)
	err = s.transientLogs.Walk(s.ctx, collections.NewPrefixedPairRange[uint64, uint64](s.execIndex), func(key collections.Pair[uint64, uint64], log evmtypes.Log) (stop bool, err error) {
		logs[key.K2()] = log
		return false, nil
	})
	if err != nil {
		panic(err)
	}

	return logs
}

// AddPreimage implements vm.StateDB.
func (s *StateDB) AddPreimage(common.Hash, []byte) {
	// no-op
}
