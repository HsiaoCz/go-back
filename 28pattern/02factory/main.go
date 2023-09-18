package main

import "fmt"

// 工厂模式
// 简单工厂方法模式不满足开闭原则
// 工厂方法模式在简单工厂方法模式之上增加了开闭原则

type Fruit interface {
	Show()
}

// 抽象的工厂
type AbstractFactory interface {
	CreateFruit() Fruit
}

// 实现层
type Apple struct{}

func (a *Apple) Show() {
	fmt.Println("this is apple")
}

type Banana struct{}

func (b *Banana) Show() {
	fmt.Println("this is banana")
}

type Pear struct{}

func (p *Pear) Show() {
	fmt.Println("this is pear")
}

// 工厂类
type AppleFactory struct{}

func (af *AppleFactory) CreateFruit() Fruit {
	return new(Apple)
}

type BananaFactory struct{}

func (bf *BananaFactory) CreateFruit() Fruit {
	return new(Banana)
}

type PearFactory struct{}

func (pf *PearFactory) CreateFruit() Fruit {
	return new(Pear)
}

func main() {
	// 这里比如需要一个苹果
	af := new(AppleFactory)
	apple := af.CreateFruit()
	apple.Show()

	bf := new(BananaFactory)
	banana := bf.CreateFruit()
	banana.Show()
}
