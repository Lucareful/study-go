package main

import "fmt"

func main() {
	// slice 切片 底层也是数组，可以动态改变长度
	names := []string{"lynn", "luenci"}

	// 对于一个切片不仅仅只有长度概念，还有'容量'概念
	// 追加元素前 长度是 2 容量是 2
	// 追加元素前 长度是 3 容量是 4
	// 在一定量级的时候,动态追加元素, 容量一般是2倍增长
	fmt.Printf("追加元素前 长度是 %d 容量是 %d \n", len(names), cap(names))

	// append 追加元素
	names = append(names, "kk")
	fmt.Printf("追加元素前 长度是 %d 容量是 %d \n", len(names), cap(names))

	// 在一定量级的时候,动态追加元素, 容量一般是2倍增长
	num1 := [] int {}
	for i := 0; i < 50; i++ {
		num1 = append(num1, i)
		fmt.Printf("长度 %d 容量 %d \n", len(num1), cap(num1))
	}

	// 使用make创建数组
	//mnu2 := make([]int, 10)
	//mnu2
}
