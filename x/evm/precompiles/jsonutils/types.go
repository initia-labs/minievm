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

type UnmarshalJSONArguments struct {
	JSONBytes []byte `abi:"json_bytes"`
}

const (
	MERGE_GAS          storetypes.Gas = 100
	STRINGIFY_JSON_GAS storetypes.Gas = 100
	UNMARSHAL_JSON_GAS storetypes.Gas = 100

	GAS_PER_BYTE      storetypes.Gas = 1
	GAS_PER_SORT_ITEM storetypes.Gas = 10

	MaxDepth    = 16
	MaxJSONSize = 10 * 1024
)
