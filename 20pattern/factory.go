package main

import "fmt"

// 工厂模式
// 简单工厂模式不符合开闭
// 给简单工厂模式加上开闭原则，就是工厂模式

// 抽象的水果类
type FFruit interface {
	ShowFruit()
}

// 抽象的工厂类
type Factory interface {
	CreateFruit()
}

// 具体的水果类
type FApple struct{}

func (f *FApple) ShowFruit() { fmt.Println("this is fruit") }

type FBanana struct{}

func (f *FBanana) ShowFruit() { fmt.Println("this is banana") }

type FPear struct{}

func (f *FPear) ShowFruit() { fmt.Println("this is pear") }

// 具体的工厂类
type AppleFactory struct{}

func (a *AppleFactory) CreateFruit() FFruit {
	return new(FApple)
}

type BananaFactory struct{}

func (b *BananaFactory) CreateFruit() FFruit {
	return new(FBanana)
}

type PearFactory struct{}

func (p *PearFactory) CreateFruit() FFruit {
	return new(FPear)
}
