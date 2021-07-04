package main

// sub 是文件夹名，也是package名
import (
	"fmt"
	"luenci/github.com/day01/12-import/sub"
	// SUB "luenci/github.com/day01/12-import/sub" SUB使我们重命名的包名 可以 SUB.Sub 调用
	// . "luenci/github.com/day01/12-import/sub" 导入sub文件中的所有函数，可直接使用函数 Sub (不建议使用,可能会混淆)
)

func main() {
	// 调用方法：包名.函数
	// 想调用某个包里面的函数，此函数必须首字母大写（public），
	// 首字母小写的函数（private）只能在相同包名的文件中使用

	res := sub.Sub(20, 10)
	fmt.Println("sub res", res)


}
