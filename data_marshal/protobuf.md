# Protobuf

protocol buffer是Google的一种独立和轻量级的数据交换格式。以二进制结构进行存储。

协议缓冲区是谷歌的语言无关、平台无关、用于序列化结构化数据的可扩展机制——考虑XML，但更小、更快、更简单。一旦定义了数据的结构化方式，就可以使用特殊生成的源代码轻松地向各种数据流写入和读取结构化数据，并使用各种语言。

```protobuf
message Person {
  required string name = 1;
  required int32 id = 2;
  optional string email = 3;
}
```

* 参考
    * [protobuf](https://github.com/protocolbuffers/protobuf)
    * [golang/protobuf](https://github.com/golang/protobuf)