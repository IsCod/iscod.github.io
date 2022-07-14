# 快速排序

`快速排序`是由英国计算机科学家 `Tony Hoare` 发明的一种排序算法。
`快速排序`是一种分治法策略算法既一种分而治之思想典型应用，本质上看`快速排序`是在`冒泡排序`基础上的递归分治法

## 算法介绍

1, 先从数列中取出一个数作为基准数（pivot,一般取第一个数）

2, 将所有元素与基准数对比，比基准数小的放到基准数前面，比基准数大的放到基准数后面

3, 对基准数两侧重复步骤1，2（递归）直到各区间只有一个数


如下乱序数列：`[10,6,11,21,16]`

```
步骤1：
    选出基准数 10
步骤2：
    比较剩余数列
    6 < 10 放到基准数前 6,10
    11 > 10 放到基准数后 6,10,11
    21 > 10 放到基准数后 6,10,11,21
    16 > 10 放到基准数后 6,10,11,21,16
    基准数(10)左侧数列进行步骤1
       选出基准数(6)
        比较剩余数列（无数值直接返回）
        结果返回：[6]
    基准数(10)右侧数列进行步骤1
        选出基准数11
        比较剩余数列
        21 > 11 放到基准数后 11,21
        16 > 11 放到基准数后 11,21,16
        结果返回：11,21,16
        基准数(11)左侧数列进行步骤1
           无数值返回空
           结果返回[]
        基准数(11)右侧数列进行步骤1
            选出基准数21
            比较剩余数列
            21 > 16 放到基准数后 16,21
            结果返回[16,21]

迭代结束
返回结果:[6,10,11,16,21]
```

## 复杂度

`O(n^2)` 大多数情况下为`o(n logn)`

## 代码实现

Golang: 
```go
package main

import "fmt"

func quickSort(arr []int, pivot int) []int {
	if len(arr) < 1 {
		return []int{pivot}
	}
	var leftArr, rightArr, returnArr []int
	for _, v := range arr {
		if v > pivot {
			leftArr = append(leftArr, v)
		} else {
			rightArr = append(rightArr, v)
		}
	}

	if len(leftArr) > 0 {
		leftArr = quickSort(leftArr[1:], leftArr[0])
	}
	if len(rightArr) > 0 {
		rightArr = quickSort(rightArr[1:], rightArr[0])
	}

	for _, i := range rightArr {
		returnArr = append(returnArr, i)
	}
	returnArr = append(returnArr, pivot)

	for _, i := range leftArr {
		returnArr = append(returnArr, i)
	}
	return returnArr
}

func main() {
	var arr = []int{10, 6, 11, 100, 21, 7, 4, 89, 70, 10}
	sArr := quickSort(arr[1:], arr[0])
	fmt.Printf("%d", sArr)
}
```

输出： 
```[4 6 7 10 10 11 21 70 89 100]```

PHP: 
```php
function quick_sort(int $pivot, array $arr): array
{
    if (count($arr) < 1) {
        return [$pivot];
    }
    $leftArr = $rightArr = [];
    foreach ($arr as $val) {
        if ($val > $pivot) {
            $leftArr[] = $val;
        } else {
            $rightArr[] = $val;
        }
    }
    if (count($leftArr) > 0) {
        $leftArr = quick_sort(array_shift($leftArr), $leftArr);
    }
    if (count($rightArr) > 0) {
        $rightArr = quick_sort(array_shift($rightArr), $rightArr);
    }
    return array_merge($rightArr, [$pivot], $leftArr);
}

$arr = [10, 6, 11, 100, 21, 7, 4, 89, 70, 10];
$arr = quick_sort(array_shift($arr), $arr);
print_r($arr);
```

输出： 
```[4 6 7 10 10 11 21 70 89 100]```