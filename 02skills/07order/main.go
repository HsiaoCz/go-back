package mian

import "fmt"

// 命令模式
// 命令模式将一个请求封装成为一个对象，从而让我们可以用不同的请求
// 对客户进行参数化
// 命令模式可以将请求发生者和接收者完全解耦，发送者与接收者之间没有直接引用关系
// 发生请求的对象只需要知道如何发送请求，而不必知道如何完成请求

// 这里举个例子，假如路边烤串，有烤羊肉串，烤鸡肉串，有烤串师傅和服务员

type Cooker struct{}

func (c *Cooker) MakeChuaner() { fmt.Println("烤串师傅烤了羊肉串...") }
func (c *Cooker) MakeJiChi()   { fmt.Println("烤串师傅烤了鸡翅...") }

// 抽象的命令
type Command interface{ Make() }

type CommandCookChicken struct{ cooker *Cooker }

func (cmd *CommandCookChicken) Make() { cmd.cooker.MakeJiChi() }

type CommandCookChuaner struct{ cooker *Cooker }

func (cmd *CommandCookChuaner) Make() { cmd.cooker.MakeChuaner() }

// 服务员，命令的收集者
type Waiter struct {
	CmdList []Command
}

func (w Waiter) Notify() {
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
