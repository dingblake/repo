// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Owner {
    address public owner;

    constructor() {
        owner = msg.sender;
    }

    // 自定义修饰符，仅允许合约所有者调用
    modifier onlyOwner {
        require(msg.sender == owner, "only ziji");
        _;
    }

    function changeOwner(address newOwner) public onlyOwner {
        owner = newOwner;
    }
}