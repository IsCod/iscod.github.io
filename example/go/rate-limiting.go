// 速率限制
// https://gobyexample.com/rate-limiting
package main

import (
	"fmt"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var burstyLimiter = make(chan int, 3)

	go func() {
		for i := 0; i < 3; i++ { //设置初始三个缓存
			burstyLimiter <- i
		}
		tick := time.NewTicker(time.Millisecond * 1000)
		for {
			select {
			case <-tick.C:
				burstyLimiter <- 1
			}
		}
	}()

	burstyRequests := make(chan int, 5)

	for i := 0; i < 5; i++ { //五个请求到达
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter //通过burstyLimiter控制处理速度
		fmt.Printf("req id %d %s\n", req, time.Now())
	}
}
