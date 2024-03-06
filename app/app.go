package app

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"reflect"

	"github.com/gorilla/mux"
	"github.com/rakyll/statik/fs"
	"github.com/spf13/cast"

	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
	reflectionv1 "cosmossdk.io/api/cosmos/reflection/v1"
	"cosmossdk.io/log"
	"cosmossdk.io/math"
	storetypes "cosmossdk.io/store/types"
	"cosmossdk.io/x/feegrant"
	feegrantkeeper "cosmossdk.io/x/feegrant/keeper"
	feegrantmodule "cosmossdk.io/x/feegrant/module"
	"cosmossdk.io/x/upgrade"
	upgradekeeper "cosmossdk.io/x/upgrade/keeper"
	upgradetypes "cosmossdk.io/x/upgrade/types"

	abci "github.com/cometbft/cometbft/abci/types"
	tmjson "github.com/cometbft/cometbft/libs/json"
	tmos "github.com/cometbft/cometbft/libs/os"

	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/grpc/cmtservice"
	nodeservice "github.com/cosmos/cosmos-sdk/client/grpc/node"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	runtimeservices "github.com/cosmos/cosmos-sdk/runtime/services"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/std"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	cosmosante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	authcodec "github.com/cosmos/cosmos-sdk/x/auth/codec"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/cosmos/cosmos-sdk/x/auth/posthandler"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	authzmodule "github.com/cosmos/cosmos-sdk/x/authz/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/consensus"
	consensusparamkeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	consensusparamtypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	genutil "github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/group"
	groupkeeper "github.com/cosmos/cosmos-sdk/x/group/keeper"
	groupmodule "github.com/cosmos/cosmos-sdk/x/group/module"
	"github.com/cosmos/gogoproto/proto"

	// ibc imports
	packetforward "github.com/cosmos/ibc-apps/middleware/packet-forward-middleware/v8/packetforward"
	packetforwardkeeper "github.com/cosmos/ibc-apps/middleware/packet-forward-middleware/v8/packetforward/keeper"
	packetforwardtypes "github.com/cosmos/ibc-apps/middleware/packet-forward-middleware/v8/packetforward/types"
	icq "github.com/cosmos/ibc-apps/modules/async-icq/v8"
	icqkeeper "github.com/cosmos/ibc-apps/modules/async-icq/v8/keeper"
	icqtypes "github.com/cosmos/ibc-apps/modules/async-icq/v8/types"
	"github.com/cosmos/ibc-go/modules/capability"
	capabilitykeeper "github.com/cosmos/ibc-go/modules/capability/keeper"
	capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"
	ica "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts"
	icacontroller "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/controller"
	icacontrollerkeeper "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/controller/keeper"
	icacontrollertypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/controller/types"
	icahost "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/host"
	icahostkeeper "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/host/keeper"
	icahosttypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/host/types"
	icatypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/types"
	ibcfee "github.com/cosmos/ibc-go/v8/modules/apps/29-fee"
	ibcfeekeeper "github.com/cosmos/ibc-go/v8/modules/apps/29-fee/keeper"
	ibcfeetypes "github.com/cosmos/ibc-go/v8/modules/apps/29-fee/types"
	ibctransfer "github.com/cosmos/ibc-go/v8/modules/apps/transfer"
	ibctransferkeeper "github.com/cosmos/ibc-go/v8/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	ibc "github.com/cosmos/ibc-go/v8/modules/core"
	porttypes "github.com/cosmos/ibc-go/v8/modules/core/05-port/types"
	ibcexported "github.com/cosmos/ibc-go/v8/modules/core/exported"
	ibckeeper "github.com/cosmos/ibc-go/v8/modules/core/keeper"
	solomachine "github.com/cosmos/ibc-go/v8/modules/light-clients/06-solomachine"
	ibctm "github.com/cosmos/ibc-go/v8/modules/light-clients/07-tendermint"

	// initia imports

	initiaapplanes "github.com/initia-labs/initia/app/lanes"
	initiaappparams "github.com/initia-labs/initia/app/params"
	"github.com/initia-labs/initia/x/ibc/fetchprice"
	fetchpricekeeper "github.com/initia-labs/initia/x/ibc/fetchprice/keeper"
	fetchpricetypes "github.com/initia-labs/initia/x/ibc/fetchprice/types"
	ibctestingtypes "github.com/initia-labs/initia/x/ibc/testing/types"
	icaauth "github.com/initia-labs/initia/x/intertx"
	icaauthkeeper "github.com/initia-labs/initia/x/intertx/keeper"
	icaauthtypes "github.com/initia-labs/initia/x/intertx/types"

	// OPinit imports
	opchild "github.com/initia-labs/OPinit/x/opchild"
	opchildkeeper "github.com/initia-labs/OPinit/x/opchild/keeper"
	opchildtypes "github.com/initia-labs/OPinit/x/opchild/types"

	// skip imports
	mevabci "github.com/skip-mev/block-sdk/abci"
	signer_extraction "github.com/skip-mev/block-sdk/adapters/signer_extraction_adapter"
	"github.com/skip-mev/block-sdk/block"
	blockbase "github.com/skip-mev/block-sdk/block/base"
	mevlane "github.com/skip-mev/block-sdk/lanes/mev"
	"github.com/skip-mev/block-sdk/x/auction"
	auctionante "github.com/skip-mev/block-sdk/x/auction/ante"
	auctionkeeper "github.com/skip-mev/block-sdk/x/auction/keeper"
	auctiontypes "github.com/skip-mev/block-sdk/x/auction/types"
	"github.com/skip-mev/slinky/x/oracle"
	oraclekeeper "github.com/skip-mev/slinky/x/oracle/keeper"
	oracletypes "github.com/skip-mev/slinky/x/oracle/types"

	// local imports
	appante "github.com/initia-labs/minievm/app/ante"
	apphook "github.com/initia-labs/minievm/app/hook"
	evmibcmiddleware "github.com/initia-labs/minievm/app/ibc-middleware"
	appkeepers "github.com/initia-labs/minievm/app/keepers"
	applanes "github.com/initia-labs/minievm/app/lanes"

	"github.com/initia-labs/minievm/x/evm"
	evmconfig "github.com/initia-labs/minievm/x/evm/config"
	evmkeeper "github.com/initia-labs/minievm/x/evm/keeper"
	evmtypes "github.com/initia-labs/minievm/x/evm/types"

	// unnamed import of statik for swagger UI support
	_ "github.com/initia-labs/minievm/client/docs/statik"
)

