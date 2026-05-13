// Two strings are equal or not
package pgms

import (
	"fmt"
	"strings"
)

func equalOrNotV1(s1, s2 string) bool {
	return s1 == s2
}

func equalOrNotV2(s1, s2 string) bool {
	return strings.EqualFold(s1, s2)
}

func Program013() {
	s1 := "abc"
	s2 := "abc"

	if equalOrNotV2(s1, s2) {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not equal")
	}
}
