package main

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
)

func test(key string, i int) {
	fmt.Println("doing insert key:", key)
	//设置redis服务器地址
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// 从redis获取hash数据
	var index int
	for j := 1; j <= 10; j++ {
		index = i*100 + j
		client.HSet(key, strconv.Itoa(index), index)
	}

}

func main() {
	for i := 1; i <= 1000; i++ {
		test("users"+strconv.Itoa(i), i)
	}
}
