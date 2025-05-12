package keeper_test

import (
	"crypto/rand"
	"testing"

	"cosmossdk.io/math"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/initia-labs/minievm/x/evm/contracts/erc20"
	"github.com/initia-labs/minievm/x/evm/contracts/erc20_factory"
	"github.com/initia-labs/minievm/x/evm/keeper"
	"github.com/initia-labs/minievm/x/evm/types"
)

func createERC20(s *ERC20TestSuite, caller common.Address, symbol string) common.Address {
	abi, err := erc20_factory.Erc20FactoryMetaData.GetAbi()
	s.Require().NoError(err)

	inputBz, err := abi.Pack("createERC20", symbol, symbol, uint8(6))
	s.Require().NoError(err)

	factoryAddr, err := s.input.EVMKeeper.GetERC20FactoryAddr(s.ctx)
	s.Require().NoError(err)

	ret, _, err := s.input.EVMKeeper.EVMCall(s.ctx, caller, factoryAddr, inputBz, nil, nil)
	s.Require().NoError(err)

	return common.BytesToAddress(ret[12:])
}

func create2ERC20(s *ERC20TestSuite, caller common.Address, symbol string) common.Address {
	salt := func() [32]byte {
		var salt [32]byte
		rand.Read(salt[:])
		return salt
	}()
	abi, err := erc20_factory.Erc20FactoryMetaData.GetAbi()
	s.Require().NoError(err)

	inputBz, err := abi.Pack("createERC200", symbol, symbol, uint8(6), salt)
	s.Require().NoError(err)

	factoryAddr, err := s.input.EVMKeeper.GetERC20FactoryAddr(s.ctx)
	s.Require().NoError(err)
	ret, _, err := s.input.EVMKeeper.EVMCall(s.ctx, caller, factoryAddr, inputBz, nil, nil)
	s.Require().NoError(err)

	inputBz, err = abi.Pack("computeERC20Address", symbol, symbol, uint8(6), salt)
	s.Require().NoError(err)
	ret2, _, err := s.input.EVMKeeper.EVMCall(s.ctx, caller, factoryAddr, inputBz, nil, nil)

	s.Require().Equal(ret2[12:], ret[12:])
	s.Require().NoError(err)

	return common.BytesToAddress(ret[12:])
}

func (s *ERC20TestSuite) burnERC20(caller, from common.Address, amount sdk.Coin, expectErr bool) {
	erc20ContractAddr, err := types.DenomToContractAddr(s.ctx, &s.input.EVMKeeper, amount.Denom)
	s.Require().NoError(err)

	abi, err := erc20.Erc20MetaData.GetAbi()
	s.Require().NoError(err)

	inputBz, err := abi.Pack("approve", caller, amount.Amount.BigInt())
	s.Require().NoError(err)

	_, _, err = s.input.EVMKeeper.EVMCall(s.ctx, from, erc20ContractAddr, inputBz, nil, nil)
	s.Require().NoError(err)

	inputBz, err = abi.Pack("burnFrom", from, amount.Amount.BigInt())
	s.Require().NoError(err)

	_, _, err = s.input.EVMKeeper.EVMCall(s.ctx, caller, erc20ContractAddr, inputBz, nil, nil)
	if expectErr {
		s.Require().Error(err)
	} else {
		s.Require().NoError(err)
	}
}

func (s *ERC20TestSuite) mintERC20(caller, recipient common.Address, amount sdk.Coin, expectErr bool) {
	abi, err := erc20.Erc20MetaData.GetAbi()
	s.Require().NoError(err)

	inputBz, err := abi.Pack("mint", recipient, amount.Amount.BigInt())
	s.Require().NoError(err)

	erc20ContractAddr, err := types.DenomToContractAddr(s.ctx, &s.input.EVMKeeper, amount.Denom)
	s.Require().NoError(err)

	_, _, err = s.input.EVMKeeper.EVMCall(s.ctx, caller, erc20ContractAddr, inputBz, nil, nil)
	if expectErr {
		s.Require().Error(err)
	} else {
		s.Require().NoError(err)
	}
}

