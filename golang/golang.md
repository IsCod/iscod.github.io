
## 值类型与引用类型

 哪些是引用类型？

`map`, `slice`, `interface`, `chan`, `pointer` 

哪些是值类型？

`array`, `struct` ,`int`, `float`, `bool`, `string`

在函数中传参，如果是值类型传递，修改被传递的参数，函数内修改不会影响原有变量. 如果是引用类型传递，函数内修改参数会影响原值

> 引用类型初始化时，都可赋值为`nil`, 而值类型不能初始化为`nil`

```go
func change(slice []int) {
	for k, v := range slice {
		slice[k] = v + 10
	}
}

func main() {
	var slice = make([]int, 0)
	for i := 0; i < 10; i++ {
		slice = append(slice, i)
	}
	slice1 := slice
	change(slice)
	fmt.Println("slice:", slice)
	fmt.Println("slice1:", slice1)
}
```

输出结果：

```sh
# go run main.go
slice: [10 11 12 13 14 15 16 17 18 19]
slice1: [10 11 12 13 14 15 16 17 18 19]
```

可以看到引用类型，两个变量都发生了修改

> Go函数参数中都是值传递，只是引用类型形参和实参指向的内存地址相同。如果想获取一个互不影响的值，可以使用深`copy`


## panic 与 defer 谁先执行？

先说结论：`panic`相当于一个`return`。所以在函数`panic`前执行`defer`

```go
func main() {
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	panic("err panic")
}
```

输出结果：

```
 # go run main.go
defer 2
defer 1
panic: err panic

goroutine 1 [running]:
main.main()
        /Users/ning/Data/xz_server/src/ningserver/main.go:126 +0xc5
exit status 2
```

## goroutine 需要错误处理吗？

在一个程序中存在多个`goroutine`。如果不进行错误处理，当一个`goroutine`发生`panic`会造成整个程序崩溃。所以为了保证程序还能正常运行，一般在子`goroutine`可以使用`recover`捕获异常，进而对异常处理

```go
func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		panic("err panic")
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			fmt.Printf("%d,", i)
		}
	}()
	wg.Wait()
}

```

输出结果：

```
# go run main.go
err panic
0,1,2,3,4,5,6,7,8,9,%       
```


##  有趣的`defer`返回值

```go
func def1(i int) (t int) {
	t = i
	defer func() { t += 3 }()
	return t
}

func def2(i int) int {
	t := i
	defer func() { t += 3 }()
	return t
}

func def3(i int) (t int) {
	defer func() { t += i }()
	return 2
}
func main() {
	fmt.Println(def1(1), def2(1), def3(1))
}

```

执行结果：
```sh
# go run main.go
4,1,3
```

## iota

```go
const (
	a = iota
	b
	c = "zz"
	d
	e
	f = iota
	g
)

func main() {
	fmt.Println(a, b, c, d, e, f, g)

}

执行结果：
```sh
# go run main.go
0 1 zz zz zz 5 6
```

## unsafe.Pointer

`unsafe`包提供了两个重要的能力：

1. 任何类型的指针和`unsafe.Pointer`可以相互转换
2. `uintptr`类型和`unsafe.Pointer`可以相互转换

#### `unsafe.Pointer`实现更改私有成员值

```go
type User struct {
	name  string
	age   int
	score int
}

func (u *User) SetName(n string) {
	u.name = n
}

func main() {
	user := User{}
	user.SetName("iscod")
	fmt.Println(user)
	name := (*string)(unsafe.Pointer(&user))
	*name = "ascoon"
	age := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&user)) + unsafe.Sizeof(user.name)))
	*age = 18
	score := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&user)) + unsafe.Offsetof(user.score)))
	*score = 1
	fmt.Println(user)
}
```

执行结果：
```sh
# go run main.go
{iscod 0 0}
{ascoon 18 1}
```

> `unsafe.Sizeof`是返回字段大小，`unsafe.Offsetof`返回字段在结构内的偏移量

#### `unsafe.Pointer`实现`string`与`[]byte`无copy转换

```go
func StringToByte(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

func ByteToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
```

## sync

`sync`包提供了三个常用的功能:

1. `sync.WaitGroup` 等待组常用于保证在并发环境中完成指定数量的任务
1. `sync.Once` 可以保证函数只执行一次的实现，比如加载配置文件场景，且是并发安全的
1. `sync.Cond` 条件变量用来协调想要访问共享资源的那些 goroutine，当共享资源的状态发生变化的时候，它可以用来通知被互斥锁阻塞的 goroutine


示例：

```go
//sync.Once
func main() {
	wg := sync.WaitGroup{}
	once := sync.Once{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Printf("%d", i)
			once.Do(func() {
				fmt.Println("once one")
			})
		}
	}()
	wg.Wait()
}

```


```go
func main() {
	//sync.Cond
	wg := sync.WaitGroup{}
	var b bool
	cond := sync.NewCond(&sync.RWMutex{})
	wg.Add(1)
	go func() {
		defer wg.Done()
		cond.L.Lock()
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
		}
		b = true
		cond.L.Unlock()
		cond.Broadcast()
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("run wait 1")
		cond.L.Lock()
		for !b {
			cond.Wait()
		}
		fmt.Println("run 1")
		cond.L.Unlock()
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("run wait 2")
		cond.L.Lock()
		for !b {
			cond.Wait()
		}
		fmt.Println("run 2")
		cond.L.Unlock()
	}()

	wg.Wait()
}
```

## 内存回收

Golang的垃圾回收是通过`GC`来实现的。

可以通过手动和自动方式进行垃圾回收。自动垃圾回收

```go
runtime.GC()//程序中手动调用GC
```

* `GC`采用的是`三色标记`算法来实现垃圾回收. 

`三色标记`可以理解成将一个内存标记为三种状态

1. 白色对象 — 潜在的垃圾，其内存可能会被垃圾收集器回收
1. 黑色对象 — 活跃的对象，包括不存在任何引用外部指针的对象 以及从根对象可达的对象
1. 灰色对象 — 活跃的对象，因为存在指向白色对象的外部指针，垃圾收集器会扫描这些对象的子对象

* 垃圾收回的过程

1. 初始状态所有对象都是白色
1. 从根节点开始变了所有对象, 把遍历对象标记为`灰色对象`
1. 遍历`灰色对象`, 将灰色对象引用的对象也变成`灰色对象`, 然后将遍历的`灰色对象`标记为`黑色对象`
1. 循环步骤 3, 直到`灰色对象`全部变为`黑色对象`

当三色标记工作完成后, 应用程序中就不存在任何`灰色对象`。我们只能看到`黑色对象`和`白色对象`. 垃圾回收器就可以回收这些`白色对象`垃圾.

> 内存回收有多种算法如: 引用标记, 标记清除, 分代收集等 


* 参考
	* [Golang调度器](https://studygolang.com/articles/9610)
	* [Golang并发模型](https://www.jianshu.com/p/f9024e250ac6)
	* [redis引用计数内存回收](https://iscod.github.io/#/nosql/redis?id=%e5%86%85%e5%ad%98%e5%9b%9e%e6%94%b6)
