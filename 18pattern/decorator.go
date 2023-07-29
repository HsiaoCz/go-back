package main

import "fmt"

// 装饰器模式
// 动态的为一个对象增加一些额外的职责

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

func (x *XiaoMi) Show() { fmt.Println("秀出了小米手机") }

type HuaWei struct{}

func (h *HuaWei) Show() { fmt.Println("秀出了华为手机") }

// 具体的装饰器
type MoDecorator struct {
	Decorator
}

func NewMoDecorator(phone Phone) Phone {
	return &MoDecorator{Decorator{phone: phone}}
}

func (md *MoDecorator) Show() {
	md.phone.Show()
	fmt.Println("戴上了壳的手机")
}

// 具体的装饰器
type keDecorator struct {
	Decorator
}

func NewKeDecorator(phone Phone) Phone {
	return &keDecorator{Decorator{phone: phone}}
}

func (kd *keDecorator) Show() {
	kd.phone.Show()
	fmt.Println("戴上壳的手机")
}
