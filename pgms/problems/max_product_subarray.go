package pgms

import "fmt"

func maxProductSubarray(arr []int) int {
	n := len(arr)
	maxProd := arr[0]

	for i := 0; i < n; i++ {
		mul := 1
		for j := i; j < n; j++ {
			mul *= arr[j]
			maxProd = max(maxProd, mul)
		}
	}

	return maxProd
}

func Program039() {
	arr := []int{-2, 6, -3, -10, 0, 2};
	result := maxProductSubarray(arr)
	fmt.Println(result)
}