package cosmosprecompile_test

import (
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
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/gogoproto/proto"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"

	contracts "github.com/initia-labs/minievm/x/evm/contracts/i_cosmos"
	precompiles "github.com/initia-labs/minievm/x/evm/precompiles/cosmos"
	"github.com/initia-labs/minievm/x/evm/types"

	oracletypes "github.com/skip-mev/connect/v2/x/oracle/types"
)

func setup() (sdk.Context, codec.Codec, address.Codec, types.AccountKeeper, types.BankKeeper) {
	kv := db.NewMemDB()
	cms := store.NewCommitMultiStore(kv, log.NewNopLogger(), storemetrics.NewNoOpMetrics())

	ctx := sdk.NewContext(cms, cmtproto.Header{}, false, log.NewNopLogger())

	interfaceRegistry, _ := codectypes.NewInterfaceRegistryWithOptions(codectypes.InterfaceRegistryOptions{
		ProtoFiles: proto.HybridResolver,
		SigningOptions: signing.Options{
			AddressCodec:          codecaddress.NewBech32Codec("init"),
			ValidatorAddressCodec: codecaddress.NewBech32Codec("initvaloper"),
		},
	})
	banktypes.RegisterInterfaces(interfaceRegistry)
	types.RegisterInterfaces(interfaceRegistry)

	cdc := codec.NewProtoCodec(interfaceRegistry)
	ac := codecaddress.NewBech32Codec("init")

	return ctx, cdc, ac,
		&MockAccountKeeper{ac: ac, accounts: make(map[string]sdk.AccountI)},
		&MockBankKeeper{ac: ac, blockedAddresses: make(map[string]bool)}
}

func Test_CosmosPrecompile_IsBlockedAddress(t *testing.T) {
	ctx, cdc, ac, ak, bk := setup()
	authorityAddr := authtypes.NewModuleAddress(govtypes.ModuleName).String()

	stateDB := NewMockStateDB(ctx)
	cosmosPrecompile, err := precompiles.NewCosmosPrecompile(stateDB, cdc, ac, ak, bk, nil, nil, nil, nil, authorityAddr)
	require.NoError(t, err)

	evmAddr := common.HexToAddress("0x1")
	cosmosAddr, err := ac.BytesToString(evmAddr.Bytes())
	require.NoError(t, err)

	abi, err := contracts.ICosmosMetaData.GetAbi()
	require.NoError(t, err)

	// is blocked address
	inputBz, err := abi.Pack(precompiles.METHOD_IS_BLOCKED_ADDRESS, evmAddr)
	require.NoError(t, err)

	// out of gas error
	_, _, err = cosmosPrecompile.ExtendedRun(vm.AccountRef(evmAddr), inputBz, precompiles.IS_BLOCKED_ADDRESS_GAS-1, true)
	require.ErrorIs(t, err, vm.ErrOutOfGas)

	retBz, _, err := cosmosPrecompile.ExtendedRun(vm.AccountRef(evmAddr), inputBz, precompiles.IS_BLOCKED_ADDRESS_GAS+uint64(len(inputBz)), true)
	require.NoError(t, err)

	ret, err := abi.Unpack(precompiles.METHOD_IS_BLOCKED_ADDRESS, retBz)
	require.NoError(t, err)
	require.False(t, ret[0].(bool))

	// block address
	bk.(*MockBankKeeper).blockedAddresses[cosmosAddr] = true

	// is blocked address
	inputBz, err = abi.Pack(precompiles.METHOD_IS_BLOCKED_ADDRESS, evmAddr)
	require.NoError(t, err)

	retBz, _, err = cosmosPrecompile.ExtendedRun(vm.AccountRef(evmAddr), inputBz, precompiles.IS_BLOCKED_ADDRESS_GAS+uint64(len(inputBz)), true)
	require.NoError(t, err)

	ret, err = abi.Unpack(precompiles.METHOD_IS_BLOCKED_ADDRESS, retBz)
	require.NoError(t, err)
	require.True(t, ret[0].(bool))
}

