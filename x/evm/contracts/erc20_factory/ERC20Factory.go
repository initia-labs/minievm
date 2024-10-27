// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20_factory

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

// Erc20FactoryMetaData contains all meta data concerning the Erc20Factory contract.
var Erc20FactoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"erc20\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC20Created\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"createERC20\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b50612a858061001c5f395ff3fe608060405234801561000f575f80fd5b5060043610610029575f3560e01c806306ef1a861461002d575b5f80fd5b6100476004803603810190610042919061036c565b61005d565b6040516100549190610433565b60405180910390f35b5f8084848460405161006e906101dc565b61007a939291906104bb565b604051809103905ff080158015610093573d5f803e3d5ffd5b50905060f273ffffffffffffffffffffffffffffffffffffffff1663d126274a826040518263ffffffff1660e01b81526004016100d09190610433565b6020604051808303815f875af11580156100ec573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906101109190610533565b508073ffffffffffffffffffffffffffffffffffffffff1663f2fde38b336040518263ffffffff1660e01b815260040161014a9190610433565b5f604051808303815f87803b158015610161575f80fd5b505af1158015610173573d5f803e3d5ffd5b505050503373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f85e892981b234101136bc30081e0a5c44345bebc0940193230c20a43b279e2d160405160405180910390a3809150509392505050565b6124f18061055f83390190565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61024882610202565b810181811067ffffffffffffffff8211171561026757610266610212565b5b80604052505050565b5f6102796101e9565b9050610285828261023f565b919050565b5f67ffffffffffffffff8211156102a4576102a3610212565b5b6102ad82610202565b9050602081019050919050565b828183375f83830152505050565b5f6102da6102d58461028a565b610270565b9050828152602081018484840111156102f6576102f56101fe565b5b6103018482856102ba565b509392505050565b5f82601f83011261031d5761031c6101fa565b5b813561032d8482602086016102c8565b91505092915050565b5f60ff82169050919050565b61034b81610336565b8114610355575f80fd5b50565b5f8135905061036681610342565b92915050565b5f805f60608486031215610383576103826101f2565b5b5f84013567ffffffffffffffff8111156103a05761039f6101f6565b5b6103ac86828701610309565b935050602084013567ffffffffffffffff8111156103cd576103cc6101f6565b5b6103d986828701610309565b92505060406103ea86828701610358565b9150509250925092565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61041d826103f4565b9050919050565b61042d81610413565b82525050565b5f6020820190506104465f830184610424565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f61047e8261044c565b6104888185610456565b9350610498818560208601610466565b6104a181610202565b840191505092915050565b6104b581610336565b82525050565b5f6060820190508181035f8301526104d38186610474565b905081810360208301526104e78185610474565b90506104f660408301846104ac565b949350505050565b5f8115159050919050565b610512816104fe565b811461051c575f80fd5b50565b5f8151905061052d81610509565b92915050565b5f60208284031215610548576105476101f2565b5b5f6105558482850161051f565b9150509291505056fe608060405234801561000f575f80fd5b506040516124f13803806124f183398181016040528101906100319190610235565b335f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550826003908161007f91906104ca565b50816004908161008f91906104ca565b508060055f6101000a81548160ff021916908360ff160217905550505050610599565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610111826100cb565b810181811067ffffffffffffffff821117156101305761012f6100db565b5b80604052505050565b5f6101426100b2565b905061014e8282610108565b919050565b5f67ffffffffffffffff82111561016d5761016c6100db565b5b610176826100cb565b9050602081019050919050565b8281835e5f83830152505050565b5f6101a361019e84610153565b610139565b9050828152602081018484840111156101bf576101be6100c7565b5b6101ca848285610183565b509392505050565b5f82601f8301126101e6576101e56100c3565b5b81516101f6848260208601610191565b91505092915050565b5f60ff82169050919050565b610214816101ff565b811461021e575f80fd5b50565b5f8151905061022f8161020b565b92915050565b5f805f6060848603121561024c5761024b6100bb565b5b5f84015167ffffffffffffffff811115610269576102686100bf565b5b610275868287016101d2565b935050602084015167ffffffffffffffff811115610296576102956100bf565b5b6102a2868287016101d2565b92505060406102b386828701610221565b9150509250925092565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061030b57607f821691505b60208210810361031e5761031d6102c7565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026103807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610345565b61038a8683610345565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6103ce6103c96103c4846103a2565b6103ab565b6103a2565b9050919050565b5f819050919050565b6103e7836103b4565b6103fb6103f3826103d5565b848454610351565b825550505050565b5f90565b61040f610403565b61041a8184846103de565b505050565b5b8181101561043d576104325f82610407565b600181019050610420565b5050565b601f8211156104825761045381610324565b61045c84610336565b8101602085101561046b578190505b61047f61047785610336565b83018261041f565b50505b505050565b5f82821c905092915050565b5f6104a25f1984600802610487565b1980831691505092915050565b5f6104ba8383610493565b9150826002028217905092915050565b6104d3826102bd565b67ffffffffffffffff8111156104ec576104eb6100db565b5b6104f682546102f4565b610501828285610441565b5f60209050601f831160018114610532575f8415610520578287015190505b61052a85826104af565b865550610591565b601f19841661054086610324565b5f5b8281101561056757848901518255600182019150602085019450602081019050610542565b868310156105845784890151610580601f891682610493565b8355505b6001600288020188555050505b505050505050565b611f4b806105a65f395ff3fe608060405234801561000f575f80fd5b5060043610610114575f3560e01c806342966c68116100a057806395d89b411161006f57806395d89b41146102f0578063a9059cbb1461030e578063dd62ed3e1461033e578063f2fde38b1461036e578063fe1195ec1461038a57610114565b806342966c681461025657806370a082311461027257806379cc6790146102a25780638da5cb5b146102d257610114565b80631988513b116100e75780631988513b146101b457806323b872dd146101d05780632d688ca814610200578063313ce5671461021c57806340c10f191461023a57610114565b806301ffc9a71461011857806306fdde0314610148578063095ea7b31461016657806318160ddd14610196575b5f80fd5b610132600480360381019061012d91906116b1565b6103a6565b60405161013f91906116f6565b60405180910390f35b61015061041f565b60405161015d919061177f565b60405180910390f35b610180600480360381019061017b919061182c565b6104ab565b60405161018d91906116f6565b60405180910390f35b61019e610598565b6040516101ab9190611879565b60405180910390f35b6101ce60048036038101906101c99190611892565b61059e565b005b6101ea60048036038101906101e59190611892565b61061d565b6040516101f791906116f6565b60405180910390f35b61021a6004803603810190610215919061182c565b610838565b005b6102246108b5565b60405161023191906118fd565b60405180910390f35b610254600480360381019061024f919061182c565b6108c7565b005b610270600480360381019061026b9190611916565b6109e6565b005b61028c60048036038101906102879190611941565b610aae565b6040516102999190611879565b60405180910390f35b6102bc60048036038101906102b7919061182c565b610ac3565b6040516102c991906116f6565b60405180910390f35b6102da610cdc565b6040516102e7919061197b565b60405180910390f35b6102f8610cff565b604051610305919061177f565b60405180910390f35b6103286004803603810190610323919061182c565b610d8b565b60405161033591906116f6565b60405180910390f35b61035860048036038101906103539190611994565b610e5c565b6040516103659190611879565b60405180910390f35b61038860048036038101906103839190611941565b610e7c565b005b6103a4600480360381019061039f919061182c565b610fc4565b005b5f7f8da6da19000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161480610418575061041782611041565b5b9050919050565b6003805461042c906119ff565b80601f0160208091040260200160405190810160405280929190818152602001828054610458906119ff565b80156104a35780601f1061047a576101008083540402835291602001916104a3565b820191905f5260205f20905b81548152906001019060200180831161048657829003601f168201915b505050505081565b5f8160025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516105869190611879565b60405180910390a36001905092915050565b60065481565b600173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461060d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161060490611a79565b60405180910390fd5b6106188383836110aa565b505050565b5f8260f173ffffffffffffffffffffffffffffffffffffffff1663f2af9ac9826040518263ffffffff1660e01b8152600401610659919061197b565b602060405180830381865afa158015610674573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106989190611ac1565b156106d8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106cf90611b5c565b60405180910390fd5b8260025f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015610793576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161078a90611bea565b60405180910390fd5b8260025f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461081a9190611c35565b9250508190555061082c8585856110aa565b60019150509392505050565b600173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146108a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161089e90611a79565b60405180910390fd5b6108b18282611335565b5050565b60055f9054906101000a900460ff1681565b8160f173ffffffffffffffffffffffffffffffffffffffff1663f2af9ac9826040518263ffffffff1660e01b8152600401610902919061197b565b602060405180830381865afa15801561091d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109419190611ac1565b15610981576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097890611cb2565b60405180910390fd5b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146109d7575f80fd5b6109e18383611335565b505050565b3360f173ffffffffffffffffffffffffffffffffffffffff166360dc402f826040518263ffffffff1660e01b8152600401610a21919061197b565b602060405180830381865afa158015610a3c573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a609190611ac1565b15610aa0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a9790611d1a565b60405180910390fd5b610aaa3383611504565b5050565b6001602052805f5260405f205f915090505481565b5f8260f173ffffffffffffffffffffffffffffffffffffffff166360dc402f826040518263ffffffff1660e01b8152600401610aff919061197b565b602060405180830381865afa158015610b1a573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b3e9190611ac1565b15610b7e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b7590611d1a565b60405180910390fd5b8260025f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015610c39576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c3090611da8565b60405180910390fd5b8260025f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610cc09190611c35565b92505081905550610cd18484611504565b600191505092915050565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60048054610d0c906119ff565b80601f0160208091040260200160405190810160405280929190818152602001828054610d38906119ff565b8015610d835780601f10610d5a57610100808354040283529160200191610d83565b820191905f5260205f20905b815481529060010190602001808311610d6657829003601f168201915b505050505081565b5f8260f173ffffffffffffffffffffffffffffffffffffffff1663f2af9ac9826040518263ffffffff1660e01b8152600401610dc7919061197b565b602060405180830381865afa158015610de2573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e069190611ac1565b15610e46576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e3d90611b5c565b60405180910390fd5b610e513385856110aa565b600191505092915050565b6002602052815f5260405f20602052805f5260405f205f91509150505481565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610ed2575f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610f09575f80fd5b8073ffffffffffffffffffffffffffffffffffffffff165f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3805f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611033576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161102a90611a79565b60405180910390fd5b61103d8282611504565b5050565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b8160f273ffffffffffffffffffffffffffffffffffffffff16634e25ab64826040518263ffffffff1660e01b81526004016110e5919061197b565b602060405180830381865afa158015611100573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111249190611ac1565b6111a45760f273ffffffffffffffffffffffffffffffffffffffff1663ceeae52a826040518263ffffffff1660e01b8152600401611162919061197b565b6020604051808303815f875af115801561117e573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111a29190611ac1565b505b8160015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015611224576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161121b90611e36565b60405180910390fd5b8160015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546112709190611c35565b925050819055508160015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546112c39190611e54565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516113279190611879565b60405180910390a350505050565b8160f273ffffffffffffffffffffffffffffffffffffffff16634e25ab64826040518263ffffffff1660e01b8152600401611370919061197b565b602060405180830381865afa15801561138b573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906113af9190611ac1565b61142f5760f273ffffffffffffffffffffffffffffffffffffffff1663ceeae52a826040518263ffffffff1660e01b81526004016113ed919061197b565b6020604051808303815f875af1158015611409573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061142d9190611ac1565b505b8160015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461147b9190611e54565b925050819055508160065f8282546114939190611e54565b925050819055508273ffffffffffffffffffffffffffffffffffffffff165f73ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516114f79190611879565b60405180910390a3505050565b8060015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015611584576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161157b90611ef7565b60405180910390fd5b8060015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546115d09190611c35565b925050819055508060065f8282546115e89190611c35565b925050819055505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161164c9190611879565b60405180910390a35050565b5f80fd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6116908161165c565b811461169a575f80fd5b50565b5f813590506116ab81611687565b92915050565b5f602082840312156116c6576116c5611658565b5b5f6116d38482850161169d565b91505092915050565b5f8115159050919050565b6116f0816116dc565b82525050565b5f6020820190506117095f8301846116e7565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f6117518261170f565b61175b8185611719565b935061176b818560208601611729565b61177481611737565b840191505092915050565b5f6020820190508181035f8301526117978184611747565b905092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6117c88261179f565b9050919050565b6117d8816117be565b81146117e2575f80fd5b50565b5f813590506117f3816117cf565b92915050565b5f819050919050565b61180b816117f9565b8114611815575f80fd5b50565b5f8135905061182681611802565b92915050565b5f806040838503121561184257611841611658565b5b5f61184f858286016117e5565b925050602061186085828601611818565b9150509250929050565b611873816117f9565b82525050565b5f60208201905061188c5f83018461186a565b92915050565b5f805f606084860312156118a9576118a8611658565b5b5f6118b6868287016117e5565b93505060206118c7868287016117e5565b92505060406118d886828701611818565b9150509250925092565b5f60ff82169050919050565b6118f7816118e2565b82525050565b5f6020820190506119105f8301846118ee565b92915050565b5f6020828403121561192b5761192a611658565b5b5f61193884828501611818565b91505092915050565b5f6020828403121561195657611955611658565b5b5f611963848285016117e5565b91505092915050565b611975816117be565b82525050565b5f60208201905061198e5f83018461196c565b92915050565b5f80604083850312156119aa576119a9611658565b5b5f6119b7858286016117e5565b92505060206119c8858286016117e5565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680611a1657607f821691505b602082108103611a2957611a286119d2565b5b50919050565b7f45524332303a2063616c6c6572206973206e6f742074686520636861696e00005f82015250565b5f611a63601e83611719565b9150611a6e82611a2f565b602082019050919050565b5f6020820190508181035f830152611a9081611a57565b9050919050565b611aa0816116dc565b8114611aaa575f80fd5b50565b5f81519050611abb81611a97565b92915050565b5f60208284031215611ad657611ad5611658565b5b5f611ae384828501611aad565b91505092915050565b7f45524332303a207472616e7366657220746f20626c6f636b65642061646472655f8201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b5f611b46602283611719565b9150611b5182611aec565b604082019050919050565b5f6020820190508181035f830152611b7381611b3a565b9050919050565b7f45524332303a207472616e7366657220616d6f756e74206578636565647320615f8201527f6c6c6f77616e6365000000000000000000000000000000000000000000000000602082015250565b5f611bd4602883611719565b9150611bdf82611b7a565b604082019050919050565b5f6020820190508181035f830152611c0181611bc8565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611c3f826117f9565b9150611c4a836117f9565b9250828203905081811115611c6257611c61611c08565b5b92915050565b7f45524332303a206d696e7420746f20626c6f636b6564206164647265737300005f82015250565b5f611c9c601e83611719565b9150611ca782611c68565b602082019050919050565b5f6020820190508181035f830152611cc981611c90565b9050919050565b7f45524332303a206275726e2066726f6d206d6f64756c652061646472657373005f82015250565b5f611d04601f83611719565b9150611d0f82611cd0565b602082019050919050565b5f6020820190508181035f830152611d3181611cf8565b9050919050565b7f45524332303a206275726e20616d6f756e74206578636565647320616c6c6f775f8201527f616e636500000000000000000000000000000000000000000000000000000000602082015250565b5f611d92602483611719565b9150611d9d82611d38565b604082019050919050565b5f6020820190508181035f830152611dbf81611d86565b9050919050565b7f45524332303a207472616e7366657220616d6f756e74206578636565647320625f8201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b5f611e20602683611719565b9150611e2b82611dc6565b604082019050919050565b5f6020820190508181035f830152611e4d81611e14565b9050919050565b5f611e5e826117f9565b9150611e69836117f9565b9250828201905080821115611e8157611e80611c08565b5b92915050565b7f45524332303a206275726e20616d6f756e7420657863656564732062616c616e5f8201527f6365000000000000000000000000000000000000000000000000000000000000602082015250565b5f611ee1602283611719565b9150611eec82611e87565b604082019050919050565b5f6020820190508181035f830152611f0e81611ed5565b905091905056fea264697066735822122002e9ad3a5509c74fe59dd4b570d5950dea623e81de69a8349bb3a7b6861f41c164736f6c63430008190033a2646970667358221220803ccecb32a334711235eb8ba3541387923d08549658060d4efb4d070a1ab80f64736f6c63430008190033",
}

