package main

import "fmt"

// 适配器模式

// 适配的目标
type V5 interface {
	Use5V()
}

// 业务类，依赖5v接口
type Phone struct {
	v V5
}

func NewPhone(v V5) *Phone {
	return &Phone{v: v}
}

func (p *Phone) Charge() {
	fmt.Println("对Phone 进行充电...")
	p.v.Use5V()
}

// 被适配的角色
type V220 struct{}

func (v *V220) Use220V() {
	fmt.Println("使用220v的电压进行充电....")
}

// 电源适配器
type Adapter struct {
	v220 *V220
}

func (a *Adapter) Use5V() {
	fmt.Println("使用电源适配器进行充电...")
	// 调用适配者的方法
	a.v220.Use220V()
}

func NewAdapter(V220 *V220) *Adapter {
	return &Adapter{v220: V220}
}

func main() {
	ipone := NewPhone(NewAdapter(new(V220)))
	ipone.Charge()
}
