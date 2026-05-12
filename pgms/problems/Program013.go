// Two strings are equal or not
package pgms

import (
	"fmt"
)

func equalOrNot(s1, s2 string) bool {
	return s1 == s2
}

func Program013() {
	s1 := "abc"
	s2 := "abcd"

	if equalOrNot(s1, s2) {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not equal")
	}
}