func (s *ERC20TestSuite) transferERC20(caller common.Address, recipient common.Address, amount sdk.Coin, expectErr bool) {
	abi, err := erc20.Erc20MetaData.GetAbi()
	s.Require().NoError(err)

	inputBz, err := abi.Pack("transfer", recipient, amount.Amount.BigInt())
	s.Require().NoError(err)

	erc20ContractAddr, err := types.DenomToContractAddr(s.ctx, &s.input.EVMKeeper, amount.Denom)
	s.Require().NoError(err)

	_, _, err = s.input.EVMKeeper.EVMCall(s.ctx, caller, erc20ContractAddr, inputBz, nil, nil)
	if expectErr {
		s.Require().Error(err)
	} else {
		s.Require().NoError(err)
	}

}

func (s *ERC20TestSuite) approveERC20(caller, spender common.Address, amount sdk.Coin, expectErr bool) {
	abi, err := erc20.Erc20MetaData.GetAbi()
	s.Require().NoError(err)

	inputBz, err := abi.Pack("approve", spender, amount.Amount.BigInt())
	s.Require().NoError(err)

	erc20ContractAddr, err := types.DenomToContractAddr(s.ctx, &s.input.EVMKeeper, amount.Denom)
	s.Require().NoError(err)

	_, _, err = s.input.EVMKeeper.EVMCall(s.ctx, caller, erc20ContractAddr, inputBz, nil, nil)
	if expectErr {
		s.Require().Error(err)
	} else {
		s.Require().NoError(err)
	}
}

func (s *ERC20TestSuite) transferFromERC20(caller, from, to common.Address, amount sdk.Coin, expectErr bool) {
	abi, err := erc20.Erc20MetaData.GetAbi()
	s.Require().NoError(err)

	inputBz, err := abi.Pack("transferFrom", from, to, amount.Amount.BigInt())
	s.Require().NoError(err)

	erc20ContractAddr, err := types.DenomToContractAddr(s.ctx, &s.input.EVMKeeper, amount.Denom)
	s.Require().NoError(err)

	_, _, err = s.input.EVMKeeper.EVMCall(s.ctx, caller, erc20ContractAddr, inputBz, nil, nil)
	if expectErr {
		s.Require().Error(err)
	} else {
		s.Require().NoError(err)
	}
}

func (s *ERC20TestSuite) updateMetadataERC20(caller common.Address, denom, name, symbol string, decimals uint8) error {
	abi, err := erc20.Erc20MetaData.GetAbi()
	s.Require().NoError(err)

	inputBz, err := abi.Pack("updateMetadata", name, symbol, decimals)
	s.Require().NoError(err)

	erc20ContractAddr, err := types.DenomToContractAddr(s.ctx, &s.input.EVMKeeper, denom)
	s.Require().NoError(err)

	_, _, err = s.input.EVMKeeper.EVMCall(s.ctx, caller, erc20ContractAddr, inputBz, nil, nil)
	return err
}

type ERC20Creator struct {
	name string
	fn   func(s *ERC20TestSuite, caller common.Address, symbol string) common.Address
}

type ERC20TestSuite struct {
	suite.Suite

	ctx         sdk.Context
	input       TestKeepers
	accAddrs    []sdk.AccAddress
	evmAddrs    []common.Address
	erc20Keeper *keeper.ERC20Keeper
	creator     ERC20Creator
}

func NewERC20TestSuite(creator ERC20Creator) *ERC20TestSuite {
	return &ERC20TestSuite{
		creator:  creator,
		evmAddrs: []common.Address{},
		accAddrs: []sdk.AccAddress{},
	}
}

func TestERC20AllCreators(t *testing.T) {
	creators := []ERC20Creator{
		{"CREATE", createERC20},
		{"CREATE2", create2ERC20},
	}

	for _, creator := range creators {
		suite.Run(t, NewERC20TestSuite(creator))
	}
}

func (s *ERC20TestSuite) SetupTest() {
	s.ctx, s.input = createDefaultTestInput(s.T())
	for i := 0; i < 2; i++ {
		_, _, addr := keyPubAddr()
		s.accAddrs = append(s.accAddrs, addr)
		s.evmAddrs = append(s.evmAddrs, common.BytesToAddress(addr.Bytes()))
	}

	var err error
	erc20KeeperInterface, err := keeper.NewERC20Keeper(&s.input.EVMKeeper)
	s.Require().NoError(err)
	s.erc20Keeper = erc20KeeperInterface.(*keeper.ERC20Keeper)
	s.Require().NoError(err)
}

