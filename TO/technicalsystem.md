# Technical system

技术体系包含`编程语言`、`工程组件`、`分布式中间件`、`云原生`、`场景化解决方案`五个部分

围绕这五个方面形成自身的的一套技术体系，才能成为一名合格的架构师，对项目全局才能有更深层的理解和应用。

## 编程语言(golang)

#### 1. golang的特性

#### 2. golang并发编程


## 工程组件

#### 1. 微服务工具集go-kit

#### 2. 高性能grpc

* RPC与HTTP的区别在哪里？

RPC与HTTP的区别是：

RPC是一个远程调用`方案`，HTTP是仅仅是一个通信传输`协议`

RPC包含了整套的调用方案内容，其中包括 服务发现、通信传输（也可以用HTTP）、序列化、负载均衡策略 等组件

* 构建高性能grpc需要考虑的问题

1. 用户鉴权
1. grpc数据传递
1. 拦截器
1. 负载均衡
1. 服务的健康检查
1. 数据传输方式
1. 服务间的认证
1. 服务限流、熔断与降级
1. 日志追踪

#### 3. 网关grpc-gateway

#### 4. web组件[gin](https://gin-gonic.com/zh-cn/docs/introduction/)

gin是一个以性能著称的web框架，gin 具有快速、支持中间件、JSON验证、路由组、错误管理、内置渲染等特性。
同时gin采用RESTful api设计风格，并能轻松实现实现api版本管理。




```go
// 快速构建gin http服务
// https://github.com/iscod/iscod.github.io/tree/master/example/gin
package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "net"
    "net/http"
)

func main() {
    c, err := net.Listen("tcp", "localhost:8080")
    if err != nil {
        fmt.Printf("Error creating %s", err)
        return
    }
    g := gin.New()
    g.Handle("GET", "/", func(context *gin.Context) {
        context.String(http.StatusOK, "ok")
    })
    g.POST("/post", func(context *gin.Context) {
        context.JSON(http.StatusOK, gin.H{"message": "ok"})
    })

    //g.Use(gin.Logger(), gin.Recovery()) //添加logger和Recovery中间件
    if err = g.RunListener(c); err != nil {
        fmt.Printf("Error creating %s", err)
    }
}

```

#### 5. 命令行Cli组件 [Cobra](https://github.com/spf13/cobra/blob/main/site/content/user_guide.md)

```go
package main

import (
    "fmt"
    "github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
    Short:     "Usage: cobra create Command",
    Long:      "详细描述",
    Use:       "create",                                                   //命令名称
    Args:      cobra.MatchAll(cobra.OnlyValidArgs, cobra.MinimumNArgs(1)), // 校验至少一个arg
    ValidArgs: []string{"a", "b", "c"},                                    // 校验arg必须在指定选项内
    PreRun: func(cmd *cobra.Command, args []string) {//钩子函数
        fmt.Println("pre run")
    },
    PersistentPreRun: func(cmd *cobra.Command, args []string) {//钩子函数
        fmt.Println("persistent pre run")
    },
    PostRun: func(cmd *cobra.Command, args []string) {//钩子函数
        fmt.Println("post run")
    },
    PersistentPostRun: func(cmd *cobra.Command, args []string) {//钩子函数
        fmt.Println("persistent post run")
    },
    //PersistentPostRunE: func(cmd *cobra.Command, args []string) error {//钩子函数
        //	fmt.Println("Persistent PostRunE")
        //	return nil
    //},
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Printf("create Command Command Run success args: %v \n", args)
    },
}

var subCommand = &cobra.Command{
    Use: "sub",
    Run: func(cmd *cobra.Command, args []string) {
        number, err := cmd.Flags().GetInt("number")
        if err != nil {
            fmt.Printf("sub command run Error: %s \n", err)
            return
        }
        fmt.Printf("sub command run Flags number: %d \n", number)
    },
}

var rootCmd = &cobra.Command{ //根命令不设置RUN
    Short: "Usage: cobra [OPTIONS] Command", //简要描述
    Long:  "Usage: cobra [OPTIONS] Command", //详细描述
    Use:   "",
}

func init() {
    rootCmd.AddCommand(createCmd)
    createCmd.AddCommand(subCommand) //createCmd命令的子命令
    subCommand.Flags().Int("number", 0, "数量")// 添加flags参数
}

func main() {
    rootCmd.Execute()
}
```

