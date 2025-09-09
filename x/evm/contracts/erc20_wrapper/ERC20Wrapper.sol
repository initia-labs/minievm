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
    // constant
    uint8 constant REMOTE_DECIMALS = 6;
    uint8 constant LOCAL_DECIMALS = 18;
    string constant NAME_PREFIX = "Wrapped";
    string constant SYMBOL_PREFIX = "W";
    ERC20Factory public factory;
    
    // event
    event Converted(
        string srcDenom,
        string dstDenom,
        uint8 srcDecimal,
        uint8 dstDecimal,
        uint amount
    );

    mapping(address => address) public remoteTokens; // localToken -> remoteToken
    mapping(address => uint8) public remoteDecimals; // localToken -> remoteDecimals
    mapping(address => mapping(uint8 => address)) public localTokens; // remoteToken -> decimals -> localToken

    modifier onlyContract() {
        require(
            msg.sender == address(this),
            "only the contract itself can call this function"
        );
        _;
    }

    constructor(address erc20Factory) {
        factory = ERC20Factory(erc20Factory);
    }

    /**
     * @notice Sets the factory address for creating ERC20 tokens
     * @param newFactory The address of the new factory to set
     * @dev Only the chain can call this function at upgrade
     */
    function setFactory(address newFactory) external onlyChain {
        require(newFactory != address(0), "invalid factory address");
        factory = ERC20Factory(newFactory);
    }

    /**
     * @notice Converts a remote token into a local ERC20 token with 18 decimals
     * @notice The remote token must be approved for spending by this contract before calling
     * @dev This function is a convenience wrapper that uses the sender's full balance
     * @dev For wrapping a specific amount, use the overloaded toLocal function with amount parameter
     * @param receiver The address that will receive the wrapped local tokens
     * @param remoteDenom The denomination identifier of the remote token (e.g. "uatom", "uosmo")
     * @param _remoteDecimals The number of decimal places used by the remote token (typically 6)
     */
    function toLocal(
        address receiver,
        string memory remoteDenom,
        uint8 _remoteDecimals
    ) public {
        // remoteDenom -> remoteToken
        address remoteToken = COSMOS_CONTRACT.to_erc20(remoteDenom);

        // load balance and allowance of remote token
        uint remoteAllowance = ERC20(remoteToken).allowance(msg.sender, address(this));
        uint remoteBalance = ERC20(remoteToken).balanceOf(msg.sender);

        // use min(remoteAllowance, remoteBalance) as the remote amount
        uint remoteAmount = remoteAllowance > remoteBalance ? remoteBalance : remoteAllowance;

        toLocal(receiver, remoteDenom, remoteAmount, _remoteDecimals);
    }

    /**
     * @notice Converts a remote token into a local ERC20 token with 18 decimals
     * @notice The remote token must be approved for spending by this contract before calling
     * @dev This function handles both minting new wrapped tokens and transferring existing ones
     * @param receiver The address that will receive the wrapped local tokens
     * @param remoteDenom The denomination identifier of the remote token
     * @param remoteAmount The amount of remote tokens to wrap
     * @param _remoteDecimals The number of decimal places used by the remote token
     */
    function toLocal(
        address receiver,
        string memory remoteDenom,
        uint remoteAmount,
        uint8 _remoteDecimals
    ) public {
        // remoteDenom -> remoteToken
        address remoteToken = COSMOS_CONTRACT.to_erc20(remoteDenom);

        // ensure the local token exists if not create it
        address localToken = _ensureLocalTokenExists(
            remoteDenom,
            remoteToken,
            _remoteDecimals
        );

        // if the remote amount is 0, do nothing
        if (remoteAmount == 0) {
            return;
        }

        // check if the remote token is owned by this contract
        if (_isOwner(remoteToken)) {
            // burn the remote token from the msg.sender
            ERC20(remoteToken).burnFrom(msg.sender, remoteAmount);
        } else {
            // (lock) transfer the remote token to this contract
            ERC20(remoteToken).transferFrom(
                msg.sender,
                address(this),
                remoteAmount
            );
        }

        uint8 _localDecimals = ERC20(localToken).decimals();
        // convert the remote amount to the local amount
        uint localAmount = _convertDecimal(
            remoteAmount,
            _remoteDecimals,
            _localDecimals
        );

        // check if the local token is owned by this contract
        if (_isOwner(localToken)) {
            // mint the local token to receiver
            ERC20(localToken).mint(receiver, localAmount);
        } else {
            // (unlock) transfer the local token to receiver from this contract
            ERC20(localToken).transfer(receiver, localAmount);
        }

        emit Converted(
            remoteDenom,
            COSMOS_CONTRACT.to_denom(localToken),
            _remoteDecimals,
            LOCAL_DECIMALS,
            remoteAmount
        );
    }

    /**
     * @notice Converts a local token into a remote token with arbitrary decimals
     * @notice If the remote token does not exist, it will be created with 6 decimals
     * @notice The local token must be approved for spending by this contract before calling
     * @dev This function handles both burning existing wrapped tokens and transferring them to the receiver
     * @param receiver The address that will receive the unwrapped remote tokens
     * @param localDenom The denomination identifier of the local token
     * @param localAmount The amount of local tokens to unwrap
     */
    function toRemote(
        address receiver,
        string memory localDenom,
        uint localAmount
    )
        public
        returns (address remoteToken, uint remoteAmount, uint8 _remoteDecimals)
    {
        // localDenom -> localToken
        address localToken = COSMOS_CONTRACT.to_erc20(localDenom);

        // ensure the remote token exists if not create it
        remoteToken = _ensureRemoteTokenExists(localDenom, localToken);
        _remoteDecimals = remoteDecimals[localToken];

        // if the local amount is 0, do nothing
        if (localAmount == 0) {
            return (remoteToken, 0, _remoteDecimals);
        }

        // check if the local token is owned by this contract
        if (_isOwner(localToken)) {
            // burn the local token from the msg.sender
            ERC20(localToken).burnFrom(msg.sender, localAmount);
        } else {
            // (lock) transfer the local token to this contract
            ERC20(localToken).transferFrom(
                msg.sender,
                address(this),
                localAmount
            );
        }

        // convert the local amount to the remote amount
        uint8 _localDecimals = IERC20(localToken).decimals();
        remoteAmount = _convertDecimal(
            localAmount,
            _localDecimals,
            _remoteDecimals
        );

        // check if the remote token is owned by this contract
        if (_isOwner(remoteToken)) {
            // mint the remote token to receiver
            ERC20(remoteToken).mint(receiver, remoteAmount);
        } else {
            // (unlock) transfer the remote token to receiver from this contract
            ERC20(remoteToken).transfer(receiver, remoteAmount);
        }

        emit Converted(
            localDenom,
            COSMOS_CONTRACT.to_denom(remoteToken),
            _localDecimals,
            _remoteDecimals,
            localAmount
        );
        return (remoteToken, remoteAmount, _remoteDecimals);
    }

    /////////////////////////////
    // Convert and OP Withdraw //
    /////////////////////////////

    /// @notice Converts local tokens to remote tokens and initiates an OP withdraw using a specific amount of local tokens
    /// @dev This is a convenience wrapper that automatically converts and withdraws the specified amount of local tokens
    /// @param receiver The destination address that will receive the unwrapped remote tokens
    /// @param localDenom The denomination identifier of the local wrapped token to convert
    /// @param localAmount The amount of local tokens to convert
    function toRemoteAndOPWithdraw(
        string memory receiver,
        string memory localDenom,
        uint localAmount
    ) public {
        toRemoteAndOPWithdraw(receiver, localDenom, localAmount, 250_000);
    }

    /// @notice Converts local tokens to remote tokens and initiates an OP withdraw using a specific amount of local tokens
    /// @dev This is a convenience wrapper that automatically converts and withdraws the specified amount of local tokens
    /// @param receiver The destination address that will receive the unwrapped remote tokens
    /// @param localDenom The denomination identifier of the local wrapped token to convert
    /// @param localAmount The amount of local tokens to convert
    /// @param gasLimit The gas limit for the OP withdraw
    function toRemoteAndOPWithdraw(
        string memory receiver,
        string memory localDenom,
        uint localAmount,
        uint64 gasLimit
    ) public {
        (address remoteToken, uint remoteAmount, ) = toRemote(
            address(this),
            localDenom,
            localAmount
        );

        // if the remote amount is 0, do nothing
        if (remoteAmount == 0) {
            return;
        }

        string memory message = _op_withdraw(
            receiver,
            remoteToken,
            remoteAmount
        );
        COSMOS_CONTRACT.execute_cosmos(message, gasLimit);
    }

    //////////////////////////////
    // Convert and IBC Transfer //
    //////////////////////////////

    struct IbcCallBack {
        address sender;
        address remoteToken;
        uint remoteAmount;
        uint8 remoteDecimals;
    }

    uint64 callbackId = 0;
    mapping(uint64 => IbcCallBack) private ibcCallBack; // id -> CallBackInfo

    /// @notice Converts local tokens to remote tokens and initiates an IBC transfer
    /// @dev This is a convenience wrapper that automatically converts and transfers local tokens.
    /// @param localDenom The denomination identifier of the local wrapped token to convert (e.g. "evm/123...")
    /// @param localAmount The amount of local tokens to convert and transfer
    /// @param channel The IBC channel identifier to use for the transfer (e.g. "channel-0")
    /// @param receiver The destination address that will receive the unwrapped remote tokens on the target chain
    /// @param timeout The Unix timestamp in nanoseconds after which the IBC transfer will timeout and revert
    function toRemoteAndIBCTransfer(
        string memory localDenom,
        uint localAmount,
        // args for IBC transfer
        string memory channel,
        string memory receiver,
        uint timeout
    ) public {
        toRemoteAndIBCTransfer(
            localDenom,
            localAmount,
            channel,
            receiver,
            timeout,
            "{}",
            250_000
        );
    }

    /// @notice Converts local tokens to remote tokens and initiates an IBC transfer
    /// @dev This is a convenience wrapper that automatically converts and transfers local tokens.
    /// @param localDenom The denomination identifier of the local wrapped token to convert (e.g. "evm/123...")
    /// @param localAmount The amount of local tokens to convert and transfer
    /// @param channel The IBC channel identifier to use for the transfer (e.g. "channel-0")
    /// @param receiver The destination address that will receive the unwrapped remote tokens on the target chain
    /// @param timeout The Unix timestamp in nanoseconds after which the IBC transfer will timeout and revert
    /// @param memo Optional memo string to include with the IBC transfer
    function toRemoteAndIBCTransfer(
        string memory localDenom,
        uint localAmount,
        // args for IBC transfer
        string memory channel,
        string memory receiver,
        uint timeout,
        string memory memo
    ) public {
        toRemoteAndIBCTransfer(
            localDenom,
            localAmount,
            channel,
            receiver,
            timeout,
            memo,
            250_000
        );
    }

    /// @notice Converts local tokens to remote tokens and initiates an IBC transfer
    /// @dev This is a convenience wrapper that automatically converts and transfers local tokens.
    /// @param localDenom The denomination identifier of the local wrapped token to convert (e.g. "erc20/0x123...")
    /// @param localAmount The amount of local tokens to convert and transfer
    /// @param channel The IBC channel identifier to use for the transfer (e.g. "channel-0")
    /// @param receiver The destination address that will receive the unwrapped remote tokens on the target chain
    /// @param timeout The Unix timestamp in nanoseconds after which the IBC transfer will timeout and revert
    /// @param memo Optional memo string to include with the IBC transfer
    /// @param gasLimit The gas limit for the IBC transfer
    function toRemoteAndIBCTransfer(
        string memory localDenom,
        uint localAmount,
        // args for IBC transfer
        string memory channel,
        string memory receiver,
        uint timeout,
        string memory memo,
        uint64 gasLimit
    ) public {
        (
            address remoteToken,
            uint remoteAmount,
            uint8 _remoteDecimals
        ) = toRemote(address(this), localDenom, localAmount);

        // if the remote amount is 0, do nothing
        if (remoteAmount == 0) {
            return;
        }

        callbackId += 1;

        // store the callback data
        ibcCallBack[callbackId] = IbcCallBack({
            sender: msg.sender,
            remoteToken: remoteToken,
            remoteAmount: remoteAmount,
            remoteDecimals: _remoteDecimals
        });

        string memory message = _ibc_transfer(
            channel,
            remoteToken,
            remoteAmount,
            timeout,
            receiver,
            memo
        );
        COSMOS_CONTRACT.execute_cosmos(message, gasLimit);
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

    /////////////////////////////
    // External View functions //
    /////////////////////////////
    /**
     * Get the local wrapped contract address
     * @param remoteDenom remote coin denom. it can be not exists in this chain.
     * @param name remote token name
     * @param symbol remote token symbol
     * @param decimal retmoe token decimals
     */
    function getToLocalERC20Address(
        string memory remoteDenom,
        string memory name,
        string memory symbol,
        uint8 decimal
    ) external view returns (address) {
        return
            factory.computeERC20Address(
                address(this),
                string.concat(NAME_PREFIX, name),
                string.concat(SYMBOL_PREFIX, symbol),
                LOCAL_DECIMALS,
                keccak256(abi.encodePacked(remoteDenom, decimal))
            );
    }

    /**
     * Get the remote wrapped contract address
     * @param localDenom local coin denom
     */
    function getToRemoteERC20Address(
        string memory localDenom
    ) external view returns (address) {
        address token = COSMOS_CONTRACT.to_erc20(localDenom);
        return
            factory.computeERC20Address(
                address(this),
                string.concat(NAME_PREFIX, ERC20(token).name()),
                string.concat(SYMBOL_PREFIX, ERC20(token).symbol()),
                REMOTE_DECIMALS,
                keccak256(abi.encodePacked(localDenom, ERC20(token).decimals()))
            );
    }

    ////////////////////////
    // Internal functions //
    ////////////////////////
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

        // The remote token is already held by this contract from the failed IBC transfer.
        // If this contract owns the remote token, burn it directly.
        // No transfer needed since we already have custody of the tokens.
        if (_isOwner(callback.remoteToken)) {
            ERC20(callback.remoteToken).burn(callback.remoteAmount);
        }

        // check if this contract owns the local token to determine whether to mint or transfer
        if (_isOwner(localToken)) {
            ERC20(localToken).mint(callback.sender, localAmount);
        } else {
            ERC20(localToken).transfer(callback.sender, localAmount);
        }

        delete ibcCallBack[callback_id];
    }

    /**
     * @notice Ensures that a remote token exists for the given local token
     * @notice If no remote token exists, creates a new one with 6 decimals (REMOTE_DECIMALS)
     * @param localToken The address of the local token to check/create remote token for
     * @dev Updates remoteTokens, remoteDecimals, and localTokens mappings if a new token is created
     */
    function _ensureRemoteTokenExists(
        string memory localDenom,
        address localToken
    ) internal returns (address remoteToken) {
        remoteToken = remoteTokens[localToken];
        if (remoteToken == address(0)) {
            remoteToken = factory.createERC20WithSalt(
                string.concat(NAME_PREFIX, IERC20(localToken).name()),
                string.concat(SYMBOL_PREFIX, IERC20(localToken).symbol()),
                REMOTE_DECIMALS,
                keccak256(
                    abi.encodePacked(localDenom, IERC20(localToken).decimals())
                )
            );
            remoteTokens[localToken] = remoteToken;
            remoteDecimals[localToken] = REMOTE_DECIMALS;
            localTokens[remoteToken][REMOTE_DECIMALS] = localToken;
        }

        return remoteToken;
    }

    /**
     * @notice Ensures that a local token exists for the given remote token and decimal precision
     * @notice If no local token exists, creates a new one with 18 decimals (LOCAL_DECIMALS)
     * @param remoteToken The address of the remote token to check/create local token for
     * @param _remoteDecimals The number of decimal places used by the remote token
     * @dev Updates localTokens, remoteTokens, and remoteDecimals mappings if a new token is created
     */
    function _ensureLocalTokenExists(
        string memory remoteDenom,
        address remoteToken,
        uint8 _remoteDecimals
    ) internal returns (address localToken) {
        localToken = localTokens[remoteToken][_remoteDecimals];
        if (localToken == address(0)) {
            localToken = factory.createERC20WithSalt(
                string.concat(NAME_PREFIX, IERC20(remoteToken).name()),
                string.concat(SYMBOL_PREFIX, IERC20(remoteToken).symbol()),
                LOCAL_DECIMALS,
                keccak256(abi.encodePacked(remoteDenom, _remoteDecimals))
            );
            localTokens[remoteToken][_remoteDecimals] = localToken;
            remoteTokens[localToken] = remoteToken;
            remoteDecimals[localToken] = _remoteDecimals;
        }

        return localToken;
    }

    // view
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
                '"source_channel": "',
                channel,
                '",',
                '"token": { "denom": "',
                COSMOS_CONTRACT.to_denom(token),
                '",',
                '"amount": "',
                Strings.toString(amount),
                '"},',
                '"sender": "',
                COSMOS_CONTRACT.to_cosmos_address(address(this)),
                '",',
                '"receiver": "',
                receiver,
                '",',
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

    function _op_withdraw(
        string memory receiver,
        address token,
        uint amount
    ) internal view returns (string memory message) {
        // Construct OPbridge withdraw message
        message = string(
            abi.encodePacked(
                '{"@type": "/opinit.opchild.v1.MsgInitiateTokenWithdrawal"',
                ',"amount": { "denom": "',
                COSMOS_CONTRACT.to_denom(token),
                '","amount": "',
                Strings.toString(amount),
                '"},"sender": "',
                COSMOS_CONTRACT.to_cosmos_address(address(this)),
                '","to": "',
                receiver,
                '"}'
            )
        );
    }

    /**
     * @notice Checks if the contract is the owner of the given token
     * @param token The address of the token to check ownership of
     * @return true if the contract is the owner, false otherwise
     * @dev This function is safe to call even if the token does not support ownership
     */
    function _isOwner(address token) internal view returns (bool) {
        try ERC20(token).owner() returns (address owner) {
            return owner == address(this);
        } catch {
            return false;
        }
    }

    // pure
    /**
     * @notice Converts an amount from one decimal precision to another by scaling up or down
     * @param amount The amount to convert
     * @param decimal The original decimal precision (e.g. 6 for USDC)
     * @param newDecimal The desired decimal precision (e.g. 18 for wrapped USDC)
     * @return convertedAmount The amount converted to the new decimal precision
     * @dev When scaling down (decimal > newDecimal), checks for and prevents dust amounts
     * @dev When scaling up (decimal < newDecimal), multiplies by the scaling factor
     * @dev When decimal precisions match, returns original amount unchanged
     */
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
    }
}
