package main

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/go-redis/redis"
)

const CHAN_SIZE = 100

var wg sync.WaitGroup
var l *sync.Mutex

var data map[string]string

func test(key string) {
	defer wg.Done()
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
		return
	}

	l.Lock()
	for index := range list {
		data[index] = list[index]
		//fmt.Println(index, list[index])
	}
	l.Unlock()
}

func main() {
	l = new(sync.Mutex)
	data = make(map[string]string)

	for i := 1; i <= CHAN_SIZE; i++ {
		wg.Add(1)
		go test("users" + strconv.Itoa(i))
	}

	wg.Wait()
	fmt.Println(data)
	fmt.Println("data len:", len(data))
}
