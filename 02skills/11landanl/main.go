package main

import (
	"sync"
	"sync/atomic"
)

// 单例的懒汉式

type singleton struct{}

var instance *singleton

var lock sync.Mutex

var instag uint32

func GetInstance() *singleton {
	if atomic.LoadUint32(&instag) == 1 {
		return instance
	}

	lock.Lock()
	defer lock.Unlock()

	if instag == 0 {
		instance = new(singleton)
		atomic.StoreUint32(&instag, 1)
	}
	return instance
}
