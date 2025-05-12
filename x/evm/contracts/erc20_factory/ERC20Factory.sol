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
        // try to create the ERC20 contract with create2
        // if it fails, create the ERC20 contract with the fallback constructor
        address erc20Addr;
        try this.createERC20WithCreate2(msg.sender, name, symbol, decimals) returns (
            address _erc20Addr
        ) {
            erc20Addr = _erc20Addr;
        } catch {
            ERC20 erc20 = new ERC20(name, symbol, decimals, msg.sender != CHAIN_ADDRESS);
            erc20Addr = address(erc20);
        }

        // register the ERC20 contract with the ERC20 registry
        ERC20_REGISTRY_CONTRACT.register_erc20_from_factory(erc20Addr);

        // transfer ownership of the ERC20 contract to the sender
        ERC20(erc20Addr).transferOwnership(msg.sender);

        emit ERC20Created(erc20Addr, msg.sender);
        return erc20Addr;
    }

    /*
     * @notice Create a new ERC20 contract with create2
     * @param name The name of the ERC20 contract
     * @param symbol The symbol of the ERC20 contract
     * @param decimals The decimals of the ERC20 contract
     * @return The address of the new ERC20 contract
     */
    function createERC20WithCreate2(
        address creator,
        string memory name,
        string memory symbol,
        uint8 decimals
    ) external returns (address) {
        require(msg.sender == address(this), "ERC20Factory: only the factory can call this function");

        address erc20Addr;
        bytes32 salt = keccak256(abi.encodePacked(creator, symbol, decimals));

        // prepare the bytecode for the ERC20 contract
        bytes memory bytecode = abi.encodePacked(
            type(ERC20).creationCode,
            abi.encode(name, symbol, decimals, creator != CHAIN_ADDRESS)
        );

        // deploy the ERC20 contract
        assembly {
            erc20Addr := create2(0, add(bytecode, 0x20), mload(bytecode), salt)
            if iszero(extcodesize(erc20Addr)) {
                revert(0, 0)
            }
        }

        return erc20Addr;
    }
}
