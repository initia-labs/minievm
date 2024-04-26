// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

/// @dev The IERC721Registry contract's address.
address constant ERC721_REGISTRY_ADDRESS = 0x00000000000000000000000000000000000000f3;

/// @dev The IERC721Registry contract's instance.
IERC721Registry constant ERC721_REGISTRY_CONTRACT = IERC721Registry(
    ERC721_REGISTRY_ADDRESS
);

interface IERC721Registry {
    function register_erc721() external;
    function register_erc721_store(address account) external;
    function is_erc721_store_registered(
        address account
    ) external view returns (bool registered);
}
