package main

// 插入排序
// 将一个待排序的序列的第一个元素看做一个有序序列
// 将后面的元素看作未排序的序列
// 如果从小到大排序
// 那么从第二个元素开始，依次取出，和前面的序列进行比较
// 如果比第一个大，放在右边，否则放在左边

import (
	"fmt"
)

type uint64Slice []uint64

func main() {
	numbers := []uint64{5, 4, 20, 3, 8, 2, 9}
	insertSort(numbers)
	fmt.Println(numbers)
}

func insertSort(numbers uint64Slice) {
	for i := 1; i < len(numbers); i++ {
		tmp := numbers[i]
		// 从待排序序列开始比较,找到比其小的数
		j := i
		for j > 0 && tmp < numbers[j-1] {
			numbers[j] = numbers[j-1]
			j--
		}
		// 存在比其小的数插入
		if j != i {
			numbers[j] = tmp
		}
	}
}
