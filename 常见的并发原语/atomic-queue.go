package queue

import (
	"sync/atomic"
	"unsafe"
)

// LKQueue lock-free 的queue
type LKQueue struct {
	head unsafe.Pointer
	tail unsafe.Pointer
}

// node 通过链表实现，这个数据结构代表链表中的节点
type node struct {
	value interface{}
	next  unsafe.Pointer
}

// NewLKQueue 初始化 queue
func NewLKQueue() *LKQueue {
	n := unsafe.Pointer(&node{})
	return &LKQueue{head: n, tail: n}
}

// Enqueue 出队，没有元素则返回 nil
func (q *LKQueue) Enqueue(v interface{}) interface{} {
	n := &node{value: v}
	for {
		head := load(&q.head)
		tail := load(&q.tail)
		next := load(&head.next)
		if head == load(&q.head) { // head 一致
			if head == tail { // head 和 tail 一样
				if next == nil { // 空队列
					return nil
				}
				// 尾指针还没有调整，尝试调整它指向下一个
				cas(&q.tail, tail, next)
			} else {
				// 读取队列的数据
				v := next.value
				// 数据出队，头指针移先下一个
				if cas(&q.head, head, next) {
					return v
				}
			}
		}
	}
}

// load 将 unsafe.Pointer 原子加载转换成 node
func load(p *unsafe.Pointer) (n *node) {
	return (*node)(atomic.LoadPointer(p))
}

// 封装cas 避免直接将 *node 转化为 unsafe.Pointer
func cas(p *unsafe.Pointer, old, new *node) (ok bool) {
	return atomic.CompareAndSwapPointer(
		p, unsafe.Pointer(old), unsafe.Pointer(new))
}
