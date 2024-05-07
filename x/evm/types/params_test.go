package types_test

import (
	"encoding/binary"
	"math/rand"
	"testing"

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
}
