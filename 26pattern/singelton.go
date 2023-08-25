package main

import "fmt"

// 单例模式
// 饿汉式单例
// 单例的几个特点
// 一个类只有一个实例
// 它必须自行创建这个实例
// 它必须自行向整个系统提供这个实例

type singel struct{}

var instance = new(singel)

func GetInstance() *singel {
	return instance
}

func (s *singel) DOSomeThing() {
	fmt.Println("单例的某个方法")
}

// 饿汉式单例的好处是，没有并发安全问题