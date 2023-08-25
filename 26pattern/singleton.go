package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 单例的懒汉式
// 懒汉式单例需要考虑并发问题

type single struct{}

var instancee *single

var lock sync.Mutex

var initlization uint32

func GetInstancee() *single {
	if atomic.LoadUint32(&initlization) == 1 {
		return instancee
	}
	lock.Lock()
	defer lock.Unlock()
	if initlization == 0 {
		instancee = new(single)
		atomic.StoreUint32(&initlization, 1)
	}
	return instancee
}
func (s *single) DOSomeThing() {
	fmt.Println("单例的某个方法")
}
