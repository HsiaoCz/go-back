package main

import "fmt"

// 模板方法模式
// 模板方法模式是一种代码复用技术
// 它提取了类库中的公共行为，将公共行为放在父类中
// 而通过子类来实现不同的行为，它鼓励我们恰当的使用继承来实现代码的复用技术

type Beverage interface {
	BoilWater()
	Brew()
	PourInCup()
	AddThings()
	WantAddThings() bool
}

// 封装一套流程模板
type template struct {
	b Beverage
}

// 封装固定的模板
func (t *template) MakeBeverage() {
	if t == nil {
		return
	}

	t.b.BoilWater()
	t.b.Brew()
	t.b.PourInCup()

	if t.b.WantAddThings() {
		t.b.AddThings()
	}
}

// 具体的模板子类，制作咖啡
type MakeCoffee struct {
	template //继承模板
}

func NewMakeCoffee() *MakeCoffee {
	makeCoffee := new(MakeCoffee)
	// b 为Beverage，是makeCoffee的接口，这里需要给接口复制，指向具体的子类对象
	// 来触发b全部的接口方法的多态特性
	makeCoffee.b = makeCoffee
	return makeCoffee
}

func (mc *MakeCoffee) BoilWater()          { fmt.Println("将水煮到100摄氏度") }
func (mc *MakeCoffee) Brew()               { fmt.Println("用水冲咖啡豆...") }
func (mc *MakeCoffee) PourInCup()          { fmt.Println("将充好的咖啡到如陶瓷杯中...") }
func (mc *MakeCoffee) AddThings()          { fmt.Println("添加牛奶和糖...") }
func (mc *MakeCoffee) WantAddThings() bool { return true }
