package main

import "fmt"

// 简单工厂模式

// 抽象的水果类
type Fruit interface {
	Show()
}

// 具体的水果类
type Apple struct{}

func (a *Apple) Show() { fmt.Println("this is apple") }

type Banana struct{}

func (b *Banana) Show() { fmt.Println("this is banana") }

type Pear struct{}

func (p *Pear) Show() { fmt.Println("this is pear") }

// 工厂类
type SimpleFactory struct{}

func (s *SimpleFactory) CreateFruit(kind string) Fruit {
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
