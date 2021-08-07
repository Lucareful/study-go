package main

import (
	"fmt"
	"reflect"
)

// createCase 收集 select 中的 chan case
func createCase(chs ...chan int) []reflect.SelectCase {
	var cases []reflect.SelectCase

	// 创建 recv case
	for _, ch := range chs {
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ch),
		})
	}

	// 创建 send case
	for i, ch := range chs {
		v := reflect.ValueOf(i)
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectSend,
			Chan: reflect.ValueOf(ch),
			Send: v,
		})
	}

	return cases
}

func main() {
	var ch1 = make(chan int, 10)
	var ch2 = make(chan int, 10)

	// 创建 SelectCae
	var cases = createCase(ch1, ch2)

	// 执行 10 次 select

	for i := 0; i < 10; i++ {
		chosen, recv, ok := reflect.Select(cases)
		if recv.IsValid() {
			fmt.Println("recv:", cases[chosen].Dir, recv, ok)
		} else {
			fmt.Println("send:", cases[chosen].Dir, ok)
		}
	}

}
