package main

import (
	"fmt"
	"sync"
)

// 使用sync.Once来实现单例

type singellton struct{}

var instancees *singellton

var once sync.Once

func GetInstancees() *singellton {
	once.Do(func() {
		instancees = new(singellton)
	})
	return instancees
}

func (s *singellton) DoSomething() {
	fmt.Println("单例的某个方法")
}
