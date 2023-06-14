package main

import "fmt"

// 单例模式
// 单例保证一个类永远只有一个对象，且该对象的功能依然能够被其他模块使用
// 单例有三个要点
// 1、某个类只有一个实例
// 2、它必须自行创建这个实例
// 3、它必须向整个系统提供这个实例

// 保证一个类只有一个实例，那么这个类就不能共有化

type singleton struct{}

// 这个类必须像外提供一个实例

var instance *singleton = new(singleton)

// 这个类必须向外提供这个实例

func GetInstance() *singleton {
	return instance
}

func (s *singleton) SomeThing() { fmt.Println("do something") }

func main() {
	// 单例的饿汉模式
	// 这种单例无论使用不使用 都已经开辟了内存
	// 饿汉式的好处是，不会出现线程并发创建，出现多个单例
	// 缺点是即使业务没有被使用，也申请了内存
	s := GetInstance()
	s.SomeThing()
}
