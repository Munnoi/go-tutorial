/*
Variables in go.
*/
package pgms

import (
	"fmt"
)

func Variables() {
	/*
		There are different types of datatypes in go.
		int, float, string, bool
	*/

	var a int // Variable declaration. Syntax -> `var <variable_name> datatype`.
	a = 10    // Variable initialization.
	fmt.Println(a)

	/*
	 Variable declaration + initialization. Syntax -> `var <variable_name> datatype = <value>`.
	 Can be used everywhere.
	*/
	var b int = 11
	fmt.Println(b)

	/*
	 Variable declaration + initialization. Syntax -> `<variable_name> := <value>`.
	 Only be used to initialize local variables.
	*/
	c := 11 //
	fmt.Println(c)

	var num int        // Defaults to 0.
	var name string    // Defalts to "".
	var salary float32 // Defaults to 0.
	var isMarried bool // Defaults to false.
	fmt.Println(num, name, salary, isMarried)

	// Multiple variable declaration + initialization.
	var one, two, three int = 1, 2, 3
	four, five := 4, 5
	fmt.Println(one, two, three, four, five)

	// Block declaration + initialization.
	var (
		x int = 10
		y string = "Hello"
		z bool = true
	)
	fmt.Println(x, y, z)
}
