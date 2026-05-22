package pgms

import (
	"fmt"
)

func sortArr(arr []int) {
	for i := 0; i < len(arr) - 1; i+=2 {
		arr[i], arr[i+1] = arr[i+1], arr[i]
	}
}

func Program036() {
	arr := []int{1, 2, 3, 4, 5}
	sortArr(arr)
	fmt.Println(arr)
}