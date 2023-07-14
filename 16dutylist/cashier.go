package main

import "fmt"

// 具体的处理者

type Cashier struct {
	next Department
}

func (c *Cashier) Execute(p *Patient) {
	if p.PaymentDone {
		fmt.Println("payment done")
		return
	}
	fmt.Println("all done")
}

func (c *Cashier) SetNext(next Department) {
	c.next = next
}
