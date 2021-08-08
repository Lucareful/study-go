package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// Config 配置结构体
type Config struct {
	NodeName string
	Addr     string
	Count    int32
}

// loadNewConfig 加载配置
func loadNewConfig() Config {
	return Config{
		NodeName: "上海",
		Addr:     "127.1.0.1",
		Count:    rand.Int31(),
	}
}

func main() {
	var config atomic.Value
	config.Store(loadNewConfig())

	cond := sync.NewCond(&sync.Mutex{})

	// 设置新的config
	go func() {
		for {
			time.Sleep(time.Duration(5+rand.Int63n(5)) * time.Second)
			config.Store(loadNewConfig())
			cond.Broadcast() // 通知等待配置变更
		}
	}()

	go func() {
		for {
			cond.L.Lock()
			cond.Wait()                 // 等待变更信号
			c := config.Load().(Config) // 读取新的配置
			fmt.Printf("New config: %v\n", c)
			cond.L.Unlock()
		}
	}()

	select {}

}
