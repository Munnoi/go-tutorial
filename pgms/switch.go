package pgms

import "fmt"

func Switch() {
	/*
		Switch statement in Go.
		Syntax:
		switch expression {
		case value1:
			// code to be executed if expression == value1
		case value2:
			// code to be executed if expression == value2
		default:
			// code to be executed if expression is different from all cases
		}
	*/
	month := 5

	switch month {
	case 1:
		fmt.Println("January")
	case 2:
		fmt.Println("February")
	case 3:
		fmt.Println("March")
	case 4:
		fmt.Println("April")
	case 5:
		fmt.Println("May")
	case 6:
		fmt.Println("June")
	case 7:
		fmt.Println("July")
	case 8:
		fmt.Println("August")
	case 9:
		fmt.Println("September")
	case 10:
		fmt.Println("October")
	case 11:
		fmt.Println("November")
	case 12:
		fmt.Println("December")
	default:
		fmt.Println("Invalid month")
	}

	// Switch with multiple case values.
	newMonth := "February"
	switch newMonth {
	case "January", "February", "March":
		fmt.Println("First Quarter")
	case "April", "May", "June":
		fmt.Println("Second Quarter")
	default:
		fmt.Println("Other Quarter")
	}

	temperature := 25
	switch {
	case temperature < 0:
		fmt.Println("Freezing")
	case temperature >= 0 && temperature < 10:
		fmt.Println("Cold")
	case temperature >= 10 && temperature < 20:
		fmt.Println("Cool")
	case temperature >= 20 && temperature < 30:
		fmt.Println("Warm")
	default:
		fmt.Println("Hot")
	}

	// Fallthrough keyword
	// The `fallthrough` keyword is used to force the execution flow to fall through to the next case.
	// This is unlike other languages where cases fall through by default unless a `break` is used.
	// In Go, cases do not fall through by default.
	grade := "B"
	switch grade {
	case "A":
		fmt.Println("Excellent")
	case "B":
		fmt.Println("Good")
		fallthrough
	case "C":
		fmt.Println("Average")
	case "D":
		fmt.Println("Poor")
	default:
		fmt.Println("Invalid grade")
	}
}
