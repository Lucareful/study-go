package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once

	f1 := func() {
		fmt.Println("f1 exce！")
	}
	once.Do(f1)

	f2 := func() {
		fmt.Println("f2 exce")
	}
	once.Do(f2)
}
