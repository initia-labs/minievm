package types

import (
	"crypto/sha256"
	"math/big"
	"strings"
)

var (
	ClassTraceClassIdPrefixIBC = "ibc/"
)

func TokenIdToBigInt(classId string, tokenId string) (*big.Int, bool) {
	if strings.HasPrefix(classId, ClassTraceClassIdPrefixIBC) {
		hash := sha256.New()
		hash.Write([]byte(tokenId))
		return new(big.Int).SetBytes(hash.Sum(nil)), true
	}
	return new(big.Int).SetString(tokenId, 10)
}
