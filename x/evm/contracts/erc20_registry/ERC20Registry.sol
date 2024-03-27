// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "../i_erc20_registry/IERC20Registry.sol";

/**
 * @title ERC20Registry
 */
contract ERC20Registry {
    modifier register_erc20() {
        ERC20_REGISTRY_CONTRACT.register_erc20();

        _;
    }

    modifier register_erc20_store(address account) {
        if (
            !ERC20_REGISTRY_CONTRACT.is_erc20_store_registered(
                account
            )
        ) {
            ERC20_REGISTRY_CONTRACT.register_erc20_store(
                account
            );
        }

        _;
    }
}
