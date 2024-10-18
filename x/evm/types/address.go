package types

import (
	"cosmossdk.io/core/address"

	"github.com/ethereum/go-ethereum/common"
)

// 0x0 null address
var NullAddress common.Address = common.HexToAddress("0x0")

// 0x1 std address
var StdAddress common.Address = common.HexToAddress("0x1")

// ERC20FactorySalt is the salt used to create the ERC20 factory address
var ERC20FactorySalt = uint64(1)
var ERC20WrapperSalt = uint64(2)

// 0xf1 Cosmos precompile address
var CosmosPrecompileAddress common.Address = common.HexToAddress("0xf1")

// 0xf2 ERC20Registry precompile address
var ERC20RegistryPrecompileAddress common.Address = common.HexToAddress("0xf2")

// 0xf3 ERC721Registry precompile address
var ERC721RegistryPrecompileAddress common.Address = common.HexToAddress("0xf3")

// IsPrecompileAddress checks if the address is a precompile address
func IsPrecompileAddress(addr common.Address) bool {
	return addr == ERC20RegistryPrecompileAddress || addr == CosmosPrecompileAddress
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
