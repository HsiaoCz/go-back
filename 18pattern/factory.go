package main

import "fmt"

// 工厂模式

type FactoryFruit interface {
	Show()
}

type FApple struct{}

func (fa *FApple) Show() { fmt.Println("this is apple") }

type FBanana struct{}

func (fb *FBanana) Show() { fmt.Println("this is banana") }

type FPear struct{}

func (fp *FPear) Show() { fmt.Println("this is pear") }

// 抽象的工厂类
type FFactory interface {
	CreateFFruit() FactoryFruit
}

type AppleFactory struct{}

func (af *AppleFactory) CreateFFruit() FactoryFruit {
	return new(FApple)
}

type BananaFactory struct{}

func (bf *BananaFactory) CreateFFruit() FactoryFruit {
	return new(FBanana)
}

type PearFactory struct{}

func (pf *PearFactory) CreateFFruit() FactoryFruit {
	return new(FPear)
}
