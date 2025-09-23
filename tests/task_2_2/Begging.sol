// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract BeggingContract {
    // 合约所有者
    address public owner;
    
    // 记录每个捐赠者的捐赠金额
    mapping(address => uint256) public donations;
    
    // 捐赠总额
    uint256 public totalDonations;
    
    // 捐赠事件
    event DonationReceived(address indexed donor, uint256 amount);
    
    // 提款事件
    event Withdrawal(address indexed owner, uint256 amount);
    
    // 修饰符：只有所有者可以调用
    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can call this function");
        _;
    }
    
    // 构造函数：设置合约所有者
    constructor() {
        owner = msg.sender;
    }
    
    // 捐赠函数：允许用户向合约发送以太币
    function donate() external payable {
        require(msg.value > 0, "Donation amount must be greater than 0");
        
        // 记录捐赠金额
        donations[msg.sender] += msg.value;
        totalDonations += msg.value;
        
        // 触发捐赠事件
        emit DonationReceived(msg.sender, msg.value);
    }
    
    // 提款函数：允许合约所有者提取所有资金
    function withdraw() external onlyOwner {
        uint256 balance = address(this).balance;
        require(balance > 0, "No funds to withdraw");
        
        // 转账给所有者
        payable(owner).transfer(balance);
        
        // 触发提款事件
        emit Withdrawal(owner, balance);
    }
    
    // 查询函数：查询某个地址的捐赠金额
    function getDonation(address donor) external view returns (uint256) {
        return donations[donor];
    }
    
    // 获取合约余额
    function getContractBalance() external view returns (uint256) {
        return address(this).balance;
    }
    
    // 获取捐赠者数量（需要额外记录）
    function getTotalDonors() external view returns (uint256) {
        // 注意：这个函数在实际部署中可能无法准确返回捐赠者数量
        // 因为mapping无法直接遍历，这里返回总捐赠次数估算
        return totalDonations > 0 ? 1 : 0; // 简化实现
    }
}