func Test_CosmosPrecompile_IsModuleAddress(t *testing.T) {
	ctx, cdc, ac, ak, bk := setup()
	authorityAddr := authtypes.NewModuleAddress(govtypes.ModuleName).String()

	stateDB := NewMockStateDB(ctx)
	cosmosPrecompile, err := precompiles.NewCosmosPrecompile(stateDB, cdc, ac, ak, bk, nil, nil, nil, nil, authorityAddr)
	require.NoError(t, err)

	evmAddr := common.HexToAddress("0x1")
	cosmosAddr, err := ac.BytesToString(evmAddr.Bytes())
	require.NoError(t, err)

	abi, err := contracts.ICosmosMetaData.GetAbi()
	require.NoError(t, err)

	// is module address
	inputBz, err := abi.Pack(precompiles.METHOD_IS_MODULE_ADDRESS, evmAddr)
	require.NoError(t, err)

	// out of gas error
	_, _, err = cosmosPrecompile.ExtendedRun(vm.AccountRef(evmAddr), inputBz, precompiles.IS_MODULE_ADDRESS_GAS-1, true)
	require.ErrorIs(t, err, vm.ErrOutOfGas)

	retBz, _, err := cosmosPrecompile.ExtendedRun(vm.AccountRef(evmAddr), inputBz, precompiles.IS_MODULE_ADDRESS_GAS+uint64(len(inputBz)), true)
	require.NoError(t, err)

	ret, err := abi.Unpack(precompiles.METHOD_IS_MODULE_ADDRESS, retBz)
	require.NoError(t, err)
	require.False(t, ret[0].(bool))

	// module address
	ak.(*MockAccountKeeper).accounts[cosmosAddr] = authtypes.NewEmptyModuleAccount("test")

	// is module address
	inputBz, err = abi.Pack(precompiles.METHOD_IS_MODULE_ADDRESS, evmAddr)
	require.NoError(t, err)

	retBz, _, err = cosmosPrecompile.ExtendedRun(vm.AccountRef(evmAddr), inputBz, precompiles.IS_MODULE_ADDRESS_GAS+uint64(len(inputBz)), true)
	require.NoError(t, err)

	ret, err = abi.Unpack(precompiles.METHOD_IS_MODULE_ADDRESS, retBz)
	require.NoError(t, err)
	require.True(t, ret[0].(bool))
}

func Test_CosmosPrecompile_IsAuthorityAddress(t *testing.T) {
	ctx, cdc, ac, ak, bk := setup()
	authorityAddr := authtypes.NewModuleAddress(govtypes.ModuleName).String()

	stateDB := NewMockStateDB(ctx)
	cosmosPrecompile, err := precompiles.NewCosmosPrecompile(stateDB, cdc, ac, ak, bk, nil, nil, nil, nil, authorityAddr)
	require.NoError(t, err)

	evmAddr := common.HexToAddress("0x1")
	require.NoError(t, err)

	abi, err := contracts.ICosmosMetaData.GetAbi()
	require.NoError(t, err)

	// is authority address
	inputBz, err := abi.Pack(precompiles.METHOD_IS_AUTHORITY_ADDRESS, evmAddr)
	require.NoError(t, err)

	// out of gas error
	_, _, err = cosmosPrecompile.ExtendedRun(vm.AccountRef(evmAddr), inputBz, precompiles.IS_AUTHORITY_ADDRESS_GAS-1, true)
	require.ErrorIs(t, err, vm.ErrOutOfGas)

	retBz, _, err := cosmosPrecompile.ExtendedRun(vm.AccountRef(evmAddr), inputBz, precompiles.IS_AUTHORITY_ADDRESS_GAS+uint64(len(inputBz)), true)
	require.NoError(t, err)

	ret, err := abi.Unpack(precompiles.METHOD_IS_AUTHORITY_ADDRESS, retBz)
	require.NoError(t, err)
	require.False(t, ret[0].(bool))

	// update input to authority address
	authorityEVMAddr := common.BytesToAddress(authtypes.NewModuleAddress(govtypes.ModuleName).Bytes())
	inputBz, err = abi.Pack(precompiles.METHOD_IS_AUTHORITY_ADDRESS, authorityEVMAddr)
	require.NoError(t, err)

	retBz, _, err = cosmosPrecompile.ExtendedRun(vm.AccountRef(evmAddr), inputBz, precompiles.IS_AUTHORITY_ADDRESS_GAS+uint64(len(inputBz)), true)
	require.NoError(t, err)

	ret, err = abi.Unpack(precompiles.METHOD_IS_AUTHORITY_ADDRESS, retBz)
	require.NoError(t, err)
	require.True(t, ret[0].(bool))
}

