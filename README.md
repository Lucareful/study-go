# go语法学习笔记

:ledger:go语言学习笔记

```shell
.
├── README.md
├── chan编排模式
│   ├── chan_fanin.go
│   ├── chan_fanout.go
│   └── chan_or-done.go
├── chan进阶使用
│   ├── Token-channe任务编排.go
│   ├── chan-wait&notify.go
│   ├── channel任务编排.go
│   ├── chan实现锁.go
│   ├── select-超时控制.go
│   └── 反射Select-动态处理chan.go
├── context使用
│   ├── main.go
│   └── withValue
├── error和panic
│   ├── panic
│   └── wrapError
├── excelize使用
│   ├── 1.excel操作.go
│   ├── 2.流式写入.go
│   ├── Book1.xlsx
│   ├── demo.xlsx
│   ├── go.mod
│   └── go.sum
├── go.mod
├── go.sum
├── goroutine和chan
│   ├── 1.goroutine(go程).go
│   ├── 2.提前退出go程.go
│   ├── 3.无缓冲管道.go
│   ├── 4.有缓冲管道.go
│   ├── 5.for range读取管道.go
│   ├── 6.判断管道是否关闭.go
│   ├── 7.单向管道.go
│   ├── 8.利用管道控制go程执行情况.go
│   └── 9.select用法.go
├── pointerplay
│   ├── go.mod
│   ├── pointerplay.go
│   └── pointerplay_test.go
├── rpc使用
│   ├── 1.rpc-server.go
│   ├── 2.rpc-client.go
│   ├── 3.design.go
│   └── python-client.py
├── run-go-vet.sh
├── runtime调用堆栈信息
│   └── main.go
├── wire依赖注入
│   ├── case1
│   └── case2
├── 基础语法
│   ├── 1.txt
│   ├── 1.变量.go
│   ├── 10.内存逃逸.go
│   ├── 12-import
│   ├── 2.自增语法.go
│   ├── 3.指针.go
│   ├── 4.string.go
│   ├── 5.定长数组.go
│   ├── 6.slice切片.go
│   ├── 7.切片2.go
│   ├── 8.map字典.go
│   ├── 9.函数func.go
│   └── main.go
├── 小知识实践
│   └── defer解析.go
├── 常见的并发原语
│   ├── RWMutex读写锁.go
│   ├── atomic-queue.go
│   ├── atomic-value.go
│   ├── cond唤醒等待.go
│   ├── context.http
│   ├── context使用.go
│   ├── flagDemo.go
│   ├── model
│   ├── once加强版.go
│   └── once案例1.go
├── 并发原语之Mutex
│   ├── 02.RecursiveMutex.go
│   ├── 03.TokenRecursiveMutex.go
│   ├── 04.TryLock.go
│   ├── 05.SliceQueen.go
│   ├── Deadlock.go
│   └── Mutex使用.go
├── 序列化和反序列化
│   ├── 1.scoket-server.go
│   ├── 2.json的使用.go
│   └── 3.结构体标签.go
├── 结构体标签等语法
│   ├── 1.switch.go
│   ├── 2.标签LABEL.go
│   ├── 3.枚举iota.go
│   ├── 4.结构体.go
│   ├── 5.init函数
│   └── 6.defer函数-匿名函数-文件操作.go
└── 面向对象
    ├── 1.类的封装和绑定方法.go
    ├── 2.类的继承.go
    ├── 3.类成员访问权限
    ├── 4.接口.go
    ├── 5.多态.go
    └── 组合和继承
```
