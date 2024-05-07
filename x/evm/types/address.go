package types

import (
	"cosmossdk.io/core/address"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/holiman/uint256"

	"github.com/initia-labs/minievm/x/evm/contracts/erc20_factory"
)

// 0x0 null address
var NullAddress common.Address = common.HexToAddress("0x0")

// 0x1 std address
var StdAddress common.Address = common.HexToAddress("0x1")

// ERC20FactorySalt is the salt used to create the ERC20 factory address
var ERC20FactorySalt = uint64(1)

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

// factoryCodeHash is the hash of the factory code
var factoryCodeHash common.Hash
var factoryAddr common.Address

func init() {
	// to avoid repeated hashing, we hash the factory code once at init
	bz, err := hexutil.Decode(erc20_factory.Erc20FactoryBin)
	if err != nil {
		panic(err)
	}

	factoryCodeHash = crypto.Keccak256Hash(bz)
	factoryAddr = crypto.CreateAddress2(StdAddress, uint256.NewInt(ERC20FactorySalt).Bytes32(), factoryCodeHash.Bytes())
}

// ERC20FactoryAddress returns the address of the ERC20 factory
func ERC20FactoryAddress() common.Address {
	return factoryAddr
}
