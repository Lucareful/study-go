package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	/* --------------------- 方式一 ------------------*/
	// 通用的序列化和反序列化 -- json、protobuf
	// 使用 Dail 链接服务器 -- Dail()
	//conn, err := rpc.Dial("tcp", "127.0.0.1:8800")

	//conn, err := jsonrpc.Dial("tcp", "127.0.0.1:8800")
	//if err != nil {
	//	fmt.Println("Dail error:", err)
	//	return
	//}
	//defer conn.Close()
	//
	//// 调用远程函数
	//var response string // 接收返回值 -- 传出参数
	//if err := conn.Call("luenci.HelloWorld", "lynn", &response); err != nil {
	//	fmt.Println("Call error", err)
	//	return
	//}
	//fmt.Println("response", response)

	/* --------------------- 方式二 ------------------*/

	MyClient := InitClient("127.0.0.1:8800")
	defer func(c *rpc.Client) {
		err := c.Close()
		if err != nil {
			fmt.Println("关闭链接出错.error", err)
			return
		}
	}(MyClient.c)
	
	var resp string
	err := MyClient.HelloWorld("luenci", &resp)
	if err != nil {
		fmt.Println("出错了,error", err)
		return
	}
}
