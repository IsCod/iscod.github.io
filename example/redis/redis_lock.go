package main

import "C"
import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"math/rand"
	"runtime"
	"time"
)

func getRdb() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: ":6379",
	})

	return rdb
}

/**
* @param lockName 锁名称
* @param lockTime 锁的时间
* @param outTime 获取锁的超时时间
* @return string 锁的标示
 */
func getLockWithTimeout(rdb *redis.Client, lockName string, lockTime time.Duration, outTime time.Duration) string {

	lockKey := "DLock:" + lockName

	timer := time.NewTimer(outTime)
	identifier := fmt.Sprintf("%d%d", time.Now().Unix(), rand.Intn(1000))
	ttl := fmt.Sprintf("%d", int(lockTime.Seconds()))
	script := `
local key = KEYS[1]
local required = KEYS[2]
local ttl = KEYS[3]

local result = redis.call('SETNX', key, required)

if result == 1 then
    --设置成功，则设置过期时间
   redis.call('EXPIRE', key, ttl)
else
    local value = redis.call('get', key)
    if value == result then
        --如果跟之前的锁一样，则重新设置时间
        result = 1
       
    end
end
--成功则返回1
return result
		`
	for {
		select {
		case <-timer.C:
			identifier = ""
			return identifier
		default:
			if r, err := rdb.Eval(context.Background(), script, []string{lockKey, identifier, ttl}).Int(); err == nil && r == 1 {
				return identifier
			} else {
				//fmt.Println(err, r)
			}
		}
	}
}

func releaseLock(rdb *redis.Client, lockName, identifier string) bool {
	script := `
--当锁匹配的钥匙相同时才可以删除锁
local key = KEYS[1]
local required = KEYS[2]
local value = redis.call('GET', key)
if value == required then
    redis.call('DEL', key);
    return 1;
end
return 1;
	`
	lockKey := "DLock:" + lockName
	if r, err := rdb.Eval(context.Background(), script, []string{lockKey, identifier}).Int(); err == nil && r == 1 {
		return true
	} else {
		if err != nil {
			fmt.Printf("del lock error : %d,%s\n", r, err)
		}

	}
	return false
}

func main() {

	var max = 10

	rdb := getRdb()

	runtime.GOMAXPROCS(runtime.NumCPU())
	rdb.Set(context.Background(), "sellNum", 0, 0)
	for i := 0; i < 5000; i++ {
		userName := 1000 + i
		go func() {
			identifier := getLockWithTimeout(rdb, "Huawei Mate 10", 1*time.Second, time.Second*20)

			if identifier != "" {
				if n, err := rdb.Get(context.Background(), "sellNum").Int(); n < max && err == nil {
					rdb.Incr(context.Background(), "sellNum")
					fmt.Printf("正在为用户：%d 处理订单 购买第 %d 台 剩余 %d 台\n", userName, n, max)
				} else {
					//fmt.Printf("%s %d %d\n", err, n, max)
					//fmt.Printf("用户：%d 无法购买，已售罄！\n", userName)
				}
				//fmt.Printf("identifier: %s", identifier)
				releaseLock(rdb, "Huawei Mate 10", identifier)
			} else {
				//fmt.Printf("===")
			}
		}()
	}
	go func() {
		for {
			time.Sleep(1 * time.Second)
			fmt.Printf("go runtime: %d\n", runtime.NumGoroutine())
		}
	}()
	time.Sleep(20 * time.Second)

	sellnum := rdb.Get(context.Background(), "sellNum").String()
	fmt.Printf("库存 %d 台， 共卖出 %s 台\n", max, sellnum)
	time.Sleep(1 * time.Second)
}
