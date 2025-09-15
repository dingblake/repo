// SPDX-License-Identifier: MIT
pragma solidity >=0.8.2 <0.9.0;
//用 solidity 实现整数转罗马数字
contract RomanNumerals {
   struct RomanSymbol {
        uint256 value;
        string symbol;
    }

    RomanSymbol[] private romanSymbols;

    constructor() {
        // 按值从大到小初始化罗马数字符号
        romanSymbols.push(RomanSymbol(1000, "M"));
        romanSymbols.push(RomanSymbol(900, "CM"));
        romanSymbols.push(RomanSymbol(500, "D"));
        romanSymbols.push(RomanSymbol(400, "CD"));
        romanSymbols.push(RomanSymbol(100, "C"));
        romanSymbols.push(RomanSymbol(90, "XC"));
        romanSymbols.push(RomanSymbol(50, "L"));
        romanSymbols.push(RomanSymbol(40, "XL"));
        romanSymbols.push(RomanSymbol(10, "X"));
        romanSymbols.push(RomanSymbol(9, "IX"));
        romanSymbols.push(RomanSymbol(5, "V"));
        romanSymbols.push(RomanSymbol(4, "IV"));
        romanSymbols.push(RomanSymbol(1, "I"));
    }
     function toRoman(uint256 num) public view returns (string memory) {
        require(num > 0 && num < 4000, "Number out of range (1-3999)");
        
        bytes memory roman;
        uint256 remaining = num;

        // 贪心算法：从最大的符号开始遍历
        for (uint256 i = 0; i < romanSymbols.length; i++) {
            RomanSymbol memory symbol = romanSymbols[i];
            while (remaining >= symbol.value) {
                roman = abi.encodePacked(roman, symbol.symbol);
                remaining -= symbol.value;
            }
        }
        return string(roman);
    }


    
}