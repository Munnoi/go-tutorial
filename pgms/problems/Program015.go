// Remove all occurrences of a character in a string
package pgms

import (
	"fmt"
	"strings"
)

func removeOccurencesV1(s string, c rune) string {
	result := ""
	for _, char := range s {
		if char != c {
			result += string(char)
		}
	}
	return result
}

func removeOccurencesV2(s string, c rune) string {
	result := strings.ReplaceAll(s, string(c), "")
	return result
}

func Program015() {
	s := "geeksforgeeks"
	c := 'g'

	fmt.Println(removeOccurencesV2(s, c))
}