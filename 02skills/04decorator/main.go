package main

import "fmt"

// 装饰器模式
// 动态的给一个对象增加一些额外的职责
// 这里拿手机举个例子，给一台手机贴膜和戴壳

type Phone interface {
	Show()
}

// 装饰器的基础类
// 这里为结构体是因为go语言中interface不可以有成员属性

type Decorator struct {
	phone Phone
}

func (d *Decorator) Show() {}

// 具体的手机
type XiaoMi struct{}

func (x *XiaoMi) Show() { fmt.Println("秀出了小米手机...") }

type HuaWei struct{}

func (h *HuaWei) Show() { fmt.Println("秀出了华为手机...") }

// 具体的装饰器类
type MoDecorator struct {
	Decorator // 继承基础的装饰器类(主要继承phone的成员属性)
}

func NewMoDecorator(phone Phone) Phone {
	return &MoDecorator{Decorator{phone: phone}}
}

func (md *MoDecorator) Show() {
	md.phone.Show()
	fmt.Println("戴上壳的手机...")
}

type KeDecorator struct {
	Decorator
}

func NewKeDecorator(phone Phone) Phone {
	return &KeDecorator{Decorator{phone: phone}}
}

func (kd *KeDecorator) Show() {
	kd.phone.Show()
	fmt.Println("贴上了膜的手机...")
}
