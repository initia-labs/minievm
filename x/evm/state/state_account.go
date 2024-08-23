package state

import (
	"encoding/binary"
)

type StateAccount struct {
	Nonce    uint64
	CodeHash []byte
}

func EmptyStateAccount() *StateAccount {
	return &StateAccount{
		Nonce:    0,
		CodeHash: []byte{},
	}
}

func (sa *StateAccount) IsEmpty() bool {
	return sa.Nonce == 0 && len(sa.CodeHash) == 0
}

func (sa StateAccount) Marshal() []byte {
	bz := make([]byte, 8+len(sa.CodeHash))

	binary.BigEndian.PutUint64(bz, sa.Nonce)
	copy(bz[8:], sa.CodeHash)
	return bz
}

func (sa *StateAccount) Unmarshal(bz []byte) *StateAccount {
	sa.Nonce = bytesToUint64(bz)
	sa.CodeHash = bz[8:]
	return sa
}
