package main

import (
	"context"
	"fmt"
	"github.com/bits-and-blooms/bloom/v3"
	"github.com/go-redis/redis/v8"
)

func main() {
	filter := bloom.New(1000000, 4)

	filter.Add([]byte("test1"))
	filter.AddString("test2")

	fmt.Printf("Filter test1 : %v\n", filter.TestString("test1"))
	fmt.Printf("Filter test2 : %v\n", filter.TestString("test2"))

	//存储到redis可以下次直接使用
	client := redis.NewClient(&redis.Options{Addr: ":6379"})

	buf, err := filter.GobEncode()
	if err != nil {
		fmt.Println(err.Error())
	}

	err = client.Set(context.Background(), "fb", buf, 0).Err()
	if err != nil {
		fmt.Println(err)
		return
	}

	// redis中初始化
	buf1, err := client.Get(context.Background(), "fb").Bytes()
	if err != nil {
		fmt.Println(err)
		return
	}

	var f1 bloom.BloomFilter
	err = f1.GobDecode(buf1)
	//err = gob.NewDecoder(bytes.NewBuffer(buf1)).Decode(&f1)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("New Bloom Filter test1 : %v\n", filter.TestString("test1"))
	fmt.Printf("New Bloom Filter test2 : %v\n", filter.TestString("test2"))
	fmt.Printf("New Bloom Filter test3 : %v\n", filter.TestString("test3"))
}
