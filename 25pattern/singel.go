package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 单例的懒汉式
// 懒汉式的单例，只有在使用的时候，才会创建实例
var initialized uint32

type singelton struct{}

var lock sync.Mutex

var instancee *singelton

func GetInstancee() *singelton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instancee
	}
	lock.Lock()
	defer lock.Unlock()

	if initialized == 0 {
		instancee = new(singelton)
		atomic.StoreUint32(&initialized, 1)
	}
	return instancee
}

func (s *singelton) SomeThing() {
	fmt.Println("some thing")
}
