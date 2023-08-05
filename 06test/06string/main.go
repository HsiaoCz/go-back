package main

import "fmt"

// string 的驻留机制
// 很显然，go里面的string并不驻留
func main() {
	str := "hello"
	str1 := str
	fmt.Printf("%v\n", &str)
	fmt.Printf("%v\n", &str1)
}
