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

    // convert an ERC20 address to a Cosmos denom
    function to_denom(address erc20_address) external returns (string memory denom);

    // convert a Cosmos denom to an ERC20 address
    function to_erc20(string memory denom) external returns (address erc20_address);

    // record a cosmos message to be executed
    // after the current message execution.
    //
    // msg should be in json string format like:
    // {
    //    "@type": "/cosmos.bank.v1beta1.MsgSend",
    //    "from_address": "init13vhzmdmzsqlxkdzvygue9zjtpzedz7j87c62q4",
    //    "to_address": "init1enjh88u7c9s08fgdu28wj6umz94cetjy0hpcxf",
    //    "amount": [
    //        {
    //            "denom": "stake",
    //            "amount": "100"
    //        }
    //    ]
    // }
    //
    function execute_cosmos(string memory msg) external;

    // query a whitelisted cosmos querys.
    //
    // example)
    // path: "/slinky.oracle.v1.Query/GetPrices"
    // req: {
    //    "currency_pair_ids": ["BITCOIN/USD", "ETHEREUM/USD"]
    // }
    //
    function query_cosmos(
        string memory path,
        string memory req
    ) external returns (string memory result);
}
