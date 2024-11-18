package indexer_test

import (
	"crypto/ecdsa"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	minitiaapp "github.com/initia-labs/minievm/app"
	minievmtypes "github.com/initia-labs/minievm/types"

	"github.com/stretchr/testify/require"

	"cosmossdk.io/math"
	dbm "github.com/cosmos/cosmos-db"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	abci "github.com/cometbft/cometbft/abci/types"

	"github.com/initia-labs/minievm/indexer"
)

// Bond denom should be set for staking test
const baseDenom = minievmtypes.BaseDenom

var (
	genCoins = sdk.NewCoins(sdk.NewCoin(baseDenom, math.NewInt(1_000_000_000_000_000_000).MulRaw(1_000_000))).Sort()
)

func checkBalance(t *testing.T, app *minitiaapp.MinitiaApp, addr common.Address, balances sdk.Coins) {
	ctxCheck := app.BaseApp.NewContext(true)
	require.True(t, balances.Equal(app.BankKeeper.GetAllBalances(ctxCheck, addr.Bytes())))
}

func createApp(t *testing.T) (*minitiaapp.MinitiaApp, []common.Address, []*ecdsa.PrivateKey) {
	addrs, privKeys := generateKeys(t, 2)
	genAccs := authtypes.GenesisAccounts{}
	for _, addr := range addrs {

		genAccs = append(genAccs, &authtypes.BaseAccount{Address: sdk.AccAddress(addr.Bytes()).String()})
	}

	genBalances := []banktypes.Balance{}
	for _, addr := range addrs {
		genBalances = append(genBalances, banktypes.Balance{Address: sdk.AccAddress(addr.Bytes()).String(), Coins: genCoins})
	}

	app := minitiaapp.SetupWithGenesisAccounts(nil, genAccs, genBalances...)
	for _, addr := range addrs {
		checkBalance(t, app, addr, genCoins)
	}

	_, err := app.FinalizeBlock(&abci.RequestFinalizeBlock{Height: app.LastBlockHeight() + 1})
	require.NoError(t, err)

	app.Commit()

	return app, addrs, privKeys
}

func setupIndexer(t *testing.T) (*minitiaapp.MinitiaApp, indexer.EVMIndexer, []common.Address, []*ecdsa.PrivateKey) {
	app, addrs, privKeys := createApp(t)

	db := dbm.NewMemDB()
	indexer, err := indexer.NewEVMIndexer(db, app.AppCodec(), app.Logger(), app.TxConfig(), app.EVMKeeper)
	require.NoError(t, err)

	return app, indexer, addrs, privKeys
}
