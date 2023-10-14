package pattern

import "fmt"

// 单例模式
type singelton struct{}

var instance = new(singelton)

func GetInstance() *singelton {
	return instance
}

func (s *singelton) DoSomething() { fmt.Println("单例的某个方法") }

// 单例的懒汉模式
// 没有并发问题
// 但是无论会不会使用，都会分配内存
