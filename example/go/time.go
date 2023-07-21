package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	time1 := time.Now()

	fmt.Printf("%s \n", time1.Format(time.RFC3339Nano))

	time2, _ := time.Parse(time.RFC3339Nano, time1.Format(time.RFC3339Nano))
	time2 = time2.Add(time.Second)

	diff := time1.Sub(time2).Milliseconds()
	fmt.Println(diff)
	fmt.Println(time1.Before(time2))
	fmt.Println(time1.After(time2))
	fmt.Println(time1.Equal(time2))

	switch time1.Weekday() {
	case time.Sunday, time.Saturday:
		fmt.Println("sunday")
	default:
		fmt.Println("week day")
	}

	rand1 := rand.New(rand.NewSource(10))
	rand2 := rand.New(rand.NewSource(10))
	fmt.Println(rand1.Intn(100))
	fmt.Println(rand2.Intn(100))
}
