package newmath

//for 循环

func Recur(x int) int {
	if x <= 1 {
		return 1
	} else {
		return x * Recur(x-1)
	}
}
