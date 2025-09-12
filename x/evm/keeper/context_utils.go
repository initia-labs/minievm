package keeper

import (
	"context"
	"errors"
	"fmt"
	"math"
	"math/big"

	"cosmossdk.io/collections"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/tracing"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/params"

	"github.com/initia-labs/initia/crypto/ethsecp256k1"

	evmstate "github.com/initia-labs/minievm/x/evm/state"
	"github.com/initia-labs/minievm/x/evm/types"
)

func (k Keeper) LoadFee(ctx context.Context, params types.Params) (types.Fee, error) {
	feeContract, err := types.DenomToContractAddr(ctx, k, params.FeeDenom)
	if err != nil && !errors.Is(err, collections.ErrNotFound) {
		return types.Fee{}, err
	}

	decimals := uint8(0)
	if (feeContract != common.Address{} &&
		// erc20Keeper.Decimals is also calling LoadFee, so we need to check this call is not recursive
		sdk.UnwrapSDKContext(ctx).Value(types.CONTEXT_KEY_LOAD_DECIMALS) == nil) {
		decimals = k.erc20Keeper.Decimals(ctx, feeContract)
	}

	return types.NewFee(params.FeeDenom, feeContract, decimals), nil
}

func (k Keeper) extractGasPriceFromContext(ctx context.Context, fee types.Fee) (*big.Int, error) {
	if (fee.Contract() == common.Address{}) {
		return big.NewInt(0), nil
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	value := sdkCtx.Value(types.CONTEXT_KEY_GAS_PRICES)
	if value == nil {
		return big.NewInt(0), nil
	}

	gasPrices := value.(sdk.DecCoins)
	gasPriceDec := gasPrices.AmountOf(fee.Denom())
	if !gasPriceDec.IsPositive() {
		return big.NewInt(0), nil
	}

	// multiply by 1e9 to prevent decimal drops
	gasPrice := gasPriceDec.
		MulTruncate(sdkmath.LegacyNewDec(1e9)).
		TruncateInt().BigInt()

	return types.ToEthersUnit(fee.Decimals()+9, gasPrice), nil
}

func (k Keeper) baseFee(ctx context.Context, fee types.Fee) (*big.Int, error) {
	gasPriceDec, err := k.gasPriceKeeper.GasPrice(ctx, fee.Denom())
	if err != nil {
		return nil, err
	}

	// multiply by 1e9 to prevent decimal drops
	gasPrice := gasPriceDec.
		MulTruncate(sdkmath.LegacyNewDec(1e9)).
		TruncateInt().BigInt()

	return types.ToEthersUnit(fee.Decimals()+9, gasPrice), nil
}

func (k Keeper) BaseFee(ctx context.Context) (*big.Int, error) {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	fee, err := k.LoadFee(ctx, params)
	if err != nil {
		return nil, err
	}

	return k.baseFee(ctx, fee)
}

// prepareSDKContext prepares the SDK context for EVM execution.
// 1. sets the cosmos messages to context
// 2. checks the recursive depth and increments it (the maximum depth is 8)
func prepareSDKContext(ctx sdk.Context) (sdk.Context, error) {
	// set cosmos messages to context
	ctx = ctx.WithValue(types.CONTEXT_KEY_EXECUTE_REQUESTS, &[]types.ExecuteRequest{})

	depth := 1
	if val := ctx.Value(types.CONTEXT_KEY_RECURSIVE_DEPTH); val != nil {
		depth = val.(int) + 1
		if depth > types.MAX_RECURSIVE_DEPTH {
			return ctx, types.ErrExceedMaxRecursiveDepth
		}
	}

	// set recursive depth to context
	return ctx.WithValue(types.CONTEXT_KEY_RECURSIVE_DEPTH, depth), nil
}

// decorateTracing setup the tracing for the EVM execution if there is tracing configuration or not return evm without tracing.
// 1. sets the tracer to the EVM and the stateDB to the tracing VMContext.
// 2. returns a cleanup function to rollback the tracing.
func decorateTracing(ctx context.Context, evm *vm.EVM, stateDB *evmstate.StateDB) (func(), error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	// if tracing is not enabled, return the evm with pure stateDB
	if sdkCtx.Value(types.CONTEXT_KEY_TRACING) == nil || sdkCtx.Value(types.CONTEXT_KEY_TRACE_EVM) == nil {
		evm.StateDB = stateDB
		return func() {}, nil
	}
	// setup hooked stateDB and cleanup function
	tracing := sdkCtx.Value(types.CONTEXT_KEY_TRACING).(*types.Tracing)
	evmPointer := sdkCtx.Value(types.CONTEXT_KEY_TRACE_EVM).(**vm.EVM)

	originalStateDB := tracing.VMContext().StateDB

	evm.Config.Tracer = tracing.Tracer()
	evm.StateDB = evmstate.NewHookedState(stateDB, tracing.Tracer())
	tracing.VMContext().StateDB = evm.StateDB

	originalEVM := *evmPointer
	*evmPointer = evm

	return func() {
		tracing.VMContext().StateDB = originalStateDB
		*evmPointer = originalEVM
	}, nil
}

// chargeIntrinsicGas charges the intrinsic gas for the given data, list, and authList.
func chargeIntrinsicGas(gasBalance uint64, isContractCreation bool, data []byte, list coretypes.AccessList, authList []coretypes.SetCodeAuthorization, rules params.Rules) (uint64, error) {
	intrinsicGas, err := core.IntrinsicGas(data, list, authList, isContractCreation, rules.IsHomestead, rules.IsIstanbul, rules.IsShanghai)
	if err != nil {
		return 0, err
	}
	if gasBalance < intrinsicGas {
		return 0, fmt.Errorf("%w: have %d, want %d", core.ErrIntrinsicGas, gasBalance, intrinsicGas)
	}
	return gasBalance - intrinsicGas, nil
}

// buildDefaultBlockContext builds the default block context for the given context.
func buildDefaultBlockContext(ctx context.Context) (vm.BlockContext, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	headerHash := sdkCtx.HeaderHash()
	if len(headerHash) == 0 {
		headerHash = make([]byte, 32)
	}

	return vm.BlockContext{
		BlockNumber: big.NewInt(sdkCtx.BlockHeight()),
		Time:        uint64(sdkCtx.BlockTime().Unix()),
		Random:      (*common.Hash)(headerHash),
	}, nil
}

// computeGasLimit computes the gas limit for the given context.
func (k Keeper) computeGasLimit(sdkCtx sdk.Context) uint64 {
	gasLimit := sdkCtx.GasMeter().Limit() - sdkCtx.GasMeter().GasConsumedToLimit()
	if sdkCtx.ExecMode() == sdk.ExecModeSimulate {
		gasLimit = k.config.ContractSimulationGasLimit
	}

	return gasLimit
}

// consumeGas consumes gas
func consumeGas(ctx sdk.Context, gasUsed, gasRemaining uint64, description string) {
	// evm sometimes return 0 gasRemaining, but it's not an out of gas error.
	// cosmos use infinite gas meter at simulation and block operations.
	//
	// to prevent uint64 overflow, we don't consume gas when gas meter is infinite
	// and gasRemaining is 0.
	if ctx.GasMeter().Limit() == math.MaxUint64 && gasRemaining == 0 {
		return
	}

	ctx.GasMeter().ConsumeGas(gasUsed, description)
}

// applyAuthorizations applies the given set code authorizations to the EVM state.
func applyAuthorizations(ctx sdk.Context, evm *vm.EVM, authList []coretypes.SetCodeAuthorization) error {
	for _, auth := range authList {
		authority, pub, err := auth.AuthorityWithPubKey()
		if err != nil {
			return fmt.Errorf("%w: %v", core.ErrAuthorizationInvalidSignature, err)
		}

		evm.StateDB.AddAddressToAccessList(authority)
		code := evm.StateDB.GetCode(authority)
		if _, ok := coretypes.ParseDelegation(code); len(code) != 0 && !ok {
			return core.ErrAuthorizationDestinationHasCode
		}
		if have := evm.StateDB.GetNonce(authority); have != auth.Nonce {
			return core.ErrAuthorizationNonceMismatch
		}
		// If the account already exists in state, refund the new account cost
		// charged in the intrinsic calculation.
		if evm.StateDB.Exist(authority) {
			evm.StateDB.AddRefund(params.CallNewAccountGas - params.TxAuthTupleGas)
		}
		// parse pubkey
		pubKey, err := ethsecp256k1.NewPubKeyFromBytes(pub)
		if err != nil {
			return fmt.Errorf("%w: %v", core.ErrAuthorizationInvalidSignature, err)
		}
		// Update nonce, pubkey and account code.
		evm.StateDB.(*evmstate.StateDB).SetNonceAndPubKey(authority, auth.Nonce+1, pubKey, tracing.NonceChangeAuthorization)
		if auth.Address == (common.Address{}) {
			// Delegation to zero address means clear.
			evm.StateDB.SetCode(authority, nil)
			return nil
		}

		// Otherwise install delegation to auth.Address.
		evm.StateDB.SetCode(authority, coretypes.AddressToDelegation(auth.Address))
	}

	return nil
}
