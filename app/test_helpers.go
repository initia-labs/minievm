package app

// DONTCOVER

import (
	"encoding/json"
	"time"

	"cosmossdk.io/log"
	abci "github.com/cometbft/cometbft/abci/types"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	tmtypes "github.com/cometbft/cometbft/types"
	dbm "github.com/cosmos/cosmos-db"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	opchildtypes "github.com/initia-labs/OPinit/x/opchild/types"

	"github.com/initia-labs/minievm/types"
	evmconfig "github.com/initia-labs/minievm/x/evm/config"
	evmtypes "github.com/initia-labs/minievm/x/evm/types"
)

// defaultConsensusParams defines the default Tendermint consensus params used in
// MinitiaApp testing.
var defaultConsensusParams = &tmproto.ConsensusParams{
	Block: &tmproto.BlockParams{
		MaxBytes: 8000000,
		MaxGas:   1234000000,
	},
	Evidence: &tmproto.EvidenceParams{
		MaxAgeNumBlocks: 302400,
		MaxAgeDuration:  504 * time.Hour, // 3 weeks is the max duration
		MaxBytes:        10000,
	},
	Validator: &tmproto.ValidatorParams{
		PubKeyTypes: []string{
			tmtypes.ABCIPubKeyTypeEd25519,
		},
	},
}

func getOrCreateMemDB(db *dbm.DB) dbm.DB {
	if db != nil {
		return *db
	}
	return dbm.NewMemDB()
}

func setup(db *dbm.DB, withGenesis bool) (*MinitiaApp, GenesisState) {
	encCdc := MakeEncodingConfig()
	app := NewMinitiaApp(
		log.NewNopLogger(),
		getOrCreateMemDB(db),
		dbm.NewMemDB(),
		nil,
		true,
		evmconfig.DefaultEVMConfig(),
		EmptyAppOptions{},
	)

	_ = app.InitializeIndexer(client.Context{})

	if withGenesis {
		return app, NewDefaultGenesisState(encCdc.Codec, app.BasicModuleManager, types.BaseDenom)
	}

	return app, GenesisState{}
}

// SetupWithGenesisAccounts setup initiaapp with genesis account
func SetupWithGenesisAccounts(
	valSet *tmtypes.ValidatorSet,
	genAccs []authtypes.GenesisAccount,
	balances ...banktypes.Balance,
) *MinitiaApp {
	app, genesisState := setup(nil, true)

	if len(genAccs) == 0 {
		privAcc := secp256k1.GenPrivKey()
		genAccs = []authtypes.GenesisAccount{
			authtypes.NewBaseAccount(privAcc.PubKey().Address().Bytes(), privAcc.PubKey(), 0, 0),
		}
	}

	// set genesis accounts
	authGenesis := authtypes.NewGenesisState(authtypes.DefaultParams(), genAccs)
	genesisState[authtypes.ModuleName] = app.AppCodec().MustMarshalJSON(authGenesis)

	// allow empty validator
	if valSet == nil || len(valSet.Validators) == 0 {
		privVal := ed25519.GenPrivKey()
		pubKey, err := cryptocodec.ToTmPubKeyInterface(privVal.PubKey()) //nolint:staticcheck
		if err != nil {
			panic(err)
		}

		validator := tmtypes.NewValidator(pubKey, 1)
		valSet = tmtypes.NewValidatorSet([]*tmtypes.Validator{validator})
	}

	validators := make([]opchildtypes.Validator, 0, len(valSet.Validators))
	for _, val := range valSet.Validators {
		pk, err := cryptocodec.FromTmPubKeyInterface(val.PubKey) //nolint:staticcheck
		if err != nil {
			panic(err)
		}
		pkAny, err := codectypes.NewAnyWithValue(pk)
		if err != nil {
			panic(err)
		}

		validator := opchildtypes.Validator{
			Moniker:         "test-validator",
			OperatorAddress: sdk.ValAddress(val.Address).String(),
			ConsensusPubkey: pkAny,
			ConsPower:       1,
		}

		validators = append(validators, validator)
	}

	// set validators and delegations
	var opchildGenesis opchildtypes.GenesisState
	app.AppCodec().MustUnmarshalJSON(genesisState[opchildtypes.ModuleName], &opchildGenesis)
	opchildGenesis.Params.Admin = sdk.AccAddress(valSet.Validators[0].Address.Bytes()).String()
	opchildGenesis.Params.BridgeExecutors = []string{sdk.AccAddress(valSet.Validators[0].Address.Bytes()).String()}
	opchildGenesis.Params.MinGasPrices = sdk.NewDecCoins(sdk.NewDecCoin(types.BaseDenom, math.NewInt(1_000_000_000)))

	// set validators and delegations
	opchildGenesis = *opchildtypes.NewGenesisState(opchildGenesis.Params, validators, nil)
	genesisState[opchildtypes.ModuleName] = app.AppCodec().MustMarshalJSON(&opchildGenesis)

	// set evm genesis params
	var evmGenesis evmtypes.GenesisState
	app.AppCodec().MustUnmarshalJSON(genesisState[evmtypes.ModuleName], &evmGenesis)
	evmGenesis.Params.GasRefundRatio = math.LegacyZeroDec()
	genesisState[evmtypes.ModuleName] = app.AppCodec().MustMarshalJSON(&evmGenesis)

	// update total supply
	bankGenesis := banktypes.NewGenesisState(banktypes.DefaultGenesisState().Params, balances, sdk.NewCoins(), []banktypes.Metadata{}, []banktypes.SendEnabled{})
	genesisState[banktypes.ModuleName] = app.AppCodec().MustMarshalJSON(bankGenesis)

	stateBytes, err := json.MarshalIndent(genesisState, "", " ")
	if err != nil {
		panic(err)
	}

	_, err = app.InitChain(
		&abci.RequestInitChain{
			Validators:      []abci.ValidatorUpdate{},
			ConsensusParams: defaultConsensusParams,
			AppStateBytes:   stateBytes,
		},
	)
	if err != nil {
		panic(err)
	}

	_, err = app.FinalizeBlock(&abci.RequestFinalizeBlock{Height: 1})
	if err != nil {
		panic(err)
	}

	_, err = app.Commit()
	if err != nil {
		panic(err)
	}

	return app
}
