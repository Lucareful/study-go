package main

import "fmt"

// iota 是常量组计数器
// iota从0开始，每换行递增1,从第一次出现iota开始计算
// 常量组如果不赋值，默认上一行表达式相同
// iota是以行为单位递增

// 每个常量组的iota是独立的，都是从零开始递增

const (
	MONDAY = iota // 0
	TUESDAY
	WEDNESDAY
	THURSDAY
	FRIDAY
	SATURDAY
	M, N = iota, iota // const 属于预编译，所以不需要 := 自动推导

)

// go语言中没有枚举类型，但是可以使用 const + iota（常量累加器）来进行模拟
func main() {
	// 变量组统一命名变量
	var (
		number int
		name   string
		flag   bool
	)
	fmt.Println(number, name, flag)

	fmt.Println(MONDAY)
	fmt.Println(TUESDAY)
	fmt.Println(WEDNESDAY)
	fmt.Println(THURSDAY)
	fmt.Println(FRIDAY)
	fmt.Println(SATURDAY)
	fmt.Println(M, N)

}
