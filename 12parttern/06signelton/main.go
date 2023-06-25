package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 之前那种单例是饿汉式的
// 无论有没有使用，都已经分配内存了

var initialized uint32

type singelton struct{}

var lock sync.Mutex

var instance *singelton

func GetInstance() *singelton {
	// 这里考虑线程安全
	// 避免重复创建
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	lock.Lock()
	defer lock.Unlock()

	if initialized == 0 {
		instance = new(singelton)
		atomic.AddUint32(&initialized, 1)
	}
	return instance
}

func (s *singelton) SomeThing() { fmt.Println("some thing") }

func main() {
	s := GetInstance()
	s.SomeThing()
}
