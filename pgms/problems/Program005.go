// Third largest element in an array
package pgms

import "fmt"

func thirdLargestV1(arr []int) (int, bool) {
	if len(arr) < 3 {
		return 0, false
	}

	first, second, third := -1<<63, -1<<63, -1<<63

	for _, v := range arr {
		if v > first {
			third = second
			second = first
			first = v
		} else if v > second && v != first {
			third = second
			second = v
		} else if v > third && v != second && v != first {
			third = v
		}
	}

	if third == -1<<63 {
		return 0, false
	}

	return third, true
}

func Program005() {
	arr := []int{12, 45, 7, 23, 89, 56}

	value, status := thirdLargestV1(arr)
	if status {
		fmt.Println("Third largest number: ", value)
	} else {
		fmt.Println("There is no third largest number")
	}
}
