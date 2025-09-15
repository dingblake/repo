// SPDX-License-Identifier: MIT
pragma solidity >=0.8.2 <0.9.0;
//4.合并两个有序数组 (Merge Sorted Array)
//题目描述：将两个有序数组合并为一个有序数组。
pragma solidity ^0.8.0;

contract MergeSortedArrays {

     function merge(uint256[] memory arr1, uint256[] memory arr2) 
        public 
        pure 
        returns (uint256[] memory) 
    {
        uint256 n1 = arr1.length;
        uint256 n2 = arr2.length;
        uint256 total = n1 + n2;
        uint256[] memory result = new uint256[](total);
        
        uint256 i = 0;
        uint256 j = 0;
        
        // 使用一个循环完成合并
        for (uint256 k = 0; k < total; k++) {
            if (i < n1 && (j >= n2 || arr1[i] <= arr2[j])) {
                result[k] = arr1[i];
                i++;
            } else if (j < n2) {
                result[k] = arr2[j];
                j++;
            }
        }
        
        return result;
    }

}