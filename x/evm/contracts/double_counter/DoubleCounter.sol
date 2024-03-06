// SPDX-License-Identifier: MIT 
pragma solidity ^0.8.19;

contract DoubleCounter {
    uint256 public count;

    function increase() external {
        count+=2;
    }
}
