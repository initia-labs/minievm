package main

import (
	"context"
	"encoding/json"
	"math/big"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/initia-labs/OPinit/contrib/launchtools"
	"github.com/initia-labs/OPinit/contrib/launchtools/steps"
	"github.com/initia-labs/OPinit/contrib/launchtools/utils"
	ophosttypes "github.com/initia-labs/OPinit/x/ophost/types"

	"github.com/initia-labs/initia/app/params"

	minitiaapp "github.com/initia-labs/minievm/app"
	"github.com/initia-labs/minievm/types"
	"github.com/initia-labs/minievm/x/evm/contracts/erc20_wrapper"
	evmtypes "github.com/initia-labs/minievm/x/evm/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// DefaultLaunchStepFactories is a list of default launch step factories.
var DefaultLaunchStepFactories = []launchtools.LauncherStepFuncFactory[*launchtools.Config]{
	steps.InitializeConfig,
	steps.InitializeRPCHelpers,

	// Initialize genesis
	steps.InitializeGenesis,

	// Add system keys to the keyring
	steps.InitializeKeyring,

	// Run the app
	steps.RunApp,

	// Establish IBC channels for fungible and NFT transfer
	steps.EstablishIBCChannelsWithNFTTransfer(func() (string, string, string) {
		return "nft-transfer", "nft-transfer", "ics721-1"
	}),

	// Create OP Bridge, using open channel states
	steps.InitializeOpBridge,

	// Set bridge info and update clients
	steps.SetBridgeInfo,

	// Get the L1 and L2 heights
	steps.GetL1Height,
	steps.GetL2Height,

	// Initiate deposit to L2 with 0 amount to create wrapped INIT token on L2
	CreateINITWrappedToken,

	// Cleanup
	steps.StopApp,
}

const (
	INIT_DENOM = "uinit"
	MIN_DENOM  = "umin"
)

func LaunchCommand(ac *appCreator, enc params.EncodingConfig, mbm module.BasicManager) *cobra.Command {
	return launchtools.LaunchCmd(
		ac,
		func(denom string) map[string]json.RawMessage {
			// default denom in OPinit is "umin"
			if denom == MIN_DENOM {
				// convert to "GAS"
				denom = types.BaseDenom
			}

			return minitiaapp.NewDefaultGenesisState(enc.Codec, mbm, denom)
		},
		DefaultLaunchStepFactories,
	)
}

func CreateINITWrappedToken(config *launchtools.Config) launchtools.LauncherStepFunc {
	return func(ctx launchtools.Launcher) error {
		ctx.Logger().Info("Initiating token deposit...")

		bridgeId := *ctx.GetBridgeId()
		clientCtx := (*ctx.ClientContext()).WithChainID(config.L2Config.ChainID)
		executorL1AddrStr := config.SystemKeys.BridgeExecutor.L1Address
		executorL2AddrStr := config.SystemKeys.BridgeExecutor.L2Address
		executorMnemonic := config.SystemKeys.BridgeExecutor.Mnemonic

		// use operator account to create hook message to avoid sequence mismatch
		// because the bridge executor need to send a tx to relay the initiate token deposit message
		operatorL2AddrStr := config.SystemKeys.Validator.L2Address
		operatorMnemonic := config.SystemKeys.Validator.Mnemonic

		// generate op bridge hook message
		hookMsg, err := generateOPBridgeHookMessage(clientCtx, bridgeId, operatorL2AddrStr, operatorMnemonic)
		if err != nil {
			return errors.Wrap(err, "failed to generate op bridge hook message")
		}

		// create initiate token deposit message
		initiateTokenDepositMsg := ophosttypes.NewMsgInitiateTokenDeposit(
			executorL1AddrStr,
			bridgeId,
			executorL2AddrStr,
			sdk.NewCoin(INIT_DENOM, sdkmath.NewInt(0)),
			hookMsg,
		)

		ctx.Logger().Info("broadcasting tx to L1...",
			"from-address", executorL1AddrStr,
		)

		// already validated in config.go
		gasPrices, _ := sdk.ParseDecCoins(config.L1Config.GasPrices)
		gasFees, _ := gasPrices.MulDec(sdkmath.LegacyNewDecFromInt(sdkmath.NewInt(200000))).TruncateDecimal()

		// send initiate token deposit message to host (L1)
		res, err := ctx.GetRPCHelperL1().BroadcastTxAndWait(
			executorL1AddrStr,
			executorMnemonic,
			200000,
			gasFees,
			initiateTokenDepositMsg,
		)
		if err != nil {
			return errors.Wrap(err, "failed to broadcast initiate token deposit tx")
		}

		// if transaction is failed, return error
		if res.TxResult.Code != 0 {
			ctx.Logger().Error("tx failed", "code", res.TxResult.Code, "log", res.TxResult.Log)
			return errors.Errorf("tx failed with code %d", res.TxResult.Code)
		}

		ctx.Logger().Info("token deposit initiated")
		return nil
	}
}

// generateOPBridgeHookMessage generates a hook message for the OP bridge to create
// a wrapped token on L2.
func generateOPBridgeHookMessage(
	clientCtx client.Context,
	bridgeId uint64,
	operatorL2AddrStr string,
	operatorMnemonic string,
) ([]byte, error) {
	wrapperABI, err := erc20_wrapper.Erc20WrapperMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get abi")
	}

	// get operator address bytes
	operatorAddr, err := utils.L2AddressCodec().StringToBytes(operatorL2AddrStr)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get l2 address")
	}

	// compute l2 denom
	l2Denom := ophosttypes.L2Denom(bridgeId, INIT_DENOM)

	// fetch erc20 wrapper address
	wrapperAddr, err := loadERC20WrapperAddress(clientCtx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load erc20 wrapper address")
	}

	// pack abi
	input, err := wrapperABI.Pack("toLocal", common.BytesToAddress(operatorAddr), l2Denom, big.NewInt(0), uint8(6))
	if err != nil {
		return nil, errors.Wrap(err, "failed to pack abi")
	}

	// get bridge info
	msg := evmtypes.MsgCall{
		Sender:       operatorL2AddrStr,
		ContractAddr: wrapperAddr.Hex(),
		Input:        hexutil.Encode(input),
		Value:        sdkmath.NewInt(0),
	}

	// load account info
	acc, err := loadAccountInfo(clientCtx, operatorAddr)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load account info")
	}

	// sign the transaction
	signedTx, err := utils.SignTxOffline(
		&clientCtx,
		operatorMnemonic,
		1,
		acc.GetAccountNumber(),
		acc.GetSequence(),
		sdk.NewCoins(),
		&msg,
	)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to sign transaction")
	}

	// encode the transaction
	txBytes, err := clientCtx.TxConfig.TxEncoder()(signedTx)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to encode transaction")
	}

	return txBytes, nil
}

// loadERC20WrapperAddress loads the address of the ERC20 wrapper contract.
func loadERC20WrapperAddress(clientCtx client.Context) (common.Address, error) {
	queryClient := evmtypes.NewQueryClient(clientCtx)
	res, err := queryClient.ERC20Wrapper(context.Background(), &evmtypes.QueryERC20WrapperRequest{})
	if err != nil {
		return common.Address{}, errors.Wrap(err, "failed to get erc20 wrapper")
	}

	return common.HexToAddress(res.Address), nil
}

// loadAccountInfo loads the account info of the given address.
func loadAccountInfo(clientCtx client.Context, addr sdk.AccAddress) (client.Account, error) {
	ar := authtypes.AccountRetriever{}
	return ar.GetAccount(
		clientCtx,
		addr,
	)
}
