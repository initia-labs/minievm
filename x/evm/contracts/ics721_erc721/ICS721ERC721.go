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
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"classURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_tokenUri\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_tokenOriginId\",\"type\":\"string\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_tokenUri\",\"type\":\"string\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenOriginId\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50604051612d4a380380612d4a83398181016040528101906100319190610255565b8282815f90816100419190610506565b5080600190816100519190610506565b5050503360065f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060f373ffffffffffffffffffffffffffffffffffffffff1663379da8466040518163ffffffff1660e01b81526004015f604051808303815f87803b1580156100da575f80fd5b505af11580156100ec573d5f803e3d5ffd5b5050505080600790816100ff9190610506565b505050506105d5565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61016782610121565b810181811067ffffffffffffffff8211171561018657610185610131565b5b80604052505050565b5f610198610108565b90506101a4828261015e565b919050565b5f67ffffffffffffffff8211156101c3576101c2610131565b5b6101cc82610121565b9050602081019050919050565b8281835e5f83830152505050565b5f6101f96101f4846101a9565b61018f565b9050828152602081018484840111156102155761021461011d565b5b6102208482856101d9565b509392505050565b5f82601f83011261023c5761023b610119565b5b815161024c8482602086016101e7565b91505092915050565b5f805f6060848603121561026c5761026b610111565b5b5f84015167ffffffffffffffff81111561028957610288610115565b5b61029586828701610228565b935050602084015167ffffffffffffffff8111156102b6576102b5610115565b5b6102c286828701610228565b925050604084015167ffffffffffffffff8111156102e3576102e2610115565b5b6102ef86828701610228565b9150509250925092565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061034757607f821691505b60208210810361035a57610359610303565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026103bc7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610381565b6103c68683610381565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61040a610405610400846103de565b6103e7565b6103de565b9050919050565b5f819050919050565b610423836103f0565b61043761042f82610411565b84845461038d565b825550505050565b5f90565b61044b61043f565b61045681848461041a565b505050565b5b818110156104795761046e5f82610443565b60018101905061045c565b5050565b601f8211156104be5761048f81610360565b61049884610372565b810160208510156104a7578190505b6104bb6104b385610372565b83018261045b565b50505b505050565b5f82821c905092915050565b5f6104de5f19846008026104c3565b1980831691505092915050565b5f6104f683836104cf565b9150826002028217905092915050565b61050f826102f9565b67ffffffffffffffff81111561052857610527610131565b5b6105328254610330565b61053d82828561047d565b5f60209050601f83116001811461056e575f841561055c578287015190505b61056685826104eb565b8655506105cd565b601f19841661057c86610360565b5f5b828110156105a35784890151825560018201915060208501945060208101905061057e565b868310156105c057848901516105bc601f8916826104cf565b8355505b6001600288020188555050505b505050505050565b612768806105e25f395ff3fe608060405234801561000f575f80fd5b506004361061012a575f3560e01c806370a08231116100ab578063b88d4fde1161006f578063b88d4fde1461033e578063c87b56dd1461035a578063d3fc98641461038a578063e985e9c5146103a6578063f2fde38b146103d65761012a565b806370a08231146102985780638da5cb5b146102c857806395d89b41146102e6578063a22cb46514610304578063b0a7fd4d146103205761012a565b80632fb102cf116100f25780632fb102cf146101e457806342842e0e1461020057806342966c681461021c5780636352211e146102385780636c8a5e77146102685761012a565b806301ffc9a71461012e57806306fdde031461015e578063081812fc1461017c578063095ea7b3146101ac57806323b872dd146101c8575b5f80fd5b61014860048036038101906101439190611c4c565b6103f2565b6040516101559190611c91565b60405180910390f35b6101666104d3565b6040516101739190611d1a565b60405180910390f35b61019660048036038101906101919190611d6d565b610562565b6040516101a39190611dd7565b60405180910390f35b6101c660048036038101906101c19190611e1a565b61057d565b005b6101e260048036038101906101dd9190611e58565b610593565b005b6101fe60048036038101906101f99190611fd4565b610692565b005b61021a60048036038101906102159190611e58565b61081f565b005b61023660048036038101906102319190611d6d565b61083e565b005b610252600480360381019061024d9190611d6d565b6108ac565b60405161025f9190611dd7565b60405180910390f35b610282600480360381019061027d9190611d6d565b6108bd565b60405161028f9190611d1a565b60405180910390f35b6102b260048036038101906102ad9190612070565b61095e565b6040516102bf91906120aa565b60405180910390f35b6102d0610a14565b6040516102dd9190611dd7565b60405180910390f35b6102ee610a39565b6040516102fb9190611d1a565b60405180910390f35b61031e600480360381019061031991906120ed565b610ac9565b005b610328610adf565b6040516103359190611d1a565b60405180910390f35b610358600480360381019061035391906121c9565b610b6f565b005b610374600480360381019061036f9190611d6d565b610c7b565b6040516103819190611d1a565b60405180910390f35b6103a4600480360381019061039f9190612249565b610d1c565b005b6103c060048036038101906103bb91906122b5565b610e7a565b6040516103cd9190611c91565b60405180910390f35b6103f060048036038101906103eb9190612070565b610f08565b005b5f7f80ac58cd000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614806104bc57507f5b5e139f000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b806104cc57506104cb82611055565b5b9050919050565b60605f80546104e190612320565b80601f016020809104026020016040519081016040528092919081815260200182805461050d90612320565b80156105585780601f1061052f57610100808354040283529160200191610558565b820191905f5260205f20905b81548152906001019060200180831161053b57829003601f168201915b5050505050905090565b5f61056c826110be565b5061057682611144565b9050919050565b61058f828261058a61117d565b611184565b5050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610603575f6040517f64a0ae920000000000000000000000000000000000000000000000000000000081526004016105fa9190611dd7565b60405180910390fd5b5f610616838361061161117d565b611196565b90508373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461068c578382826040517f64283d7b00000000000000000000000000000000000000000000000000000000815260040161068393929190612350565b60405180910390fd5b50505050565b60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106ea575f80fd5b8360f373ffffffffffffffffffffffffffffffffffffffff1663fa75f257826040518263ffffffff1660e01b81526004016107259190611dd7565b602060405180830381865afa158015610740573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107649190612399565b6107d05760f373ffffffffffffffffffffffffffffffffffffffff1663d6e69551826040518263ffffffff1660e01b81526004016107a29190611dd7565b5f604051808303815f87803b1580156107b9575f80fd5b505af11580156107cb573d5f803e3d5ffd5b505050505b6107da85856113a1565b8260085f8681526020019081526020015f2090816107f89190612561565b508160095f8681526020019081526020015f2090816108179190612561565b505050505050565b61083983838360405180602001604052805f815250610b6f565b505050565b5f610848826110be565b90506108558133846113be565b61089f5761086161117d565b826040517f177e802f000000000000000000000000000000000000000000000000000000008152600401610896929190612630565b60405180910390fd5b6108a88261147e565b5050565b5f6108b6826110be565b9050919050565b606060095f8381526020019081526020015f2080546108db90612320565b80601f016020809104026020016040519081016040528092919081815260200182805461090790612320565b80156109525780601f1061092957610100808354040283529160200191610952565b820191905f5260205f20905b81548152906001019060200180831161093557829003601f168201915b50505050509050919050565b5f8073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036109cf575f6040517f89c62b640000000000000000000000000000000000000000000000000000000081526004016109c69190611dd7565b60405180910390fd5b60035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b606060018054610a4890612320565b80601f0160208091040260200160405190810160405280929190818152602001828054610a7490612320565b8015610abf5780601f10610a9657610100808354040283529160200191610abf565b820191905f5260205f20905b815481529060010190602001808311610aa257829003601f168201915b5050505050905090565b610adb610ad461117d565b8383611500565b5050565b606060078054610aee90612320565b80601f0160208091040260200160405190810160405280929190818152602001828054610b1a90612320565b8015610b655780601f10610b3c57610100808354040283529160200191610b65565b820191905f5260205f20905b815481529060010190602001808311610b4857829003601f168201915b5050505050905090565b8260f373ffffffffffffffffffffffffffffffffffffffff1663fa75f257826040518263ffffffff1660e01b8152600401610baa9190611dd7565b602060405180830381865afa158015610bc5573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610be99190612399565b610c555760f373ffffffffffffffffffffffffffffffffffffffff1663d6e69551826040518263ffffffff1660e01b8152600401610c279190611dd7565b5f604051808303815f87803b158015610c3e575f80fd5b505af1158015610c50573d5f803e3d5ffd5b505050505b610c60858585610593565b610c74610c6b61117d565b86868686611669565b5050505050565b606060085f8381526020019081526020015f208054610c9990612320565b80601f0160208091040260200160405190810160405280929190818152602001828054610cc590612320565b8015610d105780601f10610ce757610100808354040283529160200191610d10565b820191905f5260205f20905b815481529060010190602001808311610cf357829003601f168201915b50505050509050919050565b60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610d74575f80fd5b8260f373ffffffffffffffffffffffffffffffffffffffff1663fa75f257826040518263ffffffff1660e01b8152600401610daf9190611dd7565b602060405180830381865afa158015610dca573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610dee9190612399565b610e5a5760f373ffffffffffffffffffffffffffffffffffffffff1663d6e69551826040518263ffffffff1660e01b8152600401610e2c9190611dd7565b5f604051808303815f87803b158015610e43575f80fd5b505af1158015610e55573d5f803e3d5ffd5b505050505b610e7484848460405180602001604052805f815250610692565b50505050565b5f60055f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905092915050565b60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610f60575f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610f97575f80fd5b8073ffffffffffffffffffffffffffffffffffffffff1660065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a38060065f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b5f806110c983611815565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361113b57826040517f7e27328900000000000000000000000000000000000000000000000000000000815260040161113291906120aa565b60405180910390fd5b80915050919050565b5f60045f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b5f33905090565b611191838383600161184e565b505050565b5f806111a184611815565b90505f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16146111e2576111e1818486611a0d565b5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461126d576112215f855f8061184e565b600160035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825403925050819055505b5f73ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff16146112ec57600160035f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055505b8460025f8681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550838573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4809150509392505050565b6113ba828260405180602001604052805f815250611ad0565b5050565b5f8073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415801561147557508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16148061143657506114358484610e7a565b5b8061147457508273ffffffffffffffffffffffffffffffffffffffff1661145c83611144565b73ffffffffffffffffffffffffffffffffffffffff16145b5b90509392505050565b5f61148a5f835f611196565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036114fc57816040517f7e2732890000000000000000000000000000000000000000000000000000000081526004016114f391906120aa565b60405180910390fd5b5050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361157057816040517f5b08ba180000000000000000000000000000000000000000000000000000000081526004016115679190611dd7565b60405180910390fd5b8060055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c318360405161165c9190611c91565b60405180910390a3505050565b5f8373ffffffffffffffffffffffffffffffffffffffff163b111561180e578273ffffffffffffffffffffffffffffffffffffffff1663150b7a02868685856040518563ffffffff1660e01b81526004016116c794939291906126a9565b6020604051808303815f875af192505050801561170257506040513d601f19601f820116820180604052508101906116ff9190612707565b60015b611783573d805f8114611730576040519150601f19603f3d011682016040523d82523d5f602084013e611735565b606091505b505f81510361177b57836040517f64a0ae920000000000000000000000000000000000000000000000000000000081526004016117729190611dd7565b60405180910390fd5b805181602001fd5b63150b7a0260e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161461180c57836040517f64a0ae920000000000000000000000000000000000000000000000000000000081526004016118039190611dd7565b60405180910390fd5b505b5050505050565b5f60025f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b808061188657505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b156119b8575f611895846110be565b90505f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141580156118ff57508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b801561191257506119108184610e7a565b155b1561195457826040517fa9fbf51f00000000000000000000000000000000000000000000000000000000815260040161194b9190611dd7565b60405180910390fd5b81156119b657838573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45b505b8360045f8581526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050565b611a188383836113be565b611acb575f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603611a8c57806040517f7e273289000000000000000000000000000000000000000000000000000000008152600401611a8391906120aa565b60405180910390fd5b81816040517f177e802f000000000000000000000000000000000000000000000000000000008152600401611ac2929190612630565b60405180910390fd5b505050565b611ada8383611af3565b611aee611ae561117d565b5f858585611669565b505050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603611b63575f6040517f64a0ae92000000000000000000000000000000000000000000000000000000008152600401611b5a9190611dd7565b60405180910390fd5b5f611b6f83835f611196565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614611be1575f6040517f73c6ac6e000000000000000000000000000000000000000000000000000000008152600401611bd89190611dd7565b60405180910390fd5b505050565b5f604051905090565b5f80fd5b5f80fd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b611c2b81611bf7565b8114611c35575f80fd5b50565b5f81359050611c4681611c22565b92915050565b5f60208284031215611c6157611c60611bef565b5b5f611c6e84828501611c38565b91505092915050565b5f8115159050919050565b611c8b81611c77565b82525050565b5f602082019050611ca45f830184611c82565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f611cec82611caa565b611cf68185611cb4565b9350611d06818560208601611cc4565b611d0f81611cd2565b840191505092915050565b5f6020820190508181035f830152611d328184611ce2565b905092915050565b5f819050919050565b611d4c81611d3a565b8114611d56575f80fd5b50565b5f81359050611d6781611d43565b92915050565b5f60208284031215611d8257611d81611bef565b5b5f611d8f84828501611d59565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f611dc182611d98565b9050919050565b611dd181611db7565b82525050565b5f602082019050611dea5f830184611dc8565b92915050565b611df981611db7565b8114611e03575f80fd5b50565b5f81359050611e1481611df0565b92915050565b5f8060408385031215611e3057611e2f611bef565b5b5f611e3d85828601611e06565b9250506020611e4e85828601611d59565b9150509250929050565b5f805f60608486031215611e6f57611e6e611bef565b5b5f611e7c86828701611e06565b9350506020611e8d86828701611e06565b9250506040611e9e86828701611d59565b9150509250925092565b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b611ee682611cd2565b810181811067ffffffffffffffff82111715611f0557611f04611eb0565b5b80604052505050565b5f611f17611be6565b9050611f238282611edd565b919050565b5f67ffffffffffffffff821115611f4257611f41611eb0565b5b611f4b82611cd2565b9050602081019050919050565b828183375f83830152505050565b5f611f78611f7384611f28565b611f0e565b905082815260208101848484011115611f9457611f93611eac565b5b611f9f848285611f58565b509392505050565b5f82601f830112611fbb57611fba611ea8565b5b8135611fcb848260208601611f66565b91505092915050565b5f805f8060808587031215611fec57611feb611bef565b5b5f611ff987828801611e06565b945050602061200a87828801611d59565b935050604085013567ffffffffffffffff81111561202b5761202a611bf3565b5b61203787828801611fa7565b925050606085013567ffffffffffffffff81111561205857612057611bf3565b5b61206487828801611fa7565b91505092959194509250565b5f6020828403121561208557612084611bef565b5b5f61209284828501611e06565b91505092915050565b6120a481611d3a565b82525050565b5f6020820190506120bd5f83018461209b565b92915050565b6120cc81611c77565b81146120d6575f80fd5b50565b5f813590506120e7816120c3565b92915050565b5f806040838503121561210357612102611bef565b5b5f61211085828601611e06565b9250506020612121858286016120d9565b9150509250929050565b5f67ffffffffffffffff82111561214557612144611eb0565b5b61214e82611cd2565b9050602081019050919050565b5f61216d6121688461212b565b611f0e565b90508281526020810184848401111561218957612188611eac565b5b612194848285611f58565b509392505050565b5f82601f8301126121b0576121af611ea8565b5b81356121c084826020860161215b565b91505092915050565b5f805f80608085870312156121e1576121e0611bef565b5b5f6121ee87828801611e06565b94505060206121ff87828801611e06565b935050604061221087828801611d59565b925050606085013567ffffffffffffffff81111561223157612230611bf3565b5b61223d8782880161219c565b91505092959194509250565b5f805f606084860312156122605761225f611bef565b5b5f61226d86828701611e06565b935050602061227e86828701611d59565b925050604084013567ffffffffffffffff81111561229f5761229e611bf3565b5b6122ab86828701611fa7565b9150509250925092565b5f80604083850312156122cb576122ca611bef565b5b5f6122d885828601611e06565b92505060206122e985828601611e06565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061233757607f821691505b60208210810361234a576123496122f3565b5b50919050565b5f6060820190506123635f830186611dc8565b612370602083018561209b565b61237d6040830184611dc8565b949350505050565b5f81519050612393816120c3565b92915050565b5f602082840312156123ae576123ad611bef565b5b5f6123bb84828501612385565b91505092915050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026124207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826123e5565b61242a86836123e5565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61246561246061245b84611d3a565b612442565b611d3a565b9050919050565b5f819050919050565b61247e8361244b565b61249261248a8261246c565b8484546123f1565b825550505050565b5f90565b6124a661249a565b6124b1818484612475565b505050565b5b818110156124d4576124c95f8261249e565b6001810190506124b7565b5050565b601f821115612519576124ea816123c4565b6124f3846123d6565b81016020851015612502578190505b61251661250e856123d6565b8301826124b6565b50505b505050565b5f82821c905092915050565b5f6125395f198460080261251e565b1980831691505092915050565b5f612551838361252a565b9150826002028217905092915050565b61256a82611caa565b67ffffffffffffffff81111561258357612582611eb0565b5b61258d8254612320565b6125988282856124d8565b5f60209050601f8311600181146125c9575f84156125b7578287015190505b6125c18582612546565b865550612628565b601f1984166125d7866123c4565b5f5b828110156125fe578489015182556001820191506020850194506020810190506125d9565b8683101561261b5784890151612617601f89168261252a565b8355505b6001600288020188555050505b505050505050565b5f6040820190506126435f830185611dc8565b612650602083018461209b565b9392505050565b5f81519050919050565b5f82825260208201905092915050565b5f61267b82612657565b6126858185612661565b9350612695818560208601611cc4565b61269e81611cd2565b840191505092915050565b5f6080820190506126bc5f830187611dc8565b6126c96020830186611dc8565b6126d6604083018561209b565b81810360608301526126e88184612671565b905095945050505050565b5f8151905061270181611c22565b92915050565b5f6020828403121561271c5761271b611bef565b5b5f612729848285016126f3565b9150509291505056fea264697066735822122000ab35d32bcda9fbbf255ff09d782986ad8bf39aa110b0588748e02d9656888564736f6c63430008190033",
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

