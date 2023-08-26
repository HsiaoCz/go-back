package fb

func FeB(num int) int {
	if num == 1 {
		return 1
	}
	if num == 2 {
		return 1
	}
	return FeB(num-1) + FeB(num-2)
}
