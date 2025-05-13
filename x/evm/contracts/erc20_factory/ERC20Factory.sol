// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import "../erc20/ERC20.sol";
import "../i_erc20_registry/IERC20Registry.sol";

contract ERC20Factory is ERC20Registry {
    event ERC20Created(address indexed erc20, address indexed owner);

    /**
     * @notice Create a new ERC20 contract
     * @param name The name of the ERC20 contract
     * @param symbol The symbol of the ERC20 contract
     * @param decimals The decimals of the ERC20 contract
     * @return The address of the new ERC20 contract
     */
    function createERC20(
        string memory name,
        string memory symbol,
        uint8 decimals
    ) external returns (address) {
        ERC20 erc20 = new ERC20(
            name,
            symbol,
            decimals,
            msg.sender != CHAIN_ADDRESS
        );
        address erc20Addr = address(erc20);

        _handlePostCreation(erc20Addr);
        return erc20Addr;
    }

    /**
     * @notice Create a new ERC20 contract with a salt
     * @param name The name of the ERC20 contract
     * @param symbol The symbol of the ERC20 contract
     * @param decimals The decimals of the ERC20 contract
     * @param salt The salt to use for the ERC20 contract. it will be hashed with the sender's address to avoid collisions
     * @return The address of the new ERC20 contract
     */
    function createERC20WithSalt(
        string memory name,
        string memory symbol,
        uint8 decimals,
        bytes32 salt
    ) external returns (address) {
        address erc20Addr;
        bytes32 _salt = keccak256(abi.encodePacked(msg.sender, salt));

        // prepare the bytecode for the ERC20 contract
        bytes memory bytecode = abi.encodePacked(
            type(ERC20).creationCode,
            abi.encode(name, symbol, decimals, msg.sender != CHAIN_ADDRESS)
        );

        // deploy the ERC20 contract
        assembly {
            erc20Addr := create2(0, add(bytecode, 0x20), mload(bytecode), _salt)
            if iszero(extcodesize(erc20Addr)) {
                revert(0, 0)
            }
        }

        _handlePostCreation(erc20Addr);
        return erc20Addr;
    }

    /**
     * @notice Create a new ERC20 contract with a salt and a custom bytecode
     * @param creator The address of the creator of the ERC20 contract
     * @param name The name of the ERC20 contract
     * @param symbol The symbol of the ERC20 contract
     * @param decimals The decimals of the ERC20 contract
     * @param salt The salt to use for the ERC20 contract. it will be hashed with the sender's address to avoid collisions
     * @return addr The address of the new ERC20 contract
     */
    function computeERC20Address(
        address creator,
        string memory name,
        string memory symbol,
        uint8 decimals,
        bytes32 salt
    ) external pure returns (address addr) {
        bytes memory bytecode = abi.encodePacked(
            type(ERC20).creationCode,
            abi.encode(name, symbol, decimals, creator!= CHAIN_ADDRESS)
        );
        bytes32 bytecodeHash = keccak256(bytecode);
        assembly ("memory-safe") {
            let ptr := mload(0x40)
            mstore(add(ptr, 0x40), bytecodeHash)
            mstore(add(ptr, 0x20), salt)
            mstore(ptr, creator) 
            let start := add(ptr, 0x0b)
            mstore8(start, 0xff)
            addr := and(
                keccak256(start, 85),
                0xffffffffffffffffffffffffffffffffffffffff
            )
        }
    }

    /**
     * @notice Post-action to be performed after the ERC20 contract is created
     * @param erc20Addr The address of the ERC20 contract
     */
    function _handlePostCreation(address erc20Addr) internal {
        // register the ERC20 contract with the ERC20 registry
        ERC20_REGISTRY_CONTRACT.register_erc20_from_factory(erc20Addr);

        // transfer ownership of the ERC20 contract to the sender
        ERC20(erc20Addr).transferOwnership(msg.sender);

        emit ERC20Created(erc20Addr, msg.sender);
    }
}
