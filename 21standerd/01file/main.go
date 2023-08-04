package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

// go的文件操作
func main() {
	file, err := os.Open("./1233.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	b := make([]byte, 1024)
	n, err := file.Read(b)
	if err == io.EOF {
		fmt.Println("文件读取完毕")
		return
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("读取的文件内容:")
	fmt.Println(string(b[:n]))
}

// 通过buffio读取文件

func ReadFileByBuffIO(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			fmt.Println("文件读取完毕")
			break
		}
		if err != nil {
			fmt.Println("read file failed,err:", err)
			return
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

// 文件的写入操作
func WriteFile(filename string) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	str := "yizhi"
	file.Write([]byte(str))    // 写入字节
	file.WriteString("hihiii") // 写入字符串

}

// 使用bufio的NewWriter
func WriteFileByBuffio(filename string) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("hellohello mand\n")
	}
	writer.Flush()
}

// os.WriteFile
func WriteFileByOs(filename string) {
	err := os.WriteFile(filename, []byte("meizi"), 0666)
	if err != nil {
		log.Fatal(err)
	}
}
