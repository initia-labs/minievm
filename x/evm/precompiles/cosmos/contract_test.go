package cosmosprecompile_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"

	"cosmossdk.io/core/address"
	"cosmossdk.io/log"
	"cosmossdk.io/math"
	"cosmossdk.io/store"
	storemetrics "cosmossdk.io/store/metrics"
	"cosmossdk.io/x/tx/signing"
	db "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/codec"
	codecaddress "github.com/cosmos/cosmos-sdk/codec/address"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/gogoproto/proto"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"

	contracts "github.com/initia-labs/minievm/x/evm/contracts/i_cosmos"
	precompiles "github.com/initia-labs/minievm/x/evm/precompiles/cosmos"
	"github.com/initia-labs/minievm/x/evm/types"
)

func setup() (sdk.Context, codec.Codec, address.Codec, types.AccountKeeper) {
	kv := db.NewMemDB()
	cms := store.NewCommitMultiStore(kv, log.NewNopLogger(), storemetrics.NewNoOpMetrics())

	ctx := sdk.NewContext(cms, cmtproto.Header{}, false, log.NewNopLogger()).WithValue(types.CONTEXT_KEY_COSMOS_MESSAGES, &[]sdk.Msg{})

	interfaceRegistry, _ := codectypes.NewInterfaceRegistryWithOptions(codectypes.InterfaceRegistryOptions{
		ProtoFiles: proto.HybridResolver,
		SigningOptions: signing.Options{
			AddressCodec:          codecaddress.NewBech32Codec("init"),
			ValidatorAddressCodec: codecaddress.NewBech32Codec("initvaloper"),
		},
	})
	banktypes.RegisterInterfaces(interfaceRegistry)

	cdc := codec.NewProtoCodec(interfaceRegistry)
	ac := codecaddress.NewBech32Codec("init")

	return ctx, cdc, ac, &MockAccountKeeper{accounts: make(map[string]sdk.AccountI)}
}

func Test_CosmosPrecompile_ToCosmosAddress(t *testing.T) {
	ctx, cdc, ac, ak := setup()

	cosmosPrecompile, err := precompiles.NewCosmosPrecompile(cdc, ac, ak)
	require.NoError(t, err)

	cosmosPrecompile = cosmosPrecompile.WithContext(ctx).(precompiles.CosmosPrecompile)

	evmAddr := common.HexToAddress("0x1")
	cosmosAddr, err := ac.BytesToString(evmAddr.Bytes())
	require.NoError(t, err)

	abi, err := contracts.ICosmosMetaData.GetAbi()
	require.NoError(t, err)

	// to cosmos address
	inputBz, err := abi.Pack(precompiles.METHOD_TO_COSMOS_ADDRESS, evmAddr)
	require.NoError(t, err)

	// out of gas panic
	require.Panics(t, func() {
		_, _, _ = cosmosPrecompile.ExtendedRun(vm.AccountRef(evmAddr), inputBz, precompiles.TO_COSMOS_ADDRESS_GAS-1, true)
	})

	retBz, _, err := cosmosPrecompile.ExtendedRun(vm.AccountRef(evmAddr), inputBz, precompiles.TO_COSMOS_ADDRESS_GAS+uint64(len(inputBz)), true)
	require.NoError(t, err)

	ret, err := abi.Unpack(precompiles.METHOD_TO_COSMOS_ADDRESS, retBz)
	require.NoError(t, err)
	require.Equal(t, cosmosAddr, ret[0].(string))
}

func Test_CosmosPrecompile_ToEVMAddress(t *testing.T) {
	ctx, cdc, ac, ak := setup()
	cosmosPrecompile, err := precompiles.NewCosmosPrecompile(cdc, ac, ak)
	require.NoError(t, err)

	cosmosPrecompile = cosmosPrecompile.WithContext(ctx).(precompiles.CosmosPrecompile)

	evmAddr := common.HexToAddress("0x1")
	cosmosAddr, err := ac.BytesToString(evmAddr.Bytes())
	require.NoError(t, err)

	abi, err := contracts.ICosmosMetaData.GetAbi()
	require.NoError(t, err)

	// to cosmos address
	inputBz, err := abi.Pack(precompiles.METHOD_TO_EVM_ADDRESS, cosmosAddr)
	require.NoError(t, err)

	// out of gas panic
	require.Panics(t, func() {
		_, _, _ = cosmosPrecompile.ExtendedRun(vm.AccountRef(evmAddr), inputBz, precompiles.TO_EVM_ADDRESS_GAS-1, true)
	})

	retBz, _, err := cosmosPrecompile.ExtendedRun(vm.AccountRef(evmAddr), inputBz, precompiles.TO_EVM_ADDRESS_GAS+uint64(len(inputBz)), true)
	require.NoError(t, err)

	ret, err := abi.Unpack(precompiles.METHOD_TO_EVM_ADDRESS, retBz)
	require.NoError(t, err)
	require.Equal(t, evmAddr, ret[0].(common.Address))
}

