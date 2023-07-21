# kafka

## kafka特点

1. 消息持久化
1. 高吞吐量
1. 扩展性强（可以动态扩展）
1. 多客户端支持
1. Kafka stream（流处理）
1. 安全机制
1. 数据备份
1. 轻量级
1. 支持消息压缩（支持自定义压缩）

> 不支持延迟队列、优先级队列、消息重试

## kafka适用场景

日志收集、消息系统、流式处理。不适合限时订单

kafka有以下优点：多生产者多消费，基于磁盘的数据存储，高伸缩性，高性能

## 基础概念

1. 主题
1. 生产者&消费者、偏移量、消费者群组

## 副本

`kafka`副本主要用于提高数据可靠性。
`kafka`默认生成1个副本，一般生产环境设置2-3个副本，过多的副本会增加磁盘存储空间和网络数据传输降低效率

`kafka`中副本分为`Leader`和`Follower`。生产者只会发送数据到`Leader`。`Follower`会自动从`Leader`中同步

AR=Isr+Osr

`Isr`: 表示和`Leader`同步的`Follower`集合

Osr: 表示`Follower`和`Leader`同步时，延迟过多（超时）的副本

#### 副本选举策略

根据ISR中存活的，且在Osr中排序较前的做为`Leader`

```bash
kafka-topics --create --topic test --replication-factor 2 --partitions 4 --bootstrap-server 127.0.0.1:9092 # 创建了四个分区，两个副本的topic
```

### Partition

#### 分区分配策略

    * range

    首先对Topic里面的分区按照序号排序，并对消费者按照字母排序，通过 `partition`数量/`consumer`数量 决定每个消费者应该消费几个分区，余数由前面的几个分区消费。
    当Topic数量较多时容易产生数据倾斜，应注意。

    * cooperative-sticky

    粘性分区分配策略是尽量均衡的放置分区到消费者上面，当一个消费者宕机时, 会尽量保持原有的分配的分区不变化

    * roundrobin

    针对所有的topic, 把所有的`partition`和`consumer`都罗列出来, 然后进行hashcode进行排序, 最后通过轮询算法来分配`partition`到各个消费者

默认分区策略是：`range` + `roundrobin`

> 每一个分区只能有一个消费者消费，当消费者组内的消费者数量大于分区数时，多出的消费者必然不会工作

### kafka配置

#### 系统配置文件

* broker.id

在集群下的唯一id,要求是整数。如果服务器ip发生变化，而broker.id没有变化，则不影响consumers消费情况

* listeners

监听列表，逗号分割，如果hostname为`0.0.0.0`则绑定所有的网卡地址，如果为空，则绑定默认网卡

* zookeeper.connect

zookeeper集群地址，多个采用逗号分隔

* auto.create.topics.enable=true

是否运行自动创建主题，如果设置为true, 那么product,consumer,fetach一个不存在的主题时，会自动创建。一般处于安全考虑会设置为false

* log.dirs

kafka消息数据存放的目录，可以设置多个，采用逗号分割，如果设置多个，kafka会根据`最少使用`原则，把同一分区的日志片段保存到同一路径下，会往拥有最少分区的路径新增分区

1. 主题配置

* num.partitions

