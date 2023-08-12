package main

import "fmt"

// 快速排序
type uint64Slice []uint64

func (us uint64Slice) swap(i, j int) {
	us[i], us[j] = us[j], us[i]
}

func quickSort(numbers uint64Slice, start, end int) {
	var middle int
	tempStart := start
	tempEnd := end

	if tempStart >= tempEnd {
		return
	}

	pivot := numbers[start]
	for start < end {
		for start < end && numbers[end] > pivot {
			end--
		}
		if start < end {
			numbers.swap(start, end)
			start++
		}
		for start < end && numbers[start] < pivot {
			start++
		}
		if start < end {
			numbers.swap(start, end)
			end--
		}
		numbers[start] = pivot
		middle = start

		quickSort(numbers, tempStart, middle-1)
		quickSort(numbers, middle+1, tempEnd)
	}
}

func main() {
	numbers := []uint64{5, 4, 20, 3, 8, 2, 8}
	quickSort(numbers, 0, len(numbers)-1)
	fmt.Println(numbers)
}
