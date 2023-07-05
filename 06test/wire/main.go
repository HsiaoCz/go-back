package main

import "fmt"

// wire的使用
// 假如黑暗的世界里有一头怪兽
type Monster struct {
	Name string
}

func NewMonster() Monster {
	return Monster{Name: "kitty"}
}

// 假如黑暗的世界里还有一位勇士
type Player struct {
	Name string
}

func NewPlayer(name string) Player {
	return Player{Name: name}
}

// 很显然 勇士打死了恶龙
type Mission struct {
	Player  Player
	Monster Monster
}

func NewMission(p Player, m Monster) Mission {
	return Mission{Player: p, Monster: m}
}

func (m Mission) Start() {
	fmt.Printf("%s defeats %s,world peace!\n", m.Player.Name, m.Monster.Name)
}

func main() {
	monster := NewMonster()
	player := NewPlayer("jenger")
	mission := NewMission(player, monster)

	mission.Start()
}
