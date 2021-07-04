package main

import "fmt"

func main() {
	// 变量定义: var
	// 常量定义: const

	// 1.先定义变量，var 变量名 类型 ；再赋值 变量=XX
	var name string
	name = "Luenci"

	var age int
	age = 22

	fmt.Println("name:", name)

	// 2.定义时直接赋值
	var sex = "man"

	fmt.Printf("sex is %s \n", sex)

	fmt.Printf("name is %s age is %d \n", name, age)

	// 3.定义时直接赋值，使用时自动推导类型（常用）
	hobby := "篮球"

	fmt.Println("my hobby is ", hobby)

	// 灰色部分表示形参
	test(10, "ll")

	// 4. 平行赋值
	// 同时赋值给两个变量
	i, j := 10, 22
	fmt.Println(i)
	fmt.Println(j)
	i, j = j, i
	fmt.Println(i)
	fmt.Println(j)
}

func test(a int, b string) {
	fmt.Println(a)
	fmt.Println(b)
}
