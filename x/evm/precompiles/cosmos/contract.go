package cosmosprecompile

import (
	"bytes"
	"context"
	"errors"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/vm"

	abci "github.com/cometbft/cometbft/abci/types"

	"cosmossdk.io/core/address"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/initia-labs/minievm/x/evm/contracts/i_cosmos"
	"github.com/initia-labs/minievm/x/evm/types"
)

var _ vm.ExtendedPrecompiledContract = CosmosPrecompile{}
var _ vm.PrecompiledContract = CosmosPrecompile{}
var _ types.WithContext = CosmosPrecompile{}

type CosmosPrecompile struct {
	*abi.ABI

	ctx context.Context
	cdc codec.Codec
	ac  address.Codec

	ak         types.AccountKeeper
	bk         types.BankKeeper
	edk        types.ERC20DenomKeeper
	grpcRouter types.GRPCRouter

	queryWhitelist types.QueryCosmosWhitelist
}

func NewCosmosPrecompile(
	cdc codec.Codec,
	ac address.Codec,
	ak types.AccountKeeper,
	bk types.BankKeeper,
	edk types.ERC20DenomKeeper,
	grpcRouter types.GRPCRouter,
	queryWhitelist types.QueryCosmosWhitelist,
) (CosmosPrecompile, error) {
	abi, err := i_cosmos.ICosmosMetaData.GetAbi()
	if err != nil {
		return CosmosPrecompile{}, err
	}

	return CosmosPrecompile{
		ABI:            abi,
		cdc:            cdc,
		ac:             ac,
		ak:             ak,
		bk:             bk,
		edk:            edk,
		grpcRouter:     grpcRouter,
		queryWhitelist: queryWhitelist,
	}, nil
}

func (e CosmosPrecompile) WithContext(ctx context.Context) vm.PrecompiledContract {
	e.ctx = ctx
	return e
}

func (e CosmosPrecompile) originAddress(ctx context.Context, addrBz []byte) (sdk.AccAddress, error) {
	account := e.ak.GetAccount(ctx, addrBz)
	if shorthandCallerAccount, ok := account.(types.ShorthandAccountI); ok {
		addr, err := shorthandCallerAccount.GetOriginalAddress(e.ac)
		if err != nil {
			return nil, types.ErrPrecompileFailed.Wrap(err.Error())
		}

		addrBz = addr.Bytes()
	}

	return addrBz, nil
}

// ExtendedRun implements vm.ExtendedPrecompiledContract.
func (e CosmosPrecompile) ExtendedRun(caller vm.ContractRef, input []byte, suppliedGas uint64, readOnly bool) (resBz []byte, usedGas uint64, err error) {
	defer func() {
		if r := recover(); r != nil {
			switch r.(type) {
			case storetypes.ErrorOutOfGas:
				// convert cosmos out of gas error to EVM out of gas error
				usedGas = suppliedGas
				err = vm.ErrOutOfGas
			default:
				panic(r)
			}
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

	ctx := sdk.UnwrapSDKContext(e.ctx).WithGasMeter(storetypes.NewGasMeter(suppliedGas))

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
		addr, err := e.originAddress(ctx, isBlockedAddressArguments.Address.Bytes())
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
		addr, err := e.originAddress(ctx, isModuleAddressArguments.Address.Bytes())
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
	case METHOD_EXECUTE_COSMOS:
		ctx.GasMeter().ConsumeGas(EXECUTE_COSMOS_GAS, "execute_cosmos")

		if readOnly {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrNonReadOnlyMethod.Wrap(method.Name)
		}

		var executeCosmosArguments ExecuteCosmosArguments
		if err := method.Inputs.Copy(&executeCosmosArguments, args); err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
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

		messages := ctx.Value(types.CONTEXT_KEY_COSMOS_MESSAGES).(*[]sdk.Msg)
		*messages = append(*messages, sdkMsg)

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
	default:
		return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrUnknownPrecompileMethod.Wrap(method.Name)
	}

	usedGas = ctx.GasMeter().GasConsumedToLimit()
	return resBz, usedGas, nil
}

// RequiredGas implements vm.PrecompiledContract.
func (e CosmosPrecompile) RequiredGas(input []byte) uint64 {
	return 0
}

// Run implements vm.PrecompiledContract.
func (e CosmosPrecompile) Run(input []byte) ([]byte, error) {
	return nil, errors.New("the CosmosPrecompile works exclusively with ExtendedRun")
}
