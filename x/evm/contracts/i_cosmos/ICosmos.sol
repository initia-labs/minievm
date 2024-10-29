// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

/// @dev The ICosmos contract's address.
address constant COSMOS_ADDRESS = 0x00000000000000000000000000000000000000f1;

/// @dev The ICosmos contract's instance.
ICosmos constant COSMOS_CONTRACT = ICosmos(COSMOS_ADDRESS);

interface ICosmos {
    // check if an address is blocked in bank module
    function is_blocked_address(
        address account
    ) external view returns (bool blocked);

    // check if an address is a module account
    function is_module_address(
        address account
    ) external view returns (bool module);

    // convert an EVM address to a Cosmos address
    function to_cosmos_address(
        address evm_address
    ) external returns (string memory cosmos_address);

    // convert a Cosmos address to an EVM address
    function to_evm_address(
        string memory cosmos_address
    ) external returns (address evm_address);

    // convert an ERC20 address to a Cosmos denom
    function to_denom(
        address erc20_address
    ) external returns (string memory denom);

    // convert a Cosmos denom to an ERC20 address
    function to_erc20(
        string memory denom
    ) external returns (address erc20_address);

    // record a cosmos message to be executed after the current message execution.
    // - if execution fails, whole transaction will be reverted.
    //
    // `msg` format (json string):
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
    function execute_cosmos(string memory msg) external returns (bool dummy);

    // @args
    // - `allow_failure`: if `true`, the transaction will not be reverted even if the execution fails.
    // - `callback_id`: the callback id to be called after the execution. `0` means no callback.
    struct Options {
        bool allow_failure;
        uint64 callback_id;
    }

    // record a cosmos message to be executed after the current message execution.
    //
    // `msg` format (json string):
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
    // `callback` function signature in the caller contract (see ICosmosCallback.sol):
    // - function callback(uint64 callback_id, bool success) external;
    function execute_cosmos_with_options(
        string memory msg,
        Options memory options
    ) external returns (bool dummy);

    // query a whitelisted cosmos querys.
    //
    // example)
    // path: "/connect.oracle.v2.Query/GetPrices"
    // req: {
    //    "currency_pair_ids": ["BITCOIN/USD", "ETHEREUM/USD"]
    // }
    //
    // res: {
    //    "prices": [
    //        {
    //        "price": {
    //            "price": "5796264752",
    //            "block_timestamp": "2024-08-16T04:18:25.372878802Z",
    //            "block_height": "4231677"
    //        },
    //        "nonce": "4230787",
    //        "decimals": "5",
    //        "id": "2"
    //        }
    //    ]
    // }
    function query_cosmos(
        string memory path,
        string memory req
    ) external returns (string memory result);
}
