// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// 导入OpenZeppelin合约
import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol"; // 新增导入
import "@openzeppelin/contracts/utils/Counters.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract MyNFT is ERC721URIStorage, Ownable {
    using Counters for Counters.Counter;
    Counters.Counter private _tokenIds;
    
    // 基础URI，用于构建tokenURI
    string private _baseTokenURI;
    
    // 构造函数：设置NFT名称和符号
   constructor(
    string memory name,
    string memory symbol,
    string memory baseTokenURI
) ERC721(name, symbol) Ownable(msg.sender) { 
    _baseTokenURI = baseTokenURI;
}
    
    // 铸造NFT函数
    function mintNFT(address recipient, string memory tokenURI_) 
        public 
        returns (uint256) 
    {
        _tokenIds.increment();
        uint256 newItemId = _tokenIds.current();
        
        _mint(recipient, newItemId);
        _setTokenURI(newItemId, tokenURI_);
        
        return newItemId;
    }
    
    // 设置基础URI（可选，仅所有者可调用）
    function setBaseURI(string memory baseURI) public onlyOwner {
        _baseTokenURI = baseURI;
    }
    
    // 重写_tokenURI函数，返回完整的tokenURI
    function _baseURI() internal view virtual override returns (string memory) {
        return _baseTokenURI;
    }
    function tokenURI(uint256 tokenId)
    public
    view
    virtual
    override(ERC721URIStorage)
    returns (string memory)
{
    return super.tokenURI(tokenId);
}
    
    // 获取当前tokenId计数
    function getCurrentTokenId() public view returns (uint256) {
        return _tokenIds.current();
    }
}