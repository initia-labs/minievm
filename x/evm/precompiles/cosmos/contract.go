package cosmosprecompile

import (
	"bytes"
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/pkg/errors"

	abci "github.com/cometbft/cometbft/abci/types"

	"cosmossdk.io/core/address"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/initia-labs/minievm/x/evm/contracts/i_cosmos"
	"github.com/initia-labs/minievm/x/evm/types"
)

var _ vm.ExtendedPrecompiledContract = &CosmosPrecompile{}
var _ vm.PrecompiledContract = &CosmosPrecompile{}

var erc20CosmosABI *abi.ABI

func init() {
	var err error
	erc20CosmosABI, err = i_cosmos.ICosmosMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
}

type CosmosPrecompile struct {
	*abi.ABI

	stateDB types.StateDB
	cdc     codec.Codec
	ac      address.Codec

	ak         types.AccountKeeper
	bk         types.BankKeeper
	edk        types.ERC20DenomKeeper
	grpcRouter types.GRPCRouter

	queryWhitelist types.QueryCosmosWhitelist

	authorityAddr sdk.AccAddress
}

func NewCosmosPrecompile(
	stateDB types.StateDB,
	cdc codec.Codec,
	ac address.Codec,
	ak types.AccountKeeper,
	bk types.BankKeeper,
	edk types.ERC20DenomKeeper,
	grpcRouter types.GRPCRouter,
	queryWhitelist types.QueryCosmosWhitelist,
	authority string,
) (*CosmosPrecompile, error) {
	authorityAddr, err := sdk.AccAddressFromBech32(authority)
	if err != nil {
		return nil, err
	}

	return &CosmosPrecompile{
		ABI:            erc20CosmosABI,
		cdc:            cdc,
		ac:             ac,
		ak:             ak,
		bk:             bk,
		edk:            edk,
		stateDB:        stateDB,
		grpcRouter:     grpcRouter,
		queryWhitelist: queryWhitelist,
		authorityAddr:  authorityAddr,
	}, nil
}

func (e *CosmosPrecompile) originAddress(ctx context.Context, addrBz []byte) (sdk.AccAddress, error) {
	account := e.ak.GetAccount(ctx, addrBz)
	if shorthandCallerAccount, ok := account.(types.ShorthandAccountI); ok {
		addr, err := shorthandCallerAccount.GetOriginalAddress(e.ac)
		if err != nil {
			return nil, err
		}

		addrBz = addr.Bytes()
	}

	return addrBz, nil
}

