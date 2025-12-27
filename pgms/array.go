/*
Arrays in go.
*/
package pgms

import "fmt"

func Array() {
	/*
		Arrays are used to store multiple values of the same type in a single variable,
		instead of declaring separate variables for each value.
		Syntax -> `var <name> = [<length>]dt{val1, val2, ... }`; similar for the shorthand syntax too.
	*/
	var arr = [4]int{1, 2, 3, 4} // Here the size of the array is 4 and it can store 4 integer values.
	fmt.Println(arr)

	list := [3]int{9, 8, 7} // This is the shorthand array initialization syntax.
	fmt.Println(list)

	var nums1 = [...]int{1, 2, 3, 4, 5} // Here the array length is inferred.
	fmt.Println(nums1)

	var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	fmt.Println(cars)

	var nums2 = [4]int{1: 10, 3: 20} // In this eg only index 1 & 3 gets initialized.
	fmt.Println(nums2)

	fmt.Println(len(nums2)) // `len()` calculates the length of the array (similar in Python!).
}
