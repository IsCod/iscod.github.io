# TCP/IP

## 分层

* OSI7层参考模型和TCP/IP五层模型

<table>
    <tr>
        <th>OSI七层参考模型</th>
        <th>TCP/IP五层模型</th>
        <th>使用协议</th>
        <th>描述</th>
    </tr>
    <tr>
        <td>应用层</td>
        <td rowspan="3">应用层</td>
        <td>HTTP, HTP, SSH</td>
        <td>用户空间</td>
    </tr>
    <tr>
        <td>表示层</td>
        <td>NCP, XDR</td>
        <td></td>
    </tr>
    <tr>
        <td>会话层</td>
        <td>sockets, BSD</td>
        <td></td>
    </tr>
     <tr>
        <td>传输层</td>
        <td>传输层</td>
        <td>TCP,UDP</td>
        <td></td>
    </tr>
     <tr>
        <td>网络层</td>
        <td>网络层</td>
        <td>IP, ICMP</td>
        <td></td>
    </tr>
     <tr>
        <td>数据链路层</td>
        <td>数据链路层</td>
        <td>IEEE 802.11,FDDI</td>
        <td>数据链路层实现了网卡接口的网络驱动程序, 他们实现了IP地址和机器物理地址(通常是MAC地址, 802.1无线网络,以太网都使用MAC地址)之间的相互转换</td>
    </tr>
     <tr>
        <td>物理层</td>
        <td>物理层</td>
        <td>无线电,光纤</td>
        <td>物理传输媒介</td>
    </tr>
</table>

```
TCP (传输控制协议) - TCP 用于从应用程序到网络的数据传输控制。它负责在数据传输到达之前将它们分割为`IP`包,然后在它们到达后进行重组
UDP (用户数据包协议) - 应用程序之间的简单通信
IP (网际协议) - 计算机之间的通信
ICMP (因特网消息控制协议) - 针对错误和状态
DHCP (动态主机配置协议) - 针对动态寻址
```

## 连接方式

#### 建立连接(三次握手)
 
建立连接三次握手的基本思想是`让我知道你已经知道`。这种约定俗成的可靠信息交互的基本方式在我们的生活非常常见，经典的例子就是投递简历：

```
1, 你(客户端)投递简历给某公司(服务端)，
2, 然后公司(服务端)回复你(客户端)确认收到你的简历
3, 你(客户端)收到这个通知后一般需要回复公司(服务端)我已知道
```

这三个信息交互过程实现了你(客户端)和公司(服务端)双方`让我知道你已经知道`信息表达。
这种思想应用到信息交互框架就叫做协议，而TCP协议就是这种类型的具体实现

**TCP链接三次握手:**

