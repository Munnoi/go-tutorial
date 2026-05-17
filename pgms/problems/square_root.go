package pgms

import "fmt"

func squareRootV1(n int) int {
	i := 1

	for i*i <= n {
		i++
	}
	return i - 1
}

func squareRootV2(n int) int {
	left := 0
	right := n
	res := 0

	for left <= right {
		mid := (left + right) / 2
		square := mid * mid

		if square == n {
			return mid
		} else if square < n {
			res = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return res
}

func Program033() {
	n := 16
	fmt.Println(squareRootV2(n))
}