// Remove spaces from a string
package pgms

import (
	"fmt"
	"strings"
)

func removeSpacesV1(s string) string {
	result := ""
	for _, char := range s {
		if char != ' ' {
			result += string(char)
		}
	}
	return result
}

func removeSpacesV2(s string) string {
	return strings.ReplaceAll(s, " ", "")
}

func Program016() {
	s := "g  eeks   for ge  eeks"
	fmt.Println(removeSpacesV2(s))
}