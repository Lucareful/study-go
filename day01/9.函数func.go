package main

import "fmt"

func test1(a int, b int, c string) (int, string, bool) {

	return a + b, c, false
}

func test2(a, b int, c string) (res int, str string, bl bool) {

	// 直接使用返回值变量名参与运算
	res = a + b
	str = c
	bl = false

	// 当返回值有名称时候，可以直接return
	return
}

func main() {
	a, b, c := test1(1, 2, "luenci")
	fmt.Println("test1函数返回值为", a, b, c)

	res, str, bl := test2(1, 2, "luenci")
	fmt.Println("test2函数返回值为", res, str, bl)
}
