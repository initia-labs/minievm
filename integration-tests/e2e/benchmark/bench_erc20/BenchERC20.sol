// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

/// @title BenchERC20. A minimal ERC20 for benchmarking transfer throughput.
/// Open mint allows any account to self-fund. Each transfer writes 2 storage
/// slots (sender balance, recipient balance), the standard ERC20 workload.
contract BenchERC20 {
    mapping(address => uint256) public balanceOf;
    uint256 public totalSupply;

    event Transfer(address indexed from, address indexed to, uint256 value);

    function mint(address to, uint256 amount) external {
        balanceOf[to] += amount;
        totalSupply += amount;
        emit Transfer(address(0), to, amount);
    }

    function transfer(address to, uint256 amount) external returns (bool) {
        require(balanceOf[msg.sender] >= amount, "insufficient balance");
        balanceOf[msg.sender] -= amount;
        balanceOf[to] += amount;
        emit Transfer(msg.sender, to, amount);
        return true;
    }
}
