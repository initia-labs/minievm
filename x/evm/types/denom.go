package types

import (
	"context"
	"errors"
	"strings"

	"cosmossdk.io/collections"
	"github.com/ethereum/go-ethereum/common"
)

const DENOM_PREFIX = "evm/"

type ERC20Keeper interface {
	GetContractAddrByDenom(context.Context, string) (common.Address, error)
	GetDenomByContractAddr(context.Context, common.Address) (string, error)
}

func DenomToContractAddr(ctx context.Context, k ERC20Keeper, denom string) (common.Address, error) {
	if strings.HasPrefix(denom, DENOM_PREFIX) {
		contractAddrInString := strings.TrimPrefix(denom, DENOM_PREFIX)
		if !common.IsHexAddress(contractAddrInString) {
			return common.Address{}, ErrInvalidDenom
		}

		return common.HexToAddress(contractAddrInString), nil
	}

	return k.GetContractAddrByDenom(ctx, denom)
}

func ContractAddrToDenom(ctx context.Context, k ERC20Keeper, contractAddr common.Address) (string, error) {
	denom, err := k.GetDenomByContractAddr(ctx, contractAddr)
	if err != nil && errors.Is(err, collections.ErrNotFound) {
		return DENOM_PREFIX + strings.TrimPrefix(contractAddr.Hex(), "0x"), nil
	} else if err != nil {
		return "", err
	}

	return denom, nil
}

func IsERC20Denom(denom string) bool {
	return strings.HasPrefix(denom, DENOM_PREFIX)
}
