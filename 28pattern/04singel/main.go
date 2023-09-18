package main

import "fmt"

// 单例
// 单例的三个特点
// 保证一个类只有一个实例
// 它必须自行创建这个实例
// 它必须自行向整个系统提供这个实例

type singleton struct{}

var instance *singleton = new(singleton)

func GetInstance() *singleton {
	return instance
}

func (s *singleton) SomeThing() {
	fmt.Println("单例的某个方法")
}
