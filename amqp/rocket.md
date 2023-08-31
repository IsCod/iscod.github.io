# RocketMQ

## Rocket特点

1. 支持事务消息
1. 支持死信队列
1. 支持延迟队列
1. 支持消息重试
1. 支持广播消息

> 不支持优先级别队列

## Rocket持久化机制

1. commitLog：日志数据文件，被所有queue共享，大小1G
2. consumeQueue: 逻辑queue，记录了queue在commitLog中的物理偏移量offset和消息内容大小以及消息tag的hash值，大小约为600w字节
2. indexFile

Rocker所有队列共享一个日志数据文件，避免了kafka分区数过多、日志文件过多导致的磁盘IO读写压力较大造成的性能瓶颈（超过64个分区时性能下降明显）。
Rocket的queue只存储少量数据，更轻量化，对于磁盘的访问是串行化避免磁盘竞争。缺点在于：写虽然是顺序的，但读是随机的，先读consumeQueue（获取queue的位置），再读commitLog，会降低消息读的效率。

* 参考
* [docker-rabbit](https://hub.docker.com/_/rabbitmq)
* [rabbit](https://www.rabbitmq.com/documentation.html)
* [Kafka、RabbitMQ和RocketMQ差异](https://support.huaweicloud.com/productdesc-kafka/kafka_pd_0003.html)