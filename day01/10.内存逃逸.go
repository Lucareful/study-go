package main

import "fmt"

func main() {
	resPtr := test3()
	fmt.Println("rePtr:", *resPtr)
}

func test3() *string {
	// 没有被返回，没有逃逸
	name := "lynn"
	p0 := &name
	fmt.Println("p0", *p0)

	// 地址返回 内存逃逸
	city := "上海"
	ptr := &city
	fmt.Println("地址为:", ptr)

	return ptr
}
