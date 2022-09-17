# 并发不是并行

理解并发与并行的区别首先想到的是如下的程序执行图：

![并发&并行](https://iscod.github.io/images/cp1.png)

> 并发是同时处理很多事情（事件是交叉处理的）。并行是同时做很多事情（事情是并行处理的）

### goroutines

goroutines是并发执行的关键

### channels

channels实现同步与消息的传递

### select 

多路并发控制

## 并行




思考如下程序的输出

```go
func Num(proc int) int {
	runtime.GOMAXPROCS(proc)
	var count int
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			count += 1
			wg.Done()
		}()
	}
	wg.Wait()
	return count
}

func main() {
	fmt.Println(Num(1), Num(2)) //输出: 1000 985
}
```