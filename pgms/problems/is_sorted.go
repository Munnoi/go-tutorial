package pgms

import (
	"fmt"
	"sort"
)

func isSortedV1(arr []int) bool {
	n := len(arr)
	for i := 1; i < n; i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

func isSortedV2(arr []int) bool {
	return sort.IntsAreSorted(arr)
}

func Program023() {
	arr := []int{10, 20, 30, 40, 50}
	fmt.Println(isSortedV2(arr))
}