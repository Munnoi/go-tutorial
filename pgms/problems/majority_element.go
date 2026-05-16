package pgms

import "fmt"

func majorityElementV1(arr []int) int {
	n := len(arr)

	for i := 0; i < n; i++ {
		count := 0
		for j := 0; j < n; j++ {
			if arr[i] == arr[j] {
				count++
			}
		}

		if count >= n/2 {
			return arr[i]
		}
	}
	return -1
}

func majorityElementV2(arr []int) int {
	n := len(arr)

	candidate := -1
	count := 0

	for _, num := range arr {
		if count == 0 {
			candidate = num
			count = 1
		} else if num == candidate {
			count++
		} else {
			count--
		}
	}

	count = 0
	for _, num := range arr {
		if num == candidate {
			count++
		}
	}

	if count > n/2 {
		return candidate
	}

	return -1
}

func Program029() {
	arr := []int{1, 1, 2, 1, 3, 5, 1}
	fmt.Println(majorityElementV2(arr))
}
