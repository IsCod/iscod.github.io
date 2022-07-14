# 插入排序

插入排序是一种最简单直观的排序算法，它的原理是每次把一个新数插入到已经排好序的数列里面形成新的排好序的数列。
玩过扑克牌应该很容易理解这种插入算法模式。

## 算法介绍

1, 将待排序序列第一个元素看做一个有序序列，把第二个元素到最后一个元素当成未排序的数列

2, 从头到尾依次扫描未排序序列，将扫描到的每个元素插入有序数列的适当位置

如下乱序数列：`[10,6,11,21,16]`

```
步骤1：
   拿出第一个元素组成有序数列：[10] 未排序数列： [6,11,21,16]
步骤2：
   扫描未排序数列: [6,11,21,16]
   扫描第一个元素 6 插入到排好序的数列: [10]
   与排好序的数列[10]进行比较
   6 < 10 插入到10之前
   结束: [6,10]
   扫描第二个元素 11 插入到排好序的数列: [6,10]
   与排好序的数列[6,10]进行比较
   11 > 10 插入到10之后
   结束: [6,10,11]
   扫描第三个元素 21 插入到排好序的数列: [6,10,11]
   与排好序的数列[6,10,11]进行比较
   21 > 11 插入到11之后
   结束: [6,10,11,21]
   扫描第四个元素 16 插入到排好序的数列: [6,10,11,21]
   与排好序的数列[6,10,11,21]进行比较
   21 > 16 插入到21之前
   11 < 16 插入到11后
   结束: [6,10,11,16,21]
结果：[6,10,11,16,21]
```

## 复杂度

`O(n^2)`

## 代码实现

Golang: 
```go
package main

import "fmt"

func insertSort(arr []int) []int {
	for k, _ := range arr {
		if k == 0 {
			continue
		}
		for i, v := range arr[0:k] {
			if v > arr[k] {
				arr[i], arr[k] = arr[k], arr[i]
			}
		}
	}
	return arr
}

func main() {
	var arr = []int{10, 6, 11, 100, 21, 7, 4, 89, 70, 10}
	sArr := insertSort(arr)
	fmt.Printf("%d", sArr)
}
```

输出： 
```[4 6 7 10 10 11 21 70 89 100]```

PHP: 
```php
function insert_sort(array $arr): array
{
    foreach ($arr as $key => $value) {
        if ($key == 0) {
            continue;
        }

        for ($i = 0; $i < $key; $i++) {
            if ($arr[$i] > $value) {
                $kValue = $arr[$key];
                $arr[$key] = $arr[$i];
                $arr[$i] = $kValue;
            }
        }
    }
    return $arr;
}

$arr = insert_sort([10, 6, 11, 100, 21, 7, 4, 89, 70, 10]);
print_r($arr);
```

输出： 
```[4 6 7 10 10 11 21 70 89 100]```

* 参考
    * [数据结构和算法](https://www.bookstack.cn/read/JS-Sorting-Algorithm/3.insertionSort.md)
    * [经典排序算法](https://www.bookstack.cn/read/hunterhug-goa.c/algorithm-sort-insert_sort.md)
