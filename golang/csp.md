# CSP通讯

Golang核心的一个概念是*不要使用共享内存方式进行通讯，而是使用通讯共享内存*

CSP通讯的内容

## channel

### 读写已关闭的chan，会发生什么

先说结论：

1. 写已关闭的`chan`会发生`panic`
1. 读已关闭的`chan`能一直读到东西，但是读到的内容会根据 通道是否关闭和是否有元素 而不同
	* 如果`chan`关闭, `buffer`内有元素还未读, 会正确读到`chan`内的值, 且第二个返回值为`true`
	* 如果`chan`关闭, `buffer`内无元素, 接下来所有的读都会非阻塞直接成功返回，但是`channel`返回的元素值是`0`. 且第二个返回值一直为`false`

### 为什么是这样？

想要了解深层的机制，那么阅读源码是最好的方式。接下来我们一步步分析为什么是这样。

首先了解一下`chan`结构体

```go
type hchan struct {
	qcount   uint           // total data in the queue
	dataqsiz uint           // size of the circular queue
	buf      unsafe.Pointer // points to an array of dataqsiz elements
	elemsize uint16
	closed   uint32         //是否关闭标识符
	//...
}
```

### `close`函数做了什么？

```go
func closechan(c *hchan) {
	//...
	c.closed = 1
	//...
}
```

可以看到`close`函数将`closed`标识符设为`1`

---------------


### 写的时候怎么判断`closed`标识符的？

```go
//runtime/chan.go
//go chan写
func chansend(c *hchan, ep unsafe.Pointer, block bool, callerpc uintptr) bool {
	//...
	if c.closed != 0 {
		unlock(&c.lock)
		panic(plainError("send on closed channel"))
	}
	//...
}
```

`c.closed != 0`直接`panic`

------------


### 读的时候怎么判断`closed`标识符的？

```go
//runtime/chan.go
func chanrecv(c *hchan, ep unsafe.Pointer, block bool) (selected, received bool) {
	//...
	if c.closed != 0 && c.qcount == 0 {
		if raceenabled {
			raceacquire(c.raceaddr())
		}
		unlock(&c.lock)
		if ep != nil {
			typedmemclr(c.elemtype, ep)
		}
		return true, false
	}
	//...
}
```

读函数仅在`c.closed != 0` 和 `c.qcount == 0`时返回false。即：当`buffer`内无元素时返回元素值0，和第二返回值false



