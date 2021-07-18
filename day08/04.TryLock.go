package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

// 复制 mutex 定义的常量
const (
	mutexLocked      = 1 << iota // 加锁标识位置 << 左移
	mutexWork                    // 唤醒标识位置
	mutexStaring                 // 锁饥饿标识位置
	mutexWaiterShift = iota      // 标识waiter的起始bit位置
)

// Mutex 扩展
type Mutex struct {
	sync.Mutex
}

// TryLock 尝试获取锁
func (m *Mutex) TryLock() bool {
	if atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), 0, mutexLocked) {
		return true
	}

	// 如果处于唤醒、加锁或者饥饿状态，这次请求就不参与竞争了，返回false
	old := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	if old&(mutexLocked|mutexStaring|mutexWork) != 0 {
		return false
	}
	// 尝试在竞争的状态下请求锁
	new := old | mutexLocked

	return atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), old, new)
}

// Count 统计获取 state 字段的值
func (m *Mutex) Count() int {
	v := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	v = v >> mutexWaiterShift
	v = v + (v & mutexLocked)

	return int(v)
}

// IsLocked 判断锁是否被持有
func (m *Mutex) IsLocked() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&mutexLocked == mutexLocked
}

// IsWork 是否有等待者被唤醒
func (m *Mutex) IsWork() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&mutexWork == mutexWork
}

// IsStarving 锁是否处于饥饿状态
func (m *Mutex) IsStarving() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&mutexStaring == mutexStaring
}

// testCount 测试锁的 state 状态
func testCount() {
	var mu Mutex
	for i := 0; i < 1000; i++ {
		// 开启 1000 个 goroutine
		go func() {
			mu.Lock()
			time.Sleep(time.Second)
			mu.Unlock()
		}()
	}
	time.Sleep(time.Second)
	// 输出锁的信息
	fmt.Printf("waitings:%d \nisLock:%t \nwoken:%t \nStaring:%t \n", mu.Count(), mu.IsLocked(), mu.IsWork(), mu.IsStarving())
}

// try 测试获取函数锁
func try() {
	var mu Mutex
	go func() {
		mu.Lock()
		time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
		mu.Unlock()
	}()

	time.Sleep(time.Second)

	if ok := mu.TryLock(); ok {
		fmt.Println("获取锁成功")
		// do something
		mu.Unlock()
		return
	}
	// 未获取到
	fmt.Println("未获取到锁")
}

func main() {
	// fmt.Println(mutexLocked)
	// fmt.Println(mutexWork)
	// fmt.Println(mutexStaring)
	// fmt.Println(mutexWaiterShift)

	// try()
	testCount()
}
