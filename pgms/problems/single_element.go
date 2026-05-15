package pgms

import (
	"fmt"
)

func singleElementV1(arr []int) int {
	n := len(arr)

	for i := 0; i < n; i+=2 {
		if arr[i] != arr[i+1] {
			return arr[i]
		}
	}

	return arr[n-1]
}

func singleElementV2(arr []int) int {
	xor := 0
	for _, val := range arr {
		xor ^= val
	}

	return xor
}

func singleElementV3(arr []int) int {
	n := len(arr)
	left := 0
	right := n - 1

	for left < right {
		mid := left + (right - left) / 2

		if mid % 2 == 1 {
			mid--
		}

		if arr[mid] == arr[mid+1] {
			left = mid + 2
		} else {
			right = mid
		}
	}

	return arr[left]
}

func Program027() {
	arr := []int{1, 1, 2, 2, 3, 4, 4}
	fmt.Println(singleElementV3(arr))
}