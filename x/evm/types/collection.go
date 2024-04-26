package types

import (
	"context"
	"crypto/sha256"
	"errors"
	"math/big"
	"strings"

	"cosmossdk.io/collections"
	"github.com/ethereum/go-ethereum/common"
)

var (
	ClassTraceClassIdPrefixIBC = "ibc/"
	ClassTraceClassIdPrefixEVM = "evm/"
)

const (
	MaxNftCollectionNameLength   = 256
	MaxNftCollectionSymbolLength = 256
	MaxSftCollectionNameLength   = 256
	MaxSftCollectionSymbolLength = 256
)

type ERC721ClassIdKeeper interface {
	GetContractAddrByClassId(context.Context, string) (common.Address, error)
	GetClassIdByContractAddr(context.Context, common.Address) (string, error)
}

func ClassIdToContractAddr(ctx context.Context, k ERC721ClassIdKeeper, classId string) (common.Address, error) {
	if strings.HasPrefix(classId, ClassTraceClassIdPrefixEVM) {
		contractAddrInString := strings.TrimPrefix(classId, ClassTraceClassIdPrefixEVM)
		if !common.IsHexAddress(contractAddrInString) {
			return common.Address{}, ErrInvalidDenom
		}

		return common.HexToAddress(contractAddrInString), nil
	}

	return k.GetContractAddrByClassId(ctx, classId)
}

func ContractAddrToClassId(ctx context.Context, k ERC721ClassIdKeeper, contractAddr common.Address) (string, error) {
	denom, err := k.GetClassIdByContractAddr(ctx, contractAddr)
	if err != nil && errors.Is(err, collections.ErrNotFound) {
		return DENOM_PREFIX + strings.TrimPrefix(contractAddr.Hex(), "0x"), nil
	} else if err != nil {
		return "", err
	}

	return denom, nil
}

func IsERC721ClassId(classId string) bool {
	return strings.HasPrefix(classId, DENOM_PREFIX)
}

func TokenIdToBigInt(tokenId string) *big.Int {
	hash := sha256.New()
	hash.Write([]byte(tokenId))
	return new(big.Int).SetBytes(hash.Sum(nil))
}
