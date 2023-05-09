package main

import "fmt"

// 策略模式
// 这个举个例子
// 英雄使用武器
type WeaponStategy interface {
	UseWeapon() // 使用武器
}

// 具体的策略
type Ak47 struct{}

func (ak *Ak47) UseWeapon() { fmt.Println("使用Ak47去战斗...") }

type Knife struct{}

func (k *Knife) UseWeapon() { fmt.Println("使用小刀去战斗...") }

// 环境类
type Hero struct {
	strategy WeaponStategy // 环境类拥有抽象的策略
}

// 设置一个策略
func (h *Hero) SetWeaponStategy(s WeaponStategy) {
	h.strategy = s
}

func (h *Hero) Fight() {
	h.strategy.UseWeapon() // 在这里使用策略
}

func main() {
	hero := new(Hero)
	hero.SetWeaponStategy(new(Ak47))
	hero.Fight()

	hero.SetWeaponStategy(new(Knife))
	hero.Fight()
}
