package types

import (
	"fmt"

	"cosmossdk.io/core/address"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"

	"gopkg.in/yaml.v3"
)

// MAX_RECURSIVE_DEPTH is the maximum depth of the x/evm call stack.
const MAX_RECURSIVE_DEPTH = 8

func DefaultParams() Params {
	return Params{
		AllowCustomERC20:     true,
		FeeDenom:             sdk.DefaultBondDenom,
		GasRefundRatio:       math.LegacyNewDecWithPrec(5, 1),
		NumRetainBlockHashes: 256,
		// no limit on gas price or gas limit per evm transaction
	}
}

func (p Params) String() string {
	out, err := yaml.Marshal(p)
	if err != nil {
		panic(err)
	}
	return string(out)
}

// normalize to checksum hex addresses
func (p *Params) NormalizeAddresses(ac address.Codec) error {
	if normalizedPublishers, err := normalizeAddrs(ac, p.AllowedPublishers); err != nil {
		return err
	} else {
		p.AllowedPublishers = normalizedPublishers
	}

	if normalizedCustomERC20s, err := normalizeAddrs(ac, p.AllowedCustomERC20s); err != nil {
		return err
	} else {
		p.AllowedCustomERC20s = normalizedCustomERC20s
	}

	if p.GasEnforcement != nil {
		if normalizedUnlimitedGasSenders, err := normalizeAddrs(ac, p.GasEnforcement.UnlimitedGasSenders); err != nil {
			return err
		} else {
			p.GasEnforcement.UnlimitedGasSenders = normalizedUnlimitedGasSenders
		}
	}
	return nil
}

func (p Params) Validate(ac address.Codec) error {
	if err := validateCheckSumHexAddrs(ac, p.AllowedPublishers); err != nil {
		return err
	}
	if err := validateCheckSumHexAddrs(ac, p.AllowedCustomERC20s); err != nil {
		return err
	}
	if p.GasRefundRatio.IsNegative() || p.GasRefundRatio.GT(math.LegacyOneDec()) {
		return ErrInvalidGasRefundRatio
	}

	if p.NumRetainBlockHashes != 0 && p.NumRetainBlockHashes < 256 {
		return ErrInvalidNumRetainBlockHashes
	}

	if p.GasEnforcement != nil {
		if p.GasEnforcement.MaxGasFeeCap != nil && p.GasEnforcement.MaxGasFeeCap.IsNegative() {
			return ErrInvalidGasEnforcement
		}
		if err := validateCheckSumHexAddrs(ac, p.GasEnforcement.UnlimitedGasSenders); err != nil {
			return err
		}
	}

	return nil
}

func (p Params) ToExtraEIPs() []int {
	extraEIPs := make([]int, len(p.ExtraEIPs))
	for i, eip := range p.ExtraEIPs {
		extraEIPs[i] = int(eip)
	}

	return extraEIPs
}

func validateCheckSumHexAddrs(ac address.Codec, addrs []string) error {
	for _, addr := range addrs {
		ethAddr, err := ContractAddressFromString(ac, addr)
		if err != nil || ethAddr == (common.Address{}) {
			return fmt.Errorf("invalid address: %s: %w", addr, err)
		}
		if addr != ethAddr.Hex() {
			return fmt.Errorf("address must be in ethereum checksum hex format: %s", addr)
		}
	}
	return nil
}

func normalizeAddrs(ac address.Codec, addrs []string) ([]string, error) {
	if len(addrs) == 0 {
		return nil, nil
	}
	normalized := make([]string, len(addrs))
	for i, addr := range addrs {
		ethAddr, err := ContractAddressFromString(ac, addr)
		if err != nil {
			return nil, fmt.Errorf("invalid address: %s: %w", addr, err)
		}
		if ethAddr == (common.Address{}) {
			return nil, fmt.Errorf("address cannot be empty: %s", addr)
		}
		normalized[i] = ethAddr.Hex()
	}
	return normalized, nil
}