func (s *ERC20TestSuite) createContract(symbol string, deployer common.Address) (common.Address, string) {
	addr := s.creator.fn(s, deployer, symbol)
	denom, err := types.ContractAddrToDenom(s.ctx, &s.input.EVMKeeper, addr)
	s.Require().NoError(err)
	s.Require().Equal("evm/"+addr.Hex()[2:], denom)
	return addr, denom
}

func (s *ERC20TestSuite) Test_BalanceOf() {
	addr, addr1 := s.accAddrs[0], s.accAddrs[1]
	s.input.Faucet.Fund(s.ctx, addr, sdk.NewCoin("foo", math.NewInt(100)))

	amount, err := s.input.EVMKeeper.ERC20Keeper().GetBalance(s.ctx, addr, "foo")
	s.Require().NoError(err)
	s.Require().Equal(math.NewInt(100), amount)

	amount, err = s.input.EVMKeeper.ERC20Keeper().GetBalance(s.ctx, addr1, "foo")
	s.Require().NoError(err)
	s.Require().Equal(math.NewInt(0), amount)
}

func (s *ERC20TestSuite) Test_SendCoins() {
	erc20Keeper, err := keeper.NewERC20Keeper(&s.input.EVMKeeper)
	addr, addr2 := s.accAddrs[0], s.accAddrs[1]
	s.Require().NoError(err)

	err = erc20Keeper.MintCoins(s.ctx, addr, sdk.NewCoins(
		sdk.NewCoin("bar", math.NewInt(200)),
		sdk.NewCoin("foo", math.NewInt(100)),
	))
	s.Require().NoError(err)

	err = erc20Keeper.SendCoins(s.ctx, addr, addr2, sdk.NewCoins(
		sdk.NewCoin("bar", math.NewInt(100)),
		sdk.NewCoin("foo", math.NewInt(50)),
	))
	s.Require().NoError(err)

	res, _, err := erc20Keeper.GetPaginatedBalances(s.ctx, nil, addr)
	s.Require().NoError(err)
	s.Require().Equal(sdk.NewCoins(
		sdk.NewCoin("bar", math.NewInt(100)),
		sdk.NewCoin("foo", math.NewInt(50)),
	), res)

	res2, _, err := erc20Keeper.GetPaginatedBalances(s.ctx, nil, addr2)
	s.Require().NoError(err)
	s.Require().Equal(sdk.NewCoins(
		sdk.NewCoin("bar", math.NewInt(100)),
		sdk.NewCoin("foo", math.NewInt(50)),
	), res2)
}

func (s *ERC20TestSuite) Test_TransferToModuleAccount() {
	s.input.Faucet.Fund(s.ctx, s.accAddrs[0], sdk.NewCoin("foo", math.NewInt(100)))

	feeCollectorAddr := authtypes.NewModuleAddress(authtypes.FeeCollectorName)
	s.transferERC20(s.evmAddrs[0], common.BytesToAddress(feeCollectorAddr.Bytes()), sdk.NewCoin("foo", math.NewInt(50)), true)

	_, _, addr2 := keyPubAddr()
	evmAddr2 := common.BytesToAddress(addr2.Bytes())
	s.transferERC20(s.evmAddrs[0], evmAddr2, sdk.NewCoin("foo", math.NewInt(50)), false)
}

func (s *ERC20TestSuite) TestMintBurn() {
	evmAddr, addr := s.evmAddrs[0], s.accAddrs[0]
	_, denom := s.createContract("foo", evmAddr)

	s.mintERC20(evmAddr, evmAddr, sdk.NewCoin(denom, math.NewInt(100)), false)

	amount, err := s.erc20Keeper.GetBalance(s.ctx, addr, denom)
	s.Require().NoError(err)
	s.Require().Equal(math.NewInt(100), amount)

	err = s.erc20Keeper.BurnCoins(s.ctx, addr, sdk.NewCoins(
		sdk.NewCoin(denom, math.NewInt(50)),
	))
	s.Require().NoError(err)
	s.Require().Equal(math.NewInt(50), s.input.CommunityPoolKeeper.CommunityPool.AmountOf(denom))
}

