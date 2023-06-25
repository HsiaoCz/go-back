package main

import (
	"fmt"
)

// 抽象工厂模式
// 针对产品族和产品等级结构

// 抽象的水果
// 抽象的是产品的等级结构
type AbstractApple interface {
	ShowApple()
}

type AbstractBanana interface {
	ShowBanana()
}

type AbstractPear interface {
	ShowPear()
}

// 抽象得到工厂

type AbstractFactory interface {
	CreateApple() AbstractApple
	CreateBanana() AbstractBanana
	CreatePear() AbstractPear
}

// 实现类
// 中国的产品族
type ChinaApple struct{}

func (ca *ChinaApple) ShowApple() { fmt.Println("this is china apple") }

type ChinaBanana struct{}

func (cb *ChinaBanana) ShowBanana() { fmt.Println("this is china banana") }

type ChinaPear struct{}

func (cp *ChinaPear) ShowPear() { fmt.Println("this is china pear") }

// 中国的工厂
type ChinaFactory struct{}

func (cf *ChinaFactory) CreateApple() AbstractApple {
	// china_apple := new(ChinaApple)
	// return china_apple
	return new(ChinaApple)
}

func (cf *ChinaFactory) CreateBanana() AbstractBanana {
	return new(ChinaBanana)
}

func (cf *ChinaFactory) CreatePear() AbstractPear {
	return new(ChinaPear)
}

// 同样的原理可以实现日本和美国的产品族

func main() {
	cf := new(ChinaFactory)
	cApple := cf.CreateApple()
	cApple.ShowApple()
}
