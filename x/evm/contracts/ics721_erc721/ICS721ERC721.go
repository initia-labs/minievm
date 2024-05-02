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
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_tokenUri\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_tokenOriginId\",\"type\":\"string\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_tokenUri\",\"type\":\"string\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenOriginId\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801562000010575f80fd5b50604051620029ab380380620029ab83398181016040528101906200003691906200022a565b8181815f9081620000489190620004e4565b5080600190816200005a9190620004e4565b5050503360065f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050620005c8565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6200010682620000be565b810181811067ffffffffffffffff82111715620001285762000127620000ce565b5b80604052505050565b5f6200013c620000a5565b90506200014a8282620000fb565b919050565b5f67ffffffffffffffff8211156200016c576200016b620000ce565b5b6200017782620000be565b9050602081019050919050565b5f5b83811015620001a357808201518184015260208101905062000186565b5f8484015250505050565b5f620001c4620001be846200014f565b62000131565b905082815260208101848484011115620001e357620001e2620000ba565b5b620001f084828562000184565b509392505050565b5f82601f8301126200020f576200020e620000b6565b5b815162000221848260208601620001ae565b91505092915050565b5f8060408385031215620002435762000242620000ae565b5b5f83015167ffffffffffffffff811115620002635762000262620000b2565b5b6200027185828601620001f8565b925050602083015167ffffffffffffffff811115620002955762000294620000b2565b5b620002a385828601620001f8565b9150509250929050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680620002fc57607f821691505b602082108103620003125762000311620002b7565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302620003767fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8262000339565b62000382868362000339565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f620003cc620003c6620003c0846200039a565b620003a3565b6200039a565b9050919050565b5f819050919050565b620003e783620003ac565b620003ff620003f682620003d3565b84845462000345565b825550505050565b5f90565b6200041562000407565b62000422818484620003dc565b505050565b5b8181101562000449576200043d5f826200040b565b60018101905062000428565b5050565b601f8211156200049857620004628162000318565b6200046d846200032a565b810160208510156200047d578190505b620004956200048c856200032a565b83018262000427565b50505b505050565b5f82821c905092915050565b5f620004ba5f19846008026200049d565b1980831691505092915050565b5f620004d48383620004a9565b9150826002028217905092915050565b620004ef82620002ad565b67ffffffffffffffff8111156200050b576200050a620000ce565b5b620005178254620002e4565b620005248282856200044d565b5f60209050601f8311600181146200055a575f841562000545578287015190505b620005518582620004c7565b865550620005c0565b601f1984166200056a8662000318565b5f5b8281101562000593578489015182556001820191506020850194506020810190506200056c565b86831015620005b35784890151620005af601f891682620004a9565b8355505b6001600288020188555050505b505050505050565b6123d580620005d65f395ff3fe608060405234801561000f575f80fd5b506004361061011f575f3560e01c80636c8a5e77116100ab578063b88d4fde1161006f578063b88d4fde14610315578063c87b56dd14610331578063d3fc986414610361578063e985e9c51461037d578063f2fde38b146103ad5761011f565b80636c8a5e771461025d57806370a082311461028d5780638da5cb5b146102bd57806395d89b41146102db578063a22cb465146102f95761011f565b806323b872dd116100f257806323b872dd146101bd5780632fb102cf146101d957806342842e0e146101f557806342966c68146102115780636352211e1461022d5761011f565b806301ffc9a71461012357806306fdde0314610153578063081812fc14610171578063095ea7b3146101a1575b5f80fd5b61013d600480360381019061013891906118de565b6103c9565b60405161014a9190611923565b60405180910390f35b61015b6104aa565b60405161016891906119c6565b60405180910390f35b61018b60048036038101906101869190611a19565b610539565b6040516101989190611a83565b60405180910390f35b6101bb60048036038101906101b69190611ac6565b610554565b005b6101d760048036038101906101d29190611b04565b61056a565b005b6101f360048036038101906101ee9190611c80565b610669565b005b61020f600480360381019061020a9190611b04565b61070f565b005b61022b60048036038101906102269190611a19565b61072e565b005b61024760048036038101906102429190611a19565b61079c565b6040516102549190611a83565b60405180910390f35b61027760048036038101906102729190611a19565b6107ad565b60405161028491906119c6565b60405180910390f35b6102a760048036038101906102a29190611d1c565b61084e565b6040516102b49190611d56565b60405180910390f35b6102c5610904565b6040516102d29190611a83565b60405180910390f35b6102e3610929565b6040516102f091906119c6565b60405180910390f35b610313600480360381019061030e9190611d99565b6109b9565b005b61032f600480360381019061032a9190611e75565b6109cf565b005b61034b60048036038101906103469190611a19565b6109f4565b60405161035891906119c6565b60405180910390f35b61037b60048036038101906103769190611ef5565b610a95565b005b61039760048036038101906103929190611f61565b610b0c565b6040516103a49190611923565b60405180910390f35b6103c760048036038101906103c29190611d1c565b610b9a565b005b5f7f80ac58cd000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916148061049357507f5b5e139f000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b806104a357506104a282610ce7565b5b9050919050565b60605f80546104b890611fcc565b80601f01602080910402602001604051908101604052809291908181526020018280546104e490611fcc565b801561052f5780601f106105065761010080835404028352916020019161052f565b820191905f5260205f20905b81548152906001019060200180831161051257829003601f168201915b5050505050905090565b5f61054382610d50565b5061054d82610dd6565b9050919050565b6105668282610561610e0f565b610e16565b5050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036105da575f6040517f64a0ae920000000000000000000000000000000000000000000000000000000081526004016105d19190611a83565b60405180910390fd5b5f6105ed83836105e8610e0f565b610e28565b90508373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610663578382826040517f64283d7b00000000000000000000000000000000000000000000000000000000815260040161065a93929190611ffc565b60405180910390fd5b50505050565b60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106c1575f80fd5b6106cb8484611033565b8160075f8581526020019081526020015f2090816106e991906121ce565b508060085f8581526020019081526020015f20908161070891906121ce565b5050505050565b61072983838360405180602001604052805f8152506109cf565b505050565b5f61073882610d50565b9050610745813384611050565b61078f57610751610e0f565b826040517f177e802f00000000000000000000000000000000000000000000000000000000815260040161078692919061229d565b60405180910390fd5b61079882611110565b5050565b5f6107a682610d50565b9050919050565b606060085f8381526020019081526020015f2080546107cb90611fcc565b80601f01602080910402602001604051908101604052809291908181526020018280546107f790611fcc565b80156108425780601f1061081957610100808354040283529160200191610842565b820191905f5260205f20905b81548152906001019060200180831161082557829003601f168201915b50505050509050919050565b5f8073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036108bf575f6040517f89c62b640000000000000000000000000000000000000000000000000000000081526004016108b69190611a83565b60405180910390fd5b60035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60606001805461093890611fcc565b80601f016020809104026020016040519081016040528092919081815260200182805461096490611fcc565b80156109af5780601f10610986576101008083540402835291602001916109af565b820191905f5260205f20905b81548152906001019060200180831161099257829003601f168201915b5050505050905090565b6109cb6109c4610e0f565b8383611192565b5050565b6109da84848461056a565b6109ee6109e5610e0f565b858585856112fb565b50505050565b606060075f8381526020019081526020015f208054610a1290611fcc565b80601f0160208091040260200160405190810160405280929190818152602001828054610a3e90611fcc565b8015610a895780601f10610a6057610100808354040283529160200191610a89565b820191905f5260205f20905b815481529060010190602001808311610a6c57829003601f168201915b50505050509050919050565b60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610aed575f80fd5b610b0783838360405180602001604052805f815250610669565b505050565b5f60055f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905092915050565b60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610bf2575f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610c29575f80fd5b8073ffffffffffffffffffffffffffffffffffffffff1660065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a38060065f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b5f80610d5b836114a7565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610dcd57826040517f7e273289000000000000000000000000000000000000000000000000000000008152600401610dc49190611d56565b60405180910390fd5b80915050919050565b5f60045f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b5f33905090565b610e2383838360016114e0565b505050565b5f80610e33846114a7565b90505f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614610e7457610e7381848661169f565b5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610eff57610eb35f855f806114e0565b600160035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825403925050819055505b5f73ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1614610f7e57600160035f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055505b8460025f8681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550838573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4809150509392505050565b61104c828260405180602001604052805f815250611762565b5050565b5f8073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415801561110757508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614806110c857506110c78484610b0c565b5b8061110657508273ffffffffffffffffffffffffffffffffffffffff166110ee83610dd6565b73ffffffffffffffffffffffffffffffffffffffff16145b5b90509392505050565b5f61111c5f835f610e28565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361118e57816040517f7e2732890000000000000000000000000000000000000000000000000000000081526004016111859190611d56565b60405180910390fd5b5050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361120257816040517f5b08ba180000000000000000000000000000000000000000000000000000000081526004016111f99190611a83565b60405180910390fd5b8060055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31836040516112ee9190611923565b60405180910390a3505050565b5f8373ffffffffffffffffffffffffffffffffffffffff163b11156114a0578273ffffffffffffffffffffffffffffffffffffffff1663150b7a02868685856040518563ffffffff1660e01b81526004016113599493929190612316565b6020604051808303815f875af192505050801561139457506040513d601f19601f820116820180604052508101906113919190612374565b60015b611415573d805f81146113c2576040519150601f19603f3d011682016040523d82523d5f602084013e6113c7565b606091505b505f81510361140d57836040517f64a0ae920000000000000000000000000000000000000000000000000000000081526004016114049190611a83565b60405180910390fd5b805181602001fd5b63150b7a0260e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161461149e57836040517f64a0ae920000000000000000000000000000000000000000000000000000000081526004016114959190611a83565b60405180910390fd5b505b5050505050565b5f60025f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b808061151857505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b1561164a575f61152784610d50565b90505f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415801561159157508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b80156115a457506115a28184610b0c565b155b156115e657826040517fa9fbf51f0000000000000000000000000000000000000000000000000000000081526004016115dd9190611a83565b60405180910390fd5b811561164857838573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45b505b8360045f8581526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050565b6116aa838383611050565b61175d575f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361171e57806040517f7e2732890000000000000000000000000000000000000000000000000000000081526004016117159190611d56565b60405180910390fd5b81816040517f177e802f00000000000000000000000000000000000000000000000000000000815260040161175492919061229d565b60405180910390fd5b505050565b61176c8383611785565b611780611777610e0f565b5f8585856112fb565b505050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036117f5575f6040517f64a0ae920000000000000000000000000000000000000000000000000000000081526004016117ec9190611a83565b60405180910390fd5b5f61180183835f610e28565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614611873575f6040517f73c6ac6e00000000000000000000000000000000000000000000000000000000815260040161186a9190611a83565b60405180910390fd5b505050565b5f604051905090565b5f80fd5b5f80fd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6118bd81611889565b81146118c7575f80fd5b50565b5f813590506118d8816118b4565b92915050565b5f602082840312156118f3576118f2611881565b5b5f611900848285016118ca565b91505092915050565b5f8115159050919050565b61191d81611909565b82525050565b5f6020820190506119365f830184611914565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015611973578082015181840152602081019050611958565b5f8484015250505050565b5f601f19601f8301169050919050565b5f6119988261193c565b6119a28185611946565b93506119b2818560208601611956565b6119bb8161197e565b840191505092915050565b5f6020820190508181035f8301526119de818461198e565b905092915050565b5f819050919050565b6119f8816119e6565b8114611a02575f80fd5b50565b5f81359050611a13816119ef565b92915050565b5f60208284031215611a2e57611a2d611881565b5b5f611a3b84828501611a05565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f611a6d82611a44565b9050919050565b611a7d81611a63565b82525050565b5f602082019050611a965f830184611a74565b92915050565b611aa581611a63565b8114611aaf575f80fd5b50565b5f81359050611ac081611a9c565b92915050565b5f8060408385031215611adc57611adb611881565b5b5f611ae985828601611ab2565b9250506020611afa85828601611a05565b9150509250929050565b5f805f60608486031215611b1b57611b1a611881565b5b5f611b2886828701611ab2565b9350506020611b3986828701611ab2565b9250506040611b4a86828701611a05565b9150509250925092565b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b611b928261197e565b810181811067ffffffffffffffff82111715611bb157611bb0611b5c565b5b80604052505050565b5f611bc3611878565b9050611bcf8282611b89565b919050565b5f67ffffffffffffffff821115611bee57611bed611b5c565b5b611bf78261197e565b9050602081019050919050565b828183375f83830152505050565b5f611c24611c1f84611bd4565b611bba565b905082815260208101848484011115611c4057611c3f611b58565b5b611c4b848285611c04565b509392505050565b5f82601f830112611c6757611c66611b54565b5b8135611c77848260208601611c12565b91505092915050565b5f805f8060808587031215611c9857611c97611881565b5b5f611ca587828801611ab2565b9450506020611cb687828801611a05565b935050604085013567ffffffffffffffff811115611cd757611cd6611885565b5b611ce387828801611c53565b925050606085013567ffffffffffffffff811115611d0457611d03611885565b5b611d1087828801611c53565b91505092959194509250565b5f60208284031215611d3157611d30611881565b5b5f611d3e84828501611ab2565b91505092915050565b611d50816119e6565b82525050565b5f602082019050611d695f830184611d47565b92915050565b611d7881611909565b8114611d82575f80fd5b50565b5f81359050611d9381611d6f565b92915050565b5f8060408385031215611daf57611dae611881565b5b5f611dbc85828601611ab2565b9250506020611dcd85828601611d85565b9150509250929050565b5f67ffffffffffffffff821115611df157611df0611b5c565b5b611dfa8261197e565b9050602081019050919050565b5f611e19611e1484611dd7565b611bba565b905082815260208101848484011115611e3557611e34611b58565b5b611e40848285611c04565b509392505050565b5f82601f830112611e5c57611e5b611b54565b5b8135611e6c848260208601611e07565b91505092915050565b5f805f8060808587031215611e8d57611e8c611881565b5b5f611e9a87828801611ab2565b9450506020611eab87828801611ab2565b9350506040611ebc87828801611a05565b925050606085013567ffffffffffffffff811115611edd57611edc611885565b5b611ee987828801611e48565b91505092959194509250565b5f805f60608486031215611f0c57611f0b611881565b5b5f611f1986828701611ab2565b9350506020611f2a86828701611a05565b925050604084013567ffffffffffffffff811115611f4b57611f4a611885565b5b611f5786828701611c53565b9150509250925092565b5f8060408385031215611f7757611f76611881565b5b5f611f8485828601611ab2565b9250506020611f9585828601611ab2565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680611fe357607f821691505b602082108103611ff657611ff5611f9f565b5b50919050565b5f60608201905061200f5f830186611a74565b61201c6020830185611d47565b6120296040830184611a74565b949350505050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261208d7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82612052565b6120978683612052565b95508019841693508086168417925050509392505050565b5f819050919050565b5f6120d26120cd6120c8846119e6565b6120af565b6119e6565b9050919050565b5f819050919050565b6120eb836120b8565b6120ff6120f7826120d9565b84845461205e565b825550505050565b5f90565b612113612107565b61211e8184846120e2565b505050565b5b81811015612141576121365f8261210b565b600181019050612124565b5050565b601f8211156121865761215781612031565b61216084612043565b8101602085101561216f578190505b61218361217b85612043565b830182612123565b50505b505050565b5f82821c905092915050565b5f6121a65f198460080261218b565b1980831691505092915050565b5f6121be8383612197565b9150826002028217905092915050565b6121d78261193c565b67ffffffffffffffff8111156121f0576121ef611b5c565b5b6121fa8254611fcc565b612205828285612145565b5f60209050601f831160018114612236575f8415612224578287015190505b61222e85826121b3565b865550612295565b601f19841661224486612031565b5f5b8281101561226b57848901518255600182019150602085019450602081019050612246565b868310156122885784890151612284601f891682612197565b8355505b6001600288020188555050505b505050505050565b5f6040820190506122b05f830185611a74565b6122bd6020830184611d47565b9392505050565b5f81519050919050565b5f82825260208201905092915050565b5f6122e8826122c4565b6122f281856122ce565b9350612302818560208601611956565b61230b8161197e565b840191505092915050565b5f6080820190506123295f830187611a74565b6123366020830186611a74565b6123436040830185611d47565b818103606083015261235581846122de565b905095945050505050565b5f8151905061236e816118b4565b92915050565b5f6020828403121561238957612388611881565b5b5f61239684828501612360565b9150509291505056fea26469706673582212207cc5afdb042e85f38f661674372347e93a09868e6857d22a0a49908f7b5a374164736f6c63430008180033",
}

// Ics721Erc721ABI is the input ABI used to generate the binding from.
// Deprecated: Use Ics721Erc721MetaData.ABI instead.
var Ics721Erc721ABI = Ics721Erc721MetaData.ABI

// Ics721Erc721Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Ics721Erc721MetaData.Bin instead.
var Ics721Erc721Bin = Ics721Erc721MetaData.Bin

// DeployIcs721Erc721 deploys a new Ethereum contract, binding an instance of Ics721Erc721 to it.
func DeployIcs721Erc721(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string) (common.Address, *types.Transaction, *Ics721Erc721, error) {
	parsed, err := Ics721Erc721MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Ics721Erc721Bin), backend, name_, symbol_)
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
