package pgms

import (
	"fmt"
	"math"
	"sort"
)

func minDiffPair(arr []int) int {
	sort.Ints(arr)

	minDiff := math.MaxInt32

	for i := 1; i < len(arr); i++ {
		diff := arr[i] - arr[i-1]
		if diff < minDiff {
			minDiff = diff
		}
	}

	return minDiff
}

func Program038() {
	arr := []int{1, 5, 3, 19, 18, 25}
	fmt.Println(minDiffPair(arr))
}