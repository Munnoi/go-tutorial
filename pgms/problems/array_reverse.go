// Reverse an array
package pgms

import "fmt"

func reverseArray_v1(arr []int) {
	left := 0
	right := len(arr) - 1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}

func reverseArray_v2(arr []int) {
	n := len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	}
}


func Program001() {
	arr := []int{1, 2, 3, 4, 5}
	reverseArray_v2(arr)
	fmt.Println(arr)
}