func (s *ERC20TestSuite) Test_MintToModuleAccount() {
	evmAddr := s.evmAddrs[0]
	_, fooDenom := s.createContract("foo", evmAddr)

	// deploy erc20 contract
	feeCollectorAddr := authtypes.NewModuleAddress(authtypes.FeeCollectorName)
	s.mintERC20(evmAddr, common.BytesToAddress(feeCollectorAddr.Bytes()), sdk.NewCoin(fooDenom, math.NewInt(50)), true)

	_, _, addr2 := keyPubAddr()
	evmAddr2 := common.BytesToAddress(addr2.Bytes())
	s.mintERC20(evmAddr, evmAddr2, sdk.NewCoin(fooDenom, math.NewInt(50)), false)
}

func (s *ERC20TestSuite) Test_BurnFromModuleAccount() {
	evmAddr, evmAddr2 := s.evmAddrs[0], s.evmAddrs[1]
	addr, addr2 := s.accAddrs[0], s.accAddrs[1]
	_, fooDenom := s.createContract("foo", evmAddr)
	// register fee collector module account
	s.input.AccountKeeper.GetModuleAccount(s.ctx, authtypes.FeeCollectorName)

	erc20Keeper, err := keeper.NewERC20Keeper(&s.input.EVMKeeper)
	s.Require().NoError(err)
	// mint coins
	feeCollectorAddr := authtypes.NewModuleAddress(authtypes.FeeCollectorName)
	s.mintERC20(evmAddr, evmAddr, sdk.NewCoin(fooDenom, math.NewInt(100)), false)
	erc20Keeper.SendCoins(s.ctx, addr, feeCollectorAddr, sdk.NewCoins(sdk.NewCoin(fooDenom, math.NewInt(50))))
	erc20Keeper.SendCoins(s.ctx, addr, addr2, sdk.NewCoins(sdk.NewCoin(fooDenom, math.NewInt(50))))

	// should not be able to burn from module account
	s.burnERC20(evmAddr, common.BytesToAddress(feeCollectorAddr.Bytes()), sdk.NewCoin(fooDenom, math.NewInt(50)), true)

	// should be able to burn from other account
	s.burnERC20(evmAddr, evmAddr2, sdk.NewCoin(fooDenom, math.NewInt(50)), false)
}

func (s *ERC20TestSuite) Test_MintBurn() {
	evmAddr, addr := s.evmAddrs[0], s.accAddrs[0]
	_, fooDenom := s.createContract("foo", evmAddr)

	// cannot mint erc20 from cosmos side
	cacheCtx, _ := s.ctx.CacheContext()
	err := s.erc20Keeper.MintCoins(cacheCtx, addr, sdk.NewCoins(
		sdk.NewCoin("bar", math.NewInt(200)),
		sdk.NewCoin(fooDenom, math.NewInt(100)),
	))
	s.Require().Error(err)

	// mint success
	err = s.erc20Keeper.MintCoins(s.ctx, addr, sdk.NewCoins(
		sdk.NewCoin("bar", math.NewInt(200)),
	))
	s.Require().NoError(err)

	// mint erc20
	s.mintERC20(evmAddr, evmAddr, sdk.NewCoin(fooDenom, math.NewInt(100)), false)

	amount, err := s.erc20Keeper.GetBalance(s.ctx, addr, "bar")
	s.Require().NoError(err)
	s.Require().Equal(math.NewInt(200), amount)

	amount, err = s.erc20Keeper.GetBalance(s.ctx, addr, fooDenom)
	s.Require().NoError(err)
	s.Require().Equal(math.NewInt(100), amount)

	// erc20(foo) coins will be sent to community pool
	err = s.erc20Keeper.BurnCoins(s.ctx, addr, sdk.NewCoins(
		sdk.NewCoin("bar", math.NewInt(50)),
		sdk.NewCoin(fooDenom, math.NewInt(50)),
	))
	s.Require().NoError(err)

	res, _, err := s.erc20Keeper.GetPaginatedBalances(s.ctx, nil, addr)
	s.Require().NoError(err)
	s.Require().Equal(sdk.NewCoins(
		sdk.NewCoin("bar", math.NewInt(150)),
		sdk.NewCoin(fooDenom, math.NewInt(50)),
	), res)

	// check community pool
	s.Require().Equal(math.NewInt(50), s.input.CommunityPoolKeeper.CommunityPool.AmountOf(fooDenom))

}

