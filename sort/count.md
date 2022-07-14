# 计数和桶排序

#### 计数排序

计数排序的核心是将输入的数据转化为键存储在新开辟的数组空间中。作为一种先线性时间复杂度的排序，计数排序要求输入的数据必须是有确定范围的整数。

计数排序在输入数据的偏值较大时内存占用比较浪费，在偏离值比较小时能实现快速排序

#### 桶排序


## 算法介绍

## 代码实现

Golang: 
```go
package main

import "fmt"

func countSort(arr []int, max int) []int {
	var carr = make([]int, max+1)

	for _, v := range arr {
		carr[v]++
	}

	var rArr = []int{}
	for k, v := range carr {
		for i := 0; i < v; i++ {
			rArr = append(rArr, k)
		}
	}

	return rArr
}

func main() {
	var arr = []int{10, 6, 11, 100, 21, 7, 4, 89, 70, 10}
	sArr := countSort(arr, 100)
	fmt.Printf("%d", sArr)
}
```