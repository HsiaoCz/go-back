package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
)

func main() {
	// 初始化切片容量的时候
	// 指定的初始值会成为切片的长度
	// 长度代表能存多少数据
	// 这里有一点需要注意
	// 如果写成s:=make([]string,0,10)
	// 通过s[0],s[1]这种赋值的形式添加数据是会出错的
	// 可以通过append追加的方式来添加
	s := make([]string, 2)
	s[0] = "hello"
	s[1] = "hi"
	for k, v := range s {
		fmt.Println(k, v)
	}

	ss := []string{"hello", "hi", "welcome", "and", "bullet"}
	value, sl := GetValue(ss)
	fmt.Println("value:", value)
	for _, v := range sl {
		fmt.Println("sl value:", v)
	}
	dq := NewDequeue()
	dq.InQueue("hello")
	dq.InQueue("hi")
	val, err := dq.OutQueue()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("dequeue:", val)
}

func GetValue(s []string) (string, []string) {
	return s[0], s[1:]
}

type Dequeue struct {
	data []any
	lock sync.Mutex
}

func NewDequeue() *Dequeue {
	return &Dequeue{
		data: make([]any, 0),
	}
}

// 这里在入队的时候应该设置最大不能超过多少
func (d *Dequeue) InQueue(v any) {
	d.lock.Lock()
	defer d.lock.Unlock()
	d.data = append(d.data, v)
}

// 出队的时候，应该先判断队列里是否有元素，没有就返回一个空队列的错误
func (d *Dequeue) OutQueue() (v any, err error) {
	d.lock.Lock()
	defer d.lock.Unlock()
	if len(d.data) == 0 {
		return nil, errors.New("队列为空")
	}
	v = d.data[0]
	d.data = d.data[1:]
	return v, nil
}