func Test_CosmosPrecompile_ToCosmosAddress(t *testing.T) {
	ctx, cdc, ac, ak, bk := setup()
	authorityAddr := authtypes.NewModuleAddress(govtypes.ModuleName).String()

	stateDB := NewMockStateDB(ctx)
	cosmosPrecompile, err := precompiles.NewCosmosPrecompile(stateDB, cdc, ac, ak, bk, nil, nil, nil, nil, authorityAddr)
	require.NoError(t, err)

	evmAddr := common.HexToAddress("0x1")
	cosmosAddr, err := ac.BytesToString(evmAddr.Bytes())
	require.NoError(t, err)

	abi, err := contracts.ICosmosMetaData.GetAbi()
	require.NoError(t, err)

	// to cosmos address
	inputBz, err := abi.Pack(precompiles.METHOD_TO_COSMOS_ADDRESS, evmAddr)
	require.NoError(t, err)

	// out of gas error
	_, _, err = cosmosPrecompile.ExtendedRun(vm.AccountRef(evmAddr), inputBz, precompiles.TO_COSMOS_ADDRESS_GAS-1, true)
	require.ErrorIs(t, err, vm.ErrOutOfGas)

	retBz, _, err := cosmosPrecompile.ExtendedRun(vm.AccountRef(evmAddr), inputBz, precompiles.TO_COSMOS_ADDRESS_GAS+uint64(len(inputBz)), true)
	require.NoError(t, err)

	ret, err := abi.Unpack(precompiles.METHOD_TO_COSMOS_ADDRESS, retBz)
	require.NoError(t, err)
	require.Equal(t, cosmosAddr, ret[0].(string))
}

func Test_CosmosPrecompile_ToEVMAddress(t *testing.T) {
	ctx, cdc, ac, ak, bk := setup()
	authorityAddr := authtypes.NewModuleAddress(govtypes.ModuleName).String()

	stateDB := NewMockStateDB(ctx)
	cosmosPrecompile, err := precompiles.NewCosmosPrecompile(stateDB, cdc, ac, ak, bk, nil, nil, nil, nil, authorityAddr)
	require.NoError(t, err)

	evmAddr := common.HexToAddress("0x1")
	cosmosAddr, err := ac.BytesToString(evmAddr.Bytes())
	require.NoError(t, err)

	abi, err := contracts.ICosmosMetaData.GetAbi()
	require.NoError(t, err)

	// to cosmos address
	inputBz, err := abi.Pack(precompiles.METHOD_TO_EVM_ADDRESS, cosmosAddr)
	require.NoError(t, err)

	// out of gas error
	_, _, err = cosmosPrecompile.ExtendedRun(vm.AccountRef(evmAddr), inputBz, precompiles.TO_EVM_ADDRESS_GAS-1, true)
	require.ErrorIs(t, err, vm.ErrOutOfGas)

	retBz, _, err := cosmosPrecompile.ExtendedRun(vm.AccountRef(evmAddr), inputBz, precompiles.TO_EVM_ADDRESS_GAS+uint64(len(inputBz)), true)
	require.NoError(t, err)

	ret, err := abi.Unpack(precompiles.METHOD_TO_EVM_ADDRESS, retBz)
	require.NoError(t, err)
	require.Equal(t, evmAddr, ret[0].(common.Address))
}

