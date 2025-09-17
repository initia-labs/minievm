package test

import (
	"encoding/hex"
	"encoding/json"
	"math/big"
	"time"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	transfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"

	ibctesting "github.com/initia-labs/initia/x/ibc/testing"
	"github.com/initia-labs/minievm/x/evm/types"
	evmtypes "github.com/initia-labs/minievm/x/evm/types"

	"github.com/ethereum/go-ethereum/common"
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

func (suite *KeeperTestSuite) TestE2ELocalTokenWrapper() {
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

	// mint tokenA 10 ether in A chain
	amount, _ = new(big.Int).SetString("10000000000000000000", 10)
	userA = pathA2B.EndpointA.Chain.SenderAccount.GetAddress()
	userB = pathA2B.EndpointB.Chain.SenderAccount.GetAddress()
	tokenA = suite.createAndMintERC20(pathA2B.EndpointA, userA, amount, 18)
	initialBalancesA = bankKeeperA.GetAllBalances(pathA2B.EndpointA.Chain.GetContext(), userA)
	initialBalancesB = bankKeeperB.GetAllBalances(pathA2B.EndpointB.Chain.GetContext(), userB)

	var tokenARemoteDenom string // denom of wrapped tokenA in source chain
	var tokenB sdk.Coin
	var wrapperAddr common.Address

	// test timeout
	tokenARemoteDenom, tokenB, wrapperAddr = suite.wrapLocal(pathA2B, tokenA, userA, userB, amount, true)

	// timeout should revert transfer
	suite.Require().Equal(initialBalancesB, bankKeeperB.GetAllBalances(pathA2B.EndpointB.Chain.GetContext(), userB))
	suite.Require().Equal(initialBalancesA, bankKeeperA.GetAllBalances(pathA2B.EndpointA.Chain.GetContext(), userA))

	// Wrap tokenA and transfer token from A chain to B chain
	tokenARemoteDenom, tokenB, wrapperAddr = suite.wrapLocal(pathA2B, tokenA, userA, userB, amount, false)

	// check local decimals => remote decimals properly converted
	suite.Require().Equal(tokenB.Amount, math.NewIntFromBigInt(amount).QuoRaw(1e12))

	// Transfer tokenB from B chain to A chain, unwrap tokenB
	suite.unwrapLocal(pathA2B, tokenARemoteDenom, tokenB, wrapperAddr, userB, userA)

	// Have the same balance as the initial state
	suite.Require().Equal(initialBalancesB, bankKeeperB.GetAllBalances(pathA2B.EndpointB.Chain.GetContext(), userB))
	suite.Require().Equal(initialBalancesA, bankKeeperA.GetAllBalances(pathA2B.EndpointA.Chain.GetContext(), userA))
}

func (suite *KeeperTestSuite) TestE2ERemoteTokenWrapper() {
	suite.SetupTest()
	pathA2B := NewTransferPath(suite.chainA, suite.chainB)
	suite.coordinator.SetupConnections(pathA2B)
	suite.coordinator.CreateTransferChannels(pathA2B)

	bankKeeperA := getMinitiaApp(suite.chainA).BankKeeper
	bankKeeperB := getMinitiaApp(suite.chainB).BankKeeper

	var tokenAContractAddr common.Address
	var amount *big.Int
	var initialBalancesA sdk.Coins
	var initialBalancesB sdk.Coins
	var userA sdk.AccAddress
	var userB sdk.AccAddress

	// Mint tokenA 10 ether in B chain
	amount, _ = new(big.Int).SetString("10000000", 10)
	userA = pathA2B.EndpointA.Chain.SenderAccount.GetAddress()
	userB = pathA2B.EndpointB.Chain.SenderAccount.GetAddress()
	tokenAContractAddr = suite.createAndMintERC20(pathA2B.EndpointA, userA, amount, 6)
	initialBalancesA = bankKeeperA.GetAllBalances(pathA2B.EndpointA.Chain.GetContext(), userA)
	initialBalancesB = bankKeeperB.GetAllBalances(pathA2B.EndpointB.Chain.GetContext(), userB)

	// transfer token A from A chain to B chain and wrap it with hook
	var tokenA, tokenB sdk.Coin
	var tokenBContractAddr common.Address
	tokenA, tokenB, tokenBContractAddr = suite.wrapRemote(pathA2B, tokenAContractAddr, userA, userB, amount)

	// Have the expected balance
	expectedBalanceA := initialBalancesA.Sub(tokenA)
	expectedBalanceB := initialBalancesB.Add(tokenB)
	suite.Require().Equal(expectedBalanceA, bankKeeperA.GetAllBalances(pathA2B.EndpointA.Chain.GetContext(), userA))
	suite.Require().Equal(expectedBalanceB, bankKeeperB.GetAllBalances(pathA2B.EndpointB.Chain.GetContext(), userB))

	// test timeout
	tokenA = suite.unwrapRemote(pathA2B, tokenBContractAddr, userB, userA, tokenB.Amount.BigInt(), true)

	// Have the expected balance
	suite.Require().Equal(expectedBalanceA, bankKeeperA.GetAllBalances(pathA2B.EndpointA.Chain.GetContext(), userA))
	suite.Require().Equal(expectedBalanceB, bankKeeperB.GetAllBalances(pathA2B.EndpointB.Chain.GetContext(), userB))

	// test unwrap
	// Unwrap tokenB and transfer token from B chain to A chain
	tokenA = suite.unwrapRemote(pathA2B, tokenBContractAddr, userB, userA, tokenB.Amount.BigInt(), false)

	// Have the same balance with initial balance
	suite.Require().Equal(initialBalancesA, bankKeeperA.GetAllBalances(pathA2B.EndpointA.Chain.GetContext(), userA))
	suite.Require().Equal(initialBalancesB, bankKeeperB.GetAllBalances(pathA2B.EndpointB.Chain.GetContext(), userB))
}

func (suite *KeeperTestSuite) createAndMintERC20(endpoint *ibctesting.Endpoint, to sdk.AccAddress, amount *big.Int, decimals uint8) common.Address {
	ctx := endpoint.Chain.GetContext()
	toAddr := common.BytesToAddress(to)
	evmKeeper := getMinitiaApp(endpoint.Chain).EVMKeeper
	erc20Keeper := evmKeeper.ERC20Keeper()

	ethFactoryAddr, err := evmKeeper.GetERC20FactoryAddr(ctx)
	suite.Require().NoError(err)

	abi := erc20Keeper.GetERC20FactoryABI()

	// Create
	inputBz, err := abi.Pack("createERC20", "foo", "foo", decimals)
	suite.Require().NoError(err)

	result, _, err := evmKeeper.EVMCall(ctx, toAddr, ethFactoryAddr, inputBz, nil, nil, nil)
	suite.Require().NoError(err)
	tokenAddr := common.BytesToAddress(result)

	// Mint
	abi = evmKeeper.ERC20Keeper().GetERC20ABI()
	suite.Require().NoError(err)

	inputBz, err = abi.Pack("mint", toAddr, amount)
	suite.Require().NoError(err)

	_, _, err = evmKeeper.EVMCall(ctx, toAddr, tokenAddr, inputBz, nil, nil, nil)
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
	timeout bool,
) (string, sdk.Coin, common.Address) {
	fromEndpoint := path.EndpointA
	toEndpoint := path.EndpointB
	fromCtx := fromEndpoint.Chain.GetContext()
	toCtx := toEndpoint.Chain.GetContext()

	timeoutTime := fromEndpoint.Chain.CurrentHeader.GetTime().Add(time.Minute * 1)

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
	_, _, err = evmKeeper.EVMCall(fromCtx, senderAddr, tokenAddress, inputBz, nil, nil, nil)
	suite.Require().NoError(err)

	// wrap 18dp token to 6dp token
	denom, err := evmtypes.ContractAddrToDenom(fromCtx, evmKeeper, tokenAddress)
	suite.Require().NoError(err)
	// get wrapped token contract address
	inputBz, err = erc20Keeper.GetERC20WrapperABI().Pack("getToRemoteERC20Address", denom)
	suite.Require().NoError(err)
	ret, err := evmKeeper.EVMStaticCall(fromCtx, senderAddr, wrapperAddr, inputBz, nil)
	suite.Require().NoError(err)
	expectedAddr := common.BytesToAddress(ret[12:])
	inputBz, err = erc20Keeper.GetERC20WrapperABI().Pack("toRemoteAndIBCTransfer0", denom, amount, fromEndpoint.ChannelID, receiver.String(), big.NewInt(timeoutTime.UnixNano()))
	suite.Require().NoError(err)

	senderStr, err := suite.chainA.Codec.InterfaceRegistry().SigningContext().AddressCodec().BytesToString(sender)
	suite.Require().NoError(err)
	msgWrap := &evmtypes.MsgCall{
		Sender:       senderStr,
		ContractAddr: wrapperAddr.Hex(),
		Input:        "0x" + common.Bytes2Hex(inputBz),
	}
	res, err := fromEndpoint.Chain.SendMsgs(msgWrap)
	suite.Require().NoError(err)

	events := res.GetEvents()
	for _, event := range events {
		if event.Type == evmtypes.EventTypeERC20Created {
			suite.Require().Equal(event.Attributes[1].Value, expectedAddr.Hex())
		}
	}
	packet, err := ibctesting.ParsePacketFromEvents(events)
	suite.Require().NoError(err)

	var data transfertypes.FungibleTokenPacketData
	err = suite.chainA.Codec.UnmarshalJSON(packet.GetData(), &data)
	suite.Require().NoError(err)

	if timeout {
		// raise timeout by setting the current time to the timeout time
		suite.coordinator.IncrementTimeBy(time.Minute * 10)
		suite.coordinator.UpdateTime()

		err = toEndpoint.UpdateClient()
		suite.Require().NoError(err)

		err = fromEndpoint.UpdateClient()
		suite.Require().NoError(err)

		// trigger timeout
		err = fromEndpoint.TimeoutPacket(packet)
		suite.Require().NoError(err)
	} else {
		err = toEndpoint.UpdateClient()
		suite.Require().NoError(err)

		err = toEndpoint.RecvPacket(packet)
		suite.Require().NoError(err)
	}

	receivedToken := transfertypes.GetTransferCoin(toEndpoint.ChannelConfig.PortID, toEndpoint.ChannelID, data.Denom, math.ZeroInt())
	receivedToken = bankKeeper.GetBalance(toEndpoint.Chain.GetContext(), receiver, receivedToken.Denom)
	return data.Denom, receivedToken, wrapperAddr
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
	inputBz, err := erc20Keeper.GetERC20WrapperABI().Pack("toLocal0", common.BytesToAddress(receiver), tokenADenom, uint8(6))
	suite.Require().NoError(err)
	hook, err := unwrapHook(HookData{
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

	err = toEndpoint.UpdateClient()
	suite.Require().NoError(err)

	err = toEndpoint.RecvPacket(packet)
	suite.Require().NoError(err)
}

// wrap the remote tokens and transfer token from A to B
func (suite *KeeperTestSuite) wrapRemote(
	path *ibctesting.Path,
	tokenAddress common.Address,
	sender sdk.AccAddress,
	receiver sdk.AccAddress,
	amount *big.Int,
) (sdk.Coin, sdk.Coin, common.Address) {
	fromEndpoint := path.EndpointA
	toEndpoint := path.EndpointB
	fromCtx := fromEndpoint.Chain.GetContext()
	toCtx := toEndpoint.Chain.GetContext()
	fromEvmKeeper := getMinitiaApp(fromEndpoint.Chain).EVMKeeper
	fromErc20Keeper := fromEvmKeeper.ERC20Keeper()
	toEvmKeeper := getMinitiaApp(toEndpoint.Chain).EVMKeeper
	toErc20Keeper := toEvmKeeper.ERC20Keeper()

	bankKeeper := getMinitiaApp(toEndpoint.Chain).BankKeeper
	coins := bankKeeper.GetAllBalances(toCtx, receiver)
	suite.Require().Equal(1, coins.Len())

	wrapperAddr, err := fromEvmKeeper.GetERC20WrapperAddr(fromCtx)
	suite.Require().NoError(err)

	receiverAddr := common.BytesToAddress(receiver)
	denom, err := types.ContractAddrToDenom(fromCtx, fromEvmKeeper, tokenAddress)
	suite.Require().NoError(err)
	sendToken := sdk.NewCoin(denom, math.NewIntFromBigInt(amount))

	// get wrapped token contract address
	inputBz, err := toErc20Keeper.GetERC20WrapperABI().Pack("getToLocalERC20Address", denom, denom, denom, uint8(6))
	suite.Require().NoError(err)
	ret, err := toEvmKeeper.EVMStaticCall(toCtx, common.HexToAddress("0x1"), wrapperAddr, inputBz, nil)
	suite.Require().NoError(err)
	expectedAddr := common.BytesToAddress(ret[12:])

	// create wrap hook message
	receivedToken := transfertypes.GetTransferCoin(toEndpoint.ChannelConfig.PortID, toEndpoint.ChannelID, denom, math.NewIntFromBigInt(amount))
	inputBz, err = fromErc20Keeper.GetERC20WrapperABI().Pack("toLocal", receiverAddr, receivedToken.Denom, amount, uint8(6))
	suite.Require().NoError(err)

	hook, err := unwrapHook(HookData{
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
		sendToken,
		sender.String(),
		wrapperAddr.Hex(),
		toEndpoint.Chain.GetTimeoutHeight(),
		uint64(toEndpoint.Chain.CurrentHeader.Time.UnixNano()+1000000000000000000),
		hook,
	)

	res, err := fromEndpoint.Chain.SendMsgs(msgTransfer)
	suite.Require().NoError(err)

	events := res.GetEvents()
	for _, event := range events {
		if event.Type == evmtypes.EventTypeERC20Created {
			suite.Require().Equal(event.Attributes[1].Value, expectedAddr.Hex())
		}
	}
	packet, err := ibctesting.ParsePacketFromEvents(events)
	suite.Require().NoError(err)

	var data transfertypes.FungibleTokenPacketData
	err = suite.chainA.Codec.UnmarshalJSON(packet.GetData(), &data)
	suite.Require().NoError(err)

	err = commitRecvPacket(fromEndpoint, toEndpoint, packet)
	suite.Require().NoError(err)

	coins = bankKeeper.GetAllBalances(toCtx, receiver)
	suite.Require().Equal(2, coins.Len())
	suite.Require().Equal(math.NewIntFromBigInt(amount).MulRaw(1e12), coins[0].Amount)

	// compute received token contract address
	contractAddr, err := types.DenomToContractAddr(toCtx, toEvmKeeper, coins[0].Denom)
	suite.Require().NoError(err)
	return sendToken, coins[0], contractAddr
}

// unwrap the remote tokens and transfer token from B to A
func (suite *KeeperTestSuite) unwrapRemote(
	path *ibctesting.Path,
	tokenAddress common.Address,
	sender sdk.AccAddress,
	receiver sdk.AccAddress,
	amount *big.Int,
	timeout bool,
) sdk.Coin {
	fromEndpoint := path.EndpointB
	toEndpoint := path.EndpointA
	fromCtx := fromEndpoint.Chain.GetContext()
	toCtx := toEndpoint.Chain.GetContext()

	timeoutTime := fromEndpoint.Chain.CurrentHeader.GetTime().Add(time.Minute * 1)

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
	_, _, err = evmKeeper.EVMCall(fromCtx, senderAddr, tokenAddress, inputBz, nil, nil, nil)
	suite.Require().NoError(err)

	// unwrap 18dp token to 6dp token
	denom, err := types.ContractAddrToDenom(fromCtx, evmKeeper, tokenAddress)
	suite.Require().NoError(err)
	inputBz, err = erc20Keeper.GetERC20WrapperABI().Pack("toRemoteAndIBCTransfer0", denom, amount, fromEndpoint.ChannelID, receiver.String(), big.NewInt(timeoutTime.UnixNano()))
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

	if timeout {
		// raise timeout by setting the current time to the timeout time
		suite.coordinator.IncrementTimeBy(time.Minute * 10)
		suite.coordinator.UpdateTime()

		err = toEndpoint.UpdateClient()
		suite.Require().NoError(err)

		err = fromEndpoint.UpdateClient()
		suite.Require().NoError(err)

		// trigger timeout
		err = fromEndpoint.TimeoutPacket(packet)
		suite.Require().NoError(err)
	} else {
		err = toEndpoint.UpdateClient()
		suite.Require().NoError(err)

		err = toEndpoint.RecvPacket(packet)
		suite.Require().NoError(err)
	}

	receivedToken := transfertypes.GetTransferCoin(toEndpoint.ChannelConfig.PortID, toEndpoint.ChannelID, data.Denom, math.ZeroInt())
	receivedToken = bankKeeper.GetBalance(toEndpoint.Chain.GetContext(), receiver, receivedToken.Denom)
	return receivedToken
}

type HookData struct {
	EVM struct {
		Message struct {
			ContractAddr string   `json:"contract_addr"`
			Input        string   `json:"input"`
			Value        string   `json:"value"`
			AccessList   []string `json:"access_list"`
		} `json:"message"`
	} `json:"evm"`
}

func unwrapHook(data HookData) (string, error) {
	bz, err := json.Marshal(data)

	return string(bz), err
}
