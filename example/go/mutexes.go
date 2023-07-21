// 互斥锁
// https://gobyexample.com/mutexes
package main

import (
	"fmt"
	"sync"
)

type Container struct {
	sync.Mutex
	counters map[string]int
}

func (c *Container) incr(k string, incr int) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	c.counters[k] += incr
}

func main() {
	c := Container{counters: map[string]int{"a": 0, "b": 0}}

	doIncr := func(name string, n int, wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 0; i < n; i++ {
			c.incr(name, 1)
		}
	}

	wg := sync.WaitGroup{}
	wg.Add(4)
	go doIncr("a", 1000, &wg)
	go doIncr("b", 1000, &wg)
	go doIncr("b", 1000, &wg)
	go doIncr("b", 1000, &wg)
	wg.Wait()
	fmt.Println(c.counters)

}
