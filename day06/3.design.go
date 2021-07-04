package main

import (
	"net/rpc"
)

type MyInterface interface {
	HelloWorld(string, *string) error
}

// 服务端注册函数
func RegisterName(i MyInterface) error {
	err := rpc.RegisterName("Hello", i)
	if err != nil {
		return err
	}
	return nil
}

type MyClient struct {
	c *rpc.Client
}

func InitClient(addr string) MyClient {
	conn, _ := rpc.Dial("tcp", addr)
	return MyClient{c: conn}
}

// 实现函数参照interface实现
func (this *MyClient) HelloWorld(name string, resp *string) error {
	return this.c.Call("Hello.HelloWorld", name, resp)
}