func (s *ERC20TestSuite) Test_BurnMultipleCoins() {
	erc20Keeper, err := keeper.NewERC20Keeper(&s.input.EVMKeeper)
	s.Require().NoError(err)

	evmAddr, addr := s.evmAddrs[0], s.accAddrs[0]
	_, denom0 := s.createContract("foo", evmAddr)
	_, denom1 := s.createContract("bar", evmAddr)
	// cannot mint erc20 from cosmos side
	cacheCtx, _ := s.ctx.CacheContext()
	err = erc20Keeper.MintCoins(cacheCtx, addr, sdk.NewCoins(
		sdk.NewCoin(denom0, math.NewInt(100)),
		sdk.NewCoin(denom1, math.NewInt(100)),
	))
	s.Require().Error(err)

	s.mintERC20(evmAddr, evmAddr, sdk.NewCoin(denom0, math.NewInt(100)), false)
	s.mintERC20(evmAddr, evmAddr, sdk.NewCoin(denom1, math.NewInt(100)), false)

	res, _, err := erc20Keeper.GetPaginatedBalances(s.ctx, nil, addr)
	s.Require().NoError(err)
	s.Require().Equal(sdk.NewCoins(
		sdk.NewCoin(denom0, math.NewInt(100)),
		sdk.NewCoin(denom1, math.NewInt(100)),
	), res)

	s.Require().True(s.input.CommunityPoolKeeper.CommunityPool.IsZero())
	err = erc20Keeper.BurnCoins(s.ctx, addr, sdk.NewCoins(
		sdk.NewCoin(denom0, math.NewInt(50)),
		sdk.NewCoin(denom1, math.NewInt(50)),
	))
	s.Require().NoError(err)

	s.Require().Equal(math.NewInt(50), s.input.CommunityPoolKeeper.CommunityPool.AmountOf(denom0))
	s.Require().Equal(math.NewInt(50), s.input.CommunityPoolKeeper.CommunityPool.AmountOf(denom1))

	res, _, err = erc20Keeper.GetPaginatedBalances(s.ctx, nil, addr)
	s.Require().NoError(err)
	s.Require().Equal(sdk.NewCoins(
		sdk.NewCoin(denom0, math.NewInt(50)),
		sdk.NewCoin(denom1, math.NewInt(50)),
	), res)
}

func (s *ERC20TestSuite) Test_GetSupply() {
	evmAddr, addr, addr2 := s.evmAddrs[0], s.accAddrs[0], s.accAddrs[1]

	// deploy erc20 contract
	_, fooDenom := s.createContract("foo", evmAddr)

	// mint erc20
	s.mintERC20(evmAddr, evmAddr, sdk.NewCoin(fooDenom, math.NewInt(100)), false)

	// mint native coin
	err := s.erc20Keeper.MintCoins(s.ctx, addr, sdk.NewCoins(
		sdk.NewCoin("bar", math.NewInt(200)),
	))
	s.Require().NoError(err)

	err = s.erc20Keeper.SendCoins(s.ctx, addr, addr2, sdk.NewCoins(
		sdk.NewCoin("bar", math.NewInt(100)),
		sdk.NewCoin(fooDenom, math.NewInt(50)),
	))
	s.Require().NoError(err)

	amount, err := s.erc20Keeper.GetSupply(s.ctx, fooDenom)
	s.Require().NoError(err)
	s.Require().Equal(math.NewInt(100), amount)

	has, err := s.erc20Keeper.HasSupply(s.ctx, fooDenom)
	s.Require().NoError(err)
	s.Require().True(has)

	amount, err = s.erc20Keeper.GetSupply(s.ctx, "bar")
	s.Require().NoError(err)
	s.Require().Equal(math.NewInt(200), amount)

	has, err = s.erc20Keeper.HasSupply(s.ctx, "bar")
	s.Require().NoError(err)
	s.Require().True(has)

	s.erc20Keeper.IterateSupply(s.ctx, func(supply sdk.Coin) (bool, error) {
		s.Require().True(supply.Denom == "bar" || supply.Denom == fooDenom || supply.Denom == sdk.DefaultBondDenom)
		switch supply.Denom {
		case "bar":
			s.Require().Equal(math.NewInt(200), supply.Amount)
		case fooDenom:
			s.Require().Equal(math.NewInt(100), supply.Amount)
		case sdk.DefaultBondDenom:
			s.Require().Equal(math.NewInt(1_000_000), supply.Amount)
		}
		return false, nil
	})

	supply, _, err := s.erc20Keeper.GetPaginatedSupply(s.ctx, nil)
	s.Require().NoError(err)
	s.Require().Equal(sdk.NewCoins(
		sdk.NewCoin("bar", math.NewInt(200)),
		sdk.NewCoin(fooDenom, math.NewInt(100)),
		sdk.NewCoin(sdk.DefaultBondDenom, math.NewInt(1_000_000)),
	), supply)
}

