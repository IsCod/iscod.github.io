# 栈（stack）和队列 (queue)

## 栈（stack） 

* 先进后出

如同一叠纸张，这个队列是垂直的，你将一张纸放到上面，先放的纸肯定是最后拿走的，因为上面的纸挡住了它，所以先进的后出

## 队列 (queue)

* 先进先出

如同买东西时，人太多需要排队，排到前面的人肯定先买到，后到的人在前面人买完后才能买，故是先进先出

`栈（stack）`和`队列 (queue)`可以通过`链表` `数组` 结构实现

## 代码实现

```go
package main

import (
	"fmt"
	"sync"
)

type stack struct {
	arr  []*interface{}
	len  int
	lock sync.Mutex
}

func (s *stack) Push(v interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.arr = append(s.arr, &v)
	s.len++
}

func (s *stack) Pop() interface{} {
	if s.len < 1 {
		return nil
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	v := s.arr[s.len-1]
	s.len--
	//构建新数组，减少cap
	newArr := make([]*interface{}, s.len)
	copy(newArr, s.arr[:s.len])
	s.arr = newArr
	return *v
}

func main() {
	var l = &stack{}
	l.Push(1)
	l.Push(10)
	l.Push(3)
	l.Push(2)
	v := l.Pop()
	for v != nil {
		fmt.Printf("%d\n", v)
		v = l.Pop()
	}
}
``` 