// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

/// @dev The IJSONUtils contract's address.
address constant JSONUTILS_ADDRESS = 0x00000000000000000000000000000000000000f3;

/// @dev The IJSONUtils contract's instance.
IJSONUtils constant JSONUTILS_CONTRACT = IJSONUtils(JSONUTILS_ADDRESS);

interface IJSONUtils {
    // recursively merges the src and dst maps. Key conflicts are resolved by
    // preferring src, or recursively descending, if both src and dst are maps.
    function merge_json(
        string memory dst_json,
        string memory src_json
    ) external view returns (string memory);

    // stringify the json string
    function stringify_json(
        string memory json
    ) external view returns (string memory);

    // JSON marshal/unmarshal functions

    struct JSONElement {
        string key;
        bytes value;
    }

    struct JSONObject {
        JSONElement[] elements;
    }

    function unmarshal_to_object(
        bytes memory json_bytes
    ) external view returns (JSONObject memory);

    function unmarshal_to_array(
        bytes memory json_bytes
    ) external view returns (bytes[] memory);

    function unmarshal_to_string(
        bytes memory json_bytes
    ) external view returns (string memory);

    // allow both "123" or 1234 to be converted to uint
    function unmarshal_to_uint(
        bytes memory json_bytes
    ) external view returns (uint256);

    function unmarshal_to_bool(
        bytes memory json_bytes
    ) external view returns (bool);

    // RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00" to unix timestamp (nano second)
    function unmarshal_iso_to_unix(
        bytes memory json_bytes
    ) external view returns (uint256);
}
