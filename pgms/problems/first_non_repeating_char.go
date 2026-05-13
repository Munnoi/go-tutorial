// First non-repeating character
package pgms

import "fmt"

func firstNonRepeatingCharV1(s string) string {
	charCount := make(map[rune]int)

	for _, char := range s {
		charCount[char]++
	}

	for _, char := range s {
		if charCount[char] == 1 {
			return string(char)
		}
	}

	return "$"
}

func Program018() {
	s := "geeksforgeeks"
	fmt.Println(firstNonRepeatingCharV1(s))
}