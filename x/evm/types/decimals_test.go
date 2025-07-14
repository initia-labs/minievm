package types

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToEthersUnit(t *testing.T) {
	val := big.NewInt(1000)
	// decimals == EtherDecimals
	res := ToEthersUnit(EtherDecimals, val)
	require.Equal(t, val, res)

	// decimals > EtherDecimals
	val2 := big.NewInt(1000000000)
	res = ToEthersUnit(EtherDecimals+2, val2)
	exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(2), nil)
	require.Equal(t, new(big.Int).Div(val2, exp), res)

	// decimals < EtherDecimals
	val3 := big.NewInt(1000)
	res = ToEthersUnit(EtherDecimals-2, val3)
	exp = new(big.Int).Exp(big.NewInt(10), big.NewInt(2), nil)
	require.Equal(t, new(big.Int).Mul(val3, exp), res)
}

func TestToGweiUint(t *testing.T) {
	val := big.NewInt(1000)
	// decimals == GweiDecimals
	res := ToGweiUint(GweiDecimals, val)
	require.Equal(t, val, res)

	// decimals > GweiDecimals
	val2 := big.NewInt(1000000000)
	res = ToGweiUint(GweiDecimals+2, val2)
	exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(2), nil)
	require.Equal(t, new(big.Int).Div(val2, exp), res)

	// decimals < GweiDecimals
	val3 := big.NewInt(1000)
	res = ToGweiUint(GweiDecimals-2, val3)
	exp = new(big.Int).Exp(big.NewInt(10), big.NewInt(2), nil)
	require.Equal(t, new(big.Int).Mul(val3, exp), res)
}

func TestFromGweiUnit(t *testing.T) {
	val := big.NewInt(1000)
	// decimals == GweiDecimals
	res := FromGweiUnit(GweiDecimals, val)
	require.Equal(t, val, res)

	// decimals > GweiDecimals
	val2 := big.NewInt(1000)
	res = FromGweiUnit(GweiDecimals+2, val2)
	exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(2), nil)
	require.Equal(t, new(big.Int).Mul(val2, exp), res)

	// decimals < GweiDecimals
	val3 := big.NewInt(1000)
	res = FromGweiUnit(GweiDecimals-2, val3)
	exp = new(big.Int).Exp(big.NewInt(10), big.NewInt(2), nil)
	require.Equal(t, new(big.Int).Div(val3, exp), res)
}

func TestFromEthersUnit(t *testing.T) {
	val := big.NewInt(1000)
	// decimals == EtherDecimals
	res := FromEthersUnit(EtherDecimals, val)
	require.Equal(t, val, res)

	// decimals > EtherDecimals
	val2 := big.NewInt(1000)
	res = FromEthersUnit(EtherDecimals+2, val2)
	exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(2), nil)
	require.Equal(t, new(big.Int).Mul(val2, exp), res)

	// decimals < EtherDecimals
	val3 := big.NewInt(1000)
	res = FromEthersUnit(EtherDecimals-2, val3)
	exp = new(big.Int).Exp(big.NewInt(10), big.NewInt(2), nil)
	require.Equal(t, new(big.Int).Div(val3, exp), res)
}
