package main

import "fmt"

// 命令模式
// 将一个请求封装成为一个对象，从而让我们可以用不同的请求对客户
// 进行参数化
// 命令模式可以将请求的发送者和接收者完全解耦，发送者和接收者之间没有直接引用关系
// 发送请求的对象只需要知道如何发送请求，而不需要知道如何完成请求

type Cooker struct{}

func (c *Cooker) MakeChuaner() { fmt.Println("师傅烤了羊肉串...") }
func (c *Cooker) MakeJiChi()   { fmt.Println("师傅烤了鸡翅...") }

// 抽象的命令
type Command interface {
	Make()
}

type CommandCookChuaner struct {
	cooker *Cooker
}

func (ccc *CommandCookChuaner) Make() { ccc.cooker.MakeChuaner() }

type CommandCookChicken struct {
	cooker *Cooker
}

func (cck *CommandCookChicken) Make() { cck.cooker.MakeJiChi() }

// 服务员，命令的收集者
type Waiter struct {
	CmdList []Command
}

func (w *Waiter) Notify() {
	if w.CmdList == nil {
		return
	}

	for _, cmd := range w.CmdList {
		cmd.Make()
	}
}
func main() {
	cooker := new(Cooker)
	cmdChicken := CommandCookChicken{cooker: cooker}
	cmdChuaner := CommandCookChuaner{cooker: cooker}

	waiter := new(Waiter)
	waiter.CmdList = append(waiter.CmdList, &cmdChicken)
	waiter.CmdList = append(waiter.CmdList, &cmdChuaner)

	waiter.Notify()
}
