package main

// 抽象的处理者
type Department interface {
	Execute(*Patient)
	SetNext(Department)
}
