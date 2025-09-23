// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0 <0.9.0;


contract Demo {

    function test () public  {
// msg.data;
// msg.gas;
// msg.sender;
// msg.value;
// tx.origin
// block.blockhash(blockNumber);
    }
}

contract BlockInfo {
    function getBlockDetails() public view returns (uint, uint) {
        return (block.number, block.timestamp);
    }
}