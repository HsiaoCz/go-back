package main

import "fmt"

// 具体的处理者
// 前台
type Reception struct {
	next Department
}

func (r *Reception) Execute(p *Patient) {
	if p.RegisterionDone {
		fmt.Println("patient registration already done")
		r.next.Execute(p)
		return
	}

	fmt.Println("Reception registering patient")
	p.RegisterionDone = true
	r.next.Execute(p)
}

func (r *Reception) SetNext(next Department) {
	r.next = next
}
