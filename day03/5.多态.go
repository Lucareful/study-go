package main

import "fmt"

// go语言多态，需要实现定义接口

// 定义一个接口，类型是interface
type IAttack interface {
	// 如果定义了多个接口函数，那么实现的类必须全部实现这些接口，才可以赋值
	Attack() // 接口函数可以有多个，但是只能有函数原型，不可以有实现
}

// 低等级玩家
type HumanLowLevel struct {
	name  string
	level int
}

func (a *HumanLowLevel) Attack() {
	fmt.Println("我是", a.name, "等级为", a.level, "造成伤害1000")
}

// 高等级玩家
type HumanHighLevel struct {
	name  string
	level int
}

func (a *HumanHighLevel) Attack() {
	fmt.Println("我是", a.name, "等级为", a.level, "造成伤害50000")
}

// 定义一个多态的通用接口，传入不同的对象，调用相同的函数，实现不同的效果
func DoAttack(a IAttack) {
	a.Attack()
}

func main() {
	//var player IAttack // 定义一个包含Attack的接口变量

	lowLevel1 := HumanLowLevel{
		name:  "luenci",
		level: 0,
	}

	HighLevel1 := HumanHighLevel{
		name:  "lynn",
		level: 1000,
	}

	// 两个不同对象调用相同方法
	DoAttack(&HighLevel1)
	DoAttack(&lowLevel1)

	//lowLevel1.Attack()

	// 对player赋值为lowLevel1，接口需要使用指针类型来赋值
	//player = &lowLevel1
	//player.Attack()
}
