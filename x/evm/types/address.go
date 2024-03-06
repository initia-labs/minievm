package types

import (
	"cosmossdk.io/core/address"
	"github.com/ethereum/go-ethereum/common"
)

// Parse string contract address to sdk.AccAddress
func ContractAddressFromString(ac address.Codec, contractAddrInString string) (contractAddr common.Address, err error) {
	if common.IsHexAddress(contractAddrInString) {
		contractAddr = common.HexToAddress(contractAddrInString)
	} else if contractAddrBytes, err := ac.StringToBytes(contractAddrInString); err != nil {
		return common.Address{}, err
	} else {
		contractAddr = common.BytesToAddress(contractAddrBytes)
	}

	return contractAddr, nil
}
