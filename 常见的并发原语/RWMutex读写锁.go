package main

import (
	"fmt"
	"sync"
	"time"
)

// Counter 一个线程安全的计数器
type Counter struct {
	mu    sync.RWMutex
	count uint64
}

// Inc 使用写锁保护
func (c *Counter) Inc() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

// Count 使用读锁保护
func (c *Counter) Count() uint64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.count
}

func main() {
	var counter Counter
	for i := 0; i < 10; i++ {
		go func() {
			count := counter.Count() // 计数器操作
			fmt.Println(count)
			time.Sleep(time.Millisecond)
		}()
	}

	for {
		// 单个 writer
		counter.Inc()
		time.Sleep(time.Second)
	}

}