func Test_ExecuteCosmos(t *testing.T) {
	ctx, cdc, ac, ak, bk := setup()
	authorityAddr := authtypes.NewModuleAddress(govtypes.ModuleName).String()

	stateDB := NewMockStateDB(ctx)
	cosmosPrecompile, err := precompiles.NewCosmosPrecompile(stateDB, cdc, ac, ak, bk, nil, MockMessageRouter{
		routes: map[string]baseapp.MsgServiceHandler{
			"/minievm.evm.v1.MsgCall": func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
				ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeEVM, sdk.NewAttribute(types.AttributeKeyLog, "dummy")))
				ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeEVM, sdk.NewAttribute(types.AttributeKeyLog, "dummy2")))
				resp, err := proto.Marshal(&types.MsgCallResponse{
					Logs: []types.Log{
						{
							Address: "0x1",
							Topics:  []string{"hello", "world"},
							Data:    "0x1234",
						},
						{
							Address: "0x2",
							Topics:  []string{"hello", "world"},
							Data:    "0x4567",
						},
					},
				})
				require.NoError(t, err)
				return &sdk.Result{
					Data: resp,
				}, nil
			},
			"/minievm.evm.v1.MsgCreate": func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
				ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeEVM, sdk.NewAttribute(types.AttributeKeyLog, "dummy")))
				ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeEVM, sdk.NewAttribute(types.AttributeKeyLog, "dummy2")))
				return &sdk.Result{}, fmt.Errorf("FORCE_REVERT")
			},
		},
	}, nil, nil, authorityAddr)
	require.NoError(t, err)

	evmAddr := common.HexToAddress("0x1")
	cosmosAddr, err := ac.BytesToString(evmAddr.Bytes())
	require.NoError(t, err)

	abi, err := contracts.ICosmosMetaData.GetAbi()
	require.NoError(t, err)

	// 1. execute cosmos message
	inputBz, err := abi.Pack(precompiles.METHOD_EXECUTE_COSMOS, fmt.Sprintf(`{
		"@type": "/minievm.evm.v1.MsgCall",
		"sender": "%s",
		"contract_addr": "0x1",
		"input": "0x"
	}`, cosmosAddr))
	require.NoError(t, err)

	// out of gas error
	_, _, err = cosmosPrecompile.ExtendedRun(vm.AccountRef(evmAddr), inputBz, precompiles.EXECUTE_COSMOS_GAS-1, false)
	require.ErrorIs(t, err, vm.ErrOutOfGas)

	// cannot call execute in readonly mode
	_, _, err = cosmosPrecompile.ExtendedRun(vm.AccountRef(evmAddr), inputBz, precompiles.EXECUTE_COSMOS_GAS+uint64(len(inputBz)), true)
	require.ErrorIs(t, err, vm.ErrExecutionReverted)

	// succeed
	_, _, err = cosmosPrecompile.ExtendedRun(vm.AccountRef(evmAddr), inputBz, precompiles.EXECUTE_COSMOS_GAS+uint64(len(inputBz)), false)
	require.NoError(t, err)

	// should have empty events
	require.Empty(t, stateDB.ctx.EventManager().Events())
	// should have 2 logs
	require.Len(t, stateDB.Logs(), 2)

	// 2. execute create to revert the call
	inputBz, err = abi.Pack(precompiles.METHOD_EXECUTE_COSMOS, fmt.Sprintf(`{
		"@type": "/minievm.evm.v1.MsgCreate",
		"sender": "%s",
		"code": "0x"
	}`, cosmosAddr))
	require.NoError(t, err)

	// failed with unauthorized error
	ret, _, err := cosmosPrecompile.ExtendedRun(vm.AccountRef(evmAddr), inputBz, precompiles.EXECUTE_COSMOS_GAS+uint64(len(inputBz)), false)
	require.ErrorIs(t, err, vm.ErrExecutionReverted)
	require.Contains(t, types.NewRevertError(ret).Error(), "FORCE_REVERT")

	// 3. wrong signer message
	inputBz, err = abi.Pack(precompiles.METHOD_EXECUTE_COSMOS, `{
		"@type": "/minievm.evm.v1.MsgCall",
		"sender": "init1enjh88u7c9s08fgdu28wj6umz94cetjy0hpcxf",
		"contract_addr": "0x1",
		"input": "0x"
	}`)
	require.NoError(t, err)

	// failed with unauthorized error
	ret, _, err = cosmosPrecompile.ExtendedRun(vm.AccountRef(evmAddr), inputBz, precompiles.EXECUTE_COSMOS_GAS+uint64(len(inputBz)), false)
	require.ErrorIs(t, err, vm.ErrExecutionReverted)
	require.Contains(t, types.NewRevertError(ret).Error(), sdkerrors.ErrUnauthorized.Error())

	// 4. invalid message
	inputBz, err = abi.Pack(precompiles.METHOD_EXECUTE_COSMOS, fmt.Sprintf(`{
		"@type": "/minievm.evm.v2.MsgCall",
		"sender": "%s",
		"contract_addr": "0x1",
		"input": "0x"
	}`, cosmosAddr))
	require.NoError(t, err)

	// failed with unauthorized error
	ret, _, err = cosmosPrecompile.ExtendedRun(vm.AccountRef(evmAddr), inputBz, precompiles.EXECUTE_COSMOS_GAS+uint64(len(inputBz)), false)
	require.ErrorIs(t, err, vm.ErrExecutionReverted)
	require.Contains(t, types.NewRevertError(ret).Error(), "unable to resolve type")

	// 5. invalid message 2
	inputBz, err = abi.Pack(precompiles.METHOD_EXECUTE_COSMOS, fmt.Sprintf(`{
		"@type": "/cosmos.bank.v1beta1.MsgSend",
		"from_address": "%s",
		"to_address": "init1enjh88u7c9s08fgdu28wj6umz94cetjy0hpcxf",
		"amount": []
	}`, cosmosAddr))
	require.NoError(t, err)

	// failed with unauthorized error
	ret, _, err = cosmosPrecompile.ExtendedRun(vm.AccountRef(evmAddr), inputBz, precompiles.EXECUTE_COSMOS_GAS+uint64(len(inputBz)), false)
	require.ErrorIs(t, err, vm.ErrExecutionReverted)
	require.Contains(t, types.NewRevertError(ret).Error(), types.ErrNotSupportedCosmosMessage.Error())
}

