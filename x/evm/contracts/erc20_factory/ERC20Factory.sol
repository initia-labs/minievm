// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import "../erc20/ERC20.sol";
import "../i_erc20_registry/IERC20Registry.sol";

contract ERC20Factory is ERC20Registry {
    event ERC20Created(address indexed erc20, address indexed owner);

    function createERC20(
        string memory name,
        string memory symbol,
        uint8 decimals
    ) external returns (address) {
        // create the ERC20 contract with a deterministic address using create2
        address erc20Addr;
        bytes32 salt = keccak256(abi.encodePacked(msg.sender, symbol, decimals));

        // prepare the bytecode for the ERC20 contract
        bytes memory bytecode = abi.encodePacked(
            type(ERC20).creationCode,
            abi.encode(name, symbol, decimals, msg.sender != CHAIN_ADDRESS)
        );

        // deploy the ERC20 contract
        assembly {
            erc20Addr := create2(0, add(bytecode, 0x20), mload(bytecode), salt)
            if iszero(extcodesize(erc20Addr)) {
                revert(0, 0)
            }
        }

        // register the ERC20 contract with the ERC20 registry
        ERC20_REGISTRY_CONTRACT.register_erc20_from_factory(erc20Addr);

        // transfer ownership of the ERC20 contract to the sender
        ERC20(erc20Addr).transferOwnership(msg.sender);

        emit ERC20Created(erc20Addr, msg.sender);
        return erc20Addr;
    }
}
