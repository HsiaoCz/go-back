package main

import "fmt"

// 反转一个字符串
// "hello" --->"olleh"

func devstr(str []byte) string {
	var str1 []byte
	for i := len(str) - 1; i >= 0; i-- {
		str1 = append(str1, str[i])
	}
	return string(str1)
}

func main() {
	str := "hello"
	str1 := devstr([]byte(str))
	fmt.Println(str1)
}
