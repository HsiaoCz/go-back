package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 饿汉式单例
// 只有在使用的时候才开辟内存

type singleton struct{}

var instancet *singleton

var initialize uint32

var lock sync.Mutex

func GetInstancet() *singleton {
	if atomic.LoadUint32(&initialize) == 1 {
		return instancet
	}

	lock.Lock()
	defer lock.Unlock()

	if initialize == 0 {
		instancet = new(singleton)
		atomic.StoreUint32(&initialize, 1)
	}
	return instancet
}

func (s *singleton) DoSomeThing() { fmt.Println("单例的某个方法") }