func Test_QueryCosmos(t *testing.T) {
	ctx, cdc, ac, ak, bk := setup()
	authorityAddr := authtypes.NewModuleAddress(govtypes.ModuleName).String()

	queryPath := "/connect.oracle.v2.Query/Prices"
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

	stateDB := NewMockStateDB(ctx)
	cosmosPrecompile, err := precompiles.NewCosmosPrecompile(stateDB, cdc, ac, ak, bk, nil, nil, MockGRPCRouter{
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
	}, authorityAddr)
	require.NoError(t, err)

	evmAddr := common.HexToAddress("0x1")

	abi, err := contracts.ICosmosMetaData.GetAbi()
	require.NoError(t, err)

	// pack query_cosmos
	inputBz, err := abi.Pack(precompiles.METHOD_QUERY_COSMOS, queryPath, `{"currency_pair_ids": ["BITCOIN/USD"]}`)
	require.NoError(t, err)

	// out of gas error
	_, _, err = cosmosPrecompile.ExtendedRun(vm.AccountRef(evmAddr), inputBz, precompiles.QUERY_COSMOS_GAS-1, false)
	require.ErrorIs(t, err, vm.ErrOutOfGas)

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
	ctx, cdc, ac, ak, bk := setup()
	authorityAddr := authtypes.NewModuleAddress(govtypes.ModuleName).String()

	erc20Addr := common.HexToAddress("0x123")
	denom := "evm/0000000000000000000000000000000000000123"

	stateDB := NewMockStateDB(ctx)
	cosmosPrecompile, err := precompiles.NewCosmosPrecompile(stateDB, cdc, ac, ak, bk, &MockERC20DenomKeeper{
		denomMap: map[string]common.Address{
			denom: erc20Addr,
		},
		addrMap: map[common.Address]string{
			erc20Addr: denom,
		},
	}, nil, nil, nil, authorityAddr)
	require.NoError(t, err)

	evmAddr := common.HexToAddress("0x1")

	abi, err := contracts.ICosmosMetaData.GetAbi()
	require.NoError(t, err)

	// pack to_denom
	inputBz, err := abi.Pack(precompiles.METHOD_TO_DENOM, erc20Addr)
	require.NoError(t, err)

	// out of gas error
	_, _, err = cosmosPrecompile.ExtendedRun(vm.AccountRef(evmAddr), inputBz, precompiles.TO_DENOM_GAS-1, false)
	require.ErrorIs(t, err, vm.ErrOutOfGas)

	// succeed
	retBz, _, err := cosmosPrecompile.ExtendedRun(vm.AccountRef(evmAddr), inputBz, precompiles.TO_DENOM_GAS+uint64(len(inputBz)), true)
	require.NoError(t, err)

	// unpack response
	unpackedRet, err := abi.Methods["to_denom"].Outputs.Unpack(retBz)
	require.NoError(t, err)
	require.Equal(t, denom, unpackedRet[0].(string))
}

func Test_ToErc20(t *testing.T) {
	ctx, cdc, ac, ak, bk := setup()
	authorityAddr := authtypes.NewModuleAddress(govtypes.ModuleName).String()

	erc20Addr := common.HexToAddress("0x123")
	denom := "evm/0000000000000000000000000000000000000123"

	stateDB := NewMockStateDB(ctx)
	cosmosPrecompile, err := precompiles.NewCosmosPrecompile(stateDB, cdc, ac, ak, bk, &MockERC20DenomKeeper{
		denomMap: map[string]common.Address{
			denom: erc20Addr,
		},
		addrMap: map[common.Address]string{
			erc20Addr: denom,
		},
	}, nil, nil, nil, authorityAddr)
	require.NoError(t, err)

	evmAddr := common.HexToAddress("0x1")

	abi, err := contracts.ICosmosMetaData.GetAbi()
	require.NoError(t, err)

	// pack to_erc20
	inputBz, err := abi.Pack(precompiles.METHOD_TO_ERC20, denom)
	require.NoError(t, err)

	// out of gas panic
	_, _, err = cosmosPrecompile.ExtendedRun(vm.AccountRef(evmAddr), inputBz, precompiles.TO_ERC20_GAS-1, false)
	require.ErrorIs(t, err, vm.ErrOutOfGas)

	// succeed
	retBz, _, err := cosmosPrecompile.ExtendedRun(vm.AccountRef(evmAddr), inputBz, precompiles.TO_ERC20_GAS+uint64(len(inputBz)), true)
	require.NoError(t, err)

	// unpack response
	unpackedRet, err := abi.Methods["to_erc20"].Outputs.Unpack(retBz)
	require.NoError(t, err)
	require.Equal(t, erc20Addr, unpackedRet[0].(common.Address))
}
