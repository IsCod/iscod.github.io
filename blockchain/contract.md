# Solidity Contract


### 关键词

1. immutable
    `immutable`关键词用于声明一个不可变的变量，即在合约部署后就不能再被改变的变量
```solidity
   contract FundMe{
   address public immutable i_owner; //声明合约所有人，immutable设置的变量在构造函数中设置值，即部署时可以设置
   constructor() {
        i_owner = msg.sender;
    }
   }
```

2. constant
    `constant`关键词用于声明一个常量。常量在编译时就已确定，并且不能被修改，可以在合约级别和函数级别定义
```solidity
contract FundMe{
    uint256 public constant MINNUM_USD = 1 * 1e8; //通常使用大写定义常量
}
```

### 构造函数

1. constructor

同其他语言类型，智能合约提供了构造函数 `constructor`, 用于在合约被部署时执行初始化操作

```solidity
contract FundMe{
    address public immutable i_owner;
    constructor() {
        i_owner = msg.sender;
    }
}
```

constructor

### Receive & fallback 特殊函数

`receive` & `fallback` 是`Solidity`中的特殊函数，
主要用于处理以太币转账或者调用合约时没有匹配到具体函数的情况。

1. `fallblock`函数： 
    如果合约接收到以太币但没有对应的`receive`函数或者匹配的函数可以执行, 就会触发fallback函数。fallback函数可以有任意的实现逻辑, 比如记录日志、触发事件等
    
    ```solidity
    fallback() external payable {
        // 处理合约接收到以太币的逻辑
    }
    ```

2. `receive`函数：
    receive函数也是一个特殊函数，用于接受以太币转账。当合约接收到以太币时, 如果存在`receive`函数，以太币直接发送到`receive`函数处理

    ```solidity
    receive() external payable {
    // 处理合约接收到以太币的逻辑
    }
    ```

> 一个合约只能有一个`receive`函数， 而`fallback`函数是可选的。两个函数都是未匹配到具体函数时才调用


* 参考
    * [smartcontractkit](https://github.com/smartcontractkit/full-blockchain-solidity-course-js?tab=readme-ov-file)
    * [eth-converter](https://eth-converter.com/)
    * [chain.link-data-feeds](https://docs.chain.link/data-feeds)
