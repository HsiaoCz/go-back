package main

import "fmt"

// 适配器模式
// 将一个类的接口动态的适配成客户希望的另外一个接口
// 使得原本不能兼容的接口可以一起工作

type V5 interface {
	Use5V()
}

// 业务类
type Phone struct {
	v V5
}

func NewPhone(v V5) *Phone {
	return &Phone{
		v: v,
	}
}

func (p *Phone) Charge() {
	fmt.Println("对phone进行充电...")
	p.v.Use5V()
}

// 被适配的角色
type V220 struct{}

func (v *V220) Use220V() {
	fmt.Println("使用220v的电压进行充电...")
}

// 电源适配器
type Adapter struct {
	v220 *V220
}

func (a *Adapter) Use5V() {
	fmt.Println("使用电源适配进行充电...")
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
