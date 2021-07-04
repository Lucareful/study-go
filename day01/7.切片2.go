package main

import "fmt"

func main() {
	citys := [6]string{"shanghai", "beijing", "wuhan", "hangzhou", "nanjing", "shenzhen"}
	fmt.Printf("city 源地址 %p \n", &citys[0])

	// 使用索引切片访问 前开后闭 浅拷贝（副本）
	loveCity := citys[0:1]
	//fmt.Println("my love city", loveCity)
	fmt.Printf("切片地址 %p \n", &loveCity[0])
	loveCity[0] = "kk"

	fmt.Println("修改元素",loveCity)
	fmt.Println("citys",citys[0])

	// 如果想拷贝一份独立与源数组的 使用自由函数 copy()
	loveCitys := copy(loveCity, citys[:])
	fmt.Printf("copy 地址 %p \n", &loveCitys)

}
