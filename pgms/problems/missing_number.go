// Missing number
package pgms

import "fmt"

func missingNumberV1(arr []int) int {
	totalSum, arrSum := 0, 0
	for i := 1; i <= len(arr)+1; i++ {
		totalSum += i
	}
	for _, val := range arr {
		arrSum += val
	}
	return totalSum - arrSum
}

func missingNumberV2(arr []int) int {
	xor1, xor2 := 0, 0
	for _, val := range arr {
		xor1 ^= val
	}
	for i := 1; i <= len(arr)+1; i++ {
		xor2 ^= i
	}

	return xor1 ^ xor2
}

func Program007() {
	arr := []int{8, 2, 4, 5, 3, 7, 1}
	fmt.Println(missingNumberV2(arr))
}