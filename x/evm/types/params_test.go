package types_test

import (
	"encoding/binary"
	"math/rand"
	"strings"
	"testing"

	"cosmossdk.io/math"
	"github.com/stretchr/testify/require"

	"github.com/cometbft/cometbft/crypto/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authcodec "github.com/cosmos/cosmos-sdk/x/auth/codec"

	"github.com/initia-labs/minievm/x/evm/types"
)

func Test_ParamsValidate(t *testing.T) {
	params := types.DefaultParams()

	ac := authcodec.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix())
	err := params.Validate(ac)
	require.NoError(t, err)

	seed := make([]byte, 8)
	binary.BigEndian.PutUint64(seed, rand.Uint64())

	key := ed25519.GenPrivKeyFromSecret(seed)
	pub := key.PubKey()
	addr := sdk.AccAddress(pub.Address())

	params.AllowedPublishers = append(params.AllowedPublishers, addr.String())
	params.AllowedCustomERC20s = append(params.AllowedCustomERC20s, addr.String())
	maxGasFeeCap := math.NewInt(1000000000)
	params.GasEnforcement = &types.GasEnforcement{
		MaxGasFeeCap: maxGasFeeCap,
		MaxGasLimit:  1000000,
		UnlimitedGasSenders: []string{
			"0x000000000000000000000000000000000000000a",
			"0x000000000000000000000000000000000000000b",
		},
	}
	params.NormalizeAddresses(ac)
	err = params.Validate(ac)
	require.NoError(t, err)

	// invalid address
	params.AllowedPublishers = append(params.AllowedPublishers, addr.String()+"abc")

	err = params.Validate(ac)
	require.Error(t, err)

	// invalid erc20 address
	params.AllowedPublishers = params.AllowedPublishers[:1]
	params.AllowedCustomERC20s = append(params.AllowedCustomERC20s, addr.String()+"abc")
	err = params.Validate(ac)
	require.Error(t, err)

	// invalid gas enforcement
	params.GasEnforcement = &types.GasEnforcement{
		MaxGasFeeCap: maxGasFeeCap,
		MaxGasLimit:  1000000,
		UnlimitedGasSenders: []string{
			"0x000000000000000000000000000000000000000a",
			"0x000000000000000000000000000000000000000B",
		},
	}
	err = params.Validate(ac)
	require.Error(t, err)
}

func Test_ParamsNormalizeAddresses(t *testing.T) {
	params := types.DefaultParams()
	ac := authcodec.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix())

	var testAddrs []string
	var expectedAddrs []string
	for i := range 3 {
		seed := make([]byte, 8)
		binary.BigEndian.PutUint64(seed, rand.Uint64())
		key := ed25519.GenPrivKeyFromSecret(seed)
		pub := key.PubKey()
		addr := sdk.AccAddress(pub.Address())
		bech32Addr := addr.String()
		ethAddr, _ := types.ContractAddressFromString(ac, bech32Addr)
		expectedChecksum := ethAddr.Hex()
		expectedAddrs = append(expectedAddrs, expectedChecksum)
		// testaddrs: checksum, lower, bech32
		switch i {
		case 0:
			testAddrs = append(testAddrs, expectedChecksum)
		case 1:
			testAddrs = append(testAddrs, strings.ToLower(expectedChecksum))
		default:
			testAddrs = append(testAddrs, bech32Addr)
		}
	}

	// set the test addresses in params
	params.AllowedPublishers = testAddrs
	params.AllowedCustomERC20s = testAddrs
	params.GasEnforcement = &types.GasEnforcement{
		UnlimitedGasSenders: testAddrs,
		MaxGasFeeCap:        math.NewInt(0),
		MaxGasLimit:         0,
	}

	err := params.NormalizeAddresses(ac)
	require.NoError(t, err)
	require.Equal(t, expectedAddrs, params.AllowedPublishers)
	require.Equal(t, expectedAddrs, params.AllowedCustomERC20s)
	require.Equal(t, expectedAddrs, params.GasEnforcement.UnlimitedGasSenders)

	err = params.Validate(ac)
	require.NoError(t, err)
}

