package types

import (
	"cosmossdk.io/core/address"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// 0x0 null address
var NullAddress common.Address = common.HexToAddress("0x0")

// 0x1 std address
var StdAddress common.Address = common.HexToAddress("0x1")

// 0xf1 ERC20 precompile address
var ERC20PrecompileAddress common.Address = common.HexToAddress("0xf1")

func FactoryAddress() common.Address {
	return crypto.CreateAddress(StdAddress, 0)
}

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
