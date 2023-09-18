package main

import (
	"sync"
	"sync/atomic"
)

// 单例的懒汉式

type singleton struct{}

var lock sync.Mutex

var instance *singleton

var initlize uint32

func GetInstance() *singleton {
	if atomic.LoadUint32(&initlize) == 1 {
		return instance
	}
	lock.Lock()
	defer lock.Unlock()

	if initlize == 0 {
		instance = new(singleton)
		atomic.StoreUint32(&initlize, 1)
	}
	return instance
}
