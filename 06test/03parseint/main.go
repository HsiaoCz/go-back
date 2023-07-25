package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	a := "+12345"
	b, err := strconv.Atoi(a)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(b)
}
