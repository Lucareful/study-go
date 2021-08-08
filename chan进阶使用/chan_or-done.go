package main

import (
	"fmt"
	"reflect"
	"time"
)

// or or-done 模式 递归实现
func or(channels ...<-chan interface{}) <-chan interface{} {
	// 特殊情况，只有零个,1个或2个 chan
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	case 2:
		select {
		case <-channels[0]:
		case <-channels[1]:
		}
	}

	orDone := make(chan interface{})
	go func() {
		defer close(orDone)
		// fmt.Println("执行")
		if len(channels) > 2 {
			m := len(channels) / 2
			select {
			case <-or(channels[:m]...):
			case <-or(channels[m:]...):
			}
		}
	}()
	return orDone
}

func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()

	return c
}

// orSelect 反射⽅式 实现
func orSelect(channels ...<-chan interface{}) <-chan interface{} {
	// 特殊情况，只有零个,1个或2个 chan
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	orDone := make(chan interface{})
	go func() {
		defer close(orDone)
		// 利用反射构建SelectCase
		var cases []reflect.SelectCase
		for _, c := range channels {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			})
		}
		// 随机选取一个可用 case
		reflect.Select(cases)
	}()

	return orDone
}

func main() {
	start := time.Now()

	<-orSelect(
		sig(10*time.Second),
		sig(20*time.Second),
		sig(30*time.Second),
		sig(40*time.Second),
		sig(50*time.Second),
		sig(01*time.Second),
	)

	fmt.Printf("done after %v", time.Since(start))

}
