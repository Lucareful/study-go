package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}

	// 单向通道：为了明确语义,一般用于函数参数
	// 单向读通道：
	//var numChanReadOnly <-chan int
	// 单向写通道:
	//var numChanWriteOnly chan<- int

	//numChanReadOnly = make(chan int, 10)
	//numChanWriteOnly = make(chan int, 10)

	// 双向管道可以赋值给单向管道，单向不能转双向
	numChan := make(chan int, 10) // 双向管道

	// 生产者消费者模型
	wg.Add(1)
	go producer(numChan, &wg)
	wg.Add(1)
	go consumer(numChan, &wg)

	wg.Wait()

}

// producer :生产者  ===> 提供一个只写通道

func producer(in chan<- int, wg *sync.WaitGroup) {

	for i := 0; i < 10; i++ {
		in <- i
		//data:= <-out 写通道不允许读操作
		fmt.Println("======> 向管道中写入数据:", i)
	}
	close(in)
	wg.Done()

}

// consumer :消费者  ===> 只提供一个只读通道

func consumer(out <-chan int, wg *sync.WaitGroup) {

	//out <-10 读通道不允许有写入数据
	for v := range out {
		fmt.Println(v, "<======= 从管道中读取数据")
	}
	wg.Done()
}
