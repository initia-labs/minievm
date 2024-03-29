package cosmosprecompile

import (
	"bytes"
	"context"
	"errors"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/vm"

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

	ak types.AccountKeeper
}

func NewCosmosPrecompile(cdc codec.Codec, ac address.Codec, ak types.AccountKeeper) (CosmosPrecompile, error) {
	abi, err := i_cosmos.ICosmosMetaData.GetAbi()
	if err != nil {
		return CosmosPrecompile{}, err
	}

	return CosmosPrecompile{ABI: abi, cdc: cdc, ac: ac, ak: ak}, nil
}

func (e CosmosPrecompile) WithContext(ctx context.Context) vm.PrecompiledContract {
	e.ctx = ctx
	return e
}

const (
	METHOD_TO_COSMOS_ADDRESS      = "to_cosmos_address"
	METHOD_TO_EVM_ADDRESS         = "to_evm_address"
	METHOD_EXECUTE_COSMOS_MESSAGE = "execute_cosmos_message"
)

// ExtendedRun implements vm.ExtendedPrecompiledContract.
func (e CosmosPrecompile) ExtendedRun(caller vm.ContractRef, input []byte, suppliedGas uint64, readOnly bool) (resBz []byte, usedGas uint64, err error) {
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

		resBz, err = method.Outputs.Pack(common.BytesToAddress(addr))
		if err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}
	case METHOD_EXECUTE_COSMOS_MESSAGE:
		ctx.GasMeter().ConsumeGas(EXECUTE_COSMOS_MESSAGE_GAS, "execute_cosmos_message")

		if readOnly {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrNonReadOnlyMethod.Wrap(method.Name)
		}

		var executeCosmosMessageArguments ExecuteCosmosMessageArguments
		if err := method.Inputs.Copy(&executeCosmosMessageArguments, args); err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		var sdkMsg sdk.Msg
		if err := e.cdc.UnmarshalInterfaceJSON([]byte(executeCosmosMessageArguments.Msg), &sdkMsg); err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		// check required signers are the same with the caller
		signers, _, err := e.cdc.GetMsgV1Signers(sdkMsg)
		if err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		for _, signer := range signers {
			if bytes.Equal(caller.Address().Bytes(), signer) {
				continue
			}

			// if signer is different from the caller, check if the signer is a shorthand account.
			// and then check shorthand account's original address is the same with the caller.
			if len(signer) != common.AddressLength {
				signerAccount := e.ak.GetAccount(ctx, signer)
				if shorthandCallerAccount, ok := signerAccount.(types.ShorthandAccountI); ok {
					addr, err := shorthandCallerAccount.GetOriginalAddress(e.ac)
					if err != nil {
						return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
					}

					if bytes.Equal(addr.Bytes(), signer) {
						continue
					}
				}
			}

			return nil, ctx.GasMeter().GasConsumedToLimit(), sdkerrors.ErrUnauthorized.Wrapf(
				"required signer: `%s`, given signer: `%s`",
				hexutil.Encode(signer), caller.Address(),
			)
		}

		messages := ctx.Value(types.CONTEXT_KEY_COSMOS_MESSAGES).(*[]sdk.Msg)
		*messages = append(*messages, sdkMsg)

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
