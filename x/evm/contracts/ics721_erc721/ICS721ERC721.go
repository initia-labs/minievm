// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ics721_erc721

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// Ics721Erc721MetaData contains all meta data concerning the Ics721Erc721 contract.
var Ics721Erc721MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"classURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_tokenUri\",\"type\":\"string\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50604051612a2d380380612a2d83398181016040528101906100319190610255565b8282815f90816100419190610506565b5080600190816100519190610506565b5050503360065f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060f373ffffffffffffffffffffffffffffffffffffffff1663379da8466040518163ffffffff1660e01b81526004015f604051808303815f87803b1580156100da575f80fd5b505af11580156100ec573d5f803e3d5ffd5b5050505080600790816100ff9190610506565b505050506105d5565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61016782610121565b810181811067ffffffffffffffff8211171561018657610185610131565b5b80604052505050565b5f610198610108565b90506101a4828261015e565b919050565b5f67ffffffffffffffff8211156101c3576101c2610131565b5b6101cc82610121565b9050602081019050919050565b8281835e5f83830152505050565b5f6101f96101f4846101a9565b61018f565b9050828152602081018484840111156102155761021461011d565b5b6102208482856101d9565b509392505050565b5f82601f83011261023c5761023b610119565b5b815161024c8482602086016101e7565b91505092915050565b5f805f6060848603121561026c5761026b610111565b5b5f84015167ffffffffffffffff81111561028957610288610115565b5b61029586828701610228565b935050602084015167ffffffffffffffff8111156102b6576102b5610115565b5b6102c286828701610228565b925050604084015167ffffffffffffffff8111156102e3576102e2610115565b5b6102ef86828701610228565b9150509250925092565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061034757607f821691505b60208210810361035a57610359610303565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026103bc7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610381565b6103c68683610381565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61040a610405610400846103de565b6103e7565b6103de565b9050919050565b5f819050919050565b610423836103f0565b61043761042f82610411565b84845461038d565b825550505050565b5f90565b61044b61043f565b61045681848461041a565b505050565b5b818110156104795761046e5f82610443565b60018101905061045c565b5050565b601f8211156104be5761048f81610360565b61049884610372565b810160208510156104a7578190505b6104bb6104b385610372565b83018261045b565b50505b505050565b5f82821c905092915050565b5f6104de5f19846008026104c3565b1980831691505092915050565b5f6104f683836104cf565b9150826002028217905092915050565b61050f826102f9565b67ffffffffffffffff81111561052857610527610131565b5b6105328254610330565b61053d82828561047d565b5f60209050601f83116001811461056e575f841561055c578287015190505b61056685826104eb565b8655506105cd565b601f19841661057c86610360565b5f5b828110156105a35784890151825560018201915060208501945060208101905061057e565b868310156105c057848901516105bc601f8916826104cf565b8355505b6001600288020188555050505b505050505050565b61244b806105e25f395ff3fe608060405234801561000f575f80fd5b5060043610610114575f3560e01c80638da5cb5b116100a0578063b88d4fde1161006f578063b88d4fde146102dc578063c87b56dd146102f8578063d3fc986414610328578063e985e9c514610344578063f2fde38b1461037457610114565b80638da5cb5b1461026657806395d89b4114610284578063a22cb465146102a2578063b0a7fd4d146102be57610114565b806323b872dd116100e757806323b872dd146101b257806342842e0e146101ce57806342966c68146101ea5780636352211e1461020657806370a082311461023657610114565b806301ffc9a71461011857806306fdde0314610148578063081812fc14610166578063095ea7b314610196575b5f80fd5b610132600480360381019061012d91906119cb565b610390565b60405161013f9190611a10565b60405180910390f35b610150610471565b60405161015d9190611a99565b60405180910390f35b610180600480360381019061017b9190611aec565b610500565b60405161018d9190611b56565b60405180910390f35b6101b060048036038101906101ab9190611b99565b61051b565b005b6101cc60048036038101906101c79190611bd7565b610531565b005b6101e860048036038101906101e39190611bd7565b610630565b005b61020460048036038101906101ff9190611aec565b61064f565b005b610220600480360381019061021b9190611aec565b6106bd565b60405161022d9190611b56565b60405180910390f35b610250600480360381019061024b9190611c27565b6106ce565b60405161025d9190611c61565b60405180910390f35b61026e610784565b60405161027b9190611b56565b60405180910390f35b61028c6107a9565b6040516102999190611a99565b60405180910390f35b6102bc60048036038101906102b79190611ca4565b610839565b005b6102c661084f565b6040516102d39190611a99565b60405180910390f35b6102f660048036038101906102f19190611e0e565b6108df565b005b610312600480360381019061030d9190611aec565b6109eb565b60405161031f9190611a99565b60405180910390f35b610342600480360381019061033d9190611f2c565b610a8c565b005b61035e60048036038101906103599190611f98565b610bf9565b60405161036b9190611a10565b60405180910390f35b61038e60048036038101906103899190611c27565b610c87565b005b5f7f80ac58cd000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916148061045a57507f5b5e139f000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b8061046a575061046982610dd4565b5b9050919050565b60605f805461047f90612003565b80601f01602080910402602001604051908101604052809291908181526020018280546104ab90612003565b80156104f65780601f106104cd576101008083540402835291602001916104f6565b820191905f5260205f20905b8154815290600101906020018083116104d957829003601f168201915b5050505050905090565b5f61050a82610e3d565b5061051482610ec3565b9050919050565b61052d8282610528610efc565b610f03565b5050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036105a1575f6040517f64a0ae920000000000000000000000000000000000000000000000000000000081526004016105989190611b56565b60405180910390fd5b5f6105b483836105af610efc565b610f15565b90508373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461062a578382826040517f64283d7b00000000000000000000000000000000000000000000000000000000815260040161062193929190612033565b60405180910390fd5b50505050565b61064a83838360405180602001604052805f8152506108df565b505050565b5f61065982610e3d565b9050610666813384611120565b6106b057610672610efc565b826040517f177e802f0000000000000000000000000000000000000000000000000000000081526004016106a7929190612068565b60405180910390fd5b6106b9826111e0565b5050565b5f6106c782610e3d565b9050919050565b5f8073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361073f575f6040517f89c62b640000000000000000000000000000000000000000000000000000000081526004016107369190611b56565b60405180910390fd5b60035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6060600180546107b890612003565b80601f01602080910402602001604051908101604052809291908181526020018280546107e490612003565b801561082f5780601f106108065761010080835404028352916020019161082f565b820191905f5260205f20905b81548152906001019060200180831161081257829003601f168201915b5050505050905090565b61084b610844610efc565b8383611262565b5050565b60606007805461085e90612003565b80601f016020809104026020016040519081016040528092919081815260200182805461088a90612003565b80156108d55780601f106108ac576101008083540402835291602001916108d5565b820191905f5260205f20905b8154815290600101906020018083116108b857829003601f168201915b5050505050905090565b8260f373ffffffffffffffffffffffffffffffffffffffff1663fa75f257826040518263ffffffff1660e01b815260040161091a9190611b56565b602060405180830381865afa158015610935573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061095991906120a3565b6109c55760f373ffffffffffffffffffffffffffffffffffffffff1663d6e69551826040518263ffffffff1660e01b81526004016109979190611b56565b5f604051808303815f87803b1580156109ae575f80fd5b505af11580156109c0573d5f803e3d5ffd5b505050505b6109d0858585610531565b6109e46109db610efc565b868686866113cb565b5050505050565b606060085f8381526020019081526020015f208054610a0990612003565b80601f0160208091040260200160405190810160405280929190818152602001828054610a3590612003565b8015610a805780601f10610a5757610100808354040283529160200191610a80565b820191905f5260205f20905b815481529060010190602001808311610a6357829003601f168201915b50505050509050919050565b60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610ae4575f80fd5b8260f373ffffffffffffffffffffffffffffffffffffffff1663fa75f257826040518263ffffffff1660e01b8152600401610b1f9190611b56565b602060405180830381865afa158015610b3a573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b5e91906120a3565b610bca5760f373ffffffffffffffffffffffffffffffffffffffff1663d6e69551826040518263ffffffff1660e01b8152600401610b9c9190611b56565b5f604051808303815f87803b158015610bb3575f80fd5b505af1158015610bc5573d5f803e3d5ffd5b505050505b610bd48484611577565b8160085f8581526020019081526020015f209081610bf2919061226b565b5050505050565b5f60055f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905092915050565b60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610cdf575f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610d16575f80fd5b8073ffffffffffffffffffffffffffffffffffffffff1660065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a38060065f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b5f80610e4883611594565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610eba57826040517f7e273289000000000000000000000000000000000000000000000000000000008152600401610eb19190611c61565b60405180910390fd5b80915050919050565b5f60045f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b5f33905090565b610f1083838360016115cd565b505050565b5f80610f2084611594565b90505f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614610f6157610f6081848661178c565b5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610fec57610fa05f855f806115cd565b600160035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825403925050819055505b5f73ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff161461106b57600160035f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055505b8460025f8681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550838573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4809150509392505050565b5f8073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141580156111d757508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16148061119857506111978484610bf9565b5b806111d657508273ffffffffffffffffffffffffffffffffffffffff166111be83610ec3565b73ffffffffffffffffffffffffffffffffffffffff16145b5b90509392505050565b5f6111ec5f835f610f15565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361125e57816040517f7e2732890000000000000000000000000000000000000000000000000000000081526004016112559190611c61565b60405180910390fd5b5050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036112d257816040517f5b08ba180000000000000000000000000000000000000000000000000000000081526004016112c99190611b56565b60405180910390fd5b8060055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31836040516113be9190611a10565b60405180910390a3505050565b5f8373ffffffffffffffffffffffffffffffffffffffff163b1115611570578273ffffffffffffffffffffffffffffffffffffffff1663150b7a02868685856040518563ffffffff1660e01b8152600401611429949392919061238c565b6020604051808303815f875af192505050801561146457506040513d601f19601f8201168201806040525081019061146191906123ea565b60015b6114e5573d805f8114611492576040519150601f19603f3d011682016040523d82523d5f602084013e611497565b606091505b505f8151036114dd57836040517f64a0ae920000000000000000000000000000000000000000000000000000000081526004016114d49190611b56565b60405180910390fd5b805181602001fd5b63150b7a0260e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161461156e57836040517f64a0ae920000000000000000000000000000000000000000000000000000000081526004016115659190611b56565b60405180910390fd5b505b5050505050565b611590828260405180602001604052805f81525061184f565b5050565b5f60025f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b808061160557505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b15611737575f61161484610e3d565b90505f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415801561167e57508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b8015611691575061168f8184610bf9565b155b156116d357826040517fa9fbf51f0000000000000000000000000000000000000000000000000000000081526004016116ca9190611b56565b60405180910390fd5b811561173557838573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45b505b8360045f8581526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050565b611797838383611120565b61184a575f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361180b57806040517f7e2732890000000000000000000000000000000000000000000000000000000081526004016118029190611c61565b60405180910390fd5b81816040517f177e802f000000000000000000000000000000000000000000000000000000008152600401611841929190612068565b60405180910390fd5b505050565b6118598383611872565b61186d611864610efc565b5f8585856113cb565b505050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036118e2575f6040517f64a0ae920000000000000000000000000000000000000000000000000000000081526004016118d99190611b56565b60405180910390fd5b5f6118ee83835f610f15565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614611960575f6040517f73c6ac6e0000000000000000000000000000000000000000000000000000000081526004016119579190611b56565b60405180910390fd5b505050565b5f604051905090565b5f80fd5b5f80fd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6119aa81611976565b81146119b4575f80fd5b50565b5f813590506119c5816119a1565b92915050565b5f602082840312156119e0576119df61196e565b5b5f6119ed848285016119b7565b91505092915050565b5f8115159050919050565b611a0a816119f6565b82525050565b5f602082019050611a235f830184611a01565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f611a6b82611a29565b611a758185611a33565b9350611a85818560208601611a43565b611a8e81611a51565b840191505092915050565b5f6020820190508181035f830152611ab18184611a61565b905092915050565b5f819050919050565b611acb81611ab9565b8114611ad5575f80fd5b50565b5f81359050611ae681611ac2565b92915050565b5f60208284031215611b0157611b0061196e565b5b5f611b0e84828501611ad8565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f611b4082611b17565b9050919050565b611b5081611b36565b82525050565b5f602082019050611b695f830184611b47565b92915050565b611b7881611b36565b8114611b82575f80fd5b50565b5f81359050611b9381611b6f565b92915050565b5f8060408385031215611baf57611bae61196e565b5b5f611bbc85828601611b85565b9250506020611bcd85828601611ad8565b9150509250929050565b5f805f60608486031215611bee57611bed61196e565b5b5f611bfb86828701611b85565b9350506020611c0c86828701611b85565b9250506040611c1d86828701611ad8565b9150509250925092565b5f60208284031215611c3c57611c3b61196e565b5b5f611c4984828501611b85565b91505092915050565b611c5b81611ab9565b82525050565b5f602082019050611c745f830184611c52565b92915050565b611c83816119f6565b8114611c8d575f80fd5b50565b5f81359050611c9e81611c7a565b92915050565b5f8060408385031215611cba57611cb961196e565b5b5f611cc785828601611b85565b9250506020611cd885828601611c90565b9150509250929050565b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b611d2082611a51565b810181811067ffffffffffffffff82111715611d3f57611d3e611cea565b5b80604052505050565b5f611d51611965565b9050611d5d8282611d17565b919050565b5f67ffffffffffffffff821115611d7c57611d7b611cea565b5b611d8582611a51565b9050602081019050919050565b828183375f83830152505050565b5f611db2611dad84611d62565b611d48565b905082815260208101848484011115611dce57611dcd611ce6565b5b611dd9848285611d92565b509392505050565b5f82601f830112611df557611df4611ce2565b5b8135611e05848260208601611da0565b91505092915050565b5f805f8060808587031215611e2657611e2561196e565b5b5f611e3387828801611b85565b9450506020611e4487828801611b85565b9350506040611e5587828801611ad8565b925050606085013567ffffffffffffffff811115611e7657611e75611972565b5b611e8287828801611de1565b91505092959194509250565b5f67ffffffffffffffff821115611ea857611ea7611cea565b5b611eb182611a51565b9050602081019050919050565b5f611ed0611ecb84611e8e565b611d48565b905082815260208101848484011115611eec57611eeb611ce6565b5b611ef7848285611d92565b509392505050565b5f82601f830112611f1357611f12611ce2565b5b8135611f23848260208601611ebe565b91505092915050565b5f805f60608486031215611f4357611f4261196e565b5b5f611f5086828701611b85565b9350506020611f6186828701611ad8565b925050604084013567ffffffffffffffff811115611f8257611f81611972565b5b611f8e86828701611eff565b9150509250925092565b5f8060408385031215611fae57611fad61196e565b5b5f611fbb85828601611b85565b9250506020611fcc85828601611b85565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061201a57607f821691505b60208210810361202d5761202c611fd6565b5b50919050565b5f6060820190506120465f830186611b47565b6120536020830185611c52565b6120606040830184611b47565b949350505050565b5f60408201905061207b5f830185611b47565b6120886020830184611c52565b9392505050565b5f8151905061209d81611c7a565b92915050565b5f602082840312156120b8576120b761196e565b5b5f6120c58482850161208f565b91505092915050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261212a7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826120ef565b61213486836120ef565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61216f61216a61216584611ab9565b61214c565b611ab9565b9050919050565b5f819050919050565b61218883612155565b61219c61219482612176565b8484546120fb565b825550505050565b5f90565b6121b06121a4565b6121bb81848461217f565b505050565b5b818110156121de576121d35f826121a8565b6001810190506121c1565b5050565b601f821115612223576121f4816120ce565b6121fd846120e0565b8101602085101561220c578190505b612220612218856120e0565b8301826121c0565b50505b505050565b5f82821c905092915050565b5f6122435f1984600802612228565b1980831691505092915050565b5f61225b8383612234565b9150826002028217905092915050565b61227482611a29565b67ffffffffffffffff81111561228d5761228c611cea565b5b6122978254612003565b6122a28282856121e2565b5f60209050601f8311600181146122d3575f84156122c1578287015190505b6122cb8582612250565b865550612332565b601f1984166122e1866120ce565b5f5b82811015612308578489015182556001820191506020850194506020810190506122e3565b868310156123255784890151612321601f891682612234565b8355505b6001600288020188555050505b505050505050565b5f81519050919050565b5f82825260208201905092915050565b5f61235e8261233a565b6123688185612344565b9350612378818560208601611a43565b61238181611a51565b840191505092915050565b5f60808201905061239f5f830187611b47565b6123ac6020830186611b47565b6123b96040830185611c52565b81810360608301526123cb8184612354565b905095945050505050565b5f815190506123e4816119a1565b92915050565b5f602082840312156123ff576123fe61196e565b5b5f61240c848285016123d6565b9150509291505056fea2646970667358221220325ab67cd4d5c0bf45a41c2727ed406c5e0c2fa8fd3ed83593d57f42aa94a18064736f6c63430008190033",
}

