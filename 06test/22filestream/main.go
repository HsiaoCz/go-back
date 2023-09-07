package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

// go 的文件操作

func ReadFileByBuffio(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		// 读取字符串,读到换行
		line, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			fmt.Println("文件读取完毕")
			break
		}
		fmt.Println(line)
	}
}

// ReadFile 方法可以读取全部的文件
func ReadFileAll(filename string) {
	b, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}

// 文件的写入
func WriteFileByBuffio(filename string) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("hello my man \n")
	}
	writer.Flush()
}

// 文件写入操作

func WriteFile(filename string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	str := "hello helo"
	file.Write([]byte(str)) // 写入字节
	file.WriteString(str)
}

// os.WriteFile可以直接写入文件
func WriteFileByOs(filename string) {
	err := os.WriteFile(filename, []byte("meizi"), 0666)
	if err != nil {
		log.Fatal(err)
	}
}
