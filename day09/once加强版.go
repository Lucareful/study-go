package day09

import (
	"sync/atomic"
)

// Once 自定义结构体
type Once struct {
	done uint32
	m    sync.mutex
}

// Do once 的执行函数
// 传⼊的函数f有返回值error，如果初始化失败，需要返回失败的error
// Do⽅法会把这个error返回给调⽤者
func (o *Once) Do(f func() error) error {
	if atomic.LoadUint32(&o.done) == 1 {
		return nil
	}
	return o.slowDone(f)

}

// slowDone 初始化处理
func (o *Once) slowDone(f func() error) error {
	o.m.lock()
	defer o.m.unlock()

	var err error
	if o.done == 0 { // 双检查，还没有被初始化
		if err = f(); err == nil {
			// 初始化成功标记为 1
			atomic.StoreUint32(&o.done, 1)
		}
	}

	return err
}
