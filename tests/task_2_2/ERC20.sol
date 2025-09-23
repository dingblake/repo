// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract SimpleERC20 {

    string public name;
    string public symbol;
    uint8 public decimals;

    // 代币供应数据 - 确保在此处声明 _totalSupply
    uint256 private _totalSupply;


    mapping (address => uint256) private _balances; 
    mapping (address => mapping (address => uint256)) private  _allowances;

    event  Transfer(address indexed from ,address indexed  to,uint256 value); 
    event  Approval(address indexed owner ,address indexed  spender,uint256 value); 

address public owner;

 // 修饰器：只有所有者可调用
    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can call this function");
        _;
    }
// 构造函数：初始化代币
    constructor(string memory _name, string memory _symbol, uint8 _decimals, uint256 _initialSupply) {
        name = _name;
        symbol = _symbol;
        decimals = _decimals;
        owner = msg.sender;
        
        // 初始铸造给部署者
        _mint(msg.sender, _initialSupply * (10 ** uint256(_decimals)));
    }

    // 查询总供应量
    function totalSupply() public view returns (uint256) {
        return _totalSupply;
    }
    
    // 查询账户余额
    function balanceOf(address _account) public view returns (uint256) {
        return _balances[_account];
    }
    
    // 转账功能
    function transfer(address _to, uint256 _value) public returns (bool) {
        _transfer(msg.sender, _to, _value);
        return true;
    }
    
    // 查询授权额度
    function allowance(address _owner, address _spender) public view returns (uint256) {
        return _allowances[_owner][_spender];
    }
    
    // 授权功能
    function approve(address _spender, uint256 _value) public returns (bool) {
        _approve(msg.sender, _spender, _value);
        return true;
    }
    
    // 代扣转账功能
    function transferFrom(address _from, address _to, uint256 _value) public returns (bool) {
        _spendAllowance(_from, msg.sender, _value);
        _transfer(_from, _to, _value);
        return true;
    }
    
    // 增发代币（仅所有者）
    function mint(address _to, uint256 _amount) public onlyOwner {
        _mint(_to, _amount);
    }
    
    // 内部转账实现
    function _transfer(address _from, address _to, uint256 _value) internal {
        require(_from != address(0), "ERC20: transfer from the zero address");
        require(_to != address(0), "ERC20: transfer to the zero address");
        require(_balances[_from] >= _value, "ERC20: transfer amount exceeds balance");
        
        _balances[_from] -= _value;
        _balances[_to] += _value;
        
        emit Transfer(_from, _to, _value);
    }
    
    // 内部铸造实现
    function _mint(address _account, uint256 _amount) internal {
        require(_account != address(0), "ERC20: mint to the zero address");
        
        _totalSupply += _amount;
        _balances[_account] += _amount;
        
        emit Transfer(address(0), _account, _amount);
    }
    
    // 内部授权实现
    function _approve(address _owner, address _spender, uint256 _value) internal {
        require(_owner != address(0), "ERC20: approve from the zero address");
        require(_spender != address(0), "ERC20: approve to the zero address");
        
        _allowances[_owner][_spender] = _value;
        
        emit Approval(_owner, _spender, _value);
    }
    
    // 内部消耗授权额度实现
    function _spendAllowance(address _owner, address _spender, uint256 _amount) internal {
        uint256 currentAllowance = allowance(_owner, _spender);
        if (currentAllowance != type(uint256).max) {
            require(currentAllowance >= _amount, "ERC20: insufficient allowance");
            _approve(_owner, _spender, currentAllowance - _amount);
        }
    }
}