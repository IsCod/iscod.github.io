# 性能分析

## pprof


```go

package main

import (
	"net/http"
	_ "net/http/pprof"
)

func main {
	go func() {
		http.ListenAndServe("127.0.0.1:8080", nil)
	}()
	//you coding...
}
```

然后运行程序后,直接访问: `http://127.0.01:8080/debug/pprof/`,可以看到go运行的信息:

```html
/debug/pprof/

Types of profiles available:
Count	Profile
360	allocs
0	block
0	cmdline
45	goroutine
360	heap
0	mutex
0	profile
15	threadcreate
0	trace
full goroutine stack dump
Profile Descriptions:

allocs: A sampling of all past memory allocations
block: Stack traces that led to blocking on synchronization primitives
cmdline: The command line invocation of the current program
goroutine: Stack traces of all current goroutines
heap: A sampling of memory allocations of live objects. You can specify the gc GET parameter to run GC before taking the heap sample.
mutex: Stack traces of holders of contended mutexes
profile: CPU profile. You can specify the duration in the seconds GET parameter. After you get the profile file, use the go tool pprof command to investigate the profile.
threadcreate: Stack traces that led to the creation of new OS threads
trace: A trace of execution of the current program. You can specify the duration in the seconds GET parameter. After you get the trace file, use the go tool trace command to investigate the trace.
```

### 生成火焰图

1. 获取cpuprofile

获取最近10秒程序运行的cpuprofile,-seconds参数不填默认为30。

```bash
go tool pprof -seconds 10 http://127.0.0.1:8080/debug/pprof/profile
```

等10s后会生成一个: pprof.samples.cpu.001.pb.gz文件

![flamegraph](https://iscod.github.io/images/pprof1.png)

2. 生成火焰图

```bash
go tool pprof -http=:8081 ~/pprof/pprof.samples.cpu.001.pb.gz
```

其中-http=:8081会启动一个http服务,端口为8081,然后浏览器会弹出此文件的图解:

![flamegraph](https://iscod.github.io/images/flamegraph1.png)

## trace

