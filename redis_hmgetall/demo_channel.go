package main

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
)

const CHAN_SIZE = 1000

var ch chan map[string]string

var data map[string]string

func test(key string) {
	fmt.Println("doing search key:", key)
	//设置redis服务器地址
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// 从redis获取hash数据
	list, err := client.HGetAll(key).Result()
	if err != nil {
		ch <- list
		return
	}

	// 消息传递
	ch <- list
	return
}

func main() {
	ch = make(chan map[string]string, CHAN_SIZE)
	data = make(map[string]string)

	for i := 1; i <= CHAN_SIZE; i++ {
		go test("users" + strconv.Itoa(i))
	}

	for j := 1; j <= CHAN_SIZE; j++ {
		tmps := <-ch
		for index := range tmps {
			data[index] = tmps[index]
		}
	}
	fmt.Println(data)
	fmt.Println("data len:", len(data))
}
