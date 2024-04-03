# Solidity Contract

## solidity

#### `immutable`

`immutable`关键词用于声明一个不可变的变量，即在合约部署后就不能再被改变的变量

```solidity
contract FundMe{
   address public immutable i_owner; //声明合约所有人，immutable设置的变量在构造函数中设置值，即部署时可以设置
   constructor() {
        i_owner = msg.sender;
    }
}
```
#### `constant`

`constant`关键词用于声明一个常量。常量在编译时就已确定，并且不能被修改，可以在合约级别和函数级别定义

```solidity
contract FundMe{
    uint256 public constant MINNUM_USD = 1 * 1e8; //通常使用大写定义常量
}
```

### 构造函数`constructor`

`constructor`构造函数是一种特殊类型的函数，用于在合约被部署（deployed）时执行初始化操作。功能类型与其他编程语言，都是用于初始化操作

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

#### `fallblock` 

如果合约接收到以太币但没有对应的`receive`函数或者匹配的函数可以执行, 就会触发fallback函数。fallback函数可以有任意的实现逻辑, 比如记录日志、触发事件等
    
```solidity
fallback() external payable {
    // 处理合约接收到以太币的逻辑
}
```

#### `receive`

receive函数也是一个特殊函数，用于接受以太币转账。当合约接收到以太币时, 如果存在`receive`函数，以太币直接发送到`receive`函数处理

```solidity
receive() external payable {
// 处理合约接收到以太币的逻辑
}
```

> 一个合约只能有一个`receive`函数， 而`fallback`函数是可选的。两个函数都是未匹配到具体函数时才调用

#### `library`

`library`是一种特殊的合约类型，通过`library`，开发者可以将常用的功能封装成库，提高代码的重复性和可维护性

```solidity
//定义一个价格相关的library
import {AggregatorV3Interface} from "@chainlink/contracts@0.8.0/src/v0.8/interfaces/AggregatorV3Interface.sol";

library PriceConverter {
       function getPrice() public view  returns (uint256)  {
         AggregatorV3Interface feed = AggregatorV3Interface(0x694AA1769357215DE4FAC081bf1f309aDC325306);
        (
            /* uint80 roundID */,
            int answer,
            /*uint startedAt*/,
            /*uint timeStamp*/,
            /*uint80 answeredInRound*/
        ) = feed.latestRoundData();
        // answer / 1e8 = 1 ether
        return uint256(answer * 1e10);// 1 wei price
    }
}

contract FundMe {
    useing PriceConverter for uint256;
}
```

### solidity

```sol
// SPDX-License-Identifier: MIT
// get funds form users
//withdraw funds
// set min funding value for usd
pragma solidity 0.8.24;

import "./PriceConverter.sol";

error NotOwer();

contract FundMe {
    using PriceConverter for uint256;
    uint256 public constant MINNUM_USD = 1 * 1e18;
    //391
    //2490

    address[] public fundusers;
    mapping (address => uint256) public fundUserToAmount;

    address public immutable i_owner;

    constructor() {
        i_owner = msg.sender;
    }
    

    function fund() public payable {
        //mas.value is wei 1ether = 1e18 wei
        require(msg.value.getConversionRate() >= MINNUM_USD, "Didn't send min eth");
        fundusers.push(msg.sender);
        fundUserToAmount[msg.sender] += msg.value;
    }

    function withdraw() public onlyOwer {
        for (uint i = 0; i< fundusers.length; i++) 
        {
            // address fundaddress = fundusers[i];
            payable(msg.sender).transfer(address(this).balance);
            bool sendSucess = payable(msg.sender).send(address(this).balance);
            require(sendSucess, "Send Fail");

            (bool callSucess, /* bytes memory callReturn */) = payable(msg.sender).call{value: address(this).balance}("");

            require(callSucess, "Call fail");
        }
    }

    modifier onlyOwer {
        if (msg.sender != i_owner) {revert NotOwer();}
        _;
    }
    receive() external payable { 
        fund();
    }

    fallback() external payable {
        fund();
     }
}
```

## Ganache 网络

### install ganache

```bash
brew install --cask ganache
```
## ethers.js

### ethers部署合约

```js
//"ethers": "^6.11.1", "fs-extra": "^11.2.0", "solc": "0.8.7-fixed"
const ethers = require("ethers");
const fs = require("fs-extra");

async function main () {

  const provier = new ethers.JsonRpcProvider("http://127.0.0.1:7545"); //Ganache网络api

  const wallet = new  ethers.Wallet("0x535ca033907d6aa8e05da5db576b1890d07d2f8458babda12f169b081b01f6e5", provier);

  abi  = fs.readFileSync("./simpleStorage_sol_SimpleStorage.abi", "utf8");
  bin  = fs.readFileSync("./simpleStorage_sol_SimpleStorage.bin", "utf8");

  const contractFactory = new ethers.ContractFactory(abi, bin, wallet);
  console.log("Deploying, please wait...")
  const contracts = await contractFactory.deploy({ gasPrice: 100000000000 });
  const deployment = await contracts.waitForDeployment()
  console.log("This is the contract deployment: ")
  console.log(deployment)
}

main().then(() => process.exit(0)).catch((error) => {
  console.log(error);
  process.exit(1);
});
```

* 参考
    * [smartcontractkit](https://github.com/smartcontractkit/full-blockchain-solidity-course-js?tab=readme-ov-file)
    * [eth-converter](https://eth-converter.com/)
    * [chain.link-data-feeds](https://docs.chain.link/data-feeds)
