package types

import "github.com/ethereum/go-ethereum/common"

// private type creates an interface key for Context that cannot be accessed by any other package
type contextKey int

const (
	ContextKeyLoadDecimals contextKey = iota
)

// Fee is a struct that represents a fee denom and a contract address
type Fee struct {
	denom    string
	contract common.Address
	decimals uint8
}

func NewFee(denom string, contract common.Address, decimals uint8) Fee {
	return Fee{
		denom:    denom,
		contract: contract,
		decimals: decimals,
	}
}

func (f Fee) Denom() string {
	return f.denom
}

func (f Fee) Contract() common.Address {
	return f.contract
}

func (f Fee) HasContract() bool {
	return f.contract != common.Address{}
}

func (f Fee) Decimals() uint8 {
	return f.decimals
}
