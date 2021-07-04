package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	cmds := os.Args

	fmt.Println("cmds 类型为", reflect.TypeOf(cmds))

	switch cmds[1] {
	// case 中默认加了 break 不需要手动break
	case "luenci":
		// 如果想向下穿透（执行下一个条件中的代码）,使用 fallthrough 关键字
		fmt.Println("i am luenci")
		//fallthrough
	case "Lynn":
		fmt.Println("i am Lynn")
	case "kk":
		fmt.Println("i am kk")
	default:
		fmt.Println("默认值")
	}

	for key, value := range cmds {
		fmt.Println("key", key, "value", value)
	}

}
