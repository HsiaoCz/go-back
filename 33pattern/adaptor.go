package pattern

import "fmt"

// 适配器模式
// 将一个接口转换成客户希望的另一个接口
// 使得原本由于接口不兼容而不能一起工作的那些类可以一起工作

type V5 interface {
	Use5V()
}

// 业务类，依赖5v接口
type IPhone struct {
	v V5
}

func NewPhone(v V5) *IPhone {
	return &IPhone{v: v}
}

func (i *IPhone) Charge() {
	fmt.Println("对iphone进行充电...")
	i.v.Use5V()
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
