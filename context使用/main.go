package main

import (
	"context"
	"fmt"
	"time"
)

type Tracker struct {
	ch   chan string
	stop chan struct{}
}

func (t *Tracker) Event(ctx context.Context, data string) error {
	select {
	case t.ch <- data:
		fmt.Println("writer data...")
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (t *Tracker) Run() {
	// chan 在写数据的时候，读阻塞
	fmt.Println("开始执行读")
	for data := range t.ch {
		fmt.Println("read data...")
		fmt.Println(data)
	}
	fmt.Println("读完了")
	// 控制 t.ch 读取完毕
	t.stop <- struct{}{}
}

func (t *Tracker) Shutdown(ctx context.Context) {
	close(t.ch)
	fmt.Println("通道关闭了")

	select {
	case <-t.stop:
	case <-ctx.Done():
	}
}

func NewTracker() *Tracker {
	return &Tracker{
		ch: make(chan string, 10),
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	t := NewTracker()
	go t.Run()
	_ = t.Event(ctx, "test1")
	_ = t.Event(ctx, "test2")
	_ = t.Event(ctx, "test3")
	defer cancel()
	t.Shutdown(ctx)
}