var (
	// DefaultNodeHome default home directories for the application daemon
	DefaultNodeHome string

	// module account permissions
	maccPerms = map[string][]string{
		authtypes.FeeCollectorName:  nil,
		icatypes.ModuleName:         nil,
		ibcfeetypes.ModuleName:      nil,
		ibctransfertypes.ModuleName: {authtypes.Minter, authtypes.Burner},
		// x/auction's module account must be instantiated upon genesis to accrue auction rewards not
		// distributed to proposers
		auctiontypes.ModuleName: nil,
		opchildtypes.ModuleName: {authtypes.Minter, authtypes.Burner},

		// slinky oracle permissions
		oracletypes.ModuleName: nil,

		// this is only for testing
		authtypes.Minter: {authtypes.Minter},
	}
)

var (
	_ servertypes.Application = (*MinitiaApp)(nil)
)

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DefaultNodeHome = filepath.Join(userHomeDir, "."+AppName)
}

// MinitiaApp extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type MinitiaApp struct {
	*baseapp.BaseApp

	legacyAmino       *codec.LegacyAmino
	appCodec          codec.Codec
	txConfig          client.TxConfig
	interfaceRegistry types.InterfaceRegistry

	// keys to access the substores
	keys    map[string]*storetypes.KVStoreKey
	tkeys   map[string]*storetypes.TransientStoreKey
	memKeys map[string]*storetypes.MemoryStoreKey

	// keepers
	AccountKeeper         *authkeeper.AccountKeeper
	BankKeeper            bankkeeper.Keeper
	CapabilityKeeper      *capabilitykeeper.Keeper
	UpgradeKeeper         *upgradekeeper.Keeper
	GroupKeeper           *groupkeeper.Keeper
	ConsensusParamsKeeper *consensusparamkeeper.Keeper
	IBCKeeper             *ibckeeper.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
	TransferKeeper        *ibctransferkeeper.Keeper
	AuthzKeeper           *authzkeeper.Keeper
	FeeGrantKeeper        *feegrantkeeper.Keeper
	ICAHostKeeper         *icahostkeeper.Keeper
	ICAControllerKeeper   *icacontrollerkeeper.Keeper
	ICAAuthKeeper         *icaauthkeeper.Keeper
	IBCFeeKeeper          *ibcfeekeeper.Keeper
	EVMKeeper             *evmkeeper.Keeper
	OPChildKeeper         *opchildkeeper.Keeper
	AuctionKeeper         *auctionkeeper.Keeper // x/auction keeper used to process bids for POB auctions
	PacketForwardKeeper   *packetforwardkeeper.Keeper
	ICQKeeper             *icqkeeper.Keeper
	OracleKeeper          *oraclekeeper.Keeper // x/oracle keeper used for the slinky oracle
	FetchPriceKeeper      *fetchpricekeeper.Keeper

	// make scoped keepers public for test purposes
	ScopedIBCKeeper           capabilitykeeper.ScopedKeeper
	ScopedTransferKeeper      capabilitykeeper.ScopedKeeper
	ScopedICAHostKeeper       capabilitykeeper.ScopedKeeper
	ScopedICAControllerKeeper capabilitykeeper.ScopedKeeper
	ScopedICAAuthKeeper       capabilitykeeper.ScopedKeeper
	ScopedICQKeeper           capabilitykeeper.ScopedKeeper
	ScopedFetchPriceKeeper    capabilitykeeper.ScopedKeeper

	// the module manager
	ModuleManager      *module.Manager
	BasicModuleManager module.BasicManager

	// the configurator
	configurator module.Configurator

	// Override of BaseApp's CheckTx
	checkTxHandler mevlane.CheckTx
}

