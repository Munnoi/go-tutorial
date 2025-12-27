/*
Printing & reading values to and from the console.
*/
package pgms

/*
Importing the package 'fmt'.
'fmt' package implements the formatted I/O operations.
*/
import (
	"fmt"
)

func Hello() {
	/*
		Printing the string "Hello World!" to the console.
		`Println()` prints the value inside + '\n' (newline character).
	*/
	fmt.Println("Hello World!")
	/*
		Printing the string "Hello, World!" to the console.
		`Print()` prints the value inside without '\n'.
	*/
	fmt.Print("Hello, ")
	fmt.Print("World!\n")

	name := "Bob" // This is a method for declaring & initializing in go.
	/*
		Printing the string "Hello, World!" to the console.
		`Printf()` prints the value inside formatted  without '\n'.
		There are various format specifiers in go like '%s' for strings.
		Look at the official docs for more info.
	*/
	fmt.Printf("Hello, %s\n", name)
	age := 23
	fmt.Printf("Age: %d\n", age) // '%d' is for int values.

	var n int
	/*
		`Scanln()` reads user input from the console until a '\n' (newline char) is encountered.
		Stores successive space-seperated values into successive args.
	*/
	fmt.Print("Enter a number: ")
	fmt.Scanln(&n)
	fmt.Println("Num: ", n)

	var id int
	var username string
	/*
		`Scanf()` reads user input from the console (similar to scanf() fn in C language).
		Stores successive space-seperated values into successive args.
	*/
	fmt.Print("Enter id and username: ")
	fmt.Scanf("%d %s", &id, &username)
	fmt.Println("Id: ", id)
	fmt.Println("Username: ", username)

	var num int // This is a method for declaring a variable in go.
	/*
		`Scan()` reads user input from the console.
		Stores successive space-seperated values into successive args.
	*/
	fmt.Print("Enter another number: ")
	fmt.Scan(&num)
	fmt.Println("Num: ", num)

	// TODO: Implement `Sprint`, `Sprintln`, `Sprintf` functions.
	// TODO: Implement `Sscan`, `Sscanln`, `Sscanf` functions.
}
