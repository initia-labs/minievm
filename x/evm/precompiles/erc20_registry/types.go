package erc20registry

import "github.com/ethereum/go-ethereum/common"

type RegisterArguments struct {
	Account common.Address `abi:"account"`
}

type IsRegisteredArguments struct {
	Account common.Address `abi:"account"`
}
