package main

import "fmt"

// 策略模式

type WeaponStategy interface {
	UseWeapon()
}

// 具体的策略
type Ak47 struct{}

func (a *Ak47) UseWeapon() { fmt.Println("使用Ak47去战斗") }

type Knife struct{}

func (k *Knife) UseWeapon() { fmt.Println("使用小刀去战斗") }

// 环境类
type Hero struct {
	strategy WeaponStategy
}

func (h *Hero) SetWeaponStategy(s WeaponStategy) {
	h.strategy = s
}

func (h *Hero) Fight() {
	h.strategy.UseWeapon()
}
