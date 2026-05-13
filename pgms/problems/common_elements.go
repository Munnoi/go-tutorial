// Common elements
package pgms

import "fmt"

func commonElementsV1(arr1 []int, arr2 []int, arr3 []int) []int {
	mp := make(map[int]int)
	for _, val := range arr1 {
		mp[val] = 1
	}

	for _, val := range arr2 {
		_, ok := mp[val]
		if ok {
			mp[val] = 2
		}
	}

	for _, val := range arr3 {
		value, ok := mp[val]
		if ok && value == 2 {
			mp[val] = 3
		}
	}

	result := []int{}
	for key, val := range mp {
		if val == 3 {
			result = append(result, key)
		}
	}

	return result
}

func Program008() {
	arr1 := []int{1, 5, 10, 20, 30}
	arr2 := []int{5, 13, 15, 20}
	arr3 := []int{5, 20}

	fmt.Println(commonElementsV1(arr1, arr2, arr3))
}