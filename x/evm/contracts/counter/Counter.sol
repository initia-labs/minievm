// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import "../i_ibc_async_callback/IIBCAsyncCallback.sol";
import "../i_cosmos/ICosmos.sol";
import "../strings/Strings.sol";

contract Counter is IIBCAsyncCallback {
    uint256 public count;

    event increased(uint256 oldCount, uint256 newCount);
    event execute_reverted(bool reverted);
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
    ) external returns (string memory result) {
        return COSMOS_CONTRACT.query_cosmos(path, req);
    }

    function execute_cosmos(
        string memory exec_msg,
        bool try_catch
    ) external {
        if (try_catch) {
            try COSMOS_CONTRACT.execute_cosmos(exec_msg) {
                emit execute_reverted(false);
            } catch {
                emit execute_reverted(true);
            }
        } else {
            COSMOS_CONTRACT.execute_cosmos(exec_msg);
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

        COSMOS_CONTRACT.execute_cosmos(_recursive(n));

        // to test branching
        COSMOS_CONTRACT.execute_cosmos(_recursive(n));
    }

    function _recursive(uint64 n) internal returns (string memory message) {
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
}
