package main

import "fmt"

// 命令模式
// 命令模式将一个请求封装成一个对象
// 从而让我们可以用不同的请求，对客户进行参数化
// 命令模式可以将请求发生着和接收者完全解耦，发送者和接收者之间没有直接引用关系

type Cooker struct{}

func (c *Cooker) MakeChuaner() { fmt.Println("烤串师傅烤了羊肉串....") }
func (c *Cooker) MakeJiChi()   { fmt.Println("烤串师傅烤了鸡翅") }

// 抽象的命令
type Command interface{ Make() }

type CommandCookChuaner struct {
	cooker *Cooker
}

func (cmd *CommandCookChuaner) Make() { cmd.cooker.MakeChuaner() }

type CommandCookChicken struct {
	cooker *Cooker
}

func (cmd *CommandCookChicken) Make() { cmd.cooker.MakeJiChi() }

// 服务员，命令的收集者
type Waiter struct {
	CmdList []Command
}

func (w Waiter) Notify() {
	if w.CmdList == nil {
		return
	}

	for _,cmd:= range w.CmdList{
		cmd.Make()
	}
}
