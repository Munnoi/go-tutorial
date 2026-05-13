package pgms

import "fmt"

func largestElemV1(arr []int) int {
	max := arr[0]
	
	for _, elem := range arr {
		if elem > max {
			max = elem
		}
	}

	return max
}

func Program021() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(largestElemV1(arr))
}