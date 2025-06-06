// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import "../i_ibc_async_callback/IIBCAsyncCallback.sol";
import "../i_cosmos/ICosmos.sol";
import "../strings/Strings.sol";
import "../test/Test.sol";

contract Counter is IIBCAsyncCallback {
    uint256 public count;

    event increased(uint256 oldCount, uint256 newCount);
    event callback_received(uint64 callback_id, bool success);
    event recursive_called(uint64 n);

    constructor() payable {}

    function increase_for_fuzz(uint64 num) public {
        if (num == 0) {
            return;
        }

        increase();
        increase_for_fuzz(num - 1);
    }

    function increase() public payable {
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
    ) external view returns (string memory result) {
        return COSMOS_CONTRACT.query_cosmos(path, req);
    }

    function execute_cosmos(
        string memory exec_msg, 
        uint64 gas_limit,
        bool call_revert
    ) external {
        COSMOS_CONTRACT.execute_cosmos(exec_msg, gas_limit);

        if (call_revert) {
            revert("revert reason dummy value for test");
        }
    }

    function execute_cosmos_with_options(
        string memory exec_msg,
        uint64 gas_limit,
        bool allow_failure,
        uint64 callback_id
    ) external {
        COSMOS_CONTRACT.execute_cosmos_with_options(
            exec_msg,
            gas_limit,
            ICosmos.Options(allow_failure, callback_id)
        );
    }

    function disable_and_execute(
        string memory exec_msg, 
        uint64 gas_limit
    ) external {
        COSMOS_CONTRACT.disable_execute_cosmos();
        COSMOS_CONTRACT.execute_cosmos(exec_msg, gas_limit);
    }

     function disable_and_execute_in_child(
        address test_addr,
        string memory exec_msg, 
        uint64 gas_limit
    ) external {
        COSMOS_CONTRACT.disable_execute_cosmos();

        // execute_cosmos should be failed because the child context is affected
        ITest(test_addr).execute_cosmos(exec_msg, gas_limit);
    }

    function disable_and_execute_in_parent(
        address test_addr,
        string memory exec_msg, 
        uint64 gas_limit
    ) external {
        // execute other contract which is disabling execute cosmos
        ITest(test_addr).disable();

        // execute_cosmos should be successful because the parent context is not affected
        COSMOS_CONTRACT.execute_cosmos(exec_msg, gas_limit);
    }

    function callback(uint64 callback_id, bool success) external {
        emit callback_received(callback_id, success);

        if (callback_id == 7) {
            revert("revert reason dummy value for test");
        }
    }

    function get_blockhash(uint64 n) external view returns (bytes32) {
        return blockhash(n);
    }

    function recursive(uint64 n) public {
        emit recursive_called(n);

        if (n == 0) {
            return;
        }

        
        COSMOS_CONTRACT.execute_cosmos(_recursive(n), uint64(n * (2**(n+1)-1) * (30_000 + 10_000 * n)));

        // to test branching
        COSMOS_CONTRACT.execute_cosmos(_recursive(n), uint64(n * (2**(n+1)-1) * (30_000 + 10_000 * n)));
    }

    function _recursive(uint64 n) internal view returns (string memory message) {
        message = string(
            abi.encodePacked(
                '{"@type": "/minievm.evm.v1.MsgCall",',
                '"sender": "',
                COSMOS_CONTRACT.to_cosmos_address(address(this)),
                '",',
                '"contract_addr": "',
                Strings.toHexString(address(this)),
                '",',
                '"input": "',
                Strings.toHexString(
                    abi.encodePacked(this.recursive.selector, abi.encode(n - 1))
                ),
                '",',
                '"value": "0",',
                '"access_list": []}'
            )
        );
    }

    function recursive_revert(uint64 n) public {
        emit recursive_called(n);

        if (n == 0) {
            return;
        }

        try this.nested_recursive_revert(n) {} catch {}
    }

    function nested_recursive_revert(uint64 n) external {
        COSMOS_CONTRACT.execute_cosmos(_recursive(n), uint64(n * (2**(n+1)-1) * (30_000 + 10_000 * n)));

        revert();
    }    
}
