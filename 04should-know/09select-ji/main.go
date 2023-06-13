package main

import "fmt"

// 使用select 实现打印10以内的奇数

func SelectJi() {
	ch := make(chan int, 1)
	for i := 1; i <= 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}

func SelectOu() {
	ch := make(chan int, 1)
	for i := 2; i <= 11; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}

func main() {
	SelectJi()
	SelectOu()
}
