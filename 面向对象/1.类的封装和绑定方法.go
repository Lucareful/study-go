package main

import "fmt"

// go语言没有class关键字生成类
// 使用struct声明类

type Person struct {
	// 成员属性
	name   string
	age    int
	gender string
}

// 类外边绑定方法
// 类的方法，可以使用自己的成员
// 使用指针可以修改类的成员变量等

func (p *Person) Eat() {
	fmt.Println("使用 *Person 指针 修改前")
	fmt.Println(p.name + " is eating")
	p.name = "luenci"
	fmt.Println("使用 *Person 指针 修改后")
	fmt.Println(p.name + " is eating")
}

func (p Person) Eat2() {
	fmt.Println("使用 Person 不是指针 修改前")
	fmt.Println(p.name + " is eating")
	p.name = "luenci"
	fmt.Println("使用 Person 不是指针 修改后")
	fmt.Println(p.name + " is eating")
}

func main() {
	lynn := Person{
		name:   "lynn",
		age:    20,
		gender: "girl",
	}

	lynn.Eat()
	lynn.Eat2()

}
