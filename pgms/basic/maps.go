/*
Maps in go.
*/
package pgms

import "fmt"

func Maps() {
	// A map is a collection of unordered pairs of key-value.
	// It is also called a hash map or dictionary.

	// Create a map using make.
	var m1 = make(map[string]int)
	m1["jose"] = 1
	m1["sam"] = 2

	fmt.Println("m1 :", m1)

	// Declare and initialize.
	m2 := map[string]int{
		"apple":  1,
		"orange": 2,
	}
	fmt.Println("m2 :", m2)

	// Add a new element.
	m2["banana"] = 3
	fmt.Println("m2 after adding banana:", m2)

	// Delete an element.
	delete(m2, "apple")
	fmt.Println("m2 after deleting apple:", m2)

	// Iterate over map.
	for key, value := range m2 {
		fmt.Printf("key : %v, value : %v\n", key, value)
	}

	// Check if key exists.
	val, ok := m2["apple"]
	fmt.Println("val :", val, "present :", ok)

	val1, ok1 := m2["orange"]
	fmt.Println("val :", val1, "present :", ok1)
}
