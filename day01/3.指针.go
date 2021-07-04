package main

func main() {
	// go语言在使用指针时，会使用内部的垃圾回收机制（garbage collector），开发人员不需要手动释放内存
	// go语言可以返回栈上指针,程序在编译的时候就确定了变量的分配位置
	// 在编译的时候，如果发现有必要的话，就将变量分配到堆上

	// 定义指针，方式一
	name := "luenci"
	namePtr := &name

	println("指针地址是", namePtr)
	println("指针内容是", *namePtr)

	// 定义指针关键字 new，方式二
	name2Ptr := new(string)
	*name2Ptr = "luenci"

	println("指针地址是", name2Ptr)
	println("指针内容是", *name2Ptr)

	res := testPtr()

	if res == nil {
		println("指针为空")
	} else {
		println("指针值为", *res)
	}

}

// 定义一个函数,返回一个string类型的指针,go语言返回值写在参数列表后面
func testPtr() *string {
	city := "shanghai"
	cityPtr := &city

	return cityPtr
}
