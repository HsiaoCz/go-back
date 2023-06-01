package main

import "fmt"

// 饿汉式单例
// 饿汉式单例，无论使不使用都创建了内存

type singleton struct{}

var instance *singleton = new(singleton)

func GetInstance() *singleton {
	return instance
}

func (s *singleton) SomeThing() {
	fmt.Println("do something...")
}

func main() {
	ins := GetInstance()
	ins.SomeThing()
}