// Erc20FactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc20FactoryMetaData.ABI instead.
var Erc20FactoryABI = Erc20FactoryMetaData.ABI

// Erc20FactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Erc20FactoryMetaData.Bin instead.
var Erc20FactoryBin = Erc20FactoryMetaData.Bin

// DeployErc20Factory deploys a new Ethereum contract, binding an instance of Erc20Factory to it.
func DeployErc20Factory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Erc20Factory, error) {
	parsed, err := Erc20FactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Erc20FactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Erc20Factory{Erc20FactoryCaller: Erc20FactoryCaller{contract: contract}, Erc20FactoryTransactor: Erc20FactoryTransactor{contract: contract}, Erc20FactoryFilterer: Erc20FactoryFilterer{contract: contract}}, nil
}

// Erc20Factory is an auto generated Go binding around an Ethereum contract.
type Erc20Factory struct {
	Erc20FactoryCaller     // Read-only binding to the contract
	Erc20FactoryTransactor // Write-only binding to the contract
	Erc20FactoryFilterer   // Log filterer for contract events
}

// Erc20FactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type Erc20FactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20FactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc20FactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20FactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc20FactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20FactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc20FactorySession struct {
	Contract     *Erc20Factory     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc20FactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc20FactoryCallerSession struct {
	Contract *Erc20FactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// Erc20FactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc20FactoryTransactorSession struct {
	Contract     *Erc20FactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// Erc20FactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type Erc20FactoryRaw struct {
	Contract *Erc20Factory // Generic contract binding to access the raw methods on
}

