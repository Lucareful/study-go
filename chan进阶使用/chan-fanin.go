package main

import (
	"fmt"
	"reflect"
	"time"
)

// 扇入模式 反射实现
func fanInReflect(channels ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)
		// 构造 SelectCases slice
		var cases []reflect.SelectCase
		for _, c := range channels {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			})
		}

		// 循环，从 cases 中选择一个可用的
		for len(cases) > 0 {
			i, v, ok := reflect.Select(cases)
			if !ok { // chan 关闭
				cases = append(cases[:i], cases[i+1:]...)
				continue
			}
			out <- v.Interface()
		}
	}()

	return out
}

func fanInRec(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		c := make(chan interface{})
		close(c)
		return c
	case 1:
		return channels[0]
	case 2:
		return mergeTow(channels[0], channels[1])
	default:
		m := len(channels) / 2
		return mergeTow(
			fanInRec(channels[:m]...),
			fanInRec(channels[m:]...))
	}
}

// 合并两个 chan
func mergeTow(a, b <-chan interface{}) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		for a != nil || b != nil {
			select {
			case v, ok := <-a:
				if !ok { // a 已关闭，设置为nil
					a = nil
					continue
				}
				c <- v

			case v, ok := <-b:
				if !ok { // b已关闭，设置为nil
					b = nil
					continue
				}
				c <- v
			}
		}
	}()

	return c
}

func sigs(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()

	return c
}

func main() {
	start := time.Now()

	<-fanInReflect(
		sigs(10*time.Second),
		sigs(02*time.Second),
		sigs(03*time.Second),
		sigs(04*time.Second),
		sigs(05*time.Second),
		sigs(01*time.Second),
	)

	fmt.Printf("done after %v", time.Since(start))
}
