package main

import "fmt"

type signelton struct{}

var instance = new(signelton)

func GetInstance() *signelton {
	return instance
}

func (s *signelton) DoSomeThing() {
	fmt.Println("单例的某个方法")
}
