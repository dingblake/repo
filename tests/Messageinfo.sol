// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;


contract Messageinfo {

    function getContractDetails() public view returns (address, uint) {
        return (address(this), address(this).balance);
    }

    function getContractAddress() public view returns (address) {
        return address(this);
    }

    function getContractBalance() public view returns (uint) {
        return address(this).balance;
    }
}