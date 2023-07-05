package main

import "fmt"

// 抽象工厂模式
// 针对产品族和产品等级结构
// 比如中国的产品族和中国的产品等级结构
// 产品族，产地相同的不同产品
// 产品等级结构，不同产地的相同产品

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

func main() {
	// 抽象工厂针对产品族和产品等级结构是符合开闭原则的
	// 但是当想增加一个产品的时候，就不符合了

	cf := new(ChinaFactory)
	china_apple := cf.CreateApple()
	china_apple.ShowApple()
}
