package main

import "fmt"

// 适配器模式
// 将一个类的接口转换成客户端希望的另外一个接口
// 使得原本由于接口不兼容的而不能一起工作的那些类可以一起工作

// 这里举个iphone充电器的例子

type V5 interface {
	Use5V()
}

// 业务类，依赖5v的接口
type Phone struct {
	v V5
}

func NewPhone(v V5) *Phone {
	return &Phone{v: v}
}

func (p *Phone) Charge() {
	fmt.Println("对Phone进行充电...")
	p.v.Use5V()
}

type V220 struct{}

func (v *V220) Use220V() {
	fmt.Println("使用220v的电压进行充电...")
}

// 电源适配器
type Adapter struct {
	v220 *V220
}

func (a *Adapter) Use5V() {
	fmt.Println("使用电源适配器进行充电")
	// 调用适配者的方法
	a.v220.Use220V()
}

func NewAdapter(v220 *V220) *Adapter {
	return &Adapter{v220: v220}
}

func main() {
	iphone := NewPhone(NewAdapter(new(V220)))
	iphone.Charge()
}
