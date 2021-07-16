package main

import "sync"

// TokenRecursiveMutex 令牌重入锁
type TokenRecursiveMutex struct {
	sync.Mutex
	token int64
	recursive int32
}

// Lock 加锁
func (* TokenRecursiveMutex) Lock(){
	
}

func main()  {
	
}
