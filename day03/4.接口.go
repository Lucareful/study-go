package main

import (
	"fmt"
)

// go语言中,interface代表接口
// interface 不仅仅是用来处理多态，它可以接收任意的数据类型

func main() {
	// func Println(a ...interface{}) (n int, err error) { ... }

	// 定义三个接口类型
	var i, j, k interface{}

	i = []string{"lynn", "luenci"}
	j = 98
	k = false

	// 快速判断接口的类型
	//jType := reflect.TypeOf(j) 反射
	jType, ok := j.(int)
	fmt.Println("ok", ok)
	if !ok {
		fmt.Println("j不是int类型")
	}

	fmt.Println("jType", jType)
	fmt.Println(i, j, k)

	// 最常用的场景：把interface当成一个函数的参数，（类似于上面的 Println函数） ，使用switch来判断用户输入的不同类型
	// 根据不同类型，做相关的逻辑处理

	// 创建一个具有三个接口类型的切片
	array := make([]interface{}, 3)
	array[0] = 3
	array[1] = "luenci"
	array[2] = true

	for _, value := range array {
		// 获取接口中真正的数据类型
		switch v := value.(type) {
		case int:
			fmt.Println("当前数据类型为int，内容为:", v)
		case string:
			fmt.Println("当前数据类型为string，内容为:", v)
		case bool:
			fmt.Println("当前数据类型为bool，内容为:", v)
		default:
			fmt.Println("不是合理的数据类型")
		}
	}

}
