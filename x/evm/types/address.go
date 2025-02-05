package types

import (
	"cosmossdk.io/core/address"

	"github.com/ethereum/go-ethereum/common"
)

// 0x0 null address
var NullAddress common.Address = common.HexToAddress("0x0")

// 0x1 std address
var StdAddress common.Address = common.HexToAddress("0x1")

// salt for precompiled contracts
var (
	ERC20FactorySalt  uint64 = 1
	ERC20WrapperSalt  uint64 = 2
	ConnectOracleSalt uint64 = 3
)

// 0xf1 Cosmos precompile address
var CosmosPrecompileAddress common.Address = common.HexToAddress("0xf1")

// 0xf2 ERC20Registry precompile address
var ERC20RegistryPrecompileAddress common.Address = common.HexToAddress("0xf2")

// 0xf3 JSONUtils precompile address
var JSONUtilsPrecompileAddress common.Address = common.HexToAddress("0xf3")

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
