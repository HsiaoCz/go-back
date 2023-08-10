package main

import "fmt"

// 开闭原则
// 对扩展开放，对修改关闭

// 不符合开闭原则的设计
type Banker struct{}

func (b *Banker) Save() { fmt.Println("银行职员进行了存款的业务") }

func (b *Banker) Trans() { fmt.Println("银行职员进行了转账的业务") }

func (b *Banker) Stack() { fmt.Println("银行职员进行了股票的业务") }

// 符合开闭的设计

// 抽象的银行职员
type IBanker interface {
	DoBuz()
}

// 具体的银行职员
type SaveBanker struct{}

func (s *SaveBanker) DoBuz() { fmt.Println("银行职员进行了存款业务") }

type TransBanker struct{}

func (t *TransBanker) DoBuz() { fmt.Println("银行职员进行了转账业务") }

type StackBanker struct{}

func (s *StackBanker) DoBuz() { fmt.Println("银行职员进行了股票业务") }

// 再抽象一层

func DoBuzz(banker IBanker) {
	banker.DoBuz()
}
