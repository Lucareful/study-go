package sub

import "fmt"

// 在go语言同一层级目录，不允许出现多个 package 名称

func init() {
	fmt.Println("sub库下的init函数")
}

func Sub(a, b int) int {
	test5()
	return a - b
}
