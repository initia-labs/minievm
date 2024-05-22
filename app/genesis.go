package app

import (
	"encoding/hex"
	"encoding/json"

	"cosmossdk.io/core/address"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	icacontrollertypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/controller/types"
	icagenesistypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/genesis/types"
	icahosttypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/host/types"
	icatypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/types"
	ibcexported "github.com/cosmos/ibc-go/v8/modules/core/exported"
	ibctypes "github.com/cosmos/ibc-go/v8/modules/core/types"

	l2slinky "github.com/initia-labs/OPinit/x/opchild/l2slinky"
	opchildtypes "github.com/initia-labs/OPinit/x/opchild/types"
	"github.com/initia-labs/initia/app/genesis_markets"

	auctiontypes "github.com/skip-mev/block-sdk/v2/x/auction/types"
	slinkytypes "github.com/skip-mev/slinky/pkg/types"
	marketmaptypes "github.com/skip-mev/slinky/x/marketmap/types"
	oracletypes "github.com/skip-mev/slinky/x/oracle/types"
)

// GenesisState - The genesis state of the blockchain is represented here as a map of raw json
// messages key'd by a identifier string.
// The identifier is used to determine which module genesis information belongs
// to so it may be appropriately routed during init chain.
// Within this application default genesis information is retrieved from
// the ModuleBasicManager which populates json from each BasicModule
// object provided to it during init.
type GenesisState map[string]json.RawMessage

// NewDefaultGenesisState generates the default state for the application.
func NewDefaultGenesisState(cdc codec.Codec, mbm module.BasicManager, denom string) GenesisState {
	return GenesisState(mbm.DefaultGenesis(cdc)).
		ConfigureMinGasPrices(cdc).
		ConfigureICA(cdc).
		ConfigureIBCAllowedClients(cdc).
		ConfigureAuctionFee(cdc, denom).
		AddMarketData(cdc, cdc.InterfaceRegistry().SigningContext().AddressCodec())
}

func (genState GenesisState) AddMarketData(cdc codec.JSONCodec, ac address.Codec) GenesisState {
	var oracleGenState oracletypes.GenesisState
	cdc.MustUnmarshalJSON(genState[oracletypes.ModuleName], &oracleGenState)

	var marketGenState marketmaptypes.GenesisState
	cdc.MustUnmarshalJSON(genState[marketmaptypes.ModuleName], &marketGenState)

	// Load initial markets
	markets, err := genesis_markets.ReadMarketsFromFile(genesis_markets.GenesisMarkets)
	if err != nil {
		panic(err)
	}
	marketGenState.MarketMap = genesis_markets.ToMarketMap(markets)

	// Skip Admin account.
	adminAddrBz, err := hex.DecodeString("51B89E89D58FFB3F9DB66263FF10A216CF388A0E")
	if err != nil {
		panic(err)
	}

	adminAddr, err := ac.BytesToString(adminAddrBz)
	if err != nil {
		panic(err)
	}

	marketGenState.Params.MarketAuthorities = []string{adminAddr}
	marketGenState.Params.Admin = adminAddr

	var id uint64

	// Initialize all markets plus ReservedCPTimestamp
	currencyPairGenesis := make([]oracletypes.CurrencyPairGenesis, len(markets)+1)
	cp, err := slinkytypes.CurrencyPairFromString(l2slinky.ReservedCPTimestamp)
	if err != nil {
		panic(err)
	}
	currencyPairGenesis[id] = oracletypes.CurrencyPairGenesis{
		CurrencyPair:      cp,
		CurrencyPairPrice: nil,
		Nonce:             0,
		Id:                id,
	}
	id++
	for i, market := range markets {
		currencyPairGenesis[i+1] = oracletypes.CurrencyPairGenesis{
			CurrencyPair:      market.Ticker.CurrencyPair,
			CurrencyPairPrice: nil,
			Nonce:             0,
			Id:                id,
		}
		id++
	}

	oracleGenState.CurrencyPairGenesis = currencyPairGenesis
	oracleGenState.NextId = id

	// write the updates to genState
	genState[marketmaptypes.ModuleName] = cdc.MustMarshalJSON(&marketGenState)
	genState[oracletypes.ModuleName] = cdc.MustMarshalJSON(&oracleGenState)
	return genState
}

func (genState GenesisState) ConfigureAuctionFee(cdc codec.JSONCodec, denom string) GenesisState {
	var auctionGenState auctiontypes.GenesisState
	cdc.MustUnmarshalJSON(genState[auctiontypes.ModuleName], &auctionGenState)
	auctionGenState.Params.ReserveFee.Denom = denom
	auctionGenState.Params.MinBidIncrement.Denom = denom
	genState[auctiontypes.ModuleName] = cdc.MustMarshalJSON(&auctionGenState)

	return genState
}

// ConfigureMinGasPrices generates the default state for the application.
func (genState GenesisState) ConfigureMinGasPrices(cdc codec.JSONCodec) GenesisState {
	var opChildGenState opchildtypes.GenesisState
	cdc.MustUnmarshalJSON(genState[opchildtypes.ModuleName], &opChildGenState)
	opChildGenState.Params.MinGasPrices = nil
	genState[opchildtypes.ModuleName] = cdc.MustMarshalJSON(&opChildGenState)

	return genState
}

func (genState GenesisState) ConfigureICA(cdc codec.JSONCodec) GenesisState {
	// create ICS27 Controller submodule params
	controllerParams := icacontrollertypes.Params{
		ControllerEnabled: true,
	}

	// create ICS27 Host submodule params
	hostParams := icahosttypes.Params{
		HostEnabled: true,
		AllowMessages: []string{
			authzMsgExec,
			authzMsgGrant,
			authzMsgRevoke,
			bankMsgSend,
			bankMsgMultiSend,
			feegrantMsgGrantAllowance,
			feegrantMsgRevokeAllowance,
			groupCreateGroup,
			groupCreateGroupPolicy,
			groupExec,
			groupLeaveGroup,
			groupSubmitProposal,
			groupUpdateGroupAdmin,
			groupUpdateGroupMember,
			groupUpdateGroupPolicyAdmin,
			groupUpdateGroupPolicyDecisionPolicy,
			groupVote,
			groupWithdrawProposal,
			transferMsgTransfer,
			nftTransferMsgTransfer,
			sftTransferMsgTransfer,
			moveMsgPublishModuleBundle,
			moveMsgExecuteEntryFunction,
			moveMsgExecuteScript,
		},
	}

	var icaGenState icagenesistypes.GenesisState
	cdc.MustUnmarshalJSON(genState[icatypes.ModuleName], &icaGenState)
	icaGenState.ControllerGenesisState.Params = controllerParams
	icaGenState.HostGenesisState.Params = hostParams
	genState[icatypes.ModuleName] = cdc.MustMarshalJSON(&icaGenState)

	return genState
}

func (genState GenesisState) ConfigureIBCAllowedClients(cdc codec.JSONCodec) GenesisState {
	var ibcGenesis ibctypes.GenesisState
	cdc.MustUnmarshalJSON(genState[ibcexported.ModuleName], &ibcGenesis)

	allowedClients := ibcGenesis.ClientGenesis.Params.AllowedClients
	for i, client := range allowedClients {
		if client == ibcexported.Localhost {
			allowedClients = append(allowedClients[:i], allowedClients[i+1:]...)
			break
		}
	}

	ibcGenesis.ClientGenesis.Params.AllowedClients = allowedClients
	genState[ibcexported.ModuleName] = cdc.MustMarshalJSON(&ibcGenesis)

	return genState
}
