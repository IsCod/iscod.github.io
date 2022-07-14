# 树

树是一种比较高级的基础数据结构，由多个有限节点组成的具有层次关系的集合

树的定义:

    1.有节点间的层次关系，分为父节点和子节点
    2.有唯一一个根节点，该节点没有父节点
    3.除了根节点，每个节点都有一个父节点
    4.每个节点本身以及它的子节点都是一颗树，是一个递归结构
    5.没有后代的节点称为叶子节点，没有节点的树称为空树

根据子节点的多寡，有二叉树、三叉树、四叉树等。

## 二叉树

二叉树：每个节点最多只有两个子节点的树

满二叉树：叶子节点与叶子节点之前的高度差为`0`的二叉树，既整颗树都是满的

完全二叉树：完全二叉树是满二叉树引申出来的，设二叉树的深度为`K`，除第`K`层外，其它各层的节点数都达到最大值，且第k层所有节点都连续集中在最左边

### 二叉树实现

二叉树节点：

```go
type treeNode struct {
	Value       interface{}
	Times       int64
	Left, Right *treeNode
	Lock sync.Mutex
}
```

### 遍历二叉树

遍历整个二叉树有四种方法：

    1.先序遍历：先访问根节点，在访问左节点，最后访问右节点（也可先访问根节点，在访问右节点，最后访问左节点）
    2.后序遍历：先访问左节点，后访问右节点，最后访问根节点（也可先访问右节点，后访问左节点，最后访问根节点一致）
    3.中序遍历：先访问左节点，后访问根节点，最后访问右节点（也可先访问右节点，后访问根节点，最后访问左节点）
    4.层次遍历：从左向右顺序访问每一层节点

```go
package main

import (
	"fmt"
	"sync"
)

type Tree struct {
	Root *treeNode
}
type treeNode struct {
	Value       interface{}
	Times       int64
	Left, Right *treeNode
	Lock        sync.Mutex
}

func (t *treeNode) Add(value interface{}) {
	t.Lock.Lock()
	defer t.Lock.Unlock()
	if t.Value.(int) > value.(int) {
		if t.Left != nil {
			t.Left.Add(value)
		} else {
			t.Left = &treeNode{Value: value}
		}
	} else if t.Value.(int) < value.(int) {
		if t.Right != nil {
			t.Right.Add(value)
		} else {
			t.Right = &treeNode{Value: value}
		}
	} else {
		t.Times++
	}
	return
}

func (t *Tree) Add(value interface{}) {
	if t.Root == nil {
		t.Root = &treeNode{Value: value}
		return
	} else {
		t.Root.Add(value)
	}
}

//先序
func (tree *treeNode) PreTree() {
	if tree == nil {
		return
	}
	fmt.Printf("%d ", tree.Value)
	tree.Left.PreTree()
	tree.Right.PreTree()
}

//后序
func (tree *treeNode) LastTree() {
	if tree == nil {
		return
	}

	tree.Left.LastTree()
	tree.Right.LastTree()
	fmt.Printf("%d ", tree.Value)
}

//中序
func (tree *treeNode) MidTree() {
	if tree == nil {
		return
	}

	tree.Left.MidTree()
	fmt.Printf("%d ", tree.Value)
	tree.Right.MidTree()
}

type linkNode struct {
	Next  *linkNode
	Value interface{}
}

func main() {
	var tree = &Tree{}
	tree.Add(3)
	tree.Add(1)
	tree.Add(2)
	tree.Add(6)
	tree.Add(4)
	tree.Add(7)
	tree.Add(5)
	fmt.Printf("先序：")
	tree.Root.PreTree()
	fmt.Printf("\n")
	fmt.Printf("后序：")
	tree.Root.LastTree()
	fmt.Printf("\n")
	fmt.Printf("中序：")
	tree.Root.MidTree()
}
```

输出：
```
先序：3 1 2 6 4 5 7 
后序：2 1 5 4 7 6 3 
中序：1 2 3 4 5 6 7
```

## 二叉查找树

二叉树查找是有特定规律的二叉树，具有以下特点

    1.它是一个二叉树，或者空树
    2.它的左子树节点值都小于它的父节点，右子树所有节点的值都大于的父节点值
    3.左右子树也是一个二叉树

以上特点可以保证二叉树一直向左查找可以找到最小的元素，一直向右可以找到最大元素

```go
type SearchTree struct {
	Root *SearchTreeNode
}

type SearchTreeNode struct {
	Value       interface{}
	Times       int64
	Left, Right *SearchTreeNode
	Lock sync.Mutex
}

func NewSearchTree() *SearchTree {
    return new(SearchTree)
}
```

## 2-3树和2-3-4树等

### `2-3`树

2-3树是一种严格子的自平衡的多路查找树。由1986年图灵奖得主，美国计算机科学家 John Edward Hopcroft 发明，又称`3阶B树`（B为Balance既平衡的意思）

三叉树有一下特征

    1.内部节点要么有一元素和两个子节点，要么有两个元素和三个子节点，叶子节点没有孩子，但是有1或者两个元素
    2.所有叶子节点到根节点的长度一致
    3.每个节点的数据元素保持从小到大排序，两个元素之间的子树的所有值大小介于两个元素质检

由于`2-3`树的第二个特征，它是一个完全平衡的数，除了叶子节点，其它节点都没有空儿子，所有树的高度相对较小

> [B+Trees模拟器](https://www.cs.usfca.edu/~galles/visualization/BPlusTree.html)

### 2-3-4树


## 代码实现

### 二叉树

```go
package main

import (
	"sync"
)

type Tree struct {
	Root *treeNode
}

type treeNode struct {
	Value       interface{}
	Times       int64
	Left, Right *treeNode
	Lock sync.Mutex
}

func (t *treeNode) Add(value interface{}) {
	t.Lock.Lock()
	defer t.Lock.Unlock()
	if t.Value.(int) < value.(int) {
		if t.Left != nil {
			t.Left.Add(value)
		} else {
			t.Left = &treeNode{Value: value}
		}
	} else if t.Value.(int) > value.(int) {
		if t.Right != nil {
			t.Right.Add(value)
		} else {
			t.Right = &treeNode{Value: value}
		}
	} else {
		t.Times++
	}
	return
}

func (t *Tree) Add(value interface{}) {
	if t.Root == nil {
		t.Root = &treeNode{Value: value}
		return
	} else {
		t.Root.Add(value)
	}
}
```

### 二叉查找树

```go
package main

import (
	"sync"
)

type SearchTree struct {
	Root *SearchTreeNode
}

type SearchTreeNode struct {
	Value       interface{}
	Times       int64
	Left, Right *SearchTreeNode
	Lock sync.Mutex
}

func NewSearchTree() *SearchTree {
    return new(SearchTree)
}
```

