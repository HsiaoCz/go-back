package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
)

func main() {
	m := make(map[string]string)
	m["zhangsan"] = "lisi"
	m["wangwu"] = "zhaoliu"
	for k, v := range m {
		fmt.Println("key:", k)
		fmt.Println("value:", v)
	}
	andi := NewAndi[string]()
	andi.InAndi("hello")
	andi.InAndi("hi")
	andi.InAndi("anddd")
	andi.InAndi("makesss")
	val, err := andi.OutAndi()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("andi value:", val)
}

// 类似于栈，先进后出结构
type Andi[T any] struct {
	data  []T
	lengh int
	l     sync.Mutex
}

func NewAndi[T any]() *Andi[T] {
	return &Andi[T]{
		data:  make([]T, 0),
		lengh: 0,
	}
}

func (a *Andi[T]) InAndi(v T) {
	a.l.Lock()
	defer a.l.Unlock()
	a.data = append(a.data, v)
}

func (a *Andi[T]) OutAndi() (v T, err error) {
	a.l.Lock()
	defer a.l.Unlock()

	a.lengh = len(a.data)
	if a.lengh == 0 {
		return v, errors.New("数据为空")
	}
	v = a.data[0]
	a.data = a.data[1:]
	return v, nil
}
