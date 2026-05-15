package pgms

import (
	"fmt"
	"sort"
)

func numberOfOccurrencesV1(arr []int, target int) int {
	var count int
	for _, val := range arr {
		if val == target {
			count++
		}
	}

	return count
}

func numberOfOccurrencesV2(arr []int, target int) int {
	lower := sort.Search(len(arr), func (i int) bool {
		return arr[i] >= target
	})

	upper := sort.Search(len(arr), func (i int) bool {
		return arr[i] > target
	})

	return upper - lower
}

func Program028() {
	arr := []int{1, 1, 2, 2, 2, 2, 3}
	target := 2
	fmt.Println(numberOfOccurrencesV2(arr, target))
}