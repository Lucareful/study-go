package main

import (
	"fmt"
	"time"
)

func display(num int) {
	count := 1
	for {
		fmt.Println("============> 这是子go程：", num, "当前count值", count)
		count++
	}

}

func main() {
	// 启动子go程
	for i := 0; i < 3; i++ {
		go display(i)
	}

	// 主go程
	count := 1
	for {
		fmt.Println("============> 这是主go程：", count)
		count++
		time.Sleep(1 * time.Second)
	}

}
