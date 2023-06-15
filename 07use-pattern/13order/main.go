package main

import "fmt"

// 命令模式
// 命令模式将一个请求封装成一个对象，从而让我们可以使用不同的请求
// 对客户端进行参数化

// 这里举个例子，假如路边烤串
type Cooker struct{}

func (c *Cooker) MakeChuaner() { fmt.Println("烤羊肉串") }
func (c *Cooker) MakeJiChi()   { fmt.Println("烤鸡翅") }

// 抽象的命令

type Command interface {
	Make()
}

type CommandMakeChuan struct {
	cooker *Cooker
}

func (cc *CommandMakeChuan) Make() {
	cc.cooker.MakeChuaner()
}

type CommandMakeJiChi struct {
	cooker *Cooker
}

func (cc *CommandMakeJiChi) Make() {
	cc.cooker.MakeJiChi()
}

// 服务员，命令的收集者
type Water struct {
	Cmdlist []Command
}

func (w Water) Notify() {
	if w.Cmdlist == nil {
		return
	}

	for _, cmd := range w.Cmdlist {
		cmd.Make()
	}
}

func main() {
	cooker := new(Cooker)
	cmdChicken := CommandMakeJiChi{cooker: cooker}
	cmdChuaner := CommandMakeChuan{cooker: cooker}

	waiter := new(Water)
	waiter.Cmdlist = append(waiter.Cmdlist, &cmdChicken)
	waiter.Cmdlist = append(waiter.Cmdlist, &cmdChuaner)

	waiter.Notify()
}
