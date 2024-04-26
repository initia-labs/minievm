// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "../i_erc721_registry/IERC721Registry.sol";

/**
 * @title ERC721Registry
 */
contract ERC721Registry {
    modifier register_erc721() {
        ERC721_REGISTRY_CONTRACT.register_erc721();

        _;
    }

    modifier register_erc721_store(address account) {
        if (
            !ERC721_REGISTRY_CONTRACT.is_erc721_store_registered(
                account
            )
        ) {
            ERC721_REGISTRY_CONTRACT.register_erc721_store(
                account
            );
        }

        _;
    }
}
