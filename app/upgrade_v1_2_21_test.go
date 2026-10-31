package app

import (
	"testing"
	"time"

	upgradetypes "cosmossdk.io/x/upgrade/types"

	abci "github.com/cometbft/cometbft/abci/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"

	"github.com/stretchr/testify/require"

	"github.com/initia-labs/minievm/app/upgrades/v1_2_21"
)

// The restart upgrade is applied by block time on the restart chains, not by a stored plan.
// It must run exactly once, in the first block at or after RestartAfter, clear forwarding
// allowed_denoms, leave no plan behind, and never run on a chain outside the set.
func TestRestartUpgradeAppliesOnceByBlockTime(t *testing.T) {
	app := SetupWithGenesisAccounts(nil, nil) // height 1 committed

	cutoff := time.Date(2026, time.November, 1, 0, 0, 0, 0, time.UTC)
	oldAfter := v1_2_21.RestartAfter
	v1_2_21.RestartAfter = cutoff
	v1_2_21.RestartChains[app.ChainID()] = true
	t.Cleanup(func() { v1_2_21.RestartAfter = oldAfter; delete(v1_2_21.RestartChains, app.ChainID()) })

	// halt-time state
	ctx := app.NewUncachedContext(false, cmtproto.Header{Height: 1, ChainID: app.ChainID()})
	require.NoError(t, app.ForwardingKeeper.AllowedDenoms.Set(ctx, "*"))

	// height 2, before the cutoff nothing happens
	finalizeAt(t, app, 2, cutoff.Add(-time.Second))
	ctx = app.NewUncachedContext(false, cmtproto.Header{Height: 2, ChainID: app.ChainID()})
	done, err := app.UpgradeKeeper.GetDoneHeight(ctx, v1_2_21.UpgradeName)
	require.NoError(t, err)
	require.Zero(t, done)
	require.Equal(t, []string{"*"}, app.ForwardingKeeper.GetAllowedDenoms(ctx))

	// height 3, at the cutoff the handler runs in the pre-block
	finalizeAt(t, app, 3, cutoff)
	ctx = app.NewUncachedContext(false, cmtproto.Header{Height: 3, ChainID: app.ChainID()})
	done, err = app.UpgradeKeeper.GetDoneHeight(ctx, v1_2_21.UpgradeName)
	require.NoError(t, err)
	require.Equal(t, int64(3), done)
	require.Empty(t, app.ForwardingKeeper.GetAllowedDenoms(ctx))
	_, err = app.UpgradeKeeper.GetUpgradePlan(ctx)
	require.ErrorIs(t, err, upgradetypes.ErrNoUpgradePlanFound, "no plan may be left behind")

	// height 4 not applied again, even if someone re-adds a denom
	require.NoError(t, app.ForwardingKeeper.AllowedDenoms.Set(ctx, "umin"))
	finalizeAt(t, app, 4, cutoff.Add(time.Minute))
	ctx = app.NewUncachedContext(false, cmtproto.Header{Height: 4, ChainID: app.ChainID()})
	done, err = app.UpgradeKeeper.GetDoneHeight(ctx, v1_2_21.UpgradeName)
	require.NoError(t, err)
	require.Equal(t, int64(3), done)
}

// A real denom list (no "*") is not touched, the wiring fix makes it work as configured.
func TestRestartUpgradeKeepsRealDenomList(t *testing.T) {
	app := SetupWithGenesisAccounts(nil, nil)
	v1_2_21.RestartChains[app.ChainID()] = true
	t.Cleanup(func() { delete(v1_2_21.RestartChains, app.ChainID()) })

	ctx := app.NewUncachedContext(false, cmtproto.Header{Height: 1, ChainID: app.ChainID()})
	require.NoError(t, app.ForwardingKeeper.AllowedDenoms.Set(ctx, "umin"))
	require.NoError(t, app.ForwardingKeeper.AllowedDenoms.Set(ctx, "uinit"))

	finalizeAt(t, app, 2, v1_2_21.RestartAfter.Add(time.Hour))
	ctx = app.NewUncachedContext(false, cmtproto.Header{Height: 2, ChainID: app.ChainID()})
	done, err := app.UpgradeKeeper.GetDoneHeight(ctx, v1_2_21.UpgradeName)
	require.NoError(t, err)
	require.Equal(t, int64(2), done, "the upgrade still applies, it just has nothing to clear")
	require.ElementsMatch(t, []string{"umin", "uinit"}, app.ForwardingKeeper.GetAllowedDenoms(ctx))
}

func TestRestartUpgradeSkipsOtherChains(t *testing.T) {
	app := SetupWithGenesisAccounts(nil, nil)
	require.False(t, v1_2_21.RestartChains[app.ChainID()])

	ctx := app.NewUncachedContext(false, cmtproto.Header{Height: 1, ChainID: app.ChainID()})
	require.NoError(t, app.ForwardingKeeper.AllowedDenoms.Set(ctx, "*"))

	finalizeAt(t, app, 2, v1_2_21.RestartAfter.Add(time.Hour))
	ctx = app.NewUncachedContext(false, cmtproto.Header{Height: 2, ChainID: app.ChainID()})
	done, err := app.UpgradeKeeper.GetDoneHeight(ctx, v1_2_21.UpgradeName)
	require.NoError(t, err)
	require.Zero(t, done)
	require.Equal(t, []string{"*"}, app.ForwardingKeeper.GetAllowedDenoms(ctx))
}

func finalizeAt(t *testing.T, app *MinitiaApp, height int64, blockTime time.Time) {
	t.Helper()
	_, err := app.FinalizeBlock(&abci.RequestFinalizeBlock{Height: height, Time: blockTime})
	require.NoError(t, err)
	_, err = app.Commit()
	require.NoError(t, err)
}
