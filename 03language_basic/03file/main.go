package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

// go 文件操作

func main() {
	// 打开文件
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
	fmt.Println("读取的文件的内容")
	fmt.Println(string(b[:n]))
}

// 使用bufio读取文件的内容
func ReadFileByBuffIO(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
				return
			}
		}
		fmt.Println(line)
	}
}

// ReadFile 方法可以读取全部文件
func ReadFileAll(filename string) {
	b, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}

// 文件操作
func WriteFile(filename string) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	str := "yizhi"
	file.Write([]byte(str))       // 写入字节
	file.WriteString("mammamamm") // 写入字符串
}

// 使用bufio的NewWriter
func WriteFileByBuffio(filename string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		write.WriteString("hahhahh ,mamm\n")
	}
	write.Flush()
}

// os.WriteFile
func WriteFileByOs(filename string) {
	err := os.WriteFile(filename, []byte("meizi"), 0666)
	if err != nil {
		log.Fatal(err)
	}
}