// Erc20FactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc20FactoryCallerRaw struct {
	Contract *Erc20FactoryCaller // Generic read-only contract binding to access the raw methods on
}

// Erc20FactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc20FactoryTransactorRaw struct {
	Contract *Erc20FactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErc20Factory creates a new instance of Erc20Factory, bound to a specific deployed contract.
func NewErc20Factory(address common.Address, backend bind.ContractBackend) (*Erc20Factory, error) {
	contract, err := bindErc20Factory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc20Factory{Erc20FactoryCaller: Erc20FactoryCaller{contract: contract}, Erc20FactoryTransactor: Erc20FactoryTransactor{contract: contract}, Erc20FactoryFilterer: Erc20FactoryFilterer{contract: contract}}, nil
}

// NewErc20FactoryCaller creates a new read-only instance of Erc20Factory, bound to a specific deployed contract.
func NewErc20FactoryCaller(address common.Address, caller bind.ContractCaller) (*Erc20FactoryCaller, error) {
	contract, err := bindErc20Factory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20FactoryCaller{contract: contract}, nil
}

// NewErc20FactoryTransactor creates a new write-only instance of Erc20Factory, bound to a specific deployed contract.
func NewErc20FactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*Erc20FactoryTransactor, error) {
	contract, err := bindErc20Factory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20FactoryTransactor{contract: contract}, nil
}

