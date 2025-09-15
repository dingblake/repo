// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract ExampleContract {
    // 状态变量（存储在 storage => stateDB）
    uint256 public stateVar = 100;
    mapping(address => uint256) public balances;
    uint256[] someStorageArray;
    
     // 常量（编译时确定，在字节码中）
    uint256 public constant CONSTANT_VAR = 1000;
    
    // immutable（部署时确定，在代码区）
    uint256 public immutable IMMUTABLE_VAR;
    
    // 部署的时候 调用的
    constructor() {

        IMMUTABLE_VAR = 500;
    }
    
    

    function testMemory() public pure returns(string memory) {

        return "testMemory";
    }
    
    function claimVar() public view returns (uint256) {
        
        //  === 存在 memory
        uint256[] memory memoryArray = new uint256[](3);
        string memory str = "Hello";
        bytes memory tempBytes = new bytes(10);
        // ===
        
        // ====  存在 stack
        uint256 stackVar = 123;
        bool flag = true;
        address sender = msg.sender;
        uint8 smallNum = 255;
        // ===
        
        // === storage 指针 指向 stateDB
        uint256[] storage storageArray = someStorageArray;
        mapping(address => uint256) storage storageMap = balances;
        

        TempStruct memory tempStruct = TempStruct({
            value: 100,
            name: "Test"
        });
        
        // 6. 循环中的变量（栈上）
        for (uint8 i = 0; i < 3; i++) {
            uint256 loopVar = i * 2;
            memoryArray[i] = loopVar;
        }
        
        return stackVar;
    }
   
    
   
    // 纯函数中的变量（全部在栈上）
    function pureFunction(uint256 x) public pure returns (uint256) {
        uint256 y = x * 2;
        return y;
    }

    // sload 操作 
    function viewFunction() public view returns (uint256) {
        return stateVar;
    }
    
    // 引用传递示例  
    function referenceExample(uint256[] memory inputArray) public {
        // memory 到 memory 的引用
        uint256[] memory localArray = inputArray;
        
        // storage 引用 sload => someStorageArray， sstore 操作 修改了 stateDB的 someStorageArray
        someStorageArray.push(100);
    }

    struct TempStruct {
        uint256 value;
        string name;
    }

    function memoryTempStruct(string calldata name, uint val) public view returns(TempStruct memory) {

        // name = string.concat(name, "abc");

        TempStruct memory ts = TempStruct({
            name: name,
            value: val
        });
        return ts;
    }

    function _call(TempStruct calldata ts) public pure returns (TempStruct memory) {
        return ts;
    }

    // memory 可以基于 calldata 得到
    // calldata不能在函数声明
    // https://remix-ide.readthedocs.io/zh-cn/latest/udapp.html  结构体传参交互
    function calldataTempStruct(TempStruct calldata ts) external view returns (TempStruct memory) {

        // ts.name = "qaq";

        // TempStruct memory ts2 = ts;
        // return ts2;
        // TempStruct calldata ttt = TempStruct(12, "zood");
        return _call(ts);

        //  ts;

    }

    // function _call(TempStruct calldata ts) internal  view returns (TempStruct memory) {
    //     TempStruct memory ts2 = ts;
    //     return  ts2;
    // }
    

    
}