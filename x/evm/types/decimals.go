package types

import "math/big"

const (
	EtherDecimals = uint8(18)
	GweiDecimals  = uint8(9)
)

func ToEthersUint(decimals uint8, val *big.Int) *big.Int {
	if decimals == EtherDecimals {
		return new(big.Int).Set(val)
	}

	if decimals > EtherDecimals {
		decimalDiff := decimals - EtherDecimals
		exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimalDiff)), nil)
		return new(big.Int).Div(val, exp)
	}

	decimalDiff := EtherDecimals - decimals

	exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimalDiff)), nil)
	return new(big.Int).Mul(val, exp)
}

func ToGweiUint(decimals uint8, val *big.Int) *big.Int {
	if decimals == GweiDecimals {
		return new(big.Int).Set(val)
	}

	if decimals > GweiDecimals {
		decimalDiff := decimals - GweiDecimals
		exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimalDiff)), nil)
		return new(big.Int).Div(val, exp)
	}

	decimalDiff := GweiDecimals - decimals

	exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimalDiff)), nil)
	return new(big.Int).Mul(val, exp)
}

func FromGweiUnit(decimals uint8, val *big.Int) *big.Int {
	if decimals == GweiDecimals {
		return new(big.Int).Set(val)
	}

	if decimals > GweiDecimals {
		decimalDiff := decimals - GweiDecimals
		exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimalDiff)), nil)
		return new(big.Int).Mul(val, exp)
	}

	decimalDiff := GweiDecimals - decimals

	exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimalDiff)), nil)
	return new(big.Int).Div(val, exp)
}

func FromEthersUnit(decimals uint8, val *big.Int) *big.Int {
	if decimals == EtherDecimals {
		return new(big.Int).Set(val)
	}

	if decimals > EtherDecimals {
		decimalDiff := decimals - EtherDecimals
		exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimalDiff)), nil)
		return new(big.Int).Mul(val, exp)
	}

	decimalDiff := EtherDecimals - decimals

	exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimalDiff)), nil)
	return new(big.Int).Div(val, exp)
}
