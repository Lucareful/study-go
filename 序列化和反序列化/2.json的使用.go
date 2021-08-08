package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// 将 结构体 --> 字符串  编码
// 将 字符串 --> 结构体  解码

// 结构体的成员须大写，不然不参与编码

type Student struct {
	Name  string
	Sex   string
	Age   int
	Score int
}

func main() {
	st1 := Student{
		Name:  "luenci",
		Sex:   "man",
		Age:   22,
		Score: 99,
	}
	// 编码 序列化
	encodeInfo, err := json.Marshal(st1)
	if err != nil {
		fmt.Println("序列化发生错误,error", err)
		return
	}
	fmt.Println(reflect.TypeOf(encodeInfo))
	fmt.Println(string(encodeInfo))

	// 解码 反序列化
	var st2 Student

	if err := json.Unmarshal([]byte(encodeInfo), &st2); err != nil {
		fmt.Println("反序列化发生错误,", err)
		return
	}
	fmt.Println(st2.Name)

}