// ExtendedRun implements vm.ExtendedPrecompiledContract.
func (e *CosmosPrecompile) ExtendedRun(caller vm.ContractRef, input []byte, suppliedGas uint64, readOnly bool) (resBz []byte, usedGas uint64, err error) {
	snapshot := e.stateDB.Snapshot()
	ctx := e.stateDB.Context().WithGasMeter(storetypes.NewGasMeter(suppliedGas))

	defer func() {
		if r := recover(); r != nil {
			switch r.(type) {
			case storetypes.ErrorOutOfGas:
				// set the used gas to the supplied gas
				usedGas = suppliedGas

				// convert cosmos out of gas error to normal error
				err = errors.New("out of gas in precompile")
			default:
				panic(r)
			}
		}

		if err != nil {
			// convert cosmos error to EVM error
			resBz = types.NewRevertReason(err)
			err = vm.ErrExecutionReverted

			// revert the stateDB to the snapshot
			e.stateDB.RevertToSnapshot(snapshot)
		}
	}()

	method, err := e.ABI.MethodById(input)
	if err != nil {
		return nil, 0, types.ErrPrecompileFailed.Wrap(err.Error())
	}

	args, err := method.Inputs.Unpack(input[4:])
	if err != nil {
		return nil, 0, types.ErrPrecompileFailed.Wrap(err.Error())
	}

	// charge input gas
	ctx.GasMeter().ConsumeGas(storetypes.Gas(len(input))*GAS_PER_BYTE, "input bytes")

	switch method.Name {
	case METHOD_IS_BLOCKED_ADDRESS:
		ctx.GasMeter().ConsumeGas(IS_BLOCKED_ADDRESS_GAS, "is_blocked_address")

		var isBlockedAddressArguments IsBlockedAddressArguments
		if err := method.Inputs.Copy(&isBlockedAddressArguments, args); err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		// convert shorthand account to original address
		addr, err := e.originAddress(ctx, isBlockedAddressArguments.Account.Bytes())
		if err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		isBlocked := e.bk.BlockedAddr(addr)

		// abi encode the response
		resBz, err = method.Outputs.Pack(isBlocked)
		if err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}
	case METHOD_IS_MODULE_ADDRESS:
		ctx.GasMeter().ConsumeGas(IS_MODULE_ADDRESS_GAS, "is_module_address")

		var isModuleAddressArguments IsModuleAddressArguments
		if err := method.Inputs.Copy(&isModuleAddressArguments, args); err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		// convert shorthand account to original address
		addr, err := e.originAddress(ctx, isModuleAddressArguments.Account.Bytes())
		if err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		// check if the address is a module account
		account := e.ak.GetAccount(ctx, addr)
		_, isModuleAccount := account.(sdk.ModuleAccountI)

		// abi encode the response
		resBz, err = method.Outputs.Pack(isModuleAccount)
		if err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}
	case METHOD_IS_AUTHORITY_ADDRESS:
		ctx.GasMeter().ConsumeGas(IS_AUTHORITY_ADDRESS_GAS, "is_authority_address")

		var isAuthorityAddressArguments IsAuthorityAddressArguments
		if err := method.Inputs.Copy(&isAuthorityAddressArguments, args); err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		// convert shorthand account to original address
		addr, err := e.originAddress(ctx, isAuthorityAddressArguments.Account.Bytes())
		if err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		// check if the address is the authority address
		isAuthorityAddr := addr.Equals(e.authorityAddr)

		// abi encode the response
		resBz, err = method.Outputs.Pack(isAuthorityAddr)
		if err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}
	case METHOD_TO_COSMOS_ADDRESS:
		ctx.GasMeter().ConsumeGas(TO_COSMOS_ADDRESS_GAS, "to_cosmos_address")

		var toCosmosAddressArguments ToCosmosAddressArguments
		if err := method.Inputs.Copy(&toCosmosAddressArguments, args); err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		addr, err := e.ac.BytesToString(toCosmosAddressArguments.EVMAddress.Bytes())
		if err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		// abi encode the response
		resBz, err = method.Outputs.Pack(addr)
		if err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}
	case METHOD_TO_EVM_ADDRESS:
		ctx.GasMeter().ConsumeGas(TO_EVM_ADDRESS_GAS, "to_evm_address")

		var toEVMAddressArguments ToEVMAddressArguments
		if err := method.Inputs.Copy(&toEVMAddressArguments, args); err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		addr, err := e.ac.StringToBytes(toEVMAddressArguments.CosmosAddress)
		if err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		// check address length
		if len(addr) != common.AddressLength {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrInvalidAddressLength.Wrap(hexutil.Encode(addr))
		}

		// abi encode the response
		resBz, err = method.Outputs.Pack(common.BytesToAddress(addr))
		if err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}
	case METHOD_EXECUTE_COSMOS, METHOD_EXECUTE_COSMOS_WITH_OPTIONS:
		ctx.GasMeter().ConsumeGas(EXECUTE_COSMOS_GAS, "execute_cosmos")

		if readOnly {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrNonReadOnlyMethod.Wrap(method.Name)
		}

		// check if execute cosmos is disabled
		disabled, ok := ctx.Value(types.CONTEXT_KEY_DISABLE_EXECUTE_COSMOS).(bool)
		if ok && disabled {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrExecuteCosmosDisabled.Wrap(method.Name)
		}

		var executeCosmosArguments ExecuteCosmos
		if method.Name == METHOD_EXECUTE_COSMOS {
			if err := method.Inputs.Copy(&executeCosmosArguments, args); err != nil {
				return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
			}
		} else {
			var executeCosmosArgumentsWithOptions ExecuteCosmosWithOptions
			if err := method.Inputs.Copy(&executeCosmosArgumentsWithOptions, args); err != nil {
				return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
			}

			executeCosmosArguments = executeCosmosArgumentsWithOptions.ToExecuteCosmos()
		}

		var sdkMsg sdk.Msg
		if err := e.cdc.UnmarshalInterfaceJSON([]byte(executeCosmosArguments.Msg), &sdkMsg); err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		// check required signers are the same with the caller
		signers, _, err := e.cdc.GetMsgV1Signers(sdkMsg)
		if err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		// convert shorthand account to original address
		callerAddr, err := e.originAddress(ctx, caller.Address().Bytes())
		if err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		for _, signer := range signers {
			if !bytes.Equal(callerAddr, signer) {
				return nil, ctx.GasMeter().GasConsumedToLimit(), sdkerrors.ErrUnauthorized.Wrapf(
					"required signer: `%s`, given signer: `%s`",
					hexutil.Encode(signer), caller.Address(),
				)
			}
		}

		// pre-charge the gas for the execute cosmos
		ctx.GasMeter().ConsumeGas(executeCosmosArguments.GasLimit, "pre-charge execute cosmos gas")

		messages := ctx.Value(types.CONTEXT_KEY_EXECUTE_REQUESTS).(*[]types.ExecuteRequest)
		*messages = append(*messages, types.ExecuteRequest{
			Caller: caller,
			Msg:    sdkMsg,

			AllowFailure: executeCosmosArguments.Options.AllowFailure,
			CallbackId:   executeCosmosArguments.Options.CallbackId,

			GasLimit: executeCosmosArguments.GasLimit,
		})

		// abi encode the response
		resBz, err = method.Outputs.Pack(true)
		if err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}
	case METHOD_QUERY_COSMOS:
		ctx.GasMeter().ConsumeGas(QUERY_COSMOS_GAS, "query_cosmos")

		var queryCosmosArguments QueryCosmosArguments
		if err := method.Inputs.Copy(&queryCosmosArguments, args); err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		route := e.grpcRouter.Route(queryCosmosArguments.Path)
		if route == nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrNotSupportedCosmosQuery.Wrap(queryCosmosArguments.Path)
		}

		protoSet, found := e.queryWhitelist[queryCosmosArguments.Path]
		if !found {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrNotSupportedCosmosQuery.Wrap(queryCosmosArguments.Path)
		}

		reqData, err := types.ConvertJSONToProto(e.cdc, protoSet.Request, []byte(queryCosmosArguments.Req))
		if err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		res, err := route(ctx, &abci.RequestQuery{
			Data: reqData,
			Path: queryCosmosArguments.Path,
		})
		if err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		resBz, err = types.ConvertProtoToJSON(e.cdc, protoSet.Response, res.Value)
		if err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		// abi encode the response
		resBz, err = method.Outputs.Pack(string(resBz))
		if err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}
	case METHOD_TO_DENOM:
		ctx.GasMeter().ConsumeGas(TO_DENOM_GAS, "to_denom")

		var toDenomArguments ToDenomArguments
		if err := method.Inputs.Copy(&toDenomArguments, args); err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		denom, err := types.ContractAddrToDenom(ctx, e.edk, toDenomArguments.ERC20Address)
		if err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		// abi encode the response
		resBz, err = method.Outputs.Pack(denom)
		if err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}
	case METHOD_TO_ERC20:
		ctx.GasMeter().ConsumeGas(TO_ERC20_GAS, "to_erc20")

		var toERC20Arguments ToERC20Arguments
		if err := method.Inputs.Copy(&toERC20Arguments, args); err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		contractAddr, err := types.DenomToContractAddr(ctx, e.edk, toERC20Arguments.Denom)
		if err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		// abi encode the response
		resBz, err = method.Outputs.Pack(contractAddr)
		if err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}
	case METHOD_DISABLE_EXECUTE_COSMOS:
		ctx.GasMeter().ConsumeGas(DISABLE_EXECUTE_COSMOS_GAS, "disable_execute_cosmos")
		e.stateDB.SetContextValue(types.CONTEXT_KEY_DISABLE_EXECUTE_COSMOS, true)

		resBz, err = method.Outputs.Pack(true)
		if err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}
	default:
		return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrUnknownPrecompileMethod.Wrap(method.Name)
	}

	usedGas = ctx.GasMeter().GasConsumedToLimit()
	return resBz, usedGas, nil
}

// RequiredGas implements vm.PrecompiledContract.
func (e *CosmosPrecompile) RequiredGas(input []byte) uint64 {
	return 0
}

// Run implements vm.PrecompiledContract.
func (e *CosmosPrecompile) Run(input []byte) ([]byte, error) {
	return nil, errors.New("the CosmosPrecompile works exclusively with ExtendedRun")
}
