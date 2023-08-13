package main

import "fmt"

// go的泛型
// 类型形参
// 类型约束
// 类型参数列表

type MySlice[T int | int64 | string] []T

// 这里T被称为类型形参
// int|int64|string 被称为类型约束
// [] 这个中括号里面的被称为类型参数列表

// 定义好泛型以后，并不能直接使用
// 需要先将泛型实例化之后才能使用

// 定义一个泛型函数
func Add[T int | int64 | float32 | string](a, b T) T {
	return a + b
}

func main() {
	var a MySlice[string] = []string{"hello", "hi"}
	fmt.Println(a)

	s1 := "hello"
	s2 := "!hi"
	s3 := Add[string](s1, s2)
	fmt.Println(s3)
}