// Ics721Erc721ABI is the input ABI used to generate the binding from.
// Deprecated: Use Ics721Erc721MetaData.ABI instead.
var Ics721Erc721ABI = Ics721Erc721MetaData.ABI

// Ics721Erc721Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Ics721Erc721MetaData.Bin instead.
var Ics721Erc721Bin = Ics721Erc721MetaData.Bin

// DeployIcs721Erc721 deploys a new Ethereum contract, binding an instance of Ics721Erc721 to it.
func DeployIcs721Erc721(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string, uri_ string) (common.Address, *types.Transaction, *Ics721Erc721, error) {
	parsed, err := Ics721Erc721MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Ics721Erc721Bin), backend, name_, symbol_, uri_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ics721Erc721{Ics721Erc721Caller: Ics721Erc721Caller{contract: contract}, Ics721Erc721Transactor: Ics721Erc721Transactor{contract: contract}, Ics721Erc721Filterer: Ics721Erc721Filterer{contract: contract}}, nil
}

// Ics721Erc721 is an auto generated Go binding around an Ethereum contract.
type Ics721Erc721 struct {
	Ics721Erc721Caller     // Read-only binding to the contract
	Ics721Erc721Transactor // Write-only binding to the contract
	Ics721Erc721Filterer   // Log filterer for contract events
}

