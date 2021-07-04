package main

import "fmt"

type Human struct {
	name string
	sex  string
	age  int
}

type Student struct {
	hum    Human // 包含 Human类型的变量 是嵌套类
	school string
}

type Teacher struct {
	Human  // 直接声明Human类型，没有定义变量 类继承
	school string
}

// 类外面绑定方法

func (h *Human) Eat() {
	fmt.Println(h.name + " is eating")
}

func main() {
	st1 := Student{
		hum: Human{
			name: "lynn",
			sex:  "girl",
			age:  20,
		},
		school: "一中",
	}
	fmt.Println("st1", st1)
	fmt.Println("st1 name", st1.hum.name)

	t1 := Teacher{}
	t1.school = "一中"
	t1.name = "lynn"
	t1.sex = "girl"
	t1.age = 20

	fmt.Println("t1", t1)
	fmt.Println("t1 name", t1.name)
	// 继承的时候虽然我们没有声明变量名称，但是默认自动会给类型创建一个同名字段
	// 这是为了能在子类中操作父类，因为：子类父类可能出现同名字段
	fmt.Println("t1 age", t1.Human.age)
}
