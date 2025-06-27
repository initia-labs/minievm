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
