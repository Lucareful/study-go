package main

// sub 是文件夹名，也是package名
import (
	_ "luenci/github.com/day02/5.init函数/sub" // 只会调用sub中的init函数
)

// init 函数没有参数和返回值,使用如下
// 同一个包中包含多个init函数时候,调用顺序是不确定的（同一个package下的多个文件都可以有init）
// init 函数是不允许调用(显示调用)的
// 如果只想调用一个package中的 init 函数,只需在导包前加上 _
// https://juejin.cn/post/6844903550095458312
func main() {

	//res := sub.Sub(10, 5)
	//fmt.Println("sub res", res)
}