func Test_ParamsString(t *testing.T) {
	params := types.DefaultParams()

	// Test that String() method works without panicking
	str := params.String()
	require.NotEmpty(t, str)
	require.Contains(t, str, "allowcustomerc20")
	require.Contains(t, str, "fee_denom")
}

func Test_ParamsToExtraEIPs(t *testing.T) {
	params := types.DefaultParams()

	// Test with empty ExtraEIPs
	extraEIPs := params.ToExtraEIPs()
	require.Empty(t, extraEIPs)

	// Test with some ExtraEIPs
	params.ExtraEIPs = []int64{1, 2, 3, 4, 5}
	extraEIPs = params.ToExtraEIPs()
	require.Equal(t, []int{1, 2, 3, 4, 5}, extraEIPs)
}

func Test_ParamsValidate_EdgeCases(t *testing.T) {
	ac := authcodec.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix())

	// Test negative gas refund ratio
	params := types.DefaultParams()
	params.GasRefundRatio = math.LegacyNewDecWithPrec(-1, 1)
	err := params.Validate(ac)
	require.Error(t, err)
	require.Contains(t, err.Error(), "Invalid gas refund ratio")

	// Test gas refund ratio greater than 1
	params.GasRefundRatio = math.LegacyNewDecWithPrec(11, 1)
	err = params.Validate(ac)
	require.Error(t, err)
	require.Contains(t, err.Error(), "Invalid gas refund ratio")

	// Test invalid NumRetainBlockHashes
	params = types.DefaultParams()
	params.NumRetainBlockHashes = 100 // less than 256
	err = params.Validate(ac)
	require.Error(t, err)
	require.Contains(t, err.Error(), "Invalid num retain block hashes")

	// Test nil MaxGasFeeCap
	params = types.DefaultParams()
	params.GasEnforcement = &types.GasEnforcement{
		MaxGasFeeCap: math.Int{},
		MaxGasLimit:  1000000,
	}
	err = params.Validate(ac)
	require.Error(t, err)
	require.Contains(t, err.Error(), "Invalid gas enforcement parameters")

	// Test negative MaxGasFeeCap
	params.GasEnforcement.MaxGasFeeCap = math.NewInt(-1)
	err = params.Validate(ac)
	require.Error(t, err)
	require.Contains(t, err.Error(), "Invalid gas enforcement parameters")
}

func Test_ParamsNormalizeAddresses_EdgeCases(t *testing.T) {
	ac := authcodec.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix())

	// Test with empty addresses
	params := types.DefaultParams()
	params.AllowedPublishers = []string{}
	params.AllowedCustomERC20s = []string{}
	params.GasEnforcement = &types.GasEnforcement{
		UnlimitedGasSenders: []string{},
		MaxGasFeeCap:        math.NewInt(0),
		MaxGasLimit:         0,
	}

	err := params.NormalizeAddresses(ac)
	require.NoError(t, err)
	require.Empty(t, params.AllowedPublishers)
	require.Empty(t, params.AllowedCustomERC20s)
	require.Empty(t, params.GasEnforcement.UnlimitedGasSenders)

	// Test with nil GasEnforcement
	params = types.DefaultParams()
	params.GasEnforcement = nil
	err = params.NormalizeAddresses(ac)
	require.NoError(t, err)

	// Test with invalid address
	params = types.DefaultParams()
	params.AllowedPublishers = []string{"invalid_address"}
	err = params.NormalizeAddresses(ac)
	require.Error(t, err)
	require.Contains(t, err.Error(), "invalid address")
}

func Test_DefaultParams(t *testing.T) {
	params := types.DefaultParams()

	require.True(t, params.AllowCustomERC20)
	require.Equal(t, sdk.DefaultBondDenom, params.FeeDenom)
	require.Equal(t, math.LegacyNewDecWithPrec(5, 1), params.GasRefundRatio)
	require.Equal(t, uint64(256), params.NumRetainBlockHashes)
	require.Nil(t, params.GasEnforcement)
	require.Empty(t, params.ExtraEIPs)
}
