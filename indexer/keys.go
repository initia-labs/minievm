package indexer

import "encoding/binary"

var (
	// block indexes
	prefixBlock             = []byte{0, 0, 1, 1}
	prefixBlockHashToNumber = []byte{0, 0, 1, 2}

	// tx indexes
	prefixTx                    = []byte{0, 0, 2, 1}
	prefixTxReceipt             = []byte{0, 0, 2, 2}
	prefixBlockAndIndexToTxHash = []byte{0, 0, 2, 3}
)

func keyBlock(height uint64) []byte {
	key := make([]byte, 8)
	binary.BigEndian.PutUint64(key, height)
	return append(prefixBlock, key...)
}

func keyBlockHashToNumber(hash []byte) []byte {
	return append(prefixBlockHashToNumber, hash...)
}

func keyTx(hash []byte) []byte {
	return append(prefixTx, hash...)
}

func keyTxReceipt(hash []byte) []byte {
	return append(prefixTxReceipt, hash...)
}

func keyBlockAndIndexToTxHash(blockHeight uint64, index uint64) []byte {
	key := make([]byte, 16)
	binary.BigEndian.PutUint64(key[:8], blockHeight)
	binary.BigEndian.PutUint64(key[8:], index)
	return append(prefixBlockAndIndexToTxHash, key...)
}
