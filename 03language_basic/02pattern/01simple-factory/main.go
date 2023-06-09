package main

import "fmt"

// 简单工厂模式
// 简单工厂实现了创建和使用的分离
// 但是简单工厂不满足开闭原则
type Fruit interface {
	Show()
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
type SimpleFactory struct{}

func (s *SimpleFactory) NewFruit(kind string) Fruit {
	var fruit Fruit
	if kind == "apple" {
		fruit = &Apple{}
	}
	if kind == "banana" {
		fruit = &Banana{}
	}
	if kind == "pear" {
		fruit = &Pear{}
	}
	return fruit
}

func main() {
	sf := &SimpleFactory{}
	fruit := sf.NewFruit("apple")
	fruit.Show()
}
