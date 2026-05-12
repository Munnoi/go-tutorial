// Toggle case; toggling the case of chars in a string
package pgms

import (
	"fmt"
	"strings"
	"unicode"
)

func toggleCaseV1(s string) string {
	result := ""
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			result += string(char - 32)
		} else if char >= 'A' && char <= 'Z' {
			result += string(char + 32)
		}
	}
	return result
}

func toggleCaseV2(s string) string {
	result := ""
	for _, char := range s {
		if unicode.IsUpper(char) {
			result += strings.ToLower(string(char))
		} else {
			result += strings.ToUpper(string(char))
		}
	}
	return result
}

func toggleCaseV3(s string) string {
	runes := []rune(s)

	for i, char := range runes {
		runes[i] = char ^ (1 << 5)
	}

	return string(runes)
}

func Program014() {
	s := "Hello World"
	fmt.Println(toggleCaseV3(s))
}