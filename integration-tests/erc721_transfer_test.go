package test

import (
	"math/big"

	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"

	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"

	nfttransfertypes "github.com/initia-labs/initia/x/ibc/nft-transfer/types"
	ibctesting "github.com/initia-labs/initia/x/ibc/testing"

	evmkeeper "github.com/initia-labs/minievm/x/evm/keeper"
	evmtypes "github.com/initia-labs/minievm/x/evm/types"

	"github.com/ethereum/go-ethereum/common"
)

func NewNftTransferPath(chainA, chainB *ibctesting.TestChain) *ibctesting.Path {
	path := ibctesting.NewPath(chainA, chainB)
	path.EndpointA.ChannelConfig.PortID = nfttransfertypes.PortID
	path.EndpointB.ChannelConfig.PortID = nfttransfertypes.PortID
	path.EndpointA.ChannelConfig.Version = nfttransfertypes.Version
	path.EndpointB.ChannelConfig.Version = nfttransfertypes.Version
	path.EndpointA.ClientConfig.(*ibctesting.TendermintConfig).TrustingPeriod = TrustingPeriod
	path.EndpointA.ClientConfig.(*ibctesting.TendermintConfig).UnbondingPeriod = UnbondingPeriod
	path.EndpointB.ClientConfig.(*ibctesting.TendermintConfig).TrustingPeriod = TrustingPeriod
	path.EndpointB.ClientConfig.(*ibctesting.TendermintConfig).UnbondingPeriod = UnbondingPeriod
	return path
}

func (suite *KeeperTestSuite) CreateNftClass(
	endpoint *ibctesting.Endpoint,
	name, uri string,
) string {
	evmKeeper := getMinitiaApp(endpoint.Chain).EVMKeeper
	nftKeeper := evmKeeper.ERC721Keeper().(*evmkeeper.ERC721Keeper)

	ctx := endpoint.Chain.GetContext()

	createAccount := endpoint.Chain.SenderAccounts[5].SenderAccount.GetAddress()
	createAccountAddr := common.BytesToAddress(createAccount)

	inputBz, err := nftKeeper.ABI.Pack("", name, name)
	suite.Require().NoError(err)

	_, contractAddr, _, err := evmKeeper.EVMCreate(ctx, createAccountAddr, append(nftKeeper.ERC721Bin, inputBz...), nil, nil)
	suite.Require().NoError(err)

	classId, err := evmtypes.ClassIdFromCollectionAddress(endpoint.Chain.GetContext(), nftKeeper, contractAddr)
	suite.Require().NoError(err)
	return classId
}

func (suite *KeeperTestSuite) MintNft(
	endpoint *ibctesting.Endpoint,
	receiver sdk.AccAddress,
	classId, className, tokenUri string, tokenId math.Int,
) {
	evmKeeper := getMinitiaApp(endpoint.Chain).EVMKeeper
	nftKeeper := evmKeeper.ERC721Keeper().(*evmkeeper.ERC721Keeper)

	ctx := endpoint.Chain.GetContext()

	createAccount := endpoint.Chain.SenderAccounts[5].SenderAccount.GetAddress()
	createAccountAddr := common.BytesToAddress(createAccount)
	receiverAddr := common.BytesToAddress(receiver)

	bigTokenId, ok := new(big.Int).SetString(tokenId.String(), 10)
	suite.Require().True(ok)

	inputBz, err := nftKeeper.ABI.Pack("mint", receiverAddr, bigTokenId, tokenUri, "")
	suite.Require().NoError(err)

	contractAddr, err := evmtypes.ContractAddressFromClassId(ctx, nftKeeper, classId)
	suite.Require().NoError(err)

	_, _, err = nftKeeper.EVMCall(ctx, createAccountAddr, contractAddr, inputBz, nil, nil)
	suite.Require().NoError(err)
}

func (suite *KeeperTestSuite) ConfirmClassId(endpoint *ibctesting.Endpoint, classId, targetClassId string) {
	if classId == targetClassId {
		return
	}
	ctx := endpoint.Chain.GetContext()
	classIdPath, err := getMinitiaApp(endpoint.Chain).NftTransferKeeper.ClassIdPathFromHash(ctx, targetClassId)
	suite.Require().NoError(err, "ClassIdPathFromHash error on chain %s", endpoint.Chain.ChainID)

	baseClassId := nfttransfertypes.ParseClassTrace(classIdPath).BaseClassId
	suite.Equal(classId, baseClassId, "wrong classId on chain %s", endpoint.Chain.ChainID)
}

