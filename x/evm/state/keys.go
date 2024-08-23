package state

import (
	"encoding/binary"

	"github.com/ethereum/go-ethereum/common"
)

var (
	AccountKeyPrefix  = []byte("account")
	CodeKeyPrefix     = []byte("code")
	CodeSizeKeyPrefix = []byte("codesize")
	StateKeyPrefix    = []byte("state")
)

func accountKey(addr common.Address) []byte {
	return append(addr.Bytes(), AccountKeyPrefix...)
}

func codeKey(addr common.Address, codeHash []byte) []byte {
	return append(addr.Bytes(), append(CodeKeyPrefix, codeHash...)...)
}

func codeSizeKey(addr common.Address, codeHash []byte) []byte {
	return append(addr.Bytes(), append(CodeSizeKeyPrefix, codeHash...)...)
}

func stateKey(addr common.Address, slot common.Hash) []byte {
	return append(addr.Bytes(), append(StateKeyPrefix, slot.Bytes()...)...)
}

func uint64ToBytes(v uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, v)
	return bz
}

func bytesToUint64(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
