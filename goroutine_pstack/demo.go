package main

import (
	"fmt"
	"runtime"
	"time"
)

func sleep(s string) {
	for {
		fmt.Println(s)
		time.Sleep(time.Second)
	}
}

func main() {
	fmt.Println(runtime.GOMAXPROCS(0))
	//runtime.GOMAXPROCS(1)
	fmt.Println(runtime.GOMAXPROCS(0))
	go sleep("Hello")
	go sleep("World")
	time.Sleep(1000 * time.Second)
	//for {
	//}
}
