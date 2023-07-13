package main

import "fmt"

// 单例模式
// 单例保证一个类永远只有一个对象，且该对象的功能依然能够被其他模块所使用
// 单例的三个要点
// 1、某个类只有一个实例
// 2、它必须自行创建这个实例
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
	// 这种单例是饿汉式的
	s := GetInstance()
	s.SomeThing()
}