func (suite *KeeperTestSuite) ConfirmNFTOwner(endpoint *ibctesting.Endpoint, classId, tokenId string, receiver sdk.Address) {
	evmKeeper := getMinitiaApp(endpoint.Chain).EVMKeeper
	nftKeeper := evmKeeper.ERC721Keeper().(*evmkeeper.ERC721Keeper)

	ctx := endpoint.Chain.GetContext()

	receiverAddr := common.BytesToAddress(receiver.Bytes())
	owner, err := nftKeeper.OwnerOf(ctx, tokenId, classId)
	suite.Require().NoError(err)
	suite.Require().Equal(receiverAddr, owner, "wrong owner on chain %s", endpoint.Chain.ChainID)
}

// The following test describes the entire cross-chain process of nft-transfer.
// The execution sequence of the cross-chain process is:
// A -> B -> C -> B ->A
func (suite *KeeperTestSuite) TestSendAndReceive() {
	suite.SetupTest()

	var classId string
	classUri := "uri"
	className := "name"
	nftId := math.NewInt(128379128731)
	nftIdStr := nftId.String()
	nftUri := "kitty_uri"

	var targetClassId string
	var packet channeltypes.Packet

	// WARNING : be careful not to be confused with endpoint names
	// pathB2C.EndpointA is ChainB endpoint (source of path)`
	// pathB2C.EndpointB is ChainC endpoint (destination of path)
	// pathA2B.EndpointB.Chain.SenderAccount is same with receiver account of pathA2B before testing`
	pathA2B := NewNftTransferPath(suite.chainA, suite.chainB)
	suite.Run("transfer forward A->B", func() {
		{
			suite.coordinator.SetupConnections(pathA2B)
			suite.coordinator.CreateChannels(pathA2B)

			sender := pathA2B.EndpointA.Chain.SenderAccount.GetAddress()
			receiver := pathA2B.EndpointB.Chain.SenderAccount.GetAddress()

			classId = suite.CreateNftClass(pathA2B.EndpointA, className, classUri)
			suite.MintNft(pathA2B.EndpointA, sender, classId, className, nftUri, nftId)

			packet = suite.transferNft(
				pathA2B.EndpointA,
				pathA2B.EndpointB,
				classId,
				nftIdStr,
				sender.String(),
				receiver.String(),
			)

			targetClassId = suite.receiverNft(
				pathA2B.EndpointA,
				pathA2B.EndpointB,
				packet,
			)

			suite.ConfirmClassId(pathA2B.EndpointB, classId, targetClassId)
			suite.ConfirmNFTOwner(pathA2B.EndpointB, targetClassId, nftIdStr, receiver)
		}
	})

	// transfer from chainB to chainC
	pathB2C := NewNftTransferPath(suite.chainB, suite.chainC)
	suite.Run("transfer forward B->C", func() {
		{
			suite.coordinator.SetupConnections(pathB2C)
			suite.coordinator.CreateChannels(pathB2C)

			sender := pathA2B.EndpointB.Chain.SenderAccount.GetAddress()
			receiver := pathB2C.EndpointB.Chain.SenderAccount.GetAddress()

			packet = suite.transferNft(
				pathB2C.EndpointA,
				pathB2C.EndpointB,
				targetClassId,
				nftIdStr,
				sender.String(),
				receiver.String(),
			)

			targetClassId = suite.receiverNft(
				pathB2C.EndpointA,
				pathB2C.EndpointB,
				packet,
			)

			suite.ConfirmClassId(pathB2C.EndpointB, classId, targetClassId)
			suite.ConfirmNFTOwner(pathB2C.EndpointB, targetClassId, nftIdStr, receiver)
		}
	})

	// transfer from chainC to chainB
	suite.Run("transfer back C->B", func() {
		{
			sender := pathB2C.EndpointB.Chain.SenderAccount.GetAddress()
			receiver := pathB2C.EndpointA.Chain.SenderAccount.GetAddress()

			packet = suite.transferNft(
				pathB2C.EndpointB,
				pathB2C.EndpointA,
				targetClassId,
				nftIdStr,
				sender.String(),
				receiver.String(),
			)

			targetClassId = suite.receiverNft(
				pathB2C.EndpointB,
				pathB2C.EndpointA,
				packet,
			)

			suite.ConfirmClassId(pathB2C.EndpointA, classId, targetClassId)
			suite.ConfirmNFTOwner(pathB2C.EndpointA, targetClassId, nftIdStr, receiver)
		}
	})

	// transfer from chainB to chainA
	suite.Run("transfer back B->A", func() {
		{
			sender := pathA2B.EndpointB.Chain.SenderAccount.GetAddress()
			receiver := pathA2B.EndpointA.Chain.SenderAccount.GetAddress()

			packet = suite.transferNft(
				pathA2B.EndpointB,
				pathA2B.EndpointA,
				targetClassId,
				nftIdStr,
				sender.String(),
				receiver.String(),
			)

			targetClassId = suite.receiverNft(
				pathA2B.EndpointB,
				pathA2B.EndpointA,
				packet,
			)

			suite.ConfirmClassId(pathA2B.EndpointA, classId, targetClassId)
			suite.ConfirmNFTOwner(pathA2B.EndpointA, targetClassId, nftIdStr, receiver)
		}
	})
}

