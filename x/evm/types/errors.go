package types

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// EVM Errors
var (
	// ErrInvalidAddressLength error for the invalid address length
	ErrInvalidAddressLength        = errorsmod.Register(ModuleName, 2, "Invalid address length: address must be 20 bytes to use EVM")
	ErrEVMCallFailed               = errorsmod.Register(ModuleName, 3, "EVMCall failed")
	ErrEVMCreateFailed             = errorsmod.Register(ModuleName, 4, "EVMCreate failed")
	ErrUnknownPrecompileMethod     = errorsmod.Register(ModuleName, 5, "Unknown precompile method")
	ErrInvalidHexString            = errorsmod.Register(ModuleName, 6, "Invalid hex string")
	ErrFailedToDecodeOutput        = errorsmod.Register(ModuleName, 7, "Failed to decode output")
	ErrInvalidDenom                = errorsmod.Register(ModuleName, 8, "Invalid denom")
	ErrInvalidRequest              = errorsmod.Register(ModuleName, 9, "Invalid request")
	ErrFailedToPackABI             = errorsmod.Register(ModuleName, 10, "Failed to pack ABI")
	ErrFailedToUnpackABI           = errorsmod.Register(ModuleName, 11, "Failed to unpack ABI")
	ErrNonReadOnlyMethod           = errorsmod.Register(ModuleName, 12, "Failed to call precompile in readonly mode")
	ErrAddressAlreadyExists        = errorsmod.Register(ModuleName, 13, "Address already exists")
	ErrFailedToEncodeLogs          = errorsmod.Register(ModuleName, 14, "Failed to encode logs")
	ErrPrecompileFailed            = errorsmod.Register(ModuleName, 16, "Precompile failed")
	ErrNotSupportedCosmosMessage   = errorsmod.Register(ModuleName, 17, "Not supported cosmos message")
	ErrNotSupportedCosmosQuery     = errorsmod.Register(ModuleName, 18, "Not supported cosmos query")
	ErrInvalidTokenId              = errorsmod.Register(ModuleName, 19, "Invalid token id")
	ErrInvalidClassId              = errorsmod.Register(ModuleName, 20, "Invalid class id")
	ErrCustomERC20NotAllowed       = errorsmod.Register(ModuleName, 21, "Custom ERC20 is not allowed")
	ErrInvalidERC20FactoryAddr     = errorsmod.Register(ModuleName, 22, "Invalid ERC20 factory address")
	ErrReverted                    = errorsmod.Register(ModuleName, 23, "Reverted")
	ErrInvalidValue                = errorsmod.Register(ModuleName, 24, "Invalid value")
	ErrFailedToGetERC20FactoryAddr = errorsmod.Register(ModuleName, 25, "Failed to get ERC20 factory address")
	ErrInvalidFeeDenom             = errorsmod.Register(ModuleName, 26, "Invalid fee denom")
	ErrInvalidGasRefundRatio       = errorsmod.Register(ModuleName, 27, "Invalid gas refund ratio")
)

func NewRevertError(revert []byte) error {
	err := ErrReverted

	reason, errUnpack := abi.UnpackRevert(revert)
	if errUnpack == nil {
		return err.Wrapf("reason: %v, revert: %v", reason, hexutil.Encode(revert))
	}

	return err.Wrapf("revert: %v", hexutil.Encode(revert))
}
