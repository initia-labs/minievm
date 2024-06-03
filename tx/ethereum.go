package tx

import (
	"context"

	signingv1beta1 "cosmossdk.io/api/cosmos/tx/signing/v1beta1"
	"cosmossdk.io/x/tx/signing"
	sdk "github.com/cosmos/cosmos-sdk/types"

	coretypes "github.com/ethereum/go-ethereum/core/types"
)

const SignMode_SIGN_MODE_ETHEREUM = 9999

type convertCosmosTxToEthereumTxFunc func(sdkTx sdk.Tx) (*coretypes.Transaction, error)

// SignModeEthereumHandler defines the SIGN_MODE_DIRECT SignModeHandler
type SignModeEthereumHandler struct {
	convertCosmosTxToEthereumTxFunc
}

// NewSignModeEthereumHandler returns a new SignModeEthereumHandler.
func NewSignModeEthereumHandler(fun convertCosmosTxToEthereumTxFunc) *SignModeEthereumHandler {
	return &SignModeEthereumHandler{
		convertCosmosTxToEthereumTxFunc: fun,
	}
}

var _ signing.SignModeHandler = SignModeEthereumHandler{}

// Mode implements signing.SignModeHandler.Mode.
func (SignModeEthereumHandler) Mode() signingv1beta1.SignMode {
	return SignMode_SIGN_MODE_ETHEREUM /* custom code */
}

// GetSignBytes implements SignModeHandler.GetSignBytes
func (h SignModeEthereumHandler) GetSignBytes(
	ctx context.Context, data signing.SignerData, txData signing.TxData,
) ([]byte, error) {
	// TODO - convert cosmos tx to ethereum tx
	// and return the sign bytes

	// need to refer cosmos-sdk/client/tx/aux_builder.go
	// to convert txData to sdk.Tx
	h.convertCosmosTxToEthereumTxFunc(txData)

	return nil, nil
}
