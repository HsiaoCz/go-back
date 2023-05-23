package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 单例的懒汉式
type singelton struct{}

var initialized uint32

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
		atomic.AddUint32(&initialized, 1)
	}
	return instance
}

func (s *singelton) SomeThing() {
	fmt.Println("something...")
}

func main() {
	ins := GetInstance()
	ins.SomeThing()
}
