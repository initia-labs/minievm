package ante

import (
	"context"
	"fmt"

	signingv1beta1 "cosmossdk.io/api/cosmos/tx/signing/v1beta1"
	errorsmod "cosmossdk.io/errors"
	txsigning "cosmossdk.io/x/tx/signing"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/crypto/types/multisig"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"

	evmkeeper "github.com/initia-labs/minievm/x/evm/keeper"
	evmtypes "github.com/initia-labs/minievm/x/evm/types"

	coretypes "github.com/ethereum/go-ethereum/core/types"
)

// internalSignModeToAPI converts a signing.SignMode to a protobuf SignMode.
func internalSignModeToAPI(mode signing.SignMode) (signingv1beta1.SignMode, error) {
	switch mode {
	case signing.SignMode_SIGN_MODE_DIRECT:
		return signingv1beta1.SignMode_SIGN_MODE_DIRECT, nil
	case signing.SignMode_SIGN_MODE_LEGACY_AMINO_JSON:
		return signingv1beta1.SignMode_SIGN_MODE_LEGACY_AMINO_JSON, nil
	case signing.SignMode_SIGN_MODE_TEXTUAL:
		return signingv1beta1.SignMode_SIGN_MODE_TEXTUAL, nil
	case signing.SignMode_SIGN_MODE_DIRECT_AUX:
		return signingv1beta1.SignMode_SIGN_MODE_DIRECT_AUX, nil
	case signing.SignMode_SIGN_MODE_EIP_191:
		return signingv1beta1.SignMode_SIGN_MODE_EIP_191, nil //nolint:staticcheck
	default:
		return signingv1beta1.SignMode_SIGN_MODE_UNSPECIFIED, fmt.Errorf("unsupported sign mode %s", mode)
	}
}

// verifySignature verifies a transaction signature contained in SignatureData abstracting over different signing
// modes. It differs from verifySignature in that it uses the new txsigning.TxData interface in x/tx.
func verifySignature(
	ctx context.Context,
	pubKey cryptotypes.PubKey,
	signerData txsigning.SignerData,
	signatureData signing.SignatureData,
	handler *txsigning.HandlerMap,
	txData txsigning.TxData,
	// required to verify EVM signatures
	ek *evmkeeper.Keeper,
	tx sdk.Tx,
) error {
	switch data := signatureData.(type) {
	case *signing.SingleSignatureData:
		if data.SignMode == evmkeeper.SignMode_SIGN_MODE_ETHEREUM {
			// eth sign mode
			ethTx, expectedSender, err := evmkeeper.NewTxUtils(ek).ConvertCosmosTxToEthereumTx(ctx, tx)
			if err != nil {
				return err
			}
			if ethTx == nil {
				return fmt.Errorf("failed to convert tx to ethereum tx")
			}

			sdkCtx := sdk.UnwrapSDKContext(ctx)
			ethChainID := evmtypes.ConvertCosmosChainIDToEthereumChainID(sdkCtx.ChainID())
			signer := coretypes.LatestSignerForChainID(ethChainID)
			sender, err := signer.Sender(ethTx)
			if err != nil {
				return errorsmod.Wrapf(sdkerrors.ErrorInvalidSigner, "failed to recover sender address: %v", err)
			}

			// check if the recovered sender matches the expected sender
			if expectedSender == nil || *expectedSender != sender {
				return errorsmod.Wrapf(sdkerrors.ErrorInvalidSigner, "expected sender %s, got %s", expectedSender, sender)
			}

			return nil
		}

		signMode, err := internalSignModeToAPI(data.SignMode)
		if err != nil {
			return err
		}
		signBytes, err := handler.GetSignBytes(ctx, signMode, signerData, txData)
		if err != nil {
			return err
		}

		if !pubKey.VerifySignature(signBytes, data.Signature) {
			return fmt.Errorf("unable to verify single signer signature")
		}

		return nil

	case *signing.MultiSignatureData:
		multiPK, ok := pubKey.(multisig.PubKey)
		if !ok {
			return fmt.Errorf("expected %T, got %T", (multisig.PubKey)(nil), pubKey)
		}
		err := multiPK.VerifyMultisignature(func(mode signing.SignMode) ([]byte, error) {
			signMode, err := internalSignModeToAPI(mode)
			if err != nil {
				return nil, err
			}
			return handler.GetSignBytes(ctx, signMode, signerData, txData)
		}, data)
		if err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("unexpected SignatureData %T", signatureData)
	}
}
