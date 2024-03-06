package types

import (
	"cosmossdk.io/core/address"
	"gopkg.in/yaml.v3"
)

func DefaultParams() Params {
	return Params{}
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

	return nil
}
