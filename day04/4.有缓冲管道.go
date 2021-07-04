package main

import "fmt"

/* 当管道的读写次数不一致的时候
1.如果阻塞在主go程，程序会崩溃
2.如果阻塞在子go程，会发生内存泄露
*/
func main() {
	// 当缓冲写满的时候，写阻塞，被读取后，在恢复写入
	// 当缓冲区读取完毕，读阻塞，开始写入
	// 如果没有使用make分配空间，那么管道默认nil的，读取，写入都会阻塞
	numChan := make(chan int, 10) // 有缓冲通道

	// 子go程1
	go func() {
		for i := 0; i < 25; i++ {
			// 只能 <- 数据流向
			data := <-numChan
			fmt.Println("子go程1 读取data", data)
		}
	}()

	// 父go程
	for i := 0; i < 50; i++ {
		// 向管道中写入数据
		numChan <- i
		fmt.Println("====> 主go程，写入数据", i)
	}

	var names chan string
	// 因为names是nil的，写操作会一直阻塞在这里
	// 并发生 deadlock 的 error
	names <- "luenci"
	fmt.Println("names", <-names)
}
