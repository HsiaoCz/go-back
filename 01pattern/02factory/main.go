package main

import "fmt"

// 工厂方法模式
// 简单工厂方法模式不满足开闭
// 所以可以给简单工厂方法模式加上开闭原则
// 还是以刚刚的水果举例
// 这里要满足开闭，就得抽象出工厂的接口

// 这里就满足了开闭原则
// 这时候我们需要增加水果，或是增加工厂就是完全符合开闭的
// 缺点在于类会比较多，代码比较麻烦

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

func (a Apple) Show() { fmt.Println("this is apple") }

type Banana struct{}

func (b Banana) Show() { fmt.Println("this is banana") }

type Pear struct{}

func (p Pear) Show() { fmt.Println("this is pear") }

// 工厂类
// 这里针对不同种类的水果实现不同的工厂
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

// 业务层

func main() {

	// 这里比如需要一个苹果
	af := new(AppleFactory)
	apple := af.CreateFruit()
	apple.Show()
}
