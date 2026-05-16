package pgms

import "fmt"

func sortedSubsequenceOfSize3V1(arr []int) []int {
	n := len(arr)

	if n < 3 {
		return []int{}
	}

	smaller := make([]int, n)
	minIndex := 0
	smaller[0] = -1
	for i := 1; i < n; i++ {
		if arr[i] <= arr[minIndex] {
			minIndex = i
			smaller[i] = -1
		} else {
			smaller[i] = minIndex
		}
	}
	
	greater := make([]int, n)
	maxIndex := n - 1
	greater[n-1] = -1
	for i := n - 2; i >= 0; i-- {
		if arr[i] >= arr[maxIndex] {
			maxIndex = i
			greater[i] = -1
		} else {
			greater[i] = maxIndex
		}
	}

	for i := 0; i < n; i++ {
		if smaller[i] != -1 && greater[i] != -1 {
			return []int{arr[smaller[i]], arr[i], arr[greater[i]]}
		}
	}
	return []int{}
}

func Program031() {
	arr := []int{12, 11, 10, 5, 6, 2, 30}
	fmt.Println(sortedSubsequenceOfSize3V1(arr))
}