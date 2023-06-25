package main

import "fmt"

// 简单工厂模式
// 简单工厂模式实现了对象创建与使用的分离

type Fruit interface {
	Show()
}

// 实现层
type Apple struct{}

func (a *Apple) Show() { fmt.Println("this is apple") }

type Banana struct{}

func (b *Banana) Show() { fmt.Println("this is banana") }

type Pear struct{}

func (p *Pear) Show() { fmt.Println("this is pear") }

// 工厂类
type SimpleFactory struct{}

func (s *SimpleFactory) NewFruit(kind string) Fruit {
	var fruit Fruit

	if kind == "apple" {
		fruit = new(Apple)
	}
	if kind == "banana" {
		fruit = new(Banana)
	}
	if kind == "pear" {
		fruit = new(Pear)
	}
	return fruit
}

func main() {
	sf := new(SimpleFactory)
	fruit := sf.NewFruit("apple")
	fruit.Show()
}
