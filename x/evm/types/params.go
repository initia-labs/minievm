package types

import (
	"cosmossdk.io/core/address"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"gopkg.in/yaml.v3"
)

func DefaultParams() Params {
	return Params{
		AllowCustomERC20: true,
		FeeDenom:         sdk.DefaultBondDenom,
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

	return nil
}
