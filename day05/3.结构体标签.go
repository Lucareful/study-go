package main

import (
	"encoding/json"
	"fmt"
)

// 结构体标签由一个或多个键值对组成。
// 键与值使用冒号分隔，值用双引号括起来。键值对之间使用一个空格分隔。

type Teacher struct {
	Name    string `json:"-"`                 // 在使用json编码时候，这个字段不参与编码
	Subject string `json:"subject_name"`      // 在使用json编码时候，这个字段会编码成 subject_name
	Age     int    `json:"age,string"`        // 在使用json编码时候，这个字段类型会变成age and 类型会变为 string
	Address string `json:"address,omitempty"` // 在使用json编码时候，如果这个字段是空的,就会忽略掉不编码

	// 小写的结构体成员在json编码时候会被忽略掉
	gender string
}

func main() {
	t1 := Teacher{
		Name:    "lynn",
		Subject: "math",
		Age:     0,
		Address: "123",
		gender:  "girl",
	}

	marshal, err := json.Marshal(t1)
	if err != nil {
		return
	}
	fmt.Println("编码后结果为:", string(marshal))
}
