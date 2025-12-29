/*
Conditions in go.
*/
package pgms

import "fmt"

func Conditions() {
	/*
		Syntax:
		if condition {
			// code to be executed if condition is true
		}
	*/

	if 20 > 18 {
		fmt.Println("20 is greater than 18")
	}

	a := 20
	b := 18
	if a > b {
		fmt.Printf("%d is greater than %d\n", a, b)
	}

	/*
	 	Syntax:
		if condition {
			// code to be executed if condition is true
		} else {
			// code to be executed if condition is false
		}
	*/

	if a < b {
		fmt.Printf("%d is greater than %d\n", b, a)
	} else {
		fmt.Printf("%d is greater than %d\n", a, b)
	}

	/*
		Syntax:
		if condition1 {
			// code to be executed if condition1 is true
		} else if condition2 {
			// code to be executed if condition1 is false and condition2 is true
		} else {
			// code to be executed if both condition1 and condition2 are false
		}
	*/

	time := 22
	if time < 10 {
		fmt.Println("Good morning.")
	} else if time < 20 {
		fmt.Println("Good day.")
	} else {
		fmt.Println("Good evening.")
	}

	// Nested if-else.
	num := 10
	if num >= 0 {
		if num % 2 == 0 {
			fmt.Println("Positive and Even")
		} else {
			fmt.Println("Positive and Odd")
		}
	} else {
		fmt.Println("Negative")
	}

	// if with a short statement.
	// The variable `x` is only accessible within the if-else block.
	if x := 10; x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is not greater than 5")
	}
	// fmt.Println(x) // This would cause a compile-time error: undefined: x.
	
}
