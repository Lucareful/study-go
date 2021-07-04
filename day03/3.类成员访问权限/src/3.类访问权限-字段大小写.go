package src

import "fmt"

// 类外面绑定方法
// 即使在类的定义前也没关系，go应该不是逐行运行

func (h *Human) Eat() {
	fmt.Println(h.Name + " is eating")
}

type Human struct {
	Name string
	Sex  string
	Age  int
}

type Student struct {
	Hum    Human // 包含 Human类型的变量 是嵌套类
	School string
}

type Teacher struct {
	Human  // 直接声明Human类型，没有定义变量 类继承
	School string
}
