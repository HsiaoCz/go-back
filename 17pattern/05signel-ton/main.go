package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 单例的懒汉式
// 只有在使用的时候，才创建实例
var initialized uint32

type singelton struct{}

var lock sync.Mutex

var instance *singelton

func GetInstance() *singelton {
	// 这里要考虑线程安全
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	lock.Lock()
	defer lock.Unlock()

	if initialized == 0 {
		instance = new(singelton)
		atomic.StoreUint32(&initialized, 1)
	}
	return instance
}

func (s *singelton) SomeThing() { fmt.Println("some thing") }

func main() {
	s := GetInstance()
	s.SomeThing()
}
