package erc20registry_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"cosmossdk.io/log"
	"cosmossdk.io/store"
	storemetrics "cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	db "github.com/cosmos/cosmos-db"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"

	contracts "github.com/initia-labs/minievm/x/evm/contracts/i_erc20_registry"
	precompiles "github.com/initia-labs/minievm/x/evm/precompiles/erc20_registry"
	"github.com/initia-labs/minievm/x/evm/types"
)

func setup() (sdk.Context, types.IERC20StoresKeeper) {
	kv := db.NewMemDB()
	cms := store.NewCommitMultiStore(kv, log.NewNopLogger(), storemetrics.NewNoOpMetrics())
	return sdk.NewContext(cms, cmtproto.Header{}, false, log.NewNopLogger()), ERC20StoresKeeper{
		erc20s: make(map[string]bool),
		stores: make(map[string]map[string]bool),
	}
}

var _ types.IERC20StoresKeeper = ERC20StoresKeeper{}

type ERC20StoresKeeper struct {
	erc20s map[string]bool
	stores map[string]map[string]bool
}

const REGISTER_GAS storetypes.Gas = 300
const REGISTER_STORE_GAS storetypes.Gas = 200
const IS_STORE_REGISTERED_GAS storetypes.Gas = 100

func (e ERC20StoresKeeper) Register(ctx context.Context, contractAddr common.Address) error {
	sdk.UnwrapSDKContext(ctx).GasMeter().ConsumeGas(REGISTER_GAS, "register gas")

	e.erc20s[contractAddr.Hex()] = true
	return nil
}

// IsRegistered implements types.IERC20StoresKeeper.
func (e ERC20StoresKeeper) IsStoreRegistered(ctx context.Context, addr sdk.AccAddress, contractAddr common.Address) (bool, error) {
	sdk.UnwrapSDKContext(ctx).GasMeter().ConsumeGas(IS_STORE_REGISTERED_GAS, "is_register gas")

	store, ok := e.stores[addr.String()]
	if !ok {
		return false, nil
	}

	_, ok = store[contractAddr.Hex()]
	return ok, nil
}

// Register implements types.IERC20StoresKeeper.
func (e ERC20StoresKeeper) RegisterStore(ctx context.Context, addr sdk.AccAddress, contractAddr common.Address) error {
	sdk.UnwrapSDKContext(ctx).GasMeter().ConsumeGas(REGISTER_STORE_GAS, "register gas")

	_, ok := e.stores[addr.String()]
	if !ok {
		e.stores[addr.String()] = make(map[string]bool)
	}

	e.stores[addr.String()][contractAddr.Hex()] = true
	return nil
}

func Test_ERC20Registry(t *testing.T) {
	ctx, k := setup()

	registry, err := precompiles.NewERC20Registry(k)
	require.NoError(t, err)

	// set context
	registry = registry.WithContext(ctx).(precompiles.ERC20Registry)

	erc20Addr := common.HexToAddress("0x1")
	accountAddr := common.HexToAddress("0x2")
	abi, err := contracts.IErc20RegistryMetaData.GetAbi()
	require.NoError(t, err)

	// register erc20
	bz, err := abi.Pack(precompiles.METHOD_REGISTER)
	require.NoError(t, err)

	// out of gas panic
	require.Panics(t, func() {
		_, _, _ = registry.ExtendedRun(vm.AccountRef(erc20Addr), bz, REGISTER_GAS-1, false)
	})

	// non read only method fail
	_, _, err = registry.ExtendedRun(vm.AccountRef(erc20Addr), bz, REGISTER_GAS, true)
	require.Error(t, err)

	// success
	_, usedGas, err := registry.ExtendedRun(vm.AccountRef(erc20Addr), bz, REGISTER_GAS, false)
	require.NoError(t, err)
	require.Equal(t, usedGas, uint64(REGISTER_GAS))

	// check erc20 registered
	require.True(t, k.(ERC20StoresKeeper).erc20s[erc20Addr.Hex()])

	// check unregistered
	bz, err = abi.Pack(precompiles.METHOD_IS_STORE_REGISTERED, accountAddr)
	require.NoError(t, err)

	resBz, usedGas, err := registry.ExtendedRun(vm.AccountRef(erc20Addr), bz, IS_STORE_REGISTERED_GAS, true)
	require.NoError(t, err)
	require.Equal(t, usedGas, uint64(IS_STORE_REGISTERED_GAS))

	res, err := abi.Methods[precompiles.METHOD_IS_STORE_REGISTERED].Outputs.Unpack(resBz)
	require.NoError(t, err)
	require.False(t, res[0].(bool))

	// register store
	bz, err = abi.Pack(precompiles.METHOD_REGISTER_STORE, accountAddr)
	require.NoError(t, err)

	// out of gas panic
	require.Panics(t, func() {
		_, _, _ = registry.ExtendedRun(vm.AccountRef(erc20Addr), bz, REGISTER_STORE_GAS-1, false)
	})

	// non read only method fail
	_, _, err = registry.ExtendedRun(vm.AccountRef(erc20Addr), bz, REGISTER_STORE_GAS, true)
	require.Error(t, err)

	// success
	_, usedGas, err = registry.ExtendedRun(vm.AccountRef(erc20Addr), bz, REGISTER_STORE_GAS, false)
	require.NoError(t, err)
	require.Equal(t, usedGas, uint64(REGISTER_STORE_GAS))

	// check registered
	bz, err = abi.Pack(precompiles.METHOD_IS_STORE_REGISTERED, accountAddr)
	require.NoError(t, err)

	// out of gas panic
	require.Panics(t, func() {
		_, _, _ = registry.ExtendedRun(vm.AccountRef(erc20Addr), bz, IS_STORE_REGISTERED_GAS-1, true)
	})

	resBz, usedGas, err = registry.ExtendedRun(vm.AccountRef(erc20Addr), bz, IS_STORE_REGISTERED_GAS, true)
	require.NoError(t, err)
	require.Equal(t, usedGas, uint64(IS_STORE_REGISTERED_GAS))

	res, err = abi.Methods[precompiles.METHOD_IS_STORE_REGISTERED].Outputs.Unpack(resBz)
	require.NoError(t, err)
	require.True(t, res[0].(bool))
}
