package main

import (
	"sync"
	"sync/atomic"
)

type signelton struct{}

var instance *signelton

var initlization uint32

var lock sync.Mutex

func GetInstance() *signelton {
	if atomic.LoadUint32(&initlization) == 1 {
		return instance
	}
	lock.Lock()
	defer lock.Unlock()

	if initlization == 0 {
		instance = new(signelton)
		atomic.StoreUint32(&initlization, 1)
	}
	return instance
}
