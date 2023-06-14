package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 单例的懒汉式
// 只有在使用的时候，才会创建实例
// 这里设置一个标志位
var initialized uint32

type singelton struct{}

func (s *singelton) SomeThing() { fmt.Println("do something") }

var lock sync.Mutex

var instance *singelton

func GetInstance() *singelton {
	// 这里要考虑线程安全
	// 避免重复创建
	// 这里通过原子操作实现
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	lock.Lock()
	defer lock.Unlock()

	if initialized == 0 {
		instance = new(singelton)
		atomic.AddUint32(&initialized, 1)
	}

	// 还可以使用sync.Once来实现
	// once.DO(func(){instance=new(singleton)})
	return instance
}

func main() {
	s := GetInstance()
	s.SomeThing()
}
