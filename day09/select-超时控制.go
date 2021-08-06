package main

import (
	"fmt"
	"time"
)

// server 服务处理逻辑
func server(ch chan bool) {
	time.Sleep(time.Second * 6)
	ch <- true
}

func main() {
	ch := make(chan bool, 1)
	go server(ch)
	select {
	case ret := <-ch:
		fmt.Println(ret)
	case <-time.After(time.Second * 5):
		fmt.Println("Timeout")
	}
}
