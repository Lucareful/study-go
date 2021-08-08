package main

import (
	"fmt"
	"runtime"
	"time"
)

// GOEXIT ===> 提前退出go程
// return ===> 返回当前函数
// exit ===> 退出当前进程

func main() {
	go func() {
		func() {
			fmt.Println("子go程内部的函数!")
			//return // 退出当前函数
			//os.Exit(-1) // 退出进程
			runtime.Goexit() // 退出当前go程
		}()

		fmt.Println("子go程结束！")
	}()

	// 主go程需要等待子go程退出
	fmt.Println("主go程~")
	time.Sleep(5 * time.Second)
	fmt.Println("OVER!")
}
