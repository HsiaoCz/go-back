package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 单例模式的饿汉式
type singeltone struct{}

var instancee *singeltone

var lock sync.Mutex

var initlize uint32

func GetInstancee() *singeltone {
	if initlize == 1 {
		return instancee
	}

	lock.Lock()
	defer lock.Unlock()

	if atomic.LoadUint32(&initlize) == 0 {
		instancee = new(singeltone)
		atomic.StoreUint32(&initlize, 1)
	}
	return instancee
}

func (s *singeltone) DoSomeThing() {
	fmt.Println("单例的某个方法")
}
