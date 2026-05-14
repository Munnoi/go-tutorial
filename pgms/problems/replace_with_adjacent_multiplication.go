package pgms

import "fmt"

func replaceWithAdjacentMultiplicationV1(arr []int) {
	n := len(arr)
	temp := make([]int, n)
	var previous, next int
	for i := 0; i < n; i++ {
		if i == 0 {
			previous = 1
		} else {
			previous = arr[i-1]
		}
		if i == n-1 {
			next = 1
		} else {
			next = arr[i+1]
		}
		temp[i] = previous * arr[i] * next
	}
	copy(arr, temp)
}

func replaceWithAdjacentMultiplicationV2(arr []int) {
	n := len(arr)
	var curr, next int
	prev := 1
	for i := 0; i < n; i++ {
		curr = arr[i]
		if i == n - 1 {
			next = 1
		} else {
			next = arr[i+1]
		}

		arr[i] = prev * curr * next
		prev = curr
	}
}

func Program024() {
	arr := []int{2, 4, 5}
	replaceWithAdjacentMultiplicationV2(arr)
	fmt.Println(arr)
}