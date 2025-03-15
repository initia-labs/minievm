package test

import (
	"encoding/hex"
	"encoding/json"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	"github.com/ethereum/go-ethereum/common"
	ibctesting "github.com/initia-labs/initia/x/ibc/testing"
	"github.com/initia-labs/minievm/x/evm/types"
)

func NewTransferPath(chainA, chainB *ibctesting.TestChain) *ibctesting.Path {
	path := ibctesting.NewPath(chainA, chainB)
	path.EndpointA.ChannelConfig.Version = transfertypes.Version
	path.EndpointB.ChannelConfig.Version = transfertypes.Version
	path.EndpointA.ClientConfig.(*ibctesting.TendermintConfig).TrustingPeriod = TrustingPeriod
	path.EndpointA.ClientConfig.(*ibctesting.TendermintConfig).UnbondingPeriod = UnbondingPeriod
	path.EndpointB.ClientConfig.(*ibctesting.TendermintConfig).TrustingPeriod = TrustingPeriod
	path.EndpointB.ClientConfig.(*ibctesting.TendermintConfig).UnbondingPeriod = UnbondingPeriod

	return path
}

func (suite *KeeperTestSuite) TestE2ETokenWrapper() {
	suite.SetupTest()
	pathA2B := NewTransferPath(suite.chainA, suite.chainB)
	suite.coordinator.SetupConnections(pathA2B)
	suite.coordinator.CreateTransferChannels(pathA2B)

	bankKeeperA := getMinitiaApp(suite.chainA).BankKeeper
	bankKeeperB := getMinitiaApp(suite.chainB).BankKeeper

	var tokenA common.Address
	var amount *big.Int
	var initialBalancesA sdk.Coins
	var initialBalancesB sdk.Coins
	var userA sdk.AccAddress
	var userB sdk.AccAddress
	suite.Run("Mint tokenA 10 ether in A chain", func() {
		amount, _ = new(big.Int).SetString("10000000000000000000", 10)
		userA = pathA2B.EndpointA.Chain.SenderAccount.GetAddress()
		userB = pathA2B.EndpointB.Chain.SenderAccount.GetAddress()
		tokenA = suite.createAndMintERC20(pathA2B.EndpointA, userA, amount)
		initialBalancesA = bankKeeperA.GetAllBalances(pathA2B.EndpointA.Chain.GetContext(), userA)
		initialBalancesB = bankKeeperB.GetAllBalances(pathA2B.EndpointB.Chain.GetContext(), userB)
	})

	var tokenARemoteDenom string // denom of wrapped tokenA in source chain
	var tokenB sdk.Coin
	var wrapperAddr common.Address
	suite.Run("Wrap tokenA and transfer token from A chain to B chain", func() {
		tokenARemoteDenom, tokenB, wrapperAddr = suite.wrapLocal(pathA2B, tokenA, userA, userB, amount, big.NewInt(suite.chainB.CurrentHeader.Time.UnixNano()+1000000000000000000))
	})

	suite.Run("Transfer tokenB from B chain to A chain, unwrap tokenB", func() {
		suite.unwrapLocal(pathA2B, tokenARemoteDenom, tokenB, wrapperAddr, userB, userA)
	})

	suite.Run("Have the same balance as the initial state", func() {
		suite.Require().Equal(initialBalancesB, bankKeeperB.GetAllBalances(pathA2B.EndpointB.Chain.GetContext(), userB))
		suite.Require().Equal(initialBalancesA, bankKeeperA.GetAllBalances(pathA2B.EndpointA.Chain.GetContext(), userA))
	})
}

func (suite *KeeperTestSuite) createAndMintERC20(endpoint *ibctesting.Endpoint, to sdk.AccAddress, amount *big.Int) common.Address {
	ctx := endpoint.Chain.GetContext()
	toAddr := common.BytesToAddress(to)
	evmKeeper := getMinitiaApp(endpoint.Chain).EVMKeeper
	erc20Keeper := evmKeeper.ERC20Keeper()

	ethFactoryAddr, err := evmKeeper.GetERC20FactoryAddr(ctx)
	suite.Require().NoError(err)

	abi := erc20Keeper.GetERC20FactoryABI()

	// Create
	inputBz, err := abi.Pack("createERC20", "foo", "foo", uint8(18))
	suite.Require().NoError(err)

	result, _, err := evmKeeper.EVMCall(ctx, toAddr, ethFactoryAddr, inputBz, nil, nil)
	suite.Require().NoError(err)
	tokenAddr := common.BytesToAddress(result)

	// Mint
	abi = evmKeeper.ERC20Keeper().GetERC20ABI()
	suite.Require().NoError(err)

	inputBz, err = abi.Pack("mint", toAddr, amount)
	suite.Require().NoError(err)

	_, _, err = evmKeeper.EVMCall(ctx, toAddr, tokenAddr, inputBz, nil, nil)
	suite.Require().NoError(err)

	return tokenAddr
}

