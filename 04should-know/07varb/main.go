package main

import "fmt"

type Person struct {
	Username string
	Age      int
	Gender   string
}

func (p *Person) Hello() {
	fmt.Printf("%s-%s", p.Username, p.Gender)
}

func main() {
	var c = 12
	fmt.Println(c)
	p := &Person{
		Username: "Hello",
		Age:      12,
		Gender:   "男",
	}
	p.Hello()
	fmt.Println(p)
}
