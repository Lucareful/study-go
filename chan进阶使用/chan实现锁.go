package main

import (
	"fmt"
	"time"
)

// Mutex 结构体
type Mutex struct {
	ch chan struct{}
}

// newMutex 初始化锁
func newMutex() *Mutex {
	mu := &Mutex{make(chan struct{}, 1)}
	mu.ch <- struct{}{}
	return mu
}

// Lock 请求锁，直到获取到
func (mu *Mutex) Lock() {
	<-mu.ch
}

// Unlock 解锁
func (mu *Mutex) Unlock() {
	select {
	case mu.ch <- struct{}{}:
	default:
		panic("unlock of unlocked mutex")
	}
}

// TryLock 尝试获取锁
func (mu *Mutex) TryLock(times int) bool {
	select {
	case <-mu.ch:
		return true
	case <-time.After(time.Duration(times) * time.Second):
		return false
	}
	return false
}

// isLocked 锁是否被持有
func (mu *Mutex) isLocked() bool {
	return len(mu.ch) == 0
}

func main() {
	m := newMutex()
	ok := m.TryLock(5)
	fmt.Printf("locked v %v\n", ok)
	ok = m.TryLock(5)
	fmt.Printf("locked %v\n", ok)
}
