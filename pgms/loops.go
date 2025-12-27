/*
Loops in go.
*/
package pgms

import "fmt"

func Loops() {
	/*
		For loop in Go.
		Go has only one looping construct, the `for` loop.
		Syntax:
		for initialization; condition; post-statement {
			// code to be executed
		}
	*/

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// `for` loop as a `while` loop.
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// Infinite loop.
	// k := 0
	// for {
	// 	fmt.Println(k)
	// 	k++
	// }

	// `for-range` loop.
	// Used to iterate over arrays, slices, strings, maps, and channels.
	fruits := []string{"apple", "banana", "cherry"}
	for index, value := range fruits {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}

	// You can omit the index or value if not needed.
	for _, value := range fruits {
		fmt.Printf("Value: %s\n", value)
	}

	for index := range fruits {
		fmt.Printf("Index: %d\n", index)
	}

	// `break` statement.
	// Used to terminate the loop.
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	// `continue` statement.
	// Used to skip the current iteration and proceed to the next.
	for i := 0; i < 10; i++ {
		if i%2 != 0 { // Skip odd numbers
			continue
		}
		fmt.Println(i)
	}

	// Nested loops.
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("(%d, %d) ", i, j)
		}
		fmt.Println()
	}

	// Labels with
	// `break` and `continue`.
	// Labels are used with `break` and `continue` to specify which loop to affect in nested loops.
OuterLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				break OuterLoop // Breaks out of the OuterLoop
			}
			fmt.Printf("%d %d\n", i, j)
		}
	}

	fmt.Println("---")

AnotherOuterLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				continue AnotherOuterLoop // Continues the AnotherOuterLoop
			}
			fmt.Printf("%d %d\n", i, j)
		}
	}
}
