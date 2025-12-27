/*
Operators in go.
*/
package pgms

import "fmt"

func Operators() {
	var a = 15 + 25 // Addition.
	fmt.Println(a)

	var b = 15 - 25 // Subtraction.
	fmt.Println(b)

	var c = 15 * 25 // Multiplication.
	fmt.Println(c)

	var d = 25 / 15 // Division.
	fmt.Println(d)

	var e = 25 % 15 // Modulus.
	fmt.Println(e)

	// Assignment operators
	var x = 10
	x += 5 // x = x + 5
	fmt.Println(x)

	x -= 3 // x = x - 3
	fmt.Println(x)

	x *= 2 // x = x * 2
	fmt.Println(x)

	x /= 4 // x = x / 4
	fmt.Println(x)

	x %= 3 // x = x % 3
	fmt.Println(x)

	// Comparison operators
	var p = 10
	var q = 20
	fmt.Println(p == q) // Equal to
	fmt.Println(p != q) // Not equal to
	fmt.Println(p > q)  // Greater than
	fmt.Println(p < q)  // Less than
	fmt.Println(p >= q) // Greater than or equal to
	fmt.Println(p <= q) // Less than or equal to

	// Logical operators
	var isTrue = true
	var isFalse = false
	fmt.Println(isTrue && isFalse) // Logical AND
	fmt.Println(isTrue || isFalse) // Logical OR
	fmt.Println(!isTrue)           // Logical NOT

	// Bitwise operators
	var num1 uint = 60 // 0011 1100
	var num2 uint = 13 // 0000 1101

	fmt.Println(num1 & num2) // Bitwise AND (0000 1100 = 12)
	fmt.Println(num1 | num2) // Bitwise OR (0011 1101 = 61)
	fmt.Println(num1 ^ num2) // Bitwise XOR (0011 0001 = 49)
	fmt.Println(num1 << 2)   // Left Shift (1111 0000 = 240)
	fmt.Println(num1 >> 2)   // Right Shift (0000 1111 = 15)
	fmt.Println(^num1)       // Bitwise NOT (1100 0011 = -61)
}
