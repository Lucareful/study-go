package main

import (
	"fmt"
	"sync"
)

func main() {
	/* 当程序中有多个channel协同工作，ch1，ch2，某一时刻，ch1或者ch2触发了，程序要做响应处理
	1.使用 select 来监听多个管道，当管道被触发时（写入数据，读取数据，关闭管道）
	2.select 语法于 switch case 很像，但是所有的分支条件必须是通道I/O
	*/
	var ch1, ch2 chan int
	ch1 = make(chan int)
	ch2 = make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(3)

	// 启动一个go程，负责监听两个channel
	go func() {
		fmt.Println("开始监听....")
		for {
			select {
			case data1 := <-ch1:
				fmt.Println("从ch1中读取数据: ", data1)
			case data2 := <-ch2:
				fmt.Println("从ch2中读取数据: ", data2)
			}
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			ch2 <- i
		}
		wg.Done()
	}()

	wg.Wait()
}
