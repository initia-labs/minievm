// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

contract BenchHeavyState {
    mapping(uint256 => uint256) public sharedState;
    mapping(address => mapping(uint256 => uint256)) public localState;
    mapping(address => uint256) public senderNonce;
    uint256 public totalCalls;

    function writeMixed(uint256 sharedCount, uint256 localCount) external {
        uint256 nonce = senderNonce[msg.sender];
        senderNonce[msg.sender] = nonce + 1;

        // Each call writes to unique keys so the state tree grows continuously.
        // This creates IAVL rebalancing pressure that MemIAVL handles more efficiently.
        uint256 sharedBase = nonce * sharedCount;
        for (uint256 i = 0; i < sharedCount; i++) {
            sharedState[sharedBase + i] = block.number;
        }
        uint256 localBase = nonce * localCount;
        for (uint256 i = 0; i < localCount; i++) {
            localState[msg.sender][localBase + i] = block.number;
        }
        totalCalls++;
    }
}
