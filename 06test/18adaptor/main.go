package main

import "fmt"

// 适配器模式

type V5 interface {
	UseV5()
}

// 被适配的角色
type IPhone struct {
	v V5
}

func NewIPone(v V5) *IPhone {
	return &IPhone{v: v}
}

func (i *IPhone) Charge(v V5) {
	fmt.Println("给iPhone充电")
	i.v.UseV5()
}

// 被适配的角色
type V220 struct{}

func (v *V220) Use220V() {
	fmt.Println("使用220v的电压进行充电")
}

// 电源适配器
type Adapter struct {
	v220 *V220
}

func (a *Adapter) UseV5() {
	fmt.Println("使用电源适配器进行充电")
	a.v220.Use220V()
}

func NewAdapter(v220 *V220) *Adapter {
	return &Adapter{v220: v220}
}