// Ics721Erc721Caller is an auto generated read-only Go binding around an Ethereum contract.
type Ics721Erc721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ics721Erc721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Ics721Erc721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ics721Erc721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Ics721Erc721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ics721Erc721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Ics721Erc721Session struct {
	Contract     *Ics721Erc721     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Ics721Erc721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Ics721Erc721CallerSession struct {
	Contract *Ics721Erc721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// Ics721Erc721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Ics721Erc721TransactorSession struct {
	Contract     *Ics721Erc721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// Ics721Erc721Raw is an auto generated low-level Go binding around an Ethereum contract.
type Ics721Erc721Raw struct {
	Contract *Ics721Erc721 // Generic contract binding to access the raw methods on
}

// Ics721Erc721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Ics721Erc721CallerRaw struct {
	Contract *Ics721Erc721Caller // Generic read-only contract binding to access the raw methods on
}

// Ics721Erc721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Ics721Erc721TransactorRaw struct {
	Contract *Ics721Erc721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIcs721Erc721 creates a new instance of Ics721Erc721, bound to a specific deployed contract.
func NewIcs721Erc721(address common.Address, backend bind.ContractBackend) (*Ics721Erc721, error) {
	contract, err := bindIcs721Erc721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ics721Erc721{Ics721Erc721Caller: Ics721Erc721Caller{contract: contract}, Ics721Erc721Transactor: Ics721Erc721Transactor{contract: contract}, Ics721Erc721Filterer: Ics721Erc721Filterer{contract: contract}}, nil
}

// NewIcs721Erc721Caller creates a new read-only instance of Ics721Erc721, bound to a specific deployed contract.
func NewIcs721Erc721Caller(address common.Address, caller bind.ContractCaller) (*Ics721Erc721Caller, error) {
	contract, err := bindIcs721Erc721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Ics721Erc721Caller{contract: contract}, nil
}

// NewIcs721Erc721Transactor creates a new write-only instance of Ics721Erc721, bound to a specific deployed contract.
func NewIcs721Erc721Transactor(address common.Address, transactor bind.ContractTransactor) (*Ics721Erc721Transactor, error) {
	contract, err := bindIcs721Erc721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Ics721Erc721Transactor{contract: contract}, nil
}

// NewIcs721Erc721Filterer creates a new log filterer instance of Ics721Erc721, bound to a specific deployed contract.
func NewIcs721Erc721Filterer(address common.Address, filterer bind.ContractFilterer) (*Ics721Erc721Filterer, error) {
	contract, err := bindIcs721Erc721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Ics721Erc721Filterer{contract: contract}, nil
}

// bindIcs721Erc721 binds a generic wrapper to an already deployed contract.
func bindIcs721Erc721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Ics721Erc721MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ics721Erc721 *Ics721Erc721Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ics721Erc721.Contract.Ics721Erc721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ics721Erc721 *Ics721Erc721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ics721Erc721.Contract.Ics721Erc721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ics721Erc721 *Ics721Erc721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ics721Erc721.Contract.Ics721Erc721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ics721Erc721 *Ics721Erc721CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ics721Erc721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ics721Erc721 *Ics721Erc721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ics721Erc721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ics721Erc721 *Ics721Erc721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ics721Erc721.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Ics721Erc721 *Ics721Erc721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ics721Erc721.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Ics721Erc721 *Ics721Erc721Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Ics721Erc721.Contract.BalanceOf(&_Ics721Erc721.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Ics721Erc721 *Ics721Erc721CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Ics721Erc721.Contract.BalanceOf(&_Ics721Erc721.CallOpts, owner)
}

// ClassURI is a free data retrieval call binding the contract method 0xb0a7fd4d.
//
// Solidity: function classURI() view returns(string)
func (_Ics721Erc721 *Ics721Erc721Caller) ClassURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Ics721Erc721.contract.Call(opts, &out, "classURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ClassURI is a free data retrieval call binding the contract method 0xb0a7fd4d.
//
// Solidity: function classURI() view returns(string)
func (_Ics721Erc721 *Ics721Erc721Session) ClassURI() (string, error) {
	return _Ics721Erc721.Contract.ClassURI(&_Ics721Erc721.CallOpts)
}

// ClassURI is a free data retrieval call binding the contract method 0xb0a7fd4d.
//
// Solidity: function classURI() view returns(string)
func (_Ics721Erc721 *Ics721Erc721CallerSession) ClassURI() (string, error) {
	return _Ics721Erc721.Contract.ClassURI(&_Ics721Erc721.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Ics721Erc721 *Ics721Erc721Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Ics721Erc721.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Ics721Erc721 *Ics721Erc721Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Ics721Erc721.Contract.GetApproved(&_Ics721Erc721.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Ics721Erc721 *Ics721Erc721CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Ics721Erc721.Contract.GetApproved(&_Ics721Erc721.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Ics721Erc721 *Ics721Erc721Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Ics721Erc721.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Ics721Erc721 *Ics721Erc721Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Ics721Erc721.Contract.IsApprovedForAll(&_Ics721Erc721.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Ics721Erc721 *Ics721Erc721CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Ics721Erc721.Contract.IsApprovedForAll(&_Ics721Erc721.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Ics721Erc721 *Ics721Erc721Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Ics721Erc721.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Ics721Erc721 *Ics721Erc721Session) Name() (string, error) {
	return _Ics721Erc721.Contract.Name(&_Ics721Erc721.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Ics721Erc721 *Ics721Erc721CallerSession) Name() (string, error) {
	return _Ics721Erc721.Contract.Name(&_Ics721Erc721.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ics721Erc721 *Ics721Erc721Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ics721Erc721.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ics721Erc721 *Ics721Erc721Session) Owner() (common.Address, error) {
	return _Ics721Erc721.Contract.Owner(&_Ics721Erc721.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ics721Erc721 *Ics721Erc721CallerSession) Owner() (common.Address, error) {
	return _Ics721Erc721.Contract.Owner(&_Ics721Erc721.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Ics721Erc721 *Ics721Erc721Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Ics721Erc721.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Ics721Erc721 *Ics721Erc721Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Ics721Erc721.Contract.OwnerOf(&_Ics721Erc721.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Ics721Erc721 *Ics721Erc721CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Ics721Erc721.Contract.OwnerOf(&_Ics721Erc721.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Ics721Erc721 *Ics721Erc721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Ics721Erc721.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Ics721Erc721 *Ics721Erc721Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Ics721Erc721.Contract.SupportsInterface(&_Ics721Erc721.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Ics721Erc721 *Ics721Erc721CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Ics721Erc721.Contract.SupportsInterface(&_Ics721Erc721.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Ics721Erc721 *Ics721Erc721Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Ics721Erc721.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Ics721Erc721 *Ics721Erc721Session) Symbol() (string, error) {
	return _Ics721Erc721.Contract.Symbol(&_Ics721Erc721.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Ics721Erc721 *Ics721Erc721CallerSession) Symbol() (string, error) {
	return _Ics721Erc721.Contract.Symbol(&_Ics721Erc721.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Ics721Erc721 *Ics721Erc721Caller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Ics721Erc721.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Ics721Erc721 *Ics721Erc721Session) TokenURI(tokenId *big.Int) (string, error) {
	return _Ics721Erc721.Contract.TokenURI(&_Ics721Erc721.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Ics721Erc721 *Ics721Erc721CallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Ics721Erc721.Contract.TokenURI(&_Ics721Erc721.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Ics721Erc721 *Ics721Erc721Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ics721Erc721.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Ics721Erc721 *Ics721Erc721Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ics721Erc721.Contract.Approve(&_Ics721Erc721.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Ics721Erc721 *Ics721Erc721TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ics721Erc721.Contract.Approve(&_Ics721Erc721.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Ics721Erc721 *Ics721Erc721Transactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Ics721Erc721.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Ics721Erc721 *Ics721Erc721Session) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Ics721Erc721.Contract.Burn(&_Ics721Erc721.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Ics721Erc721 *Ics721Erc721TransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Ics721Erc721.Contract.Burn(&_Ics721Erc721.TransactOpts, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0xd3fc9864.
//
// Solidity: function mint(address receiver, uint256 tokenId, string _tokenUri) returns()
func (_Ics721Erc721 *Ics721Erc721Transactor) Mint(opts *bind.TransactOpts, receiver common.Address, tokenId *big.Int, _tokenUri string) (*types.Transaction, error) {
	return _Ics721Erc721.contract.Transact(opts, "mint", receiver, tokenId, _tokenUri)
}

// Mint is a paid mutator transaction binding the contract method 0xd3fc9864.
//
// Solidity: function mint(address receiver, uint256 tokenId, string _tokenUri) returns()
func (_Ics721Erc721 *Ics721Erc721Session) Mint(receiver common.Address, tokenId *big.Int, _tokenUri string) (*types.Transaction, error) {
	return _Ics721Erc721.Contract.Mint(&_Ics721Erc721.TransactOpts, receiver, tokenId, _tokenUri)
}

// Mint is a paid mutator transaction binding the contract method 0xd3fc9864.
//
// Solidity: function mint(address receiver, uint256 tokenId, string _tokenUri) returns()
func (_Ics721Erc721 *Ics721Erc721TransactorSession) Mint(receiver common.Address, tokenId *big.Int, _tokenUri string) (*types.Transaction, error) {
	return _Ics721Erc721.Contract.Mint(&_Ics721Erc721.TransactOpts, receiver, tokenId, _tokenUri)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Ics721Erc721 *Ics721Erc721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ics721Erc721.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Ics721Erc721 *Ics721Erc721Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ics721Erc721.Contract.SafeTransferFrom(&_Ics721Erc721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Ics721Erc721 *Ics721Erc721TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ics721Erc721.Contract.SafeTransferFrom(&_Ics721Erc721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Ics721Erc721 *Ics721Erc721Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Ics721Erc721.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Ics721Erc721 *Ics721Erc721Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Ics721Erc721.Contract.SafeTransferFrom0(&_Ics721Erc721.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Ics721Erc721 *Ics721Erc721TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Ics721Erc721.Contract.SafeTransferFrom0(&_Ics721Erc721.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Ics721Erc721 *Ics721Erc721Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Ics721Erc721.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Ics721Erc721 *Ics721Erc721Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Ics721Erc721.Contract.SetApprovalForAll(&_Ics721Erc721.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Ics721Erc721 *Ics721Erc721TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Ics721Erc721.Contract.SetApprovalForAll(&_Ics721Erc721.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Ics721Erc721 *Ics721Erc721Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ics721Erc721.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Ics721Erc721 *Ics721Erc721Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ics721Erc721.Contract.TransferFrom(&_Ics721Erc721.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Ics721Erc721 *Ics721Erc721TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ics721Erc721.Contract.TransferFrom(&_Ics721Erc721.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ics721Erc721 *Ics721Erc721Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ics721Erc721.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ics721Erc721 *Ics721Erc721Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ics721Erc721.Contract.TransferOwnership(&_Ics721Erc721.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ics721Erc721 *Ics721Erc721TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ics721Erc721.Contract.TransferOwnership(&_Ics721Erc721.TransactOpts, newOwner)
}

// Ics721Erc721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Ics721Erc721 contract.
type Ics721Erc721ApprovalIterator struct {
	Event *Ics721Erc721Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Ics721Erc721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Ics721Erc721Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Ics721Erc721Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Ics721Erc721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Ics721Erc721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Ics721Erc721Approval represents a Approval event raised by the Ics721Erc721 contract.
type Ics721Erc721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Ics721Erc721 *Ics721Erc721Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*Ics721Erc721ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Ics721Erc721.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &Ics721Erc721ApprovalIterator{contract: _Ics721Erc721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Ics721Erc721 *Ics721Erc721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Ics721Erc721Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Ics721Erc721.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Ics721Erc721Approval)
				if err := _Ics721Erc721.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Ics721Erc721 *Ics721Erc721Filterer) ParseApproval(log types.Log) (*Ics721Erc721Approval, error) {
	event := new(Ics721Erc721Approval)
	if err := _Ics721Erc721.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Ics721Erc721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Ics721Erc721 contract.
type Ics721Erc721ApprovalForAllIterator struct {
	Event *Ics721Erc721ApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Ics721Erc721ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Ics721Erc721ApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Ics721Erc721ApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Ics721Erc721ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Ics721Erc721ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Ics721Erc721ApprovalForAll represents a ApprovalForAll event raised by the Ics721Erc721 contract.
type Ics721Erc721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Ics721Erc721 *Ics721Erc721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*Ics721Erc721ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Ics721Erc721.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &Ics721Erc721ApprovalForAllIterator{contract: _Ics721Erc721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Ics721Erc721 *Ics721Erc721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *Ics721Erc721ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Ics721Erc721.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Ics721Erc721ApprovalForAll)
				if err := _Ics721Erc721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Ics721Erc721 *Ics721Erc721Filterer) ParseApprovalForAll(log types.Log) (*Ics721Erc721ApprovalForAll, error) {
	event := new(Ics721Erc721ApprovalForAll)
	if err := _Ics721Erc721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Ics721Erc721OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ics721Erc721 contract.
type Ics721Erc721OwnershipTransferredIterator struct {
	Event *Ics721Erc721OwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Ics721Erc721OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Ics721Erc721OwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Ics721Erc721OwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Ics721Erc721OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Ics721Erc721OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Ics721Erc721OwnershipTransferred represents a OwnershipTransferred event raised by the Ics721Erc721 contract.
type Ics721Erc721OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ics721Erc721 *Ics721Erc721Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Ics721Erc721OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ics721Erc721.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Ics721Erc721OwnershipTransferredIterator{contract: _Ics721Erc721.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ics721Erc721 *Ics721Erc721Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Ics721Erc721OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ics721Erc721.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Ics721Erc721OwnershipTransferred)
				if err := _Ics721Erc721.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ics721Erc721 *Ics721Erc721Filterer) ParseOwnershipTransferred(log types.Log) (*Ics721Erc721OwnershipTransferred, error) {
	event := new(Ics721Erc721OwnershipTransferred)
	if err := _Ics721Erc721.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Ics721Erc721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Ics721Erc721 contract.
type Ics721Erc721TransferIterator struct {
	Event *Ics721Erc721Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Ics721Erc721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Ics721Erc721Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Ics721Erc721Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Ics721Erc721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Ics721Erc721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Ics721Erc721Transfer represents a Transfer event raised by the Ics721Erc721 contract.
type Ics721Erc721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Ics721Erc721 *Ics721Erc721Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*Ics721Erc721TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Ics721Erc721.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &Ics721Erc721TransferIterator{contract: _Ics721Erc721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Ics721Erc721 *Ics721Erc721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Ics721Erc721Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Ics721Erc721.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Ics721Erc721Transfer)
				if err := _Ics721Erc721.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Ics721Erc721 *Ics721Erc721Filterer) ParseTransfer(log types.Log) (*Ics721Erc721Transfer, error) {
	event := new(Ics721Erc721Transfer)
	if err := _Ics721Erc721.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