// TokenOriginId is a free data retrieval call binding the contract method 0x6c8a5e77.
//
// Solidity: function tokenOriginId(uint256 tokenId) view returns(string)
func (_Ics721Erc721 *Ics721Erc721Caller) TokenOriginId(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Ics721Erc721.contract.Call(opts, &out, "tokenOriginId", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenOriginId is a free data retrieval call binding the contract method 0x6c8a5e77.
//
// Solidity: function tokenOriginId(uint256 tokenId) view returns(string)
func (_Ics721Erc721 *Ics721Erc721Session) TokenOriginId(tokenId *big.Int) (string, error) {
	return _Ics721Erc721.Contract.TokenOriginId(&_Ics721Erc721.CallOpts, tokenId)
}

// TokenOriginId is a free data retrieval call binding the contract method 0x6c8a5e77.
//
// Solidity: function tokenOriginId(uint256 tokenId) view returns(string)
func (_Ics721Erc721 *Ics721Erc721CallerSession) TokenOriginId(tokenId *big.Int) (string, error) {
	return _Ics721Erc721.Contract.TokenOriginId(&_Ics721Erc721.CallOpts, tokenId)
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

// Mint is a paid mutator transaction binding the contract method 0x2fb102cf.
//
// Solidity: function mint(address receiver, uint256 tokenId, string _tokenUri, string _tokenOriginId) returns()
func (_Ics721Erc721 *Ics721Erc721Transactor) Mint(opts *bind.TransactOpts, receiver common.Address, tokenId *big.Int, _tokenUri string, _tokenOriginId string) (*types.Transaction, error) {
	return _Ics721Erc721.contract.Transact(opts, "mint", receiver, tokenId, _tokenUri, _tokenOriginId)
}

// Mint is a paid mutator transaction binding the contract method 0x2fb102cf.
//
// Solidity: function mint(address receiver, uint256 tokenId, string _tokenUri, string _tokenOriginId) returns()
func (_Ics721Erc721 *Ics721Erc721Session) Mint(receiver common.Address, tokenId *big.Int, _tokenUri string, _tokenOriginId string) (*types.Transaction, error) {
	return _Ics721Erc721.Contract.Mint(&_Ics721Erc721.TransactOpts, receiver, tokenId, _tokenUri, _tokenOriginId)
}

// Mint is a paid mutator transaction binding the contract method 0x2fb102cf.
//
// Solidity: function mint(address receiver, uint256 tokenId, string _tokenUri, string _tokenOriginId) returns()
func (_Ics721Erc721 *Ics721Erc721TransactorSession) Mint(receiver common.Address, tokenId *big.Int, _tokenUri string, _tokenOriginId string) (*types.Transaction, error) {
	return _Ics721Erc721.Contract.Mint(&_Ics721Erc721.TransactOpts, receiver, tokenId, _tokenUri, _tokenOriginId)
}

// Mint0 is a paid mutator transaction binding the contract method 0xd3fc9864.
//
// Solidity: function mint(address receiver, uint256 tokenId, string _tokenUri) returns()
func (_Ics721Erc721 *Ics721Erc721Transactor) Mint0(opts *bind.TransactOpts, receiver common.Address, tokenId *big.Int, _tokenUri string) (*types.Transaction, error) {
	return _Ics721Erc721.contract.Transact(opts, "mint0", receiver, tokenId, _tokenUri)
}

// Mint0 is a paid mutator transaction binding the contract method 0xd3fc9864.
//
// Solidity: function mint(address receiver, uint256 tokenId, string _tokenUri) returns()
func (_Ics721Erc721 *Ics721Erc721Session) Mint0(receiver common.Address, tokenId *big.Int, _tokenUri string) (*types.Transaction, error) {
	return _Ics721Erc721.Contract.Mint0(&_Ics721Erc721.TransactOpts, receiver, tokenId, _tokenUri)
}

// Mint0 is a paid mutator transaction binding the contract method 0xd3fc9864.
//
// Solidity: function mint(address receiver, uint256 tokenId, string _tokenUri) returns()
func (_Ics721Erc721 *Ics721Erc721TransactorSession) Mint0(receiver common.Address, tokenId *big.Int, _tokenUri string) (*types.Transaction, error) {
	return _Ics721Erc721.Contract.Mint0(&_Ics721Erc721.TransactOpts, receiver, tokenId, _tokenUri)
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
