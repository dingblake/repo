// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
contract Payment {
    function pay() public payable {}
    function getBalance() public view returns (uint) {
        return address(this).balance;
    }
}


contract Auth {
    address public owner;
    constructor() { owner = msg.sender; }
    function isOwner() public view returns (bool) {
        return msg.sender == owner;
    }
}