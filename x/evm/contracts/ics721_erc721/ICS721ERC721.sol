// SPDX-License-Identifier: MIT

pragma solidity ^0.8.20;

import {ERC721} from "../erc721/ERC721.sol";
import {Ownable} from "../ownable/Ownable.sol";
import {ERC721Utils} from "../utils/ERC721Utils.sol";

contract ICS721ERC721 is ERC721, Ownable {

    mapping(uint256 => string) private tokenUris;
    mapping(uint256 => string) private tokenOriginIds;
    constructor(string memory name_, string memory symbol_) ERC721(name_, symbol_) Ownable() {}

    function burn(uint256 tokenId) public {
        address owner = _requireOwned(tokenId);
        if (!_isAuthorized(owner, msg.sender, tokenId)) {
            revert ERC721InsufficientApproval(_msgSender(), tokenId);
        }

        _burn(tokenId);
    }

    function mint(address receiver, uint256 tokenId, string memory _tokenUri) public onlyOwner {
        mint(receiver, tokenId, _tokenUri, "");
    }

    function mint(address receiver, uint256 tokenId, string memory _tokenUri, string memory _tokenOriginId) public onlyOwner {
        _safeMint(receiver, tokenId);
        tokenUris[tokenId] = _tokenUri;
        tokenOriginIds[tokenId] = _tokenOriginId;
    }

    function tokenURI(uint256 tokenId) public view override returns (string memory) {
        return tokenUris[tokenId];
    }

    function tokenOriginId(uint256 tokenId) public view returns (string memory) {
        return tokenOriginIds[tokenId];
    }

    function safeTransferFrom(address from, address to, uint256 tokenId, bytes memory data) public override {
        transferFrom(from, to, tokenId);
        ERC721Utils.checkOnERC721Received(_msgSender(), from, to, tokenId, data);
    }
}