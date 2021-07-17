package TokenRecursiveMutex

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// TokenRecursiveMutex 令牌重入锁
type TokenRecursiveMutex struct {
	sync.Mutex
	token int64
	recursive int32
}

// Lock 加锁
func (m *TokenRecursiveMutex) Lock(token int64){
	if atomic.LoadInt64(&m.token) == token{
		m.recursive++
		return
	}
	m.Mutex.Lock()
	atomic.StoreInt64(&m.token, token)
	m.recursive=1
}
// Unlock 解锁
func (m *TokenRecursiveMutex) Unlock(token int64){
	if atomic.LoadInt64(&m.token) != token{
		panic(fmt.Sprintf("wrong of owner(%d):%d",&m.token,token))
	}
	m.recursive--
	if m.recursive == 0{
		return
	}
	atomic.StoreInt64(&m.token, token)
	m.Mutex.Unlock()
}