新建主题的分区个数。该参数根据消费者处理能力和数量进行确定，分区数应大于消费者数量（1，便于后面扩充消费者数量，2，每一个分区至少一个消费者，分区数大于消费者数量时，消费者会平分分区[分区策略]([https://iscod.github.io/#/amqp/kafka?id=partition]))

* message.max.bytes

消息最大字节数，生产者消费者应该设置一致。该值一般与`fetch.message.max.bytes`配合设置。

#### 生产者配置
```go
configMap.SetKey("request.timeout.ms", "30000") //生产者请求的确认超时时间
```

#### 消费者配置

```go
configMap.SetKey("enable.auto.commit", true)     //是否自动提交offset，默认true
configMap.SetKey("auto.commit.interval.ms", 5000) //自动提交offset的时间间隔,默认5秒，可以修改时间提高频率
```

#### 手动提交`offset`

`kafka`的自动提交虽然便利，但是业务开发有时需要手动控制`offset`的提交时机，因此可以采用手动模式进行提交

手动提交有两种方法：`commitSync(同步提交)`和`commitAsync(异步提交)`两种，同步提交会阻塞当前线程，一直到提交成功，并且会自动失败重试。而异步提交没有失败重试机制，故有可能提交失败。

合理的方式是：在主线程采用异步提交，而退出线程时采用回调等形式进行一次同步提交

```go
configMap.SetKey("enable.auto.commit", false)     //首先关闭自动提交
```

#### 指定`offset`消费

#### 重复消费&漏消费

* 重复消费

采用自动提交`offset`时，默认策略是每5s进行一次提交。而在下次未提交之前, `consumer`拉取了数据就进行处理, 却因为异常退出, 而未进行`commit`。
那么`consumer`重启后，则从上次提交的`offset`处继续消费，造成重复消费。因此数据的处理要做幂等性。

* 漏消费

设置`offset`手动提交时，数据拉取后程序内存中，未处理完毕，此时消费者线程kill掉，那么offset采用的异步提交已完成。那么就导致这部分内存中数据未处理而丢失

如何避免漏消费？

`kafka`消费端将消费过程和提交`offset`过程做原子绑定（事务），此时将`kafka`的offset保存到支持事务的自定义介质中（如mysql）进行处理。这样就避免业务未处理完毕而进行了commit


#### 数据积压

1. 消费者消费能力不足

    如果是消费能力不足, 可通过增加topic的分区数, 并同时提高消费者数量, 消费者数=分区数, 从而增加消费能力

1. 下游数据处理不足（拉取速度/处理时间 < 生产速度）

    如果是下游数据处理不足, 则可提高每批次的拉取数量。并配合每次拉取最大字节数。提高数据拉取熟读

如何增加吞吐量？

```go
//partition设置
configMap.SetKey("batch.size", 1000000) //数据达到多大容量时发送，可提高至32kb
configMap.SetKey("linger.ms", 5) //每批次等待时间，超过5ms也发送一次

//consumer设置，先提高poll条数
configMap.SetKey("fetch.max.bytes", 52428800) //提高每批次抓取最大上限
```

## 消息送达语义

* At most once：消息发送或消费至多一次
* At least once：消息发送或消费至少一次
* Exactly once：消息恰好只发送一次或消费一次

Kafka默认的`Producer`消息送达语义是 `At least once` ，也就是至少投递一次，保证消息不丢失。

### 消息重复和消息幂等

大多数云消息队列 Kafka 版消费的语义是at least once, 但是这样的消费语义就无法保证消息不重复。
在出现网络问题、客户端重启时，均有可能造成少量重复消息，此时应用消费端如果对消息重复比较敏感（例如订单交易类），则应该做消息幂等。

常用做法是：

1. 发送消息时，传入key作为唯一流水号ID。
1. 消费消息时，判断key是否已经消费过，如果已经被消费，则忽略，如果没消费过，则消费一次。
1. 如果应用本身对少量消息重复不敏感，则不需要做此类幂等检查。

### 消息失败

当消费者拿到某条消息后执行逻辑失败，例如应用程序出现故障，导致消息处理失败，就需要人工干预，一般的有两种处理方式：

1. 失败后再次尝试执行消费逻辑，这种方式可能造成消息阻塞，无法向前推进业务，造成消息堆积
1. 消息队列处理失败后推送到相应的服务，或者消息（例如创建一个Topic专门存储失败消息），然后定时检查失败消息，或发送到系统报警进行人工处理

## 应用场景

#### 1. 消息传递&异步处理

场景说明：用户注册后，需要发送注册邮件和注册短信。传统做法有两种：1.串行化，2.并行处理

1. 串行化

    ![用户注册串行化](https://iscod.github.io/images/kafka1.png)

1. 并行处理

    ![用户注册并行处理](https://iscod.github.io/images/kafka2.png)

    引入消息队列，对非必须逻辑进行异步处理。改造后的架构如下：

1. 消息队列

    ![用户注册消息队列](https://iscod.github.io/images/kafka3.png)

    可以看到，引入消息队列将非必须的业务逻辑，进行异步处理后。用户的响应时间仅仅相当于是注册信息写入数据库的时间。系统的吞吐量比串行提高了3倍，比并行提高了2倍

#### 2. 应用解耦

场景说明: 用户下单后，订单系统需要通知库存系统、积分系统、物流系统等。传统做法是，订单系统调取库存等系统的API。

![传统模式用户下单](https://iscod.github.io/images/kafka4.png)

传统模式的缺点是，如果库存系统出现异常无法访问，则订单扣减库存失败，从而导致订单失败，订单系统与库存系统高度耦合

如何解决上面的问题呢？引入消息队列后的系统架构如下：

![引入消息队列用户下单](https://iscod.github.io/images/kafka5.png)

订单系统: 用户下单后，订单系统完成持久化，将消息写入消息队列，返回用户订单下单成功
库存系统: 订阅下单消息，采用拉/推送方式，获取订单信息，库存系统根据订单信息，进行库存操作

假如：在下单时库存系统不能正常使用。也不影响正常下单，因为下单后，订单系统写入消息队列就不再关心其他的后续操作了。实现订单系统与库存系统的应用解耦

#### 3. 流量削峰

流量削峰也是消息队列中的常用场景，一般在秒杀或团抢活动中使用广泛。

应用场景：秒杀活动，一般会因为流量过大，导致流量暴增，应用挂掉。为解决这个问题，一般需要在应用前端加入消息队列。

![流量削峰](https://iscod.github.io/images/kafka6.png)

用户的请求，服务器接收后，首先写入消息队列。假如消息队列长度超过最大数量，则直接抛弃用户请求或跳转至错误页面。
秒杀业务根据消息队列中的请求信息，做后续处理。

#### 4. 日志同步

Kafka设计的初衷是为了应对大量日志传输场景，通过异步处理方式将日志消息同步到消息服务，再通过其它组件对日志进行实时和离线分析。
日志同步的关键部件是：日志客户端采集，Kafka消息队列，后端日志处理程序，三部分实现

![日志同步](https://iscod.github.io/images/kafka7.png)

#### 5. 活动跟踪

跟踪用户

#### 6. 流处理

实时数据流计算



* 参考
* [kafka-tencent](https://cloud.tencent.com/developer/article/1974648)
* [kafka-huawei](https://support.huaweicloud.com/productdesc-kafka/kafka-scenarios.html)
* [Kafka、RabbitMQ和RocketMQ差异](https://support.huaweicloud.com/productdesc-kafka/kafka_pd_0003.html)