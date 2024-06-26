// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

/// @dev The IERC20Registry contract's address.
address constant ERC20_REGISTRY_ADDRESS = 0x00000000000000000000000000000000000000F2;

/// @dev The IERC20Registry contract's instance.
IERC20Registry constant ERC20_REGISTRY_CONTRACT = IERC20Registry(
    ERC20_REGISTRY_ADDRESS
);

interface IERC20Registry {
    function register_erc20() external returns (bool dummy);
    function register_erc20_from_factory(
        address erc20
    ) external returns (bool dummy);
    function register_erc20_store(
        address account
    ) external returns (bool dummy);
    function is_erc20_store_registered(
        address account
    ) external view returns (bool registered);
}
