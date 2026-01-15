package pgms

import (
	"strings"
	"fmt"
)

func contains(s string, ch rune) bool {
	return strings.ContainsRune(s, ch)
}

func pangramV1(s string) {
	for ch := 'a'; ch <= 'z'; ch++ {
		if !contains(s, ch) {
			fmt.Println("Not a pangram")
			return
		}
	}
	fmt.Println("Pangram")
}

func Program012() {
	s := "The quick brown fox jumps over the lazy dog"
	pangramV1(s)
}