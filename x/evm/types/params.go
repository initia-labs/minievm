package types

import (
	"cosmossdk.io/core/address"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"gopkg.in/yaml.v3"
)

func DefaultParams() Params {
	return Params{
		AllowCustomERC20: true,
		FeeDenom:         sdk.DefaultBondDenom,
		GasRefundRatio:   math.LegacyNewDecWithPrec(5, 1),
	}
}

func (p Params) String() string {
	out, err := yaml.Marshal(p)
	if err != nil {
		panic(err)
	}
	return string(out)
}

func (p Params) Validate(ac address.Codec) error {
	for _, addr := range p.AllowedPublishers {
		_, err := ac.StringToBytes(addr)
		if err != nil {
			return err
		}
	}

	for _, addr := range p.AllowedCustomERC20s {
		_, err := ContractAddressFromString(ac, addr)
		if err != nil {
			return err
		}
	}

	if p.GasRefundRatio.IsNegative() || p.GasRefundRatio.GT(math.LegacyOneDec()) {
		return ErrInvalidGasRefundRatio
	}

	return nil
}
