// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "../i_ibc_async_callback/IIBCAsyncCallback.sol";
import "../i_cosmos/ICosmos.sol";

contract Counter is IIBCAsyncCallback {
    uint256 public count;

    event increased(uint256 oldCount, uint256 newCount);

    constructor() payable {}

    function increase() external payable {
        count++;

        emit increased(count - 1, count);
    }

    function ibc_ack(uint64 callback_id, bool success) external {
        if (success) {
            count += callback_id;
        } else {
            count++;
        }
    }

    function ibc_timeout(uint64 callback_id) external {
        count += callback_id;
    }

    function query_cosmos(
        string memory path,
        string memory req
    ) external returns (string memory result) {
        return COSMOS_CONTRACT.query_cosmos(path, req);
    }
}
