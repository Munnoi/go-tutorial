// Pick k such that the pair wise min diff is maximized
package pgms

import (
	"fmt"
	"sort"
)

func canPlace(arr []int, k int, mid int) bool {
	count := 1
	lastPlaced := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i]-lastPlaced >= mid {
			count++
			lastPlaced = arr[i]
		}
	}

	return count >= k
}

func pickK(arr []int, k int) int {
	sort.Ints(arr)

	low := 0
	high := arr[len(arr)-1] - arr[0]

	ans := 0

	for low <= high {
		mid := (low + high) / 2

		if canPlace(arr, k, mid) {
			ans = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return ans
}

func Program035() {
	arr := []int{1, 4, 9, 0, 2, 13, 3}
	k := 4
	fmt.Println(pickK(arr, k))
}