// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "../i_erc20/IERC20.sol";
import "../ownable/Ownable.sol";
import "../erc20_registry/ERC20Registry.sol";
import "../erc20_acl/ERC20ACL.sol";
import {ERC165, IERC165} from "../erc165/ERC165.sol";

contract ERC20 is IERC20, Ownable, ERC20Registry, ERC165, ERC20ACL {
    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(
        address indexed owner,
        address indexed spender,
        uint256 value
    );

    mapping(address => uint256) public balanceOf;
    mapping(address => mapping(address => uint256)) public allowance;

    // ERC20 Metadata
    string public name;
    string public symbol;
    uint8 public decimals;
    uint256 public totalSupply;

    // metadataSealed is set to true after metadata is sealed
    bool public metadataSealed;

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

    // for custom erc20s, you should add `register_erc20` modifier to the constructor
    constructor(
        string memory _name,
        string memory _symbol,
        uint8 _decimals,
        bool _metadataSealed
    ) {
        name = _name;
        symbol = _symbol;
        decimals = _decimals;
        metadataSealed = _metadataSealed;
    }

    function _transfer(
        address sender,
        address recipient,
        uint256 amount
    ) internal register_erc20_store(recipient) {
        require(
            balanceOf[sender] >= amount,
            "ERC20: transfer amount exceeds balance"
        );
        balanceOf[sender] -= amount;
        balanceOf[recipient] += amount;
        emit Transfer(sender, recipient, amount);
    }

    function _mint(
        address to,
        uint256 amount
    ) internal register_erc20_store(to) {
        balanceOf[to] += amount;
        totalSupply += amount;
        emit Transfer(address(0), to, amount);
    }

    function _burn(address from, uint256 amount) internal {
        require(
            balanceOf[from] >= amount,
            "ERC20: burn amount exceeds balance"
        );
        balanceOf[from] -= amount;
        totalSupply -= amount;
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

    function sudoMint(address to, uint256 amount) external onlyChain {
        _mint(to, amount);
    }

    function sudoBurn(address from, uint256 amount) external onlyChain {
        _burn(from, amount);
    }

    //
    // ERC20 Metadata onetime setters only for authority(gov)
    //

    event MetadataUpdated(string name, string symbol, uint8 decimals);

    /// @notice Allows one-time update of token metadata by authority
    /// @dev Only callable when metadata is not sealed and by authority
    /// @param _name New token name
    /// @param _symbol New token symbol
    /// @param _decimals New decimal places
    function updateMetadata(
        string memory _name,
        string memory _symbol,
        uint8 _decimals
    ) external onlyAuthority {
        require(!metadataSealed, "ERC20: metadata sealed");
        require(bytes(_name).length > 0, "ERC20: empty name");
        require(bytes(_symbol).length > 0, "ERC20: empty symbol");
        require(_decimals <= 18, "ERC20: invalid decimals");

        // Update all fields at once to save gas
        (name, symbol, decimals) = (_name, _symbol, _decimals);

        // seal the metadata to prevent further updates
        metadataSealed = true;

        emit MetadataUpdated(_name, _symbol, _decimals);
    }
}
