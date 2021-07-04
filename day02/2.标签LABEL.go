package main

import "fmt"

func main() {
	// 标签 LABLE1
	// goto LABEL1 >> 下次进入循环时,i 不会保存之前状态，i=0,重新向下运行
	// break LABEL1 >> 直接跳出指定位置的循环
	// continue LABEL1 >> 跳到指定位置，会记录之前的状态，向下执行

	// 标签名称可以自定义命名
LABEL1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 3 {
				goto LABEL1
			}
			fmt.Println("i", i, ",j", j)
		}
	}
}
