package types

import (
	errorsmod "cosmossdk.io/errors"
)

// EVM Errors
var (
	// ErrInvalidAddressLength error for the invalid address length
	ErrInvalidAddressLength    = errorsmod.Register(ModuleName, 2, "Invalid address length: address must be 20 bytes to use EVM")
	ErrEVMCallFailed           = errorsmod.Register(ModuleName, 3, "EVMCall failed")
	ErrEVMCreateFailed         = errorsmod.Register(ModuleName, 4, "EVMCreate failed")
	ErrUnknownPrecompileMethod = errorsmod.Register(ModuleName, 5, "Unknown precompile method")
	ErrInvalidHexString        = errorsmod.Register(ModuleName, 6, "Invalid hex string")
	ErrFailedToDecodeOutput    = errorsmod.Register(ModuleName, 7, "Failed to decode output")
	ErrInvalidDenom            = errorsmod.Register(ModuleName, 8, "Invalid denom")
	ErrInvalidRequest          = errorsmod.Register(ModuleName, 9, "Invalid request")
	ErrFailedToPackABI         = errorsmod.Register(ModuleName, 10, "Failed to pack ABI")
	ErrFailedToUnpackABI       = errorsmod.Register(ModuleName, 11, "Failed to unpack ABI")
	ErrNonReadOnlyMethod       = errorsmod.Register(ModuleName, 12, "Failed to call precompile in readonly mode")
	ErrAddressAlreadyExists    = errorsmod.Register(ModuleName, 13, "Address already exists")
	ErrFailedToEncodeLogs      = errorsmod.Register(ModuleName, 14, "Failed to encode logs")
	ErrEmptyContractAddress    = errorsmod.Register(ModuleName, 15, "Empty contract address")
)
