/*
Constants in go.
*/
package pgms

import "fmt"

func Constants() {
	/*
		Constants in go syntax -> `const <constant_name> datatype = value`.
		Constants names are usually written in uppercase letters.
	*/
	const PI float32 = 3.14 // Typed: Here the dt is provided.
	const N = 10            // Untyped: Here the dt is not provided, it's inferred.
	fmt.Println(PI)
	fmt.Println(N)
	// Block type constants
	const (
		A = 10
		B = 20
		C = "Hi"
	)
	fmt.Println(A, B, C)
}
