package main

import "fmt"

func main() {

	numChan := make(chan int)

	go func() {
		for i := 0; i < 50; i++ {
			numChan <- i
			fmt.Println("写入数据<<", i)
		}
		fmt.Println("数据写入完成,关闭管道")
		// 从一个close的管道中读取数据时，会返回零值（不会崩溃）
		close(numChan)
	}()

	// 遍历一个管道时，只会返回一个值
	for v := range numChan {
		fmt.Println("读取数据>>", v)
	}

}
