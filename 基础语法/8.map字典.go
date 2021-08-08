package main

import (
	"fmt"
)

func main() {
	// key -> value key是经过hash运算，是无顺序的
	// 使用map之前一定要分配空间

	var idName map[int]string

	idName = make(map[int]string, 10)

	idName[0] = "luenci"
	idName[1] = "lynn"

	//遍历map
	for key, value := range idName {
		fmt.Println("id为", key, "value为", value)
	}

	// 确定key是否在map中
	// 在map中不存在访问越界，访问一个不存在的key，map不会崩溃，会返回零值
	// 零值: bool-》true/false string-》空 int-》0
	name := idName[9]
	fmt.Println("name值为", name)

	idScore := make(map[int]int, 10)
	fmt.Println("int零值", idScore[0])

	idFalse := make(map[int]bool, 10)
	fmt.Println("bool零值", idFalse[0])

	// map无法通过获取value来判断这个对应的key是否存在
	// 可用过下面方法来确定是否存在key  ok -》 bool值
	value, ok := idName[99]
	fmt.Println("ok值为", ok)
	if ok {
		fmt.Println("idname[99]存在,值为", value)
	} else {
		fmt.Println("key不存在")
	}

	// 删除map中的元素
	// 删除不存在的key也不会报错
	delete(idName, 0)
	fmt.Println("删除后的map为", idName)

	// 并发处理时需要对map进行上锁 TODO

}
