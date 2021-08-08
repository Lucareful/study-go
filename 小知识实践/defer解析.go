package main

import "fmt"

func testScope() {
	{
		defer fmt.Println("defer runs")
		fmt.Println("block ends")
	}

	fmt.Println("main ends")
	// 输出:
	// block ends
	// main ends
	// defer runs
}

func main() {
	i := 1
	defer func() {
		fmt.Println(i)
	}()
	i++
	// 输出 2

	testScope()

}

// func main(){
// 	i:=1
// 	defer func(i int) {
// 		fmt.Println(i)
// 	}(i)
// 	i++

// 	// 输出 1
// }
