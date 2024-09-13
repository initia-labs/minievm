package keeper_test

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common"

	evmante "github.com/initia-labs/minievm/x/evm/ante"
	"github.com/initia-labs/minievm/x/evm/types"
)

func Test_GasPrice(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()

	fee, err := input.EVMKeeper.LoadFee(ctx)
	require.NoError(t, err)

	gasPriceInEthersUnit := types.ToEthersUint(0, big.NewInt(123))
	gasPrice := types.FromEthersUnit(fee.Decimals(), gasPriceInEthersUnit)
	ctx = ctx.WithValue(evmante.ContextKeyGasPrices, sdk.DecCoins{sdk.NewDecCoinFromDec(fee.Denom(), math.LegacyNewDecFromBigInt(gasPrice))})

	caller := common.BytesToAddress(addr.Bytes())
	_, evm, err := input.EVMKeeper.CreateEVM(ctx, caller, nil)
	require.NoError(t, err)

	require.Equal(t, gasPriceInEthersUnit, evm.GasPrice)
}

func Test_BaseFee(t *testing.T) {
	ctx, input := createDefaultTestInput(t)
	_, _, addr := keyPubAddr()

	fee, err := input.EVMKeeper.LoadFee(ctx)
	require.NoError(t, err)

	gasPriceInEthersUnit := types.ToEthersUint(0, big.NewInt(123))
	gasPrice := types.FromEthersUnit(fee.Decimals(), gasPriceInEthersUnit)
	input.GasPriceKeeper.GasPrices[fee.Denom()] = math.LegacyNewDecFromBigInt(gasPrice)

	caller := common.BytesToAddress(addr.Bytes())
	_, evm, err := input.EVMKeeper.CreateEVM(ctx, caller, nil)
	require.NoError(t, err)

	require.Equal(t, gasPriceInEthersUnit, evm.Context.BaseFee)
}
