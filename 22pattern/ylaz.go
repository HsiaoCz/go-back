package main

import "fmt"

// 依赖倒转原则
// 张三开奔驰，李四开宝马
// 现在张三想开宝马，李四想开奔驰
// 并且现在又多了一辆车，多了一个司机
type Car interface {
	Run()
}

type Driver interface {
	Drive(car Car)
}

// 使用依赖倒转来实现
// 依赖倒转，依赖于抽象，而不是依赖于具体的类
type Benz struct{}

func (b *Benz) Run() { fmt.Println("奔驰 is running") }

type BWM struct{}

func (b *BWM) Run() { fmt.Println("宝马 is running") }

type Toyota struct{}

func (t *Toyota) Run() { fmt.Println("Toyota is running") }

// 具体的司机类
type Zhangs struct{}

func (z *Zhangs) Drive(car Car) {
	fmt.Println("张三正在开车")
	car.Run()
}

type Lis struct{}

func (l *Lis) Drive(car Car) {
	fmt.Println("李四正在开车")
	car.Run()
}

type Wangw struct{}

func (w *Wangw) Drive(car Car) {
	fmt.Println("王五正在开车")
	car.Run()
}
