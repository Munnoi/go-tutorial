// Second largest element in an array
package pgms

import (
	"fmt"
	"math"
	"sort"
)

func secondLargestV1(arr []int) int {
	if len(arr) < 2 {
		return -1
	}

	second := math.MinInt
	first := math.MinInt

	for _, v := range arr {
		if v > first {
			second = first
			first = v
		} else if v > second && v != first {
			second = v
		}
	}

	if second == math.MinInt {
		return -1
	}

	return second
}

func secondLargestV2(arr []int) int {
	if len(arr) < 2 {
		return -1
	}

	sort.Ints(arr)

	n := len(arr) - 1
	for i := n - 1; i >= 0; i-- {
		if arr[i] != arr[n] {
			return arr[i]
		}
	}

	return -1
}

func Program002() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(secondLargestV2(arr))
}

