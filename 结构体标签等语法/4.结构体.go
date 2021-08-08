package main

import "fmt"

// 结构体声明 type + 别名 + struct
type Student struct {
	name  string
	age   int
	sex   string
	score float64
}

func main() {
	lynn := Student{
		name:  "lynn",
		age:   20,
		sex:   "girl",
		score: 100, // 最后换一个变量后必须加`,`, 或者以`}`结尾
	}

	fmt.Println("lynn", lynn)

	s1 := &lynn
	fmt.Println("s1", s1.name, s1.age, s1.sex, s1.score)
	fmt.Println("s1", (*s1).name, (*s1).age, (*s1).sex, (*s1).score)

	// 如果只对结构体部分变量赋值，name应该指定变量名称
	luenci := Student{
		name: "luenci",
		age:  21,
	}
	fmt.Println("leunci", luenci)

}
