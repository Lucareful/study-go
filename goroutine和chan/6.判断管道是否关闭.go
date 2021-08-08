package main

import "fmt"

func main() {
	numChan := make(chan int, 10)

	go func() {
		for i := 0; i < 10; i++ {
			numChan <- i
		}
		close(numChan)
	}()

	for {
		// ok-idom 模式判断
		if v, ok := <-numChan; ok {
			fmt.Println("读取数据", v)
		} else {
			fmt.Println("管道已经关闭！")
			break
		}
	}

}
