// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "../erc20/ERC20.sol";
import "../i_erc20_registry/IERC20Registry.sol";

contract ERC20Factory is ERC20Registry {
    event ERC20Created(address indexed erc20, address indexed owner);

    function createERC20(
        string memory name,
        string memory symbol,
        uint8 decimals
    ) external returns (address) {
        ERC20 erc20 = new ERC20(name, symbol, decimals);
        
        // register the ERC20 contract with the ERC20 registry
        ERC20_REGISTRY_CONTRACT.register_erc20_from_factory(address(erc20));

        // transfer ownership of the ERC20 contract to the sender
        erc20.transferOwnership(msg.sender);

        emit ERC20Created(address(erc20), msg.sender);
        return address(erc20);
    }
}