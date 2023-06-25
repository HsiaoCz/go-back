package main

import "fmt"

// 简单工厂模式显然是不满足开闭原则的
// 我们实现一个满足开闭模式的

// 抽象的水果
type Fruit interface {
	Show()
}

// 抽象的工厂
type AbstractFactory interface {
	CreateFruit() Fruit
}

type Apple struct{}

func (a *Apple) Show() { fmt.Println("this is apple") }

type Banana struct{}

func (b Banana) Show() { fmt.Println("this is banana") }

type Pear struct{}

func (p Pear) Show() { fmt.Println("this is pear") }

// 工厂类

type AppleFactory struct{}

func (a *AppleFactory) CreateFruit() Fruit {
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
