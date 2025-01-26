package jsonutils

import (
	storetypes "cosmossdk.io/store/types"
)

type MergeJSONArguments struct {
	DstJSON string `abi:"dst_json"`
	SrcJSON string `abi:"src_json"`
}

type StringifyArguments struct {
	JSON string `abi:"json"`
}

const (
	MERGE_GAS          storetypes.Gas = 100
	STRINGIFY_JSON_GAS storetypes.Gas = 100

	GAS_PER_BYTE storetypes.Gas = 1

	MaxDepth = 16
)
