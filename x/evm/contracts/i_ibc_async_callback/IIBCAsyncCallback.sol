// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

interface IIBCAsyncCallback {
    function ibc_ack(uint64 callback_id, bool success) external;
    function ibc_timeout(uint64 callback_id) external;
}
