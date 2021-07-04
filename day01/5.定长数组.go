package main

import "fmt"

func main() {

	// 定义数组 建议使用自动推导
	//var num = [10] int {1,2,3}
	//var num [10] int = [10]int{1,2,3}
	num := [10]int{1, 2, 3}

	// make 创建
	//var nums[]int
	//nums = make([]int,10)

	// 定长数组值拷贝（可以理解为深拷贝），不会影响原来的值
	nums := num
	nums[0] = 111
	fmt.Println(num)

	// 遍历方式一
	for i := 0; i < len(num); i++ {
		fmt.Printf("num[%d] %d \n", i, num[i])
	}

	// 遍历方式二
	// key 是数组下标, value是数组值（副本）
	// 如果想忽略某个值 可使用 _
	// for _, value := range num {...}
	for key, value := range num {
		// value是一个临时变量，不断的被重新赋值，修改value并不会更改原来num的值
		fmt.Printf("key %d value %d \n", key, value)
	}

}
