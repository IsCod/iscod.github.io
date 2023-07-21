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

### 采样分析类型

 采样类型 | 描述   |
-------- | ----- |
allocs | 内存总分配
block | 同步原语导致的阻塞
cmdline | 进程启动命令行
goroutine | 当前所有的 goroutine
heap | 内存分析
mutex | 互斥锁
profile | CPU
threadcreate | 操作系统中的线程创建
trace | 程序执行 trace, 和其他样本数据不同的是，这个需要使用 go tool trace 来分析
full goroutine stack dump | 打印所有 goroutine 的堆栈

### 内存分析

```
// 获取内存中存留对象数据
$ go tool pprof -inuse_objects http://localhost:6060/debug/pprof/heap
```

可以获取内存分析采样类型`options`：
```
-inuse_space           程序实际占用内存
-inuse_objects         内存中存留的对象数量
-alloc_space           程序分配内存（如果该指标明显大于inuse_space指标，说明存在内存分配尖峰）
-alloc_objects         内存中分配的对象数量（如果该指标明显大于inuse_objects指标说明存在内存分配尖峰）
```

> 排除内存相关问题，除了关注内存使用和闲置相关参数，还要关注GC频率和GC时间等参数。

### 生成火焰图

1. 采样数据

获取最近10秒程序运行的cpuprofile,-seconds参数不填默认为30。

```bash
# pprof采样profile数据，时间10秒
go tool pprof -seconds 10 http://127.0.0.1:8080/debug/pprof/profile
```

等10s后会生成一个: pprof.samples.cpu.001.pb.gz文件

![flamegraph](https://iscod.github.io/images/pprof1.png)

2. 分析数据


3. 火焰图

```bash
go tool pprof -http=:8081 ~/pprof/pprof.samples.cpu.001.pb.gz
```

其中-http=:8081会启动一个http服务,端口为8081,然后浏览器会弹出此文件的图解:

![flamegraph](https://iscod.github.io/images/flamegraph1.png)

## trace

1. 采样数据

```bash
# trace采样，时间10秒
# trace日志是一个压缩的protobuf文件，需要使用`go tool trace`进行分析
curl -sS http://127.0.0.1:8080/debug/pprof/trace?seconds=10 -o trace.out
```

2. 分析数据

```
# 使用`go tool trace`进行分析下载的`trace.out`文件
go tool trace trace.out
```

随后浏览器将打开trace webUi。
在trace工具的界面上，可以看到程序的整体运行情况、goroutine的执行情况、系统调用的时间等信息。可以根据需要对程序的性能问题进行分析和优化。

![traceWebUi](https://iscod.github.io/images/tracewebui.png)



