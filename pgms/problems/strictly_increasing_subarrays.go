package pgms

import "fmt"

func strictlyIncreasingSubarraysV1(arr []int) int {
	n := len(arr)
	var count int
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if arr[j-1] >= arr[j] {
				break
			}
			count++
		}
	}

	return count
}

func strictlyIncreasingSubarraysV2(arr []int) int {
	n := len(arr)
	var count int
	len := 1

	for i := 1; i < n; i++ {
		if arr[i-1] < arr[i] {
			len++
		} else {
			count += len * (len - 1) / 2
			len = 1
		}
	}

	count += len * (len - 1) / 2

	return count
}

func Program032() {
	arr := []int{1, 4, 5, 3, 7, 9}
	fmt.Println(strictlyIncreasingSubarraysV2(arr))
}