#### 6. 验证器 [Validator](https://github.com/go-playground/validator)

`Validator`提供了多种类型的验证，包括`单个字段`，`struct`验证 和 `slice`、`map`的深入验证

```go
package main

import (
    "fmt"
    "github.com/go-playground/validator/v10"
    "time"
)

type Address struct {
    City  string `json:"city" validate:"required"`
    Phone string `json:"phone" validate:"required,e164"`
}
type User struct {
    FirstName       string         `validate:"required"`
    Age             uint8          `validate:"gte=0,lte=130"`
    Email           string         `validate:"required,email"`
    Hobby           []string       `validate:"required,gte=2,dive"` //至少两个，且深入验证不能是两个空字符
    Address         *Address       `validate:"required,dive"`       //dive深入验证
    ContactUser     []*ContactUser `validate:"required,dive"`       //设置dive深入验证才会验证ContactUser的字段
    FavouriteColor1 string         `validate:"iscolor"`
    FavouriteColor2 string         `validate:"hexcolor|rgb|rgba|hsl|hsla"` // alias for 'hexcolor|rgb|rgba|hsl|hsla'
}
type ContactUser struct {
    Name  string `validate:"required"`
    Phone string `validate:"required_without_all=Email QQ,omitempty,e164"` //Phone QQ Email三者必须有一个，且Phone符合e164格式
    QQ    string `validate:"required_without_all=Phone Email,omitempty"`
    Email string `validate:"required_without_all=Phone QQ"`
}

func main() {
    v := validator.New()
    var b bool
    if err := v.Var(b, "boolean"); err != nil { //验证boolean
        fmt.Println(err)
    }
    var s string = "a"
    if err := v.Var(s, "alpha"); err != nil { //验证字符
        fmt.Println(err)
    }
    var f1 = 100.1
    if err := v.Var(f1, "number"); err != nil { //验证数字
        fmt.Println(err)
    }
    var f2 = "100.1"
    if err := v.Var(f2, "numeric"); err != nil { //验证字符串类型的数字
        fmt.Println(err)
    }
    var slice = []int{1, 2, 3, 4}
    if err := v.Var(slice, "max=10,min=3"); err != nil { //验证切片的长度
        fmt.Println(err)
    }
    var m = map[int]int{1: 1, 2: 2, 3: 2, 4: 4}
    if err := v.Var(m, "max=10,min=3"); err != nil { //验证map的长度
        fmt.Println(err)
    }
    var t = time.Now().Format("2006-01-02 15:04:05")
    if err := v.Var(t, "datetime=2006-01-02 15:04:05"); err != nil { //验证时间格式
        fmt.Println(err)
    }

    var s1, s2 = "abc", "abc"
    if err := v.VarWithValue(s1, s2, "eqcsfield"); err != nil { //验证两个字段是否相同
        fmt.Println(err)
    }
    var i1, i2 = 10, 20
    if err := v.VarWithValue(i1, i2, "ltfield"); err != nil { //ltfield 小于，ltefield 小于等于
        fmt.Println(err)
    }

    var u = User{
        FirstName: "firstname", Age: 18, Email: "iscodd@gmail.com", FavouriteColor1: "#fff", FavouriteColor2: "rgb(0,0,0)",
        Hobby:       []string{"a", "b"},
        Address:     &Address{City: "shanghai", Phone: "+8617091900050"},
        ContactUser: []*ContactUser{{Name: "iscod", QQ: "1000"}},
    }
    if err := v.Struct(u); err != nil { //验证struct类型
        fmt.Println(err)
    }
    //验证map,注意长度校验，dive的应用
    var m1 = map[string]string{"a": "a", "b": "ab", "c": "abc", "d": "abcd"}
    if err := v.Var(m1, "gte=3,dive,keys,len=1,alpha,endkeys,required,gte=1,lte=5"); err != nil { //校验的key的长度，value的字符长度
        fmt.Println(err)
    }
    var m2 = map[string]map[string]string{"a": {"a": "a", "b": "ab", "c": "abc", "d": "abcd"}, "b": {"a": "a", "b": "ab", "c": "abc", "d": "abcd"}}
    if err := v.Var(m2, "gte=2,dive,keys,len=1,alpha,endkeys,required,gte=3,dive,keys,len=1,alpha,endkeys,required,gte=1,lte=5"); err != nil { //校验的key的长度，value的字符长度
        fmt.Println(err)
    }
}

```
#### 6. 配置解决方案viper

