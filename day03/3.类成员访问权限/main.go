package main

import (
	"fmt"
	"luenci/github.com/day03/3.类成员访问权限/src"
)

// go语言中，权限都是通过首字母大小写来控制的
// import XX 如果package名称不同，那么只有大写字母开头的才是 Public
// 对于类里面的成员、函数等，只有大写字母开头的才能在其他的包中使用
// 如果在用一个package（文件下）下则无以上限制

func main() {
	st1 := src.Student{
		Hum: src.Human{
			Name: "lynn",
			Sex:  "girl",
			Age:  20,
		},
		School: "一中",
	}

	fmt.Println("st1", st1)
	fmt.Println("st1 name", st1.Hum.Name)

	t1 := src.Teacher{}
	t1.School = "一中"
	t1.Name = "lynn"
	t1.Sex = "girl"
	t1.Age = 20

	fmt.Println("t1", t1)
	fmt.Println("t1 name", t1.Name)
	fmt.Println("t1 age", t1.Human.Age)
}
