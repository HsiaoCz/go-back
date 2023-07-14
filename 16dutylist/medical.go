package main

import "fmt"

// 具体的处理者之抓药师傅

type Medical struct {
	next Department
}

func (m *Medical) Execute(p *Patient) {
	if p.MedicineDone {
		fmt.Println("patient already given to patient")
		m.next.Execute(p)
		return
	}

	fmt.Println("Medical giving medical to patient")
	p.MedicineDone = true
	m.next.Execute(p)
}

func (m *Medical) SetNext(next Department) {
	m.next = next
}
