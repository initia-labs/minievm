package cosmosprecompile_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	abci "github.com/cometbft/cometbft/abci/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"

	"cosmossdk.io/core/address"
	"cosmossdk.io/log"
	"cosmossdk.io/math"
	"cosmossdk.io/store"
	storemetrics "cosmossdk.io/store/metrics"
	"cosmossdk.io/x/tx/signing"
	db "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/baseapp"
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

	oracletypes "github.com/skip-mev/slinky/x/oracle/types"
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

	cosmosPrecompile, err := precompiles.NewCosmosPrecompile(cdc, ac, ak, nil, nil, nil)
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
	cosmosPrecompile, err := precompiles.NewCosmosPrecompile(cdc, ac, ak, nil, nil, nil)
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

func Test_ExecuteCosmos(t *testing.T) {
	ctx, cdc, ac, ak := setup()
	cosmosPrecompile, err := precompiles.NewCosmosPrecompile(cdc, ac, ak, nil, nil, nil)
	require.NoError(t, err)

	cosmosPrecompile = cosmosPrecompile.WithContext(ctx).(precompiles.CosmosPrecompile)

	evmAddr := common.HexToAddress("0x1")
	cosmosAddr, err := ac.BytesToString(evmAddr.Bytes())
	require.NoError(t, err)

	abi, err := contracts.ICosmosMetaData.GetAbi()
	require.NoError(t, err)

	// execute cosmos message
	inputBz, err := abi.Pack(precompiles.METHOD_EXECUTE_COSMOS, fmt.Sprintf(`{
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
		_, _, _ = cosmosPrecompile.ExtendedRun(vm.AccountRef(evmAddr), inputBz, precompiles.EXECUTE_COSMOS_GAS-1, false)
	})

	// cannot call execute in readonly mode
	_, _, err = cosmosPrecompile.ExtendedRun(vm.AccountRef(evmAddr), inputBz, precompiles.EXECUTE_COSMOS_GAS+uint64(len(inputBz)), true)
	require.Error(t, err)

	// succeed
	_, _, err = cosmosPrecompile.ExtendedRun(vm.AccountRef(evmAddr), inputBz, precompiles.EXECUTE_COSMOS_GAS+uint64(len(inputBz)), false)
	require.NoError(t, err)

	messages := ctx.Value(types.CONTEXT_KEY_COSMOS_MESSAGES).(*[]sdk.Msg)
	require.Len(t, *messages, 1)
	require.Equal(t, (*messages)[0], &banktypes.MsgSend{
		FromAddress: cosmosAddr,
		ToAddress:   "init1enjh88u7c9s08fgdu28wj6umz94cetjy0hpcxf",
		Amount:      sdk.NewCoins(sdk.NewCoin("stake", math.NewInt(100))),
	})

	// wrong signer message
	inputBz, err = abi.Pack(precompiles.METHOD_EXECUTE_COSMOS, fmt.Sprintf(`{
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
	_, _, err = cosmosPrecompile.ExtendedRun(vm.AccountRef(evmAddr), inputBz, precompiles.EXECUTE_COSMOS_GAS+uint64(len(inputBz)), false)
	require.ErrorContains(t, err, sdkerrors.ErrUnauthorized.Error())
}

func Test_QueryCosmos(t *testing.T) {
	ctx, cdc, ac, ak := setup()

	queryPath := "/slinky.oracle.v1.Query/Prices"
	expectedRet := oracletypes.GetPricesResponse{
		Prices: []oracletypes.GetPriceResponse{
			{
				Price: &oracletypes.QuotePrice{
					Price:          math.NewInt(100),
					BlockTimestamp: time.Time{},
					BlockHeight:    100,
				},
			},
		},
	}
	cosmosPrecompile, err := precompiles.NewCosmosPrecompile(cdc, ac, ak, nil, MockGRPCRouter{
		routes: map[string]baseapp.GRPCQueryHandler{
			queryPath: func(ctx sdk.Context, req *abci.RequestQuery) (*abci.ResponseQuery, error) {
				resBz, err := cdc.Marshal(&expectedRet)
				if err != nil {
					return nil, err
				}

				return &abci.ResponseQuery{
					Code:  0,
					Value: resBz,
				}, nil
			},
		},
	}, types.QueryCosmosWhitelist{
		queryPath: {
			Request:  &oracletypes.GetPricesRequest{},
			Response: &oracletypes.GetPricesResponse{},
		},
	})
	require.NoError(t, err)

	cosmosPrecompile = cosmosPrecompile.WithContext(ctx).(precompiles.CosmosPrecompile)

	evmAddr := common.HexToAddress("0x1")

	abi, err := contracts.ICosmosMetaData.GetAbi()
	require.NoError(t, err)

	// pack query_cosmos
	inputBz, err := abi.Pack(precompiles.METHOD_QUERY_COSMOS, queryPath, `{"currency_pair_ids": ["BITCOIN/USD"]}`)
	require.NoError(t, err)

	// out of gas panic
	require.Panics(t, func() {
		_, _, _ = cosmosPrecompile.ExtendedRun(vm.AccountRef(evmAddr), inputBz, precompiles.QUERY_COSMOS_GAS-1, false)
	})

	// succeed
	retBz, _, err := cosmosPrecompile.ExtendedRun(vm.AccountRef(evmAddr), inputBz, precompiles.QUERY_COSMOS_GAS+uint64(len(inputBz)), true)
	require.NoError(t, err)

	// unpack response
	unpackedRet, err := abi.Methods["query_cosmos"].Outputs.Unpack(retBz)
	require.NoError(t, err)

	var ret oracletypes.GetPricesResponse
	err = cdc.UnmarshalJSON([]byte(unpackedRet[0].(string)), &ret)
	require.NoError(t, err)

	require.Equal(t, expectedRet, ret)
}

func Test_ToDenom(t *testing.T) {
	ctx, cdc, ac, ak := setup()

	erc20Addr := common.HexToAddress("0x123")
	denom := "evm/0000000000000000000000000000000000000123"

	cosmosPrecompile, err := precompiles.NewCosmosPrecompile(cdc, ac, ak, &MockERC20DenomKeeper{
		denomMap: map[string]common.Address{
			denom: erc20Addr,
		},
		addrMap: map[common.Address]string{
			erc20Addr: denom,
		},
	}, nil, nil)
	require.NoError(t, err)

	cosmosPrecompile = cosmosPrecompile.WithContext(ctx).(precompiles.CosmosPrecompile)

	evmAddr := common.HexToAddress("0x1")

	abi, err := contracts.ICosmosMetaData.GetAbi()
	require.NoError(t, err)

	// pack to_denom
	inputBz, err := abi.Pack(precompiles.METHOD_TO_DENOM, erc20Addr)
	require.NoError(t, err)

	// out of gas panic
	require.Panics(t, func() {
		_, _, _ = cosmosPrecompile.ExtendedRun(vm.AccountRef(evmAddr), inputBz, precompiles.TO_DENOM_GAS-1, false)
	})

	// succeed
	retBz, _, err := cosmosPrecompile.ExtendedRun(vm.AccountRef(evmAddr), inputBz, precompiles.TO_DENOM_GAS+uint64(len(inputBz)), true)
	require.NoError(t, err)

	// unpack response
	unpackedRet, err := abi.Methods["to_denom"].Outputs.Unpack(retBz)
	require.NoError(t, err)
	require.Equal(t, denom, unpackedRet[0].(string))
}

func Test_ToErc20(t *testing.T) {
	ctx, cdc, ac, ak := setup()

	erc20Addr := common.HexToAddress("0x123")
	denom := "evm/0000000000000000000000000000000000000123"

	cosmosPrecompile, err := precompiles.NewCosmosPrecompile(cdc, ac, ak, &MockERC20DenomKeeper{
		denomMap: map[string]common.Address{
			denom: erc20Addr,
		},
		addrMap: map[common.Address]string{
			erc20Addr: denom,
		},
	}, nil, nil)
	require.NoError(t, err)

	cosmosPrecompile = cosmosPrecompile.WithContext(ctx).(precompiles.CosmosPrecompile)

	evmAddr := common.HexToAddress("0x1")

	abi, err := contracts.ICosmosMetaData.GetAbi()
	require.NoError(t, err)

	// pack to_erc20
	inputBz, err := abi.Pack(precompiles.METHOD_TO_ERC20, denom)
	require.NoError(t, err)

	// out of gas panic
	require.Panics(t, func() {
		_, _, _ = cosmosPrecompile.ExtendedRun(vm.AccountRef(evmAddr), inputBz, precompiles.TO_ERC20_GAS-1, false)
	})

	// succeed
	retBz, _, err := cosmosPrecompile.ExtendedRun(vm.AccountRef(evmAddr), inputBz, precompiles.TO_ERC20_GAS+uint64(len(inputBz)), true)
	require.NoError(t, err)

	// unpack response
	unpackedRet, err := abi.Methods["to_erc20"].Outputs.Unpack(retBz)
	require.NoError(t, err)
	require.Equal(t, erc20Addr, unpackedRet[0].(common.Address))
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

var _ types.GRPCRouter = MockGRPCRouter{}

type MockGRPCRouter struct {
	routes map[string]baseapp.GRPCQueryHandler
}

func (router MockGRPCRouter) Route(path string) baseapp.GRPCQueryHandler {
	return router.routes[path]
}

var _ types.ERC20DenomKeeper = &MockERC20DenomKeeper{}

type MockERC20DenomKeeper struct {
	denomMap map[string]common.Address
	addrMap  map[common.Address]string
}

// GetContractAddrByDenom implements types.ERC20DenomKeeper.
func (e *MockERC20DenomKeeper) GetContractAddrByDenom(_ context.Context, denom string) (common.Address, error) {
	addr, found := e.denomMap[denom]
	if !found {
		return common.Address{}, sdkerrors.ErrNotFound
	}

	return addr, nil
}

// GetDenomByContractAddr implements types.ERC20DenomKeeper.
func (e *MockERC20DenomKeeper) GetDenomByContractAddr(_ context.Context, addr common.Address) (string, error) {
	denom, found := e.addrMap[addr]
	if !found {
		return "", sdkerrors.ErrNotFound
	}

	return denom, nil
}
