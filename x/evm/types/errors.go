package types

import (
	errorsmod "cosmossdk.io/errors"
)

// EVM Errors
var (
	// ErrInvalidAddressLength error for the invalid address length
	ErrInvalidAddressLength = errorsmod.Register(ModuleName, 2, "address must be 20 bytes to use EVM")
	ErrEVMCallFailed        = errorsmod.Register(ModuleName, 3, "EVMCall failed")
	ErrEVMCreateFailed      = errorsmod.Register(ModuleName, 4, "EVMCreate failed")
)