// NewMinitiaApp returns a reference to an initialized Initia.
func NewMinitiaApp(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	moveConfig evmconfig.EVMConfig,
	appOpts servertypes.AppOptions,
	baseAppOptions ...func(*baseapp.BaseApp),
) *MinitiaApp {
	encodingConfig := initiaappparams.MakeEncodingConfig()
	std.RegisterLegacyAminoCodec(encodingConfig.Amino)
	std.RegisterInterfaces(encodingConfig.InterfaceRegistry)

	appCodec := encodingConfig.Codec
	legacyAmino := encodingConfig.Amino
	interfaceRegistry := encodingConfig.InterfaceRegistry
	txConfig := encodingConfig.TxConfig

	bApp := baseapp.NewBaseApp(AppName, logger, db, encodingConfig.TxConfig.TxDecoder(), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetVersion(version.Version)
	bApp.SetInterfaceRegistry(interfaceRegistry)
	bApp.SetTxEncoder(txConfig.TxEncoder())

	keys := storetypes.NewKVStoreKeys(
		authtypes.StoreKey, banktypes.StoreKey, group.StoreKey, consensusparamtypes.StoreKey,
		ibcexported.StoreKey, upgradetypes.StoreKey, ibctransfertypes.StoreKey,
		capabilitytypes.StoreKey, authzkeeper.StoreKey, feegrant.StoreKey,
		icahosttypes.StoreKey, icacontrollertypes.StoreKey, icaauthtypes.StoreKey,
		ibcfeetypes.StoreKey, evmtypes.StoreKey, opchildtypes.StoreKey,
		auctiontypes.StoreKey, packetforwardtypes.StoreKey, icqtypes.StoreKey,
		oracletypes.StoreKey, fetchpricetypes.StoreKey,
	)
	tkeys := storetypes.NewTransientStoreKeys()
	memKeys := storetypes.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	// register streaming services
	if err := bApp.RegisterStreamingServices(appOpts, keys); err != nil {
		panic(err)
	}

	app := &MinitiaApp{
		BaseApp:           bApp,
		legacyAmino:       legacyAmino,
		appCodec:          appCodec,
		txConfig:          txConfig,
		interfaceRegistry: interfaceRegistry,
		keys:              keys,
		tkeys:             tkeys,
		memKeys:           memKeys,
	}

	ac := authcodec.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix())
	vc := authcodec.NewBech32Codec(sdk.GetConfig().GetBech32ValidatorAddrPrefix())
	cc := authcodec.NewBech32Codec(sdk.GetConfig().GetBech32ConsensusAddrPrefix())

	authorityAccAddr := authtypes.NewModuleAddress(opchildtypes.ModuleName)
	authorityAddr, err := ac.BytesToString(authorityAccAddr)
	if err != nil {
		panic(err)
	}

	// set the BaseApp's parameter store
	consensusParamsKeeper := consensusparamkeeper.NewKeeper(appCodec, runtime.NewKVStoreService(keys[consensusparamtypes.StoreKey]), authorityAddr, runtime.EventService{})
	app.ConsensusParamsKeeper = &consensusParamsKeeper
	bApp.SetParamStore(app.ConsensusParamsKeeper.ParamsStore)

	// add capability keeper and ScopeToModule for ibc module
	app.CapabilityKeeper = capabilitykeeper.NewKeeper(appCodec, keys[capabilitytypes.StoreKey], memKeys[capabilitytypes.MemStoreKey])

	// grant capabilities for the ibc and ibc-transfer modules
	app.ScopedIBCKeeper = app.CapabilityKeeper.ScopeToModule(ibcexported.ModuleName)
	app.ScopedTransferKeeper = app.CapabilityKeeper.ScopeToModule(ibctransfertypes.ModuleName)
	app.ScopedICAHostKeeper = app.CapabilityKeeper.ScopeToModule(icahosttypes.SubModuleName)
	app.ScopedICAControllerKeeper = app.CapabilityKeeper.ScopeToModule(icacontrollertypes.SubModuleName)
	app.ScopedICAAuthKeeper = app.CapabilityKeeper.ScopeToModule(icaauthtypes.ModuleName)
	app.ScopedICQKeeper = app.CapabilityKeeper.ScopeToModule(icqtypes.ModuleName)
	app.ScopedFetchPriceKeeper = app.CapabilityKeeper.ScopeToModule(fetchpricetypes.ModuleName)

	app.CapabilityKeeper.Seal()

	// add keepers
	app.EVMKeeper = &evmkeeper.Keeper{}

	accountKeeper := authkeeper.NewAccountKeeper(
		appCodec,
		runtime.NewKVStoreService(keys[authtypes.StoreKey]),
		authtypes.ProtoBaseAccount,
		maccPerms,
		ac,
		sdk.GetConfig().GetBech32AccountAddrPrefix(),
		authorityAddr,
	)
	app.AccountKeeper = &accountKeeper

	bankKeeper := bankkeeper.NewBaseKeeper(
		appCodec,
		runtime.NewKVStoreService(keys[banktypes.StoreKey]),
		app.AccountKeeper,
		app.ModuleAccountAddrs(),
		authorityAddr,
		logger,
	)
	app.BankKeeper = &bankKeeper

	communityPoolKeeper := appkeepers.NewCommunityPoolKeeper(app.BankKeeper, authtypes.FeeCollectorName)

	////////////////////////////////
	// OPChildKeeper Configuration //
	////////////////////////////////

	app.OPChildKeeper = opchildkeeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(keys[opchildtypes.StoreKey]),
		app.AccountKeeper,
		app.BankKeeper,
		apphook.NewEVMBridgeHook(app.EVMKeeper).Hook,
		app.MsgServiceRouter(),
		authorityAddr,
		vc,
		cc,
	)

	// get skipUpgradeHeights from the app options
	skipUpgradeHeights := map[int64]bool{}
	for _, h := range cast.ToIntSlice(appOpts.Get(server.FlagUnsafeSkipUpgrades)) {
		skipUpgradeHeights[int64(h)] = true
	}
	homePath := cast.ToString(appOpts.Get(flags.FlagHome))
	app.UpgradeKeeper = upgradekeeper.NewKeeper(
		skipUpgradeHeights,
		runtime.NewKVStoreService(keys[upgradetypes.StoreKey]),
		appCodec,
		homePath,
		app.BaseApp,
		authorityAddr,
	)

	i := 0
	moduleAddrs := make([]sdk.AccAddress, len(maccPerms))
	for name := range maccPerms {
		moduleAddrs[i] = authtypes.NewModuleAddress(name)
		i += 1
	}

	feeGrantKeeper := feegrantkeeper.NewKeeper(appCodec, runtime.NewKVStoreService(keys[feegrant.StoreKey]), app.AccountKeeper)
	app.FeeGrantKeeper = &feeGrantKeeper

	authzKeeper := authzkeeper.NewKeeper(runtime.NewKVStoreService(keys[authzkeeper.StoreKey]), appCodec, app.BaseApp.MsgServiceRouter(), app.AccountKeeper)
	app.AuthzKeeper = &authzKeeper

	groupConfig := group.DefaultConfig()
	groupKeeper := groupkeeper.NewKeeper(
		keys[group.StoreKey],
		appCodec,
		app.MsgServiceRouter(),
		app.AccountKeeper,
		groupConfig,
	)
	app.GroupKeeper = &groupKeeper

	// initialize oracle keeper
	oracleKeeper := oraclekeeper.NewKeeper(
		runtime.NewKVStoreService(keys[oracletypes.StoreKey]),
		appCodec,
		authorityAccAddr,
	)
	app.OracleKeeper = &oracleKeeper

	// Create IBC Keeper
	app.IBCKeeper = ibckeeper.NewKeeper(
		appCodec,
		keys[ibcexported.StoreKey],
		nil, // we don't need migration
		app.OPChildKeeper,
		app.UpgradeKeeper,
		app.ScopedIBCKeeper,
		authorityAddr,
	)

	ibcFeeKeeper := ibcfeekeeper.NewKeeper(
		appCodec,
		keys[ibcfeetypes.StoreKey],
		app.IBCKeeper.ChannelKeeper,
		app.IBCKeeper.ChannelKeeper,
		app.IBCKeeper.PortKeeper,
		app.AccountKeeper,
		app.BankKeeper,
	)
	app.IBCFeeKeeper = &ibcFeeKeeper

	////////////////////////////
	// Transfer configuration //
	////////////////////////////
	// Send   : transfer -> packet forward -> evm   -> fee            -> channel
	// Receive: channel  -> fee            -> evm   -> packet forward -> transfer

	var transferStack porttypes.IBCModule
	{
		packetForwardKeeper := &packetforwardkeeper.Keeper{}
		evmMiddleware := &evmibcmiddleware.IBCMiddleware{}

		// Create Transfer Keepers
		transferKeeper := ibctransferkeeper.NewKeeper(
			appCodec,
			keys[ibctransfertypes.StoreKey],
			nil, // we don't need migration
			// ics4wrapper: transfer -> packet forward
			packetForwardKeeper,
			app.IBCKeeper.ChannelKeeper,
			app.IBCKeeper.PortKeeper,
			app.AccountKeeper,
			app.BankKeeper,
			app.ScopedTransferKeeper,
			authorityAddr,
		)
		app.TransferKeeper = &transferKeeper
		transferIBCModule := ibctransfer.NewIBCModule(*app.TransferKeeper)

		// create packet forward middleware
		*packetForwardKeeper = *packetforwardkeeper.NewKeeper(
			appCodec,
			keys[packetforwardtypes.StoreKey],
			app.TransferKeeper,
			app.IBCKeeper.ChannelKeeper,
			communityPoolKeeper,
			app.BankKeeper,
			// ics4wrapper: transfer -> packet forward -> evm
			evmMiddleware,
			authorityAddr,
		)
		app.PacketForwardKeeper = packetForwardKeeper
		packetForwardMiddleware := packetforward.NewIBCMiddleware(
			// receive: packet forward -> transfer
			transferIBCModule,
			app.PacketForwardKeeper,
			0,
			packetforwardkeeper.DefaultForwardTransferPacketTimeoutTimestamp,
			packetforwardkeeper.DefaultRefundTransferPacketTimeoutTimestamp,
		)

		// create move middleware for transfer
		*evmMiddleware = evmibcmiddleware.NewIBCMiddleware(
			// receive: evm -> packet forward -> transfer
			packetForwardMiddleware,
			// ics4wrapper: transfer -> packet forward -> evm -> fee
			app.IBCFeeKeeper,
			app.EVMKeeper,
		)

		// create ibcfee middleware for transfer
		transferStack = ibcfee.NewIBCMiddleware(
			// receive: fee -> evm -> packet forward -> transfer
			evmMiddleware,
			// ics4wrapper: transfer -> packet forward -> evm -> fee -> channel
			*app.IBCFeeKeeper,
		)
	}

	///////////////////////
	// ICA configuration //
	///////////////////////

	var icaHostStack porttypes.IBCModule
	var icaControllerStack porttypes.IBCModule
	{
		icaHostKeeper := icahostkeeper.NewKeeper(
			appCodec, keys[icahosttypes.StoreKey],
			nil, // we don't need migration
			app.IBCFeeKeeper,
			app.IBCKeeper.ChannelKeeper,
			app.IBCKeeper.PortKeeper,
			app.AccountKeeper,
			app.ScopedICAHostKeeper,
			app.MsgServiceRouter(),
			authorityAddr,
		)
		app.ICAHostKeeper = &icaHostKeeper

		icaControllerKeeper := icacontrollerkeeper.NewKeeper(
			appCodec, keys[icacontrollertypes.StoreKey],
			nil, // we don't need migration
			app.IBCFeeKeeper,
			app.IBCKeeper.ChannelKeeper,
			app.IBCKeeper.PortKeeper,
			app.ScopedICAControllerKeeper,
			app.MsgServiceRouter(),
			authorityAddr,
		)
		app.ICAControllerKeeper = &icaControllerKeeper

		icaAuthKeeper := icaauthkeeper.NewKeeper(
			appCodec,
			*app.ICAControllerKeeper,
			app.ScopedICAAuthKeeper,
			ac,
		)
		app.ICAAuthKeeper = &icaAuthKeeper

		icaAuthIBCModule := icaauth.NewIBCModule(*app.ICAAuthKeeper)
		icaHostIBCModule := icahost.NewIBCModule(*app.ICAHostKeeper)
		icaHostStack = ibcfee.NewIBCMiddleware(icaHostIBCModule, *app.IBCFeeKeeper)
		icaControllerIBCModule := icacontroller.NewIBCMiddleware(icaAuthIBCModule, *app.ICAControllerKeeper)
		icaControllerStack = ibcfee.NewIBCMiddleware(icaControllerIBCModule, *app.IBCFeeKeeper)
	}

	///////////////////////
	// ICQ configuration //
	///////////////////////
	var icqStack porttypes.IBCModule
	{
		icqKeeper := icqkeeper.NewKeeper(
			appCodec,
			app.keys[icqtypes.StoreKey],
			// ics4wrapper: icq -> fee -> channel
			app.IBCFeeKeeper,
			app.IBCKeeper.ChannelKeeper,
			app.IBCKeeper.PortKeeper,
			app.ScopedICQKeeper,
			bApp.GRPCQueryRouter(),
			authorityAddr,
		)
		app.ICQKeeper = &icqKeeper

		// Create Async ICQ module
		icqModule := icq.NewIBCModule(*app.ICQKeeper)
		icqStack = ibcfee.NewIBCMiddleware(icqModule, *app.IBCFeeKeeper)
	}

	//////////////////////////////
	// FetchPrice configuration //
	//////////////////////////////

	var fetchpriceStack porttypes.IBCModule
	{
		app.FetchPriceKeeper = fetchpricekeeper.NewKeeper(
			appCodec,
			runtime.NewKVStoreService(app.keys[fetchpricetypes.StoreKey]),
			// ics4wrapper: fetchprice -> fee -> channel
			app.IBCFeeKeeper,
			app.IBCKeeper.ChannelKeeper,
			app.IBCKeeper.PortKeeper,
			app.AccountKeeper,
			app.OracleKeeper,
			app.ScopedFetchPriceKeeper,
			authorityAddr,
		)

		// Create FetchPrice module
		fetchpriceModule := fetchprice.NewIBCModule(*app.FetchPriceKeeper)
		fetchpriceStack = ibcfee.NewIBCMiddleware(
			// receive: fee -> fetchprice
			fetchpriceModule,
			*app.IBCFeeKeeper,
		)
	}

	//////////////////////////////
	// IBC router Configuration //
	//////////////////////////////

	// Create static IBC router, add transfer route, then set and seal it
	ibcRouter := porttypes.NewRouter()
	ibcRouter.AddRoute(ibctransfertypes.ModuleName, transferStack).
		AddRoute(icahosttypes.SubModuleName, icaHostStack).
		AddRoute(icacontrollertypes.SubModuleName, icaControllerStack).
		AddRoute(icaauthtypes.ModuleName, icaControllerStack).
		AddRoute(icqtypes.ModuleName, icqStack).
		AddRoute(fetchpricetypes.ModuleName, fetchpriceStack)

	app.IBCKeeper.SetRouter(ibcRouter)

	//////////////////////////////
	// WasmKeeper Configuration //
	//////////////////////////////
	evmConfig := evmconfig.GetConfig(appOpts)

	// The last arguments can contain custom message handlers, and custom query handlers,
	// if we want to allow any custom callbacks
	app.EVMKeeper = evmkeeper.NewKeeper(
		ac,
		appCodec,
		runtime.NewKVStoreService(keys[evmtypes.StoreKey]),
		authorityAddr,
		evmConfig,
	)

	// x/auction module keeper initialization

	// initialize the keeper
	auctionKeeper := auctionkeeper.NewKeeperWithRewardsAddressProvider(
		app.appCodec,
		app.keys[auctiontypes.StoreKey],
		app.AccountKeeper,
		app.BankKeeper,
		applanes.NewRewardsAddressProvider(authtypes.FeeCollectorName),
		authorityAddr,
	)
	app.AuctionKeeper = &auctionKeeper

	/****  Module Options ****/

	// NOTE: we may consider parsing `appOpts` inside module constructors. For the moment
	// we prefer to be more strict in what arguments the modules expect.

	// TODO - add crisis module
	// skipGenesisInvariants := cast.ToBool(appOpts.Get(crisis.FlagSkipGenesisInvariants))

	// NOTE: Any module instantiated in the module manager that is later modified
	// must be passed by reference here.

	app.ModuleManager = module.NewManager(
		auth.NewAppModule(appCodec, *app.AccountKeeper, nil, nil),
		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper, nil),
		opchild.NewAppModule(appCodec, *app.OPChildKeeper),
		capability.NewAppModule(appCodec, *app.CapabilityKeeper, false),
		feegrantmodule.NewAppModule(appCodec, app.AccountKeeper, app.BankKeeper, *app.FeeGrantKeeper, app.interfaceRegistry),
		upgrade.NewAppModule(app.UpgradeKeeper, ac),
		authzmodule.NewAppModule(appCodec, *app.AuthzKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
		groupmodule.NewAppModule(appCodec, *app.GroupKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
		consensus.NewAppModule(appCodec, *app.ConsensusParamsKeeper),
		evm.NewAppModule(appCodec, app.EVMKeeper),
		auction.NewAppModule(app.appCodec, *app.AuctionKeeper),
		// ibc modules
		ibc.NewAppModule(app.IBCKeeper),
		ibctransfer.NewAppModule(*app.TransferKeeper),
		ica.NewAppModule(app.ICAControllerKeeper, app.ICAHostKeeper),
		icaauth.NewAppModule(appCodec, *app.ICAAuthKeeper),
		ibcfee.NewAppModule(*app.IBCFeeKeeper),
		ibctm.NewAppModule(),
		solomachine.NewAppModule(),
		packetforward.NewAppModule(app.PacketForwardKeeper, nil),
		icq.NewAppModule(*app.ICQKeeper, nil),
		fetchprice.NewAppModule(appCodec, *app.FetchPriceKeeper),
		// slinky modules
		oracle.NewAppModule(appCodec, *app.OracleKeeper),
	)

	// BasicModuleManager defines the module BasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration and genesis verification.
	// By default it is composed of all the module from the module manager.
	// Additionally, app module basics can be overwritten by passing them as argument.
	app.BasicModuleManager = module.NewBasicManagerFromManager(
		app.ModuleManager,
		map[string]module.AppModuleBasic{
			genutiltypes.ModuleName: genutil.NewAppModuleBasic(genutiltypes.DefaultMessageValidator),
		})
	app.BasicModuleManager.RegisterLegacyAminoCodec(legacyAmino)
	app.BasicModuleManager.RegisterInterfaces(interfaceRegistry)

	// NOTE: upgrade module is required to be prioritized
	app.ModuleManager.SetOrderPreBlockers(
		upgradetypes.ModuleName,
	)

	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	// NOTE: staking module is required if HistoricalEntries param > 0
	app.ModuleManager.SetOrderBeginBlockers(
		capabilitytypes.ModuleName,
		opchildtypes.ModuleName,
		authz.ModuleName,
		ibcexported.ModuleName,
	)

	app.ModuleManager.SetOrderEndBlockers(
		opchildtypes.ModuleName,
		authz.ModuleName,
		feegrant.ModuleName,
		group.ModuleName,
	)

	// NOTE: The genutils module must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	// NOTE: Capability module must occur first so that it can initialize any capabilities
	// so that other modules that want to create or claim capabilities afterwards in InitChain
	// can do so safely.
	genesisModuleOrder := []string{
		capabilitytypes.ModuleName, authtypes.ModuleName, banktypes.ModuleName,
		opchildtypes.ModuleName, genutiltypes.ModuleName, authz.ModuleName, group.ModuleName,
		upgradetypes.ModuleName, feegrant.ModuleName, consensusparamtypes.ModuleName, ibcexported.ModuleName,
		ibctransfertypes.ModuleName, icatypes.ModuleName, icaauthtypes.ModuleName,
		ibcfeetypes.ModuleName, auctiontypes.ModuleName,
		evmtypes.ModuleName, oracletypes.ModuleName, packetforwardtypes.ModuleName,
		icqtypes.ModuleName, fetchpricetypes.ModuleName,
	}

	app.ModuleManager.SetOrderInitGenesis(genesisModuleOrder...)
	app.ModuleManager.SetOrderExportGenesis(genesisModuleOrder...)

	// TODO - crisis keeper
	// app.ModuleManager.RegisterInvariants(app.CrisisKeeper)

	app.configurator = module.NewConfigurator(app.appCodec, app.MsgServiceRouter(), app.GRPCQueryRouter())
	err = app.ModuleManager.RegisterServices(app.configurator)
	if err != nil {
		panic(err)
	}

	// register upgrade handler for later use
	app.RegisterUpgradeHandlers(app.configurator)

	autocliv1.RegisterQueryServer(app.GRPCQueryRouter(), runtimeservices.NewAutoCLIQueryService(app.ModuleManager.Modules))

	reflectionSvc, err := runtimeservices.NewReflectionService()
	if err != nil {
		panic(err)
	}
	reflectionv1.RegisterReflectionServiceServer(app.GRPCQueryRouter(), reflectionSvc)

	// initialize stores
	app.MountKVStores(keys)
	app.MountTransientStores(tkeys)
	app.MountMemoryStores(memKeys)

	// initialize BaseApp
	app.SetInitChainer(app.InitChainer)
	app.SetPreBlocker(app.PreBlocker)
	app.SetBeginBlocker(app.BeginBlocker)
	app.setPostHandler()
	app.SetEndBlocker(app.EndBlocker)

	// initialize and set the InitiaApp mempool. The current mempool will be the
	// x/auction module's mempool which will extract the top bid from the current block's auction
	// and insert the txs at the top of the block spots.
	signerExtractor := signer_extraction.NewDefaultAdapter()

	mevConfig := blockbase.LaneConfig{
		Logger:          app.Logger(),
		TxEncoder:       app.txConfig.TxEncoder(),
		TxDecoder:       app.txConfig.TxDecoder(),
		MaxBlockSpace:   math.LegacyZeroDec(),
		MaxTxs:          0,
		SignerExtractor: signerExtractor,
	}
	factor := mevlane.NewDefaultAuctionFactory(app.txConfig.TxDecoder(), signerExtractor)
	mevLane := mevlane.NewMEVLane(
		mevConfig,
		factor,
		factor.MatchHandler(),
	)

	freeConfig := blockbase.LaneConfig{
		Logger:          app.Logger(),
		TxEncoder:       app.txConfig.TxEncoder(),
		TxDecoder:       app.txConfig.TxDecoder(),
		MaxBlockSpace:   math.LegacyZeroDec(),
		MaxTxs:          100,
		SignerExtractor: signerExtractor,
	}
	freeLane := initiaapplanes.NewFreeLane(freeConfig, applanes.FreeLaneMatchHandler())

	defaultLaneConfig := blockbase.LaneConfig{
		Logger:          app.Logger(),
		TxEncoder:       app.txConfig.TxEncoder(),
		TxDecoder:       app.txConfig.TxDecoder(),
		MaxBlockSpace:   math.LegacyZeroDec(),
		MaxTxs:          0,
		SignerExtractor: signerExtractor,
	}
	defaultLane := initiaapplanes.NewDefaultLane(defaultLaneConfig)

	lanes := []block.Lane{mevLane, freeLane, defaultLane}
	mempool, err := block.NewLanedMempool(app.Logger(), lanes)
	if err != nil {
		panic(err)
	}

	app.SetMempool(mempool)
	anteHandler := app.setAnteHandler(mevLane, freeLane)

	// override the base-app's ABCI methods (CheckTx, PrepareProposal, ProcessProposal)
	proposalHandlers := mevabci.NewProposalHandler(
		app.Logger(),
		app.txConfig.TxDecoder(),
		app.txConfig.TxEncoder(),
		mempool,
	)

	// override base-app's ProcessProposal + PrepareProposal
	app.SetPrepareProposal(proposalHandlers.PrepareProposalHandler())
	app.SetProcessProposal(proposalHandlers.ProcessProposalHandler())

	// overrde base-app's CheckTx
	checkTxHandler := mevlane.NewCheckTxHandler(
		app.BaseApp,
		app.txConfig.TxDecoder(),
		mevLane,
		anteHandler,
	)
	app.SetCheckTx(checkTxHandler.CheckTx())

	// At startup, after all modules have been registered, check that all prot
	// annotations are correct.
	protoFiles, err := proto.MergedRegistry()
	if err != nil {
		panic(err)
	}
	err = msgservice.ValidateProtoAnnotations(protoFiles)
	if err != nil {
		// Once we switch to using protoreflect-based antehandlers, we might
		// want to panic here instead of logging a warning.
		fmt.Fprintln(os.Stderr, err.Error())
	}

	// Load the latest state from disk if necessary, and initialize the base-app. From this point on
	// no more modifications to the base-app can be made
	if loadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			tmos.Exit(err.Error())
		}
	}

	return app
}

