## 反射

### 1.什么是反射?

所谓的反射，就是程序在运行期间，动态的访问，检测和修改自身状态或者行为的一种能力
简单说：反射就是在程序运行时能够“观察”并且修改自己的行为

反射的本质是在程序运行期间探知对象的类型信息和内存结构
不用反射也可以实现，可以使用汇编来获取信息

### 2、go语言反射的三大法则

三大法则分别为:
1. 从interface{}变量可以反射出反射对象
2. 从反射对象可以获取interface{}变量
3. 要修改反射对象，其值必须可设置

#### 2.1、第一法则

反射的第一法则是能将go语言的interface{}变量转换成反射对象

为什么是从 interface{} 变量到反射对象？当我们执行 reflect.ValueOf(1) 时，虽然看起来是获取了基本类型 int 对应的反射类型，但是由于 reflect.TypeOf、reflect.ValueOf 两个方法的入参都是 interface{} 类型，所以在方法执行的过程中发生了类型转换

go语言的函数调用都是值传递，所以变量会在函数调用时进行类型转换
基本类型会被转换成interface{}类型

**接口到反射对象**

```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	author := "draven"
	fmt.Println("TypeOf author:", reflect.TypeOf(author))
	fmt.Println("ValueOf author:", reflect.ValueOf(author))
}

$ go run main.go
TypeOf author: string
ValueOf author: draven
```