// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

interface ICosmosCallback {
    function callback(uint64 callback_id, bool success) external;
}