func (suite *KeeperTestSuite) transferNft(
	fromEndpoint, toEndpoint *ibctesting.Endpoint,
	classId, nftId string,
	sender, receiver string,
) channeltypes.Packet {
	msgTransfer := nfttransfertypes.NewMsgTransfer(
		fromEndpoint.ChannelConfig.PortID,
		fromEndpoint.ChannelID,
		classId,
		[]string{nftId},
		sender,
		receiver,
		toEndpoint.Chain.GetTimeoutHeight(),
		0,
		"",
	)

	res, err := fromEndpoint.Chain.SendMsgs(msgTransfer)
	suite.Require().NoError(err)

	packet, err := ibctesting.ParsePacketFromEvents(res.GetEvents())
	suite.Require().NoError(err)

	var data nfttransfertypes.NonFungibleTokenPacketData
	err = suite.chainA.Codec.UnmarshalJSON(packet.GetData(), &data)
	suite.Require().NoError(err)

	return packet
}

func (suite *KeeperTestSuite) receiverNft(
	fromEndpoint, toEndpoint *ibctesting.Endpoint,
	packet channeltypes.Packet,
) string {
	var data nfttransfertypes.NonFungibleTokenPacketData
	err := suite.chainA.Codec.UnmarshalJSON(packet.GetData(), &data)
	suite.Require().NoError(err)

	// get proof of packet commitment from chainA
	err = commitRecvPacket(fromEndpoint, toEndpoint, packet)
	suite.Require().NoError(err)

	var classId string

	isAwayFromOrigin := nfttransfertypes.SenderChainIsSource(packet.GetSourcePort(), packet.GetSourceChannel(), data.GetClassId())
	if isAwayFromOrigin {
		prefixedClassId := nfttransfertypes.GetClassIdPrefix(toEndpoint.ChannelConfig.PortID, toEndpoint.ChannelID) + data.GetClassId()
		trace := nfttransfertypes.ParseClassTrace(prefixedClassId)
		classId = trace.IBCClassId()
	} else {
		unprefixedClassId, err := nfttransfertypes.RemoveClassPrefix(packet.GetSourcePort(), packet.GetSourceChannel(), data.GetClassId())
		suite.Require().NoError(err)

		classId = unprefixedClassId
		classTrace := nfttransfertypes.ParseClassTrace(unprefixedClassId)
		if classTrace.Path != "" {
			classId = classTrace.IBCClassId()
		} else {
			_, data.ClassData, err = nfttransfertypes.ConvertClassDataFromICS721(data.ClassData)
			suite.Require().NoError(err, "ConvertTokenDataFromICS721 error on chain %s", toEndpoint.Chain.ChainID)
		}
	}
	evmKeeper := getMinitiaApp(toEndpoint.Chain).EVMKeeper
	toNftKeeper := evmKeeper.ERC721Keeper().(*evmkeeper.ERC721Keeper)

	ctx := toEndpoint.Chain.GetContext()

	_, classUri, _, err := toNftKeeper.GetClassInfo(ctx, classId)
	suite.Require().NoError(err, "not found class")
	suite.Require().Equal(classUri, data.GetClassUri(), "class uri not equal")
	return classId
}
