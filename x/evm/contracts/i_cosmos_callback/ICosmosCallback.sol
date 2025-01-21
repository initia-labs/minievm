// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

/// @title ICosmosCallback
/// @notice Interface for handling callbacks from Cosmos messages
/// @dev Implement this interface to receive results of Cosmos message execution
interface ICosmosCallback {
    /// @notice Callback function called after Cosmos message execution
    /// @param callback_id Unique identifier for the callback
    /// @param success Indicates if the Cosmos message execution was successful
    function callback(uint64 callback_id, bool success) external;
}
