// SPDX-License-Identifier: MIT
pragma solidity >=0.8.2 <0.9.0;
//2.反转字符串 (Reverse String)
//题目描述：反转一个字符串。输入 "abcde"，输出 "edcba"
 contract Reverse_string {

    function reverse(string memory _str) public pure returns (string memory) {
        // 将字符串转换为字节数组以便操作
        bytes memory strBytes = bytes(_str);
        uint256 length = strBytes.length;
        
        // 处理空字符串和单字符字符串的特殊情况
        if (length == 0) {
            return "";
        }
       
        // 创建用于存储反转后字符串的字节数组
        bytes memory reversed = new bytes(length);
        
        // 使用双指针技术反转字符串
        for (uint256 i = 0; i < length; i++) {
            reversed[i] = strBytes[length - 1 - i];
        }
        
        // 将字节数组转换回字符串并返回
        return string(reversed);
    }
 }



//3. 用 solidity 实现整数转罗马数字

 

//4.合并两个有序数组 (Merge Sorted Array)
//题目描述：将两个有序数组合并为一个有序数组。


//5.二分查找 (Binary Search)
//题目描述：在一个有序数组中查找目标值。