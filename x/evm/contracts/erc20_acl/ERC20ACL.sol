// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "../i_cosmos/ICosmos.sol";

/// @dev CHAIN_ADDRESS is the address of the chain signer.
address constant CHAIN_ADDRESS = 0x0000000000000000000000000000000000000001;

/**
 * @title ERC20ACL
 */
contract ERC20ACL {
    modifier onlyChain() {
        require(msg.sender == CHAIN_ADDRESS, "ERC20: caller is not the chain");
        _;
    }

    // check if the sender is a module address
    modifier burnable(address from) {
        require(
            !COSMOS_CONTRACT.is_module_address(from),
            "ERC20: burn from module address"
        );

        _;
    }

    // check if the recipient is a blocked address
    modifier mintable(address to) {
        require(
            !COSMOS_CONTRACT.is_blocked_address(to),
            "ERC20: mint to blocked address"
        );

        _;
    }

    // check if an recipient is blocked in bank module
    modifier transferable(address to) {
        require(
            !COSMOS_CONTRACT.is_blocked_address(to),
            "ERC20: transfer to blocked address"
        );

        _;
    }
}
