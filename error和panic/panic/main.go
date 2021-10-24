package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for {
			fmt.Println("hello go1")
		}
	}()

	Go(
		func() {
			time.Sleep(1 * time.Second)
			panic("goroutine2_panic")
		})
	if err := recover(); err != nil {
		fmt.Println("主动 goroutine panic 被捕获")
	}

	time.Sleep(5 * time.Second)
}

func Go(x func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("goroutine2 panic 被捕获")
			}
		}()
		x()
	}()
}
