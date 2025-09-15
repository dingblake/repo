// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0 <0.9.0;

//1.创建一个名为Voting的合约，包含以下功能：
//一个mapping来存储候选人的得票数
//一个vote函数，允许用户投票给某个候选人
//一个getVotes函数，返回某个候选人的得票数
//一个resetVotes函数，重置所有候选人的得票数
contract Voting {
    address public owner;

    mapping(string => uint256) public votesReceived;
    // 存储所有候选人的数组
    string[] public candidates;
    mapping(address => bool) public hasVoted;
    event Voted(address indexed voter, string candidate);
    event VotesReset(address indexed resetBy);
 modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can call this function");
        _;
    }

    constructor(string[] memory _candidates) {
        owner = msg.sender;
        candidates = _candidates;
        
        // 初始化所有候选人的票数为0
        for (uint256 i = 0; i < _candidates.length; i++) {
            votesReceived[_candidates[i]] = 0;
        }
    }
function vote (string memory _candidate) public  {
    // 确保用户还没有投过票
        require(!hasVoted[msg.sender], "You have already voted");
        // 确保候选人是有效的
        require(isValidCandidate(_candidate), "Invalid candidate");
        
        // 记录用户已投票
        hasVoted[msg.sender] = true;
        
        // 增加候选人的票数
        votesReceived[_candidate] += 1;
        
        // 触发投票事件
        emit Voted(msg.sender, _candidate);
    }
        
        
        // 获取候选人得票数
    function getVotes(string memory _candidate) public view returns (uint256) {
        require(isValidCandidate(_candidate), "Invalid candidate");
        return votesReceived[_candidate];
    }
    
    // 重置所有候选人的得票数（只有所有者可以调用）
    function resetVotes() public onlyOwner {
        for (uint256 i = 0; i < candidates.length; i++) {
            votesReceived[candidates[i]] = 0;
        }
        
        // 触发重置事件
        emit VotesReset(msg.sender);
    }
    
    // 内部函数：检查候选人是否有效
    function isValidCandidate(string memory _candidate) private view returns (bool) {
        for (uint256 i = 0; i < candidates.length; i++) {
            if (keccak256(abi.encodePacked(candidates[i])) == keccak256(abi.encodePacked(_candidate))) {
                return true;
            }
        }
        return false;
    }

    // 获取所有候选人列表
    function getCandidates() public view returns (string[] memory) {
        return candidates;
    }
     // 获取总投票数
    function getTotalVotes() public view returns (uint256) {
        uint256 total = 0;
        for (uint256 i = 0; i < candidates.length; i++) {
            total += votesReceived[candidates[i]];
        }
        return total;
    }

}