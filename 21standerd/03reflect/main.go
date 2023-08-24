package main

import (
	"fmt"
	"reflect"
)

type myInt int

// go 反射
// 反射是指在程序运行期间对程序本身进行访问和修改的能力
func main() {
	var i int32
	ReflectType(i)
	var s string = "hello"
	ReflectType(s)
	ReflectValue(s)

	var m myInt
	ReflectKindAndName(m)

	stu1 := student{
		Name:  "张三",
		Score: 123,
	}

	t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(), t.Kind())
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf(field.Name, field.Tag, field.Type)
	}
}

// 获取值的类型

func ReflectType(x any) {
	v := reflect.TypeOf(x)
	fmt.Printf("%v", v)
}

func ReflectKindAndName(x any) {
	t := reflect.TypeOf(x)
	fmt.Printf("type:%v kind:%v\n", t.Name(), t.Kind())
}

// 获取值
func ReflectValue(x any) {
	v := reflect.ValueOf(x)
	fmt.Printf("%v", v)
}

// 结构体反射
type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}
