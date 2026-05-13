package pgms

import "fmt"

func rearrangeBySignV1(arr []int) []int {
	n := len(arr)

	res := make([]int, n)
	
	posIndex := 0
	negIndex := 1

	for _, val := range arr {
		if val > 0 {
			res[posIndex] = val
			posIndex += 2
		} else {
			res[negIndex] = val
			negIndex += 2
		}
	}

	return res
}

func Program020() {
	nums := []int{3, 1, -2, -5, 2, -4}
	fmt.Println(rearrangeBySignV1(nums))
}