// NewErc20FactoryFilterer creates a new log filterer instance of Erc20Factory, bound to a specific deployed contract.
func NewErc20FactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*Erc20FactoryFilterer, error) {
	contract, err := bindErc20Factory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc20FactoryFilterer{contract: contract}, nil
}

// bindErc20Factory binds a generic wrapper to an already deployed contract.
func bindErc20Factory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Erc20FactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20Factory *Erc20FactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20Factory.Contract.Erc20FactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20Factory *Erc20FactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20Factory.Contract.Erc20FactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20Factory *Erc20FactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20Factory.Contract.Erc20FactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20Factory *Erc20FactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20Factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20Factory *Erc20FactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20Factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20Factory *Erc20FactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20Factory.Contract.contract.Transact(opts, method, params...)
}

// CreateERC20 is a paid mutator transaction binding the contract method 0x06ef1a86.
//
// Solidity: function createERC20(string name, string symbol, uint8 decimals) returns(address)
func (_Erc20Factory *Erc20FactoryTransactor) CreateERC20(opts *bind.TransactOpts, name string, symbol string, decimals uint8) (*types.Transaction, error) {
	return _Erc20Factory.contract.Transact(opts, "createERC20", name, symbol, decimals)
}

// CreateERC20 is a paid mutator transaction binding the contract method 0x06ef1a86.
//
// Solidity: function createERC20(string name, string symbol, uint8 decimals) returns(address)
func (_Erc20Factory *Erc20FactorySession) CreateERC20(name string, symbol string, decimals uint8) (*types.Transaction, error) {
	return _Erc20Factory.Contract.CreateERC20(&_Erc20Factory.TransactOpts, name, symbol, decimals)
}

