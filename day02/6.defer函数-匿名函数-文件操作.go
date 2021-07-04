package main

import (
	"fmt"
	"os"
)

func main() {
	readFile("4.结构体.go")
}

func readFile(filename string) {
	// os.Open 打开文件，会返回一个文件指针和err信息，如果无错误 err 是 nil
	// defer 当你的堆栈退出的时候会调用 （必须函数调用结束）
	// func (){...}() 不声明函数名表示匿名函数 后面加括号() 调用

	fp, err := os.Open(filename)
	defer func() {
		fmt.Println("文件关闭！")
		_ = fp.Close()
	}()

	defer fmt.Println("00000")
	defer fmt.Println("00001")
	defer fmt.Println("00002")

	if err != nil {
		fmt.Println("文件读取错误,error:", err)
		return
	}
	buf := make([]byte, 1024)
	n, _ := fp.Read(buf)
	fmt.Println(string(buf))
	fmt.Println("读取文件长度为", n)
}
