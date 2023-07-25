package main

import (
	"fmt"
)

// 工厂方法模式
// 在简单工厂模式上加入了开闭原则

// 抽象的水果
type Fruit interface {
	Show()
}

// 抽象的工厂
type AbstractFactory interface {
	CreateFruit() Fruit
}

// 实现层
// apple
type Apple struct{}

func (a *Apple) Show() { fmt.Println("this is apple") }

type Banana struct{}

func (b *Banana) Show() { fmt.Println("this is banana") }

type Pear struct{}

func (p *Pear) Show() { fmt.Println("this is pear") }

// 工厂类
type AppleFactory struct{}

func (af *AppleFactory) CreateFruit() Fruit {
	fruit := new(Apple)
	return fruit
}

type BananaFactory struct{}

func (bf *BananaFactory) CreateFruit() Fruit {
	fruit := new(Banana)
	return fruit
}

type PearFactory struct{}

func (pf *PearFactory) CreateFruit() Fruit {
	fruit := new(Pear)
	return fruit
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
