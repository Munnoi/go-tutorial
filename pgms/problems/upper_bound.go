package pgms

import "fmt"

func upperBoundV1(arr []int, target int) int {
	for i, val := range arr {
		if val > target {
			return i
		}
	}
	return len(arr)
}

func upperBoundV2(arr []int, target int) int {
	var res int
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right - left) / 2
		if arr[mid] > target {
			res = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return res
}

func Program026() {
	arr := []int{2, 3, 7, 10, 11, 11, 25}
	fmt.Println(upperBoundV2(arr, 11))
}