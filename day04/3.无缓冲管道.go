package main

import (
	"fmt"
	"time"
)

/*
当通道在多个协程之间传输的是指向数据的指针是，且读写操作是由不同的协程操作，则需要提供额外的同步动作。
*/
func main() {
	// 当涉及到多go程时，c语言使用互斥量，上锁来保持资源同步，避免资源竞争问题
	// go语言更好的解决方案是管道、通道
	// 使用通道不需要手动进行加锁
	//sync.RWMutex{}

	// 创建管道 关键字 chan
	numChan := make(chan int) // 装数字的管道，无缓冲通道，未声明空间
	//numChan := make(chan int, 10) // 有缓冲通道

	// 创建两个go程，父写，子读
	// 发现子go程没有发生资源抢夺

	// 子go程1
	go func() {
		for i := 0; i < 25; i++ {
			// 只能 <- 数据流向
			data := <-numChan
			fmt.Println("子go程1 读取data", data)
		}
	}()
	// 子go程2
	go func() {
		for i := 0; i < 25; i++ {
			data := <-numChan
			fmt.Println("子go程2 读取data", data)
		}
	}()

	// 父go程
	for i := 0; i < 50; i++ {
		// 向管道中写入数据
		numChan <- i
		fmt.Println("====> 主go程，写入数据", i)
	}

	time.Sleep(5 * time.Second)

}
