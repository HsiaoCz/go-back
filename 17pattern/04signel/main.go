package main

import "fmt"

// 单例模式
// 单例有三个要点
// 1.某个类只能有一个实例
// 2.它必须自行创建这个实例
// 3.他必须自行向整个系统提供这个实例

type singleton struct{}

var instance *singleton = new(singleton)

func GetInstance() *singleton {
	return instance
}

func (s *singleton) SomeThing() { fmt.Println("单例对象的某些方法") }

func main() {
	s := GetInstance()
	s.SomeThing()
}
