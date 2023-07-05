package main

import "fmt"

// 单例模式
// 保证一个类只有一个对象
// 且该对象的功能任然能够被其他模块所用

// 单例有三个要点
// 1、某个类只能有一个实例
// 2、他必须自行创建这个实例
// 3、它必须自行向整个系统提供这个实例

type singleton struct{}

var instance *singleton = new(singleton)

func GetInstance() *singleton {
	return instance
}

func (s *singleton) SomeThing() {
	fmt.Println("单例对象的某些方法")
}

func main() {
	//	这种单例是饿汉式单例
	// 无论使用不使用都开辟了内存
	// 好处是不会出现线程并发访问出现多个单例的问题
	// 缺点是即使业务没有被使用，也始终申请了内存

	s := GetInstance()
	s.SomeThing()
}
