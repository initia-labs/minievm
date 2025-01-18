package state

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/holiman/uint256"

	"cosmossdk.io/collections"
	corestoretypes "cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/stateless"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/trie/utils"

	evmtypes "github.com/initia-labs/minievm/x/evm/types"
)

var _ vm.StateDB = &StateDB{}

type StateDB struct {
	ctx           Context
	initialCtx    Context
	logger        log.Logger
	accountKeeper evmtypes.AccountKeeper

	vmStore collections.Map[[]byte, []byte]

	// transient memory store for the current execution
	memStoreVMStore      collections.Map[[]byte, []byte]
	memStoreCreated      collections.KeySet[[]byte]
	memStoreSelfDestruct collections.KeySet[[]byte]
	memStoreLogs         collections.Map[uint64, evmtypes.Log]
	memStoreLogSize      collections.Item[uint64]
	memStoreAccessList   collections.KeySet[[]byte]
	memStoreRefund       collections.Item[uint64]
	schema               collections.Schema

	evm             *vm.EVM
	erc20ABI        *abi.ABI
	feeContractAddr common.Address // feeDenom contract address

	// Snapshot stack
	snaps []*Snapshot
}

const (
	erc20OpGasLimit = 100000
)

func NewStateDB(
	sdkCtx sdk.Context,
	cdc codec.Codec,
	logger log.Logger,
	accountKeeper evmtypes.AccountKeeper,
	// store params
	vmStore collections.Map[[]byte, []byte],
	// erc20 params
	evm *vm.EVM,
	erc20ABI *abi.ABI,
	feeContractAddr common.Address,
) (*StateDB, error) {
	sb := collections.NewSchemaBuilderFromAccessor(
		func(ctx context.Context) corestoretypes.KVStore {
			stateCtx := ctx.(Context)
			return newKVStore(stateCtx.memStore.GetKVStore(stateCtx.memStoreKey))
		},
	)

	ctx := NewContext(sdkCtx)
	s := &StateDB{
		ctx:           ctx,
		initialCtx:    ctx,
		logger:        logger,
		accountKeeper: accountKeeper,

		vmStore: vmStore,

		memStoreVMStore:      collections.NewMap(sb, memStoreVMStorePrefix, "mem_store_vm_store", collections.BytesKey, collections.BytesValue),
		memStoreCreated:      collections.NewKeySet(sb, memStoreCreatedPrefix, "mem_store_created", collections.BytesKey),
		memStoreSelfDestruct: collections.NewKeySet(sb, memStoreSelfDestructPrefix, "mem_store_self_destruct", collections.BytesKey),
		memStoreLogs:         collections.NewMap(sb, memStoreLogsPrefix, "mem_store_logs", collections.Uint64Key, codec.CollValue[evmtypes.Log](cdc)),
		memStoreLogSize:      collections.NewItem(sb, memStoreLogSizePrefix, "mem_store_log_size", collections.Uint64Value),
		memStoreAccessList:   collections.NewKeySet(sb, memStoreAccessListPrefix, "mem_store_access_list", collections.BytesKey),
		memStoreRefund:       collections.NewItem(sb, memStoreRefundPrefix, "mem_store_refund", collections.Uint64Value),

		evm:             evm,
		erc20ABI:        erc20ABI,
		feeContractAddr: feeContractAddr,
	}
	schema, err := sb.Build()
	if err != nil {
		return nil, err
	}
	s.schema = schema

	err = s.memStoreLogSize.Set(ctx, 0)
	if err != nil {
		return nil, err
	}
	err = s.memStoreRefund.Set(ctx, 0)
	if err != nil {
		return nil, err
	}

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

	_, _, err = s.evm.Call(vm.AccountRef(evmtypes.StdAddress), s.feeContractAddr, inputBz, erc20OpGasLimit, uint256.NewInt(0))
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

	_, _, err = s.evm.Call(vm.AccountRef(evmtypes.StdAddress), s.feeContractAddr, inputBz, erc20OpGasLimit, uint256.NewInt(0))
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

	retBz, _, err := s.evm.StaticCall(vm.AccountRef(evmtypes.NullAddress), s.feeContractAddr, inputBz, erc20OpGasLimit)
	if err != nil {
		s.logger.Warn("failed to check balance", "error", err)
		panic(err)
	}
	if len(retBz) == 0 {
		return uint256.NewInt(0)
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
	refund, err := s.memStoreRefund.Get(s.ctx)
	if err != nil {
		panic(err)
	}

	err = s.memStoreRefund.Set(s.ctx, refund+gas)
	if err != nil {
		panic(err)
	}
}

// SubRefund implements vm.StateDB.
func (s *StateDB) SubRefund(gas uint64) {
	refund, err := s.memStoreRefund.Get(s.ctx)
	if err != nil {
		panic(err)
	}

	if gas > refund {
		panic(fmt.Sprintf("Refund counter below zero (gas: %d > refund: %d)", gas, refund))
	}

	err = s.memStoreRefund.Set(s.ctx, refund-gas)
	if err != nil {
		panic(err)
	}
}

// AddAddressToAccessList adds the given address to the access list
func (s *StateDB) AddAddressToAccessList(addr common.Address) {
	err := s.memStoreAccessList.Set(s.ctx, addr.Bytes())
	if err != nil {
		panic(err)
	}

}

// AddSlotToAccessList adds the given (address, slot)-tuple to the access list
func (s *StateDB) AddSlotToAccessList(addr common.Address, slot common.Hash) {
	if !s.AddressInAccessList(addr) {
		s.AddAddressToAccessList(addr)
	}

	err := s.memStoreAccessList.Set(s.ctx, append(addr.Bytes(), slot[:]...))
	if err != nil {
		panic(err)
	}
}

// AddressInAccessList returns true if the given address is in the access list
func (s *StateDB) AddressInAccessList(addr common.Address) bool {
	ok, err := s.memStoreAccessList.Has(s.ctx, addr.Bytes())
	if err != nil {
		panic(err)
	}

	return ok
}

// SlotInAccessList returns true if the given (address, slot)-tuple is in the access list
func (s *StateDB) SlotInAccessList(addr common.Address, slot common.Hash) (addressOk bool, slotOk bool) {
	ok, err := s.memStoreAccessList.Has(s.ctx, addr.Bytes())
	if err != nil {
		panic(err)
	} else if !ok {
		return false, false
	}

	ok, err = s.memStoreAccessList.Has(s.ctx, append(addr.Bytes(), slot[:]...))
	if err != nil {
		panic(err)
	}

	return true, ok
}

// CreateAccount set the nonce of the account with addr to 0
func (s *StateDB) CreateAccount(addr common.Address) {
	acc := s.accountKeeper.NewAccountWithAddress(s.ctx, addr.Bytes())
	s.accountKeeper.SetAccount(s.ctx, acc)
}

// CreateContract creates a contract account with the given address
func (s *StateDB) CreateContract(contractAddr common.Address) {
	if err := s.memStoreCreated.Set(s.ctx, contractAddr.Bytes()); err != nil {
		panic(err)
	}

	// If the account is empty, converts a normal account to a contract account
	// Else, creates a contract account if the account does not exist.
	if s.accountKeeper.HasAccount(s.ctx, sdk.AccAddress(contractAddr.Bytes())) {
		acc := s.accountKeeper.GetAccount(s.ctx, sdk.AccAddress(contractAddr.Bytes()))

		// check the account is empty or not
		if !evmtypes.IsEmptyAccount(acc) {
			panic(evmtypes.ErrAddressAlreadyExists.Wrap(contractAddr.String()))
		}

		// convert base account to contract account only if this account is empty
		contractAcc := evmtypes.NewContractAccountWithAddress(contractAddr.Bytes())
		contractAcc.AccountNumber = acc.GetAccountNumber()
		s.accountKeeper.SetAccount(s.ctx, contractAcc)
	} else {
		// create contract account
		contractAcc := evmtypes.NewContractAccountWithAddress(contractAddr.Bytes())
		contractAcc.AccountNumber = s.accountKeeper.NextAccountNumber(s.ctx)
		s.accountKeeper.SetAccount(s.ctx, contractAcc)
	}

	// emit cosmos contract created event
	s.ctx.EventManager().EmitEvent(sdk.NewEvent(
		evmtypes.EventTypeContractCreated,
		sdk.NewAttribute(evmtypes.AttributeKeyContract, contractAddr.Hex()),
	))
}

func (s *StateDB) getAccount(addr common.Address) sdk.AccountI {
	return s.accountKeeper.GetAccount(s.ctx, sdk.AccAddress(addr.Bytes()))
}

func (s *StateDB) getOrNewAccount(addr common.Address) sdk.AccountI {
	acc := s.accountKeeper.GetAccount(s.ctx, sdk.AccAddress(addr.Bytes()))
	if acc == nil {
		acc = s.accountKeeper.NewAccountWithAddress(s.ctx, sdk.AccAddress(addr.Bytes()))
	}

	return acc
}

// Empty returns empty according to the EIP161 specification (balance = nonce = code = 0)
func (s *StateDB) Empty(addr common.Address) bool {
	// check if the account has non-zero balance
	if balance := s.GetBalance(addr); balance.Sign() != 0 {
		return false
	}

	acc := s.getAccount(addr)
	return acc == nil || evmtypes.IsEmptyAccount(acc)
}

// Exist reports whether the given account address exists in the state.
// Notably this also returns true for self-destructed accounts.
func (s *StateDB) Exist(addr common.Address) bool {
	acc := s.getAccount(addr)
	return acc != nil
}

// GetCode returns the code of the account with addr
func (s *StateDB) GetCode(addr common.Address) []byte {
	acc := s.getAccount(addr)
	if acc == nil {
		return nil
	}

	cacc, ok := acc.(*evmtypes.ContractAccount)
	if !ok {
		return nil
	}

	code, err := s.vmStore.Get(s.ctx, codeKey(addr, cacc.CodeHash))
	if err != nil && errors.Is(err, collections.ErrNotFound) {
		return nil
	} else if err != nil {
		panic(err)
	}

	return code
}

// SetCode store the code of the account with addr, and set the code hash to the account
// It is always used in conjunction with CreateContract, so don't need to check account conversion.
func (s *StateDB) SetCode(addr common.Address, code []byte) {
	ca := s.getOrNewAccount(addr)
	if evmtypes.IsEmptyAccount(ca) {
		an := ca.GetAccountNumber()
		ca = evmtypes.NewContractAccountWithAddress(addr.Bytes())
		if err := ca.SetAccountNumber(an); err != nil {
			panic(err)
		}
	}

	codeHash := crypto.Keccak256Hash(code).Bytes()
	ca.(*evmtypes.ContractAccount).CodeHash = codeHash
	s.accountKeeper.SetAccount(s.ctx, ca)

	// set the code in the store
	if err := s.vmStore.Set(s.ctx, codeKey(addr, codeHash), code); err != nil {
		panic(err)
	}

	// set the code size in the store
	if err := s.vmStore.Set(s.ctx, codeSizeKey(addr, codeHash), uint64ToBytes(uint64(len(code)))); err != nil {
		panic(err)
	}
}

// GetCodeHash returns the code hash of the account with addr
func (s *StateDB) GetCodeHash(addr common.Address) common.Hash {
	acc := s.getAccount(addr)
	if acc == nil {
		return common.Hash{}
	}

	cacc, ok := acc.(*evmtypes.ContractAccount)
	if !ok {
		return types.EmptyCodeHash
	}

	return common.BytesToHash(cacc.CodeHash)
}

// GetCodeSize returns the code size of the account with addr
func (s *StateDB) GetCodeSize(addr common.Address) int {
	acc := s.getAccount(addr)
	if acc == nil {
		return 0
	}

	cacc, ok := acc.(*evmtypes.ContractAccount)
	if !ok {
		return 0
	}

	codeSize, err := s.vmStore.Get(s.ctx, codeSizeKey(addr, cacc.CodeHash))
	if err != nil && errors.Is(err, collections.ErrNotFound) {
		return 0
	} else if err != nil {
		panic(err)
	}

	return int(bytesToUint64(codeSize))
}

// GetCommittedState returns the committed state of the account with addr
func (s *StateDB) GetCommittedState(addr common.Address, state common.Hash) common.Hash {
	originCtx := s.ctx

	// use initial context to get the committed state
	s.ctx = s.initialCtx
	defer func() { s.ctx = originCtx }()

	return s.GetState(addr, state)
}

// GetNonce returns the nonce of the account with addr
func (s *StateDB) GetNonce(addr common.Address) uint64 {
	acc := s.getAccount(addr)
	if acc == nil {
		return 0
	}

	return acc.GetSequence()
}

// SetNonce sets the nonce of the account with addr
func (s *StateDB) SetNonce(addr common.Address, nonce uint64) {
	acc := s.getOrNewAccount(addr)
	if err := acc.SetSequence(nonce); err != nil {
		panic(err)
	}
	s.accountKeeper.SetAccount(s.ctx, acc)
}

// GetRefund returns the refund
func (s *StateDB) GetRefund() uint64 {
	refund, err := s.memStoreRefund.Get(s.ctx)
	if err != nil {
		panic(err)
	}

	return refund
}

// GetState returns the state of the account with addr and slot
func (s *StateDB) GetState(addr common.Address, slot common.Hash) common.Hash {
	acc := s.getAccount(addr)
	if acc != nil {
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
	acc := s.getAccount(addr)
	if acc != nil {
		ok, err := s.memStoreSelfDestruct.Has(s.ctx, addr.Bytes())
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
	acc := s.getAccount(addr)
	if acc == nil {
		return
	}

	// mark the account as self-destructed
	if err := s.memStoreSelfDestruct.Set(s.ctx, addr.Bytes()); err != nil {
		panic(err)
	}

	// clear the balance of the account
	s.SubBalance(addr, s.GetBalance(addr), tracing.BalanceDecreaseSelfdestructBurn)
}

// Selfdestruct6780 calls selfdestruct and clears the account balance if the account is created in the same transaction.
func (s *StateDB) Selfdestruct6780(addr common.Address) {
	acc := s.getAccount(addr)
	if acc == nil {
		return
	}

	ok, err := s.memStoreCreated.Has(s.ctx, addr.Bytes())
	if err != nil {
		panic(err)
	} else if ok {
		s.SelfDestruct(addr)
	}
}

// SetState implements vm.StateDB.
func (s *StateDB) SetState(addr common.Address, slot common.Hash, value common.Hash) {
	if err := s.vmStore.Set(s.ctx, stateKey(addr, slot), value[:]); err != nil {
		panic(err)
	}
}

// SetTransientState sets memStore storage for a given account.
func (s *StateDB) SetTransientState(addr common.Address, key, value common.Hash) {
	prev := s.GetTransientState(addr, key)
	if prev == value {
		return
	}
	if err := s.memStoreVMStore.Set(s.ctx, append(addr.Bytes(), key.Bytes()...), value[:]); err != nil {
		panic(err)
	}
}

// GetTransientState gets memStore storage for a given account.
func (s *StateDB) GetTransientState(addr common.Address, key common.Hash) common.Hash {
	data, err := s.memStoreVMStore.Get(s.ctx, append(addr.Bytes(), key.Bytes()...))
	if err != nil && errors.Is(err, collections.ErrNotFound) {
		return common.Hash{}
	} else if err != nil {
		panic(err)
	}

	return common.BytesToHash(data)
}

// Snapshot creates new snapshot(cache context) and return the snapshot id
func (s *StateDB) Snapshot() int {
	// get a current snapshot id
	sid := len(s.snaps) - 1

	// create a new snapshot
	snap := NewSnapshot(s.ctx)
	s.snaps = append(s.snaps, snap)

	// use the new snapshot context
	s.ctx = snap.ctx

	// return the current snapshot id
	return sid
}

// RevertToSnapshot reverts the state to the snapshot with the given id
func (s *StateDB) RevertToSnapshot(i int) {
	if i == -1 {
		s.ctx = s.initialCtx
		s.snaps = s.snaps[:0]
		return
	}

	// revert to the snapshot with the given id
	snap := s.snaps[i]
	s.ctx = snap.ctx

	// clear the snapshots after the given id
	s.snaps = s.snaps[:i+1]
}

// ContextOfSnapshot returns the context of the snapshot with the given id
func (s *StateDB) ContextOfSnapshot(i int) sdk.Context {
	if i == -1 {
		return s.initialCtx.Context
	}

	return s.snaps[i].ctx.Context
}

// Context returns the current context
func (s *StateDB) Context() sdk.Context {
	return s.ctx.Context
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
}

func (s *StateDB) Commit() error {
	// clear destructed accounts
	err := s.memStoreSelfDestruct.Walk(s.ctx, nil, func(key []byte) (stop bool, err error) {
		addr := common.BytesToAddress(key)

		// If ether was sent to account post-selfdestruct it is burnt.
		if bal := s.GetBalance(addr); bal.Sign() != 0 {
			s.SubBalance(addr, bal, tracing.BalanceDecreaseSelfdestructBurn)
		}

		err = s.vmStore.Clear(s.ctx, new(collections.Range[[]byte]).Prefix(addr.Bytes()))

		// remove cosmos account
		s.accountKeeper.RemoveAccount(s.ctx, s.accountKeeper.GetAccount(s.ctx, sdk.AccAddress(addr.Bytes())))
		return false, err
	})
	if err != nil {
		return err
	}

	// commit all changes
	for i := range s.snaps {
		s.snaps[len(s.snaps)-i-1].Commit()
	}

	// use the initial context
	s.ctx = s.initialCtx

	return nil
}

// AddLog implements vm.StateDB.
func (s *StateDB) AddLog(log *types.Log) {
	logSize, err := s.memStoreLogSize.Get(s.ctx)
	if err != nil {
		panic(err)
	}

	err = s.memStoreLogSize.Set(s.ctx, logSize+1)
	if err != nil {
		panic(err)
	}

	err = s.memStoreLogs.Set(s.ctx, logSize, evmtypes.NewLog(log))
	if err != nil {
		panic(err)
	}
}

func (s *StateDB) Logs() evmtypes.Logs {
	logSize, err := s.memStoreLogSize.Get(s.ctx)
	if err != nil {
		panic(err)
	} else if logSize == 0 {
		return []evmtypes.Log{}
	}

	logs := make([]evmtypes.Log, logSize)
	err = s.memStoreLogs.Walk(s.ctx, nil, func(key uint64, log evmtypes.Log) (stop bool, err error) {
		logs[key] = log
		return false, nil
	})
	if err != nil {
		panic(err)
	}

	return logs
}

// GetStorageRoot return non-empty storage root if the account with addr has non-empty storage
// or there is non-empty cosmos account.
func (s *StateDB) GetStorageRoot(addr common.Address) common.Hash {
	nonEmptyHash := common.Hash{1}

	// check whether the non-empty account exists in the account keeper
	if s.accountKeeper.HasAccount(s.ctx, sdk.AccAddress(addr.Bytes())) {
		account := s.accountKeeper.GetAccount(s.ctx, sdk.AccAddress(addr.Bytes()))

		// check the account is empty or not
		if !evmtypes.IsEmptyAccount(account) {
			return nonEmptyHash
		}
	}

	// check whether the non-empty storage exists in the vm store
	iter, err := s.vmStore.Iterate(s.ctx, new(collections.Range[[]byte]).Prefix(stateKeyPrefix(addr)))
	if err != nil {
		panic(err)
	} else if iter.Valid() {
		return nonEmptyHash
	}

	// return empty storage root
	return common.Hash{}
}

// unused in the current implementation
func (s *StateDB) PointCache() *utils.PointCache {
	return nil
}

// unused in the current implementation
func (s *StateDB) Witness() *stateless.Witness {
	return nil
}

// unused in the current implementation
func (s *StateDB) AddPreimage(common.Hash, []byte) {
	// no-op
}
