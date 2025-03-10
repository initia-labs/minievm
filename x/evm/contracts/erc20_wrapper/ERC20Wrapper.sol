// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import "../strings/Strings.sol";
import "../i_cosmos/ICosmos.sol";
import "../erc20_factory/ERC20Factory.sol";
import "../i_erc20/IERC20.sol";
import "../ownable/Ownable.sol";
import "../erc20_acl/ERC20ACL.sol";
import "../i_ibc_async_callback/IIBCAsyncCallback.sol";
import "../i_jsonutils/IJSONUtils.sol";
import {ERC165, IERC165} from "../erc165/ERC165.sol";

contract ERC20Wrapper is Ownable, ERC165, IIBCAsyncCallback, ERC20ACL {
    struct IbcCallBack {
        address sender;
        address originToken;
        uint wrappedAmt;
    }

    uint8 constant WRAPPED_DECIMAL = 6;
    string constant NAME_PREFIX = "Wrapped";
    string constant SYMBOL_PREFIX = "W";
    uint64 callBackId = 0;
    ERC20Factory public factory;
    mapping(address => address) public wrappedTokens; // origin -> wrapped
    mapping(uint64 => IbcCallBack) private ibcCallBack; // id -> CallBackInfo

    constructor(address erc20Factory) {
        factory = ERC20Factory(erc20Factory);
    }

    // This function can only be called by the chain at upgrade.
    function setFactory(address newFactory) external onlyChain {
        require(newFactory != address(0), "invalid factory address");
        factory = ERC20Factory(newFactory);
    }

    modifier onlyContract() {
        require(
            msg.sender == address(this),
            "only the contract itself can call this function"
        );
        _;
    }

    /**
     * @notice This function wraps the tokens and transfer the tokens by ibc transfer
     * @dev This function requires sender approve to this contract to transfer the tokens.
     */
    function wrap(
        string memory channel,
        address token,
        string memory receiver,
        uint amount,
        uint timeout
    ) public {
        wrap(channel, token, receiver, amount, timeout, "{}");
    }

    function wrap(
        string memory channel,
        address token,
        string memory receiver,
        uint amount,
        uint timeout,
        string memory memo
    ) public {
        wrap(channel, token, receiver, amount, timeout, memo, 250_000);
    }

    function wrap(
        string memory channel,
        address token,
        string memory receiver,
        uint amount,
        uint timeout,
        string memory memo,
        uint64 gas_limit
    ) public {
        _ensureWrappedTokenExists(token);

        // lock origin token
        IERC20(token).transferFrom(msg.sender, address(this), amount);
        uint wrappedAmt = _convertDecimal(
            amount,
            IERC20(token).decimals(),
            WRAPPED_DECIMAL
        );
        // mint wrapped token
        ERC20(wrappedTokens[token]).mint(address(this), wrappedAmt);

        callBackId += 1;

        // store the callback data
        ibcCallBack[callBackId] = IbcCallBack({
            sender: msg.sender,
            originToken: token,
            wrappedAmt: wrappedAmt
        });

        string memory message = _ibc_transfer(
            channel,
            wrappedTokens[token],
            wrappedAmt,
            timeout,
            receiver,
            memo
        );

        // do ibc transfer wrapped token
        COSMOS_CONTRACT.execute_cosmos(message, gas_limit);
    }

    /**
     * @notice This function is executed as an IBC hook to unwrap the wrapped tokens.
     * @dev This function is used by a hook and requires sender approve to this contract to burn wrapped tokens.
     */
    function unwrap(
        address originToken,
        address receiver,
        uint wrappedAmt
    ) public {
        address wrappedToken = wrappedTokens[originToken];
        require(wrappedToken != address(0), "wrapped token doesn't exist");
        _unwrap(wrappedToken, originToken, msg.sender, receiver, wrappedAmt);
    }

    function unwrap(address originToken, address receiver) public {
        address wrappedToken = wrappedTokens[originToken];
        require(wrappedToken != address(0), "wrapped token doesn't exist");
        uint wrappedAmt = ERC20(wrappedToken).balanceOf(msg.sender);
        _unwrap(wrappedToken, originToken, msg.sender, receiver, wrappedAmt);
    }

    function ibc_ack(uint64 callback_id, bool success) external onlyContract {
        if (success) {
            return;
        }

        _handleFailedIbcTransfer(callback_id);
    }

    function ibc_timeout(uint64 callback_id) external onlyContract {
        _handleFailedIbcTransfer(callback_id);
    }

    // internal functions //
    function _unwrap(
        address wrappedToken,
        address originToken,
        address sender,
        address receiver,
        uint wrappedAmt
    ) internal {
        // burn wrapped token
        ERC20(wrappedToken).burnFrom(sender, wrappedAmt);

        // unlock origin token and transfer to receiver
        uint amount = _convertDecimal(
            wrappedAmt,
            WRAPPED_DECIMAL,
            IERC20(originToken).decimals()
        );

        ERC20(originToken).transfer(receiver, amount);
    }

    function _handleFailedIbcTransfer(uint64 callback_id) internal {
        IbcCallBack memory callback = ibcCallBack[callback_id];
        address wrappedToken = wrappedTokens[callback.originToken];
        require(wrappedToken != address(0), "wrapped token doesn't exist");

        // The wrapped token has already been sent, burn it
        ERC20(wrappedToken).burn(callback.wrappedAmt);

        // unlock origin token and transfer to receiver
        uint amount = _convertDecimal(
            callback.wrappedAmt,
            WRAPPED_DECIMAL,
            IERC20(callback.originToken).decimals()
        );

        ERC20(callback.originToken).transfer(callback.sender, amount);
    }

    function _ensureWrappedTokenExists(address token) internal {
        if (wrappedTokens[token] == address(0)) {
            address wrappedToken = factory.createERC20(
                string.concat(NAME_PREFIX, IERC20(token).name()),
                string.concat(SYMBOL_PREFIX, IERC20(token).symbol()),
                WRAPPED_DECIMAL
            );
            wrappedTokens[token] = wrappedToken;
        }
    }

    function _convertDecimal(
        uint amount,
        uint8 decimal,
        uint8 newDecimal
    ) internal pure returns (uint convertedAmount) {
        if (decimal > newDecimal) {
            uint factor = 10 ** uint(decimal - newDecimal);
            require(amount % factor == 0, "dust amount should be zero");
            convertedAmount = amount / factor;
        } else if (decimal < newDecimal) {
            uint factor = 10 ** uint(newDecimal - decimal);
            convertedAmount = amount * factor;
        } else {
            convertedAmount = amount;
        }

        require(convertedAmount != 0, "converted amount is zero");
    }

    function _ibc_transfer(
        string memory channel,
        address token,
        uint amount,
        uint timeout,
        string memory receiver,
        string memory memo
    ) internal view returns (string memory message) {
        // Construct the memo with the async callback
        string memory callback_memo = string(
            abi.encodePacked(
                '{"evm": {"async_callback": {"id": ',
                Strings.toString(callBackId),
                ',"contract_address":"',
                Strings.toHexString(address(this)),
                '"}}}'
            )
        );

        string memory merged_memo = JSONUTILS_CONTRACT.merge_json(
            memo,
            callback_memo
        );

        // Construct the IBC transfer message
        message = string(
            abi.encodePacked(
                '{"@type": "/ibc.applications.transfer.v1.MsgTransfer",',
                '"source_port": "transfer",',
                '"source_channel": ',
                JSONUTILS_CONTRACT.stringify_json(channel),
                ',',
                '"token": { "denom": "',
                COSMOS_CONTRACT.to_denom(token),
                '",',
                '"amount": "',
                Strings.toString(amount),
                '"},',
                '"sender": "',
                COSMOS_CONTRACT.to_cosmos_address(address(this)),
                '",',
                '"receiver": ',
                JSONUTILS_CONTRACT.stringify_json(receiver),
                ',',
                '"timeout_height": {"revision_number": "0","revision_height": "0"},',
                '"timeout_timestamp": "',
                Strings.toString(timeout),
                '",',
                '"memo": ',
                JSONUTILS_CONTRACT.stringify_json(merged_memo),
                "}"
            )
        );
    }
}
