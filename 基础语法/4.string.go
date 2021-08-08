package main

import (
	"fmt"
)

func main() {
	name := "Lynn"

	// 换行，原生字符串输出时候，使用反引号 ``
	usage := `./a.out <option>
			-h	help
			-a	xxx`

	fmt.Println("name:", name)
	fmt.Println("usage:", usage)

	// 长度 自由函数 len()
	fmt.Println("name len", len(name))

	// 基本循环
	for i := 0; i < len(name); i++ {
		fmt.Printf("name[%d] %c \n", i, name[i])
	}

	// 拼接
	i, j := "hello", "world"
	fmt.Println("i+j", i+j)

	// 使用const 修饰为常量不能修改
	const ip = "127.0.0.1"
	fmt.Println("const 常量:", ip)
}