// CreateERC20 is a paid mutator transaction binding the contract method 0x06ef1a86.
//
// Solidity: function createERC20(string name, string symbol, uint8 decimals) returns(address)
func (_Erc20Factory *Erc20FactoryTransactorSession) CreateERC20(name string, symbol string, decimals uint8) (*types.Transaction, error) {
	return _Erc20Factory.Contract.CreateERC20(&_Erc20Factory.TransactOpts, name, symbol, decimals)
}

// Erc20FactoryERC20CreatedIterator is returned from FilterERC20Created and is used to iterate over the raw logs and unpacked data for ERC20Created events raised by the Erc20Factory contract.
type Erc20FactoryERC20CreatedIterator struct {
	Event *Erc20FactoryERC20Created // Event containing the contract specifics and raw log

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
func (it *Erc20FactoryERC20CreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20FactoryERC20Created)
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
		it.Event = new(Erc20FactoryERC20Created)
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
func (it *Erc20FactoryERC20CreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20FactoryERC20CreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20FactoryERC20Created represents a ERC20Created event raised by the Erc20Factory contract.
type Erc20FactoryERC20Created struct {
	Erc20 common.Address
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterERC20Created is a free log retrieval operation binding the contract event 0x85e892981b234101136bc30081e0a5c44345bebc0940193230c20a43b279e2d1.
//
// Solidity: event ERC20Created(address indexed erc20, address indexed owner)
func (_Erc20Factory *Erc20FactoryFilterer) FilterERC20Created(opts *bind.FilterOpts, erc20 []common.Address, owner []common.Address) (*Erc20FactoryERC20CreatedIterator, error) {

	var erc20Rule []interface{}
	for _, erc20Item := range erc20 {
		erc20Rule = append(erc20Rule, erc20Item)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Erc20Factory.contract.FilterLogs(opts, "ERC20Created", erc20Rule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &Erc20FactoryERC20CreatedIterator{contract: _Erc20Factory.contract, event: "ERC20Created", logs: logs, sub: sub}, nil
}

// WatchERC20Created is a free log subscription operation binding the contract event 0x85e892981b234101136bc30081e0a5c44345bebc0940193230c20a43b279e2d1.
//
// Solidity: event ERC20Created(address indexed erc20, address indexed owner)
func (_Erc20Factory *Erc20FactoryFilterer) WatchERC20Created(opts *bind.WatchOpts, sink chan<- *Erc20FactoryERC20Created, erc20 []common.Address, owner []common.Address) (event.Subscription, error) {

	var erc20Rule []interface{}
	for _, erc20Item := range erc20 {
		erc20Rule = append(erc20Rule, erc20Item)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Erc20Factory.contract.WatchLogs(opts, "ERC20Created", erc20Rule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20FactoryERC20Created)
				if err := _Erc20Factory.contract.UnpackLog(event, "ERC20Created", log); err != nil {
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

// ParseERC20Created is a log parse operation binding the contract event 0x85e892981b234101136bc30081e0a5c44345bebc0940193230c20a43b279e2d1.
//
// Solidity: event ERC20Created(address indexed erc20, address indexed owner)
func (_Erc20Factory *Erc20FactoryFilterer) ParseERC20Created(log types.Log) (*Erc20FactoryERC20Created, error) {
	event := new(Erc20FactoryERC20Created)
	if err := _Erc20Factory.contract.UnpackLog(event, "ERC20Created", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
