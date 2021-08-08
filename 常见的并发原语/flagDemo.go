package main

import (
	"flag"
	"fmt"
)

func main() {

	var name = flag.String("name", "everyone", "The greeting object.")
	// 解析参数
	flag.Parse()
	fmt.Printf("Hello, %s!\n", *name)
}
