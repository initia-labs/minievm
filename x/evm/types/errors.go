package types

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
)

// EVM Errors
var (
	// ErrInvalidAddressLength error for the invalid address length
	ErrInvalidAddressLength         = errorsmod.Register(ModuleName, 2, "Invalid address length: address must be 20 bytes to use EVM")
	ErrEVMCallFailed                = errorsmod.Register(ModuleName, 3, "EVMCall failed")
	ErrEVMCreateFailed              = errorsmod.Register(ModuleName, 4, "EVMCreate failed")
	ErrUnknownPrecompileMethod      = errorsmod.Register(ModuleName, 5, "Unknown precompile method")
	ErrInvalidHexString             = errorsmod.Register(ModuleName, 6, "Invalid hex string")
	ErrFailedToDecodeOutput         = errorsmod.Register(ModuleName, 7, "Failed to decode output")
	ErrInvalidDenom                 = errorsmod.Register(ModuleName, 8, "Invalid denom")
	ErrInvalidRequest               = errorsmod.Register(ModuleName, 9, "Invalid request")
	ErrFailedToPackABI              = errorsmod.Register(ModuleName, 10, "Failed to pack ABI")
	ErrFailedToUnpackABI            = errorsmod.Register(ModuleName, 11, "Failed to unpack ABI")
	ErrNonReadOnlyMethod            = errorsmod.Register(ModuleName, 12, "Failed to call precompile in readonly mode")
	ErrAddressAlreadyExists         = errorsmod.Register(ModuleName, 13, "Address already exists")
	ErrFailedToEncodeLogs           = errorsmod.Register(ModuleName, 14, "Failed to encode logs")
	ErrPrecompileFailed             = errorsmod.Register(ModuleName, 16, "Precompile failed")
	ErrNotSupportedCosmosMessage    = errorsmod.Register(ModuleName, 17, "Not supported cosmos message")
	ErrNotSupportedCosmosQuery      = errorsmod.Register(ModuleName, 18, "Not supported cosmos query")
	ErrInvalidTokenId               = errorsmod.Register(ModuleName, 19, "Invalid token id")
	ErrInvalidClassId               = errorsmod.Register(ModuleName, 20, "Invalid class id")
	ErrCustomERC20NotAllowed        = errorsmod.Register(ModuleName, 21, "Custom ERC20 is not allowed")
	ErrInvalidERC20FactoryAddr      = errorsmod.Register(ModuleName, 22, "Invalid ERC20 factory address")
	ErrReverted                     = errorsmod.Register(ModuleName, 23, "Reverted")
	ErrInvalidValue                 = errorsmod.Register(ModuleName, 24, "Invalid value")
	ErrFailedToGetERC20FactoryAddr  = errorsmod.Register(ModuleName, 25, "Failed to get ERC20 factory address")
	ErrInvalidFeeDenom              = errorsmod.Register(ModuleName, 26, "Invalid fee denom")
	ErrInvalidGasRefundRatio        = errorsmod.Register(ModuleName, 27, "Invalid gas refund ratio")
	ErrFailedToGetERC20WrapperAddr  = errorsmod.Register(ModuleName, 28, "Failed to get ERC20 wrapper address")
	ErrInvalidNumRetainBlockHashes  = errorsmod.Register(ModuleName, 29, "Invalid num retain block hashes")
	ErrExceedMaxRecursiveDepth      = errorsmod.Register(ModuleName, 30, "Exceed max recursive depth")
	ErrTxConversionFailed           = errorsmod.Register(ModuleName, 31, "Tx conversion failed")
	ErrFailedToGetConnectOracleAddr = errorsmod.Register(ModuleName, 32, "Failed to get ConnectOracle address")
	ErrInvalidSalt                  = errorsmod.Register(ModuleName, 33, "Invalid salt")
	ErrExecuteCosmosDisabled        = errorsmod.Register(ModuleName, 34, "Execute cosmos is disabled")
	ErrInvalidGasEnforcement        = errorsmod.Register(ModuleName, 35, "Invalid gas enforcement parameters")
)

// ModError is a wrapper for the errorsmod.Error
type ModError = *errorsmod.Error

// RevertError is a wrapper for the vm.RevertError
type RevertError struct {
	ModError
	ret []byte
}

// Ret returns the revert data
func (e RevertError) Ret() []byte {
	return e.ret
}

// NewRevertError creates a new RevertError
//
//	revert: the revert data
//
// Returns:
// - RevertError: the new RevertError
func NewRevertError(revert []byte) error {
	revertErr := &RevertError{
		ModError: ErrReverted,
		ret:      revert,
	}

	return revertErr
}

// Error returns the error message
func (e *RevertError) Error() string {
	reason, errUnpack := abi.UnpackRevert(e.ret)
	if errUnpack != nil {
		return e.ModError.Error()
	}

	return e.ModError.Wrapf("revert: %v", reason).Error()
}

// revertSelector is a special function selector for revert reason unpacking.
var revertSelector = crypto.Keccak256([]byte("Error(string)"))[:4]
var revertABIArguments abi.Arguments

func init() {
	revertABIType, err := abi.NewType("string", "", nil)
	if err != nil {
		panic(err)
	}

	revertABIArguments = abi.Arguments{{Type: revertABIType}}
}

func NewRevertReason(reason error) []byte {
	bz, err := revertABIArguments.Pack(reason.Error())
	if err != nil {
		panic(err)
	}

	ret := make([]byte, 4+len(bz))
	copy(ret, revertSelector)
	copy(ret[4:], bz)

	return ret
}
