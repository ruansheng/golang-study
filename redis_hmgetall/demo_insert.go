package main

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
)

func test(key string, i int, port int) {
	fmt.Println("doing insert key:", key)
	//设置redis服务器地址
	client := redis.NewClient(&redis.Options{
		Addr:     "47.94.226.123:" + strconv.Itoa(port),
		Password: "asdfesdgrrdfgedfedsd",
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
	ports := []int{6380, 6381, 6382, 6383, 6384, 6385, 6386, 6387, 6388, 6389}
	for i := 0; i < 10; i++ {
		test("users"+strconv.Itoa(i), i, ports[i])
	}
}
