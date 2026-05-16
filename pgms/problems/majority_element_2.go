package pgms

import "fmt"

func majorityElement2V1(arr []int) []int {
	n := len(arr)

	candidate1, candidate2 := -1, -1
	count1, count2 := 0, 0

	for _, num := range arr {
		if num == candidate1 {
			count1++
		} else if num == candidate2 {
			count2++
		} else if count1 == 0 {
			candidate1 = num
			count1 = 1
		} else if count2 == 0 {
			candidate2 = num
			count2 = 1
		} else {
			count1--
			count2--
		}
	}

	count1, count2 = 0, 0
	for _, num := range arr {
		if num == candidate1 {
			count1++
		} else if num == candidate2 {
			count2++
		}
	}

	result := []int{}

	if count1 > n/3 {
		result = append(result, candidate1)
	} else if count2 > n/3 {
		result = append(result, candidate2)
	}

	return result 
}

func Program030() {
	arr := []int{3, 2, 3, 2, 2, 1, 1}
	fmt.Println(majorityElement2V1(arr))
}