// wrap the tokens and transfer token from A to B
func (suite *KeeperTestSuite) wrapLocal(
	path *ibctesting.Path,
	tokenAddress common.Address,
	sender sdk.AccAddress,
	receiver sdk.AccAddress,
	amount *big.Int,
	timeout *big.Int,
) (string, sdk.Coin, common.Address) {
	fromEndpoint := path.EndpointA
	toEndpoint := path.EndpointB
	fromCtx := fromEndpoint.Chain.GetContext()
	toCtx := toEndpoint.Chain.GetContext()

	senderAddr := common.BytesToAddress(sender)
	bankKeeper := getMinitiaApp(toEndpoint.Chain).BankKeeper
	coins := bankKeeper.GetAllBalances(toCtx, receiver)
	suite.Require().Equal(1, coins.Len())

	evmKeeper := getMinitiaApp(fromEndpoint.Chain).EVMKeeper
	erc20Keeper := evmKeeper.ERC20Keeper()

	wrapperAddr, err := evmKeeper.GetERC20WrapperAddr(fromCtx)
	suite.Require().NoError(err)
	// approve
	inputBz, err := erc20Keeper.GetERC20ABI().Pack("approve", wrapperAddr, amount)
	suite.Require().NoError(err)
	_, _, err = evmKeeper.EVMCall(fromCtx, senderAddr, tokenAddress, inputBz, nil, nil)
	suite.Require().NoError(err)
	// wrap
	inputBz, err = erc20Keeper.GetERC20WrapperABI().Pack("wrapLocal0", fromEndpoint.ChannelID, tokenAddress, receiver.String(), amount, timeout)
	suite.Require().NoError(err)

	senderStr, err := suite.chainA.Codec.InterfaceRegistry().SigningContext().AddressCodec().BytesToString(sender)
	suite.Require().NoError(err)
	msgWrap := &types.MsgCall{
		Sender:       senderStr,
		ContractAddr: wrapperAddr.Hex(),
		Input:        "0x" + common.Bytes2Hex(inputBz),
	}
	res, err := fromEndpoint.Chain.SendMsgs(msgWrap)
	suite.Require().NoError(err)
	packet, err := ibctesting.ParsePacketFromEvents(res.GetEvents())
	suite.Require().NoError(err)

	var data transfertypes.FungibleTokenPacketData
	err = suite.chainA.Codec.UnmarshalJSON(packet.GetData(), &data)
	suite.Require().NoError(err)

	err = commitRecvPacket(fromEndpoint, toEndpoint, packet)
	suite.Require().NoError(err)

	// check balance of receiver after wrap and send tokens
	coins = bankKeeper.GetAllBalances(toCtx, receiver)
	suite.Require().Equal(2, coins.Len())
	return data.Denom, coins[0], wrapperAddr
}

// Transfer token from B to A and unwrap the local tokens
func (suite *KeeperTestSuite) unwrapLocal(
	path *ibctesting.Path,
	tokenADenom string,
	tokenB sdk.Coin,
	wrapperAddr common.Address,
	sender sdk.AccAddress,
	receiver sdk.AccAddress,
) {
	fromEndpoint := path.EndpointB
	toEndpoint := path.EndpointA

	evmKeeper := getMinitiaApp(fromEndpoint.Chain).EVMKeeper
	erc20Keeper := evmKeeper.ERC20Keeper()

	// set hook message
	inputBz, err := erc20Keeper.GetERC20WrapperABI().Pack("unwrapLocal", common.BytesToAddress(receiver), tokenADenom)
	suite.Require().NoError(err)
	hook, err := unwrapHook(UnwrapHookData{
		EVM: struct {
			Message struct {
				ContractAddr string   `json:"contract_addr"`
				Input        string   `json:"input"`
				Value        string   `json:"value"`
				AccessList   []string `json:"access_list"`
			} `json:"message"`
		}{
			Message: struct {
				ContractAddr string   `json:"contract_addr"`
				Input        string   `json:"input"`
				Value        string   `json:"value"`
				AccessList   []string `json:"access_list"`
			}{
				ContractAddr: wrapperAddr.String(),
				Input:        "0x" + hex.EncodeToString(inputBz),
				Value:        "0",
				AccessList:   nil,
			},
		},
	})
	suite.Require().NoError(err)

	msgTransfer := transfertypes.NewMsgTransfer(
		fromEndpoint.ChannelConfig.PortID,
		fromEndpoint.ChannelID,
		tokenB,
		sender.String(),
		wrapperAddr.Hex(),
		toEndpoint.Chain.GetTimeoutHeight(),
		uint64(toEndpoint.Chain.CurrentHeader.Time.UnixNano()+1000000000000000000),
		hook,
	)

	res, err := fromEndpoint.Chain.SendMsgs(msgTransfer)
	suite.Require().NoError(err)

	packet, err := ibctesting.ParsePacketFromEvents(res.GetEvents())
	suite.Require().NoError(err)

	var data transfertypes.FungibleTokenPacketData
	err = suite.chainA.Codec.UnmarshalJSON(packet.GetData(), &data)
	suite.Require().NoError(err)

	err = commitRecvPacket(fromEndpoint, toEndpoint, packet)
	suite.Require().NoError(err)
}

type UnwrapHookData struct {
	EVM struct {
		Message struct {
			ContractAddr string   `json:"contract_addr"`
			Input        string   `json:"input"`
			Value        string   `json:"value"`
			AccessList   []string `json:"access_list"`
		} `json:"message"`
	} `json:"evm"`
}

func unwrapHook(data UnwrapHookData) (string, error) {
	bz, err := json.Marshal(data)

	return string(bz), err
}
