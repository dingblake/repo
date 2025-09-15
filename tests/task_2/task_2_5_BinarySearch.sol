// SPDX-License-Identifier: MIT
pragma solidity >=0.8.2 <0.9.0;
//5.二分查找 (Binary Search)
//题目描述：在一个有序数组中查找目标值。
contract BinarySearch {
    
      function linearSearch(uint256[] memory arr, uint256 target) 
        public 
        pure 
        returns (uint256) 
    {
        for (uint256 i = 0; i < arr.length; i++) {
            if (arr[i] == target) {
                return i;
            }
        }
        return type(uint256).max;
    }
}