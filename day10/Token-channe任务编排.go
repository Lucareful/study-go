package main

import (
	"fmt"
	"time"
)

type Token struct{}

func newWork(id int, ch chan Token, nextChan chan Token) {
	for {
		token := <-ch
		fmt.Println(id + 1)
		time.Sleep(time.Second)
		nextChan <- token
	}
}

func main() {
	chs := []chan Token{make(chan Token), make(chan Token), make(chan Token), make(chan Token)}

	// 创建四个works
	for i := 0; i < 4; i++ {
		go newWork(i, chs[i], chs[(i+1)%4])
	}

	// 初始化token
	chs[0] <- struct{}{}

	select {}
}
