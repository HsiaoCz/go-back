package main

import "fmt"

// 冒泡排序
// 冒泡排序，比较相邻的两个数谁大谁小
// 如果按照从小到大排序，那么，如果前面的比后面的大，就交换两个数的位置
// 反之则不变
// 重复这种比较方式，直到有序

type uint64Slice []uint64

func (us uint64Slice) swap(i, j int) {
	us[i], us[j] = us[j], us[i]
}

func sortBubble(numbers uint64Slice) {
	length := len(numbers)
	if length == 0 {
		return
	}
	flag := true

	for i := 0; i < length && flag; i++ {
		flag = false
		for j := length - 1; j > i; j-- {
			if numbers[j-1] > numbers[j] {
				numbers.swap(j-1, j)
				flag = true // 有数据才交换
			}
		}
	}
}

func main() {
	numbers := []uint64{5, 4, 2, 3, 8}
	sortBubble(numbers)
	fmt.Println(numbers)
}
