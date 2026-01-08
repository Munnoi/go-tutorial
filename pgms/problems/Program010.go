// Remove duplicates from a sorted array
package pgms

import "fmt"

func removeDuplicatesV1(arr []int) {
	n := len(arr)
	if n == 0 {
		return
	}

	j := 0
	for i := 0; i < n-1; i++ {
		if arr[i] != arr[i+1] {
			arr[j] = arr[i]
			j++
		}
	}
	arr[j] = arr[n-1]
	j++

	fmt.Println(arr[:j])
}

func Program010() {
	arr := []int{1, 2, 3, 3, 4, 4, 4, 5, 5}
	removeDuplicatesV1(arr)
}