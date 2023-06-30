package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	hello()
	dur := time.Since(start)
	fmt.Println(dur)
}

func hello() {
	for i := 0; i < 200; i++ {
		fmt.Println("hello")
	}
}
