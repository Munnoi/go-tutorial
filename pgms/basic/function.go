/*
Functions in go.
*/
package pgms

import "fmt"

func hello() {
	fmt.Println("Hello World!")
}

func greet(name string) {
	fmt.Println("Hello", name)
}

func sum(a int, b int) int {
	return a + b
}

func difference(a int, b int) (res int) {
	// Named return value.
	// The `res` variable is automatically initialized to its zero value (0 for int).
	// We can assign a value to `res` and then use a naked `return` statement.
	res = a - b
	return
}

func myFunc() (string, string) {
	return "Hello", "World" // Returning multiple values.
}

func Function() {
	/*
		Syntax:
		func FunctionName() {
			// code to be executed
		}
	*/
	hello() // Function calling.
	/*
	 	Syntax:
		func FunctionName(param1 type, param2 type, param3 type) {
			// code to be executed
		}
	*/
	greet("Bob")
	/*
		func FunctionName(param1 type, param2 type) type {
			// code to be executed
			return output
		}
	*/
	result := sum(10, 20)
	fmt.Println(result)
	res := difference(20, 10)
	fmt.Println(res)
	hlo, world := myFunc()
	fmt.Println(hlo, world)
	// Anonymous function (closure).
	// Can access variables from the surrounding scope.
	// They are often used for short, one-time operations or as callbacks.
	func(msg string) {
		fmt.Println(msg)
	}("I'm an anonymous function!")

	// Function assigned to a variable.
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println("Sum using anonymous function:", add(5, 7))

	// Higher-order function: a function that takes another function as an argument
	// or returns a function as a result.
	// Example: a function that applies an operation to a slice of integers.
	applyOperation := func(numbers []int, operation func(int) int) []int {
		result := make([]int, len(numbers))
		for i, num := range numbers {
			result[i] = operation(num)
		}
		return result
	}

	double := func(n int) int {
		return n * 2
	}

	numbers := []int{1, 2, 3, 4, 5}
	doubledNumbers := applyOperation(numbers, double)
	fmt.Println("Doubled numbers:", doubledNumbers)

	// Variadic functions: functions that can take a variable number of arguments.
	// The type of the variadic parameter is a slice.
	sumAll := func(nums ...int) int {
		total := 0
		for _, num := range nums {
			total += num
		}
		return total
	}

	fmt.Println("Sum of 1, 2, 3:", sumAll(1, 2, 3))
	fmt.Println("Sum of 10, 20, 30, 40:", sumAll(10, 20, 30, 40))
}
