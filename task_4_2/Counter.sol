// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Counter {
    uint256 private count;
    address public owner;
    
    event CountIncreased(uint256 newCount);
    
    constructor() {
        owner = msg.sender;
        count = 0;
    }
    
    function increment() public returns (uint256) {
        count += 1;
        emit CountIncreased(count);
        return count;
    }
    
    function getCount() public view returns (uint256) {
        return count;
    }
    
    function reset() public {
        require(msg.sender == owner, "Only owner can reset");
        count = 0;
        emit CountIncreased(count);
    }
}