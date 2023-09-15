package main

import "fmt"

func main() {
	maps := make(map[string]struct{})
	maps["hello"] = struct{}{}
	_, ok := maps["hello"]
	fmt.Println(ok)
}
