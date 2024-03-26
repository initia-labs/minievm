// SPDX-License-Identifier: MIT
pragma solidity >=0.8.24;

import "../erc20/ERC20.sol";

contract Factory {
    event ERC20Created(address tokenAddress);

    function deployNewERC20(
        string calldata name,
        string calldata symbol,
        uint8 decimals
    ) external returns (address) {
        ERC20 t = new ERC20(name, symbol, decimals);
        t.transferOwnership(msg.sender);

        emit ERC20Created(address(t));

        return address(t);
    }
}
