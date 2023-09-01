package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	randen := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := randen.Intn(1000000)
	fmt.Println(num)
}
