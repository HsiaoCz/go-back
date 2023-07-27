package main

import "fmt"

// 策略模式
type WeaponStategy interface {
	UseWeapon() //使用策略
}

// 具体的策略
type Ak47 struct{}

func (ak *Ak47) UseWeapon() { fmt.Println("使用AK47进行战斗") }

type Knife struct{}

func (k *Knife) UseWeapon() { fmt.Println("使用小刀进行战斗") }

// 环境类
type Hero struct {
	strategy WeaponStategy //环境类拥有抽象的策略
}

// 设置一个策略
func (h *Hero) SetWeaponStategy(s WeaponStategy) {
	h.strategy = s
}

func (h *Hero) Fight() {
	h.strategy.UseWeapon()
}
func main() {
	hero := new(Hero)
	hero.SetWeaponStategy(new(Ak47))
	hero.Fight()

	hero.SetWeaponStategy(new(Knife))
	hero.Fight()
}
