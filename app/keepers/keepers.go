package keepers

import (
	"os"

	"cosmossdk.io/core/address"
	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	"cosmossdk.io/x/feegrant"
	feegrantkeeper "cosmossdk.io/x/feegrant/keeper"
	upgradekeeper "cosmossdk.io/x/upgrade/keeper"
	upgradetypes "cosmossdk.io/x/upgrade/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/runtime"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	consensusparamkeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	consensusparamtypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	"github.com/cosmos/cosmos-sdk/x/group"
	groupkeeper "github.com/cosmos/cosmos-sdk/x/group/keeper"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	// ibc imports
	packetforward "github.com/cosmos/ibc-apps/middleware/packet-forward-middleware/v8/packetforward"
	packetforwardkeeper "github.com/cosmos/ibc-apps/middleware/packet-forward-middleware/v8/packetforward/keeper"
	packetforwardtypes "github.com/cosmos/ibc-apps/middleware/packet-forward-middleware/v8/packetforward/types"
	ratelimit "github.com/cosmos/ibc-apps/modules/rate-limiting/v8"
	ratelimitkeeper "github.com/cosmos/ibc-apps/modules/rate-limiting/v8/keeper"
	ratelimittypes "github.com/cosmos/ibc-apps/modules/rate-limiting/v8/types"
	capabilitykeeper "github.com/cosmos/ibc-go/modules/capability/keeper"
	capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"
	icacontroller "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/controller"
	icacontrollerkeeper "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/controller/keeper"
	icacontrollertypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/controller/types"
	icahost "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/host"
	icahostkeeper "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/host/keeper"
	icahosttypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/host/types"
	ibcfee "github.com/cosmos/ibc-go/v8/modules/apps/29-fee"
	ibcfeekeeper "github.com/cosmos/ibc-go/v8/modules/apps/29-fee/keeper"
	ibcfeetypes "github.com/cosmos/ibc-go/v8/modules/apps/29-fee/types"
	ibctransfer "github.com/cosmos/ibc-go/v8/modules/apps/transfer"
	ibctransferkeeper "github.com/cosmos/ibc-go/v8/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	porttypes "github.com/cosmos/ibc-go/v8/modules/core/05-port/types"
	ibcexported "github.com/cosmos/ibc-go/v8/modules/core/exported"
	ibckeeper "github.com/cosmos/ibc-go/v8/modules/core/keeper"

	// initia imports
	appheaderinfo "github.com/initia-labs/initia/app/header_info"
	ibchooks "github.com/initia-labs/initia/x/ibc-hooks"
	ibchookskeeper "github.com/initia-labs/initia/x/ibc-hooks/keeper"
	ibchookstypes "github.com/initia-labs/initia/x/ibc-hooks/types"
	ibcnfttransfer "github.com/initia-labs/initia/x/ibc/nft-transfer"
	ibcnfttransferkeeper "github.com/initia-labs/initia/x/ibc/nft-transfer/keeper"
	ibcnfttransfertypes "github.com/initia-labs/initia/x/ibc/nft-transfer/types"
	icaauth "github.com/initia-labs/initia/x/intertx"
	icaauthkeeper "github.com/initia-labs/initia/x/intertx/keeper"
	icaauthtypes "github.com/initia-labs/initia/x/intertx/types"

	// OPinit imports
	opchildkeeper "github.com/initia-labs/OPinit/x/opchild/keeper"
	opchildlanes "github.com/initia-labs/OPinit/x/opchild/lanes"
	opchildtypes "github.com/initia-labs/OPinit/x/opchild/types"

	// skip imports
	auctionkeeper "github.com/skip-mev/block-sdk/v2/x/auction/keeper"
	auctiontypes "github.com/skip-mev/block-sdk/v2/x/auction/types"
	marketmapkeeper "github.com/skip-mev/connect/v2/x/marketmap/keeper"
	marketmaptypes "github.com/skip-mev/connect/v2/x/marketmap/types"
	oraclekeeper "github.com/skip-mev/connect/v2/x/oracle/keeper"
	oracletypes "github.com/skip-mev/connect/v2/x/oracle/types"

	// local imports
	"github.com/initia-labs/minievm/app/ante"
	ibcevmhooks "github.com/initia-labs/minievm/app/ibc-hooks"
	bankkeeper "github.com/initia-labs/minievm/x/bank/keeper"
	evmconfig "github.com/initia-labs/minievm/x/evm/config"
	evmkeeper "github.com/initia-labs/minievm/x/evm/keeper"
	evmtypes "github.com/initia-labs/minievm/x/evm/types"

	// noble forwarding keeper
	forwarding "github.com/noble-assets/forwarding/v2/x/forwarding"
	forwardingkeeper "github.com/noble-assets/forwarding/v2/x/forwarding/keeper"
	forwardingtypes "github.com/noble-assets/forwarding/v2/x/forwarding/types"
)

