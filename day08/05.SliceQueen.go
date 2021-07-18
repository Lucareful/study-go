package main

import (
	"fmt"
	"sync"
)

// SliceQueue 线程安全队列
type SliceQueue struct {
	data []interface{}
	mu   sync.Mutex
}

// NewSliceQueue 初始化一个长度为 n 的队列
func NewSliceQueue(n int) (q *SliceQueue) {
	return &SliceQueue{data: make([]interface{}, 0, n)}
}

// Enqueue 入队
func (q *SliceQueue) Enqueue(v interface{}) {
	q.mu.Lock()
	q.data = append(q.data, v)
	q.mu.Unlock()
}

// Dequeue 出队并返回元素
func (q *SliceQueue) Dequeue() interface{} {
	q.mu.Lock()
	if len(q.data) == 0 {
		q.mu.Unlock()
		return nil
	}
	v := q.data[0]
	q.data = q.data[1:]
	q.mu.Unlock()

	return v
}

func main() {
	sliceQueens := NewSliceQueue(8)
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		// fmt.Println("i:",i)
		// time.Sleep(time.Second)
		go func() {
			defer wg.Done()
			// fmt.Println(i)
			sliceQueens.Enqueue(i)
		}()
	}

	wg.Wait()
	fmt.Println(sliceQueens.data)
}
