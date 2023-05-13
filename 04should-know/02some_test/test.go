package main

import "fmt"

// 这里做一个测试
// 在java中,+号两边都是数值的时候，执行相加操作
// 如果两边有一边是字符串，那么执行字符串拼接
// 1+1+"hello"="2Hello"
// 这里试一下go的
// 这里go不允许这样做
// go的+号两边不允许有不同的类型

func main() {

	var s string
	s = "hello"
	fmt.Println(s)
}
