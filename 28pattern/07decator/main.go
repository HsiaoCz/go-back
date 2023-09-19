package main

import "fmt"

// 装饰器模式
// 动态的给一个对象添加一些额外的职责

// 抽象的手机

type Phone interface {
	Show()
}

// 装饰器的基础类
type Decorator struct {
	phone Phone
}

func (d *Decorator) Show() {}

// 具体的手机
type XiaoMi struct{}

func (x *XiaoMi) Show() { fmt.Println("小米手机") }

type Huawei struct{}

func (h *Huawei) Show() { fmt.Println("华为手机") }

// 具体的装饰器类
type MoDecorator struct {
	Decorator
}

func NewMoDecorator(phone Phone) Phone {
	return &MoDecorator{Decorator{phone: phone}}
}

func (md *MoDecorator) Show() {
	md.phone.Show()
	fmt.Println("戴上了手机壳的手机...")
}

type Kedecorator struct {
	Decorator
}

func NewKeDecorator(phone Phone) Phone {
	return &Kedecorator{Decorator{phone: phone}}
}

func (kd *Kedecorator) Show() {
	kd.phone.Show()
	fmt.Println("贴上了膜的手机...")
}

func main() {
	huawei := new(Huawei)
	huawei.Show()

	// 装饰
	keDecorator := NewKeDecorator(huawei)
	keDecorator.Show()

	moDecorator := NewMoDecorator(huawei)
	moDecorator.Show()
}
