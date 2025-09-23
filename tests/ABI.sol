// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;


contract ABI {
    // function encodedata (string memory text ,uint256 num ) public  pure returns (bytes memory,bytes memory){
    //     return  (
    //         abi.encode(text,num),
    //         abi.encodePacked(text,num)
    //     );

    // }

    // function decodedata (bytes memory encodedata) public  pure  returns  (string memory text,uint256 num){
    //     return abi.decode(encodedata,(string, uint256));

    // }


    // function modular (uint256 x,uint256 y ,uint256 m ) public  pure  returns ( uint,uint) {

    //     return  (
    //         addmod(x, y, m),
    //         mulmod(x, y, m)
    //     );
    // }

    function getRandomNumber() public view returns (uint) {
    return uint(keccak256(abi.encodePacked(block.timestamp, msg.sender))) % 100;
}
}