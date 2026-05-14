package pgms

import "fmt"

func lowerBoundV1(arr []int, target int) int {
	for i, val := range arr {
		if val >= target {
			return i
		}
	}
	return len(arr)
}

func lowerBoundV2(arr []int, target int) int {
	var res int
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right - left) / 2
		if arr[mid] >= target {
			res = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return res
}

func Program025() {
	arr := []int{2, 3, 7, 10, 11, 11, 25}
	fmt.Println(lowerBoundV2(arr, 9))
}