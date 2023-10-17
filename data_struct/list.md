# 链表

链表由一个个数据节点组成，链表将数据和数据之间关联起来，从一个数据节点指向另一个数据节点。

## 链表节点

常用链表节点结构包含指向上一个(prev)和下一个节点(next)的指针，和该节点值的value

```c
//redis中链表节点
typedef struct listNode {
    struct listNode *prev;
    struct listNode *next;
    void *value;
} listNode;
```

## 链表

链表结构一般包含头部节点(head)，末尾节点(tail)的指针和链表长度(len)

```c
//redis中链表结构
//redis的链表是双端链表，链表节点带有前置(prev)和后置节点(next)两个指针
typedef struct list {
    listNode *head;
    listNode *tail;
    void *(*dup)(void *ptr);
    void (*free)(void *ptr);
    int (*match)(void *ptr, void *key);
    unsigned long len;
} list;
```

```go
// go 官方包也实现了简单的链表结构, container/list
type Element struct { //链表节点，节点是双端节点
	next, prev *Element
	list *List
	Value any
}

type List struct { //链表节点
	root Element // sentinel list element, only &root, root.prev, and root.next are used
	len  int     // current list length excluding (this) sentinel element
}
```

链表一般分为`单链表` `双链表` `循环链表`这些链表都是一个数据指向另外一个数据的结构只是链表节点的指向不同

- `单链表`
    单链表，就是链表是单向的，每个节点只有一个方向不能往回找
- `双端链表`
    顾名思义，就是每个节点可以找到它之前的节点和之后的，节点是双向的
- `循环链表`
    循环链表，就是它可以一直向下，最后回到自己的那个节点，形成了一个回路。循环链表分为单链表和双链表，既一个是只能顺着一个方向走，另外一个是两个方向走

## 代码实现

```go
//使用golang官方包内实现的list
package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	node1 := l.PushFront(10)
	l.InsertBefore(11, node1)
	l.InsertBefore(12, node1)
	l.InsertAfter(13, node1)

	for b := l.Front(); b != nil; b = b.Next() {
		fmt.Println(b.Value)
	}
}
```

```go
// 自己实现list结构
package main

import "fmt"

// 可以看到与官方包中list.Element结构类似
type linkNode struct {
	Prev, Next *linkNode
	Value      interface{}
}

type link struct {
	Head, Tail *linkNode
	Len        int
}

func (l *link) Push(node *linkNode) {
	n := &linkNode{Value: node.Value}
	if l.Len == 0 {
		n.Next = l.Tail
		n.Prev = l.Head
		l.Head, l.Tail = n, n
	} else {
		tail := l.Tail
		l.Tail.Next, l.Tail = n, n
		l.Tail.Prev = tail
	}
	l.Len++
}

func (l *link) Pop() *linkNode {
	if l.Len < 1 {
		return nil
	}
	head := l.Head
	if head.Next != nil {
		head.Next.Prev = nil
	}
	l.Head = head.Next
	l.Len--
	return head
}

func main() {
	var l = &link{}
	l.Push(&linkNode{Value: 1})
	l.Push(&linkNode{Value: 10})
	l.Push(&linkNode{Value: 3})
	l.Push(&linkNode{Value: 2})
	fmt.Printf("List len: %d\n", l.Len)
	head := l.Head
	for i := 0; i < l.Len; i++ {
		fmt.Printf("Next Scan: %d\n", head.Value)
		head = head.Next
	}
	tail := l.Tail
	for i := 0; i < l.Len; i++ {
		fmt.Printf("Prev Scan: %d\n", tail.Value)
		tail = tail.Prev
	}

	for l.Len > 0 {
		n := l.Pop()
		fmt.Printf("Pop Scan: %d, %d \n", n.Value, l.Len)
	}
}
```
* 参考
    * [redis设计与实现](https://www.bookstack.cn/read/redisbook/f8fc3415f3c66c78.md)
    * [数据结构(golang)](https://www.bookstack.cn/read/hunterhug-goa.c/algorithm-link.md)