#!/usr/bin/python
# -*- coding: UTF-8 -*-

import json
import socket               # 导入 socket 模块

s = socket.socket()         # 创建 socket 对象
host = "127.0.0.1" # 获取本地主机名
port = 8800                # 设置端口号

s.connect((host, port))
data = json.dumps({"method":"luenci.HelloWorld","params":["lynn"],"id":0})
print(data)
s.send(data)
res = s.recv(1024)
print(res)
print(type(res))
s.close()