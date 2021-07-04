package main

import (
	"fmt"
	"net"
	"net/rpc/jsonrpc"
)

type Hello struct {
}

func (this *Hello) HelloWorld(name string, resp *string) error {
	*resp = name + " 你好！"
	return nil
}

func main() {
	// 1.注册服务
	//if err := rpc.RegisterName("luenci", new(Hello)); err != nil {
	//	fmt.Println("注册服务失败:", err)
	//	return
	//}
	if err := RegisterName(new(Hello)); err != nil {
		fmt.Println("服务注册失败,error", err)
		return
	}

	// 2.设置监听着
	listenner, err := net.Listen("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("设置监听者失败:", err)
		return
	}
	defer listenner.Close()

	// 3.建立链接
	conn, err := listenner.Accept()
	if err != nil {
		fmt.Println("建立链接失败:", err)
		return
	}
	defer conn.Close()

	// 4.绑定服务 jsonrpc
	//rpc.ServeConn(conn)
	jsonrpc.ServeConn(conn)
}
