// 有状态的goroutine
// https://gobyexample.com/stateful-goroutines
package main

type ops struct {
	key   string
	value string
	resp  chan int
}

func main() {
	reads := make(chan ops, 0)
	for i := 0; i < 100; i++ {
		go func() {
			reads <- ops{key: ""}
		}()
	}
}
