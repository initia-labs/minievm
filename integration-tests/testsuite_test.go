package test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"

	"cosmossdk.io/math"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v8/modules/core/24-host"

	nfttransferkeeper "github.com/initia-labs/initia/x/ibc/nft-transfer/keeper"
	nfttransfertypes "github.com/initia-labs/initia/x/ibc/nft-transfer/types"
	ibctesting "github.com/initia-labs/initia/x/ibc/testing"

	minievmapp "github.com/initia-labs/minievm/app"
)

type KeeperTestSuite struct {
	suite.Suite

	coordinator *ibctesting.Coordinator

	// testing chains used for convenience and readability
	// MinievmAppChain is the chain used by the testing suite
	chainA *ibctesting.TestChain
	// MinievmAppChain is the chain used by the testing suite
	chainB *ibctesting.TestChain
	// InitiaAppChain is the chain used by the testing suite
	chainC *ibctesting.TestChain

	queryClient nfttransfertypes.QueryClient
}

func getMinitiaApp(chain *ibctesting.TestChain) *minievmapp.MinitiaApp {
	return chain.App.(*minievmapp.MinitiaApp)
}

const (
	TrustingPeriod  time.Duration = time.Hour * 24 * 7 * 2 / 3
	UnbondingPeriod time.Duration = time.Hour * 24 * 7
)

func (suite *KeeperTestSuite) convertAppToMApp(chain *ibctesting.TestChain) {
	genAccs := make([]authtypes.GenesisAccount, len(chain.SenderAccounts))
	genBals := make([]banktypes.Balance, len(chain.SenderAccounts))
	for i, acc := range chain.SenderAccounts {
		genAccs[i] = acc.SenderAccount.(*authtypes.BaseAccount)
		amount, ok := math.NewIntFromString("10000000000000000000")
		suite.Require().True(ok)

		// add sender account
		balance := banktypes.Balance{
			Address: genAccs[i].GetAddress().String(),
			Coins:   sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, amount)),
		}
		genBals[i] = balance
	}

	miniApp := minievmapp.SetupWithGenesisAccounts(chain.Vals.Copy(), genAccs, genBals...)
	baseapp.SetChainID(chain.ChainID)(miniApp.GetBaseApp())
	chain.App = miniApp
	chain.Codec = miniApp.AppCodec()
	chain.TxConfig = miniApp.TxConfig()

	chain.CurrentHeader = cmtproto.Header{
		ChainID:            chain.ChainID,
		Height:             chain.App.LastBlockHeight() + 1,
		AppHash:            chain.App.LastCommitID().Hash,
		Time:               chain.CurrentHeader.Time,
		ValidatorsHash:     chain.Vals.Hash(),
		NextValidatorsHash: chain.NextVals.Hash(),
		ProposerAddress:    chain.CurrentHeader.ProposerAddress,
	}
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.coordinator = ibctesting.NewCoordinator(suite.T(), 3)
	suite.chainA = suite.coordinator.GetChain(ibctesting.GetChainID(1))
	suite.chainB = suite.coordinator.GetChain(ibctesting.GetChainID(2))
	suite.chainC = suite.coordinator.GetChain(ibctesting.GetChainID(3))

	suite.convertAppToMApp(suite.chainA)
	suite.convertAppToMApp(suite.chainB)
	suite.convertAppToMApp(suite.chainC)

	queryHelper := baseapp.NewQueryServerTestHelper(suite.chainA.GetContext(), getMinitiaApp(suite.chainA).InterfaceRegistry())
	nfttransfertypes.RegisterQueryServer(queryHelper, nfttransferkeeper.NewQueryServerImpl(getMinitiaApp(suite.chainA).NftTransferKeeper))
	suite.queryClient = nfttransfertypes.NewQueryClient(queryHelper)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func commitRecvPacket(fromEndpoint *ibctesting.Endpoint, toEndpoint *ibctesting.Endpoint, packet channeltypes.Packet) error {
	err := toEndpoint.UpdateClient()
	if err != nil {
		return err
	}

	packetKey := host.PacketCommitmentKey(packet.GetSourcePort(), packet.GetSourceChannel(), packet.GetSequence())
	proof, proofHeight := fromEndpoint.QueryProof(packetKey)

	recvMsg := channeltypes.NewMsgRecvPacket(
		packet, proof, proofHeight, toEndpoint.Chain.SenderAccount.GetAddress().String())
	_, err = toEndpoint.Chain.SendMsgs(recvMsg)
	if err != nil {
		return err
	}

	return nil
}
