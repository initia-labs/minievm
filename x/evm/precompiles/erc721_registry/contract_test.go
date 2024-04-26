package erc721registryprecompile_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"cosmossdk.io/log"
	"cosmossdk.io/store"
	storemetrics "cosmossdk.io/store/metrics"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	db "github.com/cosmos/cosmos-db"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"

	contracts "github.com/initia-labs/minievm/x/evm/contracts/i_erc721_registry"
	precompiles "github.com/initia-labs/minievm/x/evm/precompiles/erc721_registry"
	"github.com/initia-labs/minievm/x/evm/types"
)

func setup() (sdk.Context, types.IERC721StoresKeeper) {
	kv := db.NewMemDB()
	cms := store.NewCommitMultiStore(kv, log.NewNopLogger(), storemetrics.NewNoOpMetrics())
	return sdk.NewContext(cms, cmtproto.Header{}, false, log.NewNopLogger()), ERC721StoresKeeper{
		erc721s: make(map[string]bool),
		stores:  make(map[string]map[string]bool),
	}
}

var _ types.IERC721StoresKeeper = ERC721StoresKeeper{}

type ERC721StoresKeeper struct {
	erc721s map[string]bool
	stores  map[string]map[string]bool
}

func (e ERC721StoresKeeper) Register(ctx context.Context, contractAddr common.Address) error {
	e.erc721s[contractAddr.Hex()] = true
	return nil
}

// IsRegistered implements types.Ierc721StoresKeeper.
func (e ERC721StoresKeeper) IsStoreRegistered(ctx context.Context, addr sdk.AccAddress, contractAddr common.Address) (bool, error) {
	store, ok := e.stores[addr.String()]
	if !ok {
		return false, nil
	}

	_, ok = store[contractAddr.Hex()]
	return ok, nil
}

// Register implements types.Ierc721StoresKeeper.
func (e ERC721StoresKeeper) RegisterStore(ctx context.Context, addr sdk.AccAddress, contractAddr common.Address) error {
	_, ok := e.stores[addr.String()]
	if !ok {
		e.stores[addr.String()] = make(map[string]bool)
	}

	e.stores[addr.String()][contractAddr.Hex()] = true
	return nil
}

func Test_erc721RegistryPrecompile(t *testing.T) {
	ctx, k := setup()

	registry, err := precompiles.NewERC721RegistryPrecompile(k)
	require.NoError(t, err)

	// set context
	registry = registry.WithContext(ctx).(precompiles.ERC721RegistryPrecompile)

	erc721Addr := common.HexToAddress("0x1")
	accountAddr := common.HexToAddress("0x2")
	abi, err := contracts.IErc721RegistryMetaData.GetAbi()
	require.NoError(t, err)

	// register erc721
	bz, err := abi.Pack(precompiles.METHOD_REGISTER)
	require.NoError(t, err)

	// out of gas panic
	require.Panics(t, func() {
		_, _, _ = registry.ExtendedRun(vm.AccountRef(erc721Addr), bz, precompiles.REGISTER_GAS-1, false)
	})

	// non read only method fail
	_, _, err = registry.ExtendedRun(vm.AccountRef(erc721Addr), bz, precompiles.REGISTER_GAS+uint64(len(bz)), true)
	require.Error(t, err)

	// success
	_, usedGas, err := registry.ExtendedRun(vm.AccountRef(erc721Addr), bz, precompiles.REGISTER_GAS+uint64(len(bz)), false)
	require.NoError(t, err)
	require.Equal(t, usedGas, uint64(precompiles.REGISTER_GAS)+uint64(len(bz)))

	// check erc721 registered
	require.True(t, k.(ERC721StoresKeeper).erc721s[erc721Addr.Hex()])

	// check unregistered
	bz, err = abi.Pack(precompiles.METHOD_IS_STORE_REGISTERED, accountAddr)
	require.NoError(t, err)

	resBz, usedGas, err := registry.ExtendedRun(vm.AccountRef(erc721Addr), bz, precompiles.IS_STORE_REGISTERED_GAS+uint64(len(bz)), true)
	require.NoError(t, err)
	require.Equal(t, usedGas, uint64(precompiles.IS_STORE_REGISTERED_GAS)+uint64(len(bz)))

	res, err := abi.Methods[precompiles.METHOD_IS_STORE_REGISTERED].Outputs.Unpack(resBz)
	require.NoError(t, err)
	require.False(t, res[0].(bool))

	// register store
	bz, err = abi.Pack(precompiles.METHOD_REGISTER_STORE, accountAddr)
	require.NoError(t, err)

	// out of gas panic
	require.Panics(t, func() {
		_, _, _ = registry.ExtendedRun(vm.AccountRef(erc721Addr), bz, precompiles.REGISTER_STORE_GAS-1, false)
	})

	// non read only method fail
	_, _, err = registry.ExtendedRun(vm.AccountRef(erc721Addr), bz, precompiles.REGISTER_STORE_GAS+uint64(len(bz)), true)
	require.Error(t, err)

	// success
	_, usedGas, err = registry.ExtendedRun(vm.AccountRef(erc721Addr), bz, precompiles.REGISTER_STORE_GAS+uint64(len(bz)), false)
	require.NoError(t, err)
	require.Equal(t, usedGas, uint64(precompiles.REGISTER_STORE_GAS)+uint64(len(bz)))

	// check registered
	bz, err = abi.Pack(precompiles.METHOD_IS_STORE_REGISTERED, accountAddr)
	require.NoError(t, err)

	// out of gas panic
	require.Panics(t, func() {
		_, _, _ = registry.ExtendedRun(vm.AccountRef(erc721Addr), bz, precompiles.IS_STORE_REGISTERED_GAS-1, true)
	})

	resBz, usedGas, err = registry.ExtendedRun(vm.AccountRef(erc721Addr), bz, precompiles.IS_STORE_REGISTERED_GAS+uint64(len(bz)), true)
	require.NoError(t, err)
	require.Equal(t, usedGas, uint64(precompiles.IS_STORE_REGISTERED_GAS)+uint64(len(bz)))

	res, err = abi.Methods[precompiles.METHOD_IS_STORE_REGISTERED].Outputs.Unpack(resBz)
	require.NoError(t, err)
	require.True(t, res[0].(bool))
}
