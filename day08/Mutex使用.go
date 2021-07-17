package day08

import (
	"fmt"
	"sync"
)

func main() {
	var count = 0
	// 互斥锁保护计数器
	var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			// 对变量 count 进行加法操作
			// count++ 不是一个原子操作，它至少包含几个步骤，
			// 比如读取变量 count 的当前值，
			// 对这个值加 1，把结果再保存到 count 中。
			// 因为不是原子操作，就可能有并发的问题。
			for i := 0; i < 10000; i++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}
	// 等待 10 个 goroutine 完成
	wg.Wait()
	fmt.Println("结果为", count)
	main2()

}
// Counter 计数器
type Counter struct {
	id    int
	name  string
	mu    sync.Mutex
	count uint
}

// Inc 写入
func (c *Counter) Inc() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}
// Count 统计
func (c *Counter) Count() uint {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main2() {
	var counter Counter

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				counter.Inc() // 受到 mutex 保护的方法
			}
		}()
	}
	wg.Wait()
	fmt.Println("结果为", counter.count)
}
