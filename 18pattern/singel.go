package main

import "fmt"

// 单例模式
// 单例的懒汉式，不管有没有使用，都已经开辟了内存
type singel struct{}

var instance *singel = new(singel)

func GetInstance() *singel {
	return instance
}

func (s *singel) SomeThing() {
	fmt.Println("单例的某个方法")
}