![TCP三次握手](https://iscod.github.io/images/tcp_1.png)

- 1, 客户端向服务端发送建立链接请求, 此时Flags标识位为：`SYN`。附带参数为`seq`(tcp.seq_raw=4272455158)随机序号值

![TCP客户端发起请求链接参数](https://iscod.github.io/images/tcp_2.png)

- 2, 服务端确认收到客户端协议后向客户端发送确认信息, 此时Flags标识位为：`SYN`+ `ACK`。并附带参数`ack`（tcp.ack_raw=客户端发送的tcp.seq_raw+1=4272455159）和 新的`seq`(tcp.seq_raw=317767926)随机序号

![服务端返回ACK确认信息](https://iscod.github.io/images/tcp_3.png)

- 3，客户端确认收到服务端协议后向服务器发送确认信息，此时Flags标识位为：`ACK`。附带参数`ack`（tcp.ack_raw=服务端seq+1=317767927）和`seq`(tcp.seq_raw=服务端发送的tcp.ack_raw=4272455159)

![客户端返回ACK确认信息](https://iscod.github.io/images/tcp_4.png)

至此链接建立完毕, 就可以发送数据

![客户端发发送信息](https://iscod.github.io/images/tcp_5.png)


#### 断开连接(四次挥手)

断开连接四次挥手的与三次握手可以用同样的`让我知道你已经知道`思想来理解，我们依然用投递简历来讲解

```
1, 你(客户端)通知公司(服务端)因为某些问题不太满意，要回退简历
2, 公司(服务端)先回复你(客户端)确认收到你(客户端)的请求，但是需要内部协商(既TCP通讯中需要发送完未发送的数据)后再正式通知
3, 公司(服务端)内部协商后向你(客户端)退回简历，(既TCP通信中所有数据发送完后下向客户端发送同意关闭协议）
4, 你(客户端)收到公司退回的简历通知后，回复公司(服务端)我已知道，但是我会等待一段时间关闭，因为我不确定公司(服务端)是否真正关闭
```

客户端发送第`4`步骤的信息进入`等待(TIME-WAIT)`状态，如果在`等待(TIME-WAIT)`时间结束后没有收到服务器新的信息则进入`关闭(CLOSE)`。

服务端在接收到客户端第`4`步骤的信息后进入`关闭(CLOSE)`。

**TCP链接四次挥手:**

![TCP四次挥手](https://iscod.github.io/images/tcp_fin_1.png)

- 1, 客户端向服务端发送断开链接请求, 此时Flags标识位为：`FIN`+ `ACK`。附带参数：`seq`(tcp.seq_raw=969185913)随机序号值

![TCP客户端发起断开链接参数](https://iscod.github.io/images/tcp_fin_2.png)

- 2, 服务端回复收到请求, 此时Flags标识位为：`ACK`。附带参数：`ack`（tcp.ack_raw=客户端发送的tcp.seq_raw+1）和 新的`seq`(tcp.seq_raw=3646450711)随机序号

![TCP客户端发起断开链接参数](https://iscod.github.io/images/tcp_fin_3.png)

- 3, 服务端发送断开请求, 此时Flags标识位为：`FIN`+ `ACK`。附带参数：`ack`和`seq`参数与上一步骤相同

![TCP客户端发起断开链接参数](https://iscod.github.io/images/tcp_fin_4.png)

- 4, 客户端回复确认收到断开, 此时Flags标识位为：`ACK`。附带参数：`ack`（tcp.ack_raw=服务端发送的tcp.seq_raw+1）

![TCP客户端发起断开链接参数](https://iscod.github.io/images/tcp_fin_5.png)


### 代码实现

go scoket 服务器：
```go
package main

import (
    "bytes"
    "fmt"
    "net"
    "strings"
    "time"
)

func main() {
    l, err := net.Listen("tcp4", ":1234")
    if err != nil {
        fmt.Printf("net listen err : %s\n", err)
    }

    for {
        c, err := l.Accept()
        if err != nil {
            fmt.Printf("accept err : %s", err)
            continue
        }

        go func() {
            fmt.Printf("client connection: , %s\n", c.RemoteAddr())
            _, err = c.Write([]byte(fmt.Sprintf("server connection success: %s\n", c.LocalAddr())))
            c.SetDeadline(time.Now().Add(time.Second * 10))
            var req = make([]byte, 1024)
            for {
                n, err := c.Read(req)
                if err == nil {
                    c.SetDeadline(time.Now().Add(time.Second * 10))
                    if n > 0 {
                        fmt.Printf("conn read : %s", req)
                    }

                    _, err = c.Write([]byte(fmt.Sprintf("server return: %s\n", req)))
                    if err != nil {
                        fmt.Printf("err : %s", err)
                    }
                    if strings.TrimSpace(fmt.Sprintf("%s", bytes.TrimSpace(req[:n]))) == "close" {
                        c.Close()
                    } else {
                        req = make([]byte, 1024)
                    }
                }

            }
        }()
    }
}
```

测试

```bash
$ nc 127.0.0.1 1234
server connection success: 127.0.0.1:1234

close
server return: close
```

* 工具
    * [wireshark](https://www.wireshark.org/)
    * [telnet]()
    * [nc]()


