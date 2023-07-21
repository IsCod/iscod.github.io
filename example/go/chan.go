package main

import (
	"fmt"
	"time"
)

//题目：
//A协程只输出0；
//B协程只输出奇数；
//C协程只输出偶数；
//三个协程完成协程间通信：
//例：PrintNumber(5)
//输出：0 1 0 2 0 3 0 4 0 5

func A(num int) {
	var ch, ch1 = make(chan int, 0), make(chan int, 0)
	var out = make(chan int, 0)
	go B(ch, out)
	go C(ch1, out)

	for i := 1; i <= num; i++ {
		fmt.Println(0)
		if i%2 == 0 {
			ch <- i
		} else {
			ch1 <- i
		}
		<-out
	}
}

func B(in chan int, out chan int) {
	for {
		select {
		case i := <-in:
			fmt.Println(i)
			out <- 1
		}
	}
}

func C(in chan int, out chan int) {
	for {
		select {
		case i := <-in:
			fmt.Println(i)
			out <- 1
		}
	}
}

func x(in chan int) {
	for {
		select {
		case t, ok := <-in:
			time.Sleep(time.Second * 1)
			fmt.Printf("%v, %v \n", t, ok)
		}
	}
}
func main() {
	var ch = make(chan int, 5)
	go x(ch)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()
	time.Sleep(time.Second * 10)
}
