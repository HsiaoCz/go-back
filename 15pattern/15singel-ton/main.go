package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 单例的懒汉式
var initialized uint32

type singelton struct{}

var lock sync.Mutex

var instance *singelton

func GetInstance() *singelton {
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

func (s *singelton) SomeThing() {
	fmt.Println("单例的某些方法")
}

func main() {
	s := GetInstance()
	s.SomeThing()
}
