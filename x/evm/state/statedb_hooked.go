package state

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/tracing"
	evmtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/holiman/uint256"
	"github.com/initia-labs/minievm/x/evm/types"
)

var _ types.StateDB = &HookedStateDB{}

// HookedStateDB represents a statedb which emits calls to tracing-hooks
// on state operations.
type HookedStateDB struct {
	*StateDB
	hooks *tracing.Hooks
}

// NewHookedState wraps the given stateDb with the given hooks
func NewHookedState(stateDb *StateDB, hooks *tracing.Hooks) *HookedStateDB {
	s := &HookedStateDB{stateDb, hooks}
	return s
}

// override to emit balance change events
func (s *HookedStateDB) SubBalance(addr common.Address, amount *uint256.Int, reason tracing.BalanceChangeReason) uint256.Int {
	s.StateDB.SubBalance(addr, amount, reason)
	if s.hooks != nil && s.hooks.OnBalanceChange != nil && !amount.IsZero() {
		prev := s.StateDB.GetBalance(addr)
		newBalance := new(uint256.Int).Sub(prev, amount)
		s.hooks.OnBalanceChange(addr, prev.ToBig(), newBalance.ToBig(), reason)
	}
	return uint256.Int{}
}

// override to emit balance change events
func (s *HookedStateDB) AddBalance(addr common.Address, amount *uint256.Int, reason tracing.BalanceChangeReason) uint256.Int {
	s.StateDB.AddBalance(addr, amount, reason)
	if s.hooks != nil && s.hooks.OnBalanceChange != nil && !amount.IsZero() {
		prev := s.StateDB.GetBalance(addr)
		newBalance := new(uint256.Int).Add(prev, amount)
		s.hooks.OnBalanceChange(addr, prev.ToBig(), newBalance.ToBig(), reason)
	}
	return uint256.Int{}
}

// override to emit nonce change events
func (s *HookedStateDB) SetNonce(address common.Address, nonce uint64, reason tracing.NonceChangeReason) {
	s.StateDB.SetNonce(address, nonce, reason)
	if s.hooks != nil && s.hooks.OnNonceChangeV2 != nil {
		s.hooks.OnNonceChangeV2(address, s.StateDB.GetNonce(address), nonce, reason)
	} else if s.hooks != nil && s.hooks.OnNonceChange != nil {
		s.hooks.OnNonceChange(address, s.StateDB.GetNonce(address), nonce)
	}
}

// override to emit code change events
func (s *HookedStateDB) SetCode(address common.Address, code []byte) []byte {
	prev := s.StateDB.SetCode(address, code)
	if s.hooks != nil && s.hooks.OnCodeChange != nil {
		prevHash := evmtypes.EmptyCodeHash
		if len(prev) != 0 {
			prevHash = crypto.Keccak256Hash(prev)
		}
		s.hooks.OnCodeChange(address, prevHash, prev, crypto.Keccak256Hash(code), code)
	}
	return prev
}

// override to emit storage change events
func (s *HookedStateDB) SetState(address common.Address, key common.Hash, value common.Hash) common.Hash {
	prev := s.StateDB.SetState(address, key, value)
	if s.hooks != nil && s.hooks.OnStorageChange != nil && prev != value {
		s.hooks.OnStorageChange(address, key, prev, value)
	}
	return prev
}

// override to emit balance change events
func (s *HookedStateDB) SelfDestruct(address common.Address) uint256.Int {
	var prevCode []byte
	var prevCodeHash common.Hash

	if s.hooks != nil && s.hooks.OnCodeChange != nil {
		prevCode = s.StateDB.GetCode(address)
		prevCodeHash = s.StateDB.GetCodeHash(address)
	}

	prev := s.StateDB.SelfDestruct(address)

	if s.hooks != nil && s.hooks.OnBalanceChange != nil && !prev.IsZero() {
		s.hooks.OnBalanceChange(address, prev.ToBig(), new(big.Int), tracing.BalanceDecreaseSelfdestruct)
	}

	if s.hooks != nil && s.hooks.OnCodeChange != nil && len(prevCode) > 0 {
		s.hooks.OnCodeChange(address, prevCodeHash, prevCode, evmtypes.EmptyCodeHash, nil)
	}

	return prev
}

// override to emit balance change events
func (s *HookedStateDB) SelfDestruct6780(address common.Address) (uint256.Int, bool) {
	var prevCode []byte
	var prevCodeHash common.Hash

	if s.hooks != nil && s.hooks.OnCodeChange != nil {
		prevCodeHash = s.StateDB.GetCodeHash(address)
		prevCode = s.StateDB.GetCode(address)
	}

	prev, changed := s.StateDB.SelfDestruct6780(address)

	if s.hooks != nil && s.hooks.OnBalanceChange != nil && changed && !prev.IsZero() {
		s.hooks.OnBalanceChange(address, prev.ToBig(), new(big.Int), tracing.BalanceDecreaseSelfdestruct)
	}

	if s.hooks != nil && s.hooks.OnCodeChange != nil && changed && len(prevCode) > 0 {
		s.hooks.OnCodeChange(address, prevCodeHash, prevCode, evmtypes.EmptyCodeHash, nil)
	}

	return prev, changed
}

// override to emit log events
func (s *HookedStateDB) AddLog(log *evmtypes.Log) {
	// The StateDB will modify the log (add fields), so invoke that first
	s.StateDB.AddLog(log)
	if s.hooks != nil && s.hooks.OnLog != nil {
		s.hooks.OnLog(log)
	}
}

// override to emit balance change events
func (s *HookedStateDB) Commit() (err error) {
	defer func() {
		err = s.StateDB.Commit()
	}()

	if s.hooks != nil && s.hooks.OnBalanceChange != nil {
		return s.memStoreSelfDestruct.Walk(s.ctx, nil, func(key []byte) (stop bool, err error) {
			addr := common.BytesToAddress(key)

			// If ether was sent to account post-selfdestruct it is burnt.
			if bal := s.GetBalance(addr); bal.Sign() != 0 {
				s.hooks.OnBalanceChange(addr, bal.ToBig(), new(big.Int), tracing.BalanceDecreaseSelfdestructBurn)
			}

			return false, err
		})
	}

	return nil
}
