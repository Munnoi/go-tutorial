/*
Slices in go.
*/
package pgms

import "fmt"

func Slice() {
	/*
		Slices are similar to arrays.
		However, unlike arrays,
		the length of a slice can grow and shrink as you see fit.
		Syntax -> `<name> := array[start:end]`.
	*/
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	mySlice := arr[2:4] // Slices the arr to include only values from index 2-3.
	fmt.Println(mySlice)
	/*
		`len()` - returns the length of the slice (the number of elements in the slice).
		`cap()` - returns the capacity of the slice (the number of elements the slice can grow or shrink to).
	*/
	fmt.Println(len(mySlice)) // [3, 4].
	fmt.Println(cap(mySlice)) // [3, 4, 5].

	otherSlice := make([]int, 5, 10) // Syntax -> `make([]dt, length, capacity)`.
	fmt.Println(otherSlice)
	fmt.Println(len(otherSlice))
	fmt.Println(cap(otherSlice))

	prices := []int{10, 20, 30}
	prices[2] = 50 // 30 changes to 50.
	fmt.Println(prices)

	/*
		`append()` appends to the end of the slice.
		Syntax -> `<name> = append(<name>, value1, value2, ... )`.
	*/
	prices = append(prices, 70, 80, 90)
	fmt.Println(prices) // [10, 20, 50, 70, 80, 90].

	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{6, 7, 8, 9, 10}

	slice := append(arr1, arr2...) // Appending two slices.
	fmt.Println(slice) // [1, 2, 3, 4, 5, 6, 7, 8, 9, 10].

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	dest := make([]int, 5)
	/*
		Copies a slice from one to another.
		Syntax -> `copy(dest, src)`; dest - destination, src - source.
	*/
	copy(dest, numbers)
	fmt.Println(numbers)
	fmt.Println(dest)
}