// CheckTx will check the transaction with the provided checkTxHandler. We override the default
// handler so that we can verify bid transactions before they are inserted into the mempool.
// With the POB CheckTx, we can verify the bid transaction and all of the bundled transactions
// before inserting the bid transaction into the mempool.
func (app *MinitiaApp) CheckTx(req *abci.RequestCheckTx) (*abci.ResponseCheckTx, error) {
	return app.checkTxHandler(req)
}

// SetCheckTx sets the checkTxHandler for the app.
func (app *MinitiaApp) SetCheckTx(handler mevlane.CheckTx) {
	app.checkTxHandler = handler
}

func (app *MinitiaApp) setAnteHandler(
	mevLane auctionante.MEVLane,
	freeLane block.Lane,
) sdk.AnteHandler {
	anteHandler, err := appante.NewAnteHandler(
		appante.HandlerOptions{
			HandlerOptions: cosmosante.HandlerOptions{
				AccountKeeper:   app.AccountKeeper,
				BankKeeper:      app.BankKeeper,
				FeegrantKeeper:  app.FeeGrantKeeper,
				SignModeHandler: app.txConfig.SignModeHandler(),
			},
			IBCkeeper:     app.IBCKeeper,
			Codec:         app.appCodec,
			OPChildKeeper: app.OPChildKeeper,
			TxEncoder:     app.txConfig.TxEncoder(),
			AuctionKeeper: *app.AuctionKeeper,
			MevLane:       mevLane,
			FreeLane:      freeLane,
		},
	)
	if err != nil {
		panic(err)
	}

	app.SetAnteHandler(anteHandler)
	return anteHandler
}

