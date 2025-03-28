// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import "../i_cosmos/ICosmos.sol";

interface ITest {
    function disable() external;
    function execute_cosmos(string memory msg, uint64 gas_limit) external;
}

contract Test is ITest {
    function disable() external {
        _disable_execute_cosmos();
    }

    function _disable_execute_cosmos() internal {
        COSMOS_CONTRACT.disable_execute_cosmos();
    }

    function execute_cosmos(string memory exec_msg, uint64 gas_limit) external {
        COSMOS_CONTRACT.execute_cosmos(exec_msg, gas_limit);
    }
}
