package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 单例的懒汉式
// 懒汉式单例
// 只有在使用的时候，才会创建实例
// 这里设置一个标记位
var initialized uint32

type singelton struct{}

var lock sync.Mutex

var instance *singelton

func GetInstance() *singelton {
	// 这里要考虑线程安全
	// 避免重复创建
	// 最简单的方法是加互斥锁
	// sync.lock
	// 还可以通过原子操作来实现
	// 如果标记被设置，直接返回
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	// 如果没有，则申请加锁
	lock.Lock()
	defer lock.Unlock()

	// 如果为0说明还没有实例化
	// 返回一个实例，并将标志加1
	if initialized == 0 {
		instance = new(singelton)
		atomic.StoreUint32(&initialized, 1)
	}

	// 还有一种更简单的办法来实现
	// 使用sync.once来实现
	// 本质上sync.once()就是刚刚这一套操作
	// once.DO(func(){ instance=new(singelton) })
	// 需要注意的是，这里需要提前声明一下once
	return instance
}

func (s *singelton) SomeThing() {
	fmt.Println("some thing...")
}

func main() {
	s := GetInstance()
	s.SomeThing()
}
