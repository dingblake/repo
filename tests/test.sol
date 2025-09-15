// SPDX-License-Identifier: MIT
pragma solidity >=0.8.2 <0.9.0;

contract hello {
    //string public  hello = "hello world!";
    
//     int public  x=-5;
//     uint public  y=2**256 -1 ;
// bytes32 public a = hex"01";
// enum    meiju {
//     active,
//     inactive
// }

// int[] public  arr;
// string[] public  arr2;
// bool[] public  arr3;
// address[]  public arr4;
// bytes10[] public arr5;
// struct   person {
//     uint age;
//     bool sex;
//     string name;

// }

// person public r =person ({age: 11,sex :true ,name : "zhangsan"});

string private hello ="hello";

function sayhello (string memory name)
public 
view 
returns (string memory)
{
   return string.concat(hello,name);
}


} 