func (app *MinitiaApp) setPostHandler() {
	postHandler, err := posthandler.NewPostHandler(
		posthandler.HandlerOptions{},
	)
	if err != nil {
		panic(err)
	}

	app.SetPostHandler(postHandler)
}

// Name returns the name of the App
func (app *MinitiaApp) Name() string { return app.BaseApp.Name() }

// PreBlocker application updates every pre block
func (app *MinitiaApp) PreBlocker(ctx sdk.Context, _ *abci.RequestFinalizeBlock) (*sdk.ResponsePreBlock, error) {
	return app.ModuleManager.PreBlock(ctx)
}

// BeginBlocker application updates every begin block
func (app *MinitiaApp) BeginBlocker(ctx sdk.Context) (sdk.BeginBlock, error) {
	return app.ModuleManager.BeginBlock(ctx)
}

// EndBlocker application updates every end block
func (app *MinitiaApp) EndBlocker(ctx sdk.Context) (sdk.EndBlock, error) {
	return app.ModuleManager.EndBlock(ctx)
}

// InitChainer application update at chain initialization
func (app *MinitiaApp) InitChainer(ctx sdk.Context, req *abci.RequestInitChain) (*abci.ResponseInitChain, error) {
	var genesisState GenesisState
	if err := tmjson.Unmarshal(req.AppStateBytes, &genesisState); err != nil {
		panic(err)
	}
	app.UpgradeKeeper.SetModuleVersionMap(ctx, app.ModuleManager.GetVersionMap())
	return app.ModuleManager.InitGenesis(ctx, app.appCodec, genesisState)
}

