// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "../util/Strings.sol";
import "../i_cosmos/ICosmos.sol";
import "../erc20_factory/ERC20Factory.sol";
import "../i_erc20/IERC20.sol";
import "../ownable/Ownable.sol";
import "../erc20_acl/ERC20ACL.sol";
import "../i_ibc_async_callback/IIBCAsyncCallback.sol";
import {ERC165, IERC165} from "../erc165/ERC165.sol";

contract ERC20Wrapper is Ownable, ERC165, IIBCAsyncCallback {
    struct IbcCallBack {
        address sender;
        address originToken;
        uint wrappedAmt;
    }

    uint8 constant WRAPPED_DECIMAL = 6;
    string constant NAME_PREFIX = "Wrapped";
    string constant SYMBOL_PREFIX = "W";
    uint64 callBackId = 0;
    ERC20Factory public immutable factory;
    mapping(address => address) public wrappedTokens; // origin -> wrapped
    mapping(uint64 => IbcCallBack) private ibcCallBack; // id -> CallBackInfo

    constructor(address erc20Factory) {
        factory = ERC20Factory(erc20Factory);
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

        // require(false, COSMOS_CONTRACT.to_cosmos_address(address(this)));
        string memory message = _ibc_transfer(
            channel,
            wrappedTokens[token],
            wrappedAmt,
            timeout,
            receiver
        );
        // do ibc transfer wrapped token
        COSMOS_CONTRACT.execute_cosmos(message);
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

        // burn wrapped token
        ERC20(wrappedToken).burnFrom(msg.sender, wrappedAmt);

        // unlock origin token and transfer to receiver
        uint amount = _convertDecimal(
            wrappedAmt,
            WRAPPED_DECIMAL,
            IERC20(originToken).decimals()
        );

        ERC20(originToken).transfer(receiver, amount);
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
        string memory receiver
    ) internal returns (string memory message) {
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
                '"memo": "{\\"evm\\": {\\"async_callback\\": {\\"id\\": ',
                Strings.toString(callBackId),
                ',\\"contract_address\\":\\"',
                Strings.toHexString(address(this)),
                '\\"}}}"}'
            )
        );
    }
}