1. 读取本地配置
1. 读取远程etcd配置中心文件
1. 读取k8s的configmaps配置

#### 7. Casbin

(Casbin)[https://github.com/casbin/casbin]是一套角色权限模型

1. 基于ACL访问控制模块
1. 基于RBAC模型的访问控制
1. 基于ABAC模型的访问控制

## 分布式中间件

#### 1. redis

#### 2. mysql

#### 3. mongodb

#### 4. etcd

etcd的服务注册与发现的流程一般是：

1. 服务端启动，并连接etcd进行注册，并通过`lease`进行不间断的续租(KeepAlive)，来保证在etcd中值一直存在
2. 客户端启动时，先拉去etcd中得列表，然后通过`watch`来维护本地缓存列表。


```go
//服务端进行注册,并进行KeepAlive，保持租约
func Register(name, target string) {
    c, _ := NewEtcdClient("127.0.0.1:2379")
    res, err := c.Lease.Grant(context.Background(), 10)
    if err != nil {
       fmt.Printf("Error getting leases %s", err)
    }
    _, err = c.Put(context.Background(), name, target, clientv3.WithLease(res.ID))
    if err != nil {
        fmt.Printf("Error Put %s", err)
    }
    ch, err := c.Lease.KeepAlive(context.Background(), res.ID)
    if err != nil {
        fmt.Printf("Error KeepAlive %s", err)
    }

    // 保持租约
    go func(ch <-chan *clientv3.LeaseKeepAliveResponse) {
        for item := range ch { //消费消息，防止消息阻塞
            fmt.Printf("Id: %v, %d\n", item.ID, item.TTL)
        }
    }(ch)
}

//客户端watch, 注意本地缓存map的并发编程锁问题
func clientWatch(kvs map[string]string) {
    ch := etcd.Watch("user-server")
    for item := range ch {
        for _, v := range item.Events {
        switch v.Type {
            case clientv3.EventTypeDelete:
                    //delete(kvs, string(v.Kv.Key))//删除本地缓存
                case clientv3.EventTypePut:
                    //kvs[string(v.Kv.Key)] = string(v.Kv.Value)//加入本地缓存
            }
        }
    }
}
```

#### 5. kafka

(kafka)[https://iscod.github.io/#/amqp/kafka]消息队列解决方案

#### 6. ElasticSearch

#### 7. OpenTelemetry

`OpenTelemetry`主要用于分布式链路追踪

1. jaeger集成
1. zipkin集成
1. prometheus集成


## 云原生

#### 1. git & gitlab

#### 2. Docker

*volumes*

volume有三种常用形式：匿名数据卷、绑定数据卷、具名数据卷

```
docker volume ls # 数据卷列表
docker run -d -v /data --name nginx nginx # 匿名数据卷
docker run -d -v /data:/data --name nginx nginx # 绑定数据卷

docker volume create my-volume # 创建具名数据卷
docker run -d -v my-volume:/data --name nginx nginx # 绑定具名数据卷

# 从nginx挂着的数据卷中备份数据 nginx -> /bak目录下
docker run --rm --volumes-from nginx -v /bak:/bak --name backup nginx cp -r /data /bak
docker run --rm --volumes-from nginx -v /bak:/bak --name backup nginx tar cvf  /bak/data.tar /data # 压缩备份

# 数据还原 /bak/data.tar -> nginx2容器
docker run --rm --volumes-from nginx2 -v /bak:/bak --name backup nginx tar xf  /bak/data.tar -C /data/
```

#### 3. Harbor

(Harbor)[https://goharbor.io/] 是云原生基金会（CNCF）毕业项目，是构建私有仓库的首选，包含权限管理、LDAP、日志审核等

#### 3. kubernetes

(kubernetes)[https://iscod.github.io/#/devops/kubernetes]

#### 4. jenkins

#### istio 服务网格


## 场景化解决方案

#### 1. OAuth2.0授权方案

#### 2. goadmin后台开发

#### 3. goim即时通信

#### 4. 微服务设计模式

1. 聚合器微服务设计模式
1. 代理微服务设计模式
1. 链式微服务设计模式
1. 分支微服务设计模式
1. 异步消息传递微服务设计模式