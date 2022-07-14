# 切片

### 类型定义

```go
// src/runtime/slice.go
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}
```

可以看到切片包含了三个参数: `len` : 长度, `cap` : 容量, `array` : 存储数据

### 切片容量是如何增长的？

一般在向`slice`追加(`append`函数)元素之后，才会引起扩容操作。

先看看`append`函数的原型

```go
//	slice = append(slice, elem1, elem2)
//	slice = append(slice, anotherSlice...)
func append(slice []Type, elems ...Type) []Type
```

可以看到`append`函数的参数长度可变，因此追加多个值，也可使用`...`传入一个`slice`类型

### 切片的扩容规则

关于`golang`切片扩容规则你听到最多的一句话是：

当原 slice 长度小于 1024 的时候，容量会每次增加 1 倍。当`slice`的长度大于1024的时候，扩容会以`1.25`倍扩容

这就话对，但是也不对。在最早的版本确实是这样的，我们看下早期的扩容版本函数：

```go
//`2014-08-01`由`Randall`提交的扩容规则函数
func growslice(t *slicetype, old sliceStruct, n int64) sliceStruct {
	//....
	newcap := old.cap
	if newcap+newcap < cap {
		newcap = cap
	} else {
		for {
			if old.len < 1024 {
				newcap += newcap
			} else {
				newcap += newcap / 4
			}
			if newcap >= cap {
				break
			}
		}
	}
	//....
}
```

可以看到该版本确实是按上诉所说进行的扩容。

-----

但是在`2021-09-08` `Randall`更新了gorwslice函数的扩容规则，改用了更平滑的扩容规则


```go
// Instead of growing 2x for < 1024 elements and 1.25x for >= 1024 elements,
// use a somewhat smoother formula for the growth factor. Start reducing
// the growth factor after 256 elements, but slowly.

// starting cap    growth factor
// 256             2.0
// 512             1.63
// 1024            1.44
// 2048            1.35
// 4096            1.30

func growslice(et *_type, old slice, cap int) slice {
	//...
	newcap := old.cap
	doublecap := newcap + newcap
	if cap > doublecap {
		newcap = cap
	} else {
		const threshold = 256
		if old.cap < threshold {
			newcap = doublecap
		} else {
			// Check 0 < newcap to detect overflow
			// and prevent an infinite loop.
			for 0 < newcap && newcap < cap {
				// Transition from growing 2x for small slices
				// to growing 1.25x for large slices. This formula
				// gives a smooth-ish transition between the two.
				newcap += (newcap + 3*threshold) / 4
			}
			// Set newcap to the requested cap when
			// the newcap calculation overflowed.
			if newcap <= 0 {
				newcap = cap
			}
		}
	}
	//....
}
```