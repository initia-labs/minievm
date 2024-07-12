package app

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"cosmossdk.io/log"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	dbm "github.com/cosmos/cosmos-db"

	"cosmossdk.io/x/upgrade"
	"github.com/cosmos/cosmos-sdk/testutil/mock"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authzmodule "github.com/cosmos/cosmos-sdk/x/authz/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/consensus"
	groupmodule "github.com/cosmos/cosmos-sdk/x/group/module"
	"github.com/cosmos/ibc-go/modules/capability"
	capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"

	ica "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts"
	"github.com/cosmos/ibc-go/v8/modules/apps/transfer"
	ibc "github.com/cosmos/ibc-go/v8/modules/core"

	opchild "github.com/initia-labs/OPinit/x/opchild"

	"github.com/initia-labs/minievm/x/evm"
	evmconfig "github.com/initia-labs/minievm/x/evm/config"
)

func TestSimAppExportAndBlockedAddrs(t *testing.T) {
	app := SetupWithGenesisAccounts(nil, nil)

	// BlockedAddresses returns a map of addresses in app v1 and a map of modules name in app v2.
	for acc := range app.ModuleAccountAddrs() {
		var addr sdk.AccAddress
		if modAddr, err := sdk.AccAddressFromBech32(acc); err == nil {
			addr = modAddr
		} else {
			addr = app.AccountKeeper.GetModuleAddress(acc)
		}

		require.True(
			t,
			app.BankKeeper.BlockedAddr(addr),
			fmt.Sprintf("ensure that blocked addresses are properly set in bank keeper: %s should be blocked", acc),
		)
	}
}

func TestGetMaccPerms(t *testing.T) {
	dup := GetMaccPerms()
	require.Equal(t, maccPerms, dup, "duplicated module account permissions differed from actual module account permissions")
}

func TestInitGenesisOnMigration(t *testing.T) {
	db := dbm.NewMemDB()
	logger := log.NewLogger(os.Stdout)
	app := NewMinitiaApp(
		logger, db, dbm.NewMemDB(),
		nil, true, evmconfig.DefaultEVMConfig(),
		EmptyAppOptions{},
	)
	ctx := app.NewContextLegacy(true, cmtproto.Header{Height: app.LastBlockHeight()})

	// Create a mock module. This module will serve as the new module we're
	// adding during a migration.
	mockCtrl := gomock.NewController(t)
	t.Cleanup(mockCtrl.Finish)
	mockModule := mock.NewMockAppModuleWithAllExtensions(mockCtrl)
	mockDefaultGenesis := json.RawMessage(`{"key": "value"}`)
	mockModule.EXPECT().DefaultGenesis(gomock.Eq(app.appCodec)).Times(1).Return(mockDefaultGenesis)
	mockModule.EXPECT().InitGenesis(gomock.Eq(ctx), gomock.Eq(app.appCodec), gomock.Eq(mockDefaultGenesis)).Times(1)
	mockModule.EXPECT().ConsensusVersion().Times(1).Return(uint64(0))

	app.ModuleManager.Modules["mock"] = mockModule
	app.ModuleManager.OrderMigrations = []string{"mock"}

	// Run migrations only for "mock" module. We exclude it from
	// the VersionMap to simulate upgrading with a new module.
	_, err := app.ModuleManager.RunMigrations(ctx, app.configurator,
		module.VersionMap{
			"bank":               bank.AppModule{}.ConsensusVersion(),
			"auth":               auth.AppModule{}.ConsensusVersion(),
			"authz":              authzmodule.AppModule{}.ConsensusVersion(),
			"upgrade":            upgrade.AppModule{}.ConsensusVersion(),
			"capability":         capability.AppModule{}.ConsensusVersion(),
			"group":              groupmodule.AppModule{}.ConsensusVersion(),
			"consensus":          consensus.AppModule{}.ConsensusVersion(),
			"ibc":                ibc.AppModule{}.ConsensusVersion(),
			"transfer":           transfer.AppModule{}.ConsensusVersion(),
			"interchainaccounts": ica.AppModule{}.ConsensusVersion(),
			"evm":                evm.AppModule{}.ConsensusVersion(),
			"opchild":            opchild.AppModule{}.ConsensusVersion(),
		},
	)
	require.NoError(t, err)
}

func TestUpgradeStateOnGenesis(t *testing.T) {
	app := SetupWithGenesisAccounts(nil, nil)

	// make sure the upgrade keeper has version map in state
	ctx := app.NewContext(true)
	vm, err := app.UpgradeKeeper.GetModuleVersionMap(ctx)
	require.NoError(t, err)

	for v, i := range app.ModuleManager.Modules {
		if i, ok := i.(module.HasConsensusVersion); ok {
			require.Equal(t, vm[v], i.ConsensusVersion())
		}
	}
}

func TestGetKey(t *testing.T) {
	db := dbm.NewMemDB()
	app := NewMinitiaApp(
		log.NewLogger(os.Stdout),
		db, dbm.NewMemDB(), nil, true,
		evmconfig.DefaultEVMConfig(),
		EmptyAppOptions{},
	)

	require.NotEmpty(t, app.GetKey(banktypes.StoreKey))
	require.NotEmpty(t, app.GetMemKey(capabilitytypes.MemStoreKey))
}
