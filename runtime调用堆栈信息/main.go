package main

import (
	"fmt"
	"runtime"
)

func main() {
	Foo()
}
func Foo() {
	fmt.Printf("我是 %s, 谁在调用我?\n", printMyName())
	Bar()
}
func Bar() {
	fmt.Printf("我是 %s, 谁又在调用我?\n", printMyName())
}
func printMyName() string {
	pc, file, line, _ := runtime.Caller(2)
	fmt.Println("file", file)
	fmt.Println("line", line)
	return runtime.FuncForPC(pc).Name()
}
