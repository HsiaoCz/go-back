package main

import "fmt"

// 策略模式
// 抽象的策略

type WeaponStrategy interface {
	UseWeapon()
}

// 具体的策略
type Ak47 struct{}

func (a *Ak47) UseWeapon() {
	fmt.Println("使用Ak进行战斗")
}

type Knife struct{}

func (k *Knife) UseWeapon() {
	fmt.Println("使用小刀进行战斗")
}

// 环境类
type Hero struct {
	ws WeaponStrategy
}

func (h *Hero) SetWeaponStategy(ws WeaponStrategy) {
	h.ws = ws
}

func (h *Hero) Fight() {
	h.ws.UseWeapon()
}