func (s *ERC20TestSuite) TestERC20Keeper_GetMetadata() {
	evmAddr, addr := s.evmAddrs[0], s.accAddrs[0]
	err := s.erc20Keeper.MintCoins(s.ctx, addr, sdk.NewCoins(
		sdk.NewCoin("bar", math.NewInt(200)),
		sdk.NewCoin("foo", math.NewInt(100)),
	))
	s.Require().NoError(err)

	supply, err := s.erc20Keeper.GetSupply(s.ctx, "foo")
	s.Require().NoError(err)
	s.Require().Equal(math.NewInt(100), supply)

	metadata, err := s.erc20Keeper.GetMetadata(s.ctx, "foo")
	s.Require().NoError(err)

	s.Require().Equal(banktypes.Metadata{
		Description: "",
		Base:        "foo",
		Display:     "foo",
		Name:        "foo",
		Symbol:      "foo",
		URI:         "",
		URIHash:     "",
		DenomUnits: []*banktypes.DenomUnit{
			{
				Denom:    "foo",
				Exponent: 0,
			},
		},
	}, metadata)

	factoryAbi, err := erc20_factory.Erc20FactoryMetaData.GetAbi()
	s.Require().NoError(err)

	callBz, err := factoryAbi.Pack("createERC20", "hey", "hey", uint8(18))
	s.Require().NoError(err)

	erc20WrapperAddr, err := s.input.EVMKeeper.ERC20FactoryAddr.Get(s.ctx)
	s.Require().NoError(err)
	retBz, _, err := s.input.EVMKeeper.EVMCall(s.ctx, evmAddr, common.BytesToAddress(erc20WrapperAddr), callBz, nil, nil)
	s.Require().NoError(err)
	s.Require().NotEmpty(retBz)

	ret, err := factoryAbi.Unpack("createERC20", retBz)
	s.Require().NoError(err)

	contractAddr := ret[0].(common.Address)
	denom, err := types.ContractAddrToDenom(s.ctx, &s.input.EVMKeeper, contractAddr)
	s.Require().NoError(err)

	metadata, err = s.erc20Keeper.GetMetadata(s.ctx, denom)
	s.Require().NoError(err)
	s.Require().Equal("hey", metadata.Name)
	s.Require().Equal("hey", metadata.Symbol)
	s.Require().Equal(uint32(18), metadata.DenomUnits[1].Exponent)
}

