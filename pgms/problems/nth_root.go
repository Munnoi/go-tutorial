package pgms

import "fmt"

func power(base int, exp int, limit int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
		if result > limit {
			return result
		}
	}
	return result
}

func nthRootV1(n int, m int) int {
	if m == 0 {
		return 0
	}

	if n == 1 {
		return m
	}

	left := 1
	right := m

	for left <= right {
		mid := (left + right) / 2
		pow := power(mid, n, m)
		if pow == m {
			return mid
		} else if pow < n {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func Program034(){
	n := 3
	m := 27
	fmt.Println(nthRootV1(n, m))
}