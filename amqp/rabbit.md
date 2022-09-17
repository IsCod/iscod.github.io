# rabbit

## rabbit特点

1. 支持优先级别队列
1. 支持延迟队列
1. 支持事务消息
1. 支持死信队列
1. 支持广播消息

> 不支持消息的顺序性、消息重试

### run

```sh
docker run -d -p5672:5672 --hostname my-rabbit --name some-rabbit rabbitmq
```

* 参考
	* [docker-rabbit](https://hub.docker.com/_/rabbitmq)
	* [rabbit](https://www.rabbitmq.com/documentation.html)
	* [Kafka、RabbitMQ和RocketMQ差异](https://support.huaweicloud.com/productdesc-kafka/kafka_pd_0003.html)