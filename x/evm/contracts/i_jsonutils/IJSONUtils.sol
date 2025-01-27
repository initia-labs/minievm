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
}