type AppKeepers struct {
	// keys to access the substores
	keys    map[string]*storetypes.KVStoreKey
	tkeys   map[string]*storetypes.TransientStoreKey
	memKeys map[string]*storetypes.MemoryStoreKey

	// keepers
	AccountKeeper         *authkeeper.AccountKeeper
	BankKeeper            *bankkeeper.BaseKeeper
	CapabilityKeeper      *capabilitykeeper.Keeper
	CrisisKeeper          *crisiskeeper.Keeper
	UpgradeKeeper         *upgradekeeper.Keeper
	GroupKeeper           *groupkeeper.Keeper
	ConsensusParamsKeeper *consensusparamkeeper.Keeper
	IBCKeeper             *ibckeeper.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
	TransferKeeper        *ibctransferkeeper.Keeper
	NftTransferKeeper     *ibcnfttransferkeeper.Keeper
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
	OracleKeeper          *oraclekeeper.Keeper // x/oracle keeper used for the connect oracle
	MarketMapKeeper       *marketmapkeeper.Keeper
	IBCHooksKeeper        *ibchookskeeper.Keeper
	ForwardingKeeper      *forwardingkeeper.Keeper
	RatelimitKeeper       *ratelimitkeeper.Keeper

	// make scoped keepers public for test purposes
	ScopedIBCKeeper           capabilitykeeper.ScopedKeeper
	ScopedTransferKeeper      capabilitykeeper.ScopedKeeper
	ScopedNftTransferKeeper   capabilitykeeper.ScopedKeeper
	ScopedICAHostKeeper       capabilitykeeper.ScopedKeeper
	ScopedICAControllerKeeper capabilitykeeper.ScopedKeeper
	ScopedICAAuthKeeper       capabilitykeeper.ScopedKeeper
}

