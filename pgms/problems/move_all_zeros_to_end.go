// Move all zeros to the end
package pgms

import "fmt"

func moveZeroesV1(arr []int) {
	n := len(arr)
	j := 0
	for i := 0; i < n; i++ {
		if arr[i] != 0 {
			arr[j] = arr[i]
			j++
		}
	}

	for j < n {
		arr[j] = 0
		j++
	}

	fmt.Println(arr)
}

func moveZeroesV2(arr []int) {
	pos := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			arr[pos], arr[i] = arr[i], arr[pos]
			pos++
		}
	}
	fmt.Println(arr)
}

func Program011() {
	arr := []int{1, 2, 0, 4, 3, 0, 5, 0}
	moveZeroesV2(arr)
}