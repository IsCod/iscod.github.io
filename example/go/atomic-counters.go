// 原子计数器
// https://gobyexample.com/atomic-counters
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var ops int32
	wg := sync.WaitGroup{}
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&ops, 1)
		}()
	}

	wg.Wait()
	fmt.Println(ops)
}
