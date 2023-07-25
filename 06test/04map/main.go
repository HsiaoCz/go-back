package main

import "fmt"

func SliceModify(s []int) {
	s[1] = 100
}

func MapModify(m map[string]int) {
	m["hello"] = 1222
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(s)
	SliceModify(s)
	fmt.Println(s)

	// map
	m := map[string]int{"hello": 345}
	fmt.Println(m)
	MapModify(m)
	fmt.Println(m)

}
