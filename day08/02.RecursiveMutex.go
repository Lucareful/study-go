package main

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/petermattis/goid"
)

// RecursiveMutex 基于 goroutine id 的可重入锁结构体
type RecursiveMutex struct {
	sync.Mutex
	owner     int64
	recursion int32
}

// Lock 上锁
func (m *RecursiveMutex) Lock() {
	gid := goid.Get()

	// 如果当前持有锁的goroutine就是这次调用的goroutine,说明是重入
	if atomic.LoadInt64(&m.owner) == gid {
		m.recursion++
		return
	}
	m.Mutex.Lock()
	// 获得锁的goroutine第一次调用，记录下它的goroutine id,调用次数加1
	atomic.StoreInt64(&m.owner, gid)

	m.recursion = 1
}

// Unlock 解锁
func (m *RecursiveMutex) Unlock() {
	gid := goid.Get()

	// 非持有锁的 goroutine 尝试释放锁 ，panic
	if atomic.LoadInt64(&m.owner) != gid {
		panic(fmt.Sprintf("worng the ower(%d) : %d", m.owner, gid))
	}
	m.recursion--
	if m.recursion != 0 {
		// 如果这个goroutine还没有完全释放，则直接返回
		return
	}
	atomic.StoreInt64(&m.owner, -1)
	m.Mutex.Unlock()
}
