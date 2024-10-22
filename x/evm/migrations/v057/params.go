package types

import (
	"gopkg.in/yaml.v3"
)

func (p Params) String() string {
	out, err := yaml.Marshal(p)
	if err != nil {
		panic(err)
	}
	return string(out)
}