func NewAppKeeper(
	ac, vc, cc address.Codec,
	appCodec codec.Codec,
	txConfig client.TxConfig,
	bApp *baseapp.BaseApp,
	legacyAmino *codec.LegacyAmino,
	maccPerms map[string][]string,
	blockedAddress map[string]bool,
	skipUpgradeHeights map[int64]bool,
	homePath string,
	invCheckPeriod uint,
	logger log.Logger,
	evmConfig evmconfig.EVMConfig,
	appOpts servertypes.AppOptions,
) AppKeepers {
	appKeepers := AppKeepers{}

	// Set keys KVStoreKey, TransientStoreKey, MemoryStoreKey
	appKeepers.GenerateKeys()

	// register streaming services
	if err := bApp.RegisterStreamingServices(appOpts, appKeepers.keys); err != nil {
		logger.Error("failed to load state streaming", "err", err)
		os.Exit(1)
	}

	authorityAccAddr := authtypes.NewModuleAddress(opchildtypes.ModuleName)
	authorityAddr, err := ac.BytesToString(authorityAccAddr)
	if err != nil {
		logger.Error("failed to retrieve authority address", "err", err)
		os.Exit(1)
	}

	// set the BaseApp's parameter store
	consensusParamsKeeper := consensusparamkeeper.NewKeeper(appCodec, runtime.NewKVStoreService(appKeepers.keys[consensusparamtypes.StoreKey]), authorityAddr, runtime.EventService{})
	appKeepers.ConsensusParamsKeeper = &consensusParamsKeeper
	bApp.SetParamStore(appKeepers.ConsensusParamsKeeper.ParamsStore)

	// add capability keeper and ScopeToModule for ibc module
	appKeepers.CapabilityKeeper = capabilitykeeper.NewKeeper(appCodec, appKeepers.keys[capabilitytypes.StoreKey], appKeepers.memKeys[capabilitytypes.MemStoreKey])

	// grant capabilities for the ibc and ibc-transfer modules
	appKeepers.ScopedIBCKeeper = appKeepers.CapabilityKeeper.ScopeToModule(ibcexported.ModuleName)
	appKeepers.ScopedTransferKeeper = appKeepers.CapabilityKeeper.ScopeToModule(ibctransfertypes.ModuleName)
	appKeepers.ScopedNftTransferKeeper = appKeepers.CapabilityKeeper.ScopeToModule(ibcnfttransfertypes.ModuleName)
	appKeepers.ScopedICAHostKeeper = appKeepers.CapabilityKeeper.ScopeToModule(icahosttypes.SubModuleName)
	appKeepers.ScopedICAControllerKeeper = appKeepers.CapabilityKeeper.ScopeToModule(icacontrollertypes.SubModuleName)
	appKeepers.ScopedICAAuthKeeper = appKeepers.CapabilityKeeper.ScopeToModule(icaauthtypes.ModuleName)

	appKeepers.CapabilityKeeper.Seal()

	// add keepers
	appKeepers.EVMKeeper = new(evmkeeper.Keeper)
	erc20Keeper := new(evmkeeper.ERC20Keeper)
	erc721Keeper := new(evmkeeper.ERC721Keeper)

	accountKeeper := authkeeper.NewAccountKeeper(
		appCodec,
		runtime.NewKVStoreService(appKeepers.keys[authtypes.StoreKey]),
		authtypes.ProtoBaseAccount,
		maccPerms,
		ac,
		sdk.GetConfig().GetBech32AccountAddrPrefix(),
		authorityAddr,
	)
	appKeepers.AccountKeeper = &accountKeeper

	bankKeeper := bankkeeper.NewBaseKeeper(
		appCodec,
		runtime.NewKVStoreService(appKeepers.keys[banktypes.StoreKey]),
		appKeepers.AccountKeeper,
		erc20Keeper,
		blockedAddress,
		authorityAddr,
	)
	appKeepers.BankKeeper = &bankKeeper

	communityPoolKeeper := newCommunityPoolKeeper(appKeepers.BankKeeper, authtypes.FeeCollectorName)

	////////////////////////////////
	// OPChildKeeper Configuration //
	////////////////////////////////

	// initialize oracle keeper
	marketMapKeeper := marketmapkeeper.NewKeeper(
		runtime.NewKVStoreService(appKeepers.keys[marketmaptypes.StoreKey]),
		appCodec,
		authorityAccAddr,
	)
	appKeepers.MarketMapKeeper = marketMapKeeper

	oracleKeeper := oraclekeeper.NewKeeper(
		runtime.NewKVStoreService(appKeepers.keys[oracletypes.StoreKey]),
		appCodec,
		marketMapKeeper,
		authorityAccAddr,
	)
	appKeepers.OracleKeeper = &oracleKeeper

	// Add the oracle keeper as a hook to market map keeper so new market map entries can be created
	// and propagated to the oracle keeper.
	appKeepers.MarketMapKeeper.SetHooks(appKeepers.OracleKeeper.Hooks())

	appKeepers.OPChildKeeper = opchildkeeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(appKeepers.keys[opchildtypes.StoreKey]),
		appKeepers.AccountKeeper,
		appKeepers.BankKeeper,
		appKeepers.OracleKeeper,
		ante.CreateAnteHandlerForOPinit(appKeepers.AccountKeeper, appKeepers.EVMKeeper, txConfig.SignModeHandler()),
		txConfig.TxDecoder(),
		bApp.MsgServiceRouter(),
		authorityAddr,
		ac,
		vc,
		cc,
		logger,
	)

	appKeepers.CrisisKeeper = crisiskeeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(appKeepers.keys[crisistypes.StoreKey]),
		invCheckPeriod,
		appKeepers.BankKeeper,
		authtypes.FeeCollectorName,
		authorityAddr,
		ac,
	)

	appKeepers.UpgradeKeeper = upgradekeeper.NewKeeper(
		skipUpgradeHeights,
		runtime.NewKVStoreService(appKeepers.keys[upgradetypes.StoreKey]),
		appCodec,
		homePath,
		bApp,
		authorityAddr,
	)

	i := 0
	moduleAddrs := make([]sdk.AccAddress, len(maccPerms))
	for name := range maccPerms {
		moduleAddrs[i] = authtypes.NewModuleAddress(name)
		i += 1
	}

	feeGrantKeeper := feegrantkeeper.NewKeeper(appCodec, runtime.NewKVStoreService(appKeepers.keys[feegrant.StoreKey]), appKeepers.AccountKeeper)
	appKeepers.FeeGrantKeeper = &feeGrantKeeper

	authzKeeper := authzkeeper.NewKeeper(runtime.NewKVStoreService(appKeepers.keys[authzkeeper.StoreKey]), appCodec, bApp.MsgServiceRouter(), appKeepers.AccountKeeper)
	appKeepers.AuthzKeeper = &authzKeeper

	groupConfig := group.DefaultConfig()
	groupKeeper := groupkeeper.NewKeeper(
		appKeepers.keys[group.StoreKey],
		appCodec,
		bApp.MsgServiceRouter(),
		appKeepers.AccountKeeper,
		groupConfig,
	)
	appKeepers.GroupKeeper = &groupKeeper

	// Create IBC Keeper
	appKeepers.IBCKeeper = ibckeeper.NewKeeper(
		appCodec,
		appKeepers.keys[ibcexported.StoreKey],
		nil, // we don't need migration
		appKeepers.OPChildKeeper,
		appKeepers.UpgradeKeeper,
		appKeepers.ScopedIBCKeeper,
		authorityAddr,
	)

	// Set IBC post handler to receive validator set updates
	appKeepers.IBCKeeper.ClientKeeper.WithPostUpdateHandler(
		appKeepers.OPChildKeeper.UpdateHostValidatorSet,
	)

	ibcFeeKeeper := ibcfeekeeper.NewKeeper(
		appCodec,
		appKeepers.keys[ibcfeetypes.StoreKey],
		appKeepers.IBCKeeper.ChannelKeeper,
		appKeepers.IBCKeeper.ChannelKeeper,
		appKeepers.IBCKeeper.PortKeeper,
		appKeepers.AccountKeeper,
		appKeepers.BankKeeper,
	)
	appKeepers.IBCFeeKeeper = &ibcFeeKeeper

	appKeepers.IBCHooksKeeper = ibchookskeeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(appKeepers.keys[ibchookstypes.StoreKey]),
		authorityAddr,
		ac,
	)

	appKeepers.ForwardingKeeper = forwardingkeeper.NewKeeper(
		appCodec,
		logger,
		runtime.NewKVStoreService(appKeepers.keys[forwardingtypes.StoreKey]),
		runtime.NewTransientStoreService(appKeepers.tkeys[forwardingtypes.TransientStoreKey]),
		appheaderinfo.NewHeaderInfoService(),
		runtime.ProvideEventService(),
		authorityAddr,
		appKeepers.AccountKeeper,
		appKeepers.BankKeeper,
		appKeepers.IBCKeeper.ChannelKeeper,
		appKeepers.TransferKeeper,
	)
	appKeepers.BankKeeper.AppendSendRestriction(appKeepers.ForwardingKeeper.SendRestrictionFn)

	////////////////////////////
	// Transfer configuration //
	////////////////////////////
	// Send   : transfer -> packet forward -> rate limit -> fee        -> channel
	// Receive: channel  -> fee            -> evm        -> rate limit -> packet forward -> forwarding -> transfer

	var transferStack porttypes.IBCModule
	{
		packetForwardKeeper := &packetforwardkeeper.Keeper{}
		rateLimitKeeper := &ratelimitkeeper.Keeper{}

		// Create Transfer Keepers
		transferKeeper := ibctransferkeeper.NewKeeper(
			appCodec,
			appKeepers.keys[ibctransfertypes.StoreKey],
			nil, // we don't need migration
			// ics4wrapper: transfer -> packet forward
			packetForwardKeeper,
			appKeepers.IBCKeeper.ChannelKeeper,
			appKeepers.IBCKeeper.PortKeeper,
			appKeepers.AccountKeeper,
			appKeepers.BankKeeper,
			appKeepers.ScopedTransferKeeper,
			authorityAddr,
		)
		appKeepers.TransferKeeper = &transferKeeper
		transferStack = ibctransfer.NewIBCModule(*appKeepers.TransferKeeper)

		// forwarding middleware
		transferStack = forwarding.NewMiddleware(
			// receive: forwarding -> transfer
			transferStack,
			appKeepers.AccountKeeper,
			appKeepers.ForwardingKeeper,
		)

		// create packet forward middleware
		*packetForwardKeeper = *packetforwardkeeper.NewKeeper(
			appCodec,
			appKeepers.keys[packetforwardtypes.StoreKey],
			appKeepers.TransferKeeper,
			appKeepers.IBCKeeper.ChannelKeeper,
			communityPoolKeeper,
			appKeepers.BankKeeper,
			// ics4wrapper: transfer -> packet forward -> rate limit
			rateLimitKeeper,
			authorityAddr,
		)
		appKeepers.PacketForwardKeeper = packetForwardKeeper
		transferStack = packetforward.NewIBCMiddleware(
			// receive: packet forward -> forwarding -> transfer
			transferStack,
			appKeepers.PacketForwardKeeper,
			0,
			packetforwardkeeper.DefaultForwardTransferPacketTimeoutTimestamp,
			packetforwardkeeper.DefaultRefundTransferPacketTimeoutTimestamp,
		)

		// create the rate limit keeper
		*rateLimitKeeper = *ratelimitkeeper.NewKeeper(
			appCodec,
			runtime.NewKVStoreService(appKeepers.keys[ratelimittypes.StoreKey]),
			paramtypes.Subspace{}, // empty params
			authorityAddr,
			appKeepers.BankKeeper,
			appKeepers.IBCKeeper.ChannelKeeper,
			// ics4wrapper: transfer -> packet forward -> rate limit -> fee
			appKeepers.IBCFeeKeeper,
		)
		appKeepers.RatelimitKeeper = rateLimitKeeper

		// rate limit middleware
		transferStack = ratelimit.NewIBCMiddleware(
			*appKeepers.RatelimitKeeper,
			// receive: rate limit -> packet forward -> forwarding -> transfer
			transferStack,
		)

		// create evm middleware for transfer
		transferStack = ibchooks.NewIBCMiddleware(
			// receive: evm -> rate limit -> packet forward -> forwarding -> transfer
			transferStack,
			ibchooks.NewICS4Middleware(
				nil, /* ics4wrapper: not used */
				ibcevmhooks.NewEVMHooks(appCodec, ac, appKeepers.EVMKeeper),
			),
			appKeepers.IBCHooksKeeper,
		)

		// create ibcfee middleware for transfer
		transferStack = ibcfee.NewIBCMiddleware(
			// receive: fee -> evm -> rate limit -> packet forward -> forwarding -> transfer
			transferStack,
			// ics4wrapper: transfer -> packet forward -> rate limit -> fee -> channel
			*appKeepers.IBCFeeKeeper,
		)
	}

	////////////////////////////////
	// Nft Transfer configuration //
	////////////////////////////////

	var nftTransferStack porttypes.IBCModule
	{
		// Create Transfer Keepers
		appKeepers.NftTransferKeeper = ibcnfttransferkeeper.NewKeeper(
			appCodec,
			runtime.NewKVStoreService(appKeepers.keys[ibcnfttransfertypes.StoreKey]),
			// ics4wrapper: nft transfer -> fee -> channel
			appKeepers.IBCFeeKeeper,
			appKeepers.IBCKeeper.ChannelKeeper,
			appKeepers.IBCKeeper.PortKeeper,
			appKeepers.AccountKeeper,
			erc721Keeper,
			appKeepers.ScopedNftTransferKeeper,
			authorityAddr,
		)
		nftTransferIBCModule := ibcnfttransfer.NewIBCModule(*appKeepers.NftTransferKeeper)

		// create move middleware for nft-transfer
		hookMiddleware := ibchooks.NewIBCMiddleware(
			// receive: evm -> nft-transfer
			nftTransferIBCModule,
			ibchooks.NewICS4Middleware(
				nil, /* ics4wrapper: not used */
				ibcevmhooks.NewEVMHooks(appCodec, ac, appKeepers.EVMKeeper),
			),
			appKeepers.IBCHooksKeeper,
		)

		nftTransferStack = ibcfee.NewIBCMiddleware(
			// receive: channel -> fee -> evm -> nft transfer
			hookMiddleware,
			*appKeepers.IBCFeeKeeper,
		)
	}

	///////////////////////
	// ICA configuration //
	///////////////////////

	var icaHostStack porttypes.IBCModule
	var icaControllerStack porttypes.IBCModule
	{
		icaHostKeeper := icahostkeeper.NewKeeper(
			appCodec, appKeepers.keys[icahosttypes.StoreKey],
			nil, // we don't need migration
			appKeepers.IBCFeeKeeper,
			appKeepers.IBCKeeper.ChannelKeeper,
			appKeepers.IBCKeeper.PortKeeper,
			appKeepers.AccountKeeper,
			appKeepers.ScopedICAHostKeeper,
			bApp.MsgServiceRouter(),
			authorityAddr,
		)
		icaHostKeeper.WithQueryRouter(bApp.GRPCQueryRouter())
		appKeepers.ICAHostKeeper = &icaHostKeeper

		icaControllerKeeper := icacontrollerkeeper.NewKeeper(
			appCodec, appKeepers.keys[icacontrollertypes.StoreKey],
			nil, // we don't need migration
			appKeepers.IBCFeeKeeper,
			appKeepers.IBCKeeper.ChannelKeeper,
			appKeepers.IBCKeeper.PortKeeper,
			appKeepers.ScopedICAControllerKeeper,
			bApp.MsgServiceRouter(),
			authorityAddr,
		)
		appKeepers.ICAControllerKeeper = &icaControllerKeeper

		icaAuthKeeper := icaauthkeeper.NewKeeper(
			appCodec,
			*appKeepers.ICAControllerKeeper,
			appKeepers.ScopedICAAuthKeeper,
			ac,
		)
		appKeepers.ICAAuthKeeper = &icaAuthKeeper

		icaAuthIBCModule := icaauth.NewIBCModule(*appKeepers.ICAAuthKeeper)
		icaHostIBCModule := icahost.NewIBCModule(*appKeepers.ICAHostKeeper)
		icaHostStack = ibcfee.NewIBCMiddleware(icaHostIBCModule, *appKeepers.IBCFeeKeeper)
		icaControllerIBCModule := icacontroller.NewIBCMiddleware(icaAuthIBCModule, *appKeepers.ICAControllerKeeper)
		icaControllerStack = ibcfee.NewIBCMiddleware(icaControllerIBCModule, *appKeepers.IBCFeeKeeper)
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
		AddRoute(ibcnfttransfertypes.ModuleName, nftTransferStack)

	appKeepers.IBCKeeper.SetRouter(ibcRouter)

	//////////////////////////////
	// EVMKeeper Configuration //
	//////////////////////////////

	gasPriceKeeper := newGasPriceKeeper(appKeepers.OPChildKeeper)
	*appKeepers.EVMKeeper = *evmkeeper.NewKeeper(
		ac,
		appCodec,
		runtime.NewKVStoreService(appKeepers.keys[evmtypes.StoreKey]),
		runtime.NewTransientStoreService(appKeepers.tkeys[evmtypes.TStoreKey]),
		accountKeeper,
		bankKeeper,
		communityPoolKeeper,
		appKeepers.IBCHooksKeeper,
		gasPriceKeeper,
		bApp.MsgServiceRouter(),
		bApp.GRPCQueryRouter(),
		authorityAddr,
		evmConfig,
		evmtypes.DefaultQueryCosmosWhitelist(),
	)
	*erc20Keeper = *appKeepers.EVMKeeper.ERC20Keeper().(*evmkeeper.ERC20Keeper)
	*erc721Keeper = *appKeepers.EVMKeeper.ERC721Keeper().(*evmkeeper.ERC721Keeper)

	// x/auction module keeper initialization

	// initialize the keeper
	auctionKeeper := auctionkeeper.NewKeeperWithRewardsAddressProvider(
		appCodec,
		appKeepers.keys[auctiontypes.StoreKey],
		appKeepers.AccountKeeper,
		appKeepers.BankKeeper,
		opchildlanes.NewRewardsAddressProvider(authtypes.FeeCollectorName),
		authorityAddr,
	)
	appKeepers.AuctionKeeper = &auctionKeeper

	return appKeepers
}
