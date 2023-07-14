package main

import "fmt"

// 具体的处理者
type Doctor struct {
	next Department
}

func (d *Doctor) Execute(p *Patient) {
	if p.DoctorCheckDone {
		fmt.Println("Doctor checkup already done")
		d.next.Execute(p)
		return
	}

	fmt.Println("Doctor checking patient")
	p.DoctorCheckDone = true
	d.next.Execute(p)
}

func (d *Doctor) SetNext(next Department) {
	d.next = next
}
