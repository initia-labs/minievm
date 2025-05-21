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
	IBCPrefix = "ibc/"
	EVMPrefix = "evm/"
)

type ERC721ClassIdKeeper interface {
	GetContractAddrByClassId(context.Context, string) (common.Address, error)
	GetClassIdByContractAddr(context.Context, common.Address) (string, error)
}

func TokenIdToBigInt(classId string, tokenId string) (*big.Int, bool) {
	if strings.HasPrefix(classId, IBCPrefix) {
		hash := sha256.New()
		hash.Write([]byte(tokenId))
		return new(big.Int).SetBytes(hash.Sum(nil)), true
	}
	return new(big.Int).SetString(tokenId, 10)
}

func ContractAddressFromClassId(ctx context.Context, k ERC721ClassIdKeeper, classId string) (common.Address, error) {
	if after, ok := strings.CutPrefix(classId, EVMPrefix); ok {
		contractAddrInString := after
		if !common.IsHexAddress(contractAddrInString) {
			return NullAddress, ErrInvalidClassId
		}

		return common.HexToAddress(contractAddrInString), nil
	}

	return k.GetContractAddrByClassId(ctx, classId)
}

func ClassIdFromCollectionAddress(ctx context.Context, k ERC721ClassIdKeeper, contractAddr common.Address) (string, error) {
	classId, err := k.GetClassIdByContractAddr(ctx, contractAddr)
	if err != nil && errors.Is(err, collections.ErrNotFound) {
		return EVMPrefix + strings.TrimPrefix(contractAddr.Hex(), "0x"), nil
	} else if err != nil {
		return "", err
	}
	return classId, nil
}

func IsEVMClassId(classId string) bool {
	return strings.HasPrefix(classId, EVMPrefix)
}
