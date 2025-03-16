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
        address remoteToken;
        uint remoteAmount;
        uint8 remoteDecimals;
        bool burnRemote;
    }

    uint8 constant REMOTE_DECIMALS = 6;
    uint8 constant LOCAL_DECIMALS = 18;
    string constant NAME_PREFIX = "Wrapped";
    string constant SYMBOL_PREFIX = "W";
    uint64 callbackId = 0;
    ERC20Factory public factory;
    mapping(address => address) public remoteTokens; // localToken -> remoteToken
    mapping(address => uint8) public remoteDecimals; // remoteToken -> remoteDecimals
    mapping(uint64 => IbcCallBack) private ibcCallBack; // id -> CallBackInfo
    mapping(address => mapping(uint8 => address)) public localTokens; // remoteToken -> decimals -> localToken

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
     * @notice This function wraps the remote tokens to 18 decimals local token
     * @dev This function requires sender approve to this contract to transfer the tokens.
     */
    function wrapRemote(
        address receiver,
        string memory remoteDenom,
        uint8 _remoteDecimals
    ) public {
        address remoteToken = COSMOS_CONTRACT.to_erc20(remoteDenom);
        uint remoteAmount = ERC20(remoteToken).balanceOf(msg.sender);
        wrapRemote(receiver, remoteDenom, remoteAmount, _remoteDecimals);
    }

    /**
     * @notice This function wraps the remote tokens to 18 decimals local token
     * @dev This function requires sender approve to this contract to transfer the tokens.
     */
    function wrapRemote(
        address receiver,
        string memory remoteDenom,
        uint remoteAmount,
        uint8 _remoteDecimals
    ) public {
        address remoteToken = COSMOS_CONTRACT.to_erc20(remoteDenom);

        // if there is no local token for the remote token and decimals, create a new one
        _ensureLocalTokenExists(remoteToken, _remoteDecimals);

        address localToken = localTokens[remoteToken][_remoteDecimals];

        // lock received token
        IERC20(remoteToken).transferFrom(
            msg.sender,
            address(this),
            remoteAmount
        );

        // convert decimal
        uint localAmount = _convertDecimal(
            remoteAmount,
            _remoteDecimals,
            LOCAL_DECIMALS
        );

        // mint wrapped token to receiver
        ERC20(localToken).mint(receiver, localAmount);
    }

    /**
     * @notice This function unwraps the remote tokens from 18 decimals local token
     * @dev This function requires sender approve to this contract to transfer the tokens.
     */
    function unwrapRemote(
        string memory channel,
        address localToken,
        string memory receiver,
        uint localAmount,
        uint timeout
    ) public {
        unwrapRemote(channel, localToken, receiver, localAmount, timeout, "{}");
    }

    /**
     * @notice This function unwraps the remote tokens from 18 decimals local token
     * @dev This function requires sender approve to this contract to transfer the tokens.
     */
    function unwrapRemote(
        string memory channel,
        address localToken,
        string memory receiver,
        uint localAmount,
        uint timeout,
        string memory memo
    ) public {
        unwrapRemote(
            channel,
            localToken,
            receiver,
            localAmount,
            timeout,
            memo,
            250_000
        );
    }

    /**
     * @notice This function unwraps the remote tokens from 18 decimals local token
     * @dev This function requires sender approve to this contract to transfer the tokens.
     */
    function unwrapRemote(
        string memory channel,
        address localToken,
        string memory receiver,
        uint localAmount,
        uint timeout,
        string memory memo,
        uint64 gasLimit
    ) public {
        _unwrapRemote(
            channel,
            localToken,
            receiver,
            localAmount,
            timeout,
            memo,
            gasLimit
        );
    }

    /**
     * @notice This function wraps the local tokens to 6 decimals remote token
     * @dev This function requires sender approve to this contract to transfer the tokens.
     */
    function wrapLocal(
        string memory channel,
        address localToken,
        string memory receiver,
        uint localAmount,
        uint timeout
    ) public {
        wrapLocal(channel, localToken, receiver, localAmount, timeout, "{}");
    }

    /**
     * @notice This function wraps the local tokens to 6 decimals remote token
     * @dev This function requires sender approve to this contract to transfer the tokens.
     */
    function wrapLocal(
        string memory channel,
        address localToken,
        string memory receiver,
        uint localAmount,
        uint timeout,
        string memory memo
    ) public {
        wrapLocal(
            channel,
            localToken,
            receiver,
            localAmount,
            timeout,
            memo,
            250_000
        );
    }

    /**
     * @notice This function wraps the local tokens to 6 decimals remote token
     * @dev This function requires sender approve to this contract to transfer the tokens.
     */
    function wrapLocal(
        string memory channel,
        address localToken,
        string memory receiver,
        uint localAmount,
        uint timeout,
        string memory memo,
        uint64 gasLimit
    ) public {
        _ensureRemoteTokenExists(localToken);

        // lock origin token
        IERC20(localToken).transferFrom(msg.sender, address(this), localAmount);
        uint remoteAmount = _convertDecimal(
            localAmount,
            IERC20(localToken).decimals(),
            REMOTE_DECIMALS
        );

        // mint wrapped token
        address remoteToken = remoteTokens[localToken];
        uint8 _remoteDecimals = IERC20(remoteToken).decimals();
        ERC20(remoteToken).mint(address(this), remoteAmount);

        callbackId += 1;

        // store the callback data
        ibcCallBack[callbackId] = IbcCallBack({
            sender: msg.sender,
            remoteToken: remoteToken,
            remoteAmount: remoteAmount,
            remoteDecimals: _remoteDecimals,
            burnRemote: true
        });

        string memory message = _ibc_transfer(
            channel,
            remoteToken,
            remoteAmount,
            timeout,
            receiver,
            memo
        );

        // do ibc transfer wrapped token
        COSMOS_CONTRACT.execute_cosmos(message, gasLimit);
    }

    /**
     * @notice This function is executed as an IBC hook to unwrap the wrapped tokens.
     * @dev This function is used by a hook and requires sender approve to this contract to burn wrapped tokens.
     */
    function unwrapLocal(address receiver, string memory remoteDenom) public {
        address remoteToken = COSMOS_CONTRACT.to_erc20(remoteDenom);
        address localToken = localTokens[remoteToken][REMOTE_DECIMALS];
        require(localToken != address(0), "local token doesn't exist");
        uint remoteAmount = ERC20(remoteToken).balanceOf(msg.sender);
        _unwrapLocal(
            remoteToken,
            localToken,
            msg.sender,
            receiver,
            remoteAmount
        );
    }

    /**
     * @notice This function is executed as an IBC hook to unwrap the wrapped tokens.
     * @dev This function is used by a hook and requires sender approve to this contract to burn wrapped tokens.
     */
    function unwrapLocal(
        address receiver,
        string memory remoteDenom,
        uint remoteAmount
    ) public {
        address remoteToken = COSMOS_CONTRACT.to_erc20(remoteDenom);
        address localToken = localTokens[remoteToken][REMOTE_DECIMALS];
        require(localToken != address(0), "local token doesn't exist");
        _unwrapLocal(
            remoteToken,
            localToken,
            msg.sender,
            receiver,
            remoteAmount
        );
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
    function _unwrapRemote(
        string memory channel,
        address localToken,
        string memory receiver,
        uint localAmount,
        uint timeout,
        string memory memo,
        uint64 gasLimit
    ) internal {
        // burn wrapped token
        ERC20(localToken).burnFrom(msg.sender, localAmount);

        // unlock origin token and transfer to receiver
        address remoteToken = remoteTokens[localToken];
        require(remoteToken != address(0), "remote token doesn't exist");

        // unlock origin token and transfer to receiver
        uint8 _remoteDecimals = remoteDecimals[remoteToken];
        require(
            localTokens[remoteToken][_remoteDecimals] != address(0),
            "local token doesn't exist"
        );

        uint remoteAmount = _convertDecimal(
            localAmount,
            LOCAL_DECIMALS,
            _remoteDecimals
        );

        callbackId += 1;

        // store the callback data
        ibcCallBack[callbackId] = IbcCallBack({
            sender: msg.sender,
            remoteToken: remoteToken,
            remoteAmount: remoteAmount,
            remoteDecimals: _remoteDecimals,
            burnRemote: false
        });

        string memory message = _ibc_transfer(
            channel,
            remoteToken,
            remoteAmount,
            timeout,
            receiver,
            memo
        );

        // do ibc transfer wrapped token
        COSMOS_CONTRACT.execute_cosmos(message, gasLimit);
    }

    function _unwrapLocal(
        address remoteToken,
        address localToken,
        address sender,
        address receiver,
        uint remoteAmount
    ) internal {
        // burn wrapped token
        ERC20(remoteToken).burnFrom(sender, remoteAmount);

        // unlock origin token and transfer to receiver
        uint localAmount = _convertDecimal(
            remoteAmount,
            REMOTE_DECIMALS,
            IERC20(localToken).decimals()
        );

        ERC20(localToken).transfer(receiver, localAmount);
    }

    function _handleFailedIbcTransfer(uint64 callback_id) internal {
        IbcCallBack memory callback = ibcCallBack[callback_id];
        address localToken = localTokens[callback.remoteToken][
            callback.remoteDecimals
        ];
        require(localToken != address(0), "local token doesn't exist");

        // compute the local amount
        uint localAmount = _convertDecimal(
            callback.remoteAmount,
            callback.remoteDecimals,
            IERC20(localToken).decimals()
        );

        // The wrapped token has already been received to this contract.
        if (callback.burnRemote) {
            ERC20(callback.remoteToken).burn(callback.remoteAmount);

            // unlock local token
            ERC20(localToken).transfer(callback.sender, localAmount);
        } else {
            // mint local token
            ERC20(localToken).mint(callback.sender, localAmount);
        }

        delete ibcCallBack[callback_id];
    }

    function _ensureRemoteTokenExists(address localToken) internal {
        if (remoteTokens[localToken] == address(0)) {
            address remoteToken = factory.createERC20(
                string.concat(NAME_PREFIX, IERC20(localToken).name()),
                string.concat(SYMBOL_PREFIX, IERC20(localToken).symbol()),
                REMOTE_DECIMALS
            );
            remoteTokens[localToken] = remoteToken;
            remoteDecimals[remoteToken] = REMOTE_DECIMALS;
            localTokens[remoteToken][REMOTE_DECIMALS] = localToken;
        }
    }

    function _ensureLocalTokenExists(
        address remoteToken,
        uint8 _remoteDecimals
    ) internal {
        if (localTokens[remoteToken][_remoteDecimals] == address(0)) {
            address localToken = factory.createERC20(
                string.concat(NAME_PREFIX, IERC20(remoteToken).name()),
                string.concat(SYMBOL_PREFIX, IERC20(remoteToken).symbol()),
                LOCAL_DECIMALS
            );
            localTokens[remoteToken][_remoteDecimals] = localToken;
            remoteTokens[localToken] = remoteToken;
            remoteDecimals[remoteToken] = _remoteDecimals;
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
                Strings.toString(callbackId),
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
                ",",
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
                ",",
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
