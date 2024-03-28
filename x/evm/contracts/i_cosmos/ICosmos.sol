// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

/// @dev The ICosmos contract's address.
address constant COSMOS_ADDRESS = 0x00000000000000000000000000000000000000F2;

/// @dev The ICosmos contract's instance.
ICosmos constant COSMOS_CONTRACT = ICosmos(COSMOS_ADDRESS);

interface ICosmos {
    // convert an EVM address to a Cosmos address
    function to_cosmos_address(
        address evm_address
    ) external returns (string memory cosmos_address);

    // convert a Cosmos address to an EVM address
    function to_evm_address(
        string memory cosmos_address
    ) external returns (address evm_address);

    // record a cosmos message to be executed
    // after the current message execution.
    function execute_cosmos_message(string memory msg) external;
}
