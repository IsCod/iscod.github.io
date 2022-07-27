# 调度器

Gloang调度器如何实现？Golang并发模型？Golang的GMP模型？这三个问题的本质都是要归属到Golang的GMP模型。

## G-M-P

* G: goroutine 调度实体, 既用户代码, 在本地队列中不断切换执行
* M: machine 内核线程, 既系统线程, 负责代码执行
* P: processor 逻辑处理器, 保存了调度上下文。也可以理解为局部的一个调度器。P的数量由`runtime.GOMAXPROCS`控制




![GMP关系](https://iscod.github.io/images/gmp_1.png)
![M阻塞](https://iscod.github.io/images/gmp_2.png)
![P闲置](https://iscod.github.io/images/gmp_3.png)


* 参考
	* [Golang调度器](https://studygolang.com/articles/9610)
	* [Golang并发模型](https://www.jianshu.com/p/f9024e250ac6)