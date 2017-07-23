package main

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
)

const CHAN_SIZE = 10

var ch chan map[string]string

var data map[string]string

func test(key string, port int) {
	fmt.Println("doing search key:", key)
	//设置redis服务器地址
	client := redis.NewClient(&redis.Options{
		Addr:     "47.94.226.123:" + strconv.Itoa(port),
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
	ports := []int{6380, 6381, 6382, 6383, 6384, 6385, 6386, 6387, 6388, 6389}

	ch = make(chan map[string]string, CHAN_SIZE)
	data = make(map[string]string)

	for i := 0; i < CHAN_SIZE; i++ {
		go test("users"+strconv.Itoa(i), ports[i])
	}

	for j := 0; j < CHAN_SIZE; j++ {
		tmps := <-ch
		for index := range tmps {
			data[index] = tmps[index]
		}
	}
	fmt.Println(data)
	fmt.Println("data len:", len(data))
}