// LoadHeight loads a particular height
func (app *MinitiaApp) LoadHeight(height int64) error {
	return app.LoadVersion(height)
}

// ModuleAccountAddrs returns all the app's module account addresses.
func (app *MinitiaApp) ModuleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

// LegacyAmino returns SimApp's amino codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *MinitiaApp) LegacyAmino() *codec.LegacyAmino {
	return app.legacyAmino
}

// AppCodec returns Initia's app codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *MinitiaApp) AppCodec() codec.Codec {
	return app.appCodec
}

// InterfaceRegistry returns Initia's InterfaceRegistry
func (app *MinitiaApp) InterfaceRegistry() types.InterfaceRegistry {
	return app.interfaceRegistry
}

// GetKey returns the KVStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *MinitiaApp) GetKey(storeKey string) *storetypes.KVStoreKey {
	return app.keys[storeKey]
}

// GetTKey returns the TransientStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *MinitiaApp) GetTKey(storeKey string) *storetypes.TransientStoreKey {
	return app.tkeys[storeKey]
}

// GetMemKey returns the MemStoreKey for the provided mem key.
//
// NOTE: This is solely used for testing purposes.
func (app *MinitiaApp) GetMemKey(storeKey string) *storetypes.MemoryStoreKey {
	return app.memKeys[storeKey]
}

// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app *MinitiaApp) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	clientCtx := apiSvr.ClientCtx

	// Register new tx routes from grpc-gateway.
	authtx.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// Register new tendermint queries routes from grpc-gateway.
	cmtservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// Register node gRPC service for grpc-gateway.
	nodeservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// Register grpc-gateway routes for all modules.
	app.BasicModuleManager.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// register swagger API from root so that other applications can override easily
	if apiConfig.Swagger {
		RegisterSwaggerAPI(apiSvr.Router)
	}
}

// Simulate customize gas simulation to add fee deduction gas amount.
func (app *MinitiaApp) Simulate(txBytes []byte) (sdk.GasInfo, *sdk.Result, error) {
	gasInfo, result, err := app.BaseApp.Simulate(txBytes)
	gasInfo.GasUsed += FeeDeductionGasAmount
	return gasInfo, result, err
}

// RegisterTxService implements the Application.RegisterTxService method.
func (app *MinitiaApp) RegisterTxService(clientCtx client.Context) {
	authtx.RegisterTxService(
		app.BaseApp.GRPCQueryRouter(), clientCtx,
		app.Simulate, app.interfaceRegistry,
	)
}

// RegisterTendermintService implements the Application.RegisterTendermintService method.
func (app *MinitiaApp) RegisterTendermintService(clientCtx client.Context) {
	cmtservice.RegisterTendermintService(
		clientCtx,
		app.BaseApp.GRPCQueryRouter(),
		app.interfaceRegistry, app.Query,
	)
}

