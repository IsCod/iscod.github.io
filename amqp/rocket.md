# RocketMQ

RocketMQ 是阿里巴巴开发团队参考了kafka的一些设计理念和架构思路。比如kafka的消息分区，分布式架构等都有类似之处。

## Rocket特点

1. 支持事务消息
1. 支持死信队列
1. 支持延迟队列
1. 支持消息重试
1. 支持广播消息

> 不支持优先级别队列

## Rocket持久化机制

1. 写入磁盘: 当生产者发送消息时，RocketMQ会将消息写入磁盘的CommitLog文件中。CommitLog是一个顺序写入的文件，消息会被追加到文件末尾
1. 异步刷盘：RocketMQ默认采用异步刷盘，即将生产者发送的消息先写入PageCache中，然后异步将消息刷写到磁盘（CommitLog文件）。这样可以避免每次写入消息都要进行磁盘IO
1. 消息索引: 在RocketMQ中，使用IndexFile存储消息索引信息的文件，每个主题下的消息队列都有一个对应的IndexFile文件用来记录消息的索引信息，比如消息的Key,消息的偏移量，消息文件所在的位置。以实现消息的快速查找
1. 消息消费: RocketMQ通过consumeQueue用以记录消费者消息消息的进度和位置。每个主题下的每个消息消费队列都有一个对应的consumeQueue文件。

RocketMQ通过共享同一个CommitLog文件，避免了kafka分区数过多导致的磁盘IO压力（kafka中每一个分区都有宁独立的日志文件）和随机读写的等造成的性能瓶颈（超过64个分区时性能下降明显）。
RocketMQ的queue只存储少量数据，更轻量化，对于磁盘的访问是串行化避免磁盘竞争。缺点在于：写虽然是顺序的，但读是随机的，先读consumeQueue（获取queue的位置），再读commitLog，会降低消息读的效率。
RocketMQ在默认情况下采用的是异步刷盘机制，只能保证在系统不掉电级别的数据不丢失。如果对数据可靠性要求较高，可以采用同步刷盘机制即消息写入PageCache后、立即将数据刷写到磁盘，而后在返回给生产者ACK。

### CommitLog

在老版本的RocketMQ（3.x版本）中所有的主题共用一个CommitLog文件，这种设计称为"全局有序"的方式，即所有主题的消息被顺序写入到同一个CommitLog文件中。

然而，从RocketMQ 4.x版本开始，引入列"主题队列"，每个主题都被划分为多个队列，每个队列都有独立的CommitLog文件和ConsumeQueue文件。这种设计可以提高消息的隔离性，水平扩展和负载均衡性。

## RocketMQ相比Kafka有哪些优劣点？
1. RocketMQ支持严格的消息顺序性，可以确保消息按照发送顺序进行消费。而Kafka只保证分区内的消息有序，无法保证全局有序
1. RocketMQ具备高可靠性，支持主从同步复制和故障自动恢复机制，能够在节点故障发生时保证消息的可靠传输。而KafkaMQ则使用分区副本机制来实现高可用性，但在节点故障恢复时需要手动进行重新分配

* 参考
* [docker-rabbit](https://hub.docker.com/_/rabbitmq)
* [rabbit](https://www.rabbitmq.com/documentation.html)
* [Kafka、RabbitMQ和RocketMQ差异](https://support.huaweicloud.com/productdesc-kafka/kafka_pd_0003.html)