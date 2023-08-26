package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go Hello()
	wg.Wait()
	fmt.Println("执行结束")
}

func Hello() {
	fmt.Println("hello")
	wg.Done()
	ch := Producker()
	res := Consumer(ch)
	fmt.Println(res)
}

// 单向通道
func Producker() chan int {
	ch := make(chan int, 2)
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 1 {
				ch <- i
			}
		}
		close(ch)
	}()
	return ch
}

func Consumer(ch chan int) int {
	sum := 0
	for v := range ch {
		sum += v
	}
	return sum
}

// 只读通道
func Producker2() <-chan int {
	ch := make(chan int, 2)
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 1 {
				ch <- i
			}
		}
		close(ch)
	}()
	return ch
}

func Consumer2(ch <-chan int) int {
	sum := 0
	for v := range ch {
		sum += v
	}
	return sum
}

// select

func Qi() {
	ch := make(chan int, 1)
	for i := 1; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}
