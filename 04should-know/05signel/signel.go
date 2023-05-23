package main

import "fmt"

// 单例模式
// 一个类只有一个实例
// 自行创建这个实例
// 自行向外提供这个实例

type signelton struct{}

// 保证一个类只有一个实例
var instance *signelton = new(signelton)

// 自行向外提供这个实例
func GetInstance() *signelton {
	return instance
}

func (s *signelton) SomeThing() {
	fmt.Println("单例的某种方法...")
}

// 这个时单例的饿汉式
// 无论用不用，都初始化了

func main() {
	ins := GetInstance()
	ins.SomeThing()
}