func (app *MinitiaApp) RegisterNodeService(clientCtx client.Context, cfg config.Config) {
	nodeservice.RegisterNodeService(clientCtx, app.GRPCQueryRouter(), cfg)
}

// RegisterSwaggerAPI registers swagger route with API Server
func RegisterSwaggerAPI(rtr *mux.Router) {
	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}

	staticServer := http.FileServer(statikFS)
	rtr.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", staticServer))
}

// GetMaccPerms returns a copy of the module account permissions
func GetMaccPerms() map[string][]string {
	dupMaccPerms := make(map[string][]string)
	for k, v := range maccPerms {
		dupMaccPerms[k] = v
	}
	return dupMaccPerms
}

//////////////////////////////////////
// TestingApp functions

// GetBaseApp implements the TestingApp interface.
func (app *MinitiaApp) GetBaseApp() *baseapp.BaseApp {
	return app.BaseApp
}

// GetAccountKeeper implements the TestingApp interface.
func (app *MinitiaApp) GetAccountKeeper() *authkeeper.AccountKeeper {
	return app.AccountKeeper
}

// GetStakingKeeper implements the TestingApp interface.
// It returns opchild instead of original staking keeper.
func (app *MinitiaApp) GetStakingKeeper() ibctestingtypes.StakingKeeper {
	return app.OPChildKeeper
}

// GetIBCKeeper implements the TestingApp interface.
func (app *MinitiaApp) GetIBCKeeper() *ibckeeper.Keeper {
	return app.IBCKeeper
}

// GetICAControllerKeeper implements the TestingApp interface.
func (app *MinitiaApp) GetICAControllerKeeper() *icacontrollerkeeper.Keeper {
	return app.ICAControllerKeeper
}

// GetICAAuthKeeper implements the TestingApp interface.
func (app *MinitiaApp) GetICAAuthKeeper() *icaauthkeeper.Keeper {
	return app.ICAAuthKeeper
}

// GetScopedIBCKeeper implements the TestingApp interface.
func (app *MinitiaApp) GetScopedIBCKeeper() capabilitykeeper.ScopedKeeper {
	return app.ScopedIBCKeeper
}

// TxConfig implements the TestingApp interface.
func (app *MinitiaApp) TxConfig() client.TxConfig {
	return app.txConfig
}

// ChainID gets chainID from private fields of BaseApp
// Should be removed once SDK 0.50.x will be adopted
func (app *MinitiaApp) ChainID() string { // TODO: remove this method once chain updates to v0.50.x
	field := reflect.ValueOf(app.BaseApp).Elem().FieldByName("chainID")
	return field.String()
}