func (s *ERC20TestSuite) Test_IterateAccountBalances() {
	evmAddr, evmAddr2, addr, addr2 := s.evmAddrs[0], s.evmAddrs[1], s.accAddrs[0], s.accAddrs[1]
	// deploy erc20 contract
	_, fooDenom := s.createContract("foo", evmAddr)

	// mint erc20
	s.mintERC20(evmAddr, evmAddr, sdk.NewCoin(fooDenom, math.NewInt(100)), false)
	s.mintERC20(evmAddr, evmAddr2, sdk.NewCoin(fooDenom, math.NewInt(200)), false)

	// mint native coin
	err := s.erc20Keeper.MintCoins(s.ctx, addr, sdk.NewCoins(
		sdk.NewCoin("bar", math.NewInt(200)),
	))
	s.Require().NoError(err)
	err = s.erc20Keeper.MintCoins(s.ctx, addr2, sdk.NewCoins(
		sdk.NewCoin("bar", math.NewInt(400)),
	))
	s.Require().NoError(err)

	s.erc20Keeper.IterateAccountBalances(s.ctx, addr, func(balance sdk.Coin) (bool, error) {
		s.Require().True(balance.Denom == "bar" || balance.Denom == fooDenom)
		switch balance.Denom {
		case "bar":
			s.Require().Equal(math.NewInt(200), balance.Amount)
		case fooDenom:
			s.Require().Equal(math.NewInt(100), balance.Amount)
		}
		return false, nil
	})

	s.erc20Keeper.IterateAccountBalances(s.ctx, addr2, func(balance sdk.Coin) (bool, error) {
		s.Require().True(balance.Denom == "bar" || balance.Denom == fooDenom)
		switch balance.Denom {
		case "bar":
			s.Require().Equal(math.NewInt(400), balance.Amount)
		case fooDenom:
			s.Require().Equal(math.NewInt(200), balance.Amount)
		}
		return false, nil
	})
}

func (s *ERC20TestSuite) Test_Approve() {
	evmAddr, evmAddr2, addr, addr2 := s.evmAddrs[0], s.evmAddrs[1], s.accAddrs[0], s.accAddrs[1]

	// deploy erc20 contract
	_, fooDenom := s.createContract("foo", evmAddr)

	// mint erc20
	s.mintERC20(evmAddr, evmAddr, sdk.NewCoin(fooDenom, math.NewInt(100)), false)

	// approve erc20
	s.approveERC20(evmAddr, evmAddr2, sdk.NewCoin(fooDenom, math.NewInt(50)), false)

	// transferFrom erc20
	s.transferFromERC20(evmAddr2, evmAddr, evmAddr2, sdk.NewCoin(fooDenom, math.NewInt(50)), false)

	amount, err := s.erc20Keeper.GetBalance(s.ctx, addr, fooDenom)
	s.Require().NoError(err)
	s.Require().Equal(math.NewInt(50), amount)
	amount, err = s.erc20Keeper.GetBalance(s.ctx, addr2, fooDenom)
	s.Require().NoError(err)
	s.Require().Equal(math.NewInt(50), amount)

	// should fail to transferFrom more than approved
	s.transferFromERC20(evmAddr2, evmAddr, evmAddr2, sdk.NewCoin(fooDenom, math.NewInt(50)), true)
}

func (s *ERC20TestSuite) Test_ERC20MetadataUpdate() {
	evmAddr := s.evmAddrs[0]
	authorityAddr, err := s.input.AccountKeeper.AddressCodec().StringToBytes(s.input.EVMKeeper.GetAuthority())
	s.Require().NoError(err)
	authorityEVMAddr := common.BytesToAddress(authorityAddr)

	// deploy erc20 contract
	_, fooDenom := s.createContract("foo", evmAddr)

	// update metadata should fail because deployer is not 0x1
	err = s.updateMetadataERC20(authorityEVMAddr, fooDenom, "new name", "new symbol", 18)
	s.Require().Error(err)

	// create erc20 contract with deployer 0x1
	fooDenom = "foo"
	err = s.input.EVMKeeper.ERC20Keeper().CreateERC20(s.ctx, fooDenom, 6)
	s.Require().NoError(err)

	// update metadata
	err = s.updateMetadataERC20(authorityEVMAddr, fooDenom, "new name", "new symbol", 18)
	s.Require().NoError(err)
	metadata, err := s.input.EVMKeeper.ERC20Keeper().GetMetadata(s.ctx, fooDenom)
	s.Require().NoError(err)
	s.Require().Equal("new name", metadata.Name)
	s.Require().Equal("new symbol", metadata.Symbol)
	s.Require().Equal(uint32(18), metadata.DenomUnits[1].Exponent)

	// update metadata again should fail
	err = s.updateMetadataERC20(authorityEVMAddr, fooDenom, "new name", "new symbol", 18)
	s.Require().Error(err)
}
