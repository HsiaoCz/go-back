package pattern

import "fmt"

// 开闭原则
// 对扩展开放，对修改关闭

// 抽象的银行职员类
type Banker interface {
	DoBuz()
}

// 具体的银行职员
type TransBanker struct{}

func (t *TransBanker) DoBuz() {
	fmt.Println("执行转账操作")
}

// 存款的银行职员类
type SaveBanker struct{}

func (s *SaveBanker) DoBuz() {
	fmt.Println("执行存款操作")
}

// 股票业务的银行职员类
type StackBanker struct{}

func (s *StackBanker) DoBuz() {
	fmt.Println("执行股票业务的银行职员")
}