func Test_ExecuteCosmosMessage(t *testing.T) {
	ctx, cdc, ac, ak := setup()
	cosmosPrecompile, err := precompiles.NewCosmosPrecompile(cdc, ac, ak)
	require.NoError(t, err)

	cosmosPrecompile = cosmosPrecompile.WithContext(ctx).(precompiles.CosmosPrecompile)

	evmAddr := common.HexToAddress("0x1")
	cosmosAddr, err := ac.BytesToString(evmAddr.Bytes())
	require.NoError(t, err)

	abi, err := contracts.ICosmosMetaData.GetAbi()
	require.NoError(t, err)

	// execute cosmos message
	inputBz, err := abi.Pack(precompiles.METHOD_EXECUTE_COSMOS_MESSAGE, fmt.Sprintf(`{
		"@type": "/cosmos.bank.v1beta1.MsgSend",
		"from_address": "%s",
		"to_address": "init1enjh88u7c9s08fgdu28wj6umz94cetjy0hpcxf",
		"amount": [
			{
				"denom": "stake",
				"amount": "100"
			}
		]
	}`, cosmosAddr))
	require.NoError(t, err)

	// out of gas panic
	require.Panics(t, func() {
		_, _, _ = cosmosPrecompile.ExtendedRun(vm.AccountRef(evmAddr), inputBz, precompiles.EXECUTE_COSMOS_MESSAGE_GAS-1, false)
	})

	// cannot call execute_cosmos_message in readonly mode
	_, _, err = cosmosPrecompile.ExtendedRun(vm.AccountRef(evmAddr), inputBz, precompiles.EXECUTE_COSMOS_MESSAGE_GAS+uint64(len(inputBz)), true)
	require.Error(t, err)

	// succeed
	_, _, err = cosmosPrecompile.ExtendedRun(vm.AccountRef(evmAddr), inputBz, precompiles.EXECUTE_COSMOS_MESSAGE_GAS+uint64(len(inputBz)), false)
	require.NoError(t, err)

	messages := ctx.Value(types.CONTEXT_KEY_COSMOS_MESSAGES).(*[]sdk.Msg)
	require.Len(t, *messages, 1)
	require.Equal(t, (*messages)[0], &banktypes.MsgSend{
		FromAddress: cosmosAddr,
		ToAddress:   "init1enjh88u7c9s08fgdu28wj6umz94cetjy0hpcxf",
		Amount:      sdk.NewCoins(sdk.NewCoin("stake", math.NewInt(100))),
	})

	// wrong signer message
	inputBz, err = abi.Pack(precompiles.METHOD_EXECUTE_COSMOS_MESSAGE, fmt.Sprintf(`{
		"@type": "/cosmos.bank.v1beta1.MsgSend",
		"from_address": "init1enjh88u7c9s08fgdu28wj6umz94cetjy0hpcxf",
		"to_address": "%s",
		"amount": [
			{
				"denom": "stake",
				"amount": "100"
			}
		]
	}`, cosmosAddr))
	require.NoError(t, err)

	// failed with unauthorized error
	_, _, err = cosmosPrecompile.ExtendedRun(vm.AccountRef(evmAddr), inputBz, precompiles.EXECUTE_COSMOS_MESSAGE_GAS+uint64(len(inputBz)), false)
	require.ErrorContains(t, err, sdkerrors.ErrUnauthorized.Error())
}

var _ types.AccountKeeper = &MockAccountKeeper{}

// mock account keeper for testing
type MockAccountKeeper struct {
	accounts map[string]sdk.AccountI
}

// GetAccount implements types.AccountKeeper.
func (k MockAccountKeeper) GetAccount(ctx context.Context, addr sdk.AccAddress) sdk.AccountI {
	return k.accounts[addr.String()]
}

// HasAccount implements types.AccountKeeper.
func (k MockAccountKeeper) HasAccount(ctx context.Context, addr sdk.AccAddress) bool {
	_, ok := k.accounts[addr.String()]
	return ok
}

// NewAccount implements types.AccountKeeper.
func (k *MockAccountKeeper) NewAccount(ctx context.Context, acc sdk.AccountI) sdk.AccountI {
	acc.SetAccountNumber(uint64(len(k.accounts)))
	return acc
}

// NewAccountWithAddress implements types.AccountKeeper.
func (k MockAccountKeeper) NewAccountWithAddress(ctx context.Context, addr sdk.AccAddress) sdk.AccountI {
	return authtypes.NewBaseAccount(addr, nil, uint64(len(k.accounts)), 0)
}

// NextAccountNumber implements types.AccountKeeper.
func (k MockAccountKeeper) NextAccountNumber(ctx context.Context) uint64 {
	return uint64(len(k.accounts))
}

// SetAccount implements types.AccountKeeper.
func (k MockAccountKeeper) SetAccount(ctx context.Context, acc sdk.AccountI) {
	k.accounts[acc.GetAddress().String()] = acc
}
