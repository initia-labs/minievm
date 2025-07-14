// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import "../i_erc20/IERC20.sol";
import "../ownable/Ownable.sol";
import "../erc20_registry/ERC20Registry.sol";
import "../erc20_acl/ERC20ACL.sol";
import {ERC165, IERC165} from "../erc165/ERC165.sol";

contract InfiniteLoopERC20 is IERC20, Ownable, ERC20Registry, ERC165, ERC20ACL {
    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(
        address indexed owner,
        address indexed spender,
        uint256 value
    );

    mapping(address => uint256) private _balanceOf;
    mapping(address => mapping(address => uint256)) public allowance;
    string private _name;
    string private _symbol;
    uint8 private _decimals;
    uint256 private _totalSupply;

    function balanceOf(address account) public view returns (uint256) {
        // infinite loop for testing
        uint256 i = 0;
        while (true) {
            i++;
        }

        return _balanceOf[account];
    }

    function name() public view returns (string memory) {
        // infinite loop for testing
        uint256 i = 0;
        while (true) {
            i++;
        }

        return _name;
    }
    
    function decimals() public view returns (uint8) {
        // infinite loop for testing
        uint256 i = 0;
        while (true) {
            i++;
        }

        return _decimals;
    }

    function totalSupply() public view returns (uint256) {
        // infinite loop for testing
        uint256 i = 0;
        while (true) {
            i++;
        }

        return _totalSupply;
    }

    function symbol() public view returns (string memory) {
        // infinite loop for testing
        uint256 i = 0;
        while (true) {
            i++;
        }

        return _symbol;
    }

    /**
     * @dev See {IERC165-supportsInterface}.
     */
    function supportsInterface(
        bytes4 interfaceId
    ) public view virtual override(IERC165, ERC165) returns (bool) {
        return
            interfaceId == type(IERC20).interfaceId ||
            super.supportsInterface(interfaceId);
    }

    constructor(
        string memory __name,
        string memory __symbol,
        uint8 __decimals
    ) register_erc20 {
        _name = __name;
        _symbol = __symbol;
        _decimals = __decimals;
    }

    function _transfer(
        address sender,
        address recipient,
        uint256 amount
    ) internal register_erc20_store(recipient) {
        require(
            _balanceOf[sender] >= amount,
            "ERC20: transfer amount exceeds balance"
        );
        _balanceOf[sender] -= amount;
        _balanceOf[recipient] += amount;
        emit Transfer(sender, recipient, amount);
    }

    function _mint(
        address to,
        uint256 amount
    ) internal register_erc20_store(to) {
        _balanceOf[to] += amount;
        _totalSupply += amount;
        emit Transfer(address(0), to, amount);
    }

    function _burn(address from, uint256 amount) internal {
        require(
            _balanceOf[from] >= amount,
            "ERC20: burn amount exceeds balance"
        );
        _balanceOf[from] -= amount;
        _totalSupply -= amount;
        emit Transfer(from, address(0), amount);
    }

    function transfer(
        address recipient,
        uint256 amount
    ) external transferable(recipient) returns (bool) {
        _transfer(msg.sender, recipient, amount);
        return true;
    }

    function approve(address spender, uint256 amount) external returns (bool) {
        allowance[msg.sender][spender] = amount;
        emit Approval(msg.sender, spender, amount);
        return true;
    }

    function transferFrom(
        address sender,
        address recipient,
        uint256 amount
    ) external transferable(recipient) returns (bool) {
        require(
            allowance[sender][msg.sender] >= amount,
            "ERC20: transfer amount exceeds allowance"
        );
        allowance[sender][msg.sender] -= amount;
        _transfer(sender, recipient, amount);
        return true;
    }

    function mint(address to, uint256 amount) external mintable(to) onlyOwner {
        _mint(to, amount);
    }

    function burn(uint256 amount) external burnable(msg.sender) {
        _burn(msg.sender, amount);
    }

    function burnFrom(
        address from,
        uint256 amount
    ) external burnable(from) returns (bool) {
        require(
            allowance[from][msg.sender] >= amount,
            "ERC20: burn amount exceeds allowance"
        );
        allowance[from][msg.sender] -= amount;
        _burn(from, amount);
        return true;
    }

    function sudoTransfer(
        address sender,
        address recipient,
        uint256 amount
    ) external onlyChain {
        _transfer(sender, recipient, amount);
    }
}
