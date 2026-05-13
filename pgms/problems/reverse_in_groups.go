// Reverse in groups
package pgms

import "fmt"

func reverseInGroupsV1(arr []int, k int) {
	n := len(arr)
	for i := 0; i < n; i+=k {
		left := i
		right := i + k - 1
		if right >= n {
			right = n - 1
		}

		for left < right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}
}

func Program006() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	k := 4
	reverseInGroupsV1(arr, k)
	fmt.Println(arr)
}