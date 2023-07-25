package main

import "fmt"

// 抽象工厂模式
// 针对产品族和产品等级结构

// 抽象的水果
// 抽象的是产品等级结构

type AbstractApple interface {
	ShowApple()
}

type AbstractBanana interface {
	ShowBanana()
}

type AbstractPear interface {
	ShowPear()
}

// 抽象的工厂
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
	china_apple := new(ChinaApple)
	return china_apple
}
func (cf *ChinaFactory) CreateBanana() AbstractBanana {
	china_banana := new(ChinaBanana)
	return china_banana
}
func (cf *ChinaFactory) CreatePear() AbstractPear {
	china_pear := new(ChinaPear)
	return china_pear
}

// 同样的原理可以实现美国的产品族和日本的产品族

// 业务实现层
func main() {
	// 抽象工厂针对产品族和产品等级结构是符合开闭原则的
	// 但是当想增加一个产品的时候，就不符合了
	// 这里比如需要一个中国的苹果
	cf := new(ChinaFactory)
	china_apple := cf.CreateApple()
	china_apple.ShowApple()

}
