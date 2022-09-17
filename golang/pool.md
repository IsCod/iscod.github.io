# sync.pool

sync.pool是Go1.3发布的一个特性，它是一个临时对象存储池

### 构建连接池

1. 构建`Client interface`
2. 在服务端构建`GetClient`

```go
// sync.poll构建连接池
package client

import (
    "fmt"
    "google.golang.org/grpc"
    "google.golang.org/grpc/connectivity"
    "sync"
)

type Client interface {
    Get() *grpc.ClientConn
    Put(*grpc.ClientConn)
}

func NewClient(target string, opts ...grpc.DialOption) Client {
    return &clientPool{
        sync.Pool{New: func() any {
            conn, err := grpc.Dial(target, opts...)
            if err != nil {
                fmt.Println("error: getting client pool", err)
            }
            fmt.Printf("new client pool traget %s", target)
            return conn
        }},
    }
}

type clientPool struct {
    sync.Pool
}

func (c *clientPool) Get() *grpc.ClientConn {
    client := c.Pool.Get().(*grpc.ClientConn)
    if client == nil || client.GetState() == connectivity.Shutdown || client.GetState() == connectivity.TransientFailure {
        fmt.Printf("error: getting client pool %s", client.GetState())
        client.Close()
        client = c.Pool.New().(*grpc.ClientConn)
    }
    return client
}

func (c *clientPool) Put(client *grpc.ClientConn) {
    if client == nil || client.GetState() == connectivity.Shutdown || client.GetState() == connectivity.TransientFailure {
        fmt.Printf("error: puting client pool %s", client.GetState())
        client.Close()
        return
    }
    c.Pool.Put(client)
}
```

每个`service`实现自身的`client`为使用者提供抽象层

```go
// service构建 NewClient,引用上一步创建的 NewClient方法
var pool client.Client
var once = sync.Once{}

func GetClient() client.Client {
    once.Do(func() {//采用once保证初始化只执行一下
        pool = client.NewClient("127.0.0.1:5000", grpc.WithTransportCredentials(insecure.NewCredentials()))
    })
    return pool
}
```

web客户端如何使用？

```go
// web客户端中使用连接
svc := server.GetClient()
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    c := svc.Get()
    defer svc.Put(c) //使用完put回去
    //s, err := client.GetName(pb.NewUserClient(c)) //new server client coding...
})
```
